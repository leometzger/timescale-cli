package hypertables

import (
	"fmt"
	"strconv"

	"github.com/leometzger/timescale-cli/internal/domain"
)

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

type HypertablesFilter struct {
	Name       string
	Compressed domain.OptionFlag
}

// representation of a displayed hypertable
// with default important fields
type HypertableInfo struct {
	HypertableName     string
	NumChunks          int64
	CompressionEnabled bool
	Size               string
}

func (h HypertableInfo) Headers() []string {
	return []string{"HYPERTABLE", "CHUNKS", "COMPRESSION ENABLED", "SIZE"}
}

func (h HypertableInfo) Values() []string {
	return []string{
		h.HypertableName,
		fmt.Sprint(h.NumChunks),
		strconv.FormatBool(h.CompressionEnabled),
		h.Size,
	}
}
