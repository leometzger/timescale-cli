package aggregations

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/leometzger/timescale-cli/internal/db"
	"github.com/leometzger/timescale-cli/internal/domain"
)

type AggregationsRepository interface {
	GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregationInfo, error)
	Refresh(viewName string, start time.Time, end time.Time) error
	SetMaxTuplesDecompressedPerDmlTransaction(value int32) error
}

type AggregationsRepositoryPgx struct {
	conn   db.PgxIface
	logger *slog.Logger
}

func NewAggregationsRepository(conn db.PgxIface, logger *slog.Logger) AggregationsRepository {
	return &AggregationsRepositoryPgx{
		conn:   conn,
		logger: logger,
	}
}

func (r *AggregationsRepositoryPgx) GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregationInfo, error) {
	query, args := r.buildQuery(filter)

	rows, err := r.conn.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("error querying continuous aggregates", "cause", err)
		return nil, err
	}

	return r.parseRows(rows)
}

func (r *AggregationsRepositoryPgx) Refresh(viewName string, start time.Time, end time.Time) error {
	r.logger.Info(
		fmt.Sprintf(
			"refreshing %s from %s to %s",
			viewName,
			start.Format("2006-01-02"),
			end.Format("2006-01-02"),
		),
	)

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

func (r *AggregationsRepositoryPgx) buildQuery(filter *AggregationsFilter) (string, []interface{}) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"hypertable_name",
		"view_name",
		"materialized_only",
		"compression_enabled",
		"finalized",
	).From("timescaledb_information.continuous_aggregates")

	if filter.HypertableName != "" {
		sb.Where(sb.Equal("hypertable_name", filter.HypertableName))
	}

	if filter.ViewName != "" {
		sb.Where(sb.Like("view_name", filter.ViewName))
	}

	if filter.Compressed == domain.OptionFlagTrue {
		sb.Where("compression_enabled = true")
	} else if filter.Compressed == domain.OptionFlagFalse {
		sb.Where("compression_enabled = false")
	}

	return sb.Build()
}

func (r *AggregationsRepositoryPgx) SetMaxTuplesDecompressedPerDmlTransaction(value int32) error {
	_, err := r.conn.Exec(
		context.Background(),
		"SET timescaledb.max_tuples_decompressed_per_dml_transaction = $1",
		value,
	)
	return err
}

func (r *AggregationsRepositoryPgx) parseRows(rows pgx.Rows) ([]ContinuousAggregationInfo, error) {
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
