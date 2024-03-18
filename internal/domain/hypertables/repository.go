package hypertables

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

type HypertablesRepository interface {
	GetHypertables() ([]HypertableInfo, error)
}

type HypertablesRepositoryPg struct {
	conn   *pgx.Conn
	logger *slog.Logger
}

func NewHypertablesRepository(conn *pgx.Conn, logger *slog.Logger) *HypertablesRepositoryPg {
	return &HypertablesRepositoryPg{
		conn:   conn,
		logger: logger,
	}
}

func (r *HypertablesRepositoryPg) GetHypertables() ([]HypertableInfo, error) {
	rows, err := r.conn.Query(context.Background(), `
		SELECT 
			hypertable_name,
			num_chunks,
			compression_enabled,
			pg_size_pretty(hypertable_size(format('%I.%I', hypertable_schema, hypertable_name)::regclass)) as size
		FROM timescaledb_information.hypertables`,
	)
	if err != nil {
		r.logger.Error("error getting hypertables information", err)
		return nil, err
	}

	hypertables, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[HypertableInfo],
	)
	if err != nil {
		r.logger.Error("error parsing hypertables query", err)
		return nil, err
	}
	return hypertables, err
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
