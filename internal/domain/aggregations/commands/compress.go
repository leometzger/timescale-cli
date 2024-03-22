package commands

import (
	"fmt"

	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/spf13/cobra"
)

func newCompressCommand(options *container.CliContainer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compress",
		Short: "",
		Long:  "",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("running compress...")
		},
	}

	return cmd
}
