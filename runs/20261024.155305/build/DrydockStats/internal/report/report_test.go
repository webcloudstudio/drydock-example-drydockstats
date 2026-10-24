package report

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"strings"
	"testing"

	"drydockstats/internal/logrecord"
)

// rec builds a record with the given identity and numeric values.
func rec(id, target, command, started string, status string, in, cached, out, elapsed int64, cost float64) logrecord.Record {
	return logrecord.Record{
		ExecutionID: id, StartedAt: started, Status: status, Target: target, Command: command,
		Provider: "claude", Model: "haiku",
		InputTokens: in, CachedTokens: cached, OutputTokens: out, ElapsedMS: elapsed, CostUSD: cost,
	}
}

func TestFormatMinutesRoundsHalfUpToTenths(t *testing.T) {
	cases := map[int64]string{
		0:      "0.0",
		2999:   "0.0",
		3000:   "0.1", // (3000+3000)/6000 = 1
		95000:  "1.6", // 98000/6000 = 16
		60000:  "1.0",
		600000: "10.0",
	}
	for ms, want := range cases {
		if got := FormatMinutes(ms); got != want {
			t.Errorf("FormatMinutes(%d) = %q, want %q", ms, got, want)
		}
	}
}

func TestFormatCostHasFourPlaces(t *testing.T) {
	cases := map[float64]string{
		0:      "0.0000",
		5.5245: "5.5245",
		0.1:    "0.1000",
	}
	for usd, want := range cases {
		if got := FormatCost(usd); got != want {
			t.Errorf("FormatCost(%v) = %q, want %q", usd, got, want)
		}
	}
}

func TestWriteCSVRoundTripsRowsWithLFEndings(t *testing.T) {
	rows := [][]string{{"a", "b c", ""}, {"1", "2,3", `say "hi"`}}
	var buf bytes.Buffer
	if err := WriteCSV(&buf, rows); err != nil {
		t.Fatalf("WriteCSV: %v", err)
	}
	if strings.Contains(buf.String(), "\r") {
		t.Errorf("output contains CR; want LF endings only")
	}
	got, err := csv.NewReader(&buf).ReadAll()
	if err != nil {
		t.Fatalf("read back: %v", err)
	}
	if !reflect.DeepEqual(got, rows) {
		t.Errorf("round trip = %q, want %q", got, rows)
	}
}

func TestApplyTargetMatchesExactlyAndCaseSensitively(t *testing.T) {
	records := []logrecord.Record{
		{ExecutionID: "1", Target: "Toml", StartedAt: "2026-10-14"},
		{ExecutionID: "2", Target: "toml", StartedAt: "2026-10-14"},
		{ExecutionID: "3", Target: "Toml ", StartedAt: "2026-10-14"},
	}
	kept := Apply(records, Filter{Target: "Toml", HasTarget: true})
	if len(kept) != 1 || kept[0].ExecutionID != "1" {
		t.Errorf("kept = %+v, want only execution 1", kept)
	}
}

func TestApplyWithoutTargetKeepsEveryRecord(t *testing.T) {
	records := []logrecord.Record{{Target: "A", StartedAt: "2026-10-14"}, {Target: "B", StartedAt: "2026-10-14"}}
	if got := Apply(records, Filter{}); len(got) != 2 {
		t.Errorf("kept = %d, want 2", len(got))
	}
}

func TestApplyDateRangeIsInclusiveOnBothBounds(t *testing.T) {
	dates := []string{"2026-10-13", "2026-10-14", "2026-10-15", "2026-10-16", "2026-10-17"}
	var records []logrecord.Record
	for i, d := range dates {
		records = append(records, logrecord.Record{ExecutionID: string(rune('a' + i)), StartedAt: d + "T00:00:00.000Z"})
	}
	got := Apply(records, Filter{Since: "2026-10-14", Until: "2026-10-16"})
	var kept []string
	for _, r := range got {
		kept = append(kept, r.StartDate())
	}
	want := []string{"2026-10-14", "2026-10-15", "2026-10-16"}
	if !reflect.DeepEqual(kept, want) {
		t.Errorf("kept dates = %v, want %v", kept, want)
	}
}

func TestApplyOpenEndedRanges(t *testing.T) {
	records := []logrecord.Record{{ExecutionID: "old", StartedAt: "2026-10-13T00:00:00.000Z"},
		{ExecutionID: "new", StartedAt: "2026-10-20T00:00:00.000Z"}}
	if got := Apply(records, Filter{Since: "2026-10-14"}); len(got) != 1 || got[0].ExecutionID != "new" {
		t.Errorf("since only kept %+v, want new", got)
	}
	if got := Apply(records, Filter{Until: "2026-10-14"}); len(got) != 1 || got[0].ExecutionID != "old" {
		t.Errorf("until only kept %+v, want old", got)
	}
}

