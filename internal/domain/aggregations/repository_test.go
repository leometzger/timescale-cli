package aggregations

import (
	"context"
	"log/slog"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
)

type AggregationsFilterTestCase struct {
	expectedSQL      string
	expectedResult   []ContinuousAggregationInfo
	databaseResponse *pgxmock.Rows
}

func TestGetAggretionsApplyingFilters(t *testing.T) {
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(context.Background())

	testCases := []AggregationsFilterTestCase{
		{
			expectedSQL: `
				SELECT 
					hypertable_name, 
					view_name, 
					materialized_only, 
					compression_enabled, 
					finalized 
				FROM timescaledb_information.ontinuous_aggregates
			`,
			expectedResult: []ContinuousAggregationInfo{},
			databaseResponse: pgxmock.NewRows([]string{"hypertable_name", "view_name", "materialized_only", "compression_enabled", "finalized"}).
				AddRow("metrics", "metrics_by_hour", true, true, true).
				AddRow("metrics", "metrics_by_day", true, true, true),
		},
	}
	repo := NewAggregationsRepository(mock, slog.Default())

	for _, test := range testCases {
		mock.ExpectQuery(test.expectedSQL).WillReturnRows(test.databaseResponse)

		result, err := repo.GetAggregations(&AggregationsFilter{})

		assert.Nil(t, err)
		assert.Equal(t, test.expectedResult, result)
	}
}
