package main

import (
	"github.com/leometzger/timescale-cli/internal/cmd/aggregation"
	"github.com/leometzger/timescale-cli/internal/cmd/hypertable"
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/spf13/cobra"
)

var options *config.CliOptions
var verbose bool
var configPath string
var root cobra.Command
var env string

func main() {
	root = cobra.Command{}

	options = config.NewCliOptions(configPath, verbose, env)

	root.AddCommand(aggregation.NewAggregationCommand(options))
	root.AddCommand(hypertable.NewHypertableCommands(options))

	root.Execute()
}
