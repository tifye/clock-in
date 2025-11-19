package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tifye/clock-in/cmd/chrono"
)

func newInCommand(_ *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "in",
		Short: "Check in and start a new session",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd.Help()
			return nil
		},
	}

	return cmd
}
