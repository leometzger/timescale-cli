package hypertables

// complete representation of hypertable from
// timescaledb_information.hypertables
type Hypertable struct {
	HypertableSchema   string
	HypertableName     string
	Owner              string
	NumDimensions      int64
	NumChunks          int64
	CompressionEnabled bool
}

// representation of a hypertable

// with just the important fields
type HypertableInfo struct {
	HypertableName     string `header:"HYPERTABLE"`
	NumChunks          int64  `header:"CHUNKS"`
	CompressionEnabled bool   `header:"COMPRESSION ENABLED"`
	Size               string `header:"SIZE"`
}

type Chunk struct{}

type ChunkInfo struct {
	ChunkName string `header:"CHUNK"`
	Size      int64  `header:"SIZE"`
}
