// Package format renders report values and writes report rows as CSV.
package format

import (
	"fmt"
	"io"
	"strings"
)

// Cost renders a summed cost_usd with exactly four decimal places.
func Cost(usd float64) string {
	return fmt.Sprintf("%.4f", usd)
}

// Minutes renders elapsed milliseconds as minutes to one decimal place, using integer rounding to
// the nearest tenth of a minute.
func Minutes(ms int64) string {
	t := (ms + 3000) / 6000
	return fmt.Sprintf("%d.%d", t/10, t%10)
}

// WriteCSV writes each row as comma-separated fields terminated by LF. rows[0] is the header.
func WriteCSV(w io.Writer, rows [][]string) error {
	for _, row := range rows {
		if _, err := io.WriteString(w, strings.Join(row, ",")+"\n"); err != nil {
			return err
		}
	}
	return nil
}
