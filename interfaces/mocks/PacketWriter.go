// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import av "github.com/daneshvar/joy4/av"
import mock "github.com/stretchr/testify/mock"

// PacketWriter is an autogenerated mock type for the PacketWriter type
type PacketWriter struct {
	mock.Mock
}

// WritePacket provides a mock function with given fields: _a0
func (_m *PacketWriter) WritePacket(_a0 av.Packet) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(av.Packet) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
