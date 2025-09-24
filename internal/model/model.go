package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyTitle       = errors.New("title cannot be empty")
	ErrEmptyDescription = errors.New("description cannot be empty")
	ErrCourseNotFound   = errors.New("course not found")
)

type NewCourseInput struct {
	Title       string
	Description string
}

type FromCourseInput struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
}

type UpdateCourseInput struct {
	Title       string
	Description string
}

type Course struct {
	ID          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

func NewCourse(input NewCourseInput) (*Course, error) {
	if input.Title == "" {
		return nil, ErrEmptyTitle
	}

	if input.Description == "" {
		return nil, ErrEmptyDescription
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	created := time.Now()

	return &Course{
		ID:          id.String(),
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   created,
	}, nil
}

func FromCourse(input FromCourseInput) *Course {
	return &Course{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   input.CreatedAt,
	}
}

func (c *Course) Update(input UpdateCourseInput) error {
	if input.Title == "" {
		return ErrEmptyTitle
	}
	if input.Description == "" {
		return ErrEmptyDescription
	}

	c.Title = input.Title
	c.Description = input.Description

	return nil
}
