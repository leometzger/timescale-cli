package internal

import (
	aggregation "github.com/leometzger/timescale-cli/internal/aggregations/commands"
	"github.com/leometzger/timescale-cli/internal/config"
	hypertable "github.com/leometzger/timescale-cli/internal/hypertables/commands"
	"github.com/spf13/cobra"
)

func NewCli(options *config.CliOptions) *cobra.Command {
	root := &cobra.Command{}

	root.AddCommand(aggregation.NewAggregationCommand(options))
	root.AddCommand(hypertable.NewHypertableCommands(options))

	return root
}
