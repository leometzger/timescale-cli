package commands

import (
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/spf13/cobra"
)

func NewHypertableCommands(options *config.CliOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "hypertable",
		Aliases: []string{"h"},
		Short:   "Hypertable commands",
	}

	cmd.AddCommand(newInspectCommand(options))
	cmd.AddCommand(newListCommand(options))

	return cmd
}
