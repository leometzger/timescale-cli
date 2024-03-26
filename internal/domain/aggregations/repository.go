package aggregations

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

type AggregationsRepository interface {
	GetAggs() ([]ContinuousAggregationInfo, error)
	GetAggsByHypertable(hypertableName string) ([]ContinuousAggregationInfo, error)
	GetAggsByViewName(viewName string) ([]ContinuousAggregationInfo, error)
	GetAggsByHypertableAndViewName(hypertableName string, viewName string) ([]ContinuousAggregationInfo, error)
	GetAggregations(hypertableName string, viewName string) ([]ContinuousAggregationInfo, error)

	Refresh(viewName string, start time.Time, end time.Time) error
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

func (r *AggregationsRepositoryPg) Refresh(viewName string, start time.Time, end time.Time) error {
	r.logger.Info("refreshing continuous aggregation " + viewName)
	command := fmt.Sprintf(
		"CALL refresh_continuous_aggregate('\"%s\"', '%s', '%s')",
		viewName,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
	)

	_, err := r.conn.Exec(context.Background(), command)
	if err != nil {
		r.logger.Error("error refreshing "+viewName, "cause", err.Error())
		return err
	}

	return nil
}

// Guess the best method to find the specified aggregations
func (r *AggregationsRepositoryPg) GetAggregations(hypertableName string, viewName string) ([]ContinuousAggregationInfo, error) {
	if viewName != "" && hypertableName != "" {
		return r.GetAggsByHypertableAndViewName(hypertableName, viewName)
	} else if viewName != "" {
		return r.GetAggsByViewName(viewName)
	} else if hypertableName != "" {
		return r.GetAggsByHypertable(hypertableName)
	} else {
		return r.GetAggs()
	}
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
