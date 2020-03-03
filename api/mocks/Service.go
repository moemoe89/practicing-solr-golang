// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import form "github.com/moemoe89/practicing-solr-golang/api/api_struct/form"
import mock "github.com/stretchr/testify/mock"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *Service) Create(data form.UserForm) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(form.UserForm) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(form.UserForm) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id string) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Select provides a mock function with given fields: key, value
func (_m *Service) Select(key string, value string) ([]map[string]interface{}, int, error) {
	ret := _m.Called(key, value)

	var r0 []map[string]interface{}
	if rf, ok := ret.Get(0).(func(string, string) []map[string]interface{}); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string, string) int); ok {
		r1 = rf(key, value)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(key, value)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
