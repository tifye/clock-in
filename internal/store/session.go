package store

import (
	"context"
	"io"
	"time"

	"github.com/charmbracelet/log"
	"github.com/tifye/chrono/internal/assert"
)

type SessionStore struct {
	logger *log.Logger
	target io.Writer
}

func NewFileSessionStore(logger *log.Logger, w io.Writer) *SessionStore {
	assert.AssertNotNil(logger)
	return &SessionStore{
		logger: logger,
	}
}

func (s *SessionStore) ClockIn(ctx context.Context, t time.Time) error {
	assert.Assert(!t.IsZero(), "zero time value")
	assert.Assert(time.Now().After(t), "expected time to be before time.Now")
	s.logger.Debug("storing [in] event")
	return nil
}

func (s *SessionStore) ClockOut(ctx context.Context, t time.Time) error {
	assert.Assert(!t.IsZero(), "zero time value")
	assert.Assert(time.Now().After(t), "expected time to be before time.Now")
	s.logger.Debug("storing [out] event")
	return nil
}

func (s *SessionStore) ProjectSet(ctx context.Context, project string) error {
	assert.AssertNotEmpty(project, "expect non-empty project name")
	assert.Assert(len(project) < 50, "project name unexpectedly long")
	s.logger.Debug("storing [project] event")
	return nil
}
