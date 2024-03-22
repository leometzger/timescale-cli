package commands

import (
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func NewAggregationCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregation",
		Short: "Aggregation commands",
	}

	cmd.AddCommand(newCompressCommand(container))
	cmd.AddCommand(newInspectCommand(container))
	cmd.AddCommand(newListCommand(container))
	cmd.AddCommand(newRefreshCommand(container))

	return cmd
}
