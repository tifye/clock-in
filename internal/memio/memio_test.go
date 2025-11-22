package memio

import (
	"io"
	"testing"
)

func TestNewBufferInitialContent(t *testing.T) {
	b := NewBuffer("mino")

	got := b.String()
	want := "mino"
	if got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}

func TestBufferReadSequential(t *testing.T) {
	b := NewBuffer("mino")
	buf := make([]byte, 4)

	n, err := b.Read(buf)
	if err != nil && err != io.EOF {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 4 {
		t.Fatalf("n = %d, want 4", n)
	}
	got := string(buf)
	want := "mino"
	if got != want {
		t.Fatalf("read %q, want %q", got, want)
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

	n, err := b.Write([]byte("mino"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 4 {
		t.Fatalf("n = %d, want 4", n)
	}
	got := b.String()
	want := "mino"
	if got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}

func TestBufferSeekAndOverwrite(t *testing.T) {
	b := NewBuffer("mino")

	pos, err := b.Seek(1, io.SeekStart)
	if err != nil {
		t.Fatalf("unexpected error from Seek: %v", err)
	}
	if pos != 1 {
		t.Fatalf("pos = %d, want 1", pos)
	}

	n, err := b.Write([]byte("eep"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 3 {
		t.Fatalf("n = %d, want 3", n)
	}

	got := b.String()
	want := "meep"
	if got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}

func TestBufferSeekEndAndAppend(t *testing.T) {
	b := NewBuffer("mino")

	pos, err := b.Seek(0, io.SeekEnd)
	if err != nil {
		t.Fatalf("unexpected error from Seek: %v", err)
	}
	if pos != 4 {
		t.Fatalf("pos = %d, want 4", pos)
	}

	n, err := b.Write([]byte(" meep"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 5 {
		t.Fatalf("n = %d, want 5", n)
	}

	got := b.String()
	want := "mino meep"
	if got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}
}

func TestBufferImplementsReadWriteSeeker(t *testing.T) {
	var _ io.ReadWriteSeeker = (*Buffer)(nil)
}
