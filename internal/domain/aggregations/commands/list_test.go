package commands

import (
	"testing"

	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/testlib/mocks"
	"github.com/stretchr/testify/assert"
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
	container, mocks := mocks.GetMockedContainer(t)
	aggs := getFakeAggregations()
	mocks.AggregationsRepository.EXPECT().GetAggs().Return(aggs, nil)

	newListCommand(container).Execute()
}

func TestListAggregationsWithFilter(t *testing.T) {
	assert.True(t, true)
}
