// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import av "github.com/daneshvar/joy4/av"
import mock "github.com/stretchr/testify/mock"

// CodecData is an autogenerated mock type for the CodecData type
type CodecData struct {
	mock.Mock
}

// Type provides a mock function with given fields:
func (_m *CodecData) Type() av.CodecType {
	ret := _m.Called()

	var r0 av.CodecType
	if rf, ok := ret.Get(0).(func() av.CodecType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(av.CodecType)
	}

	return r0
}
