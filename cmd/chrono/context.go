package chrono

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/tifye/chrono/internal/assert"
)

type SessionStorer interface {
	ClockIn(context.Context, time.Time) error
	ClockOut(context.Context, time.Time) error
	ProjectSet(context.Context, string) error
}

type Context struct {
	// The time when the context was created. Used for clocking in and out etc.
	now time.Time

	Logger       *log.Logger
	SessionStore SessionStorer
}

func NewContext(
	logger *log.Logger,
	storer SessionStorer,
	now time.Time,
) *Context {
	assert.AssertNotNil(logger)
	assert.AssertNotNil(storer)
	return &Context{
		now:          now,
		Logger:       logger,
		SessionStore: storer,
	}
}

func (c *Context) Now() time.Time {
	// todo: remove time package dependency
	assert.Assert(c.now.Before(time.Now()), "context 'now' is zero value")
	return c.now
}
