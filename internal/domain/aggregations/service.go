package aggregations

import (
	"log/slog"
	"time"
)

type AggregationsService interface {
	GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregationInfo, error)
	Refresh(conf *RefreshConfig) error
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

func (s *aggregationsService) GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregationInfo, error) {
	return s.repo.GetAggregations(filter)
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
		pointer := conf.Start
		paceDuration := time.Duration(24*conf.Pace) * time.Hour

		for pointer.Before(conf.End) {
			for _, agg := range aggs {
				var err error

				if pointer.Add(paceDuration).Before(conf.End) {
					err = s.repo.Refresh(agg.ViewName, pointer, pointer.Add(paceDuration))
				} else {
					err = s.repo.Refresh(agg.ViewName, pointer, conf.End)
				}

				if err != nil {
					return err
				}
			}

			pointer = pointer.Add(paceDuration)
		}
	} else {
		for _, agg := range aggs {
			return s.repo.Refresh(agg.ViewName, conf.Start, conf.End)
		}
	}

	return nil
}
