package report

import "drydockstats/internal/logrecord"

// Filter narrows records by Target and by an inclusive start-date range.
type Filter struct {
	Target    string
	HasTarget bool
	Since     string // YYYY-MM-DD; empty leaves the range open below
	Until     string // YYYY-MM-DD; empty leaves the range open above
}

// Apply returns the records that pass every filter, in their original order.
func Apply(records []logrecord.Record, f Filter) []logrecord.Record {
	kept := make([]logrecord.Record, 0, len(records))
	for _, r := range records {
		if f.HasTarget && r.Target != f.Target {
			continue
		}
		date := r.StartDate()
		if f.Since != "" && date < f.Since {
			continue
		}
		if f.Until != "" && date > f.Until {
			continue
		}
		kept = append(kept, r)
	}
	return kept
}
