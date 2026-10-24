package report

import (
	"slices"
	"sort"
	"strconv"

	"drydockstats/internal/logrecord"
)

// totals accumulates the numeric columns shared by every report.
type totals struct {
	calls, failed                    int64
	input, cached, output, elapsedMS int64
	costUSD                          float64
}

func (t *totals) add(r logrecord.Record) {
	t.calls++
	if !r.Succeeded() {
		t.failed++
	}
	t.input += r.InputTokens
	t.cached += r.CachedTokens
	t.output += r.OutputTokens
	t.elapsedMS += r.ElapsedMS
	t.costUSD += r.CostUSD
}

func (t totals) tokens() int64 {
	return t.input + t.cached + t.output
}

func sumOf(records []logrecord.Record) totals {
	var t totals
	for _, r := range records {
		t.add(r)
	}
	return t
}

// groupBy sums records into one totals value per key.
func groupBy[K comparable](records []logrecord.Record, key func(logrecord.Record) K) map[K]*totals {
	groups := make(map[K]*totals)
	for _, r := range records {
		k := key(r)
		g, ok := groups[k]
		if !ok {
			g = &totals{}
			groups[k] = g
		}
		g.add(r)
	}
	return groups
}

func itoa(v int64) string {
	return strconv.FormatInt(v, 10)
}

var summaryHeader = []string{"records", "skipped", "succeeded", "failed", "first", "last",
	"input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"}

// Summary writes one data row over the filtered records; skipped counts the whole file.
func Summary(records []logrecord.Record, skipped int) [][]string {
	t := sumOf(records)
	first, last := "", ""
	for i, r := range records {
		date := r.StartDate()
		if i == 0 || date < first {
			first = date
		}
		if i == 0 || date > last {
			last = date
		}
	}
	return [][]string{
		summaryHeader,
		{itoa(t.calls), itoa(int64(skipped)), itoa(t.calls - t.failed), itoa(t.failed), first, last,
			itoa(t.input), itoa(t.cached), itoa(t.output), FormatCost(t.costUSD), FormatMinutes(t.elapsedMS)},
	}
}

var commandsHeader = []string{"command", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "minutes"}

// Commands writes one row per job.command_name, most calls first, then command ascending.
func Commands(records []logrecord.Record) [][]string {
	groups := groupBy(records, func(r logrecord.Record) string { return r.Command })
	names := make([]string, 0, len(groups))
	for name := range groups {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		a, b := groups[names[i]].calls, groups[names[j]].calls
		if a != b {
			return a > b
		}
		return names[i] < names[j]
	})

	rows := [][]string{commandsHeader}
	for _, name := range names {
		g := groups[name]
		rows = append(rows, []string{name, itoa(g.calls), itoa(g.failed), itoa(g.input), itoa(g.cached),
			itoa(g.output), FormatMinutes(g.elapsedMS)})
	}
	t := sumOf(records)
	return append(rows, []string{"TOTAL", itoa(t.calls), itoa(t.failed), itoa(t.input), itoa(t.cached),
		itoa(t.output), FormatMinutes(t.elapsedMS)})
}

var targetsHeader = []string{"target", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd", "minutes"}

// Targets writes one row per job.target, ordered byte-wise ascending.
func Targets(records []logrecord.Record) [][]string {
	groups := groupBy(records, func(r logrecord.Record) string { return r.Target })
	names := make([]string, 0, len(groups))
	for name := range groups {
		names = append(names, name)
	}
	slices.Sort(names)

	rows := [][]string{targetsHeader}
	for _, name := range names {
		g := groups[name]
		rows = append(rows, []string{name, itoa(g.calls), itoa(g.failed), itoa(g.input), itoa(g.cached),
			itoa(g.output), FormatCost(g.costUSD), FormatMinutes(g.elapsedMS)})
	}
	t := sumOf(records)
	return append(rows, []string{"TOTAL", itoa(t.calls), itoa(t.failed), itoa(t.input), itoa(t.cached),
		itoa(t.output), FormatCost(t.costUSD), FormatMinutes(t.elapsedMS)})
}

var dailyHeader = []string{"date", "calls", "failed", "tokens", "minutes"}

// Daily writes one row per start date, ordered ascending.
func Daily(records []logrecord.Record) [][]string {
	groups := groupBy(records, func(r logrecord.Record) string { return r.StartDate() })
	dates := make([]string, 0, len(groups))
	for date := range groups {
		dates = append(dates, date)
	}
	slices.Sort(dates)

	rows := [][]string{dailyHeader}
	for _, date := range dates {
		g := groups[date]
		rows = append(rows, []string{date, itoa(g.calls), itoa(g.failed), itoa(g.tokens()), FormatMinutes(g.elapsedMS)})
	}
	t := sumOf(records)
	return append(rows, []string{"TOTAL", itoa(t.calls), itoa(t.failed), itoa(t.tokens()), FormatMinutes(t.elapsedMS)})
}

var modelsHeader = []string{"provider", "model", "calls", "failed", "input_tokens", "cached_tokens", "output_tokens", "cost_usd"}

// providerModel is the grouping key of the models report.
type providerModel struct {
	provider, model string
}

// Models writes one row per provider and attributed model, ordered by provider then model.
func Models(records []logrecord.Record) [][]string {
	groups := groupBy(records, func(r logrecord.Record) providerModel {
		return providerModel{provider: r.Provider, model: r.Model}
	})
	keys := make([]providerModel, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].provider != keys[j].provider {
			return keys[i].provider < keys[j].provider
		}
		return keys[i].model < keys[j].model
	})

	rows := [][]string{modelsHeader}
	for _, k := range keys {
		g := groups[k]
		rows = append(rows, []string{k.provider, k.model, itoa(g.calls), itoa(g.failed), itoa(g.input),
			itoa(g.cached), itoa(g.output), FormatCost(g.costUSD)})
	}
	t := sumOf(records)
	return append(rows, []string{"TOTAL", "", itoa(t.calls), itoa(t.failed), itoa(t.input), itoa(t.cached),
		itoa(t.output), FormatCost(t.costUSD)})
}

var failuresHeader = []string{"started_at", "target", "command", "model", "returncode", "timed_out"}

// Failures writes one row per failed record, ordered by started_at then execution_id. It has no TOTAL row.
func Failures(records []logrecord.Record) [][]string {
	var failed []logrecord.Record
	for _, r := range records {
		if !r.Succeeded() {
			failed = append(failed, r)
		}
	}
	sort.SliceStable(failed, func(i, j int) bool {
		if failed[i].StartedAt != failed[j].StartedAt {
			return failed[i].StartedAt < failed[j].StartedAt
		}
		return failed[i].ExecutionID < failed[j].ExecutionID
	})

	rows := [][]string{failuresHeader}
	for _, r := range failed {
		returnCode := ""
		if r.ReturnCode != nil {
			returnCode = itoa(*r.ReturnCode)
		}
		timedOut := "false"
		if r.TimedOut {
			timedOut = "true"
		}
		rows = append(rows, []string{r.StartedAt, r.Target, r.Command, r.Model, returnCode, timedOut})
	}
	return rows
}
