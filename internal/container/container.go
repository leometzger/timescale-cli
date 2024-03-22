package container

import (
	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
	"github.com/leometzger/timescale-cli/internal/printer"
)

type CliContainer struct {
	AggregationsRepository aggregations.AggregationsRepository
	HypertablesRepository  hypertables.HypertablesRepository
	Printer                printer.Printer
	Options                *config.CliOptions
	ConfigFile             *config.ConfigFile
}

func NewCliContainer(
	aggregationsRepository aggregations.AggregationsRepository,
	hypertablesRepository hypertables.HypertablesRepository,
	printer printer.Printer,
	options *config.CliOptions,
	configFile *config.ConfigFile,
) *CliContainer {
	return &CliContainer{
		AggregationsRepository: aggregationsRepository,
		HypertablesRepository:  hypertablesRepository,
		Printer:                printer,
		Options:                options,
		ConfigFile:             configFile,
	}
}
