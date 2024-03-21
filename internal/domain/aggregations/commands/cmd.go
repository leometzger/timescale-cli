package commands

import (
	"github.com/jackc/pgx/v5"
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/spf13/cobra"
)

func NewAggregationCommand(conn *pgx.Conn, options *config.CliOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aggregation",
		Short: "Aggregation commands",
	}

	cmd.AddCommand(newCompressCommand(options))
	cmd.AddCommand(newInspectCommand(options))
	cmd.AddCommand(newListCommand(conn, options))
	cmd.AddCommand(newRefreshCommand(options))

	return cmd
}
