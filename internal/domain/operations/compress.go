package operations

import (
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
)

type TimescaleObject interface {
	*hypertables.Hypertable | *aggregations.ContinuousAggregation
}

type Compressor interface {
	Compress(obj TimescaleObject) error
	Uncompress(hypertableName string)
}
