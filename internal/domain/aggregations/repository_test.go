package aggregations

import (
	"context"
	"log/slog"
	"testing"

	"github.com/leometzger/timescale-cli/testlib"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetAggregationsInformation(t *testing.T) {
	conn := testlib.SetupDB()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())

	aggs, err := repo.GetAggs()

	assert.Nil(t, err)
	assert.NotNil(t, aggs)
	assert.GreaterOrEqual(t, len(aggs), 1)
}

func TestShouldGetAggsByHypertable(t *testing.T) {
	conn := testlib.SetupDB()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())

	aggs, err := repo.GetAggsByHypertable("metrics")

	assert.Nil(t, err)
	assert.NotNil(t, aggs)
	assert.Equal(t, 2, len(aggs))
}

func TestGetAggsByHypertableInexistentHypertable(t *testing.T) {
	conn := testlib.SetupDB()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())

	aggs, err := repo.GetAggsByHypertable("inexistent_hypertable")

	assert.Nil(t, err)
	assert.Equal(t, 0, len(aggs))
}

func TestShouldGetAggregationsByViewNameUsingLikeExpressions(t *testing.T) {
	conn := testlib.SetupDB()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())

	aggs, err := repo.GetAggsByViewName("%by_hour")

	assert.Nil(t, err)
	assert.NotNil(t, aggs)
	assert.Equal(t, 1, len(aggs))
}

func TestGetAggregationsReturnEmptyList(t *testing.T) {
	conn := testlib.SetupDB()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepository(conn, slog.Default())

	aggs, err := repo.GetAggsByViewName("%testingnotfoundagg")

	assert.Nil(t, err, "should return empty list instead of error when there is no aggregations with view")
	assert.NotNil(t, aggs)
	assert.Equal(t, 0, len(aggs))
}
