#!/usr/bin/env python3
# Generate the DrydockStats fixture: a synthetic schema_version 1 llm.jsonl with known defects.
import json, random, sys
rng = random.Random(20261022)
plans = [
    # (target, provider, requested, executed, cost_per_mtok or None)
    ("ReadingList", "claude", "haiku", "claude-haiku-5-5", 1.0),
    ("Toml", "codex", "gpt-5.6-luna", None, None),
    ("CommonMark", "claude", "sonnet", "claude-sonnet-5-5", 3.0),
    ("Toml", "gemini", "gemini-3-flash", "gemini-3-flash", 0.5),
]
seq = [
    ("2026-10-14", "ReadingList", ["import_compass", "analyze", "plan", "build", "build", "build", "score release"]),
    ("2026-10-14", "Toml", ["analyze", "plan", "plan repair", "build", "build", "diagnose", "build"]),
    ("2026-10-15", "CommonMark", ["analyze", "plan", "build", "build", "build", "build", "diagnose", "build", "score release"]),
    ("2026-10-15", "Toml", ["build", "build", "score release"]),
    ("2026-10-16", "ReadingList", ["plan", "build", "build", "score release"]),
    ("2026-10-16", "CommonMark", ["build", "build", "score release"]),
]
providers = {"ReadingList": [plans[0]], "Toml": [plans[1], plans[3]], "CommonMark": [plans[2]]}
recs = []
n = 0
for day, target, commands in seq:
    hour = rng.randint(1, 9)
    minute = rng.randint(0, 50)
    for cmd in commands:
        n += 1
        tgt, prov, req, exe, price = rng.choice(providers[target])
        sec = rng.randint(0, 59); ms = rng.randint(0, 999)
        started = f"{day}T{hour:02d}:{minute:02d}:{sec:02d}.{ms:03d}Z"
        elapsed = rng.randint(4000, 600000)
        minute += 1 + elapsed // 60000
        if minute >= 60: hour += 1; minute -= 60
        inp = rng.randint(0, 40000) if prov != "claude" else rng.randint(0, 900)
        cached = rng.randint(10000, 900000)
        out = rng.randint(200, 30000)
        cost = None if price is None else round((inp + out * 5 + cached * 0.1) * price / 1e6, 7)
        stats = {"cached_input_tokens": cached, "elapsed_ms": elapsed, "input_tokens": inp,
                 "output_tokens": out, "turns": rng.randint(1, 40), "cost_usd": cost}
        if exe is not None: stats["model"] = exe
        eid = f"{started[:10].replace('-', '')}.{started[11:19].replace(':', '')}.{ms:03d}Z-{rng.getrandbits(32):08x}"
        recs.append({"schema_version": 1, "execution_id": eid, "started_at": started, "status": "succeeded",
                     "job": {"command_name": cmd, "target": target, "llm": prov, "model": req},
                     "prompt": {"bytes": rng.randint(4000, 200000)},
                     "result": {"returncode": 0, "timed_out": False, "stats": stats}})
# Two failures: one nonzero exit, one timeout with no return code.
recs[12]["status"] = "failed"; recs[12]["result"]["returncode"] = 1
recs[20]["status"] = "failed"; recs[20]["result"]["returncode"] = None; recs[20]["result"]["timed_out"] = True
recs[20]["result"]["stats"]["output_tokens"] = 0
# One record with no stats at all (a provider that reported nothing).
del recs[30]["result"]["stats"]
lines = [json.dumps(r, separators=(",", ":")) for r in recs]
# Defects: blank line (ignored), truncated JSON, duplicate execution_id, missing job.target.
lines.insert(8, "")
lines.insert(15, lines[14][: len(lines[14]) // 2])
lines.insert(27, lines[26])
bad = dict(recs[5]); bad = json.loads(json.dumps(bad)); bad["execution_id"] = "20261014.999999.000Z-deadbeef"
del bad["job"]["target"]
lines.insert(33, json.dumps(bad, separators=(",", ":")))
sys.stdout.write("\n".join(lines) + "\n")
