package commands

import (
	"os"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/spf13/cobra"
)

func newListCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			aggs, err := container.AggregationsRepository.GetAggs()
			if err != nil {
				os.Exit(1)
			}

			var values []any
			for _, agg := range aggs {
				values = append(values, agg)
			}

			container.Printer.Print(aggregations.ContinuousAggregationInfo{}, values)
		},
	}

	return cmd
}
