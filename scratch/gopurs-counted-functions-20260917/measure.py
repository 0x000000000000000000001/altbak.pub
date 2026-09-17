#!/usr/bin/env python3
from pathlib import Path
import os, subprocess, json, statistics, hashlib, sys
base=Path(__file__).resolve().parent
variant=sys.argv[1] if len(sys.argv)>1 else 'prototype'
logs=base/('measurements-'+variant)
logs.mkdir(exist_ok=False)
env=dict(os.environ,GOGC='800',GOWORK='off')
for key in ['PPROF']:env.pop(key,None)
runs=[]
for pair in range(1,6):
 for name in (['before',variant] if pair%2 else [variant,'before']):
  binary=base/name/'probe'
  result=subprocess.run([str(binary)],cwd=binary.parent,env=env,capture_output=True,text=True,timeout=30)
  (logs/f'{pair}-{name}.stdout.json').write_text(result.stdout)
  (logs/f'{pair}-{name}.stderr.log').write_text(result.stderr)
  assert result.returncode==0,result.stderr
  samples=json.loads(result.stdout)
  assert len(samples)==7 and all(s['result']==100000 for s in samples)
  row=dict(pair=pair,name=name,sha256=hashlib.sha256(binary.read_bytes()).hexdigest(),samples=samples,median_ns=statistics.median(s['ns_per_call'] for s in samples))
  runs.append(row)
 a,b=[next(r for r in runs if r['pair']==pair and r['name']==name) for name in ['before',variant]]
 print(f"pair {pair}: {a['median_ns']/1000:.3f} -> {b['median_ns']/1000:.3f} us ({100*(b['median_ns']/a['median_ns']-1):+.3f}%)",flush=True)
summary={name:{key:statistics.median(s[key] for r in runs if r['name']==name for s in r['samples']) for key in ['ns_per_call','allocs_per_call','bytes_per_call']} for name in ['before',variant]}
ratios=[next(r['median_ns'] for r in runs if r['pair']==p and r['name']==variant)/next(r['median_ns'] for r in runs if r['pair']==p and r['name']=='before')-1 for p in range(1,6)]
report=dict(protocol='5 alternating pairs; 10 warmups/process; 7 batches of 100 calls; GC before each batch outside timer; GOGC=800; PGO off',runs=runs,summary=summary,paired_median_change=statistics.median(ratios))
(base/('results-'+variant+'.json')).write_text(json.dumps(report,indent=2)+'\n')
print(json.dumps(dict(summary=summary,paired_median_change=statistics.median(ratios)),indent=2))
