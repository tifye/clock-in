package memio

import (
	"io"
	"testing"
)

func TestNewBufferInitialContent(t *testing.T) {
	b := NewBuffer("hello")

	if got := b.String(); got != "hello" {
		t.Fatalf("String() = %q, want %q", got, "hello")
	}
}

func TestBufferReadSequential(t *testing.T) {
	b := NewBuffer("hello")
	buf := make([]byte, 5)

	n, err := b.Read(buf)
	if err != nil && err != io.EOF {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 5 {
		t.Fatalf("n = %d, want 5", n)
	}
	if got := string(buf); got != "hello" {
		t.Fatalf("read %q, want %q", got, "hello")
	}

	n, err = b.Read(buf)
	if n != 0 {
		t.Fatalf("n = %d after EOF read, want 0", n)
	}
	if err != io.EOF {
		t.Fatalf("err = %v, want io.EOF", err)
	}
}

func TestBufferWriteAndString(t *testing.T) {
	b := NewBuffer("")

	n, err := b.Write([]byte("hello"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 5 {
		t.Fatalf("n = %d, want 5", n)
	}
	if got := b.String(); got != "hello" {
		t.Fatalf("String() = %q, want %q", got, "hello")
	}
}

func TestBufferSeekAndOverwrite(t *testing.T) {
	b := NewBuffer("hello")

	pos, err := b.Seek(1, io.SeekStart)
	if err != nil {
		t.Fatalf("unexpected error from Seek: %v", err)
	}
	if pos != 1 {
		t.Fatalf("pos = %d, want 1", pos)
	}

	n, err := b.Write([]byte("i"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 1 {
		t.Fatalf("n = %d, want 1", n)
	}

	if got := b.String(); got != "hillo" {
		t.Fatalf("String() = %q, want %q", got, "hillo")
	}
}

func TestBufferSeekEndAndAppend(t *testing.T) {
	b := NewBuffer("hi")

	pos, err := b.Seek(0, io.SeekEnd)
	if err != nil {
		t.Fatalf("unexpected error from Seek: %v", err)
	}
	if pos != 2 {
		t.Fatalf("pos = %d, want 2", pos)
	}

	n, err := b.Write([]byte(" there"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 6 {
		t.Fatalf("n = %d, want 6", n)
	}

	if got := b.String(); got != "hi there" {
		t.Fatalf("String() = %q, want %q", got, "hi there")
	}
}

func TestBufferImplementsReadWriteSeeker(t *testing.T) {
	var _ io.ReadWriteSeeker = (*Buffer)(nil)
}
