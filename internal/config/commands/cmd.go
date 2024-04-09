package commands

import (
	"log/slog"
	"os"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func NewConfigCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration commands",
	}

	cmd.AddCommand(newCheckConfigCommand(container))
	cmd.AddCommand(newAddConfigCommand(container))
	cmd.AddCommand(newListConfigsCommand(container))
	cmd.AddCommand(newRemoveConfigCommand(container))

	return cmd
}

func exitOnError(err error) {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
