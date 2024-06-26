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
	GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregation, error)
	Refresh(viewName string, start time.Time, end time.Time) error
	Compress(viewName string, olderThan *time.Time, newerThan *time.Time) error
	SetMaxTuplesDecompressedPerDmlTransaction(value int32) error
}

type aggregationsRepository struct {
	conn   db.PgxIface
	logger *slog.Logger
}

func NewAggregationsRepository(conn db.PgxIface, logger *slog.Logger) AggregationsRepository {
	return &aggregationsRepository{
		conn:   conn,
		logger: logger,
	}
}

func (r *aggregationsRepository) GetAggregations(filter *AggregationsFilter) ([]ContinuousAggregation, error) {
	query, args := r.buildQuery(filter)

	rows, err := r.conn.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("error querying continuous aggregates", "cause", err)
		return nil, err
	}

	return r.parseRows(rows)
}

func (r *aggregationsRepository) Refresh(viewName string, start time.Time, end time.Time) error {
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

func (r *aggregationsRepository) Compress(viewName string, olderThan *time.Time, newerThan *time.Time) error {
	var command string

	if olderThan == nil && newerThan == nil {
		r.logger.Info(fmt.Sprintf("compressing %s", viewName))

		command = fmt.Sprintf("SELECT compress_chunk(c) FROM show_chunks('\"%s\"') c", viewName)
	}

	if olderThan == nil && newerThan != nil {
		r.logger.Info(fmt.Sprintf("compressing %s newer than %s", viewName, newerThan.Format("2006-01-02")))

		command = fmt.Sprintf(
			"SELECT compress_chunk(c) FROM show_chunks('\"%s\"', newer_than => '%s') c",
			viewName,
			newerThan.Format("2006-01-02"),
		)
	}

	if newerThan == nil && olderThan != nil {
		r.logger.Info(fmt.Sprintf("compressing %s older than %s", viewName, olderThan.Format("2006-01-02")))

		command = fmt.Sprintf(
			"SELECT compress_chunk(c) FROM show_chunks('\"%s\"', older_than => '%s') c",
			viewName,
			olderThan.Format("2006-01-02"),
		)
	}

	if newerThan != nil && olderThan != nil {
		r.logger.Info(
			fmt.Sprintf(
				"compressing %s older than %s and newer than %s",
				viewName,
				olderThan.Format("2006-01-02"),
				newerThan.Format("2006-01-02"),
			),
		)

		command = fmt.Sprintf(
			"SELECT compress_chunk(c) FROM show_chunks('\"%s\"', older_than => '%s', newer_than => '%s') c",
			viewName,
			olderThan.Format("2006-01-02"),
			newerThan.Format("2006-01-02"),
		)
	}

	_, err := r.conn.Exec(context.Background(), command)
	if err != nil {
		return err
	}

	return nil
}

func (r *aggregationsRepository) buildQuery(filter *AggregationsFilter) (string, []interface{}) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"ca.hypertable_schema",
		"coalesce(cats.view_name, ca.hypertable_name) as hypertable_name",
		"ca.view_schema",
		"ca.view_name",
		"ca.view_owner",
		"ca.materialized_only",
		"ca.materialization_hypertable_schema",
		"ca.materialization_hypertable_name",
		"ca.compression_enabled",
		"ca.finalized",
	).
		From("timescaledb_information.continuous_aggregates ca").
		JoinWithOption(
			sqlbuilder.LeftJoin,
			"timescaledb_information.continuous_aggregates cats",
			"ca.hypertable_name = cats.materialization_hypertable_name",
		)

	if filter.HypertableName != "" {
		sb.Where(sb.Equal("coalesce(cats.view_name, ca.hypertable_name)", filter.HypertableName))
	}

	if filter.ViewName != "" {
		sb.Where(sb.Like("ca.view_name", filter.ViewName))
	}

	if filter.Compressed == domain.OptionFlagTrue {
		sb.Where("ca.compression_enabled = true")
	} else if filter.Compressed == domain.OptionFlagFalse {
		sb.Where("ca.compression_enabled = false")
	}

	return sb.Build()
}

func (r *aggregationsRepository) SetMaxTuplesDecompressedPerDmlTransaction(value int32) error {
	_, err := r.conn.Exec(
		context.Background(),
		fmt.Sprintf("SET timescaledb.max_tuples_decompressed_per_dml_transaction = %d", value),
	)
	return err
}

func (r *aggregationsRepository) parseRows(rows pgx.Rows) ([]ContinuousAggregation, error) {
	aggregations, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[ContinuousAggregation],
	)
	if err != nil {
		r.logger.Error("error parsing aggregations query", "cause", err)
		return nil, err
	}

	return aggregations, err
}
