package logread

import "testing"

func TestDuplicateFirstOccurrenceKept(t *testing.T) {
	input := lineOf(t, baseRecord("a")) + "\n" + lineOf(t, baseRecord("a")) + "\n"
	res, _ := readText(t, input)
	if len(res.Records) != 1 {
		t.Fatalf("kept %d records, want 1", len(res.Records))
	}
	if res.Skipped != 1 {
		t.Errorf("Skipped = %d, want 1", res.Skipped)
	}
}

func TestDuplicateKeepsFirstEvenWhenFieldsDiffer(t *testing.T) {
	first := baseRecord("a")
	first["job"].(map[string]any)["target"] = "Toml"
	first["result"].(map[string]any)["stats"].(map[string]any)["input_tokens"] = 1
	second := baseRecord("a")
	second["job"].(map[string]any)["target"] = "CommonMark"
	second["result"].(map[string]any)["stats"].(map[string]any)["input_tokens"] = 999

	input := lineOf(t, first) + "\n" + lineOf(t, second) + "\n"
	res, _ := readText(t, input)
	if len(res.Records) != 1 {
		t.Fatalf("kept %d records, want 1", len(res.Records))
	}
	if got := res.Records[0]; got.Target != "Toml" || got.InputTokens != 1 {
		t.Errorf("kept target=%q input=%d, want Toml and 1", got.Target, got.InputTokens)
	}
}

func TestSkippedLineDoesNotRegisterID(t *testing.T) {
	malformed := baseRecord("x")
	delete(malformed, "status")
	input := lineOf(t, malformed) + "\n" + lineOf(t, baseRecord("x")) + "\n"
	res, _ := readText(t, input)
	if got := ids(res); !equalStrings(got, []string{"x"}) {
		t.Fatalf("kept ids = %v, want [x]", got)
	}
	if res.Skipped != 1 {
		t.Errorf("Skipped = %d, want 1", res.Skipped)
	}
}

func TestDuplicateWarnsWithLineNumber(t *testing.T) {
	input := lineOf(t, baseRecord("a")) + "\n" +
		lineOf(t, baseRecord("a")) + "\n" +
		lineOf(t, baseRecord("b")) + "\n" +
		lineOf(t, baseRecord("a")) + "\n"
	_, warn := readText(t, input)
	if got := warnedLines(t, warn); !equalInts(got, []int{2, 4}) {
		t.Errorf("warned lines = %v, want [2 4]", got)
	}
}

func TestDuplicateCountedWithMalformed(t *testing.T) {
	input := lineOf(t, baseRecord("a")) + "\n" +
		"{bad\n" +
		lineOf(t, baseRecord("a")) + "\n" +
		lineOf(t, baseRecord("b")) + "\n"
	res, _ := readText(t, input)
	if got := ids(res); !equalStrings(got, []string{"a", "b"}) {
		t.Errorf("kept ids = %v, want [a b]", got)
	}
	if res.Skipped != 2 {
		t.Errorf("Skipped = %d, want 2 (one malformed, one duplicate)", res.Skipped)
	}
}
