package commands

import (
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/domain"
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
		Example: "timescale hypertable ls --env prod",
		Run: func(cmd *cobra.Command, args []string) {
			container.Connect()

			filter, err := parseFilter(cmd)
			exitOnError(err)

			aggs, err := container.AggregationsService.GetAggregations(filter)
			exitOnError(err)

			var values []printer.Printable = make([]printer.Printable, len(aggs))
			for i, agg := range aggs {
				values[i] = agg
			}

			container.Printer.Print(aggregations.ContinuousAggregation{}, values)
		},
	}

	addFilterParams(cmd)
	return cmd
}

func addFilterParams(cmd *cobra.Command) {
	cmd.Flags().StringP("view", "v", "", "filter by view name (with LIKE option)")
	cmd.Flags().StringP("hypertable", "", "", "filter by hypertable name")
	cmd.Flags().BoolP("compressed", "", false, "filter by only compressed continuous aggregations")
}

func parseFilter(cmd *cobra.Command) (*aggregations.AggregationsFilter, error) {
	viewName, err := cmd.Flags().GetString("view")
	if err != nil {
		return nil, err
	}

	hypertableName, err := cmd.Flags().GetString("hypertable")
	if err != nil {
		return nil, err
	}

	compressedParam, err := cmd.Flags().GetBool("compressed")
	if err != nil {
		return nil, err
	}
	var compressed domain.OptionFlag
	if compressedParam {
		compressed = domain.OptionFlagTrue
	}

	return &aggregations.AggregationsFilter{
		ViewName:       viewName,
		HypertableName: hypertableName,
		Compressed:     compressed,
	}, nil
}
