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

type DeleteCourseHandler struct {
	courseService port.CourseServicePort
}

func NewDeleteCourseHandler(courseService port.CourseServicePort) *DeleteCourseHandler {
	return &DeleteCourseHandler{
		courseService: courseService,
	}
}

func (h *DeleteCourseHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := web.GetLogger(ctx)

	idStr := chi.URLParam(r, "id")
	if _, err := uuid.Parse(idStr); err != nil {
		logger.Warn("invalid uuid format in url param", "id", idStr, "error", err)
		web.Error(w, r, fault.New("invalid id format, must be a valid uuid", fault.WithCode(fault.Invalid)))
		return
	}

	err := h.courseService.DeleteCourseByID(ctx, idStr)
	if err != nil {
		if errors.Is(err, model.ErrCourseNotFound) {
			logger.Warn("course not found for deletion", "id", idStr)
			web.Error(w, r, fault.New("course not found", fault.WithCode(fault.NotFound)))
			return
		}

		logger.Error("failed to delete course", "id", idStr, "error", err)
		web.Error(w, r, err)
		return
	}

	logger.Info("course deleted successfully", "course_id", idStr)
	web.Success(w, r, http.StatusNoContent, nil)
}
