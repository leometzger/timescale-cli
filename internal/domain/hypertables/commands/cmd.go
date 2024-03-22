package commands

import (
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func NewHypertableCommands(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "hypertable",
		Aliases: []string{"h"},
		Short:   "Hypertable commands",
	}

	cmd.AddCommand(newInspectCommand(container))
	cmd.AddCommand(newListCommand(container))

	return cmd
}
