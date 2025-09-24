package handler

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

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

	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Courses
	r.Route("/api/v1/courses", func(r chi.Router) {
		r.Post("/", createCourseHandler.Handle)
		r.Get("/{id}", getCourseHandler.Handle)
		r.Delete("/{id}", deleteCourseHandler.Handle)
		r.Put("/{id}", updateCourseHandler.Handle)
	})
}
