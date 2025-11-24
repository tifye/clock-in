package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tifye/chrono/cmd/chrono"
	"github.com/tifye/chrono/internal/store"
)

func newOutCommand(c *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out",
		Short: "Check out and end the active session",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.SessionStore.ClockOut(cmd.Context(), c.Now())
			if err != nil {
				if errors.Is(err, store.ErrClockOutBeforeClockIn) {
					c.Logger.Print("Not currently clocked in; run 'clock in' first")
					return nil
				}
				return err
			}
			return nil
		},
	}

	return cmd
}
