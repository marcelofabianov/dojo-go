package db

import (
	"context"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/marcelofabianov/fault"

	"github.com/marcelofabianov/dojo-go/config"
)

func NewPostgresConnection(cfg *config.DBConfig, logger *slog.Logger) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, fault.Wrap(err,
			"failed to open database connection",
			fault.WithCode(fault.Internal),
		)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.QueryTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fault.Wrap(err,
			"failed to ping database",
			fault.WithCode(fault.Internal),
		)
	}

	logger.Info("database connection pool established successfully")

	return db, nil
}
