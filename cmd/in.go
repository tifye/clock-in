package cmd

import "github.com/spf13/cobra"

func newInCommand() *cobra.Command {
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
