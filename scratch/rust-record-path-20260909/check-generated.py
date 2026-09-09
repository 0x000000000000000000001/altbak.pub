"""Count allocations and check persistence on the actual before/after output."""
from pathlib import Path
import ast,json,subprocess
HERE=Path(__file__).resolve().parent
ROOT=HERE.parents[1]
GENERATED=ROOT/'run/bak/rust/output/purust_output'
DEPS=GENERATED/'target/release/deps'
# Reuse the audited counting and persistence harnesses, without executing the probe.
nodes=ast.parse((ROOT/'scratch/rust-records-audit-20260909/probe.py').read_text()).body
harnesses={node.targets[0].id:ast.literal_eval(node.value) for node in nodes
           if isinstance(node,ast.Assign) and isinstance(node.targets[0],ast.Name)
           and node.targets[0].id in ['HEADER','HELPERS','COUNT','CHECK']}
harnesses['CHECK']=(HERE/'sharing-check.rs').read_text()
paths={'before':HERE/'build/Records-before.rs',
       'after':GENERATED/'Purs_Test_Records/src/lib.rs'}
libraries=[]
for name in ['purust_core','perceus_ptr','mimalloc']:
    matches=list(DEPS.glob(f'lib{name}-*.rlib'))
    assert len(matches)==1,(name,matches)
    libraries += ['--extern',f'{name}={matches[0]}']
results={}
for mode in ['COUNT','CHECK']:
    results[mode.lower()]={}
    for side,path in paths.items():
        source=path.read_text()
        source=source[source.index('pub fn Test_Records_updateRec('):source.index('pub fn Test_Records_describe(')]
        out=HERE/'build'/f'{mode.lower()}-{side}.rs'
        binary=out.with_suffix('')
        out.write_text(harnesses['HEADER']+source+harnesses['HELPERS']+harnesses[mode])
        subprocess.run(['rustc','--edition=2021','-C','opt-level=1',str(out),'-o',str(binary),
                        '-L',f'dependency={DEPS}',*libraries],check=True)
        results[mode.lower()][side]=subprocess.check_output([str(binary)],text=True)
assert results['count']['before'].splitlines()[-1]=='10000 10003 10003'
assert results['count']['after'].splitlines()[-1]=='10000 3 3'
(HERE/'generated-checks.json').write_text(json.dumps(results,indent=2)+'\n')
print(json.dumps(results,indent=2),flush=True)
