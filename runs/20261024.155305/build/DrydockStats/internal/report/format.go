// Package report filters and aggregates log records into CSV report rows.
package report

import (
	"encoding/csv"
	"fmt"
	"io"
)

// FormatMinutes writes milliseconds as tenths of a minute, rounded half up.
func FormatMinutes(ms int64) string {
	t := (ms + 3000) / 6000
	return fmt.Sprintf("%d.%d", t/10, t%10)
}

// FormatCost writes US dollars with four decimal places.
func FormatCost(usd float64) string {
	return fmt.Sprintf("%.4f", usd)
}

// WriteCSV writes every row as CSV with \n line endings.
func WriteCSV(w io.Writer, rows [][]string) error {
	return csv.NewWriter(w).WriteAll(rows)
}
