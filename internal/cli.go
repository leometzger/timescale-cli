package internal

import (
	"log"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/container"
	aggregation "github.com/leometzger/timescale-cli/internal/domain/aggregations/commands"
	hypertable "github.com/leometzger/timescale-cli/internal/domain/hypertables/commands"
	"github.com/spf13/cobra"
)

func NewCli(container *container.CliContainer) *cobra.Command {
	root := &cobra.Command{}

	cobra.OnInitialize(onInitialize(root, container))

	root.PersistentFlags().StringP("config", "c", config.DefaultConfPath(), "File path of configuration")
	root.PersistentFlags().StringP("env", "e", "development", "Environment of config to use")

	root.AddCommand(aggregation.NewAggregationCommand(container))
	root.AddCommand(hypertable.NewHypertableCommands(container))

	return root
}

func onInitialize(root *cobra.Command, container *container.CliContainer) func() {
	return func() {
		configPath, err := root.PersistentFlags().GetString("config")
		if err != nil {
			log.Fatal("error parsing config path")
		}

		env, err := root.PersistentFlags().GetString("env")
		if err != nil {
			log.Fatalf("%s", err)
		}

		container.Options.Env = env
		container.Options.ConfigPath = configPath
	}
}
