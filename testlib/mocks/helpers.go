package mocks

import (
	"testing"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/container"
	aggregations "github.com/leometzger/timescale-cli/internal/domain/aggregations/mocks"
	hypertables "github.com/leometzger/timescale-cli/internal/domain/hypertables/mocks"
	printer "github.com/leometzger/timescale-cli/internal/printer/mocks"
)

type Mocks struct {
	AggregationsRepository *aggregations.MockAggregationsRepository
	HypertablesRepository  *hypertables.MockHypertablesRepository
	Printer                *printer.MockPrinter
}

func GetMockedContainer(t *testing.T) (*container.CliContainer, *Mocks) {
	aggregationsRepo := aggregations.NewMockAggregationsRepository(t)
	hypertablesRepo := hypertables.NewMockHypertablesRepository(t)
	printerMock := printer.NewMockPrinter(t)

	mocks := &Mocks{
		AggregationsRepository: aggregationsRepo,
		HypertablesRepository:  hypertablesRepo,
		Printer:                printerMock,
	}

	return container.NewCliContainer(
		aggregationsRepo,
		hypertablesRepo,
		printerMock,
		config.NewCliOptions(),
		config.DefaultConfig(),
	), mocks
}
