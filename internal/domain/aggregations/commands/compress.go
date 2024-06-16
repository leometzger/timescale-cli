package commands

import (
	"fmt"
	"time"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/spf13/cobra"
)

func newCompressCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "compress",
		Aliases: []string{},
		Short:   "Compresses continuous aggregations that match the filter",
		Long:    "Compresses continuous aggregations that match the filter",
		Args:    cobra.ExactArgs(0),
		Example: `timescale aggregation compress --view %daily --hypertable metrics`,
		RunE: func(cmd *cobra.Command, args []string) error {
			container.Connect()

			conf, err := parseCompressOptions(cmd)
			if err != nil {
				return err
			}

			return container.AggregationsService.Compress(conf)
		},
	}

	addFilterParams(cmd)
	addCompressionOptions(cmd)
	return cmd
}

func addCompressionOptions(cmd *cobra.Command) {
	cmd.Flags().StringP("start", "s", "", "Start date for compression")
	cmd.Flags().StringP("end", "", "", "End date for compression")
	// @TODO to implement
	// cmd.Flags().IntP("concurrency", "", 1, "if should happen concurrently and how much concurrency should have")
}

// parse refresh options
func parseCompressOptions(cmd *cobra.Command) (*aggregations.CompressConfig, error) {
	var start, end *time.Time

	filter, err := parseFilter(cmd)
	if err != nil {
		return nil, err
	}

	startParam, err := cmd.Flags().GetString("start")
	if err != nil {
		return nil, err
	}

	if startParam != "" {
		startDate, err := time.Parse("2006-01-02", startParam)
		if err != nil {
			return nil, fmt.Errorf("could not parse start date, you should use format yyyy-MM-dd")
		}
		start = &startDate
	}

	endParam, err := cmd.Flags().GetString("end")
	if err != nil {
		return nil, err
	}

	if endParam != "" {
		endDate, err := time.Parse("2006-01-02", endParam)
		if err != nil {
			return nil, fmt.Errorf("could not parse end date, you should use format yyyy-MM-dd")
		}
		end = &endDate
	}

	return &aggregations.CompressConfig{
		NewerThan: start,
		OlderThan: end,
		Filter:    filter,
	}, nil
}
