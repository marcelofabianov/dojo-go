package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/marcelofabianov/fault"

	"github.com/marcelofabianov/dojo-go/internal/model"
	"github.com/marcelofabianov/dojo-go/internal/port"
)

type PostgresCourseRepository struct {
	db *sqlx.DB
}

func NewPostgresCourseRepository(db *sqlx.DB) port.CourseRepositoryPort {
	return &PostgresCourseRepository{db: db}
}

func (r *PostgresCourseRepository) CreateCourse(ctx context.Context, course *model.Course) error {
	query := `
		INSERT INTO courses (id, title, description, created_at)
		VALUES (:id, :title, :description, :created_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, course)
	if err != nil {
		return fault.Wrap(err,
			"failed to insert course into database",
			fault.WithCode(fault.Internal),
		)
	}

	return nil
}
