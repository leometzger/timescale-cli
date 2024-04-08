package aggregations_test

import (
	"errors"
	"log/slog"
	"testing"
	"time"

	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	mocks "github.com/leometzger/timescale-cli/internal/domain/aggregations/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRefreshContinuousAggregationWithoutPace(t *testing.T) {
	// arrange
	start, _ := time.Parse("2006-01-02", "2024-01-01")
	end, _ := time.Parse("2006-01-02", "2024-02-01")
	filter := &aggregations.AggregationsFilter{}
	repo := mocks.NewMockAggregationsRepository(t)
	repo.On("Refresh", "testing_aggregation", start, end).Return(nil)
	repo.On("SetMaxTuplesDecompressedPerDmlTransaction", int32(0)).Return(nil)
	repo.On("GetAggregations", filter).Return([]aggregations.ContinuousAggregationInfo{
		{
			HypertableName:     "metrics",
			ViewName:           "testing_aggregation",
			MaterializedOnly:   false,
			CompressionEnabled: false,
			Finalized:          false,
		},
	}, nil)
	service := aggregations.NewAggregationsService(repo, slog.Default())

	// act
	err := service.Refresh(&aggregations.RefreshConfig{
		Start:  start,
		End:    end,
		Filter: filter,
	})

	// verify
	assert.Nil(t, err)
}

func TestRefreshContinuousAggregationWithPace(t *testing.T) {
	// arrange
	start, _ := time.Parse("2006-01-02", "2024-01-01")
	end, _ := time.Parse("2006-01-02", "2024-01-20")

	filter := &aggregations.AggregationsFilter{}

	repo := mocks.NewMockAggregationsRepository(t)
	repo.On(
		"Refresh",
		"testing_aggregation",
		start,
		time.Date(2024, time.January, 8, 0, 0, 0, 0, time.UTC),
	).Return(nil)
	repo.On(
		"Refresh",
		"testing_aggregation",
		time.Date(2024, time.January, 8, 0, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC),
	).Return(nil)

	// if date of pace pass end date, it should use end date
	repo.On(
		"Refresh",
		"testing_aggregation",
		time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2024, time.January, 20, 0, 0, 0, 0, time.UTC),
	).Return(nil)
	repo.On("SetMaxTuplesDecompressedPerDmlTransaction", int32(0)).Return(nil)
	repo.On("GetAggregations", filter).Return([]aggregations.ContinuousAggregationInfo{
		{
			HypertableName:     "metrics",
			ViewName:           "testing_aggregation",
			MaterializedOnly:   false,
			CompressionEnabled: false,
			Finalized:          false,
		},
	}, nil)

	service := aggregations.NewAggregationsService(repo, slog.Default())

	// act
	err := service.Refresh(&aggregations.RefreshConfig{
		Start:  start,
		End:    end,
		Filter: filter,
		Pace:   7, // 7 days
	})

	// verify
	assert.Nil(t, err)
}

func TestRaiseErrorWhenPaceIsSetted(t *testing.T) {
	// arrange
	start, _ := time.Parse("2006-01-02", "2024-01-01")
	end, _ := time.Parse("2006-01-02", "2024-01-20")

	filter := &aggregations.AggregationsFilter{}

	repo := mocks.NewMockAggregationsRepository(t)
	repo.On("SetMaxTuplesDecompressedPerDmlTransaction", int32(0)).Return(nil)
	repo.On(
		"Refresh",
		"testing_aggregation",
		start,
		time.Date(2024, time.January, 8, 0, 0, 0, 0, time.UTC),
	).Return(errors.New("error while calling refresh"))
	repo.On("GetAggregations", filter).Return([]aggregations.ContinuousAggregationInfo{
		{
			HypertableName:     "metrics",
			ViewName:           "testing_aggregation",
			MaterializedOnly:   false,
			CompressionEnabled: false,
			Finalized:          false,
		},
	}, nil)
	service := aggregations.NewAggregationsService(repo, slog.Default())

	// act
	err := service.Refresh(&aggregations.RefreshConfig{
		Start:  start,
		End:    end,
		Filter: filter,
		Pace:   7, // 7 days
	})

	// verify
	assert.NotNil(t, err)
}
