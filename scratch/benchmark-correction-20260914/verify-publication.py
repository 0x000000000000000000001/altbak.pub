#!/usr/bin/env python3
"""Check the published tables against all 90 validated process logs."""
import hashlib
import json
from pathlib import Path
import re
import runpy
import statistics

HERE=Path(__file__).resolve().parent
ROOT=HERE.parents[1]
check=runpy.run_path(str(ROOT/'bin/benchmark/validate.py'))['validate_output']
plan=json.loads((HERE/'measurement-plan.json').read_text())
results=json.loads((HERE/'results.json').read_text())
assert len(plan)==len(results)==30
for item in plan:
 key=item['key']; runs=[]
 for n in range(1,4):
  raw=(HERE/f'measure-{n}-{key}.stdout.log').read_text()
  validated=check(raw,item['mode'])
  archived=json.loads((HERE/f'measure-{n}-{key}.json').read_text())
  assert all(archived[k]==v for k,v in validated.items()), (key,n)
  runs.append(validated)
 median=[statistics.median(values) for values in zip(*(r['times_us'] for r in runs))]
 assert results[key]['times_us']==median,key
 assert abs(results[key]['total_ms']-sum(median)/1000)<1e-9,key
keys=[['js-pure','es-pure','js-ffi','js-fficc'], ['go-pure','psgo-pure','go-ffi','go-fficc']]
keys += [[name+'-'+mode for mode in ['pure','ffi','fficc']] for name in ['scm','erl','php','rust','cpp','sharp','java']]
keys += [['wasm-pure']]
core=(ROOT/'README.md').read_text().split('### Core stresstest benchmark results')[1].split('### Reading the three implementation columns')[0]
tables=re.findall(r'(?m)^.*Benchmark[^\n]*\n(?:.*\|.*\n)+',core)
assert len(tables)==len(keys)==10
for table,columns in zip(tables,keys):
 lines=table.splitlines()[2:]
 assert len(lines)==15
 for index,line in enumerate(lines):
  values=re.findall(r'~ ([0-9.]+) (μs|ms)',line)
  assert len(values)==len(columns)
  for key,(value,unit) in zip(columns,values):
   expected=results[key]['times_us'][index] if index<14 else results[key]['total_ms']
   assert value==f'{expected:.2f}',(key,index,value,expected)
   assert unit==('μs' if index<14 else 'ms')
frozen=json.loads((HERE/'corrected-source-sha256.json').read_text())
assert all((ROOT/name).is_file() and hashlib.sha256((ROOT/name).read_bytes()).hexdigest()==h for name,h in frozen.items())
print('PASS: 90 raw logs, 1260 values, all aggregates, 30 README columns, 420 timings and 30 totals; benchmark sources unchanged since measurement freeze.')
