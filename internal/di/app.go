package di

import (
	"net/http"

	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		Config,
		Pkg,
		Repository,
		Service,
		Handler,

		//----

		fx.Invoke(func(*http.Server) {}),
	)
}
