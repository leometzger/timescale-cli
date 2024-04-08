package hypertables

import (
	"context"
	"log/slog"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/leometzger/timescale-cli/internal/db"
	"github.com/leometzger/timescale-cli/internal/domain"
)

type HypertablesRepository interface {
	GetHypertables(filter *HypertablesFilter) ([]HypertableInfo, error)
}

type hypertablesRepository struct {
	conn   db.PgxIface
	logger *slog.Logger
}

func NewHypertablesRepository(conn db.PgxIface, logger *slog.Logger) HypertablesRepository {
	return &hypertablesRepository{
		conn:   conn,
		logger: logger,
	}
}

func (r *hypertablesRepository) GetHypertables(filter *HypertablesFilter) ([]HypertableInfo, error) {
	query, args := r.buildQuery(filter)

	rows, err := r.conn.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("error getting hypertables information", err)
		return nil, err
	}

	return r.parseRows(rows)
}

func (r *hypertablesRepository) buildQuery(filter *HypertablesFilter) (string, []interface{}) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"hypertable_name",
		"num_chunks",
		"compression_enabled",
		"pg_size_pretty(hypertable_size(format('%I.%I', hypertable_schema, hypertable_name)::regclass)) as size",
	).From("timescaledb_information.hypertables")

	if filter.Name != "" {
		sb.Where(sb.Like("hypertable_name", filter.Name))
	}

	if filter.Compressed == domain.OptionFlagTrue {
		sb.Where("compression_enabled = true")
	} else if filter.Compressed == domain.OptionFlagFalse {
		sb.Where("compression_enabled = false")
	}

	return sb.Build()
}

func (r *hypertablesRepository) parseRows(rows pgx.Rows) ([]HypertableInfo, error) {
	hypertables, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[HypertableInfo],
	)
	if err != nil {
		r.logger.Error("error parsing hypertables", "cause", err)
		return nil, err
	}
	return hypertables, nil
}
