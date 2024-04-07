package aggregations

import (
	"strconv"
	"time"

	"github.com/leometzger/timescale-cli/internal/domain"
)

// complete representation of continuous_aggregation from
// timescaledb_information.continuous_aggregates
type ContinuousAggregation struct {
	HypertableSchema                string
	HypertableName                  string
	ViewSchema                      string
	ViewName                        string
	ViewOwner                       string
	MaterializedOnly                bool
	CompressionEnabled              bool
	MaterializationHypertableSchema string
	MaterializationHypertableName   string
	ViewDefinition                  string
	Finalized                       bool
}

type AggregationsFilter struct {
	HypertableName string
	ViewName       string
	Compressed     domain.OptionFlag
}

type RefreshConfig struct {
	Filter *AggregationsFilter
	Start  time.Time
	End    time.Time
	Pace   int16
}

// representation of a continuous_aggregation
// with just the important fields
type ContinuousAggregationInfo struct {
	HypertableName     string
	ViewName           string
	MaterializedOnly   bool
	CompressionEnabled bool
	Finalized          bool
}

func (c ContinuousAggregationInfo) Headers() []string {
	return []string{"HYPERTABLE_NAME", "VIEW_NAME", "MATERIALIZED_ONLY", "COMPRESSION_ENABLED", "FINALIZED"}
}

func (c ContinuousAggregationInfo) Values() []string {
	return []string{
		c.HypertableName,
		c.ViewName,
		strconv.FormatBool(c.MaterializedOnly),
		strconv.FormatBool(c.CompressionEnabled),
		strconv.FormatBool(c.Finalized),
	}
}
