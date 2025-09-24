package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/marcelofabianov/fault"

	"github.com/marcelofabianov/dojo-go/internal/model"
	"github.com/marcelofabianov/dojo-go/internal/port"
	"github.com/marcelofabianov/dojo-go/pkg/validator"
	"github.com/marcelofabianov/dojo-go/pkg/web"
)

type UpdateCourseRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateCourseHandler struct {
	validator     *validator.Validator
	courseService port.CourseServicePort
}

func NewUpdateCourseHandler(validator *validator.Validator, courseService port.CourseServicePort) *UpdateCourseHandler {
	return &UpdateCourseHandler{
		validator:     validator,
		courseService: courseService,
	}
}

func (h *UpdateCourseHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := web.GetLogger(ctx)

	idStr := chi.URLParam(r, "id")
	if _, err := uuid.Parse(idStr); err != nil {
		logger.Warn("invalid uuid format in url param", "id", idStr, "error", err)
		web.Error(w, r, fault.New("invalid id format", fault.WithCode(fault.Invalid)))
		return
	}

	var req UpdateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("failed to decode request body", "error", err)
		web.ErrDecodeRequestBody(err, w, r)
		return
	}

	if err := h.validator.Validate(req); err != nil {
		web.Error(w, r, err)
		return
	}

	input := model.UpdateCourseInput{
		Title:       req.Title,
		Description: req.Description,
	}

	updatedCourse, err := h.courseService.UpdateCourse(ctx, idStr, input)
	if err != nil {
		if errors.Is(err, model.ErrCourseNotFound) {
			logger.Warn("course not found for update", "id", idStr)
			web.Error(w, r, fault.New("course not found", fault.WithCode(fault.NotFound)))
			return
		}

		logger.Error("failed to update course", "id", idStr, "error", err)
		web.Error(w, r, err)
		return
	}

	response := CreateCourseResponse{
		ID:          updatedCourse.ID,
		Title:       updatedCourse.Title,
		Description: updatedCourse.Description,
		CreatedAt:   updatedCourse.CreatedAt.String(),
	}

	logger.Info("course updated successfully", "course_id", response.ID)
	web.Success(w, r, http.StatusOK, response)
}
