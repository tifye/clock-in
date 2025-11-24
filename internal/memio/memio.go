package memio

import (
	"io"

	"github.com/tifye/chrono/internal/assert"
)

// Buffer is an in-memory implementation of io.ReadWriteSeeker.
// It is intended for testing code that expects a ReadWriteSeeker.
type Buffer struct {
	buf []byte
	off int64
}

// NewBuffer constructs a new Buffer pre-populated with the given string.
func NewBuffer(initial string) *Buffer {
	return &Buffer{buf: []byte(initial), off: 0}
}

// Read implements io.Reader.
func (b *Buffer) Read(p []byte) (int, error) {
	assert.Assert(b != nil, "nil Buffer")

	if b.off >= int64(len(b.buf)) {
		return 0, io.EOF
	}

	n := copy(p, b.buf[b.off:])
	b.off += int64(n)
	return n, nil
}

// Write implements io.Writer.
func (b *Buffer) Write(p []byte) (int, error) {
	assert.Assert(b != nil, "nil Buffer")

	if len(p) == 0 {
		return 0, nil
	}

	end := b.off + int64(len(p))
	assert.Assert(end >= 0, "negative end offset")

	if end > int64(len(b.buf)) {
		grown := make([]byte, end)
		copy(grown, b.buf)
		b.buf = grown
	}

	n := copy(b.buf[b.off:end], p)
	b.off += int64(n)
	return n, nil
}

// Seek implements io.Seeker.
func (b *Buffer) Seek(offset int64, whence int) (int64, error) {
	assert.Assert(b != nil, "nil Buffer")

	var base int64

	switch whence {
	case io.SeekStart:
		base = 0
	case io.SeekCurrent:
		base = b.off
	case io.SeekEnd:
		base = int64(len(b.buf))
	default:
		assert.Assert(false, "invalid whence")
	}

	pos := base + offset
	assert.Assert(pos >= 0, "negative resulting offset")
	assert.Assert(pos <= int64(len(b.buf)), "resulting offset beyond buffer upper bound")

	b.off = pos
	return b.off, nil
}

// String returns the entire contents of the buffer as a string.
func (b *Buffer) String() string {
	if b == nil {
		return ""
	}
	return string(b.buf)
}
