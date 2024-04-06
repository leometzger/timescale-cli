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
		Short:   "refreshes continuous aggregations that match the filter",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			container.Connect()

			viewName, err := cmd.Flags().GetString("view")
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

			pace, err := cmd.Flags().GetInt16("pace")
			exitOnError(err)

			aggs, err := container.AggregationsRepository.GetAggregations(&aggregations.AggregationsFilter{
				HypertableName: hypertableName,
				ViewName:       viewName,
			})
			exitOnError(err)

			if pace > 0 {
				pointer := start

				for pointer.Add(time.Duration(24*pace) * time.Hour).Before(end) {
					for _, agg := range aggs {
						container.AggregationsRepository.Refresh(agg.ViewName, pointer, pointer.Add(time.Duration(pace*24)*time.Hour))
					}

					pointer = pointer.Add(time.Duration(pace*24) * time.Hour)
				}
			} else {
				for _, agg := range aggs {
					container.AggregationsRepository.Refresh(agg.ViewName, start, end)
				}
			}
		},
	}

	cmd.Flags().StringP("start", "s", "2019-01-01", "Start date for the load")
	cmd.Flags().StringP("end", "", "2020-01-01", "End date for the load")
	cmd.Flags().StringP("view", "v", "", "filter by view name (with LIKE option)")
	cmd.Flags().StringP("hypertable", "", "", "filter by hypertable name")
	cmd.Flags().Int16P("pace", "", 0, "pace that refresh should happen (in days)")

	// @TODO to implement
	cmd.Flags().StringP("decompress", "", "", "flag if should be decompressed/recompressed while refreshing")
	cmd.Flags().IntP("parallel", "", 0, "if should happen paralelly and how much paralelism should have")

	return cmd
}
