package commands

import (
	"log/slog"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func newCheckConfigCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "check",
		Aliases: []string{},
		Run: func(cmd *cobra.Command, args []string) {
			logger := slog.Default().WithGroup("config")
			logger.Info("checking if config can connect...")
			container.Connect()
			logger.Info("connected ✔")
		},
	}

	return cmd
}
