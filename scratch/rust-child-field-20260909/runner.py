"""Build and compare full runners in an isolated workspace; keep live output intact."""
from pathlib import Path
import argparse
import hashlib
import json
import re
import shutil
import statistics
import subprocess
import probe

HERE, BUILD = probe.HERE, probe.BUILD
GENERATED = probe.SOURCE.parents[2]
FULL = BUILD/'full-output'
EXPECTED = ['7','55','202950','100000','20000','125','100000','21536','22','10000000','1200','1000000','202950','5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+\(Execution time - best of 10\)\s+([0-9.]+) μs')
NAMES = ['before','red_rebuild','red_field']

def patch_paths(text, keep):
    return re.sub(r'path = "(?:\.\./)?([A-Za-z_0-9]+)"',
        lambda m:m[0] if m[1] in keep else f'path = "{GENERATED/m[1]}"',text)

def build():
    FULL.mkdir(exist_ok=True)
    manifest = (GENERATED/'Cargo.toml').read_text()
    manifest = re.sub(r'(?s)^\[workspace\].*?(?=\[package\])','[workspace]\nmembers = ["Purs_App", "Purs_Test_RBTree"]\n\n',manifest)
    (FULL/'Cargo.toml').write_text(patch_paths(manifest,{'Purs_App'}))
    shutil.copy2(GENERATED/'Cargo.lock',FULL/'Cargo.lock')
    shutil.copytree(GENERATED/'src',FULL/'src',dirs_exist_ok=True)
    for name in ['Purs_App','Purs_Test_RBTree']:
        shutil.copytree(GENERATED/name,FULL/name,dirs_exist_ok=True)
        path = FULL/name/'Cargo.toml'
        path.write_text(patch_paths(path.read_text(), {'Purs_Test_RBTree'} if name=='Purs_App' else set()))
    # APFS copy-on-write copy of the cache; Cargo may update only this copy.
    if not (FULL/'target').exists():
        subprocess.run(['cp','-cR',str(GENERATED/'target'),str(FULL/'target')],check=True)
    original = (BUILD/'RBTree-original.rs').read_text()
    signature = 'pub fn Test_RBTree_ins(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {'
    versions = {'before':original,
        'red_rebuild':original.replace(signature,signature+probe.GUARD.replace('BODY',probe.REBUILD)),
        'red_field':original.replace(signature,signature+probe.GUARD.replace('BODY',probe.FIELD))}
    for name, code in versions.items():
        (FULL/'Purs_Test_RBTree/src/lib.rs').write_text(code)
        print('Building isolated full runner:',name,flush=True)
        with (BUILD/f'cargo-{name}.log').open('w') as log:
            subprocess.run(['cargo','build','--release','--offline'],cwd=FULL,stdout=log,stderr=subprocess.STDOUT,check=True)
        binary=BUILD/f'runner-{name}'
        shutil.copy2(FULL/'target/release/purust_output',binary)
        output=subprocess.check_output([str(binary)],text=True)
        parse(output)
        (BUILD/f'runner-check-{name}.log').write_text(output)
        print(name, '14 results verified',flush=True)
    assert probe.SOURCE.read_text()==original, 'Live generated source changed during experiment'

def parse(output):
    matches=PATTERN.findall(output)
    assert [v for _,v,_ in matches]==EXPECTED,matches
    return {name:float(us) for name,_,us in matches}

def measure(selected=NAMES, output_file='runner-results.json'):
    runs={n:[] for n in selected}
    for pair in range(5):
        if len(selected)==2:
            order=selected if pair%2==0 else selected[::-1]
        else:
            order=selected[pair%len(selected):]+selected[:pair%len(selected)]
            if pair%2: order=order[::-1]
        for name in order:
            output=subprocess.check_output([str(BUILD/f'runner-{name}')],text=True)
            (BUILD/f'{Path(output_file).stem}-{name}-{pair+1}.log').write_text(output)
            values=parse(output)
            runs[name].append(values)
            print(pair+1,name,round(sum(values.values())/1000,3),'ms; 14 results verified',flush=True)
    labels=list(runs[selected[0]][0])
    medians={n:{k:statistics.median(r[k] for r in rows) for k in labels} for n,rows in runs.items()}
    data={'method':f'5 blocks of {len(selected)} full runner processes; pairs alternate, triples rotate/reverse; unchanged warmup and best of 10 per benchmark; all 14 results checked each run; O1/mimalloc; no concurrent compilation/instrumentation; isolated copies of App and RBTree, all other crate sources unchanged.',
        'runs_us':runs,'median_us':medians,'total_us':{n:sum(v.values()) for n,v in medians.items()},
        'binary_sha256':{n:hashlib.sha256((BUILD/f'runner-{n}').read_bytes()).hexdigest() for n in selected}}
    (HERE/output_file).write_text(json.dumps(data,indent=2)+'\n')
    print(json.dumps({'median_us':medians,'total_us':data['total_us']},indent=2),flush=True)

if __name__=='__main__':
    parser=argparse.ArgumentParser()
    parser.add_argument('mode',choices=['build','time'])
    parser.add_argument('--variants',nargs='+',choices=NAMES,default=NAMES)
    parser.add_argument('--output',default='runner-results.json')
    args=parser.parse_args()
    if args.mode=='build': build()
    else: measure(args.variants,args.output)
