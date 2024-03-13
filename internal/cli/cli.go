package cli

import (
	"github.com/leometzger/timescale-cli/internal/cli/aggregation"
	"github.com/leometzger/timescale-cli/internal/cli/hypertable"
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/spf13/cobra"
)

func NewCli(options *config.CliOptions) *cobra.Command {
	root := &cobra.Command{}

	root.AddCommand(aggregation.NewAggregationCommand(options))
	root.AddCommand(hypertable.NewHypertableCommands(options))

	return root
}
