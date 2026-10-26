package report

import (
	"reflect"
	"testing"

	"ddkstat/internal/logread"
)

func TestCommandsOrdering(t *testing.T) {
	records := []logread.Record{
		rec("a", "2026-10-14T01:00:00.000Z", "plan", "Toml"),
		rec("b", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		rec("c", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		rec("d", "2026-10-14T01:00:00.000Z", "alpha", "Toml"),
		rec("e", "2026-10-14T01:00:00.000Z", "alpha", "Toml"),
		rec("f", "2026-10-14T01:00:00.000Z", "Zeta", "Toml"),
		rec("g", "2026-10-14T01:00:00.000Z", "Zeta", "Toml"),
	}
	rows, _ := AggregateCommands(records)
	var got []string
	for _, r := range rows {
		got = append(got, r.Command)
	}
	if want := []string{"Zeta", "alpha", "build", "plan"}; !reflect.DeepEqual(got, want) {
		t.Errorf("commands = %v, want %v", got, want)
	}
}

func TestCommandsTotal(t *testing.T) {
	rows := readCSV(t, WriteCommands, []logread.Record{
		rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		rec("b", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		asFailed(rec("c", "2026-10-14T01:00:00.000Z", "plan", "Toml")),
	}, 0)
	total := rows[len(rows)-1]
	wantTotal := []string{"TOTAL", "3", "1", "3", "6", "9", "3.0"}
	if !reflect.DeepEqual(total, wantTotal) {
		t.Errorf("TOTAL row = %v, want %v", total, wantTotal)
	}
}

func TestCommandsEmpty(t *testing.T) {
	rows := readCSV(t, WriteCommands, nil, 0)
	wantRows(t, rows, [][]string{
		{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"},
		{"TOTAL", "0", "0", "0", "0", "0", "0.0"},
	})
}

func TestCommandsWrite(t *testing.T) {
	rows := readCSV(t, WriteCommands, []logread.Record{
		rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		asFailed(rec("b", "2026-10-14T02:00:00.000Z", "plan", "Toml")),
	}, 0)
	wantRows(t, rows, [][]string{
		{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"},
		{"build", "1", "0", "1", "2", "3", "1.0"},
		{"plan", "1", "1", "1", "2", "3", "1.0"},
		{"TOTAL", "2", "1", "2", "4", "6", "2.0"},
	})
}
