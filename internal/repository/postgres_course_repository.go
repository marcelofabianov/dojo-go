package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

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

	return nil
}
