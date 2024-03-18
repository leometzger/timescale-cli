package commands

import (
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newListCommand(options *config.CliOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			p := printer.NewTabwriterPrinter()

			var arr []any
			value := aggregations.ContinuousAggregationInfo{
				ViewName:           "test",
				MaterializedOnly:   true,
				CompressionEnabled: true,
				Finalized:          true,
			}
			arr = append(arr, value)
			p.Print(aggregations.ContinuousAggregationInfo{}, arr)
		},
	}

	return cmd
}
