package aggregation

import (
	"fmt"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/spf13/cobra"
)

func NewListCommand(options *config.CliOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ls",
		Aliases: []string{"list"},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("running compress...")
		},
	}

	return cmd
}