#!/usr/bin/env python3
"""Five alternating pairs of the fixed archived probe; no compilation here."""
from pathlib import Path
import argparse,datetime,hashlib,json,os,statistics,subprocess
base=Path(__file__).resolve().parent
parser=argparse.ArgumentParser()
parser.add_argument('--pair',type=int)
parser.add_argument('--summary',action='store_true')
args=parser.parse_args()
logs=base/'measurements'; logs.mkdir(exist_ok=True)
env=dict(os.environ,GOGC='800',GOWORK='off'); env.pop('PPROF',None)
def sha(p): return hashlib.sha256(p.read_bytes()).hexdigest()
if args.pair:
 pair=args.pair
 assert 1<=pair<=5
 out=logs/f'pair-{pair}.json'
 assert not out.exists(), 'Do not overwrite a completed pair'
 order=['before','after'] if pair%2 else ['after','before']
 records={}
 for key in order:
  binary=base/key/'probe/probe'; digest=sha(binary)
  started=datetime.datetime.now(datetime.timezone.utc).isoformat()
  result=subprocess.run([str(binary),'-mode','pure','-n','100000','-iterations','5'],cwd=binary.parent,env=env,capture_output=True,text=True,timeout=30)
  (logs/f'pair-{pair}-{key}.stdout.jsonl').write_text(result.stdout)
  (logs/f'pair-{pair}-{key}.stderr.log').write_text(result.stderr)
  assert result.returncode==0,result.stderr
  assert sha(binary)==digest
  samples=[json.loads(line) for line in result.stdout.splitlines()]
  assert len(samples)==5 and all(x['depth']==22 and x['input']==100000 and x['mode']=='pure' for x in samples)
  records[key]={'binary_sha256':digest,'started_utc':started,'samples':samples,'median_ns':statistics.median(x['nanoseconds'] for x in samples)}
 percent=(records['after']['median_ns']/records['before']['median_ns']-1)*100
 out.write_text(json.dumps({'pair':pair,'order':order,'records':records,'paired_change_percent':percent},indent=2)+'\n')
 print(f"Pair {pair} ({'/'.join(order)}): {records['before']['median_ns']/1e6:.6f} -> {records['after']['median_ns']/1e6:.6f} ms; {percent:+.3f}%",flush=True)
if args.summary:
 pairs=[json.loads((logs/f'pair-{i}.json').read_text()) for i in range(1,6)]
 summary={'protocol':'5 alternating process pairs; n=100000; 3 warmups + 5 samples/process; forced GC before each measured call; GOGC=800; GOWORK=off; PGO off; depth=22 validated for every sample','environment':{k:env.get(k) for k in ['GOGC','GOWORK','GOMAXPROCS','GOMEMLIMIT']},'variants':{}}
 for key in ['before','after']:
  samples=[s for p in pairs for s in p['records'][key]['samples']]
  summary['variants'][key]={'samples':len(samples),'processes':5}
  for field in ['nanoseconds','allocated_bytes','allocations']:
   values=[s[field] for s in samples]
   summary['variants'][key][field]={'median':statistics.median(values),'min':min(values),'max':max(values)}
 changes=[p['paired_change_percent'] for p in pairs]
 summary['paired_change_percent']={'median':statistics.median(changes),'min':min(changes),'max':max(changes),'all':changes}
 a=summary['variants']['after']['nanoseconds']['median']; b=summary['variants']['before']['nanoseconds']['median']
 summary['pooled_median_change_percent']=(a/b-1)*100
 (base/'results.json').write_text(json.dumps(summary,indent=2)+'\n')
 print(json.dumps(summary,indent=2))
