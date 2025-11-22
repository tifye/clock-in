package cmd

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/tifye/chrono/cmd/chrono"
	"github.com/tifye/chrono/internal/store"
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
	logger := log.NewWithOptions(os.Stdout, log.Options{
		Level: log.DebugLevel,
	})

	fpath, file, err := openChronoFile()
	if err != nil {
		logger.Error("Failed to open chrono file", "error", err)
		os.Exit(1)
	}
	defer file.Close()
	logger.Debug("Using config file", "path", fpath)

	chronoCtx := chrono.NewContext(
		logger,
		store.NewSessionStore(logger.WithPrefix("store"), file, time.Now),
		time.Now(),
	)

	rootCmd := newRootCommand(chronoCtx)
	addCommands(rootCmd, chronoCtx)

	err = rootCmd.ExecuteContext(context.TODO())
	if err != nil {
		logger.Error("Error executing command", "error", err)
		os.Exit(1)
	}
}

func openChronoFile() (string, *os.File, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", nil, err
	}

	chronoDir := filepath.Join(configDir, "chrono")
	if err := os.MkdirAll(chronoDir, 0o755); err != nil {
		return "", nil, err
	}

	path := filepath.Join(chronoDir, "sessions.log")
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return "", nil, err
	}

	if _, err := f.Seek(0, io.SeekEnd); err != nil {
		f.Close()
		return "", nil, err
	}

	return path, f, nil
}
