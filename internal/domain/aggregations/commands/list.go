package commands

import (
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newListCommand(conn *pgx.Conn, options *config.CliOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			p := printer.NewTabwriterPrinter()
			logger := slog.Default().WithGroup("test")
			repo := aggregations.NewAggregationsRepository(conn, logger)

			aggs, err := repo.GetAggs()
			if err != nil {
				os.Exit(1)
			}

			var values []any
			for _, agg := range aggs {
				values = append(values, agg)
			}

			p.Print(aggregations.ContinuousAggregationInfo{}, values)
		},
	}

	return cmd
}
