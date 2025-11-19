package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/tifye/clock-in/cmd/chrono"
)

func newRootCommand(_ *chrono.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use: "clock",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	return cmd
}

func addCommands(cmd *cobra.Command, chrono *chrono.Context) {
	cmd.AddCommand(
		newInCommand(chrono),
		newOutCommand(chrono),
		newSummaryCommand(chrono),
	)
}

func Execute() {
	logger := log.NewWithOptions(os.Stdout, log.Options{})
	chrono := &chrono.Context{
		Logger: logger,
	}

	rootCmd := newRootCommand(chrono)
	addCommands(rootCmd, chrono)

	err := rootCmd.ExecuteContext(context.TODO())
	if err != nil {
		logger.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}
