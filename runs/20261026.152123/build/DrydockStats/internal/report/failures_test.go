package report

import (
	"reflect"
	"testing"

	"ddkstat/internal/logread"
)

var failuresHeader = []string{"started_at", "target", "command", "model", "returncode", "timed_out"}

func TestFailuresOrdering(t *testing.T) {
	late := asFailed(rec("z", "2026-10-15T01:00:00.000Z", "build", "Toml"))
	tieB := asFailed(rec("m", "2026-10-14T05:00:00.000Z", "build", "Toml"))
	tieA := asFailed(rec("a", "2026-10-14T05:00:00.000Z", "build", "Toml"))
	var got []string
	for _, r := range SelectFailures([]logread.Record{late, tieB, tieA}) {
		got = append(got, r.ExecutionID)
	}
	if want := []string{"a", "m", "z"}; !reflect.DeepEqual(got, want) {
		t.Errorf("failures = %v, want %v", got, want)
	}
}

func TestFailuresNullReturncodeEmpty(t *testing.T) {
	null := asFailed(rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml"))
	null.ReturnCode = nil
	coded := asFailed(rec("b", "2026-10-14T02:00:00.000Z", "build", "Toml"))
	coded.ReturnCode = intPtr(2)
	rows := readCSV(t, WriteFailures, []logread.Record{null, coded}, 0)
	wantRows(t, rows, [][]string{
		failuresHeader,
		{"2026-10-14T01:00:00.000Z", "Toml", "build", "claude-haiku-5-5", "", "false"},
		{"2026-10-14T02:00:00.000Z", "Toml", "build", "claude-haiku-5-5", "2", "false"},
	})
}

func TestFailuresTimedOutText(t *testing.T) {
	timedOut := asFailed(rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml"))
	timedOut.TimedOut = true
	rows := readCSV(t, WriteFailures, []logread.Record{timedOut}, 0)
	if got := rows[1][5]; got != "true" {
		t.Errorf("timed_out = %q, want true", got)
	}
	rows = readCSV(t, WriteFailures, []logread.Record{asFailed(rec("b", "2026-10-14T01:00:00.000Z", "build", "Toml"))}, 0)
	if got := rows[1][5]; got != "false" {
		t.Errorf("timed_out = %q, want false", got)
	}
}

func TestFailuresOnlyFailed(t *testing.T) {
	records := []logread.Record{
		rec("ok", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		asFailed(rec("bad", "2026-10-14T02:00:00.000Z", "build", "Toml")),
	}
	var got []string
	for _, r := range SelectFailures(records) {
		got = append(got, r.ExecutionID)
	}
	if want := []string{"bad"}; !reflect.DeepEqual(got, want) {
		t.Errorf("failures = %v, want %v", got, want)
	}
}

func TestFailuresEmpty(t *testing.T) {
	rows := readCSV(t, WriteFailures, []logread.Record{rec("ok", "2026-10-14T01:00:00.000Z", "build", "Toml")}, 0)
	wantRows(t, rows, [][]string{failuresHeader})
}
