package aggregations

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/leometzger/timescale-cli/internal/domain"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetAggretionsApplyingFilters(t *testing.T) {
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(context.Background())

	databaseResponse := pgxmock.NewRows([]string{"hypertable_name", "view_name", "materialized_only", "compression_enabled", "finalized"})

	testCases := []struct {
		usedFilter       *AggregationsFilter
		expectedSQL      string
		expectedResult   []ContinuousAggregationInfo
		databaseResponse *pgxmock.Rows
		args             []interface{}
	}{
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

func TestRefreshContinuousAggregation(t *testing.T) {
	t.Skip()
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL
	start, _ := time.Parse("2006-01-02", "2024-01-01")
	end, _ := time.Parse("2006-01-02", "2024-02-01")

	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(context.Background())

	testCases := []struct {
		name            string
		refreshQuery    string
		refreshResponse error
		expectedError   error
	}{
		{
			name:            "refreshes aggregation",
			refreshQuery:    "CALL refresh_continuous_aggregate\\('\"metrics_by_hour\"', '2024-01-01', '2024-02-01'\\)",
			refreshResponse: nil,
			expectedError:   nil,
		},
		{
			name:            "refreshes aggregation",
			refreshQuery:    "CALL refresh_continuous_aggregate\\('\"metrics_by_hour\"', '2024-01-01', '2024-02-01'\\)",
			refreshResponse: errors.New("something bad happened"),
			expectedError:   errors.New("something bad happened"),
		},
	}

	repo := NewAggregationsRepository(mock, slog.Default())

	for _, test := range testCases {
		mock.ExpectExec(test.refreshQuery).WillReturnResult(pgconn.CommandTag{}).WillReturnError(test.refreshResponse)
		err := repo.Refresh("metrics_by_hour", start, end)
		assert.Equal(t, test.expectedError, err)
	}
}
