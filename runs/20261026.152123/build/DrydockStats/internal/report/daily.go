package report

import (
	"io"
	"sort"

	"ddkstat/internal/format"
	"ddkstat/internal/logread"
)

// DailyTotals is the totals for one start date.
type DailyTotals struct {
	Date string
	Totals
}

// AggregateDaily groups records by start date, ordered ascending in byte order. It also returns the
// totals over every record, which become the TOTAL row.
func AggregateDaily(records []logread.Record) ([]DailyTotals, Totals) {
	var all Totals
	var rows []DailyTotals
	index := make(map[string]int)
	for _, rec := range records {
		all.Add(rec)
		i, ok := index[rec.Date]
		if !ok {
			i = len(rows)
			index[rec.Date] = i
			rows = append(rows, DailyTotals{Date: rec.Date})
		}
		rows[i].Add(rec)
	}
	sort.Slice(rows, func(a, b int) bool { return rows[a].Date < rows[b].Date })
	return rows, all
}

// WriteDaily writes the daily report: a header, one row per start date and a TOTAL row.
func WriteDaily(w io.Writer, records []logread.Record, _ int) error {
	rows, all := AggregateDaily(records)
	table := [][]string{{"date", "calls", "failed", "tokens", "minutes"}}
	for _, r := range rows {
		table = append(table, dailyCells(r.Date, r.Totals))
	}
	return writeTable(w, append(table, dailyCells("TOTAL", all)))
}

// dailyCells renders one daily row with the given label in the date column.
func dailyCells(label string, t Totals) []string {
	return []string{label, itoa(t.Calls), itoa(t.Failed), itoa(t.Tokens()), format.Minutes(t.ElapsedMS)}
}
