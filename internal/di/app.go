package di

import (
	"context"
	"errors"
	"log/slog"
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

		fx.Invoke(registerHooks),
	)
}

func registerHooks(lc fx.Lifecycle, srv *http.Server, logger *slog.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("starting http server", "addr", srv.Addr)
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error("failed to start http server", "error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("stopping http server")
			return srv.Shutdown(ctx)
		},
	})
}
