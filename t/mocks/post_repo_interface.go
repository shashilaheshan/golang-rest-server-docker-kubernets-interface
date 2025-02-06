// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	models "books-api/t/models"

	mock "github.com/stretchr/testify/mock"
)

// PostRepoInterface is an autogenerated mock type for the PostRepoInterface type
type PostRepoInterface struct {
	mock.Mock
}

// AddPost provides a mock function with given fields: post
func (_m *PostRepoInterface) AddPost(post *models.Post) (bool, error) {
	ret := _m.Called(post)

	if len(ret) == 0 {
		panic("no return value specified for AddPost")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Post) (bool, error)); ok {
		return rf(post)
	}
	if rf, ok := ret.Get(0).(func(*models.Post) bool); ok {
		r0 = rf(post)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*models.Post) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPosts provides a mock function with no fields
func (_m *PostRepoInterface) GetPosts() (*[]models.Post, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPosts")
	}

	var r0 *[]models.Post
	var r1 error
	if rf, ok := ret.Get(0).(func() (*[]models.Post, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *[]models.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Post)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPostRepoInterface creates a new instance of PostRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostRepoInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostRepoInterface {
	mock := &PostRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
