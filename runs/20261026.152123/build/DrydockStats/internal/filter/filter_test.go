package filter_test

import (
	"reflect"
	"testing"

	"ddkstat/internal/filter"
	"ddkstat/internal/logread"
)

// rec builds a Record with the fields the filters read: an identifier, a Target and a Date.
func rec(id, target, date string) logread.Record {
	return logread.Record{ExecutionID: id, Target: target, Date: date}
}

// ids returns the ExecutionIDs of records in order, so tests compare returned state, not pointers.
func ids(records []logread.Record) []string {
	out := []string{}
	for _, r := range records {
		out = append(out, r.ExecutionID)
	}
	return out
}

func sample() []logread.Record {
	return []logread.Record{
		rec("a", "Toml", "2026-10-14"),
		rec("b", "CommonMark", "2026-10-15"),
		rec("c", "toml", "2026-10-15"),
		rec("d", "ReadingList", "2026-10-16"),
		rec("e", "Toml", "2026-10-16"),
	}
}

func TestByTargetExact(t *testing.T) {
	got := ids(filter.ByTarget(sample(), "Toml"))
	want := []string{"a", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ids = %v, want %v", got, want)
	}
}

func TestByTargetCaseSensitive(t *testing.T) {
	got := ids(filter.ByTarget(sample(), "TOML"))
	if len(got) != 0 {
		t.Errorf("ids = %v, want none (case-sensitive match)", got)
	}
	got = ids(filter.ByTarget(sample(), "toml"))
	want := []string{"c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ids = %v, want %v", got, want)
	}
}

func TestByTargetNoMatch(t *testing.T) {
	got := filter.ByTarget(sample(), "NoSuchTarget")
	if got == nil || len(got) != 0 {
		t.Errorf("got %v, want an empty non-nil slice", got)
	}
}

func TestByTargetUnset(t *testing.T) {
	got := ids(filter.ByTarget(sample(), ""))
	want := []string{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ids = %v, want %v", got, want)
	}
}

func TestByTargetDoesNotModifyInput(t *testing.T) {
	in := sample()
	filter.ByTarget(in, "Toml")
	if got, want := ids(in), []string{"a", "b", "c", "d", "e"}; !reflect.DeepEqual(got, want) {
		t.Errorf("input ids = %v, want %v", got, want)
	}
}

func TestSinceInclusive(t *testing.T) {
	got := ids(filter.ByDateRange(sample(), "2026-10-15", ""))
	want := []string{"b", "c", "d", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ids = %v, want %v", got, want)
	}
}

func TestUntilInclusive(t *testing.T) {
	got := ids(filter.ByDateRange(sample(), "", "2026-10-15"))
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ids = %v, want %v", got, want)
	}
}

func TestRangeBothBounds(t *testing.T) {
	got := ids(filter.ByDateRange(sample(), "2026-10-15", "2026-10-15"))
	want := []string{"b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ids = %v, want %v", got, want)
	}
}

func TestSinceAfterUntilEmpty(t *testing.T) {
	got := filter.ByDateRange(sample(), "2026-10-16", "2026-10-14")
	if got == nil || len(got) != 0 {
		t.Errorf("got %v, want an empty non-nil slice", got)
	}
}

func TestBoundsUnset(t *testing.T) {
	got := ids(filter.ByDateRange(sample(), "", ""))
	want := []string{"a", "b", "c", "d", "e"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ids = %v, want %v", got, want)
	}
}

func TestByDateRangeEmptyInput(t *testing.T) {
	got := filter.ByDateRange(nil, "2026-10-14", "2026-10-16")
	if len(got) != 0 {
		t.Errorf("got %d records, want 0", len(got))
	}
}

func TestByDateRangeDoesNotModifyInput(t *testing.T) {
	in := sample()
	filter.ByDateRange(in, "2026-10-15", "2026-10-15")
	if got, want := ids(in), []string{"a", "b", "c", "d", "e"}; !reflect.DeepEqual(got, want) {
		t.Errorf("input ids = %v, want %v", got, want)
	}
}
