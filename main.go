package main

import (
	"log/slog"
	"os"

	"github.com/huandu/go-sqlbuilder"
	"github.com/leometzger/timescale-cli/internal/config"
	configCmd "github.com/leometzger/timescale-cli/internal/config/commands"
	"github.com/leometzger/timescale-cli/internal/container"
	aggregation "github.com/leometzger/timescale-cli/internal/domain/aggregations/commands"
	hypertable "github.com/leometzger/timescale-cli/internal/domain/hypertables/commands"
	"github.com/leometzger/timescale-cli/internal/printer"
	"github.com/spf13/cobra"
)

func main() {
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL

	container := container.NewCliContainer(
		printer.NewTabwriterPrinter(os.Stdout),
		config.NewCliOptions(),
	)
	root := &cobra.Command{}

	cobra.OnInitialize(onInitialize(root, container))
	root.PersistentFlags().StringP("env", "e", "", "Environment of config to use")

	root.AddCommand(configCmd.NewConfigCommand(container))
	root.AddCommand(aggregation.NewAggregationCommand(container))
	root.AddCommand(hypertable.NewHypertableCommands(container))

	root.Execute()
}

// initializes the CLI with configuration
func onInitialize(root *cobra.Command, container *container.CliContainer) func() {
	return func() {
		env, err := root.PersistentFlags().GetString("env")
		if err != nil {
			slog.Error("could get env from flags")
			os.Exit(1)
		}

		configFile, err := config.LoadConfig(config.GetDefaultConfigPath(), env)
		if err != nil {
			configFile = config.DefaultConfig()
		}

		container.Options.Env = env
		container.ConfigFile = configFile
	}
}
