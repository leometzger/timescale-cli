package commands

import (
	"os"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
	"github.com/spf13/cobra"
)

func newListCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			tables, err := container.HypertablesRepository.GetHypertables()
			if err != nil {
				os.Exit(1)
			}

			var values []any
			for _, agg := range tables {
				values = append(values, agg)
			}

			container.Printer.Print(hypertables.HypertableInfo{}, values)
		},
	}

	return cmd
}
