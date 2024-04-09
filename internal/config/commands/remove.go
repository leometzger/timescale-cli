package commands

import (
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func newRemoveConfigCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove [env]",
		Aliases: []string{"rm"},
		Run: func(cmd *cobra.Command, args []string) {
			logger := slog.Default()

			if len(args) == 0 {
				logger.Error("could not identify environment to create")
				os.Exit(1)
			}
			env := args[0]

			home, err := os.UserHomeDir()
			if err != nil {
				logger.Error("could not identify user home dir", "cause", err)
				os.Exit(1)
			}

			fullPath := path.Join(home, ".tsctl", config.DefaultConfigFileName)
			err = config.RemoveConfig(env, fullPath)
			exitOnError(err)

			logger.Info(fmt.Sprintf("✔ removed \"%s\" config on %s", env, fullPath))
		},
	}

	return cmd
}
