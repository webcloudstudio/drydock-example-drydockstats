package report

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"testing"

	"ddkstat/internal/cli"
	"ddkstat/internal/logread"
)

func summaryFixture() []logread.Record {
	return []logread.Record{
		rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		asFailed(rec("b", "2026-10-16T01:00:00.000Z", "plan", "Toml")),
	}
}

func TestSummaryAggregate(t *testing.T) {
	got := SummaryAggregate(summaryFixture(), 3)
	want := Summary{
		Totals: Totals{
			Calls:        2,
			Failed:       1,
			InputTokens:  2,
			CachedTokens: 4,
			OutputTokens: 6,
			CostUSD:      1.0,
			ElapsedMS:    120000,
		},
		Skipped: 3,
		First:   "2026-10-14",
		Last:    "2026-10-16",
	}
	if got != want {
		t.Errorf("SummaryAggregate = %+v, want %+v", got, want)
	}
}

func TestSummaryEmptyFirstLast(t *testing.T) {
	got := SummaryAggregate(nil, 0)
	if got.First != "" || got.Last != "" {
		t.Errorf("First, Last = %q, %q; want empty", got.First, got.Last)
	}
	if got.Calls != 0 || got.Failed != 0 || got.Tokens() != 0 || got.CostUSD != 0 || got.ElapsedMS != 0 {
		t.Errorf("counters = %+v, want zero", got)
	}
}

func TestSummarySkippedIgnoresFilters(t *testing.T) {
	log := writeLog(t,
		`{"execution_id":"a","started_at":"2026-10-14T01:00:00.000Z","status":"succeeded","job":{"command_name":"build","target":"Toml"}}`,
		"{not json",
		"{also bad",
		"{bad three",
	)
	var stdout, stderr bytes.Buffer
	if code := cli.Run([]string{"summary", "--target", "NoSuchTarget", log}, &stdout, &stderr); code != 0 {
		t.Fatalf("Run exit = %d, want 0", code)
	}
	rows, err := csv.NewReader(&stdout).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v", err)
	}
	if got := rows[1][:2]; !reflect.DeepEqual(got, []string{"0", "3"}) {
		t.Errorf("records, skipped = %v, want [0 3]", got)
	}
}

func TestSummaryWrite(t *testing.T) {
	rows := readCSV(t, WriteSummary, summaryFixture(), 3)
	wantRows(t, rows, [][]string{
		{"records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{"2", "3", "1", "1", "2026-10-14", "2026-10-16", "2", "4", "6", "1.0000", "2.0"},
	})
}
