package port

import (
	"context"

	"github.com/marcelofabianov/dojo-go/internal/model"
)

type CourseServicePort interface {
	CreateCourse(ctx context.Context, input model.NewCourseInput) error
}
