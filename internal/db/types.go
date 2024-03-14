package db

import "strconv"

type Hypertable struct {
	HypertableSchema   string
	HYpertableName     string
	Owner              string
	NumDimensions      int64
	NumChunks          int64
	CompressionEnabled bool
}

type HypertableInfo struct {
	HypertableName     string `present:"HYPERTABLE"`
	NumChunks          int64  `present:"CHUNKS"`
	CompressionEnabled bool   `present:"COMPRESSION ENABLED"`
	Size               int64  `present:"SIZE"`
}

type Chunk struct{}

func (h HypertableInfo) Headers() []string {
	return []string{
		"HypertableName",
		"NumChunks",
		"CompressionEnabled",
		"Size",
	}
}

func (h HypertableInfo) Values() []string {
	return []string{
		h.HypertableName,
		strconv.FormatInt(h.NumChunks, 10),
		strconv.FormatBool(h.CompressionEnabled),
		strconv.FormatInt(h.Size, 10),
	}
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

type ContinuousAggregationInfo struct {
	ViewName           string
	MaterializedOnly   bool
	CompressionEnabled bool
	Finalized          bool
}

func (c *ContinuousAggregationInfo) Headers() []string {
	return []string{
		"ViewName",
		"MaterializedOnly",
		"CompressionEnabled",
		"Finalized",
	}
}

func (c *ContinuousAggregationInfo) Values() []string {
	return []string{
		c.ViewName,
		strconv.FormatBool(c.MaterializedOnly),
		strconv.FormatBool(c.CompressionEnabled),
		strconv.FormatBool(c.Finalized),
	}
}
