// Package report aggregates filtered log records into the six CSV reports.
//
// Every function returns rows with the header first. Values are computed from the records passed in;
// nothing is hardcoded.
package report

import (
	"sort"
	"strconv"

	"ddkstat/internal/format"
	"ddkstat/internal/log"
)

// tally sums the counted columns of a group of records.
type tally struct {
	calls, failed, input, cached, output, elapsedMS int64
	cost                                            float64
}

func (t *tally) add(r log.Record) {
	t.calls++
	if r.Failed() {
		t.failed++
	}
	t.input += r.InputTokens
	t.cached += r.CachedTokens
	t.output += r.OutputTokens
	t.cost += r.CostUSD
	t.elapsedMS += r.ElapsedMS
}

// tokens is input plus cached plus output tokens.
func (t tally) tokens() int64 {
	return t.input + t.cached + t.output
}

// counts returns calls, failed, input, cached, and output as cells.
func (t tally) counts() []string {
	return []string{itoa(t.calls), itoa(t.failed), itoa(t.input), itoa(t.cached), itoa(t.output)}
}

func tallyOf(recs []log.Record) tally {
	var t tally
	for _, r := range recs {
		t.add(r)
	}
	return t
}

// groupBy tallies records under the key each one maps to.
func groupBy[K comparable](recs []log.Record, key func(log.Record) K) map[K]*tally {
	groups := make(map[K]*tally)
	for _, r := range recs {
		k := key(r)
		t, ok := groups[k]
		if !ok {
			t = &tally{}
			groups[k] = t
		}
		t.add(r)
	}
	return groups
}

// sortedKeys returns the keys of groups in the order less defines.
func sortedKeys[K comparable](groups map[K]*tally, less func(a, b K) bool) []K {
	keys := make([]K, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return less(keys[i], keys[j]) })
	return keys
}

// Summary writes the one-row summary. skipped is the whole-file skipped count.
func Summary(recs []log.Record, skipped int) [][]string {
	t := tallyOf(recs)
	var first, last string
	for i, r := range recs {
		if i == 0 || r.StartDate < first {
			first = r.StartDate
		}
		if i == 0 || r.StartDate > last {
			last = r.StartDate
		}
	}
	return [][]string{
		{"records", "skipped", "succeeded", "failed", "first", "last", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"},
		{
			itoa(t.calls), itoa(int64(skipped)), itoa(t.calls - t.failed), itoa(t.failed), first, last,
			itoa(t.input), itoa(t.cached), itoa(t.output), format.Cost(t.cost), format.Minutes(t.elapsedMS),
		},
	}
}

// Commands writes one row per command, ordered by calls descending then command ascending, and a TOTAL row.
func Commands(recs []log.Record) [][]string {
	groups := groupBy(recs, func(r log.Record) string { return r.Command })
	names := sortedKeys(groups, func(a, b string) bool { return a < b })
	sort.SliceStable(names, func(i, j int) bool { return groups[names[i]].calls > groups[names[j]].calls })

	rows := [][]string{{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"}}
	for _, name := range names {
		t := groups[name]
		rows = append(rows, append(append([]string{name}, t.counts()...), format.Minutes(t.elapsedMS)))
	}
	total := tallyOf(recs)
	return append(rows, append(append([]string{"TOTAL"}, total.counts()...), format.Minutes(total.elapsedMS)))
}

// Targets writes one row per Target, ordered ascending, and a TOTAL row.
func Targets(recs []log.Record) [][]string {
	groups := groupBy(recs, func(r log.Record) string { return r.Target })
	names := sortedKeys(groups, func(a, b string) bool { return a < b })

	rows := [][]string{{"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"}}
	for _, name := range names {
		t := groups[name]
		rows = append(rows, append(append(append([]string{name}, t.counts()...), format.Cost(t.cost)), format.Minutes(t.elapsedMS)))
	}
	total := tallyOf(recs)
	return append(rows, append(append(append([]string{"TOTAL"}, total.counts()...), format.Cost(total.cost)), format.Minutes(total.elapsedMS)))
}

// modelKey is one provider and model pair.
type modelKey struct {
	provider, model string
}

// Models writes one row per provider and model pair, ordered by provider then model, and a TOTAL row
// whose provider and model cells are empty.
func Models(recs []log.Record) [][]string {
	groups := groupBy(recs, func(r log.Record) modelKey { return modelKey{r.Provider, r.Model} })
	keys := sortedKeys(groups, func(a, b modelKey) bool {
		if a.provider != b.provider {
			return a.provider < b.provider
		}
		return a.model < b.model
	})

	rows := [][]string{{"provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"}}
	for _, k := range keys {
		t := groups[k]
		rows = append(rows, append(append([]string{k.provider, k.model}, t.counts()...), format.Cost(t.cost)))
	}
	total := tallyOf(recs)
	return append(rows, append(append([]string{"TOTAL", ""}, total.counts()...), format.Cost(total.cost)))
}

// Daily writes one row per start date, ordered ascending, and a TOTAL row.
func Daily(recs []log.Record) [][]string {
	groups := groupBy(recs, func(r log.Record) string { return r.StartDate })
	dates := sortedKeys(groups, func(a, b string) bool { return a < b })

	rows := [][]string{{"date", "calls", "failed", "tokens", "minutes"}}
	for _, d := range dates {
		t := groups[d]
		rows = append(rows, []string{d, itoa(t.calls), itoa(t.failed), itoa(t.tokens()), format.Minutes(t.elapsedMS)})
	}
	total := tallyOf(recs)
	return append(rows, []string{"TOTAL", itoa(total.calls), itoa(total.failed), itoa(total.tokens()), format.Minutes(total.elapsedMS)})
}

// Failures writes one row per failed record ordered by started_at then execution_id. It has no TOTAL row.
func Failures(recs []log.Record) [][]string {
	var failed []log.Record
	for _, r := range recs {
		if r.Failed() {
			failed = append(failed, r)
		}
	}
	sort.Slice(failed, func(i, j int) bool {
		if failed[i].StartedAt != failed[j].StartedAt {
			return failed[i].StartedAt < failed[j].StartedAt
		}
		return failed[i].ExecutionID < failed[j].ExecutionID
	})

	rows := [][]string{{"started_at", "target", "command", "model", "returncode", "timed_out"}}
	for _, r := range failed {
		returncode := ""
		if r.ReturnCode != nil {
			returncode = itoa(*r.ReturnCode)
		}
		rows = append(rows, []string{r.StartedAt, r.Target, r.Command, r.Model, returncode, strconv.FormatBool(r.TimedOut)})
	}
	return rows
}

func itoa(n int64) string {
	return strconv.FormatInt(n, 10)
}
