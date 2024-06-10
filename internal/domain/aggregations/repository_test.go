package aggregations

import (
	"context"
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/leometzger/timescale-cli/testlib"
	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
		require.Equal(t, test.expectedError, err)
	}
}

func TestIntegrationShouldGetAggregationsByViewNameUsingLikeExpressions(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	filter := AggregationsFilter{ViewName: "%by_hour"}

	aggs, err := repo.GetAggregations(&filter)

	require.Nil(t, err)
	require.NotNil(t, aggs)
	require.Equal(t, 1, len(aggs))
}

func TestIntegrationGetAggregationsReturnEmptyList(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	filter := AggregationsFilter{ViewName: "%testingnotfoundagg"}

	aggs, err := repo.GetAggregations(&filter)

	require.Nil(t, err, "should return empty list instead of error when there is no aggregations with view")
	require.NotNil(t, aggs)
	require.Equal(t, 0, len(aggs))
}

func TestIntegrationShouldGetHierarquicalAggregations(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	filter := AggregationsFilter{HypertableName: "metrics_by_hour"}

	aggs, err := repo.GetAggregations(&filter)

	require.Nil(t, err, "should return nil")
	require.NotNil(t, aggs)
	require.Equal(t, 1, len(aggs))
	assert.Equal(t, aggs[0].HypertableName, "metrics_by_hour")
}

func TestIntegrationShouldBeAbleToRefreshAContinuousAggregation(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	start, _ := time.Parse("2006-01-02", "2023-05-31")
	end, _ := time.Parse("2006-01-02", "2023-06-01")

	err := repo.Refresh("metrics_by_day", start, end)

	require.Nil(t, err)
}
