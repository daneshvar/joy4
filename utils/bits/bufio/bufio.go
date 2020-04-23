package bufio

import (
	"io"
)

type Reader struct {
	buf [][]byte
	R   io.ReadSeeker
}

func NewReaderSize(r io.ReadSeeker, size int) *Reader {
	buf := make([]byte, size*2)
	return &Reader{
		R:   r,
		buf: [][]byte{buf[0:size], buf[size:]},
	}
}
