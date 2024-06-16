package aggregations

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/leometzger/timescale-cli/internal/pacer"
)

type AggregationsService interface {
	GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregation, error)
	Refresh(conf *RefreshConfig) error
	Compress(conf *CompressConfig) error
}

type aggregationsService struct {
	repo   AggregationsRepository
	logger *slog.Logger
}

func NewAggregationsService(repo AggregationsRepository, logger *slog.Logger) AggregationsService {
	return &aggregationsService{
		repo:   repo,
		logger: logger,
	}
}

func (s *aggregationsService) GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregation, error) {
	return s.repo.GetAggregations(filter)
}

func (s *aggregationsService) Compress(conf *CompressConfig) error {
	aggs, err := s.repo.GetAggregations(conf.Filter)
	if err != nil {
		return err
	}

	if len(aggs) == 0 {
		s.logger.Info(fmt.Sprintf("no aggregations found with %s", conf.Filter))
		return nil
	}

	for _, agg := range aggs {
		err := s.repo.Compress(agg.ViewName, conf.OlderThan, conf.NewerThan)
		if err != nil {
			s.logger.Error(fmt.Sprintf("%s: error compressing aggregation", agg.ViewName), "cause", err)
		}
	}

	return nil
}

// refreshes a continuous aggregation based on configuration
// of refreshing window
func (s *aggregationsService) Refresh(conf *RefreshConfig) error {
	aggs, err := s.repo.GetAggregations(conf.Filter)
	if err != nil {
		return err
	}

	err = s.repo.SetMaxTuplesDecompressedPerDmlTransaction(0)
	if err != nil {
		s.logger.Error("error changing database parameter on session", "cause", err)
		return err
	}

	if conf.Pace > 0 {
		err = pacer.ExecuteWithPace(
			conf.Start,
			conf.End,
			time.Duration(time.Duration(24*conf.Pace)*time.Hour),
			func(start, end time.Time) error {
				for _, agg := range aggs {
					err = s.repo.Refresh(agg.ViewName, start, end)
				}
				return err
			},
		)
		if err != nil {
			return err
		}
	} else {
		for _, agg := range aggs {
			err := s.repo.Refresh(agg.ViewName, conf.Start, conf.End)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
