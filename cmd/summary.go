package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tifye/chrono/cmd/chrono"
)

func newSummaryCommand(c *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Summary of time spent on projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			state, err := c.SessionStore.State(cmd.Context())
			if err != nil {
				return err
			}

			c.Logger.Print(state.String())
			return nil
		},
	}

	return cmd
}
