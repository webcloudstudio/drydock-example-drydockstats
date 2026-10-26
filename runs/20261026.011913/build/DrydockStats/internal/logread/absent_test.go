package logread

import "testing"

// stats returns the result.stats object of a base record.
func stats(rec map[string]any) map[string]any {
	return rec["result"].(map[string]any)["stats"].(map[string]any)
}

func TestAbsentStatsIsStillRecord(t *testing.T) {
	rec := baseRecord("a")
	delete(rec["result"].(map[string]any), "stats")
	got := readOne(t, rec)
	if got.InputTokens != 0 || got.CachedTokens != 0 || got.OutputTokens != 0 || got.CostUSD != 0 || got.ElapsedMS != 0 {
		t.Errorf("numeric fields = %+v, want all zero", got)
	}
	if got.ExecutionID != "a" {
		t.Errorf("ExecutionID = %q, want a", got.ExecutionID)
	}
}

func TestNullCostCountsZero(t *testing.T) {
	rec := baseRecord("a")
	stats(rec)["cost_usd"] = nil
	got := readOne(t, rec)
	if got.CostUSD != 0 {
		t.Errorf("CostUSD = %v, want 0", got.CostUSD)
	}
	if got.InputTokens != 843 || got.OutputTokens != 7751 || got.ElapsedMS != 54872 {
		t.Errorf("other fields = in %d out %d ms %d, want 843, 7751, 54872", got.InputTokens, got.OutputTokens, got.ElapsedMS)
	}
}

func TestAbsentNumbersCountZero(t *testing.T) {
	for _, key := range []string{"input_tokens", "cached_input_tokens", "output_tokens", "cost_usd", "elapsed_ms"} {
		t.Run(key, func(t *testing.T) {
			rec := baseRecord("a")
			delete(stats(rec), key)
			got := readOne(t, rec)
			switch key {
			case "input_tokens":
				if got.InputTokens != 0 {
					t.Errorf("InputTokens = %d, want 0", got.InputTokens)
				}
			case "cached_input_tokens":
				if got.CachedTokens != 0 {
					t.Errorf("CachedTokens = %d, want 0", got.CachedTokens)
				}
			case "output_tokens":
				if got.OutputTokens != 0 {
					t.Errorf("OutputTokens = %d, want 0", got.OutputTokens)
				}
			case "cost_usd":
				if got.CostUSD != 0 {
					t.Errorf("CostUSD = %v, want 0", got.CostUSD)
				}
			case "elapsed_ms":
				if got.ElapsedMS != 0 {
					t.Errorf("ElapsedMS = %d, want 0", got.ElapsedMS)
				}
			}
		})
	}
}

func TestWrongTypedNumbersCountZero(t *testing.T) {
	wrong := map[string]any{
		"string":  "843",
		"boolean": true,
		"array":   []any{1},
		"object":  map[string]any{"n": 1},
	}
	for name, value := range wrong {
		t.Run(name, func(t *testing.T) {
			rec := baseRecord("a")
			stats(rec)["input_tokens"] = value
			stats(rec)["cost_usd"] = value
			stats(rec)["elapsed_ms"] = value
			res, warn := readText(t, lineOf(t, rec)+"\n")
			if len(res.Records) != 1 || res.Skipped != 0 || warn != "" {
				t.Fatalf("kept=%d skipped=%d warn=%q, want the record kept", len(res.Records), res.Skipped, warn)
			}
			got := res.Records[0]
			if got.InputTokens != 0 || got.CostUSD != 0 || got.ElapsedMS != 0 {
				t.Errorf("fields = in %d cost %v ms %d, want zeros", got.InputTokens, got.CostUSD, got.ElapsedMS)
			}
		})
	}
}

func TestWholeFloatTokensAccepted(t *testing.T) {
	rec := baseRecord("a")
	stats(rec)["output_tokens"] = 12.0
	stats(rec)["input_tokens"] = 12.5
	got := readOne(t, rec)
	if got.OutputTokens != 12 {
		t.Errorf("OutputTokens = %d, want 12", got.OutputTokens)
	}
	if got.InputTokens != 0 {
		t.Errorf("InputTokens = %d, want 0 for a fractional count", got.InputTokens)
	}
}

func TestReturncodeNullOrAbsentIsAbsent(t *testing.T) {
	result := func(rec map[string]any) map[string]any { return rec["result"].(map[string]any) }

	nullCode := baseRecord("a")
	result(nullCode)["returncode"] = nil
	if got := readOne(t, nullCode); got.ReturnCode != nil {
		t.Errorf("null returncode: ReturnCode = %d, want nil", *got.ReturnCode)
	}

	absentCode := baseRecord("b")
	delete(result(absentCode), "returncode")
	if got := readOne(t, absentCode); got.ReturnCode != nil {
		t.Errorf("absent returncode: ReturnCode = %d, want nil", *got.ReturnCode)
	}

	zeroCode := baseRecord("c")
	result(zeroCode)["returncode"] = 0
	if got := readOne(t, zeroCode); got.ReturnCode == nil || *got.ReturnCode != 0 {
		t.Errorf("returncode 0: ReturnCode = %v, want 0", got.ReturnCode)
	}
}

func TestTimedOutDefaultsFalse(t *testing.T) {
	absent := baseRecord("a")
	delete(absent["result"].(map[string]any), "timed_out")
	if got := readOne(t, absent); got.TimedOut {
		t.Error("absent timed_out: TimedOut = true, want false")
	}

	notBool := baseRecord("b")
	notBool["result"].(map[string]any)["timed_out"] = "yes"
	if got := readOne(t, notBool); got.TimedOut {
		t.Error("string timed_out: TimedOut = true, want false")
	}

	set := baseRecord("c")
	set["result"].(map[string]any)["timed_out"] = true
	if got := readOne(t, set); !got.TimedOut {
		t.Error("timed_out true: TimedOut = false, want true")
	}
}
