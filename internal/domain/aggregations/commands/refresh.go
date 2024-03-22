package commands

import (
	"fmt"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func newRefreshCommand(container *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "refresh",
		Aliases: []string{},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("running compress...")
		},
	}

	return cmd
}
