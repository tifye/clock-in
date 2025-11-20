package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tifye/chrono/cmd/chrono"
)

func newOutCommand(c *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out",
		Short: "Check out and end the active session",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := c.SessionStore.ClockOut(cmd.Context(), c.Now())
			if err != nil {
				return fmt.Errorf("store clock out: %s", err)
			}
			return nil
		},
	}

	return cmd
}
