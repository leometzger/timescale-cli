package aggregations

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/leometzger/timescale-cli/testlib"
	"github.com/stretchr/testify/assert"
)

type TestCaseIntegrationGet struct {
	filter           *AggregationsFilter
	expectedLength   uint
	expectedAggNames []string
}

func TestIntegrationGetAggregationsApplyingFilters(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	pairs := []TestCaseIntegrationGet{
		{
			filter:           &AggregationsFilter{HypertableName: "metrics"},
			expectedAggNames: []string{},
			expectedLength:   0,
		},
		{
			filter:           &AggregationsFilter{ViewName: "%by_hour"},
			expectedAggNames: []string{},
			expectedLength:   0,
		},
		{
			filter:           &AggregationsFilter{ViewName: "%testingnotfoundagg"},
			expectedAggNames: []string{},
			expectedLength:   0,
		},
	}

	repo := NewAggregationsRepository(conn, slog.Default())

	for _, pair := range pairs {
		aggs, err := repo.GetAggregations(pair.filter)

		assert.Nil(t, err)
		assert.NotNil(t, aggs)
		assert.Equal(t, pair.expectedLength, len(aggs))
	}
}

func TestIntegrationRefreshContinuousAggregations(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	start, _ := time.Parse("2006-01-02", "2023-05-31")
	end, _ := time.Parse("2006-01-02", "2023-06-01")

	err := repo.Refresh("metrics_by_day", start, end)

	assert.Nil(t, err)
}
