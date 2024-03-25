package main

import (
	"log/slog"
	"os"

	cli "github.com/leometzger/timescale-cli/internal"
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/printer"
)

func main() {
	options := config.NewCliOptions()

	logger := slog.Default()
	confFile, err := config.LoadConfig(options.ConfigPath, options.Env)
	if err != nil {
		logger.Error("error loading config", err)
		os.Exit(1)
	}

	printer := printer.NewTabwriterPrinter()
	container := container.NewCliContainer(
		printer,
		options,
		confFile,
	)

	cli.NewCli(container).Execute()
}
