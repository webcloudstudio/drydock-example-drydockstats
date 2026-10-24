// Package filter narrows parsed log records by Target name and by an inclusive start-date range.
package filter

import (
	"regexp"
	"time"

	"ddkstat/internal/log"
)

var dateForm = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

// ValidDate reports whether s is a real calendar date written exactly YYYY-MM-DD.
func ValidDate(s string) bool {
	if !dateForm.MatchString(s) {
		return false
	}
	_, err := time.Parse(time.DateOnly, s)
	return err == nil
}

// Options are the optional --target, --since, and --until filters. An empty Since or Until is
// absent. HasTarget distinguishes an absent --target from an empty one.
type Options struct {
	Target    string
	HasTarget bool
	Since     string
	Until     string
}

// Match reports whether r passes every filter that is set.
func (o Options) Match(r log.Record) bool {
	if o.HasTarget && r.Target != o.Target {
		return false
	}
	if o.Since != "" && r.StartDate < o.Since {
		return false
	}
	if o.Until != "" && r.StartDate > o.Until {
		return false
	}
	return true
}

// Apply returns the records that pass Match, in their original order.
func Apply(o Options, recs []log.Record) []log.Record {
	var kept []log.Record
	for _, r := range recs {
		if o.Match(r) {
			kept = append(kept, r)
		}
	}
	return kept
}
