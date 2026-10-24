package format

import (
	"bytes"
	"strings"
	"testing"
)

func TestCostHasFourDecimalPlaces(t *testing.T) {
	tests := []struct {
		usd  float64
		want string
	}{
		{0, "0.0000"},
		{0.5, "0.5000"},
		{5.5245, "5.5245"},
		{1.23456789, "1.2346"},
		{12, "12.0000"},
	}
	for _, tt := range tests {
		if got := Cost(tt.usd); got != tt.want {
			t.Errorf("Cost(%v) = %q, want %q", tt.usd, got, tt.want)
		}
	}
}

func TestMinutesRoundsToTenths(t *testing.T) {
	tests := []struct {
		ms   int64
		want string
	}{
		{0, "0.0"},
		{95000, "1.6"},
		{2999, "0.0"},
		{3000, "0.1"},
		{60000, "1.0"},
		{5991000, "99.9"},
		{5999000, "100.0"},
	}
	for _, tt := range tests {
		if got := Minutes(tt.ms); got != tt.want {
			t.Errorf("Minutes(%d) = %q, want %q", tt.ms, got, tt.want)
		}
	}
}

func TestWriteCSVUsesLFOnly(t *testing.T) {
	rows := [][]string{{"a", "b"}, {"1", ""}, {"TOTAL", "2"}}
	var buf bytes.Buffer
	if err := WriteCSV(&buf, rows); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(buf.String(), "\r") {
		t.Errorf("output contains CR: %q", buf.String())
	}
	if got, want := buf.String(), "a,b\n1,\nTOTAL,2\n"; got != want {
		t.Errorf("WriteCSV = %q, want %q", got, want)
	}
}

func TestWriteCSVEmptyRowsWritesNothing(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteCSV(&buf, nil); err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 0 {
		t.Errorf("wrote %q for no rows", buf.String())
	}
}
