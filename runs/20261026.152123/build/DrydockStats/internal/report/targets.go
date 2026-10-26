package report

import (
	"io"
	"sort"

	"ddkstat/internal/format"
	"ddkstat/internal/logread"
)

// TargetTotals is the totals for one Target.
type TargetTotals struct {
	Target string
	Totals
}

// AggregateTargets groups records by Target, ordered by Target ascending in byte order. It also
// returns the totals over every record, which become the TOTAL row.
func AggregateTargets(records []logread.Record) ([]TargetTotals, Totals) {
	var all Totals
	var rows []TargetTotals
	index := make(map[string]int)
	for _, rec := range records {
		all.Add(rec)
		i, ok := index[rec.Target]
		if !ok {
			i = len(rows)
			index[rec.Target] = i
			rows = append(rows, TargetTotals{Target: rec.Target})
		}
		rows[i].Add(rec)
	}
	sort.Slice(rows, func(a, b int) bool { return rows[a].Target < rows[b].Target })
	return rows, all
}

// WriteTargets writes the targets report: a header, one row per Target and a TOTAL row.
func WriteTargets(w io.Writer, records []logread.Record, _ int) error {
	rows, all := AggregateTargets(records)
	table := [][]string{{"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"}}
	for _, r := range rows {
		table = append(table, targetCells(r.Target, r.Totals))
	}
	return writeTable(w, append(table, targetCells("TOTAL", all)))
}

// targetCells renders one targets row with the given label in the target column.
func targetCells(label string, t Totals) []string {
	return []string{
		label, itoa(t.Calls), itoa(t.Failed), itoa(t.InputTokens), itoa(t.CachedTokens), itoa(t.OutputTokens),
		format.Cost(t.CostUSD), format.Minutes(t.ElapsedMS),
	}
}
