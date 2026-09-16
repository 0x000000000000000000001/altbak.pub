#!/usr/bin/env python3
from pathlib import Path
import argparse
import hashlib
import json
import random
import statistics
import subprocess

base = Path(__file__).resolve().parent
parser = argparse.ArgumentParser()
parser.add_argument('--rounds', type=int, default=21)
parser.add_argument('--seed', type=int, default=916263)
parser.add_argument('--output', default='timings.json')
args = parser.parse_args()
rng = random.Random(args.seed)
cases = [(0, n) for n in ['original', 'record', 'materialized', 'scalar', 'local']]
cases += [(k, n) for k in [100, 1] for n in ['record', 'materialized', 'scalar']]
rows = []
for rnd in range(args.rounds):
    order = cases.copy()
    rng.shuffle(order)
    row = {'round': rnd, 'order': order, 'ns': {}}
    for stride, name in order:
        out = subprocess.check_output([str(base/'probe'), 'time', name, str(stride)], text=True, timeout=15)
        samples = [int(x) for x in out.splitlines()]
        assert len(samples) == 10 and min(samples) > 0
        row['ns'][f'{stride}/{name}'] = samples
    rows.append(row)
    print(rnd, ' '.join(f'{name}={min(row["ns"]["0/"+name])/1000:.2f}us' for name in ['original','record','materialized','scalar','local']), flush=True)
summary = {}
for stride, name in cases:
    key = f'{stride}/{name}'
    best = [min(r['ns'][key]) for r in rows]
    ratios = [min(r['ns'][key])/min(r['ns'][f'{stride}/record']) for r in rows]
    summary[key] = {'median_best_us': statistics.median(best)/1000,
                    'median_sample_us': statistics.median(statistics.median(r['ns'][key]) for r in rows)/1000,
                    'paired_change_vs_record_pct': 100*(statistics.median(ratios)-1),
                    'wins_vs_record': sum(x < 1 for x in ratios)}
result = {'binary_sha256': hashlib.sha256((base/'probe').read_bytes()).hexdigest(),
          'protocol': '3 warmups, 10 samples, 21 shuffled rounds by default; same MiMalloc and non-inlined cross-crate workers, no LTO; destruction timed; no counters',
          'seed': args.seed, 'rounds': rows, 'summary': summary}
(base/args.output).write_text(json.dumps(result, indent=2)+'\n')
print(json.dumps(summary, indent=2))
