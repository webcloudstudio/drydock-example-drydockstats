package report

import (
	"testing"

	"ddkstat/internal/logread"
)

// rec returns a Record with the given id, target, and start date, for filter tests.
func rec(id, target, startedAt string) logread.Record {
	return logread.Record{ExecutionID: id, Target: target, StartedAt: startedAt}
}

// ids returns the ExecutionID of each record in order, so tests compare readable values.
func ids(records []logread.Record) []string {
	out := make([]string, 0, len(records))
	for _, r := range records {
		out = append(out, r.ExecutionID)
	}
	return out
}

// equal reports whether two string slices hold the same values in the same order.
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFilterTargetKeepsExactMatches(t *testing.T) {
	in := []logread.Record{
		rec("1", "Alpha", "2026-10-14T01:00:00.000Z"),
		rec("2", "Beta", "2026-10-14T01:00:00.000Z"),
		rec("3", "Alpha", "2026-10-15T01:00:00.000Z"),
	}
	got := ids(FilterTarget(in, "Alpha"))
	if want := []string{"1", "3"}; !equal(got, want) {
		t.Errorf("FilterTarget ids = %v, want %v", got, want)
	}
}

func TestFilterTargetRejectsNearMisses(t *testing.T) {
	names := []string{"alpha", "Alpha2", "Alph", "ALPHA", " Alpha", "Alpha "}
	var in []logread.Record
	for i, n := range names {
		in = append(in, rec(string(rune('a'+i)), n, "2026-10-14T01:00:00.000Z"))
	}
	if got := FilterTarget(in, "Alpha"); len(got) != 0 {
		t.Errorf("FilterTarget kept %d near-miss records, want 0: %v", len(got), ids(got))
	}
}

func TestFilterTargetEmptyInput(t *testing.T) {
	if got := FilterTarget(nil, "Alpha"); len(got) != 0 {
		t.Errorf("FilterTarget(nil) returned %d records, want 0", len(got))
	}
}

func TestFilterTargetNoMatchReturnsEmpty(t *testing.T) {
	in := []logread.Record{rec("1", "Alpha", "2026-10-14T01:00:00.000Z")}
	if got := FilterTarget(in, "NoSuchTarget"); len(got) != 0 {
		t.Errorf("FilterTarget kept %d records, want 0", len(got))
	}
}

func TestFilterTargetDoesNotModifyInput(t *testing.T) {
	in := []logread.Record{
		rec("1", "Beta", "2026-10-14T01:00:00.000Z"),
		rec("2", "Alpha", "2026-10-14T01:00:00.000Z"),
	}
	FilterTarget(in, "Alpha")
	if got := ids(in); !equal(got, []string{"1", "2"}) || in[0].Target != "Beta" {
		t.Errorf("input changed after FilterTarget: %v", got)
	}
}

func TestFilterDatesBothBoundsInclusive(t *testing.T) {
	days := []string{"2026-10-13", "2026-10-14", "2026-10-15", "2026-10-16", "2026-10-17"}
	var in []logread.Record
	for i, d := range days {
		in = append(in, rec(string(rune('a'+i)), "T", d+"T12:00:00.000Z"))
	}
	got := ids(FilterDates(in, "2026-10-14", "2026-10-16"))
	if want := []string{"b", "c", "d"}; !equal(got, want) {
		t.Errorf("FilterDates ids = %v, want %v", got, want)
	}
}

func TestFilterDatesOneSidedBounds(t *testing.T) {
	in := []logread.Record{
		rec("a", "T", "2026-10-13T00:00:00.000Z"),
		rec("b", "T", "2026-10-14T00:00:00.000Z"),
		rec("c", "T", "2026-10-15T00:00:00.000Z"),
	}
	if got := ids(FilterDates(in, "2026-10-14", "")); !equal(got, []string{"b", "c"}) {
		t.Errorf("since only: ids = %v, want [b c]", got)
	}
	if got := ids(FilterDates(in, "", "2026-10-14")); !equal(got, []string{"a", "b"}) {
		t.Errorf("until only: ids = %v, want [a b]", got)
	}
	if got := ids(FilterDates(in, "", "")); !equal(got, []string{"a", "b", "c"}) {
		t.Errorf("no bounds: ids = %v, want [a b c]", got)
	}
}

func TestFilterDatesIgnoresTimeOfDay(t *testing.T) {
	in := []logread.Record{
		rec("a", "T", "2026-10-15T00:00:00.000Z"),
		rec("b", "T", "2026-10-15T23:59:59.999Z"),
		rec("c", "T", "2026-10-14T23:59:59.999Z"),
		rec("d", "T", "2026-10-16T00:00:00.000Z"),
	}
	if got := ids(FilterDates(in, "2026-10-15", "2026-10-15")); !equal(got, []string{"a", "b"}) {
		t.Errorf("same-day range: ids = %v, want [a b]", got)
	}
}

func TestFilterDatesEmptyRangeReturnsEmpty(t *testing.T) {
	in := []logread.Record{rec("a", "T", "2026-10-14T00:00:00.000Z")}
	if got := FilterDates(in, "2030-01-01", ""); len(got) != 0 {
		t.Errorf("FilterDates kept %d records, want 0", len(got))
	}
}

func TestFilterDatesEmptyInput(t *testing.T) {
	if got := FilterDates(nil, "2026-10-14", "2026-10-16"); len(got) != 0 {
		t.Errorf("FilterDates(nil) returned %d records, want 0", len(got))
	}
}

func TestFilterDatesDoesNotModifyInput(t *testing.T) {
	in := []logread.Record{
		rec("a", "T", "2026-10-13T00:00:00.000Z"),
		rec("b", "T", "2026-10-14T00:00:00.000Z"),
	}
	FilterDates(in, "2026-10-14", "")
	if got := ids(in); !equal(got, []string{"a", "b"}) {
		t.Errorf("input changed after FilterDates: %v", got)
	}
}

func TestFiltersCompose(t *testing.T) {
	in := []logread.Record{
		rec("a", "Alpha", "2026-10-15T01:00:00.000Z"),
		rec("b", "Beta", "2026-10-15T01:00:00.000Z"),
		rec("c", "Alpha", "2026-10-16T01:00:00.000Z"),
		rec("d", "Alpha", "2026-10-17T01:00:00.000Z"),
	}
	got := ids(FilterDates(FilterTarget(in, "Alpha"), "2026-10-15", "2026-10-16"))
	if !equal(got, []string{"a", "c"}) {
		t.Errorf("combined filters: ids = %v, want [a c]", got)
	}
}
