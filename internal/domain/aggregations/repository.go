package aggregations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type AggregationsRepository interface {
	GetAggs() ([]ContinuousAggregationInfo, error)
	GetAggsByHypertable(hypertableName string) ([]ContinuousAggregationInfo, error)
	GetAggsByViewName(viewName string) ([]ContinuousAggregationInfo, error)
}

type AggregationsRepositoryPg struct {
	conn *pgx.Conn
}

func NewAggregationsRepositoryPg(conn *pgx.Conn) AggregationsRepository {
	return &AggregationsRepositoryPg{conn: conn}
}

func (r *AggregationsRepositoryPg) GetAggs() ([]ContinuousAggregationInfo, error) {
	rows, err := r.conn.Query(context.Background(), `
			SELECT  
				view_name,
				materialized_only,
				compression_enabled,
				finalized
			FROM timescaledb_information.continuous_aggregates
		`,
	)
	if err != nil {
		return nil, err
	}

	aggregations, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[ContinuousAggregationInfo],
	)
	if err != nil {
		return nil, err
	}
	return aggregations, err
}

func (r *AggregationsRepositoryPg) GetAggsByHypertable(hypertableName string) ([]ContinuousAggregationInfo, error) {
	_, err := r.conn.Exec(
		context.Background(),
		"SELECT * FROM timescaledb_information.continuous_aggregates",
		hypertableName,
	)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *AggregationsRepositoryPg) GetAggsByViewName(viewName string) ([]ContinuousAggregationInfo, error) {
	_, err := r.conn.Exec(
		context.Background(),
		"SELECT * FROM timescaledb_information.continuous_aggregates WHERE view_name like '%$1%'",
		viewName,
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
