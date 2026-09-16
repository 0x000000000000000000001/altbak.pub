#!/usr/bin/env python3
"""Descriptive bootstrap of paired median best-of-ten gains from saved runs."""
from pathlib import Path
import json
import random
import statistics

root = Path(__file__).resolve().parent
out = {}
for phase in ["timing", "timing-checks"]:
    data = json.loads((root / phase / "results.json").read_text())
    rounds = data["rounds"]
    stats = {}
    for name in data["summary"]:
        gains = [100 * (1 - r["samples"][name]["min_us"] /
                        r["samples"]["generated"]["min_us"]) for r in rounds]
        rng = random.Random(16092026)
        boot = sorted(statistics.median(rng.choices(gains, k=len(gains)))
                      for _ in range(20000))
        stats[name] = {"median_gain_pct": statistics.median(gains),
                       "bootstrap_median_95_pct": [boot[500], boot[19499]],
                       "rounds_faster": sum(g > 0 for g in gains), "rounds": len(gains)}
    out[phase] = stats
(root / "bootstrap.json").write_text(json.dumps(out, indent=2) + "\n")
print(json.dumps(out, indent=2))
