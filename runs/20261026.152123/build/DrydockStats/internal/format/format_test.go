package format_test

import (
	"fmt"
	"strings"
	"testing"

	"ddkstat/internal/format"
)

// expectMinutes derives the tenths-of-a-minute text from the rounding formula itself, so the
// expectation is computed rather than typed by hand.
func expectMinutes(ms int64) string {
	t := (ms + 3000) / 6000
	return fmt.Sprintf("%d.%d", t/10, t%10)
}

func TestMinutesRounding(t *testing.T) {
	// Boundaries where the tenth changes: the half-tenth point is 3000 ms, then every 6000 ms.
	cases := []int64{0, 2999, 3000, 5999, 6000, 8999, 9000, 95000, 594000, 600000, 3600000}
	for _, ms := range cases {
		if got, want := format.Minutes(ms), expectMinutes(ms); got != want {
			t.Errorf("Minutes(%d) = %q, want %q", ms, got, want)
		}
	}
	// Spot values from the specification, written out so a formula error cannot pass silently.
	if got := format.Minutes(95000); got != "1.6" {
		t.Errorf("Minutes(95000) = %q, want %q", got, "1.6")
	}
	if got := format.Minutes(2999); got != "0.0" {
		t.Errorf("Minutes(2999) = %q, want %q", got, "0.0")
	}
	if got := format.Minutes(3000); got != "0.1" {
		t.Errorf("Minutes(3000) = %q, want %q", got, "0.1")
	}
}

func TestMinutesZero(t *testing.T) {
	if got := format.Minutes(0); got != "0.0" {
		t.Errorf("Minutes(0) = %q, want %q", got, "0.0")
	}
}

func TestCostFourPlaces(t *testing.T) {
	cases := []float64{0, 1, 1.5, 0.123456, 12.3, 1234.56789}
	for _, usd := range cases {
		got := format.Cost(usd)
		want := fmt.Sprintf("%.4f", usd)
		if got != want {
			t.Errorf("Cost(%v) = %q, want %q", usd, got, want)
		}
		_, frac, found := strings.Cut(got, ".")
		if !found || len(frac) != 4 {
			t.Errorf("Cost(%v) = %q, want exactly four decimal places", usd, got)
		}
	}
}

func TestCostZero(t *testing.T) {
	if got := format.Cost(0); got != "0.0000" {
		t.Errorf("Cost(0) = %q, want %q", got, "0.0000")
	}
}
