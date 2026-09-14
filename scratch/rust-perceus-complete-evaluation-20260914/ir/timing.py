"""Native timing preparation; no clocks execute unless explicitly given time."""
from pathlib import Path
import argparse
import hashlib
import json
import os
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
BUILD = HERE/'build/timing'
STAGES = ['naive', 'pushdown', 'precise', 'fusion', 'specialized_drop', 'reuse', 'retained_fields']

def sha(path): return hashlib.sha256(path.read_bytes()).hexdigest()

def build():
    manifest = json.loads((HERE/'metadata.json').read_text())['files_sha256']
    (BUILD/'src/bin').mkdir(parents=True,exist_ok=True)
    for stage in STAGES:
        source = HERE/f'build/{stage}.rs'
        assert sha(source)==manifest[f'build/{stage}.rs'], 'Validated source changed'
        text = source.read_text()
        prefix = text[:text.index('fn main() {')]
        (BUILD/f'src/bin/{stage}.rs').write_text(prefix+(HERE/'timing-main.rs').read_text())
    (BUILD/'Cargo.toml').write_text('''[workspace]
[package]
name = "perceus_ir_timing"
version = "0.1.0"
edition = "2021"
[profile.release]
opt-level = 1
debug = true
[dependencies]
mimalloc = "0.1.32"
''')
    env = dict(os.environ, RUSTFLAGS='--cfg untracked')
    subprocess.run(['cargo','build','--release','--offline'],cwd=BUILD,env=env,check=True)
    metadata = {'profile':'O1, debug=true, mimalloc, --cfg untracked direct std::rc::Rc; same generated IR functions',
                'matrix':'5 functions x unique/shared-root/shared-child; fresh input and retained alias every iteration',
                'boundary':'construction + transformation + result checksum traversal + destruction of every result/input/retained alias',
                'stage_order':'Seven cyclic rotations: each stage appears once at every position',
                'validation':'Analytic checksum/count outside timed region; runtime seed and black_box input/size/result',
                'variants':{}, 'timings_run':False}
    for stage in STAGES:
        binary = BUILD/f'target/release/{stage}'
        output = subprocess.check_output([str(binary),'smoke'],text=True).strip()
        metadata['variants'][stage] = {'binary_sha256':sha(binary), 'source_sha256':sha(BUILD/f'src/bin/{stage}.rs'),
                                      'validated_ir_source_sha256':sha(HERE/f'build/{stage}.rs'), 'smoke':output}
        print(stage,output,flush=True)
    metadata['cargo_lock_sha256']=sha(BUILD/'Cargo.lock')
    (HERE/'timing-build.json').write_text(json.dumps(metadata,indent=2)+'\n')

def measure(args):
    metadata=json.loads((HERE/'timing-build.json').read_text())
    for stage in STAGES: assert sha(BUILD/f'target/release/{stage}')==metadata['variants'][stage]['binary_sha256']
    raw={stage:{} for stage in STAGES}; order_log=[]
    for block in range(args.blocks):
        offset=block%len(STAGES); order=STAGES[offset:]+STAGES[:offset]
        if block//len(STAGES)%2: order=list(reversed(order))
        order_log.append(order)
        for stage in order:
            command=[str(BUILD/f'target/release/{stage}'),'measure',str(args.samples),str(17+block*97),str(args.divisor)]
            process=subprocess.Popen(command,stdout=subprocess.PIPE,text=True)
            lines=[]; current={}
            for line in process.stdout:
                lines.append(line); tag,case,mode,sample,loops,ns=line.split()
                assert tag=='TIME'
                key=f'{case}/{mode}'
                current.setdefault(key,[]).append(int(ns))
                if int(sample)+1==args.samples:
                    print(f'block {block+1}/{args.blocks} {stage} {key}: {statistics.median(current[key])/1e6:.3f} ms ({loops} iterations)',flush=True)
            assert process.wait()==0
            assert len(current)==15 and all(len(v)==args.samples for v in current.values())
            (BUILD/f'raw-{block+1}-{stage}.log').write_text(''.join(lines))
            for key,values in current.items(): raw[stage].setdefault(key,[]).append(values)
            checkpoint={'complete':False,'completed_block':block+1,'completed_stage':stage,
                        'blocks':args.blocks,'samples_per_process_case':args.samples,
                        'iteration_divisor':args.divisor,'order':order_log,'raw_ns':raw}
            temporary=HERE/'timings-checkpoint.json.tmp'
            temporary.write_text(json.dumps(checkpoint,indent=2)+'\n')
            temporary.replace(HERE/'timings-checkpoint.json')
    result={'method':metadata, 'blocks':args.blocks,'samples_per_process_case':args.samples,'iteration_divisor':args.divisor,
            'order':order_log,'raw_ns':raw,'median_ms':{stage:{key:statistics.median([statistics.median(row) for row in rows])/1e6 for key,rows in cases.items()} for stage,cases in raw.items()}}
    (HERE/'timings.json').write_text(json.dumps(result,indent=2)+'\n')
    checkpoint['complete']=True
    temporary.write_text(json.dumps(checkpoint,indent=2)+'\n')
    temporary.replace(HERE/'timings-checkpoint.json')
    print('All native stage timings written to timings.json',flush=True)

if __name__=='__main__':
    parser=argparse.ArgumentParser();parser.add_argument('mode',choices=['build','time'])
    parser.add_argument('--blocks',type=int,default=7);parser.add_argument('--samples',type=int,default=3)
    parser.add_argument('--divisor',type=int,default=1)
    args=parser.parse_args()
    assert args.blocks>0 and args.samples>0 and args.divisor>0
    build() if args.mode=='build' else measure(args)
