// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/szczynk/MyGram/models"
)

// CommentUsecase is an autogenerated mock type for the CommentUsecase type
type CommentUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *CommentUsecase) Delete(_a0 context.Context, _a1 uint) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: _a0, _a1
func (_m *CommentUsecase) Fetch(_a0 context.Context, _a1 *[]models.Comment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *[]models.Comment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUserID provides a mock function with given fields: _a0, _a1, _a2
func (_m *CommentUsecase) GetByUserID(_a0 context.Context, _a1 *models.Comment, _a2 uint) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Comment, uint) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *CommentUsecase) Store(_a0 context.Context, _a1 *models.Comment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Comment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *CommentUsecase) Update(_a0 context.Context, _a1 models.Comment, _a2 uint) (models.Comment, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 models.Comment
	if rf, ok := ret.Get(0).(func(context.Context, models.Comment, uint) models.Comment); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(models.Comment)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Comment, uint) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentUsecase creates a new instance of CommentUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentUsecase(t mockConstructorTestingTNewCommentUsecase) *CommentUsecase {
	mock := &CommentUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
