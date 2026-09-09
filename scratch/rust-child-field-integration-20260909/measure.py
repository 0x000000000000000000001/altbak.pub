"""Measure saved full runners after real code generation, without rebuilding."""
from pathlib import Path
import argparse
import hashlib
import json
import re
import statistics
import subprocess

HERE=Path(__file__).resolve().parent
BUILD=HERE/'build'
EXPECTED=['7','55','202950','100000','20000','125','100000','21536','22','10000000','1200','1000000','202950','5']
PATTERN=re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+\(Execution time - best of 10\)\s+([0-9.]+) μs')
def parse(output):
    rows=PATTERN.findall(output)
    assert [v for _,v,_ in rows]==EXPECTED,rows
    return {n:float(us) for n,_,us in rows}

if __name__=='__main__':
    parser=argparse.ArgumentParser()
    parser.add_argument('variants',nargs='+')
    parser.add_argument('--output',required=True)
    args=parser.parse_args()
    runs={n:[] for n in args.variants}
    for pair in range(5):
        names=args.variants
        order=(names if pair%2==0 else names[::-1]) if len(names)==2 else names[pair%len(names):]+names[:pair%len(names)]
        if len(names)>2 and pair%2:order=order[::-1]
        for name in order:
            output=subprocess.check_output([str(BUILD/f'runner-{name}')],text=True)
            (BUILD/f'{Path(args.output).stem}-{name}-{pair+1}.log').write_text(output)
            values=parse(output)
            runs[name].append(values)
            print(pair+1,name,sum(values.values())/1000,'ms; 14 results checked',flush=True)
    labels=list(next(iter(runs.values()))[0])
    medians={name:{k:statistics.median(row[k] for row in rows) for k in labels} for name,rows in runs.items()}
    result={'method':'Five paired/rotating blocks of complete runners; unchanged warmup and best of 10 per benchmark; all 14 outputs verified each process; totals sum benchmark medians. O1/mimalloc. No concurrent compilation or instrumentation.',
        'runs_us':runs,'median_us':medians,'total_us':{n:sum(v.values()) for n,v in medians.items()},
        'binary_sha256':{n:hashlib.sha256((BUILD/f'runner-{n}').read_bytes()).hexdigest() for n in args.variants}}
    (HERE/args.output).write_text(json.dumps(result,indent=2)+'\n')
    key='Red-Black Tree (100k Worst-Case Insertions):'
    print(json.dumps({n:{'RBTree_ms':medians[n][key]/1000,'total_ms':result['total_us'][n]/1000} for n in args.variants},indent=2))
