package store

import (
	"fmt"
	"time"
)

type ClockState uint

//go:generate stringer -type=ClockState
const (
	InClockState ClockState = iota
	OutClockState
)

func (c ClockState) String() string {
	switch c {
	case InClockState:
		return "in"
	case OutClockState:
		return "out"
	default:
		return fmt.Sprintf("ClockState(%d)", c)
	}
}

type State struct {
	ActiveProject string
	ClockState    ClockState
	Since         time.Time
}

func (s State) String() string {
	return fmt.Sprintf("project=%s state=%s since=%s", s.ActiveProject, s.ClockState.String(), s.Since.Format(time.RFC3339))
}
