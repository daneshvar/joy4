// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import av "github.com/daneshvar/joy4/av"
import mock "github.com/stretchr/testify/mock"

// AudioEncoder is an autogenerated mock type for the AudioEncoder type
type AudioEncoder struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *AudioEncoder) Close() {
	_m.Called()
}

// CodecData provides a mock function with given fields:
func (_m *AudioEncoder) CodecData() (av.AudioCodecData, error) {
	ret := _m.Called()

	var r0 av.AudioCodecData
	if rf, ok := ret.Get(0).(func() av.AudioCodecData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(av.AudioCodecData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encode provides a mock function with given fields: _a0
func (_m *AudioEncoder) Encode(_a0 av.AudioFrame) ([][]byte, error) {
	ret := _m.Called(_a0)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(av.AudioFrame) [][]byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(av.AudioFrame) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOption provides a mock function with given fields: _a0, _a1
func (_m *AudioEncoder) GetOption(_a0 string, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBitrate provides a mock function with given fields: _a0
func (_m *AudioEncoder) SetBitrate(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetChannelLayout provides a mock function with given fields: _a0
func (_m *AudioEncoder) SetChannelLayout(_a0 av.ChannelLayout) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(av.ChannelLayout) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetOption provides a mock function with given fields: _a0, _a1
func (_m *AudioEncoder) SetOption(_a0 string, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSampleFormat provides a mock function with given fields: _a0
func (_m *AudioEncoder) SetSampleFormat(_a0 av.SampleFormat) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(av.SampleFormat) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSampleRate provides a mock function with given fields: _a0
func (_m *AudioEncoder) SetSampleRate(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
