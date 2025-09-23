package di

import (
	"go.uber.org/fx"

	"github.com/marcelofabianov/dojo-go/config"
	"github.com/marcelofabianov/dojo-go/internal/handler"
	"github.com/marcelofabianov/dojo-go/internal/repository"
	"github.com/marcelofabianov/dojo-go/internal/service"
	"github.com/marcelofabianov/dojo-go/pkg/db"
	"github.com/marcelofabianov/dojo-go/pkg/logger"
	"github.com/marcelofabianov/dojo-go/pkg/validator"
	"github.com/marcelofabianov/dojo-go/pkg/web"
)

// --- Config ---

var Config = fx.Module("config",
	fx.Provide(
		config.NewConfig,
		func(cfg *config.Config) *config.GeneralConfig { return &cfg.General },
		func(cfg *config.Config) *config.LoggerConfig { return &cfg.Logger },
		func(cfg *config.Config) *config.ServerConfig { return &cfg.Server },
		func(cfg *config.Config) *config.DBConfig { return &cfg.DB },
		func(cfg *config.Config) *config.PasswordConfig { return &cfg.Password },
	),
)

// --- PKG ---

var Pkg = fx.Module("pkg",
	fx.Provide(
		logger.NewSlogLogger,
		db.NewPostgresConnection,
		validator.NewValidator,
		web.NewRouter,
		web.NewServer,
	),
)

// --- Repository ---

var Repository = fx.Module("repository",
	fx.Provide(
		repository.NewPostgresCourseRepository,
	),
)

// --- Service ---

var Service = fx.Module("service",
	fx.Provide(
		service.NewCourseService,
	),
)

// --- Handler ---

var Handler = fx.Module("handler",
	fx.Provide(
		handler.NewCreateUserHandler,
	),
)
