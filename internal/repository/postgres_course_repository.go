package repository

import (
	"context"
	"database/sql"
	"errors"

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

func (r *PostgresCourseRepository) GetCourseByID(ctx context.Context, id string) (*model.Course, error) {
	query := `
		SELECT id, title, description, created_at
		FROM courses
		WHERE id = $1
	`

	var course model.Course
	if err := r.db.GetContext(ctx, &course, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrCourseNotFound
		}
		return nil, fault.Wrap(err,
			"failed to get course by id from database",
			fault.WithCode(fault.Internal),
		)
	}

	return &course, nil
}

func (r *PostgresCourseRepository) DeleteCourseByID(ctx context.Context, id string) error {
	query := `DELETE FROM courses WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fault.Wrap(err,
			"failed to delete course by id from database",
			fault.WithCode(fault.Internal),
		)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fault.Wrap(err,
			"failed to get rows affected after delete",
			fault.WithCode(fault.Internal),
		)
	}

	if rowsAffected == 0 {
		return model.ErrCourseNotFound
	}

	return nil
}

func (r *PostgresCourseRepository) UpdateCourse(ctx context.Context, course *model.Course) error {
	query := `
		UPDATE courses
		SET title = :title, description = :description
		WHERE id = :id
	`
	result, err := r.db.NamedExecContext(ctx, query, course)
	if err != nil {
		return fault.Wrap(err,
			"failed to update course in database",
			fault.WithCode(fault.Internal),
		)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fault.Wrap(err,
			"failed to get rows affected after update",
			fault.WithCode(fault.Internal),
		)
	}

	if rowsAffected == 0 {
		return model.ErrCourseNotFound
	}

	return nil
}
