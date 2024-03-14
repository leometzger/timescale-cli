package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TimescaleCommander interface {
	CompressHypertable(hypertable string) error
	CompressChunk(chunkName string) error
	DecompressChunk(chunkName string) error
	RefreshContinuousAggregate(caggName string) error
}

type TimescaleCommanderPgx struct {
	conn pgx.Conn
}

func NewTimescaleCommanderPgx(conn pgx.Conn) TimescaleCommander {
	return &TimescaleCommanderPgx{
		conn: conn,
	}
}

func (t *TimescaleCommanderPgx) CompressHypertable(hypertable string) error {
	_, err := t.conn.Exec(context.Background(), "SELECT compress_hypertable($1)", hypertable)
	return err
}

func (t *TimescaleCommanderPgx) CompressChunk(chunkName string) error {
	_, err := t.conn.Exec(context.Background(), "SELECT compress_chunk($1)", chunkName)
	return err
}

func (t *TimescaleCommanderPgx) DecompressChunk(chunkName string) error {
	_, err := t.conn.Exec(context.Background(), "SELECT decompress_chunk($1)", chunkName)
	return err
}

func (t *TimescaleCommanderPgx) RefreshContinuousAggregate(caggName string) error {
	_, err := t.conn.Exec(context.Background(), "SELECT refresh_continuous_aggregate($1)", caggName)
	return err
}
