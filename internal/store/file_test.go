package store

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/charmbracelet/log"
	"github.com/tifye/chrono/internal/memio"
)

func newTestStore(t *testing.T) (*SessionStore, *memio.Buffer) {
	t.Helper()

	logger := log.New(io.Discard)
	buf := memio.NewBuffer("")
	store := NewSessionStore(logger, buf)

	return store, buf
}

func TestNewFileSessionStoreNilLoggerAllowed(t *testing.T) {
	var rw io.ReadWriteSeeker = memio.NewBuffer("")
	_ = NewSessionStore(nil, rw)
}

func TestNewFileSessionStoreNilTargetPanics(t *testing.T) {
	logger := log.New(io.Discard)

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for nil target, got none")
		}
	}()

	_ = NewSessionStore(logger, nil)
}

func TestClockInValidTimeWritesEvent(t *testing.T) {
	store, buf := newTestStore(t)

	ts := time.Unix(123, 0)
	if err := store.ClockIn(context.Background(), ts); err != nil {
		t.Fatalf("ClockIn() error = %v, want nil", err)
	}

	got := buf.String()
	want := "in 123"
	if got != want {
		t.Fatalf("buffer = %q, want %q", got, want)
	}
}

func TestClockOutValidTimeWritesEvent(t *testing.T) {
	store, buf := newTestStore(t)

	ts := time.Unix(456, 0)
	if err := store.ClockOut(context.Background(), ts); err != nil {
		t.Fatalf("ClockOut() error = %v, want nil", err)
	}

	got := buf.String()
	want := "out 456"
	if got != want {
		t.Fatalf("buffer = %q, want %q", got, want)
	}
}

func TestProjectSetValidProjectWritesEvent(t *testing.T) {
	store, buf := newTestStore(t)

	if err := store.ProjectSet(context.Background(), "proj"); err != nil {
		t.Fatalf("ProjectSet() error = %v, want nil", err)
	}

	got := buf.String()
	want := "project proj"
	if got != want {
		t.Fatalf("buffer = %q, want %q", got, want)
	}
}

func TestMultipleEventsAreAppended(t *testing.T) {
	store, buf := newTestStore(t)

	if err := store.ClockIn(context.Background(), time.Unix(1, 0)); err != nil {
		t.Fatalf("ClockIn() error = %v, want nil", err)
	}
	if err := store.ClockOut(context.Background(), time.Unix(2, 0)); err != nil {
		t.Fatalf("ClockOut() error = %v, want nil", err)
	}
	if err := store.ProjectSet(context.Background(), "p"); err != nil {
		t.Fatalf("ProjectSet() error = %v, want nil", err)
	}

	got := buf.String()
	want := "in 1out 2project p"
	if got != want {
		t.Fatalf("buffer = %q, want %q", got, want)
	}
}

func TestClockInZeroTimePanics(t *testing.T) {
	store, _ := newTestStore(t)

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for zero time, got none")
		}
	}()

	_ = store.ClockIn(context.Background(), time.Time{})
}

func TestClockOutZeroTimePanics(t *testing.T) {
	store, _ := newTestStore(t)

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for zero time, got none")
		}
	}()

	_ = store.ClockOut(context.Background(), time.Time{})
}

func TestProjectSetEmptyPanics(t *testing.T) {
	store, _ := newTestStore(t)

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for empty project, got none")
		}
	}()

	_ = store.ProjectSet(context.Background(), "")
}

func TestProjectSetTooLongPanics(t *testing.T) {
	store, _ := newTestStore(t)

	longName := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz" // 52 chars

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for long project, got none")
		}
	}()

	_ = store.ProjectSet(context.Background(), longName)
}
