package port

import (
	"context"

	"github.com/marcelofabianov/dojo-go/internal/model"
)

type CourseRepositoryPort interface {
	CreateCourse(ctx context.Context, course *model.Course) error
	GetCourseByID(ctx context.Context, id string) (*model.Course, error)
	DeleteCourseByID(ctx context.Context, id string) error
	UpdateCourse(ctx context.Context, course *model.Course) error
}
