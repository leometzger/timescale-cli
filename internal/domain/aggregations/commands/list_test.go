package commands

import (
	"testing"

	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
)

func getFakeAggregations() []aggregations.ContinuousAggregationInfo {
	return []aggregations.ContinuousAggregationInfo{
		{
			HypertableName:     "metrics",
			ViewName:           "aggregation_hourly",
			MaterializedOnly:   false,
			CompressionEnabled: false,
			Finalized:          true,
		},
	}
}

func TestListAggregationsWithoutAnyFilter(t *testing.T) {
	t.Skip()
}

func TestListAggregationsWithFilter(t *testing.T) {
	t.Skip()
}
