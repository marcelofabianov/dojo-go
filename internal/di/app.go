package di

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"github.com/marcelofabianov/dojo-go/pkg/web"
)

func New() *fx.App {
	return fx.New(
		Config,
		Pkg,
		Repository,
		Service,
		Handler,

		//----

		fx.Invoke(func(r *chi.Mux) {
			r.Get("/", web.IndexHandler)
			r.Get("/healthz", web.HealthCheckHandler)
		}),

		fx.Invoke(func(*http.Server) {}),
	)
}
