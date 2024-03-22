package aggregations

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type AggregationsRepository interface {
	GetAggs() ([]ContinuousAggregationInfo, error)
	GetAggsByHypertable(hypertableName string) ([]ContinuousAggregationInfo, error)
	GetAggsByViewName(viewName string) ([]ContinuousAggregationInfo, error)
	GetAggsByHypertableAndViewName(hypertableName string, viewName string) ([]ContinuousAggregationInfo, error)
}

type AggregationsRepositoryPg struct {
	conn   *pgx.Conn
	logger *slog.Logger
}

func NewAggregationsRepository(conn *pgx.Conn, logger *slog.Logger) AggregationsRepository {
	return &AggregationsRepositoryPg{
		conn:   conn,
		logger: logger,
	}
}

func (r *AggregationsRepositoryPg) GetAggs() ([]ContinuousAggregationInfo, error) {
	rows, err := r.conn.Query(context.Background(), `
			SELECT  
				hypertable_name,
				view_name,
				materialized_only,
				compression_enabled,
				finalized
			FROM timescaledb_information.continuous_aggregates
		`,
	)
	if err != nil {
		r.logger.Error("error querying continuous aggregates", "cause", err)
		return nil, err
	}

	return r.parseRows(rows)
}

func (r *AggregationsRepositoryPg) GetAggsByHypertable(hypertableName string) ([]ContinuousAggregationInfo, error) {
	rows, err := r.conn.Query(context.Background(), `
		SELECT 
			hypertable_name,
			view_name,
			materialized_only,
			compression_enabled,
			finalized
		FROM timescaledb_information.continuous_aggregates
		WHERE hypertable_name = $1`,
		hypertableName,
	)
	if err != nil {
		r.logger.Error("error querying continuous aggregates", "cause", err)
		return nil, err
	}

	return r.parseRows(rows)
}

func (r *AggregationsRepositoryPg) GetAggsByViewName(viewName string) ([]ContinuousAggregationInfo, error) {
	rows, err := r.conn.Query(context.Background(), `
			SELECT 
				hypertable_name,
				view_name,
				materialized_only,
				compression_enabled,
				finalized
			FROM timescaledb_information.continuous_aggregates 
			WHERE view_name LIKE $1`,
		viewName,
	)
	if err != nil {
		return nil, err
	}

	return r.parseRows(rows)
}

func (r *AggregationsRepositoryPg) GetAggsByHypertableAndViewName(hypertableName string, viewName string) ([]ContinuousAggregationInfo, error) {
	rows, err := r.conn.Query(context.Background(), `
			SELECT
				hypertable_name,
				view_name,
				materialized_only,
				compression_enabled,
				finalized
			FROM timescaledb_information.continuous_aggregates
			WHERE hypertable_name = $1 AND view_name LIKE $2`,
		hypertableName,
		viewName,
	)
	if err != nil {
		return nil, err
	}

	return r.parseRows(rows)
}

func (r *AggregationsRepositoryPg) parseRows(rows pgx.Rows) ([]ContinuousAggregationInfo, error) {
	aggregations, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[ContinuousAggregationInfo],
	)
	if err != nil {
		r.logger.Error("error parsing aggregations query", "cause", err)
		return nil, err
	}

	return aggregations, err
}
