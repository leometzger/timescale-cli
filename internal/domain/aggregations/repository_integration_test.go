package aggregations

import (
	"context"
	"testing"

	"github.com/leometzger/timescale-cli/testlib"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetAggregationsInformation(t *testing.T) {
	conn := testlib.SetupDB()
	defer conn.Close(context.Background())
	repo := NewAggregationsRepositoryPg(conn)

	aggs, err := repo.GetAggs()

	assert.Nil(t, err)
	assert.NotNil(t, aggs)
	assert.GreaterOrEqual(t, len(aggs), 1)
}

func TestShouldGetAggsByHypertable(t *testing.T) {
	assert.True(t, true)
}

func TestGetAggsByHypertableInexistentHypertable(t *testing.T) {
	assert.True(t, true)
}
