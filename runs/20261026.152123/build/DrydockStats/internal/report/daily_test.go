package report

import (
	"reflect"
	"testing"

	"ddkstat/internal/logread"
)

// dailyFixture spans two start dates. The 2026-10-15 group holds one failure and one success.
func dailyFixture() []logread.Record {
	return []logread.Record{
		rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		asFailed(rec("b", "2026-10-15T01:00:00.000Z", "build", "Toml")),
		rec("c", "2026-10-15T02:00:00.000Z", "build", "Toml"),
	}
}

func TestDailyOrdering(t *testing.T) {
	records := []logread.Record{
		rec("a", "2026-10-16T01:00:00.000Z", "build", "Toml"),
		rec("b", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		rec("c", "2026-10-15T01:00:00.000Z", "build", "Toml"),
	}
	rows, _ := AggregateDaily(records)
	var got []string
	for _, r := range rows {
		got = append(got, r.Date)
	}
	if want := []string{"2026-10-14", "2026-10-15", "2026-10-16"}; !reflect.DeepEqual(got, want) {
		t.Errorf("dates = %v, want %v", got, want)
	}
}

func TestDailyTokensSum(t *testing.T) {
	a := rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml")
	a.InputTokens, a.CachedTokens, a.OutputTokens = 10, 20, 30
	b := rec("b", "2026-10-14T02:00:00.000Z", "build", "Toml")
	b.InputTokens, b.CachedTokens, b.OutputTokens = 1, 2, 3
	rows := readCSV(t, WriteDaily, []logread.Record{a, b}, 0)
	if got := rows[1][3]; got != "66" {
		t.Errorf("tokens = %q, want 66", got)
	}
}

func TestDailyTotal(t *testing.T) {
	rows := readCSV(t, WriteDaily, dailyFixture(), 0)
	total := rows[len(rows)-1]
	wantTotal := []string{"TOTAL", "3", "1", "18", "3.0"}
	if !reflect.DeepEqual(total, wantTotal) {
		t.Errorf("TOTAL row = %v, want %v", total, wantTotal)
	}
}

func TestDailyEmpty(t *testing.T) {
	rows := readCSV(t, WriteDaily, nil, 0)
	wantRows(t, rows, [][]string{
		{"date", "calls", "failed", "tokens", "minutes"},
		{"TOTAL", "0", "0", "0", "0.0"},
	})
}

func TestDailyWrite(t *testing.T) {
	rows := readCSV(t, WriteDaily, dailyFixture(), 0)
	wantRows(t, rows, [][]string{
		{"date", "calls", "failed", "tokens", "minutes"},
		{"2026-10-14", "1", "0", "6", "1.0"},
		{"2026-10-15", "2", "1", "12", "2.0"},
		{"TOTAL", "3", "1", "18", "3.0"},
	})
}
