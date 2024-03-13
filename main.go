package main

import (
	"github.com/leometzger/timescale-cli/internal/cli"
	"github.com/leometzger/timescale-cli/internal/config"
)

func main() {
	options := config.NewCliOptions(
		"",
		false,
		"",
	)

	cli.NewCli(options).Execute()
}
