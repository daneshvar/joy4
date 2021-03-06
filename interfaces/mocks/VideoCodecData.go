// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import av "github.com/daneshvar/joy4/av"
import mock "github.com/stretchr/testify/mock"

// VideoCodecData is an autogenerated mock type for the VideoCodecData type
type VideoCodecData struct {
	mock.Mock
}

// Height provides a mock function with given fields:
func (_m *VideoCodecData) Height() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *VideoCodecData) Type() av.CodecType {
	ret := _m.Called()

	var r0 av.CodecType
	if rf, ok := ret.Get(0).(func() av.CodecType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(av.CodecType)
	}

	return r0
}

// Width provides a mock function with given fields:
func (_m *VideoCodecData) Width() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
