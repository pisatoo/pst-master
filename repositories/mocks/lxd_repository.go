// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/pisatoo/pst-master/models"

// LxdRepository is an autogenerated mock type for the LxdRepository type
type LxdRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: l
func (_m *LxdRepository) Create(l models.LXD) models.LXD {
	ret := _m.Called(l)

	var r0 models.LXD
	if rf, ok := ret.Get(0).(func(models.LXD) models.LXD); ok {
		r0 = rf(l)
	} else {
		r0 = ret.Get(0).(models.LXD)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *LxdRepository) Delete(id int) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Lxd provides a mock function with given fields: id
func (_m *LxdRepository) Lxd(id int) models.LXD {
	ret := _m.Called(id)

	var r0 models.LXD
	if rf, ok := ret.Get(0).(func(int) models.LXD); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.LXD)
	}

	return r0
}

// Lxds provides a mock function with given fields:
func (_m *LxdRepository) Lxds() []models.LXD {
	ret := _m.Called()

	var r0 []models.LXD
	if rf, ok := ret.Get(0).(func() []models.LXD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.LXD)
		}
	}

	return r0
}

// Update provides a mock function with given fields: id, l
func (_m *LxdRepository) Update(id int, l models.LXD) models.LXD {
	ret := _m.Called(id, l)

	var r0 models.LXD
	if rf, ok := ret.Get(0).(func(int, models.LXD) models.LXD); ok {
		r0 = rf(id, l)
	} else {
		r0 = ret.Get(0).(models.LXD)
	}

	return r0
}
