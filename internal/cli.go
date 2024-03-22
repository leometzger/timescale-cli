package internal

import (
	"github.com/leometzger/timescale-cli/internal/container"
	aggregation "github.com/leometzger/timescale-cli/internal/domain/aggregations/commands"
	hypertable "github.com/leometzger/timescale-cli/internal/domain/hypertables/commands"
	"github.com/spf13/cobra"
)

func NewCli(container *container.CliContainer) *cobra.Command {
	root := &cobra.Command{}

	root.AddCommand(aggregation.NewAggregationCommand(container))
	root.AddCommand(hypertable.NewHypertableCommands(container))

	return root
}
