package commands

import (
	"log/slog"
	"os"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func NewAggregationCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregation",
		Short: "Aggregation commands",
	}

	cmd.AddCommand(newInspectCommand(container))
	cmd.AddCommand(newListCommand(container))
	cmd.AddCommand(newRefreshCommand(container))
	cmd.AddCommand(newCompressCommand(container))

	return cmd
}

func exitOnError(err error) {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
