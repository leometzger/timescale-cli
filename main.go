package main

import (
	cli "github.com/leometzger/timescale-cli/internal"
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
