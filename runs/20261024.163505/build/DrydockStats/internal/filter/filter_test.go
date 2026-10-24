package filter

import (
	"testing"

	"ddkstat/internal/log"
)

func TestValidDate(t *testing.T) {
	tests := map[string]bool{
		"2026-10-14":       true,
		"2024-02-29":       true,  // leap year
		"2026-02-29":       false, // not a leap year
		"2026-02-30":       false,
		"2026-13-01":       false,
		"2026-00-10":       false,
		"2026-10-00":       false,
		"2026-1-01":        false,
		"2026/10/14":       false,
		"20261014":         false,
		"2026-10-14T00:00": false,
		"":                 false,
	}
	for in, want := range tests {
		if got := ValidDate(in); got != want {
			t.Errorf("ValidDate(%q) = %v, want %v", in, got, want)
		}
	}
}

func rec(date, target string) log.Record {
	return log.Record{StartDate: date, Target: target}
}

func TestMatchTargetIsExactAndOptional(t *testing.T) {
	alpha := rec("2026-10-14", "Alpha")
	if !(Options{}).Match(alpha) {
		t.Error("no --target should keep every record")
	}
	if !(Options{Target: "Alpha", HasTarget: true}).Match(alpha) {
		t.Error("exact --target should keep the record")
	}
	for _, other := range []string{"alpha", "Alpha ", "Alph"} {
		if (Options{Target: other, HasTarget: true}).Match(alpha) {
			t.Errorf("--target %q matched Alpha; target comparison must be exact", other)
		}
	}
	if (Options{Target: "", HasTarget: true}).Match(alpha) {
		t.Error("empty --target matched a non-empty target")
	}
}

func TestMatchDateRangeIsInclusive(t *testing.T) {
	o := Options{Since: "2026-10-14", Until: "2026-10-16"}
	tests := map[string]bool{
		"2026-10-13": false,
		"2026-10-14": true,
		"2026-10-15": true,
		"2026-10-16": true,
		"2026-10-17": false,
	}
	for date, want := range tests {
		if got := o.Match(rec(date, "T")); got != want {
			t.Errorf("date %s: Match = %v, want %v", date, got, want)
		}
	}
}

func TestInvertedRangeMatchesNothing(t *testing.T) {
	o := Options{Since: "2026-10-16", Until: "2026-10-14"}
	if got := Apply(o, []log.Record{rec("2026-10-15", "T")}); len(got) != 0 {
		t.Errorf("inverted range kept %d records, want 0", len(got))
	}
}

func TestApplyKeepsOrderAndReturnsNilWhenNothingMatches(t *testing.T) {
	recs := []log.Record{rec("2026-10-14", "A"), rec("2026-10-15", "B"), rec("2026-10-16", "A")}
	got := Apply(Options{Target: "A", HasTarget: true}, recs)
	if len(got) != 2 || got[0].StartDate != "2026-10-14" || got[1].StartDate != "2026-10-16" {
		t.Errorf("Apply = %+v, want the two A records in order", got)
	}
	if none := Apply(Options{Target: "Z", HasTarget: true}, recs); len(none) != 0 {
		t.Errorf("Apply with no match = %+v, want empty", none)
	}
}
