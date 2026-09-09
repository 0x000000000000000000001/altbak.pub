from pathlib import Path
import ast,json,subprocess,statistics
HERE=Path(__file__).resolve().parent
ROOT=HERE.parents[1]
BUILD=HERE/'build'
DEPS=ROOT/'run/bak/rust/output/purust_output/target/release/deps'
nodes=ast.parse((ROOT/'scratch/rust-records-audit-20260909/probe.py').read_text()).body
h={n.targets[0].id:ast.literal_eval(n.value) for n in nodes if isinstance(n,ast.Assign) and isinstance(n.targets[0],ast.Name) and n.targets[0].id in ['HEADER','HELPERS','COUNT','CHECK','TIME']}
original=(BUILD/'Records-before.rs').read_text()
original=original[original.index('pub fn Test_Records_updateRec('):original.index('pub fn Test_Records_describe(')]
start=original.index('    let _record_child_update_1 = ')
end=original.index('    let mut _base = purs_local_1;',start)
leaf=original[start:end]
e=leaf[leaf.index('    _base.set_e(')+16:leaf.index('\n    _base.set_f(')]
f=leaf[leaf.index('    _base.set_f(')+16:leaf.rindex('\n    _base\n};')]
assert e.endswith(');') and f.endswith(');')
changed=original[:start]+f'    let _record_leaf_update_0 = {e[:-2]};\n    let _record_leaf_update_1 = {f[:-2]};\n'+original[end:]
changed=changed.replace('    _record_child.set_d(_record_child_update_1);','''    let mut _record_leaf = _record_child.get_d();
    _record_child.set_d(crate::Value::Unit);
    _record_leaf.set_e(_record_leaf_update_0);
    _record_leaf.set_f(_record_leaf_update_1);
    _record_child.set_d(_record_leaf);''')
(BUILD/'Records-prototype.rs').write_text(changed)
libs=[]
for name in ['purust_core','perceus_ptr','mimalloc']:
 matches=list(DEPS.glob(f'lib{name}-*.rlib'));assert len(matches)==1,(name,matches)
 libs+=['--extern',f'{name}={matches[0]}']
# Check independent root/middle/leaf sharing plus combinations.
h['CHECK']=h['CHECK'].replace('println!("24 seed/count combinations:', '''for mask in 0..8 {
 for seed in [0_i64,7] {
  let r=Test_Records_updateRec(seed,Test_Records_initial());
  let old=values(&r);
  let root=if mask&1!=0 {Some(r.clone())} else {None};
  let middle=if mask&2!=0 {Some(r.get_b())} else {None};
  let leaf=if mask&4!=0 {Some(r.get_b().get_d())} else {None};
  let result=Test_Records_updateRec(10000,r);
  assert_eq!(values(&result),expected(10000,old));
  if let Some(r)=root {assert_eq!(values(&r),old);}
  if let Some(b)=middle {assert_eq!(b.get_c().unwrap_int(),old[1]); assert_eq!(b.get_d().get_e().unwrap_int(),old[2]); assert_eq!(b.get_d().get_f().unwrap_int(),old[3]);}
  if let Some(d)=leaf {assert_eq!(d.get_e().unwrap_int(),old[2]); assert_eq!(d.get_f().unwrap_int(),old[3]);}
 }
}
 println!("8 independent sharing masks and 24 seed/count combinations:''')
variants={'before':original,'prototype':changed}
for mode in ['COUNT','CHECK','TIME']:
 for name,source in variants.items():
  file=BUILD/f'{mode.lower()}-{name}.rs';file.write_text(h['HEADER']+source+h['HELPERS']+h[mode])
  subprocess.run(['rustc','--edition=2021','-C','opt-level=1',str(file),'-o',str(file.with_suffix('')),'-L',f'dependency={DEPS}',*libs],check=True)
results={}
for mode in ['count','check']:
 results[mode]={name:subprocess.check_output([str(BUILD/f'{mode}-{name}')],text=True) for name in variants}
 print(mode,results[mode],flush=True)
assert results['count']['before'].splitlines()[-1]=='10000 10003 10003'
assert results['count']['prototype'].splitlines()[-1]=='10000 3 3'
runs={name:[] for name in variants}
for pair in range(5):
 for name in (list(variants) if pair%2==0 else list(reversed(variants))):
  samples=list(map(int,subprocess.check_output([str(BUILD/f'time-{name}')],text=True).split()))
  runs[name]+=samples
  print(pair+1,name,statistics.median(samples)/1e6,flush=True)
results['timing']={'method':'Five alternating pairs; 20 samples after warmup; O1, mimalloc; all four fields observed and destruction timed; compilation completed first.', 'samples_ns':runs,'medians_ms':{n:statistics.median(v)/1e6 for n,v in runs.items()}}
(HERE/'prototype-results.json').write_text(json.dumps(results,indent=2)+'\n')
print(results['timing']['medians_ms'])
