package model

import "time"

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

func NewCourse(input NewCourseInput) *Course {
	return &Course{
		Title:       input.Title,
		Description: input.Description,
	}
}

func FromCourse(input FromCourseInput) *Course {
	return &Course{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		CreatedAt:   input.CreatedAt,
	}
}
