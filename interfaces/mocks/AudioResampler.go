// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import av "github.com/daneshvar/joy4/av"
import mock "github.com/stretchr/testify/mock"

// AudioResampler is an autogenerated mock type for the AudioResampler type
type AudioResampler struct {
	mock.Mock
}

// Resample provides a mock function with given fields: _a0
func (_m *AudioResampler) Resample(_a0 av.AudioFrame) (av.AudioFrame, error) {
	ret := _m.Called(_a0)

	var r0 av.AudioFrame
	if rf, ok := ret.Get(0).(func(av.AudioFrame) av.AudioFrame); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(av.AudioFrame)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(av.AudioFrame) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
