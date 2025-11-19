package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tifye/chrono/cmd/chrono"
)

func newSummaryCommand(c *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out",
		Short: "Summary of time spent on projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
