package aggregations

import (
	"time"
)

type AggregationsService interface {
	GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregationInfo, error)
	Refresh(conf *RefreshConfig) error
}

type AggregationsServiceImpl struct {
	repo AggregationsRepository
}

func NewAggregationsService(repo AggregationsRepository) AggregationsService {
	return &AggregationsServiceImpl{repo: repo}
}

func (s *AggregationsServiceImpl) GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregationInfo, error) {
	return s.repo.GetAggregations(filter)
}

func (s *AggregationsServiceImpl) Refresh(conf *RefreshConfig) error {
	aggs, err := s.repo.GetAggregations(conf.Filter)
	if err != nil {
		return err
	}

	if conf.Pace > 0 {
		pointer := conf.Start
		paceDuration := time.Duration(24*conf.Pace) * time.Hour

		for pointer.Before(conf.End) {
			for _, agg := range aggs {
				if pointer.Add(paceDuration).Before(conf.End) {
					s.repo.Refresh(agg.ViewName, pointer, pointer.Add(paceDuration))
				} else {
					s.repo.Refresh(agg.ViewName, pointer, conf.End)
				}
			}

			pointer = pointer.Add(paceDuration)
		}
	} else {
		for _, agg := range aggs {
			s.repo.Refresh(agg.ViewName, conf.Start, conf.End)
		}
	}

	return nil
}
