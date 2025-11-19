package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tifye/clock-in/cmd/chrono"
)

func newOutCommand(_ *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out",
		Short: "Check out and end the active session",
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd.Help()
			return nil
		},
	}

	return cmd
}
