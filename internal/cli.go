package internal

import (
	"github.com/leometzger/timescale-cli/internal/config"
	aggregation "github.com/leometzger/timescale-cli/internal/domain/aggregations/commands"
	hypertable "github.com/leometzger/timescale-cli/internal/domain/hypertables/commands"
	"github.com/spf13/cobra"
)

func NewCli(options *config.CliOptions) *cobra.Command {
	root := &cobra.Command{}

	root.AddCommand(aggregation.NewAggregationCommand(options))
	root.AddCommand(hypertable.NewHypertableCommands(options))

	return root
}
