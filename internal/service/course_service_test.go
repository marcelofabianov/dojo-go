//go:build unit

package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/marcelofabianov/dojo-go/internal/mocks"
	"github.com/marcelofabianov/dojo-go/internal/model"
	"github.com/marcelofabianov/dojo-go/internal/port"
	service "github.com/marcelofabianov/dojo-go/internal/service"
)

type courseServiceTestSuite struct {
	repoMock *mocks.MockCourseRepository
	service  port.CourseServicePort
}

func setup() *courseServiceTestSuite {
	repoMock := new(mocks.MockCourseRepository)
	svc := service.NewCourseService(repoMock)
	return &courseServiceTestSuite{
		repoMock: repoMock,
		service:  svc,
	}
}

func TestCourseService_CreateCourse(t *testing.T) {
	t.Run("should create course successfully", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		input := model.NewCourseInput{Title: "Test", Description: "Test Desc"}

		s.repoMock.On("CreateCourse", mock.Anything, mock.AnythingOfType("*model.Course")).Return(nil)

		course, err := s.service.CreateCourse(ctx, input)

		assert.NoError(t, err)
		assert.NotNil(t, course)
		assert.Equal(t, input.Title, course.Title)
		assert.NotEmpty(t, course.ID)
		s.repoMock.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		input := model.NewCourseInput{Title: "Test", Description: "Test Desc"}
		expectedErr := errors.New("database error")

		s.repoMock.On("CreateCourse", mock.Anything, mock.AnythingOfType("*model.Course")).Return(expectedErr)

		course, err := s.service.CreateCourse(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, course)
		assert.Equal(t, expectedErr, err)
		s.repoMock.AssertExpectations(t)
	})

	t.Run("should return error on validation failure from model", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		input := model.NewCourseInput{Title: "", Description: "Test Desc"}

		course, err := s.service.CreateCourse(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, course)
		assert.ErrorIs(t, err, model.ErrEmptyTitle)
		s.repoMock.AssertNotCalled(t, "CreateCourse", mock.Anything, mock.Anything)
	})
}

func TestCourseService_GetCourseByID(t *testing.T) {
	t.Run("should get course by id successfully", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		courseID := "test-id"
		mockCourse := &model.Course{ID: courseID, Title: "Test"}

		s.repoMock.On("GetCourseByID", mock.Anything, courseID).Return(mockCourse, nil)

		course, err := s.service.GetCourseByID(ctx, courseID)

		assert.NoError(t, err)
		assert.NotNil(t, course)
		assert.Equal(t, courseID, course.ID)
		s.repoMock.AssertExpectations(t)
	})

	t.Run("should return not found error", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		courseID := "not-found-id"

		s.repoMock.On("GetCourseByID", mock.Anything, courseID).Return(nil, model.ErrCourseNotFound)

		course, err := s.service.GetCourseByID(ctx, courseID)

		assert.Error(t, err)
		assert.Nil(t, course)
		assert.ErrorIs(t, err, model.ErrCourseNotFound)
		s.repoMock.AssertExpectations(t)
	})
}

func TestCourseService_UpdateCourse(t *testing.T) {
	t.Run("should update course successfully", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		courseID := "test-id"
		input := model.UpdateCourseInput{Title: "Updated Title", Description: "Updated Desc"}
		existingCourse := &model.Course{
			ID:          courseID,
			Title:       "Old Title",
			Description: "Old Desc",
			CreatedAt:   time.Now(),
		}

		s.repoMock.On("GetCourseByID", mock.Anything, courseID).Return(existingCourse, nil)
		s.repoMock.On("UpdateCourse", mock.Anything, mock.AnythingOfType("*model.Course")).Return(nil)

		updatedCourse, err := s.service.UpdateCourse(ctx, courseID, input)

		assert.NoError(t, err)
		assert.NotNil(t, updatedCourse)
		assert.Equal(t, input.Title, updatedCourse.Title)
		assert.Equal(t, input.Description, updatedCourse.Description)
		s.repoMock.AssertExpectations(t)
	})

	t.Run("should return not found error when course does not exist", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		courseID := "not-found-id"
		input := model.UpdateCourseInput{Title: "Updated Title", Description: "Updated Desc"}

		s.repoMock.On("GetCourseByID", mock.Anything, courseID).Return(nil, model.ErrCourseNotFound)

		updatedCourse, err := s.service.UpdateCourse(ctx, courseID, input)

		assert.Error(t, err)
		assert.Nil(t, updatedCourse)
		assert.ErrorIs(t, err, model.ErrCourseNotFound)
		s.repoMock.AssertNotCalled(t, "UpdateCourse", mock.Anything, mock.Anything)
	})

	t.Run("should return validation error for invalid input", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		courseID := "test-id"
		input := model.UpdateCourseInput{Title: "", Description: "Updated Desc"}
		existingCourse := &model.Course{ID: courseID}

		s.repoMock.On("GetCourseByID", mock.Anything, courseID).Return(existingCourse, nil)

		updatedCourse, err := s.service.UpdateCourse(ctx, courseID, input)

		assert.Error(t, err)
		assert.Nil(t, updatedCourse)
		assert.ErrorIs(t, err, model.ErrEmptyTitle)
		s.repoMock.AssertNotCalled(t, "UpdateCourse", mock.Anything, mock.Anything)
	})
}

func TestCourseService_DeleteCourseByID(t *testing.T) {
	t.Run("should delete course successfully", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		courseID := "test-id"

		s.repoMock.On("DeleteCourseByID", mock.Anything, courseID).Return(nil)

		err := s.service.DeleteCourseByID(ctx, courseID)

		assert.NoError(t, err)
		s.repoMock.AssertExpectations(t)
	})

	t.Run("should return not found error on delete", func(t *testing.T) {
		s := setup()
		ctx := context.Background()
		courseID := "not-found-id"

		s.repoMock.On("DeleteCourseByID", mock.Anything, courseID).Return(model.ErrCourseNotFound)

		err := s.service.DeleteCourseByID(ctx, courseID)

		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrCourseNotFound)
		s.repoMock.AssertExpectations(t)
	})
}
