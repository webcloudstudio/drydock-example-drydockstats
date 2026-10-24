package report

import (
	"reflect"
	"strings"
	"testing"

	"ddkstat/internal/log"
)

// recordsFrom parses log text the same way the CLI does, so reports are tested on real records.
func recordsFrom(t *testing.T, text string) []log.Record {
	t.Helper()
	lines, err := log.ReadLines(strings.NewReader(text))
	if err != nil {
		t.Fatal(err)
	}
	return log.Parse(lines).Records
}

func entry(id, started, status, command, target, model string, input, cached, output int64, cost, ms string) string {
	return `{"execution_id":"` + id + `","started_at":"` + started + `","status":"` + status +
		`","job":{"command_name":"` + command + `","target":"` + target + `","llm":"claude","model":"` + model +
		`"},"result":{"stats":{"input_tokens":` + itoa(input) + `,"cached_input_tokens":` + itoa(cached) +
		`,"output_tokens":` + itoa(output) + `,"cost_usd":` + cost + `,"elapsed_ms":` + ms + `}}}`
}

// fixture has two dates, two targets, two commands, and one failure.
var fixture = strings.Join([]string{
	entry("1", "2026-10-14T01:00:00.000Z", "succeeded", "build", "Beta", "haiku", 10, 100, 5, "0.25", "60000"),
	entry("2", "2026-10-14T02:00:00.000Z", "failed", "plan", "Alpha", "sonnet", 1, 2, 3, "0.5", "35000"),
	entry("3", "2026-10-15T03:00:00.000Z", "succeeded", "build", "Alpha", "haiku", 4, 0, 6, "0.0", "0"),
}, "\n")

func TestSummary(t *testing.T) {
	got := Summary(recordsFrom(t, fixture), 4)
	want := [][]string{
		{"records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"3", "4", "2", "1", "2026-10-14", "2026-10-15", "15", "102", "14", "0.7500", "1.6"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Summary = %q\nwant %q", got, want)
	}
}

func TestSummaryZeroMatch(t *testing.T) {
	got := Summary(nil, 3)
	want := [][]string{
		{"records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"0", "3", "0", "0", "", "", "0", "0", "0", "0.0000", "0.0"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Summary(nil) = %q, want %q", got, want)
	}
}

func TestCommandsOrderedByCallsThenName(t *testing.T) {
	text := strings.Join([]string{
		entry("a", "2026-10-14T01:00:00.000Z", "succeeded", "plan", "T", "m", 1, 0, 0, "0", "0"),
		entry("b", "2026-10-14T01:00:01.000Z", "succeeded", "build", "T", "m", 1, 0, 0, "0", "0"),
		entry("c", "2026-10-14T01:00:02.000Z", "succeeded", "build", "T", "m", 1, 0, 0, "0", "0"),
		entry("d", "2026-10-14T01:00:03.000Z", "succeeded", "analyze", "T", "m", 1, 0, 0, "0", "0"),
	}, "\n")
	got := Commands(recordsFrom(t, text))
	if got[1][0] != "build" || got[2][0] != "analyze" || got[3][0] != "plan" || got[4][0] != "TOTAL" {
		t.Errorf("command order = %q", got)
	}
	if got[4][1] != "4" {
		t.Errorf("TOTAL calls = %q, want 4", got[4][1])
	}
}

func TestCommandsZeroMatchWritesZeroTotal(t *testing.T) {
	got := Commands(nil)
	want := [][]string{
		{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"},
		{"TOTAL", "0", "0", "0", "0", "0", "0.0"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Commands(nil) = %q, want %q", got, want)
	}
}

func TestTargetsOrderedAscendingWithCostAndTotal(t *testing.T) {
	got := Targets(recordsFrom(t, fixture))
	want := [][]string{
		{"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"Alpha", "2", "1", "5", "2", "9", "0.5000", "0.6"},
		{"Beta", "1", "0", "10", "100", "5", "0.2500", "1.0"},
		{"TOTAL", "3", "1", "15", "102", "14", "0.7500", "1.6"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Targets = %q\nwant %q", got, want)
	}
}

func TestTargetsZeroMatch(t *testing.T) {
	got := Targets(nil)
	want := [][]string{
		{"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"TOTAL", "0", "0", "0", "0", "0", "0.0000", "0.0"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Targets(nil) = %q, want %q", got, want)
	}
}

func TestModelsGroupByProviderAndModelWithEmptyTotalCells(t *testing.T) {
	got := Models(recordsFrom(t, fixture))
	want := [][]string{
		{"provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"},
		{"claude", "haiku", "2", "0", "14", "100", "11", "0.2500"},
		{"claude", "sonnet", "1", "1", "1", "2", "3", "0.5000"},
		{"TOTAL", "", "3", "1", "15", "102", "14", "0.7500"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Models = %q\nwant %q", got, want)
	}
}

func TestModelsUsesExecutedModelWhenPresent(t *testing.T) {
	text := `{"execution_id":"x","started_at":"2026-10-14T00:00:00Z","status":"succeeded","job":{"command_name":"b","target":"T","llm":"gemini","model":"requested"},"result":{"stats":{"model":"executed"}}}`
	got := Models(recordsFrom(t, text))
	if got[1][0] != "gemini" || got[1][1] != "executed" {
		t.Errorf("row = %q, want gemini/executed", got[1])
	}
}

func TestDailyOrderedByDateWithTokensAndTotal(t *testing.T) {
	got := Daily(recordsFrom(t, fixture))
	want := [][]string{
		{"date", "calls", "failed", "tokens", "minutes"},
		{"2026-10-14", "2", "1", "121", "1.6"},
		{"2026-10-15", "1", "0", "10", "0.0"},
		{"TOTAL", "3", "1", "131", "1.6"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Daily = %q\nwant %q", got, want)
	}
}

func TestDailyZeroMatch(t *testing.T) {
	got := Daily(nil)
	want := [][]string{
		{"date", "calls", "failed", "tokens", "minutes"},
		{"TOTAL", "0", "0", "0", "0.0"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Daily(nil) = %q, want %q", got, want)
	}
}

func TestFailuresListsOnlyFailedRecordsWithoutTotal(t *testing.T) {
	text := strings.Join([]string{
		`{"execution_id":"z","started_at":"2026-10-14T05:00:00.000Z","status":"failed","job":{"command_name":"diagnose","target":"T","llm":"claude","model":"sonnet"},"result":{"returncode":null,"timed_out":true}}`,
		`{"execution_id":"y","started_at":"2026-10-14T05:00:00.000Z","status":"failed","job":{"command_name":"diagnose","target":"T","llm":"claude","model":"haiku"},"result":{"returncode":2,"timed_out":false}}`,
		entry("ok", "2026-10-13T00:00:00.000Z", "succeeded", "build", "T", "haiku", 0, 0, 0, "0", "0"),
	}, "\n")
	got := Failures(recordsFrom(t, text))
	want := [][]string{
		{"started_at", "target", "command", "model", "returncode", "timed_out"},
		{"2026-10-14T05:00:00.000Z", "T", "diagnose", "haiku", "2", "false"},
		{"2026-10-14T05:00:00.000Z", "T", "diagnose", "sonnet", "", "true"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Failures = %q\nwant %q", got, want)
	}
}

func TestFailuresHeaderOnlyWhenNoneFailed(t *testing.T) {
	got := Failures(nil)
	if len(got) != 1 {
		t.Errorf("Failures(nil) = %q, want header only", got)
	}
}
