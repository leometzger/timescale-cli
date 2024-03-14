package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type TimescaleRepository interface {
	ShowChunks(hypertableName string, newerThan time.Time, olderThan time.Time) ([]*Chunk, error)
	GetContinuousAggregations() ([]*ContinuousAggregation, error)
	GetContinuousAggregationsByHypertable(hypertableName string) ([]*ContinuousAggregation, error)
	GetContinuousAggregationsByViewName(viewName string) ([]*ContinuousAggregation, error)
	GetHypertables() ([]*ContinuousAggregation, error)
}

type TimescaleRepositoryPgx struct {
	conn *pgx.Conn
}

func NewTimescaleRepository(conn *pgx.Conn) TimescaleRepository {
	return &TimescaleRepositoryPgx{
		conn: conn,
	}
}

func (r *TimescaleRepositoryPgx) ShowChunks(hypertableName string, newerThan time.Time, olderThan time.Time) ([]*Chunk, error) {
	_, err := r.conn.Exec(context.Background(), "SELECT show_chunks($1, $2, $3)", hypertableName, newerThan, olderThan)
	return nil, err
}

func (r *TimescaleRepositoryPgx) GetContinuousAggregations() ([]*ContinuousAggregation, error) {
	_, err := r.conn.Exec(context.Background(), "SELECT * FROM timescaledb_information.continuous_aggregates")
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *TimescaleRepositoryPgx) GetContinuousAggregationsByHypertable(hypertableName string) ([]*ContinuousAggregation, error) {
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

func (r *TimescaleRepositoryPgx) GetContinuousAggregationsByViewName(viewName string) ([]*ContinuousAggregation, error) {
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

func (r *TimescaleRepositoryPgx) GetHypertables() ([]*ContinuousAggregation, error) {
	_, err := r.conn.Exec(
		context.Background(),
		"SELECT * FROM timescaledb_information.hypertables",
	)
	return nil, err
}
