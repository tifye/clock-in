package cmd

import "github.com/spf13/cobra"

func newSummaryCommand() *cobra.Command {
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
