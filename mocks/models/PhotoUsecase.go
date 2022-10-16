// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/szczynk/MyGram/models"
)

// PhotoUsecase is an autogenerated mock type for the PhotoUsecase type
type PhotoUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *PhotoUsecase) Delete(_a0 context.Context, _a1 uint) error {
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
func (_m *PhotoUsecase) Fetch(_a0 context.Context, _a1 *models.Pagination) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Pagination) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: _a0, _a1, _a2
func (_m *PhotoUsecase) GetByID(_a0 context.Context, _a1 *models.Photo, _a2 uint) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Photo, uint) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUserID provides a mock function with given fields: _a0, _a1, _a2
func (_m *PhotoUsecase) GetByUserID(_a0 context.Context, _a1 *models.Photo, _a2 uint) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Photo, uint) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *PhotoUsecase) Store(_a0 context.Context, _a1 *models.Photo) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Photo) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *PhotoUsecase) Update(_a0 context.Context, _a1 models.Photo, _a2 uint) (models.Photo, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 models.Photo
	if rf, ok := ret.Get(0).(func(context.Context, models.Photo, uint) models.Photo); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(models.Photo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Photo, uint) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPhotoUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewPhotoUsecase creates a new instance of PhotoUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPhotoUsecase(t mockConstructorTestingTNewPhotoUsecase) *PhotoUsecase {
	mock := &PhotoUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
