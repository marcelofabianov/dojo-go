package port

import (
	"context"

	"github.com/marcelofabianov/dojo-go/internal/model"
)

type CourseRepositoryPort interface {
	CreateCourse(ctx context.Context, course model.Course) error
}
