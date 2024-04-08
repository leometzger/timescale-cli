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
	conn                db.PgxIface
	AggregationsService aggregations.AggregationsService
	HypertableService   hypertables.HypertablesService
	Printer             printer.Printer
	Options             *config.CliOptions
	ConfigFile          *config.ConfigEnvironment
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

	// dependencies - aggregations
	aggregationsLogger := slog.Default().WithGroup("aggregations")
	aggregationsRepo := aggregations.NewAggregationsRepository(c.conn, aggregationsLogger)
	c.AggregationsService = aggregations.NewAggregationsService(aggregationsRepo, aggregationsLogger)

	// dependencies - hypertables
	hypertablesLogger := slog.Default().WithGroup("hypertables")
	hypertablesRepo := hypertables.NewHypertablesRepository(c.conn, hypertablesLogger)
	c.HypertableService = hypertables.NewHypertablesService(hypertablesRepo, hypertablesLogger)
}

func (c *CliContainer) Close() {
	defer c.conn.Close(context.Background())
}
