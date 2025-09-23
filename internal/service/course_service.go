package service

import (
	"context"

	"github.com/marcelofabianov/dojo-go/internal/model"
	"github.com/marcelofabianov/dojo-go/internal/port"
)

type CourseService struct {
	repo port.CourseRepositoryPort
}

func NewCourseService(repo port.CourseRepositoryPort) *CourseService {
	return &CourseService{repo: repo}
}

func (*CourseService) CreateCourse(ctx context.Context, input model.NewCourseInput) (*model.Course, error) {
	//...

	return nil, nil
}
