#!/bin/sh
# run.sh — DrydockStats grader. Do not filter, skip, or reinterpret.
#
# Builds ./cmd/ddkstat with `go build` and compares every report's exit code and standard output
# over the supplied sample log with the known answer below. Standard error is never compared.
#
#   sh sources/run.sh                    from the application root (how `drydock uat` runs it)
#   sh uat/DrydockStats/bin/run.sh APP   from anywhere, grading the application root APP
#
# Exit 0 only when the build succeeds and every check passes.
set -u

HERE=$(cd "$(dirname "$0")" && pwd)
APP=$(cd "${1:-.}" && pwd)
if [ -f "$HERE/llm_sample.jsonl" ]; then DATA=$HERE; else DATA=$(cd "$HERE/../sources" && pwd); fi
WORK=$(mktemp -d)
trap 'rm -rf "$WORK"' EXIT INT TERM
BIN="$WORK/ddkstat"
LOG="$DATA/llm_sample.jsonl"
MISSING="$WORK/no-such-log.jsonl"

if ! (cd "$APP" && go build -o "$BIN" ./cmd/ddkstat); then
    echo "error: go build ./cmd/ddkstat failed at the application root $APP." >&2
    echo "The deliverable is a Go module whose main package is ./cmd/ddkstat." >&2
    exit 1
fi

PASSED=0
FAILED=0

# check <id> <expected exit> <ddkstat args...>  — expected standard output on stdin.
check() {
    id=$1
    want_rc=$2
    shift 2
    cat > "$WORK/want"
    (cd "$APP" && "$BIN" "$@") > "$WORK/got" 2> "$WORK/err"
    rc=$?
    if [ "$rc" -eq "$want_rc" ] && cmp -s "$WORK/want" "$WORK/got"; then
        PASSED=$((PASSED + 1))
        echo "PASS $id: ddkstat $*"
    else
        FAILED=$((FAILED + 1))
        echo "FAIL $id: ddkstat $*"
        echo "  exit: expected $want_rc, got $rc"
        diff "$WORK/want" "$WORK/got" | sed 's/^/  /'
        sed 's/^/  stderr: /' "$WORK/err"
    fi
}

# ── A. Every report over the whole sample ───────────────────────────────────────────────
check A1 0 summary "$LOG" <<'EOF'
records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
33,3,31,2,2026-10-14,2026-10-16,262526,17053317,398580,5.5245,159.4
EOF
check A2 0 commands "$LOG" <<'EOF'
command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
build,17,0,129705,8268539,209318,71.1
score release,5,0,31244,2819917,46278,22.7
plan,4,0,1862,2244487,59759,25.1
analyze,3,0,32267,1662483,40833,20.2
diagnose,2,2,33682,1010463,17499,10.3
import_compass,1,0,843,263315,7751,0.9
plan repair,1,0,32923,784113,17142,9.1
TOTAL,33,2,262526,17053317,398580,159.4
EOF
check A3 0 targets "$LOG" <<'EOF'
target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
CommonMark,12,1,5547,5786766,119677,3.7919,51.5
ReadingList,11,0,4104,5161259,145664,1.2485,47.6
Toml,10,1,252875,6105292,133239,0.4841,60.3
TOTAL,33,2,262526,17053317,398580,5.5245,159.4
EOF
check A4 0 models "$LOG" <<'EOF'
provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
claude,claude-sonnet-5-5,11,1,5547,5786766,119677,3.7919
claude,sonnet,1,0,0,0,0,0.0000
codex,gpt-5.6-luna,4,0,110058,2466012,40962,0.0000
gemini,gemini-3-flash,6,1,142817,3639280,92277,0.4841
TOTAL,,33,2,262526,17053317,398580,5.5245
EOF
check A5 0 daily "$LOG" <<'EOF'
date,calls,failed,tokens,minutes
2026-10-14,14,1,6944443,86.3
2026-10-15,12,1,7016674,57.8
2026-10-16,7,0,3753306,15.3
TOTAL,33,2,17714423,159.4
EOF
check A6 0 failures "$LOG" <<'EOF'
started_at,target,command,model,returncode,timed_out
2026-10-14T02:38:41.904Z,Toml,diagnose,gemini-3-flash,1,false
2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
EOF

# ── B. Filters ───────────────────────────────────────────────────────────────────────────
check B1 0 summary --target Toml --since 2026-10-15 "$LOG" <<'EOF'
records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
3,3,3,0,2026-10-15,2026-10-15,87934,1895532,26630,0.0770,12.0
EOF
check B2 0 commands --until 2026-10-14 "$LOG" <<'EOF'
command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
build,6,0,67524,2290028,87813,37.5
analyze,2,0,31740,1346394,25842,12.0
plan,2,0,467,1010857,35477,15.6
diagnose,1,1,33526,770301,17499,9.8
import_compass,1,0,843,263315,7751,0.9
plan repair,1,0,32923,784113,17142,9.1
score release,1,0,379,110560,9949,1.3
TOTAL,14,1,167402,6575568,201473,86.3
EOF
check B3 0 daily --since 2026-10-15 --until 2026-10-15 "$LOG" <<'EOF'
date,calls,failed,tokens,minutes
2026-10-15,12,1,7016674,57.8
TOTAL,12,1,7016674,57.8
EOF
check B4 0 failures --target CommonMark "$LOG" <<'EOF'
started_at,target,command,model,returncode,timed_out
2026-10-15T06:50:39.621Z,CommonMark,diagnose,claude-sonnet-5-5,,true
EOF
check B5 0 models --target ReadingList "$LOG" <<'EOF'
provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd
claude,claude-haiku-5-5,11,0,4104,5161259,145664,1.2485
TOTAL,,11,0,4104,5161259,145664,1.2485
EOF

# ── C. A filter that matches nothing ─────────────────────────────────────────────────────
check C1 0 commands --target NoSuchTarget "$LOG" <<'EOF'
command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes
TOTAL,0,0,0,0,0,0.0
EOF
check C2 0 summary --target NoSuchTarget "$LOG" <<'EOF'
records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes
0,3,0,0,,,0,0,0,0.0000,0.0
EOF

# ── E. Operational failures ──────────────────────────────────────────────────────────────
check E1 1 summary "$MISSING" < /dev/null

# ── U. Usage errors ──────────────────────────────────────────────────────────────────────
check U1 2 < /dev/null
check U2 2 frobnicate "$LOG" < /dev/null
check U3 2 summary < /dev/null
check U4 2 summary --since 2026-13-01 "$LOG" < /dev/null
check U5 2 summary --colour red "$LOG" < /dev/null
check U6 2 summary "$LOG" "$LOG" < /dev/null

TOTAL=$((PASSED + FAILED))
echo "DrydockStats: $PASSED passed, $FAILED failed (of $TOTAL)"
[ "$FAILED" -eq 0 ]
