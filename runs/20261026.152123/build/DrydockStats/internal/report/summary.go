package report

import (
	"io"

	"ddkstat/internal/format"
	"ddkstat/internal/logread"
)

// Summary is the one-row summary: totals over the filtered records and the whole-file skipped count.
type Summary struct {
	Totals
	Skipped int64
	First   string
	Last    string
}

// SummaryAggregate totals records into a Summary. skipped is recorded as given, so the caller passes
// the whole-file count. First and Last are the earliest and latest start dates, empty when records
// is empty.
func SummaryAggregate(records []logread.Record, skipped int) Summary {
	s := Summary{Skipped: int64(skipped)}
	for _, rec := range records {
		s.Add(rec)
		if s.Calls == 1 || rec.Date < s.First {
			s.First = rec.Date
		}
		if rec.Date > s.Last {
			s.Last = rec.Date
		}
	}
	return s
}

// WriteSummary writes the summary report: a header and one data row.
func WriteSummary(w io.Writer, records []logread.Record, skipped int) error {
	s := SummaryAggregate(records, skipped)
	return writeTable(w, [][]string{
		{"records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{
			itoa(s.Calls), itoa(s.Skipped), itoa(s.Calls - s.Failed), itoa(s.Failed), s.First, s.Last,
			itoa(s.InputTokens), itoa(s.CachedTokens), itoa(s.OutputTokens),
			format.Cost(s.CostUSD), format.Minutes(s.ElapsedMS),
		},
	})
}
