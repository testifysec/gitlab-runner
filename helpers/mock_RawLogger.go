// Code generated by mockery v1.0.0. DO NOT EDIT.

package helpers

import mock "github.com/stretchr/testify/mock"

// MockRawLogger is an autogenerated mock type for the RawLogger type
type MockRawLogger struct {
	mock.Mock
}

// SendRawLog provides a mock function with given fields: args
func (_m *MockRawLogger) SendRawLog(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}
