#!/usr/bin/env python3
"""Paired runs of frozen official binaries; validate every benchmark result."""
import argparse
import hashlib
import json
from pathlib import Path
import random
import statistics
import subprocess
import sys

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / "bin/benchmark"))
from validate import CASES, validate_output

parser = argparse.ArgumentParser()
parser.add_argument("--variant", nargs=2, action="append", required=True,
                    metavar=("NAME", "BINARY"))
parser.add_argument("--rounds", type=int, default=21)
parser.add_argument("--seed", type=int, default=91626)
parser.add_argument("--output", type=Path, required=True)
args = parser.parse_args()
variants = {name: Path(binary).resolve() for name, binary in args.variant}
assert len(variants) == len(args.variant) and len(variants) > 1
baseline = next(iter(variants))
rng = random.Random(args.seed)
args.output.mkdir(parents=True, exist_ok=True)
rows = []
for rnd in range(args.rounds):
    order = list(variants)
    rng.shuffle(order)
    pair = {}
    for name in order:
        result = subprocess.run([str(variants[name])], cwd=ROOT,
                                check=True, text=True, capture_output=True)
        pair[name] = validate_output(result.stdout, "pure")
        (args.output / f"{rnd:02}-{name}.log").write_text(result.stdout)
    rows.append({"round": rnd, "order": order, "results": pair})
    print(rnd, *(f"{name}: RB={pair[name]['times_us'][8]:.2f}us "
                 f"sum={pair[name]['sum_displayed_lines_ms']:.5f}ms"
                 for name in variants), flush=True)

summary = {}
for i, case in enumerate([*CASES, "total_ms"]):
    values = {
        name: [(r["results"][name]["times_us"][i] if i < len(CASES)
                else r["results"][name]["sum_displayed_lines_ms"])
               for r in rows] for name in variants
    }
    summary[case] = {name: {"median": statistics.median(vals)}
                     for name, vals in values.items()}
    for name, vals in values.items():
        ratios = [v / b for v, b in zip(vals, values[baseline])]
        summary[case][name]["paired_change_pct"] = 100 * (statistics.median(ratios) - 1)
data = {
    "command": sys.argv,
    "baseline": baseline,
    "binaries": {n: {"path": str(p), "sha256": hashlib.sha256(p.read_bytes()).hexdigest()}
                 for n, p in variants.items()},
    "protocol": "Official native MiMalloc pure runner, global warmup, best of 10 per case; shuffled paired rounds; every output checked by official validator",
    "rounds": rows,
    "summary": summary,
}
(args.output / "results.json").write_text(json.dumps(data, indent=2) + "\n")
print(json.dumps(summary, indent=2))
