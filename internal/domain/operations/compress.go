package operations

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/leometzger/timescale-cli/internal/domain/aggregations"
	"github.com/leometzger/timescale-cli/internal/domain/chunks"
	"github.com/leometzger/timescale-cli/internal/domain/hypertables"
)

type TimescaleObject interface {
	hypertables.Hypertable | aggregations.ContinuousAggregation
}

type Compressor interface {
	Compress(hypertable hypertables.Hypertable, olderThan time.Time, newerThan time.Time) error
	Decompress(hypertable hypertables.Hypertable, olderThan time.Time, newerThan time.Time) error
	CompressAggregation(aggregation aggregations.ContinuousAggregation, olderThan time.Time, newerThan time.Time) error
	DecompressAggregation(aggregation aggregations.ContinuousAggregation, olderThan time.Time, newerThan time.Time) error

	CompressChunk(chunk chunks.Chunk) error
	DecompressChunk(chunk chunks.Chunk) error
}

type TimescaleCompressor struct {
	conn *pgx.Conn
}

func NewCompressor(conn *pgx.Conn) *TimescaleCompressor {
	return &TimescaleCompressor{conn: conn}
}

func (c *TimescaleCompressor) Compress(hypertable hypertables.Hypertable, olderThan time.Time) error {
	_, err := c.conn.Exec(
		context.Background(),
		"SELECT compress_chunk(c, true) FROM show_chunks($1)",
		hypertable.HypertableName,
	)

	return err
}

func (c *TimescaleCompressor) Decompress(hypertable hypertables.Hypertable, olderThan time.Time) error {
	_, err := c.conn.Exec(
		context.Background(),
		"SELECT decompress_chunk(c, true) FROM show_chunks($1)",
		hypertable.HypertableName,
	)

	return err
}
