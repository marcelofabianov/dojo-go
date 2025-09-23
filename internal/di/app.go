package di

import (
	"net/http"

	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		// --- Config ---

		// --- PKG ---

		// --- Repositories ---

		// --- Services ---

		fx.Invoke(func(*http.Server) {}),
	)
}
