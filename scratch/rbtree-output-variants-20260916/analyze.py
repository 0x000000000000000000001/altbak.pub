#!/usr/bin/env python3
"""Summarize saved paired rounds. Does not execute benchmarks."""
from pathlib import Path
import json
import random
import statistics

root=Path(__file__).resolve().parent
out={}
for phase in ['timing','timing-mimalloc','timing-directional']:
 data=json.loads((root/phase/'results.json').read_text()); stats={}
 for name in data['summary']:
  gains=[100*(1-r['samples'][name]['min_us']/r['samples']['generated']['min_us']) for r in data['rounds']]
  rng=random.Random(16092026)
  boot=sorted(statistics.median(rng.choices(gains,k=len(gains))) for _ in range(20000))
  stats[name]={'median_gain_pct':statistics.median(gains),'bootstrap_median_95_pct':[boot[500],boot[19499]],
               'wins':sum(g>0 for g in gains),'rounds':len(gains)}
 out[phase]=stats
(root/'statistics.json').write_text(json.dumps(out,indent=2)+'\n')
for phase,stats in out.items():
 print(phase)
 for name in ['guard_fused','guard_unique']:
  if name in stats: print(name,stats[name])
