package internal

type Hypertable struct {
	HypertableSchema   string
	HYpertableName     string
	Owner              string
	NumDimensions      int64
	NumChunks          int64
	CompressionEnabled bool
}

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
