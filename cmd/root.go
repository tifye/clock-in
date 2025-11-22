package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/tifye/chrono/cmd/chrono"
	"github.com/tifye/chrono/internal/store"
)

type rootOptions struct {
	debug       bool
	logToStdout bool
	logWriter   io.Writer
}

func newRootCommand(c *chrono.Context, opts *rootOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use: "clock",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if opts.debug {
				c.Logger.SetLevel(log.DebugLevel)
			}

			if opts.logToStdout && opts.logWriter != nil {
				c.Logger.SetOutput(io.MultiWriter(opts.logWriter, os.Stdout))
			}

			return nil
		},
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	cmd.PersistentFlags().BoolVarP(&opts.debug, "debug", "d", false, "Toggle debug logs")
	cmd.PersistentFlags().BoolVar(&opts.logToStdout, "log-stdout", false, "also log to stdout")

	return cmd
}

func addCommands(cmd *cobra.Command, c *chrono.Context) {
	cmd.AddCommand(
		newInCommand(c),
		newOutCommand(c),
		newSummaryCommand(c),
	)
}

func Execute() {
	logPath, logFile, err := openLogFile()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open log file:", err)
		os.Exit(1)
	}
	defer logFile.Close()

	logger := log.NewWithOptions(logFile, log.Options{
		Level: log.InfoLevel,
	})
	logger.Info("Using log file", "path", logPath)

	fpath, file, err := openChronoFile()
	if err != nil {
		logger.Error("Failed to open chrono file", "error", err)
		os.Exit(1)
	}
	defer file.Close()
	logger.Info("Using config file", "path", fpath)

	chronoCtx := chrono.NewContext(
		logger,
		store.NewSessionStore(logger.WithPrefix("store"), file, time.Now),
		time.Now(),
	)

	opts := &rootOptions{logWriter: logFile}
	rootCmd := newRootCommand(chronoCtx, opts)
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

func openLogFile() (string, *os.File, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", nil, err
	}

	chronoDir := filepath.Join(configDir, "chrono")
	if err := os.MkdirAll(chronoDir, 0o755); err != nil {
		return "", nil, err
	}

	path := filepath.Join(chronoDir, "clock.log")
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		return "", nil, err
	}

	return path, f, nil
}
