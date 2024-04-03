package aggregations

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/leometzger/timescale-cli/testlib"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetAggregationsInformation(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())

	repo := NewAggregationsRepository(conn, slog.Default())

	aggs, err := repo.GetAggregations(&AggregationsFilter{})

	assert.Nil(t, err)
	assert.NotNil(t, aggs)
	assert.GreaterOrEqual(t, len(aggs), 1)
}

func TestShouldGetAggsByHypertable(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	filter := AggregationsFilter{HypertableName: "metrics"}

	aggs, err := repo.GetAggregations(&filter)

	assert.Nil(t, err)
	assert.NotNil(t, aggs)
	assert.Equal(t, 2, len(aggs))
}

func TestGetAggsByHypertableInexistentHypertable(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	filter := AggregationsFilter{HypertableName: "inexistent_hypertable"}

	aggs, err := repo.GetAggregations(&filter)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(aggs))
}

func TestShouldGetAggregationsByViewNameUsingLikeExpressions(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	filter := AggregationsFilter{ViewName: "%by_hour"}

	aggs, err := repo.GetAggregations(&filter)

	assert.Nil(t, err)
	assert.NotNil(t, aggs)
	assert.Equal(t, 1, len(aggs))
}

func TestGetAggregationsReturnEmptyList(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	filter := AggregationsFilter{ViewName: "%testingnotfoundagg"}

	aggs, err := repo.GetAggregations(&filter)

	assert.Nil(t, err, "should return empty list instead of error when there is no aggregations with view")
	assert.NotNil(t, aggs)
	assert.Equal(t, 0, len(aggs))
}

func TestShouldBeAbleToRefreshAContinuousAggregation(t *testing.T) {
	conn := testlib.GetConnection()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())
	start, _ := time.Parse("2006-01-02", "2023-05-31")
	end, _ := time.Parse("2006-01-02", "2023-06-01")

	err := repo.Refresh("metrics_by_day", start, end)

	assert.Nil(t, err)
}
