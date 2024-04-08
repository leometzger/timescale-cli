package commands

import (
	"fmt"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			container.Connect()

			conf, err := parseRefreshOptions(cmd)
			if err != nil {
				return err
			}

			return container.AggregationsService.Refresh(conf)
		},
	}

	addFilterParams(cmd)
	addRefreshOptions(cmd)

	return cmd
}

func addRefreshOptions(cmd *cobra.Command) {
	cmd.Flags().StringP("start", "s", "2019-01-01", "Start date for the load")
	cmd.Flags().StringP("end", "", "2020-01-01", "End date for the load")
	cmd.Flags().Int16P("pace", "", 0, "pace that refresh should happen (in days)")
	// @TODO to implement
	// cmd.Flags().IntP("concurrency", "", 1, "if should happen concurrently and how much concurrency should have")
}

// parse refresh options
func parseRefreshOptions(cmd *cobra.Command) (*aggregations.RefreshConfig, error) {
	filter, err := parseFilter(cmd)
	if err != nil {
		return nil, err
	}

	startParam, err := cmd.Flags().GetString("start")
	if err != nil {
		return nil, err
	}

	start, err := time.Parse("2006-01-02", startParam)
	if err != nil {
		return nil, fmt.Errorf("could not parse start date, you should use format yyyy-MM-dd")
	}

	endParam, err := cmd.Flags().GetString("end")
	if err != nil {
		return nil, err
	}

	end, err := time.Parse("2006-01-02", endParam)
	if err != nil {
		return nil, fmt.Errorf("could not parse end date, you should use format yyyy-MM-dd")
	}

	pace, err := cmd.Flags().GetInt16("pace")
	if err != nil {
		return nil, err
	}

	return &aggregations.RefreshConfig{
		Start:  start,
		End:    end,
		Pace:   pace,
		Filter: filter,
	}, nil
}
