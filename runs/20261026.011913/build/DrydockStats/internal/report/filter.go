package report

import "ddkstat/internal/logread"

// FilterTarget returns the records whose Target equals name exactly, byte for byte and
// case-sensitively. It never trims, prefixes, or pattern-matches. The input slice is not modified.
func FilterTarget(records []logread.Record, name string) []logread.Record {
	kept := make([]logread.Record, 0, len(records))
	for _, r := range records {
		if r.Target == name {
			kept = append(kept, r)
		}
	}
	return kept
}

// FilterDates returns the records whose start date lies within the inclusive since and until
// bounds. Dates are YYYY-MM-DD, so string comparison is chronological. An empty bound is open on
// that side. The input slice is not modified.
func FilterDates(records []logread.Record, since, until string) []logread.Record {
	kept := make([]logread.Record, 0, len(records))
	for _, r := range records {
		date := r.Date()
		if since != "" && date < since {
			continue
		}
		if until != "" && date > until {
			continue
		}
		kept = append(kept, r)
	}
	return kept
}
