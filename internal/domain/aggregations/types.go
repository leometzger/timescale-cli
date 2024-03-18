package aggregations

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

// representation of a continuous_aggregation
// with just the important fields
type ContinuousAggregationInfo struct {
	ViewName           string `header:"VIEW_NAME"`
	MaterializedOnly   bool   `header:"MATERIALIZED_ONLY"`
	CompressionEnabled bool   `header:"COMPRESSION_ENABLED"`
	Finalized          bool   `header:"FINALIZED"`
}
