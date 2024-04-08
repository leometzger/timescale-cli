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

			filter, err := parseFilter(cmd)
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

			container.AggregationsService.Refresh(
				&aggregations.RefreshConfig{
					Start:  start,
					End:    end,
					Pace:   pace,
					Filter: filter,
				},
			)
		},
	}

	addFilterParams(cmd)
	addRefreshOptions(cmd)

	// @TODO to implement
	cmd.Flags().StringP("decompress", "", "", "flag if should be decompressed/recompressed while refreshing")
	cmd.Flags().IntP("parallel", "", 0, "if should happen paralelly and how much paralelism should have")

	return cmd
}

func addRefreshOptions(cmd *cobra.Command) {
	cmd.Flags().StringP("start", "s", "2019-01-01", "Start date for the load")
	cmd.Flags().StringP("end", "", "2020-01-01", "End date for the load")
	cmd.Flags().Int16P("pace", "", 0, "pace that refresh should happen (in days)")

}
