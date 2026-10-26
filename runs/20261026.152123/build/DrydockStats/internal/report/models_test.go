package report

import (
	"reflect"
	"testing"

	"ddkstat/internal/logread"
)

// modelFixture returns one claude Haiku success and one gemini Sonnet failure.
func modelFixture() []logread.Record {
	haiku := rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml")
	sonnet := asFailed(rec("b", "2026-10-14T02:00:00.000Z", "plan", "Toml"))
	sonnet.Provider = "gemini"
	sonnet.Model = "claude-sonnet-5-5"
	return []logread.Record{haiku, sonnet}
}

func TestModelsGroupingUsesExecutedModel(t *testing.T) {
	executed := rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml")
	executed.RequestedModel = "sonnet"
	executed.Model = "claude-sonnet-5-5"
	sameExecuted := rec("b", "2026-10-14T02:00:00.000Z", "build", "Toml")
	sameExecuted.RequestedModel = "claude-sonnet-5-5"
	sameExecuted.Model = "claude-sonnet-5-5"
	fallback := rec("c", "2026-10-14T03:00:00.000Z", "build", "Toml")
	fallback.RequestedModel = "sonnet"
	fallback.Model = "sonnet"

	rows, _ := AggregateModels([]logread.Record{executed, sameExecuted, fallback})
	type pair struct {
		model string
		calls int64
	}
	var got []pair
	for _, r := range rows {
		got = append(got, pair{r.Model, r.Calls})
	}
	want := []pair{{"claude-sonnet-5-5", 2}, {"sonnet", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("models = %v, want %v", got, want)
	}
}

func TestModelsOrdering(t *testing.T) {
	records := []logread.Record{
		rec("a", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		rec("b", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		rec("c", "2026-10-14T01:00:00.000Z", "build", "Toml"),
		rec("d", "2026-10-14T01:00:00.000Z", "build", "Toml"),
	}
	records[0].Provider, records[0].Model = "gemini", "a"
	records[1].Provider, records[1].Model = "claude", "b"
	records[2].Provider, records[2].Model = "claude", "a"
	records[3].Provider, records[3].Model = "claude", "B"
	rows, _ := AggregateModels(records)
	type pair struct{ provider, model string }
	var got []pair
	for _, r := range rows {
		got = append(got, pair{r.Provider, r.Model})
	}
	want := []pair{{"claude", "B"}, {"claude", "a"}, {"claude", "b"}, {"gemini", "a"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("models = %v, want %v", got, want)
	}
}

func TestModelsTotalEmptyModelColumn(t *testing.T) {
	rows := readCSV(t, WriteModels, modelFixture(), 0)
	total := rows[len(rows)-1]
	wantTotal := []string{"TOTAL", "", "2", "1", "2", "4", "6", "1.0000"}
	if !reflect.DeepEqual(total, wantTotal) {
		t.Errorf("TOTAL row = %v, want %v", total, wantTotal)
	}
}

func TestModelsEmpty(t *testing.T) {
	rows := readCSV(t, WriteModels, nil, 0)
	wantRows(t, rows, [][]string{
		{"provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"},
		{"TOTAL", "", "0", "0", "0", "0", "0", "0.0000"},
	})
}

func TestModelsWrite(t *testing.T) {
	rows := readCSV(t, WriteModels, modelFixture(), 0)
	wantRows(t, rows, [][]string{
		{"provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"},
		{"claude", "claude-haiku-5-5", "1", "0", "1", "2", "3", "0.5000"},
		{"gemini", "claude-sonnet-5-5", "1", "1", "1", "2", "3", "0.5000"},
		{"TOTAL", "", "2", "1", "2", "4", "6", "1.0000"},
	})
}
