package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tifye/chrono/cmd/chrono"
	"github.com/tifye/chrono/internal/store"
)

func newInCommand(c *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "in",
		Short: "Check in and start a new session",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.SessionStore.ClockIn(cmd.Context(), c.Now())
			if err != nil {
				if errors.Is(err, store.ErrClockInBeforeClockOut) {
					c.Logger.Print("Already clocked in; run 'out' first")
					return nil
				}
				return err
			}
			return nil
		},
	}

	return cmd
}
