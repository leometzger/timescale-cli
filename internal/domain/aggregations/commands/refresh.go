package commands

import (
	"time"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/spf13/cobra"
)

func newRefreshCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "refresh",
		Aliases: []string{},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			container.Connect()

			viewName, err := cmd.Flags().GetString("view-name")
			exitOnError(err)

			hypertableName, err := cmd.Flags().GetString("hypertable")
			exitOnError(err)

			startParam, err := cmd.Flags().GetString("start")
			exitOnError(err)

			start, err := time.Parse("2006-01-02", startParam)
			exitOnError(err)

			endParam, err := cmd.Flags().GetString("end")
			exitOnError(err)

			end, err := time.Parse("2006-01-02", endParam)
			exitOnError(err)

			aggs, err := container.AggregationsRepository.GetAggregations(&aggregations.AggregationsFilter{
				HypertableName: hypertableName,
				ViewName:       viewName,
			})
			exitOnError(err)

			for _, agg := range aggs {
				container.AggregationsRepository.Refresh(agg.ViewName, start, end)
			}
		},
	}

	cmd.Flags().StringP("start", "", "2019-01-01", "Start date for the load")
	cmd.Flags().StringP("end", "", "2020-01-01", "End date for the load")
	cmd.Flags().StringP("view-name", "", "", "filter by view name (with LIKE option)")
	cmd.Flags().StringP("hypertable", "", "", "filter by hypertable name")
	cmd.Flags().StringP("pace", "", "", "pace that refresh should happen")
	cmd.Flags().IntP("parallel", "", 0, "if should happen paralelly and how much paralelism should have")

	return cmd
}
