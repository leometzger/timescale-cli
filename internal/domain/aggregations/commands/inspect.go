package commands

import (
	"log/slog"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func newInspectCommand(cliContainer *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "inspect",
		Aliases: []string{"ins"},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			slog.Default().Info("@TODO (in progress)")
		},
	}

	return cmd
}
