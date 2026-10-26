// Package format holds the shared text formatters for minutes and cost columns.
//
// Every report that prints a minutes or cost column uses these functions, so the rounding rule
// lives in exactly one place. Both functions are pure: they take no filters, read no input and
// keep no state.
package format

import "fmt"

// Minutes renders a duration in milliseconds as tenths of a minute with one decimal digit. The
// duration is rounded to the nearest tenth using t = (ms + 3000) / 6000 in integer division, so
// 95000 ms gives "1.6" and 0 ms gives "0.0".
func Minutes(ms int64) string {
	t := (ms + 3000) / 6000
	return fmt.Sprintf("%d.%d", t/10, t%10)
}

// Cost renders a US-dollar amount with exactly four decimal places, so 0 gives "0.0000".
func Cost(usd float64) string {
	return fmt.Sprintf("%.4f", usd)
}
