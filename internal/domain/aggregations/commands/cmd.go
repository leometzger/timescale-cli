package commands

import (
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/spf13/cobra"
)

func NewAggregationCommand(options *config.CliOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregation",
		Short: "Aggregation commands",
	}

	cmd.AddCommand(newCompressCommand(options))
	cmd.AddCommand(newInspectCommand(options))
	cmd.AddCommand(newListCommand(options))
	cmd.AddCommand(newRefreshCommand(options))

	return cmd
}
