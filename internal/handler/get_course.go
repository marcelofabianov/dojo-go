package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/marcelofabianov/fault"

	"github.com/marcelofabianov/dojo-go/internal/model"
	"github.com/marcelofabianov/dojo-go/internal/port"
	"github.com/marcelofabianov/dojo-go/pkg/web"
)

type GetCourseHandler struct {
	courseService port.CourseServicePort
}

func NewGetCourseHandler(courseService port.CourseServicePort) *GetCourseHandler {
	return &GetCourseHandler{
		courseService: courseService,
	}
}

func (h *GetCourseHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := web.GetLogger(ctx)

	idStr := chi.URLParam(r, "id")
	if _, err := uuid.Parse(idStr); err != nil {
		logger.Warn("invalid uuid format in url param", "id", idStr, "error", err)
		web.Error(w, r, fault.New("invalid id format, must be a valid uuid", fault.WithCode(fault.Invalid)))
		return
	}

	course, err := h.courseService.GetCourseByID(ctx, idStr)
	if err != nil {
		if errors.Is(err, model.ErrCourseNotFound) {
			logger.Warn("course not found", "id", idStr)
			web.Error(w, r, fault.New("course not found", fault.WithCode(fault.NotFound)))
			return
		}

		logger.Error("failed to get course", "id", idStr, "error", err)
		web.Error(w, r, err)
		return
	}

	response := CreateCourseResponse{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		CreatedAt:   course.CreatedAt.String(),
	}

	logger.Info("course retrieved successfully", "course_id", response.ID)

	web.Success(w, r, http.StatusOK, response)
}
