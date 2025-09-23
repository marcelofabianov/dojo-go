package model

type NewCourseInput struct {
	Title       string
	Description string
}

type Course struct {
	Title       string
	Description string
}

func NewCourse(input NewCourseInput) *Course {
	return &Course{
		Title:       input.Title,
		Description: input.Description,
	}
}
