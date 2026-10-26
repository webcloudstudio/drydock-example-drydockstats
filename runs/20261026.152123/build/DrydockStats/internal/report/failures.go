package report

import (
	"io"
	"sort"
	"strconv"

	"ddkstat/internal/logread"
)

// SelectFailures returns the failed records ordered by started_at and then execution_id ascending in
// byte order. The input slice is not modified.
func SelectFailures(records []logread.Record) []logread.Record {
	var out []logread.Record
	for _, rec := range records {
		if rec.Failed {
			out = append(out, rec)
		}
	}
	sort.Slice(out, func(a, b int) bool {
		if out[a].StartedAt != out[b].StartedAt {
			return out[a].StartedAt < out[b].StartedAt
		}
		return out[a].ExecutionID < out[b].ExecutionID
	})
	return out
}

// WriteFailures writes the failures report: a header and one row per failed record. There is no
// TOTAL row, and an empty selection writes the header alone.
func WriteFailures(w io.Writer, records []logread.Record, _ int) error {
	table := [][]string{{"started_at", "target", "command", "model", "returncode", "timed_out"}}
	for _, rec := range SelectFailures(records) {
		returnCode := ""
		if rec.ReturnCode != nil {
			returnCode = strconv.Itoa(*rec.ReturnCode)
		}
		table = append(table, []string{
			rec.StartedAt, rec.Target, rec.Command, rec.Model, returnCode, strconv.FormatBool(rec.TimedOut),
		})
	}
	return writeTable(w, table)
}
