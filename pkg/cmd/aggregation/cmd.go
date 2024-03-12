package aggregation

import (
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/spf13/cobra"
)

func NewAggregationCommand(options *config.CliOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregation",
		Short: "Aggregation commands",
	}

	cmd.AddCommand(NewCompressCommand(options))
	cmd.AddCommand(NewInspectCommand(options))
	cmd.AddCommand(NewListCommand(options))
	cmd.AddCommand(NewRefreshCommand(options))

	return cmd
}
