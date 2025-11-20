package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tifye/chrono/cmd/chrono"
)

func newInCommand(c *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "in",
		Short: "Check in and start a new session",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.SessionStore.ClockIn(cmd.Context(), c.Now())
			if err != nil {
				return fmt.Errorf("store clock in: %s", err)
			}
			return nil
		},
	}

	return cmd
}
