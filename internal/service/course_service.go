package service

import (
	"context"

	"github.com/marcelofabianov/dojo-go/internal/model"
	"github.com/marcelofabianov/dojo-go/internal/port"
)

type CourseService struct {
	repo port.CourseRepositoryPort
}

func NewCourseService(repo port.CourseRepositoryPort) port.CourseServicePort {
	return &CourseService{repo: repo}
}

func (c *CourseService) CreateCourse(ctx context.Context, input model.NewCourseInput) (*model.Course, error) {
	newCourse, err := model.NewCourse(input)
	if err != nil {
		return nil, err
	}

	err = c.repo.CreateCourse(ctx, newCourse)
	if err != nil {
		return nil, err
	}

	return newCourse, nil
}

func (c *CourseService) GetCourseByID(ctx context.Context, id string) (*model.Course, error) {
	return c.repo.GetCourseByID(ctx, id)
}

func (c *CourseService) DeleteCourseByID(ctx context.Context, id string) error {
	return c.repo.DeleteCourseByID(ctx, id)
}
