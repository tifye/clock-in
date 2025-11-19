package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tifye/clock-in/cmd/chrono"
)

func newSummaryCommand(_ *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out",
		Short: "Summary of time spent on projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd.Help()
			return nil
		},
	}

	return cmd
}
