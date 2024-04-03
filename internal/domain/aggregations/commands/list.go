package commands

import (
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newListCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "List all continuous aggregations from Timescale",
		Long:    `List all continuous aggregations from Timescale given a selected environment to execute the commands.`,
		Run: func(cmd *cobra.Command, args []string) {
			container.Connect()

			viewName, err := cmd.Flags().GetString("view")
			exitOnError(err)

			hypertableName, err := cmd.Flags().GetString("hypertable")
			exitOnError(err)

			filter := &aggregations.AggregationsFilter{
				HypertableName: hypertableName,
				ViewName:       viewName,
			}

			aggs, err := container.AggregationsRepository.GetAggregations(filter)
			exitOnError(err)

			var values []printer.Printable = make([]printer.Printable, len(aggs))
			for i, agg := range aggs {
				values[i] = agg
			}

			container.Printer.Print(aggregations.ContinuousAggregationInfo{}, values)
		},
	}

	cmd.Flags().StringP("view", "v", "", "filter by view name (with LIKE option)")
	cmd.Flags().StringP("hypertable", "", "", "filter by hypertable name")

	return cmd
}
