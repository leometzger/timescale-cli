package pacer_test

import (
	"errors"
	"testing"
	"time"

	"github.com/leometzger/timescale-cli/internal/pacer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldCallExecutionWithPace(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC)
	pace := time.Hour * 2

	counter := 0
	err := pacer.ExecuteWithPace(start, end, pace, func(start, end time.Time) error {
		counter++
		require.Greater(t, end, start)
		require.Equal(t, true, start.Add(2*time.Hour).Equal(end))
		return nil
	})

	require.Nil(t, err)
	assert.Equal(t, counter, 12)
}

func TestShouldReturnAnErrorAndBreakExecution(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC)
	pace := time.Hour * 2

	counter := 0
	err := pacer.ExecuteWithPace(start, end, pace, func(start time.Time, end time.Time) error {
		counter++
		if counter >= 3 {
			return errors.New("error")
		}
		return nil
	})

	assert.Equal(t, counter, 3)
	assert.NotNil(t, err)
}
