package commands

import (
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/domain"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
	"github.com/leometzger/timescale-cli/internal/printer"
	"github.com/spf13/cobra"
)

func newListCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Lists all hypertables from an environment",
		Long:    "Lists all hypertables from an environment",
		Example: "timescale hypertable ls --env prod",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			container.Connect()

			filter, err := parseFilter(cmd)
			if err != nil {
				return err
			}

			tables, err := container.HypertableService.GetHypertables(filter)
			if err != nil {
				return err
			}

			var values []printer.Printable = make([]printer.Printable, len(tables))
			for i, table := range tables {
				values[i] = table
			}

			return container.Printer.Print(hypertables.HypertableInfo{}, values)
		},
	}

	addFilterParams(cmd)
	return cmd
}

func addFilterParams(cmd *cobra.Command) {
	cmd.Flags().String("name", "", "filter hypertables by name(like)")
	cmd.Flags().Bool("compressed", false, "shows only compressed hypertables")
	cmd.Flags().Bool("decompressed", false, "shows only decompressed hypertables")
}

func parseFilter(cmd *cobra.Command) (*hypertables.HypertablesFilter, error) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	compressedParam, err := cmd.Flags().GetBool("compressed")
	if err != nil {
		return nil, err
	}

	decompressedParam, err := cmd.Flags().GetBool("compressed")
	if err != nil {
		return nil, err
	}

	compressed := domain.OptionFlagNotDefined
	if compressedParam {
		compressed = domain.OptionFlagTrue
	} else if decompressedParam {
		compressed = domain.OptionFlagFalse
	}

	return &hypertables.HypertablesFilter{
		Name:       name,
		Compressed: compressed,
	}, nil
}
