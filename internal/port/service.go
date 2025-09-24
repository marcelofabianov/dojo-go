package port

import (
	"context"

	"github.com/marcelofabianov/dojo-go/internal/model"
)

type CourseServicePort interface {
	CreateCourse(ctx context.Context, input model.NewCourseInput) (*model.Course, error)
	GetCourseByID(ctx context.Context, id string) (*model.Course, error)
}
