package handler

import (
	"log/slog"
	"net/http"

	"github.com/marcelofabianov/dojo-go/internal/port"
	"github.com/marcelofabianov/dojo-go/pkg/validator"
)

type CreateCourseRequest struct{}

type CreateCourseResponse struct{}

type CreateCourseHandler struct {
	logger        *slog.Logger
	validator     *validator.Validator
	courseService *port.CourseServicePort
}

func NewCreateCourseHandler(
	logger *slog.Logger,
	validator *validator.Validator,
	courseService *port.CourseServicePort,
) *CreateCourseHandler {
	return &CreateCourseHandler{
		logger:        logger,
		validator:     validator,
		courseService: courseService,
	}
}

func (h *CreateCourseHandler) Handle(w http.ResponseWriter, r *http.Request) {
	//...
}
