package handler

import (
	"log/slog"
	"net/http"

	"github.com/marcelofabianov/dojo-go/internal/port"
	"github.com/marcelofabianov/dojo-go/pkg/validator"
)

type CreateCourseRequest struct{}

type CreateCourseResponse struct{}

type CreateUserHandler struct {
	logger        *slog.Logger
	validator     *validator.Validator
	courseService *port.CourseServicePort
}

func NewCreateUserHandler(
	logger *slog.Logger,
	validator *validator.Validator,
	courseService *port.CourseServicePort,
) *CreateUserHandler {
	return &CreateUserHandler{
		logger:        logger,
		validator:     validator,
		courseService: courseService,
	}
}

func (h *CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	//...
}
