// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import av "github.com/daneshvar/joy4/av"
import mock "github.com/stretchr/testify/mock"

// PacketReader is an autogenerated mock type for the PacketReader type
type PacketReader struct {
	mock.Mock
}

// ReadPacket provides a mock function with given fields:
func (_m *PacketReader) ReadPacket() (av.Packet, error) {
	ret := _m.Called()

	var r0 av.Packet
	if rf, ok := ret.Get(0).(func() av.Packet); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(av.Packet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
