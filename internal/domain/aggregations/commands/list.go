package commands

import (
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
		Run: func(cmd *cobra.Command, args []string) {
			viewName, err := cmd.Flags().GetString("view-name")
			exitOnError(err)

			hypertableName, err := cmd.Flags().GetString("hypertable")
			exitOnError(err)

			aggs, err := container.AggregationsRepository.GetAggregations(hypertableName, viewName)
			exitOnError(err)

			var values []any
			for _, agg := range aggs {
				values = append(values, agg)
			}

			container.Printer.Print(aggregations.ContinuousAggregationInfo{}, values)
		},
	}

	cmd.Flags().StringP("view-name", "", "", "filter by view name (with LIKE option)")
	cmd.Flags().StringP("hypertable", "", "", "filter by hypertable name")

	return cmd
}
