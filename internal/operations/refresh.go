package operations

import (
	"time"

	"github.com/leometzger/timescale-cli/internal/db"
)

type Refresher interface {
	Refresh(agg db.ContinuousAggregation, from time.Time, to time.Time) error
}
