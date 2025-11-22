package store

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/tifye/chrono/internal/assert"
)

type SessionStore struct {
	logger *log.Logger
	target io.ReadWriteSeeker
}

func NewSessionStore(logger *log.Logger, target io.ReadWriteSeeker) *SessionStore {
	assert.AssertNotNil(logger)
	assert.AssertNotNil(target)
	return &SessionStore{
		logger: logger,
		target: target,
	}
}

func (s *SessionStore) ClockIn(ctx context.Context, t time.Time) error {
	assert.Assert(!t.IsZero(), "zero time value")
	assert.Assert(time.Now().After(t), "expected time to be before time.Now")

	unixStr := strconv.FormatInt(t.Unix(), 10)
	data := formatEvent("in", unixStr)
	io.WriteString(s.target, data)

	return nil
}

func (s *SessionStore) ClockOut(ctx context.Context, t time.Time) error {
	assert.Assert(!t.IsZero(), "zero time value")
	assert.Assert(time.Now().After(t), "expected time to be before time.Now")

	unixStr := strconv.FormatInt(t.Unix(), 10)
	data := formatEvent("out", unixStr)
	io.WriteString(s.target, data)

	return nil
}

func (s *SessionStore) ProjectSet(ctx context.Context, project string) error {
	assert.AssertNotEmpty(project, "expect non-empty project name")
	assert.Assert(len(project) < 50, "project name unexpectedly long")

	data := formatEvent("project", project)
	io.WriteString(s.target, data)

	return nil
}

func formatEvent(event string, payload string) string {
	return fmt.Sprintf("%s %s", event, payload)
}
