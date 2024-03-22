package testlib

import (
	"testing"

	"github.com/leometzger/timescale-cli/internal/config"
	"github.com/leometzger/timescale-cli/internal/container"
	aggregations "github.com/leometzger/timescale-cli/internal/domain/aggregations/mocks"
	hypertables "github.com/leometzger/timescale-cli/internal/domain/hypertables/mocks"
	printer "github.com/leometzger/timescale-cli/internal/printer/mocks"
)

func GetMockedContainer(t *testing.T) (*container.CliContainer) {
	aggregationsRepo := aggregations.NewMockAggregationsRepository(t)
	hypertablesRepo := hypertables.NewMockHypertablesRepository(t)
	printerMock := printer.NewMockPrinter(t)

	mocks := 

	return container.NewCliContainer(
		aggregationsRepo,
		hypertablesRepo,
		printerMock,
		config.NewCliOptions(),
		config.DefaultConfig(),
	)
}