func TestSummaryCountsFilteredAndUsesWholeFileSkipped(t *testing.T) {
	records := []logrecord.Record{
		rec("1", "Alpha", "build", "2026-10-14T01:00:00.000Z", "succeeded", 10, 5, 2, 60000, 0.5),
		rec("2", "Alpha", "build", "2026-10-16T01:00:00.000Z", "failed", 3, 0, 1, 0, 0.25),
	}
	rows := Summary(records, 7)
	want := []string{"2", "7", "1", "1", "2026-10-14", "2026-10-16", "13", "5", "3", "0.7500", "1.0"}
	if !reflect.DeepEqual(rows[1], want) {
		t.Errorf("summary row = %v, want %v", rows[1], want)
	}
	if len(rows) != 2 {
		t.Errorf("summary rows = %d, want header plus one", len(rows))
	}
}

func TestSummaryEmptyMatchHasEmptyFirstLast(t *testing.T) {
	rows := Summary(nil, 3)
	want := []string{"0", "3", "0", "0", "", "", "0", "0", "0", "0.0000", "0.0"}
	if !reflect.DeepEqual(rows[1], want) {
		t.Errorf("empty summary row = %v, want %v", rows[1], want)
	}
}

func TestCommandsOrdersByCallsThenNameAndAddsTotal(t *testing.T) {
	records := []logrecord.Record{
		rec("1", "T", "plan", "2026-10-14T01:00:00.000Z", "succeeded", 1, 0, 0, 0, 0),
		rec("2", "T", "build", "2026-10-14T01:00:00.000Z", "succeeded", 1, 0, 0, 0, 0),
		rec("3", "T", "build", "2026-10-14T01:00:00.000Z", "succeeded", 1, 0, 0, 0, 0),
		rec("4", "T", "analyze", "2026-10-14T01:00:00.000Z", "failed", 1, 0, 0, 0, 0),
		rec("5", "T", "analyze", "2026-10-14T01:00:00.000Z", "succeeded", 1, 0, 0, 0, 0),
		rec("6", "T", "import_compass", "2026-10-14T01:00:00.000Z", "succeeded", 1, 0, 0, 0, 0),
		rec("7", "T", "plan repair", "2026-10-14T01:00:00.000Z", "succeeded", 1, 0, 0, 0, 0),
	}
	rows := Commands(records)
	var names []string
	for _, r := range rows[1 : len(rows)-1] {
		names = append(names, r[0])
	}
	want := []string{"analyze", "build", "import_compass", "plan", "plan repair"}
	if !reflect.DeepEqual(names, want) {
		t.Errorf("command order = %v, want %v", names, want)
	}
	total := rows[len(rows)-1]
	if total[0] != "TOTAL" || total[1] != "7" || total[2] != "1" || total[3] != "7" {
		t.Errorf("TOTAL = %v, want calls 7, failed 1, input 7", total)
	}
}

