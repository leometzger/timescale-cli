package main

import (
	"context"
	"log/slog"
	"os"

	cli "github.com/leometzger/timescale-cli/internal"
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/db"
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

	cli.NewCli(conn, options).Execute()
}
