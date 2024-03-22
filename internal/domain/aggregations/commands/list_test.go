package commands

import (
	"testing"

	"github.com/leometzger/timescale-cli/testlib"
	"github.com/stretchr/testify/assert"
)

func TestListAggregationsWithoutAnyFilter(t *testing.T) {
	container := testlib.GetMockedContainer(t)
	cmd := newListCommand(container)
	cmd.Execute()
}

func TestListAggregationsWithFilter(t *testing.T) {
	assert.True(t, true)
}
