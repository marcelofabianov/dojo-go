package handler

import (
	"encoding/json"
	"net/http"

	"github.com/marcelofabianov/dojo-go/internal/model"
	"github.com/marcelofabianov/dojo-go/internal/port"
	"github.com/marcelofabianov/dojo-go/pkg/validator"
	"github.com/marcelofabianov/dojo-go/pkg/web"
)

type CreateCourseRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateCourseResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type CreateCourseHandler struct {
	validator     *validator.Validator
	courseService port.CourseServicePort
}

func NewCreateCourseHandler(
	validator *validator.Validator,
	courseService port.CourseServicePort,
) *CreateCourseHandler {
	return &CreateCourseHandler{
		validator:     validator,
		courseService: courseService,
	}
}

func (h *CreateCourseHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := web.GetLogger(ctx)

	var req CreateCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("failed to decode request body", "error", err)
		web.ErrDecodeRequestBody(err, w, r)
		return
	}

	if err := h.validator.Validate(req); err != nil {
		logger.Warn("request validation failed", "error", err)
		web.Error(w, r, err)
		return
	}

	input := model.NewCourseInput{
		Title:       req.Title,
		Description: req.Description,
	}

	createdCourse, err := h.courseService.CreateCourse(ctx, input)
	if err != nil {
		logger.Error("failed to create course", "error", err)
		web.Error(w, r, err)
		return
	}

	response := CreateCourseResponse{
		ID:          createdCourse.ID,
		Title:       createdCourse.Title,
		Description: createdCourse.Description,
		CreatedAt:   createdCourse.CreatedAt.String(),
	}

	logger.Info("course created successfully", "course_id", response.ID)
	web.Success(w, r, http.StatusCreated, response)
}
