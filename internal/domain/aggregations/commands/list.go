package commands

import (
	"log/slog"
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
			var aggs []aggregations.ContinuousAggregationInfo

			viewName, err := cmd.Flags().GetString("view-name")
			checkErrAndExit(err)

			hypertableName, err := cmd.Flags().GetString("hypertable")
			checkErrAndExit(err)

			if viewName != "" && hypertableName != "" {
				aggs, err = container.AggregationsRepository.GetAggsByHypertableAndViewName(hypertableName, viewName)
			} else if viewName != "" {
				aggs, err = container.AggregationsRepository.GetAggsByViewName(viewName)
			} else if hypertableName != "" {
				aggs, err = container.AggregationsRepository.GetAggsByHypertable(hypertableName)
			} else {
				aggs, err = container.AggregationsRepository.GetAggs()
			}

			checkErrAndExit(err)
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

func checkErrAndExit(err error) {
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
