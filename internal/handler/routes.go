package handler

import (
	"github.com/go-chi/chi/v5"

	"github.com/marcelofabianov/dojo-go/pkg/web"
)

func RegisterRoutes(
	r *chi.Mux,
	createCourseHandler *CreateCourseHandler,
) {
	// General routes
	r.Get("/", web.IndexHandler)
	r.Get("/healthz", web.HealthCheckHandler)

	// Courses
	r.Route("/api/v1/courses", func(r chi.Router) {
		r.Post("/", createCourseHandler.Handle)
	})
}
