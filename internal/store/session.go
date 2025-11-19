package store

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/tifye/chrono/internal/assert"
)

type FilesSessionStore struct {
	logger *log.Logger
}

func NewFilesSessionStore(logger *log.Logger) *FilesSessionStore {
	assert.AssertNotNil(logger)
	return &FilesSessionStore{
		logger: logger,
	}
}

func (s *FilesSessionStore) ClockIn(t time.Time) error {
	assert.Assert(!t.IsZero(), "zero time value")
	assert.Assert(time.Now().After(t), "expected time to be before time.Now")
	s.logger.Debug("storing [in] event")
	return nil
}

func (s *FilesSessionStore) ClockOut(t time.Time) error {
	assert.Assert(!t.IsZero(), "zero time value")
	assert.Assert(time.Now().After(t), "expected time to be before time.Now")
	s.logger.Debug("storing [out] event")
	return nil
}

func (s *FilesSessionStore) ProjectSet(project string) error {
	assert.AssertNotEmpty(project, "expect non-empty project name")
	assert.Assert(len(project) < 50, "project name unexpectedly long")
	s.logger.Debug("storing [project] event")
	return nil
}
