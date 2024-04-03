package hypertables

import (
	"context"
	"log/slog"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
)

type HypertablesRepository interface {
	GetHypertables(filter *HypertablesFilter) ([]HypertableInfo, error)
}

type HypertablesFilter struct {
	Name       string
	Compressed bool
}

type HypertablesRepositoryPg struct {
	conn   *pgx.Conn
	logger *slog.Logger
}

func NewHypertablesRepository(conn *pgx.Conn, logger *slog.Logger) HypertablesRepository {
	return &HypertablesRepositoryPg{
		conn:   conn,
		logger: logger,
	}
}

func (r *HypertablesRepositoryPg) GetHypertables(filter *HypertablesFilter) ([]HypertableInfo, error) {
	query, args := r.buildQuery(filter)

	rows, err := r.conn.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("error getting hypertables information", err)
		return nil, err
	}

	return r.parseRows(rows)
}

func (r *HypertablesRepositoryPg) buildQuery(filter *HypertablesFilter) (string, []interface{}) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.
		Select(
			"hypertable_name",
			"num_chunks",
			"compression_enabled",
			"pg_size_pretty(hypertable_size(format('%I.%I', hypertable_schema, hypertable_name)::regclass)) as size",
		).
		From("timescaledb_information.hypertables")

	return sb.Build()
}

func (r *HypertablesRepositoryPg) parseRows(rows pgx.Rows) ([]HypertableInfo, error) {
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

func (r *HypertablesRepositoryPg) ShowChunks(hypertableName string, newerThan time.Time, olderThan time.Time) ([]Chunk, error) {
	_, err := r.conn.Query(
		context.Background(),
		`SELECT show_chunks($1, $2, $3)`,
		hypertableName,
		newerThan,
		olderThan,
	)
	return nil, err
}

func (r *HypertablesRepositoryPg) parseRowsChunk(rows pgx.Rows) ([]ChunkInfo, error) {
	chunks, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[ChunkInfo],
	)
	if err != nil {
		r.logger.Error("error parsing chunks", "cause", err)
		return nil, err
	}
	return chunks, nil
}
