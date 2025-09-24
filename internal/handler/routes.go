package handler

import (
	"github.com/go-chi/chi/v5"

	"github.com/marcelofabianov/dojo-go/pkg/web"
)

func RegisterRoutes(
	r *chi.Mux,
	createCourseHandler *CreateCourseHandler,
	getCourseHandler *GetCourseHandler,
	deleteCourseHandler *DeleteCourseHandler,
	updateCourseHandler *UpdateCourseHandler,
) {
	// General
	r.Get("/", web.IndexHandler)
	r.Get("/healthz", web.HealthCheckHandler)

	// Courses
	r.Route("/api/v1/courses", func(r chi.Router) {
		r.Post("/", createCourseHandler.Handle)
		r.Get("/{id}", getCourseHandler.Handle)
		r.Delete("/{id}", deleteCourseHandler.Handle)
		r.Put("/{id}", updateCourseHandler.Handle)
	})
}
