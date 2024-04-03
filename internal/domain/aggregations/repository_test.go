package aggregations

import (
	"context"
	"log/slog"
	"testing"

	"github.com/huandu/go-sqlbuilder"
	"github.com/leometzger/timescale-cli/internal/domain"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
)

type AggregationsFilterTestCase struct {
	usedFilter       *AggregationsFilter
	expectedSQL      string
	expectedResult   []ContinuousAggregationInfo
	databaseResponse *pgxmock.Rows
	args             []interface{}
}

func TestGetAggretionsApplyingFilters(t *testing.T) {
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(context.Background())

	databaseResponse := pgxmock.NewRows([]string{"hypertable_name", "view_name", "materialized_only", "compression_enabled", "finalized"})

	testCases := []AggregationsFilterTestCase{
		{
			usedFilter: &AggregationsFilter{},
			expectedSQL: `
		 		SELECT hypertable_name, view_name, materialized_only, compression_enabled, finalized
		 		FROM timescaledb_information.continuous_aggregates
		 	`,
			expectedResult:   []ContinuousAggregationInfo{},
			databaseResponse: databaseResponse,
		},

		{
			usedFilter: &AggregationsFilter{HypertableName: "metrics"},
			expectedSQL: `
		 		SELECT hypertable_name, view_name, materialized_only, compression_enabled, finalized
		 		FROM timescaledb_information.continuous_aggregates
		 		WHERE hypertable_name = \$1
		 	`,
			expectedResult:   []ContinuousAggregationInfo{},
			databaseResponse: databaseResponse,
			args:             []interface{}{"metrics"},
		},

		{
			usedFilter: &AggregationsFilter{ViewName: "%by_hour"},
			expectedSQL: `
		   		SELECT hypertable_name, view_name, materialized_only, compression_enabled, finalized
		   		FROM timescaledb_information.continuous_aggregates
		   		WHERE view_name LIKE \$1
		   	`,
			databaseResponse: databaseResponse,
			expectedResult:   []ContinuousAggregationInfo{},
			args:             []interface{}{"%by_hour"},
		},

		{
			usedFilter: &AggregationsFilter{HypertableName: "metrics", ViewName: "%by_hour"},
			expectedSQL: `
		  		SELECT hypertable_name, view_name, materialized_only, compression_enabled, finalized
		  		FROM timescaledb_information.continuous_aggregates
		  		WHERE hypertable_name = \$1 AND view_name LIKE \$2
		  	`,
			expectedResult:   []ContinuousAggregationInfo{},
			databaseResponse: databaseResponse,
			args:             []interface{}{"metrics", "%by_hour"},
		},

		{
			usedFilter: &AggregationsFilter{Compressed: domain.OptionFlagTrue},
			expectedSQL: `
		  		SELECT hypertable_name, view_name, materialized_only, compression_enabled, finalized
		  		FROM timescaledb_information.continuous_aggregates
		  		WHERE compression_enabled = true
		  	`,
			databaseResponse: databaseResponse,
			expectedResult:   []ContinuousAggregationInfo{},
		},

		{
			usedFilter: &AggregationsFilter{Compressed: domain.OptionFlagFalse},
			expectedSQL: `
		  		SELECT hypertable_name, view_name, materialized_only, compression_enabled, finalized
		  		FROM timescaledb_information.continuous_aggregates
		  		WHERE compression_enabled = false
		  	`,
			databaseResponse: databaseResponse,
			expectedResult:   []ContinuousAggregationInfo{},
		},
	}
	repo := NewAggregationsRepository(mock, slog.Default())

	for _, test := range testCases {
		mock.ExpectQuery(test.expectedSQL).WithArgs(test.args...).WillReturnRows(test.databaseResponse)

		result, err := repo.GetAggregations(test.usedFilter)

		assert.Nil(t, err)
		assert.Equal(t, test.expectedResult, result)
	}
}
