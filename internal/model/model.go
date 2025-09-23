package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyTitle       = errors.New("title cannot be empty")
	ErrEmptyDescription = errors.New("description cannot be empty")
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

type Course struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
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
