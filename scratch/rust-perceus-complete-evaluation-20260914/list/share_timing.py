#!/usr/bin/env python3
"""Separate plain-module sharing harness. Time requires explicit coordination."""
from pathlib import Path
import argparse
import hashlib
import json
import re
import statistics
import subprocess

HERE=Path(__file__).resolve().parent
DEPS=HERE.parent/'build/cargo-target/release/deps'
FROZEN=HERE.parent/'build/frozen'
BUILD=HERE/'build/sharing'
META=HERE/'sharing-build.json'
VARIANTS=['baseline','lifetimes','consuming','reuse']
SCENARIOS=['unique','root-retained','interior-retained','weak']
def digest(path):return hashlib.sha256(path.read_bytes()).hexdigest()
def save(path,value):path.write_text(json.dumps(value,indent=2)+'\n')
def run(command,log):
    result=subprocess.run(list(map(str,command)),capture_output=True,text=True)
    log.write_text(result.stdout+'\n'+result.stderr)
    if result.returncode:raise RuntimeError(f'{command}: {result.returncode}\n{result.stdout}\n{result.stderr}')
    return result.stdout
def build():
    BUILD.mkdir(parents=True,exist_ok=True)
    initial=json.loads((HERE/'metadata.json').read_text())
    manifest=(FROZEN/'Purs_Test_ListOps/Cargo.toml').read_text().split('[dependencies]')[1]
    names=[name.replace('-','_') for name in re.findall(r'^([\w-]+)\s*=',manifest,re.M)]+['mimalloc']
    externs=[];dependencies={}
    for name in names:
        paths=list(DEPS.glob(f'lib{name}-*.rlib'));assert len(paths)==1,(name,paths)
        externs+=['--extern',f'{name}={paths[0]}'];dependencies[str(paths[0])]=digest(paths[0])
    flags=['--edition=2021','-C','opt-level=1','-C','debuginfo=2','-L',f'dependency={DEPS}']
    metadata={'dependencies':dependencies,'profile':'O1, debuginfo=2, mimalloc from complete frozen runner; no counter instrumentation',
              'harness_sha256':digest(HERE/'share-timing.rs'),'script_sha256':digest(Path(__file__)),
              'rustc':subprocess.check_output(['rustc','--version'],text=True).strip(),'variants':{},
              'timed_scope':'Input construction, owner/Weak setup, filter, borrowed output sum, output/retained owner/Weak destruction. One full warmup batch then seven timed batches per process.'}
    for variant in VARIANTS:
        source=HERE/f'ListOps-{variant}.rs';assert digest(source)==initial['variants'][variant]['plain']
        library=BUILD/f'lib{variant}.rlib';binary=BUILD/f'sharing-{variant}'
        run(['rustc',*flags,'--crate-type=rlib','--crate-name=Purs_Test_ListOps','-C',f'metadata=sharing_{variant}',source,'-o',library,*externs],BUILD/f'compile-{variant}.log')
        run(['rustc',*flags,HERE/'share-timing.rs','-o',binary,'--extern',f'Purs_Test_ListOps={library}',*externs],BUILD/f'compile-harness-{variant}.log')
        metadata['variants'][variant]={'source':str(source),'source_sha256':digest(source),'binary':str(binary),'binary_sha256':digest(binary)}
        print('compiled sharing',variant,flush=True)
    save(META,metadata)
def verified():
    metadata=json.loads(META.read_text())
    assert digest(HERE/'share-timing.rs')==metadata['harness_sha256']
    assert digest(Path(__file__))==metadata['script_sha256']
    for item in metadata['variants'].values():
        assert digest(Path(item['source']))==item['source_sha256']
        assert digest(Path(item['binary']))==item['binary_sha256']
    return metadata
def smoke():
    metadata=verified();results=[]
    for variant in VARIANTS:
        output=run([metadata['variants'][variant]['binary'],variant,'smoke'],BUILD/f'smoke-{variant}.log')
        row=json.loads(output);assert row['smoke_passed'] and not row['clock_read'];results.append(row)
        print('sharing smoke',variant,row['cases'],'cases; no clock',flush=True)
    save(HERE/'sharing-smoke.json',results)
def measure(rounds):
    metadata=verified();rows=[];orders=[]
    profiles=[(n,scenario) for n in [900,9000] for scenario in SCENARIOS]
    for round_index in range(rounds):
        for profile_index,(n,scenario) in enumerate(profiles):
            shift=(round_index+profile_index)%4
            order=VARIANTS[shift:]+VARIANTS[:shift]
            orders.append({'round':round_index,'n':n,'scenario':scenario,'order':order})
            for variant in order:
                iterations=100 if n==900 else 10
                output=run([metadata['variants'][variant]['binary'],variant,'time','--coordinated',n,scenario,iterations],BUILD/f'time-r{round_index}-{n}-{scenario}-{variant}.log')
                row=json.loads(output);row['round']=round_index;rows.append(row)
            save(HERE/'sharing-timings.json',{'build_metadata_sha256':digest(META),'orders':orders,'runs':rows})
            print('sharing measured',round_index,n,scenario,flush=True)
    summaries=[]
    for n,scenario in profiles:
        medians={variant:[statistics.median(r['samples_ns'])/r['iterations'] for r in rows if r['variant']==variant and r['n']==n and r['scenario']==scenario] for variant in VARIANTS}
        medians={key:statistics.median(value) for key,value in medians.items()}
        summaries.append({'n':n,'scenario':scenario,'median_process_median_ns_per_operation':medians,'delta_percent_from_baseline':{key:100*(value/medians['baseline']-1) for key,value in medians.items()}})
    save(HERE/'sharing-summary.json',{'rounds':rounds,'samples_per_process':7,'summaries':summaries})
if __name__=='__main__':
    parser=argparse.ArgumentParser();parser.add_argument('command',choices=['build','smoke','time'])
    parser.add_argument('--coordinated',action='store_true');parser.add_argument('--rounds',type=int,default=4)
    args=parser.parse_args()
    if args.command=='time':
        assert args.coordinated,'time needs --coordinated and external coordinator permission'
        assert 1<=args.rounds<=8
        measure(args.rounds)
    else:globals()[args.command]()
