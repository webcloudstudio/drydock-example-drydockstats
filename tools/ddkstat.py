#!/usr/bin/env python3
# Reference implementation of the DrydockStats contract, used only to validate the grader.
import json, re, sys
REPORTS = ("summary", "commands", "targets", "models", "daily", "failures")
def usage(m): print(m, file=sys.stderr); sys.exit(2)
a = sys.argv[1:]
if not a or a[0] not in REPORTS: usage("report")
report = a.pop(0); target = since = until = None
while a and a[0].startswith("--"):
    o = a.pop(0)
    if not a: usage("value")
    v = a.pop(0)
    if o == "--target": target = v
    elif o in ("--since", "--until"):
        if not re.fullmatch(r"\d{4}-\d{2}-\d{2}", v): usage("date")
        import datetime
        try: datetime.date.fromisoformat(v)
        except ValueError: usage("date")
        if o == "--since": since = v
        else: until = v
    else: usage("opt")
if len(a) != 1: usage("log")
try: raw = open(a[0], "rb").read().decode("utf-8")
except OSError as e: print(e, file=sys.stderr); sys.exit(1)
recs = []; seen = set(); skipped = 0
def s(d, *ks):
    for k in ks:
        if not isinstance(d, dict): return None
        d = d.get(k)
    return d
for line in raw.split("\n"):
    if not line.strip(): continue
    try: r = json.loads(line)
    except ValueError: skipped += 1; continue
    ok = isinstance(r, dict) and all(isinstance(x, str) and x for x in (
        r.get("execution_id"), r.get("started_at"), r.get("status"), s(r, "job", "command_name"), s(r, "job", "target")))
    if not ok or r["execution_id"] in seen: skipped += 1; continue
    seen.add(r["execution_id"]); recs.append(r)
def keep(r):
    d = r["started_at"][:10]
    return (target is None or s(r, "job", "target") == target) and (since is None or d >= since) and (until is None or d <= until)
recs = [r for r in recs if keep(r)]
def n(r, k):
    v = s(r, "result", "stats", k); return v if isinstance(v, int) and not isinstance(v, bool) else 0
def cost(r):
    v = s(r, "result", "stats", "cost_usd"); return float(v) if isinstance(v, (int, float)) and not isinstance(v, bool) else 0.0
def failed(r): return r["status"] != "succeeded"
def model(r): return s(r, "result", "stats", "model") or s(r, "job", "model") or ""
def mins(ms): t = (ms + 3000) // 6000; return f"{t // 10}.{t % 10}"
def agg(rs):
    return dict(calls=len(rs), failed=sum(failed(r) for r in rs), inp=sum(n(r, "input_tokens") for r in rs),
                cached=sum(n(r, "cached_input_tokens") for r in rs), out=sum(n(r, "output_tokens") for r in rs),
                cost=sum(cost(r) for r in rs) if rs else 0.0, ms=sum(n(r, "elapsed_ms") for r in rs))
def group(keyf):
    g = {}
    for r in recs: g.setdefault(keyf(r), []).append(r)
    return g
out = []
if report == "summary":
    t = agg(recs); dates = sorted(r["started_at"][:10] for r in recs)
    out.append("records,skipped,succeeded,failed,first,last,input_tokens,cached_tokens,output_tokens,cost_usd,minutes")
    out.append(f"{t['calls']},{skipped},{t['calls'] - t['failed']},{t['failed']},{dates[0] if dates else ''},{dates[-1] if dates else ''},"
               f"{t['inp']},{t['cached']},{t['out']},{t['cost']:.4f},{mins(t['ms'])}")
elif report == "commands":
    out.append("command,calls,failed,input_tokens,cached_tokens,output_tokens,minutes")
    g = group(lambda r: s(r, "job", "command_name"))
    for k in sorted(g, key=lambda k: (-len(g[k]), k.encode())):
        x = agg(g[k]); out.append(f"{k},{x['calls']},{x['failed']},{x['inp']},{x['cached']},{x['out']},{mins(x['ms'])}")
    x = agg(recs); out.append(f"TOTAL,{x['calls']},{x['failed']},{x['inp']},{x['cached']},{x['out']},{mins(x['ms'])}")
elif report == "targets":
    out.append("target,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd,minutes")
    g = group(lambda r: s(r, "job", "target"))
    for k in sorted(g, key=lambda k: k.encode()):
        x = agg(g[k]); out.append(f"{k},{x['calls']},{x['failed']},{x['inp']},{x['cached']},{x['out']},{x['cost']:.4f},{mins(x['ms'])}")
    x = agg(recs); out.append(f"TOTAL,{x['calls']},{x['failed']},{x['inp']},{x['cached']},{x['out']},{x['cost']:.4f},{mins(x['ms'])}")
elif report == "models":
    out.append("provider,model,calls,failed,input_tokens,cached_tokens,output_tokens,cost_usd")
    g = group(lambda r: (s(r, "job", "llm") or "", model(r)))
    for k in sorted(g, key=lambda k: (k[0].encode(), k[1].encode())):
        x = agg(g[k]); out.append(f"{k[0]},{k[1]},{x['calls']},{x['failed']},{x['inp']},{x['cached']},{x['out']},{x['cost']:.4f}")
    x = agg(recs); out.append(f"TOTAL,,{x['calls']},{x['failed']},{x['inp']},{x['cached']},{x['out']},{x['cost']:.4f}")
elif report == "daily":
    out.append("date,calls,failed,tokens,minutes")
    g = group(lambda r: r["started_at"][:10])
    for k in sorted(g):
        x = agg(g[k]); out.append(f"{k},{x['calls']},{x['failed']},{x['inp'] + x['cached'] + x['out']},{mins(x['ms'])}")
    x = agg(recs); out.append(f"TOTAL,{x['calls']},{x['failed']},{x['inp'] + x['cached'] + x['out']},{mins(x['ms'])}")
else:
    out.append("started_at,target,command,model,returncode,timed_out")
    for r in sorted((r for r in recs if failed(r)), key=lambda r: (r["started_at"], r["execution_id"])):
        rc = s(r, "result", "returncode"); to = s(r, "result", "timed_out") is True
        out.append(f"{r['started_at']},{s(r, 'job', 'target')},{s(r, 'job', 'command_name')},{model(r)},{'' if rc is None else rc},{'true' if to else 'false'}")
sys.stdout.write("\n".join(out) + "\n")
