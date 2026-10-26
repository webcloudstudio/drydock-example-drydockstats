package report

import (
	"io"
	"sort"

	"ddkstat/internal/format"
	"ddkstat/internal/logread"
)

// ModelTotals is the totals for one provider and executed-model pair.
type ModelTotals struct {
	Provider string
	Model    string
	Totals
}

// modelKey identifies one provider and model pair.
type modelKey struct {
	provider string
	model    string
}

// AggregateModels groups records by provider and Record.Model, ordered by provider then model
// ascending in byte order. It also returns the totals over every record, which become the TOTAL row.
func AggregateModels(records []logread.Record) ([]ModelTotals, Totals) {
	var all Totals
	var rows []ModelTotals
	index := make(map[modelKey]int)
	for _, rec := range records {
		all.Add(rec)
		key := modelKey{provider: rec.Provider, model: rec.Model}
		i, ok := index[key]
		if !ok {
			i = len(rows)
			index[key] = i
			rows = append(rows, ModelTotals{Provider: rec.Provider, Model: rec.Model})
		}
		rows[i].Add(rec)
	}
	sort.Slice(rows, func(a, b int) bool {
		if rows[a].Provider != rows[b].Provider {
			return rows[a].Provider < rows[b].Provider
		}
		return rows[a].Model < rows[b].Model
	})
	return rows, all
}

// WriteModels writes the models report: a header, one row per provider and model, and a TOTAL row
// whose model column is empty.
func WriteModels(w io.Writer, records []logread.Record, _ int) error {
	rows, all := AggregateModels(records)
	table := [][]string{{"provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"}}
	for _, r := range rows {
		table = append(table, modelCells(r.Provider, r.Model, r.Totals))
	}
	return writeTable(w, append(table, modelCells("TOTAL", "", all)))
}

// modelCells renders one models row with the given provider and model columns.
func modelCells(provider, model string, t Totals) []string {
	return []string{
		provider, model, itoa(t.Calls), itoa(t.Failed), itoa(t.InputTokens), itoa(t.CachedTokens),
		itoa(t.OutputTokens), format.Cost(t.CostUSD),
	}
}
