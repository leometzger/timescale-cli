package commands

import (
	"log/slog"
	"os"
	"path"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newListConfigsCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all available timescale connection configs",
		Long:    "List all available timescaleDB connection configs",
		Example: "timescale config ls",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := slog.Default()

			home, err := os.UserHomeDir()
			if err != nil {
				logger.Error("could not identify user home dir", "cause", err)
				os.Exit(1)
			}
			fullPath := path.Join(home, ".tsctl", config.DefaultConfigFileName)

			confs, err := config.ListConfigs(fullPath)
			if err != nil {
				return err
			}

			var values []printer.Printable = make([]printer.Printable, len(confs))
			for i, conf := range confs {
				values[i] = conf
			}

			return container.Printer.Print(&config.ConfigEnvironment{}, values)
		},
	}

	return cmd
}
