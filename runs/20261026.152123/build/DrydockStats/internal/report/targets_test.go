package report

import (
	"reflect"
	"testing"

	"ddkstat/internal/logread"
)

func TestTargetsOrdering(t *testing.T) {
	records := []logread.Record{
		rec("a", "2026-10-14T01:00:00.000Z", "build", "beta"),
		rec("b", "2026-10-14T02:00:00.000Z", "build", "Alpha"),
		rec("c", "2026-10-14T03:00:00.000Z", "build", "alpha"),
	}
	rows, _ := AggregateTargets(records)
	var got []string
	for _, r := range rows {
		got = append(got, r.Target)
	}
	if want := []string{"Alpha", "alpha", "beta"}; !reflect.DeepEqual(got, want) {
		t.Errorf("targets = %v, want %v", got, want)
	}
}

func TestTargetsTotalCost(t *testing.T) {
	a := rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml")
	a.CostUSD = 0.1
	b := rec("b", "2026-10-14T02:00:00.000Z", "build", "Toml")
	b.CostUSD = 0.2
	rows := readCSV(t, WriteTargets, []logread.Record{a, b}, 0)
	total := rows[len(rows)-1]
	if total[0] != "TOTAL" || total[6] != "0.3000" {
		t.Errorf("TOTAL row = %v, want cost 0.3000", total)
	}
}

func TestTargetsEmpty(t *testing.T) {
	rows := readCSV(t, WriteTargets, nil, 0)
	wantRows(t, rows, [][]string{
		{"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"TOTAL", "0", "0", "0", "0", "0", "0.0000", "0.0"},
	})
}

func TestTargetsWrite(t *testing.T) {
	rows := readCSV(t, WriteTargets, []logread.Record{
		rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		rec("b", "2026-10-14T02:00:00.000Z", "plan", "CommonMark"),
	}, 0)
	wantRows(t, rows, [][]string{
		{"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"CommonMark", "1", "0", "1", "2", "3", "0.5000", "1.0"},
		{"Toml", "1", "0", "1", "2", "3", "0.5000", "1.0"},
		{"TOTAL", "2", "0", "2", "4", "6", "1.0000", "2.0"},
	})
}
