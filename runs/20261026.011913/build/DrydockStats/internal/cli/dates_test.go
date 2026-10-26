package cli

import "testing"

func TestParseDateAcceptsRealDates(t *testing.T) {
	for _, value := range []string{"2026-10-14", "2026-01-01", "2026-12-31", "2024-02-29", "2000-02-29"} {
		t.Run(value, func(t *testing.T) {
			got, err := validateDate(value)
			if err != nil {
				t.Fatalf("validateDate(%q) error: %v", value, err)
			}
			if got != value {
				t.Errorf("validateDate(%q) = %q, want the value unchanged", value, got)
			}
		})
	}
}

func TestParseDateRejectsMalformed(t *testing.T) {
	for _, value := range []string{
		"2026-1-5",
		"2026-10-5",
		"20261014",
		"2026/10/14",
		"2026-10-14 ",
		" 2026-10-14",
		"2026-10-14T00",
		"2026-10-144",
		"",
		"yesterday",
	} {
		t.Run(value, func(t *testing.T) {
			if _, err := validateDate(value); err == nil {
				t.Errorf("validateDate(%q) returned no error", value)
			}
		})
	}
}

func TestParseDateRejectsImpossibleDates(t *testing.T) {
	for _, value := range []string{
		"2026-13-01",
		"2026-00-10",
		"2026-10-00",
		"2026-02-30",
		"2025-02-29",
		"2026-04-31",
	} {
		t.Run(value, func(t *testing.T) {
			if _, err := validateDate(value); err == nil {
				t.Errorf("validateDate(%q) returned no error", value)
			}
		})
	}
}

func TestParseCarriesValidDates(t *testing.T) {
	got, err := Parse([]string{"daily", "--since", "2026-10-14", "--until", "2026-10-16", "llm.jsonl"})
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if got.Since != "2026-10-14" || got.Until != "2026-10-16" {
		t.Errorf("Since=%q Until=%q, want 2026-10-14 and 2026-10-16", got.Since, got.Until)
	}
}

func TestParseRejectsInvalidDateOptions(t *testing.T) {
	for _, args := range [][]string{
		{"summary", "--since", "2026-13-01", "llm.jsonl"},
		{"summary", "--until", "2026-02-30", "llm.jsonl"},
		{"summary", "--since", "", "llm.jsonl"},
		{"summary", "--until", "2026-1-5", "llm.jsonl"},
	} {
		if _, err := Parse(args); err == nil {
			t.Errorf("Parse(%q) returned no error", args)
		}
	}
}
