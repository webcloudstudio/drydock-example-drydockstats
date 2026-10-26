// Package report builds the six ddkstat CSV reports from records already narrowed by the filters.
//
// Each report has an aggregate or selection function, which is pure, and a writer, which formats
// its result as CSV with \n line endings. Text values are written unquoted and unescaped.
//
// The writers are registered with cli.Register from init, so cli.Run reaches them without importing
// this package. cmd/ddkstat imports this package for that side effect.
package report

import (
	"io"
	"strconv"
	"strings"

	"ddkstat/internal/cli"
	"ddkstat/internal/logread"
)

func init() { Register() }

// Register makes the writer of every report available to cli.Run. Registering again replaces the
// writers with the same ones, so it is safe to call more than once.
func Register() {
	cli.Register("summary", WriteSummary)
	cli.Register("commands", WriteCommands)
	cli.Register("targets", WriteTargets)
	cli.Register("models", WriteModels)
	cli.Register("daily", WriteDaily)
	cli.Register("failures", WriteFailures)
}

// Totals is the sum over a group of records, or over every filtered record for a TOTAL row.
type Totals struct {
	Calls        int64
	Failed       int64
	InputTokens  int64
	CachedTokens int64
	OutputTokens int64
	CostUSD      float64
	ElapsedMS    int64
}

// Add folds one record into the totals.
func (t *Totals) Add(rec logread.Record) {
	t.Calls++
	if rec.Failed {
		t.Failed++
	}
	t.InputTokens += rec.InputTokens
	t.CachedTokens += rec.CachedTokens
	t.OutputTokens += rec.OutputTokens
	t.CostUSD += rec.CostUSD
	t.ElapsedMS += rec.ElapsedMS
}

// Tokens is input plus cached plus output tokens.
func (t Totals) Tokens() int64 {
	return t.InputTokens + t.CachedTokens + t.OutputTokens
}

// writeTable writes each row as one comma-separated line ended by \n.
func writeTable(w io.Writer, rows [][]string) error {
	for _, row := range rows {
		if _, err := io.WriteString(w, strings.Join(row, ",")+"\n"); err != nil {
			return err
		}
	}
	return nil
}

// itoa renders an integer count or token total in decimal.
func itoa(n int64) string {
	return strconv.FormatInt(n, 10)
}
