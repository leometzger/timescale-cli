package commands

import (
	"context"
	"log/slog"
	"os"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/db"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
	"github.com/leometzger/timescale-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newListCommand(options *config.CliOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			logger := slog.Default()

			connectionInfo := db.NewConnectionInfo("localhost", 5432, "postgres", "postgres", "password")
			conn := db.Connect(connectionInfo)
			defer conn.Close(context.Background())

			printer := printer.NewTabwriterPrinter()
			repository := hypertables.NewHypertablesRepository(conn, logger)

			hypertableValues, err := repository.GetHypertables()
			if err != nil {
				logger.Error("could not get hypertables info" + err.Error())
				os.Exit(1)
			}

			var values []any
			for _, value := range hypertableValues {
				values = append(values, value)
			}

			printer.Print(hypertables.HypertableInfo{}, values)
		},
	}

	return cmd
}
