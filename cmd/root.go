package cmd

import (
	"context"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "clock",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	return cmd
}

func addCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		newInCommand(),
		newOutCommand(),
	)
}

func Execute() {
	logger := log.NewWithOptions(os.Stdout, log.Options{})
	logger.Print("Meep")

	rootCmd := newRootCommand()
	addCommands(rootCmd)

	err := rootCmd.ExecuteContext(context.TODO())
	if err != nil {
		logger.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}
