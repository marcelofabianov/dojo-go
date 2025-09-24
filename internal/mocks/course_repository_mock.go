package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/marcelofabianov/dojo-go/internal/model"
)

type MockCourseRepository struct {
	mock.Mock
}

func (_m *MockCourseRepository) CreateCourse(ctx context.Context, course *model.Course) error {
	ret := _m.Called(ctx, course)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Course) error); ok {
		r0 = rf(ctx, course)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockCourseRepository) GetCourseByID(ctx context.Context, id string) (*model.Course, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Course
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Course); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Course)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockCourseRepository) UpdateCourse(ctx context.Context, course *model.Course) error {
	ret := _m.Called(ctx, course)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Course) error); ok {
		r0 = rf(ctx, course)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockCourseRepository) DeleteCourseByID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
