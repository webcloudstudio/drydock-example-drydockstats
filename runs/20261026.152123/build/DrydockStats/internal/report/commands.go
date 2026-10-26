package report

import (
	"io"
	"sort"

	"ddkstat/internal/format"
	"ddkstat/internal/logread"
)

// CommandTotals is the totals for one Drydock command.
type CommandTotals struct {
	Command string
	Totals
}

// AggregateCommands groups records by command, ordered by calls descending and then command
// ascending in byte order. It also returns the totals over every record, which become the TOTAL row.
func AggregateCommands(records []logread.Record) ([]CommandTotals, Totals) {
	var all Totals
	var rows []CommandTotals
	index := make(map[string]int)
	for _, rec := range records {
		all.Add(rec)
		i, ok := index[rec.Command]
		if !ok {
			i = len(rows)
			index[rec.Command] = i
			rows = append(rows, CommandTotals{Command: rec.Command})
		}
		rows[i].Add(rec)
	}
	sort.Slice(rows, func(a, b int) bool {
		if rows[a].Calls != rows[b].Calls {
			return rows[a].Calls > rows[b].Calls
		}
		return rows[a].Command < rows[b].Command
	})
	return rows, all
}

// WriteCommands writes the commands report: a header, one row per command and a TOTAL row.
func WriteCommands(w io.Writer, records []logread.Record, _ int) error {
	rows, all := AggregateCommands(records)
	table := [][]string{{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"}}
	for _, r := range rows {
		table = append(table, commandCells(r.Command, r.Totals))
	}
	return writeTable(w, append(table, commandCells("TOTAL", all)))
}

// commandCells renders one commands row with the given label in the command column.
func commandCells(label string, t Totals) []string {
	return []string{
		label, itoa(t.Calls), itoa(t.Failed), itoa(t.InputTokens), itoa(t.CachedTokens), itoa(t.OutputTokens),
		format.Minutes(t.ElapsedMS),
	}
}
