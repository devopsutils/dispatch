// Code generated by mockery v1.0.0

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/casbin/casbin/model"

// AdapterMock is an autogenerated mock type for the AdapterMock type
type AdapterMock struct {
	mock.Mock
}

// AddPolicy provides a mock function with given fields: sec, ptype, rule
func (_m *AdapterMock) AddPolicy(sec string, ptype string, rule []string) error {
	ret := _m.Called(sec, ptype, rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string) error); ok {
		r0 = rf(sec, ptype, rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadPolicy provides a mock function with given fields: _a0
func (_m *AdapterMock) LoadPolicy(_a0 model.Model) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Model) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveFilteredPolicy provides a mock function with given fields: sec, ptype, fieldIndex, fieldValues
func (_m *AdapterMock) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	_va := make([]interface{}, len(fieldValues))
	for _i := range fieldValues {
		_va[_i] = fieldValues[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, sec, ptype, fieldIndex)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int, ...string) error); ok {
		r0 = rf(sec, ptype, fieldIndex, fieldValues...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePolicy provides a mock function with given fields: sec, ptype, rule
func (_m *AdapterMock) RemovePolicy(sec string, ptype string, rule []string) error {
	ret := _m.Called(sec, ptype, rule)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string) error); ok {
		r0 = rf(sec, ptype, rule)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SavePolicy provides a mock function with given fields: _a0
func (_m *AdapterMock) SavePolicy(_a0 model.Model) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Model) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
