package main

import (
	"context"
	"log/slog"
	"os"

	cli "github.com/leometzger/timescale-cli/internal"
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/container"
	"github.com/leometzger/timescale-cli/internal/db"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
	"github.com/leometzger/timescale-cli/internal/printer"
)

func main() {
	options := config.NewCliOptions(
		"",
		false,
		"",
	)

	logger := slog.Default()
	confFile, err := config.LoadConfig(options.ConfigPath, options.Env)
	if err != nil {
		logger.Error("error loading config", err)
		os.Exit(1)
	}

	confConn := db.LoadConnectionInfoWithConfigFile(confFile)
	conn := db.Connect(confConn)
	defer conn.Close(context.Background())

	aggsRepo := aggregations.NewAggregationsRepository(conn, slog.Default().WithGroup("aggregations"))
	hypertableRepo := hypertables.NewHypertablesRepository(conn, slog.Default().WithGroup("hypertables"))
	printer := printer.NewTabwriterPrinter()

	container := container.NewCliContainer(aggsRepo, hypertableRepo, printer, options, confFile)

	cli.NewCli(container).Execute()
}
