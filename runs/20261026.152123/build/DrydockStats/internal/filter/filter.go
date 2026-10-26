// Package filter narrows a slice of logread.Record by Target and by start date.
//
// Every function here is pure: it performs no I/O, keeps no global state, and never modifies its
// input slice. Input order is preserved.
package filter

import "ddkstat/internal/logread"

// ByTarget returns the records whose Target equals name exactly. Comparison is case-sensitive.
// An empty name means no --target was given, so every record is returned. A name that matches no
// record yields an empty slice, never an error.
func ByTarget(records []logread.Record, name string) []logread.Record {
	out := make([]logread.Record, 0, len(records))
	for _, rec := range records {
		if name == "" || rec.Target == name {
			out = append(out, rec)
		}
	}
	return out
}

// ByDateRange returns the records whose Date lies within the inclusive range [since, until].
// Dates are validated YYYY-MM-DD text, so string comparison orders them correctly. An empty bound
// is unset and does not restrict. When since is later than until the result is empty; that is not
// an error.
func ByDateRange(records []logread.Record, since, until string) []logread.Record {
	out := make([]logread.Record, 0, len(records))
	for _, rec := range records {
		if since != "" && rec.Date < since {
			continue
		}
		if until != "" && rec.Date > until {
			continue
		}
		out = append(out, rec)
	}
	return out
}
