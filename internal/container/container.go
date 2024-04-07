package container

import (
	"context"
	"log/slog"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/db"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
	"github.com/leometzger/timescale-cli/internal/printer"
)

type CliContainer struct {
	conn                  db.PgxIface
	AggregationsService   aggregations.AggregationsService
	HypertablesRepository hypertables.HypertablesRepository
	Printer               printer.Printer
	Options               *config.CliOptions
	ConfigFile            *config.ConfigEnvironment
}

func NewCliContainer(
	printer printer.Printer,
	options *config.CliOptions,
) *CliContainer {
	return &CliContainer{
		Printer: printer,
		Options: options,
	}
}

func (c *CliContainer) Connect() {
	confConn := db.LoadConnectionInfoWithConfigFile(c.ConfigFile)
	c.conn = db.Connect(confConn)

	// dependencies
	aggregationsRepo := aggregations.NewAggregationsRepository(c.conn, slog.Default().WithGroup("aggregations"))
	c.AggregationsService = aggregations.NewAggregationsService(aggregationsRepo)

	c.HypertablesRepository = hypertables.NewHypertablesRepository(c.conn, slog.Default().WithGroup("hypertables"))
}

func (c *CliContainer) Close() {
	defer c.conn.Close(context.Background())
}