func TestCommandsEmptyMatchWritesZeroTotal(t *testing.T) {
	rows := Commands(nil)
	want := []string{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"}
	if !reflect.DeepEqual(rows[0], want) {
		t.Errorf("header = %v, want %v", rows[0], want)
	}
	if !reflect.DeepEqual(rows[1], []string{"TOTAL", "0", "0", "0", "0", "0", "0.0"}) {
		t.Errorf("TOTAL = %v, want zero row", rows[1])
	}
}

func TestTargetsOrdersByteWiseAndSumsCostAndMinutes(t *testing.T) {
	records := []logrecord.Record{
		rec("1", "Zeta", "b", "2026-10-14T01:00:00.000Z", "succeeded", 1, 2, 3, 1000, 0.25),
		rec("2", "Alpha", "b", "2026-10-14T01:00:00.000Z", "succeeded", 1, 0, 0, 0, 0.5),
		rec("3", "Alpha", "b", "2026-10-14T01:00:00.000Z", "failed", 0, 0, 0, 0, 0),
	}
	rows := Targets(records)
	if rows[1][0] != "Alpha" || rows[2][0] != "Zeta" {
		t.Errorf("targets order = %q, %q; want Alpha, Zeta", rows[1][0], rows[2][0])
	}
	if rows[1][1] != "2" || rows[1][2] != "1" || rows[1][6] != "0.5000" {
		t.Errorf("Alpha row = %v", rows[1])
	}
	total := rows[len(rows)-1]
	if total[0] != "TOTAL" || total[1] != "3" || total[2] != "1" || total[6] != "0.7500" || total[7] != "0.0" {
		t.Errorf("TOTAL = %v", total)
	}
}

func TestTargetsEmptyMatchWritesZeroTotal(t *testing.T) {
	rows := Targets(nil)
	if !reflect.DeepEqual(rows[1], []string{"TOTAL", "0", "0", "0", "0", "0", "0.0000", "0.0"}) {
		t.Errorf("TOTAL = %v", rows[1])
	}
}

func TestDailyOrdersByDateAndTokensIncludeCached(t *testing.T) {
	records := []logrecord.Record{
		rec("1", "T", "b", "2026-10-15T09:00:00.000Z", "succeeded", 10, 20, 30, 1000, 0),
		rec("2", "T", "b", "2026-10-14T09:00:00.000Z", "succeeded", 10, 20, 30, 1000, 0),
		rec("3", "T", "b", "2026-10-14T10:00:00.000Z", "failed", 10, 20, 30, 1000, 0),
	}
	rows := Daily(records)
	if rows[1][0] != "2026-10-14" || rows[2][0] != "2026-10-15" {
		t.Errorf("date order = %q, %q", rows[1][0], rows[2][0])
	}
	// Two 1000 ms records are 2000 ms, which rounds to 0.0; all three are 3000 ms, which rounds to 0.1.
	if !reflect.DeepEqual(rows[1], []string{"2026-10-14", "2", "1", "120", "0.0"}) {
		t.Errorf("first group = %v", rows[1])
	}
	if !reflect.DeepEqual(rows[3], []string{"TOTAL", "3", "1", "180", "0.1"}) {
		t.Errorf("TOTAL = %v", rows[3])
	}
}

func TestDailyEmptyMatchWritesZeroTotal(t *testing.T) {
	rows := Daily(nil)
	if !reflect.DeepEqual(rows[1], []string{"TOTAL", "0", "0", "0", "0.0"}) {
		t.Errorf("TOTAL = %v", rows[1])
	}
}

func TestModelsGroupByProviderAndModelWithEmptyTotalModel(t *testing.T) {
	records := []logrecord.Record{
		{ExecutionID: "1", Provider: "codex", Model: "gpt-x", Status: "succeeded", CostUSD: 0},
		{ExecutionID: "2", Provider: "claude", Model: "haiku", Status: "succeeded", CostUSD: 0.1},
		{ExecutionID: "3", Provider: "claude", Model: "haiku", Status: "failed", CostUSD: 0.2},
		{ExecutionID: "4", Provider: "claude", Model: "alpha", Status: "succeeded", CostUSD: 0},
	}
	rows := Models(records)
	var pairs [][2]string
	for _, r := range rows[1 : len(rows)-1] {
		pairs = append(pairs, [2]string{r[0], r[1]})
	}
	want := [][2]string{{"claude", "alpha"}, {"claude", "haiku"}, {"codex", "gpt-x"}}
	if !reflect.DeepEqual(pairs, want) {
		t.Errorf("pairs = %v, want %v", pairs, want)
	}
	total := rows[len(rows)-1]
	if !reflect.DeepEqual(total, []string{"TOTAL", "", "4", "1", "0", "0", "0", "0.3000"}) {
		t.Errorf("TOTAL = %v", total)
	}
}

func TestModelsEmptyMatchWritesZeroTotal(t *testing.T) {
	rows := Models(nil)
	if !reflect.DeepEqual(rows[1], []string{"TOTAL", "", "0", "0", "0", "0", "0", "0.0000"}) {
		t.Errorf("TOTAL = %v", rows[1])
	}
}

func TestFailuresListsOnlyFailedOrderedWithoutTotal(t *testing.T) {
	rc1, rc0 := int64(1), int64(0)
	records := []logrecord.Record{
		{ExecutionID: "b", StartedAt: "2026-10-15T06:50:39.621Z", Status: "failed", Target: "C", Command: "diagnose", Model: "sonnet", ReturnCode: nil, TimedOut: true},
		{ExecutionID: "a", StartedAt: "2026-10-14T02:00:00.000Z", Status: "failed", Target: "T", Command: "diagnose", Model: "gemini", ReturnCode: &rc1},
		{ExecutionID: "c", StartedAt: "2026-10-16T00:00:00.000Z", Status: "succeeded", Target: "T", Command: "build", Model: "h", ReturnCode: &rc0},
	}
	rows := Failures(records)
	if len(rows) != 3 {
		t.Fatalf("rows = %d, want header plus two failures", len(rows))
	}
	if !reflect.DeepEqual(rows[1], []string{"2026-10-14T02:00:00.000Z", "T", "diagnose", "gemini", "1", "false"}) {
		t.Errorf("row 1 = %v", rows[1])
	}
	if !reflect.DeepEqual(rows[2], []string{"2026-10-15T06:50:39.621Z", "C", "diagnose", "sonnet", "", "true"}) {
		t.Errorf("row 2 = %v; null returncode must be empty", rows[2])
	}
}

func TestFailuresEmptyMatchWritesHeaderOnly(t *testing.T) {
	rows := Failures(nil)
	if len(rows) != 1 || !reflect.DeepEqual(rows[0], failuresHeader) {
		t.Errorf("rows = %v, want header only", rows)
	}
}
