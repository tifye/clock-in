package store

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/tifye/chrono/internal/assert"
)

const maxProjectNameLength = 50

type SessionStore struct {
	logger *log.Logger
	target io.ReadWriteSeeker
	now    func() time.Time
}

func NewSessionStore(
	logger *log.Logger,
	target io.ReadWriteSeeker,
	now func() time.Time,
) *SessionStore {
	assert.AssertNotNil(logger)
	assert.AssertNotNil(target)
	assert.AssertNotNil(now)
	return &SessionStore{
		logger: logger,
		target: target,
		now:    now,
	}
}

func (s *SessionStore) ClockIn(ctx context.Context, t time.Time) error {
	assert.Assert(!t.IsZero(), "zero time value")
	assert.Assert(s.now().After(t), "expected time to be before time.Now")

	unixStr := strconv.FormatInt(t.Unix(), 10)
	data := formatEvent("in", unixStr)
	io.WriteString(s.target, data)

	return nil
}

func (s *SessionStore) ClockOut(ctx context.Context, t time.Time) error {
	assert.Assert(!t.IsZero(), "zero time value")
	assert.Assert(s.now().After(t), "expected time to be before time.Now")

	unixStr := strconv.FormatInt(t.Unix(), 10)
	data := formatEvent("out", unixStr)
	io.WriteString(s.target, data)

	return nil
}

func (s *SessionStore) ProjectSet(ctx context.Context, project string) error {
	assert.AssertNotEmpty(project, "expect non-empty project name")
	assert.Assert(len(project) <= maxProjectNameLength, "project name unexpectedly long")

	data := formatEvent("project", project)
	io.WriteString(s.target, data)

	return nil
}

func (s *SessionStore) State(ctx context.Context) (State, error) {
	if _, err := s.target.Seek(0, io.SeekStart); err != nil {
		return State{}, err
	}

	state := State{}
	scanner := bufio.NewScanner(s.target)

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		parts := bytes.SplitN(line, []byte{' '}, 2)
		// Here I assert on a response from an IO which
		// technically could be invalid because it can be
		// altered outside the program. However I wouldn't
		// just want to skip and its unlikely to be edited
		// outside the program.
		assert.Assert(len(parts) == 2, "expected only two parts per line")

		switch string(parts[0]) {
		case "project":
			state.ActiveProject = string(parts[1])
		case "in":
			sec, err := strconv.ParseInt(string(parts[1]), 10, 64)
			if err != nil {
				return State{}, err
			}
			state.Since = time.Unix(sec, 0)
			state.ClockState = InClockState
		case "out":
			sec, err := strconv.ParseInt(string(parts[1]), 10, 64)
			if err != nil {
				return State{}, err
			}
			state.Since = time.Unix(sec, 0)
			state.ClockState = OutClockState
		}

		select {
		case <-ctx.Done():
			return State{}, ctx.Err()
		default:
		}
	}

	if err := scanner.Err(); err != nil {
		return State{}, err
	}

	assert.Assert(len(state.ActiveProject) <= maxProjectNameLength, "project name unexpectedly long")
	assert.Assert(s.now().After(state.Since), "State.Since is expected to be in the past")
	assert.Assert(s.now().Sub(state.Since) < (time.Hour*24*365), "over a year since last state change")
	return state, nil
}

func formatEvent(event string, payload string) string {
	return fmt.Sprintf("%s %s\n", event, payload)
}
