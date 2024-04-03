package commands

import (
	"os"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
	"github.com/leometzger/timescale-cli/internal/printer"
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
			container.Connect()

			filter := hypertables.HypertablesFilter{}

			tables, err := container.HypertablesRepository.GetHypertables(&filter)
			if err != nil {
				os.Exit(1)
			}

			var values []printer.Printable = make([]printer.Printable, len(tables))
			for i, table := range tables {
				values[i] = table
			}

			container.Printer.Print(hypertables.HypertableInfo{}, values)
		},
	}

	return cmd
}
