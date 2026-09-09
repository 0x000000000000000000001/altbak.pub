from pathlib import Path
import ast,json,subprocess,statistics,shutil
HERE=Path(__file__).resolve().parent
ROOT=HERE.parents[1]
BUILD=HERE/'build'
DEPS=ROOT/'run/bak/rust/output/purust_output/target/release/deps'
nodes=ast.parse((ROOT/'scratch/rust-records-audit-20260909/probe.py').read_text()).body
h={n.targets[0].id:ast.literal_eval(n.value) for n in nodes if isinstance(n,ast.Assign) and isinstance(n.targets[0],ast.Name) and n.targets[0].id in ['HEADER','HELPERS','COUNT','CHECK','TIME']}
original=(BUILD/'Records-before.rs').read_text()
original=original[original.index('pub fn Test_Records_updateRec('):original.index('pub fn Test_Records_describe(')]
start=original.index('    let _record_update_1 = ')
end=original.index('    let mut _base = purs_local_1;',start)
child=original[start:end]
c=child[child.index('    _base.set_c(')+16:child.index('\n    _base.set_d(')]
assert c.endswith(');'),repr(c)
c=c[:-2]
d=child[child.index('    _base.set_d(')+16:child.rindex('\n    _base\n};')]
assert d.endswith(');'),repr(d)
d=d[:-2]
# All RHS expressions remain byte-identical to the baseline output.
changed=original[:start]+f'    let _record_child_update_0 = {c};\n    let _record_child_update_1 = {d};\n'+original[end:]
changed=changed.replace('    _base.set_b(_record_update_1);','''    let mut _record_child = _base.get_b();
    _base.set_b(crate::Value::Unit);
    _record_child.set_c(_record_child_update_0);
    _record_child.set_d(_record_child_update_1);
    _base.set_b(_record_child);''')
(BUILD/'Records-prototype.rs').write_text(changed)
libs=[]
for name in ['purust_core','perceus_ptr','mimalloc']:
    matches=list(DEPS.glob(f'lib{name}-*.rlib'));assert len(matches)==1,(name,matches)
    libs+=['--extern',f'{name}={matches[0]}']
variants={'before':original,'prototype':changed}
for mode in ['COUNT','CHECK','TIME']:
    for name,source in variants.items():
        file=BUILD/f'{mode.lower()}-{name}.rs'
        file.write_text(h['HEADER']+source+h['HELPERS']+h[mode])
        subprocess.run(['rustc','--edition=2021','-C','opt-level=1',str(file),'-o',str(file.with_suffix('')),'-L',f'dependency={DEPS}',*libs],check=True)
results={}
for mode in ['count','check']:
    results[mode]={name:subprocess.check_output([str(BUILD/f'{mode}-{name}')],text=True) for name in variants}
    print(mode,results[mode],flush=True)
assert results['count']['before'].splitlines()[-1]=='10000 20003 20003'
assert results['count']['prototype'].splitlines()[-1]=='10000 10003 10003'
runs={name:[] for name in variants}
for pair in range(5):
    for name in (list(variants) if pair%2==0 else list(reversed(variants))):
        samples=list(map(int,subprocess.check_output([str(BUILD/f'time-{name}')],text=True).split()))
        runs[name]+=samples
        print(pair+1,name,statistics.median(samples)/1e6,flush=True)
results['timing']={'method':'Five alternating pairs; 20 samples after warmup; O1, mimalloc; all four fields observed and destruction timed; all compilation completed first.', 'samples_ns':runs,'medians_ms':{n:statistics.median(v)/1e6 for n,v in runs.items()}}
(HERE/'prototype-results.json').write_text(json.dumps(results,indent=2)+'\n')
print(results['timing']['medians_ms'])
