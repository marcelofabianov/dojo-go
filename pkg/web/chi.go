package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"github.com/marcelofabianov/dojo-go/config"
)

func NewServer(cfg *config.Config, logger *slog.Logger, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.API.Host, cfg.Server.API.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.API.ReadTimeout,
		WriteTimeout: cfg.Server.API.WriteTimeout,
		IdleTimeout:  cfg.Server.API.IdleTimeout,
	}
}

func NewRouter(cfg *config.ServerConfig, logger *slog.Logger) *chi.Mux {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(SlogLoggerMiddleware(logger))
	r.Use(httprate.Limit(
		cfg.API.RateLimit,
		1*time.Minute,
		httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
		httprate.WithResponseHeaders(headersRateLimit()),
	))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(cors.Handler(setCorsOptions(cfg.CORS)))
	r.Use(apiSecurityHeaders(cfg))

	return r
}

func apiSecurityHeaders(cfg *config.ServerConfig) func(http.Handler) http.Handler {
	apiStack := func(next http.Handler) http.Handler {
		return middleware.RequestSize(int64(cfg.API.MaxBodySize))(
			middleware.AllowContentType("application/json")(next),
		)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.URL.Path, "/api") {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("X-Frame-Options", "deny")
			w.Header().Set("X-DNS-Prefetch-Control", "off")
			w.Header().Set("X-Download-Options", "noopen")
			w.Header().Set("Content-Security-Policy", "default-src 'none'")
			w.Header().Set("Referrer-Policy", "no-referrer")
			w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
			w.Header().Set("Cache-Control", "no-store, no-cache")
			w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
			w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
			w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")

			apiStack(next).ServeHTTP(w, r)
		})
	}
}

func headersRateLimit() httprate.ResponseHeaders {
	return httprate.ResponseHeaders{
		Limit:     "X-RateLimit-Limit",
		Remaining: "X-RateLimit-Remaining",
		Reset:     "X-RateLimit-Reset",
	}
}

func setCorsOptions(cfg config.CORSConfig) cors.Options {
	return cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowedMethods:   cfg.AllowedMethods,
		AllowedHeaders:   cfg.AllowedHeaders,
		ExposedHeaders:   cfg.ExposedHeaders,
		AllowCredentials: cfg.AllowCredentials,
	}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}
	Success(w, r, http.StatusOK, status)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Success(w, r, http.StatusOK, nil)
}
