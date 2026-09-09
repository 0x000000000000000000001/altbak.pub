"""Compile a minimal fresh TAST fixture, then check its original/prototype Rust."""
import hashlib
import json
import os
import re
import subprocess
from pathlib import Path
import probe

HERE, BUILD = probe.HERE, probe.BUILD
ROOT = probe.ROOT
PURUST = ROOT.parent/'purust/purust'
PURS = os.environ.get('PURS', str(ROOT/'run/bak/js/node_modules/.bin/purs'))
BUILD.mkdir(exist_ok=True)
def run(command):
    result = subprocess.run(command, cwd=HERE, text=True, capture_output=True)
    with (BUILD/'fixture-build.log').open('a') as log:
        log.write(' '.join(map(str,command))+'\n'+result.stdout+result.stderr)
    if result.returncode:
        raise RuntimeError(f'{command[0]} failed: {result.stdout[-4000:]}\n{result.stderr[-4000:]}')
    return result.stdout

run([PURS,'compile',str(HERE/'PerceusProbe.purs'),'--codegen','corefn','--output',str(BUILD/'fixture-tast')])
tast_path = BUILD/'fixture-tast/PerceusProbe/corefn.json'
tast = json.loads(tast_path.read_text())
assert len(tast['dataDecls']) == 2
run(['node','--stack-size=65536',str(PURUST/'bin/purust.js'),'--source',str(BUILD/'fixture-tast'),
    '--out',str(BUILD/'fixture-rust'),'--main','PerceusProbe'])
generated = (BUILD/'fixture-rust/Purs_PerceusProbe/src/lib.rs').read_text()
starts = list(re.finditer(r'^(?:pub )?fn (PerceusProbe_\w+)\(',generated,re.M))
target = next(i for i,m in enumerate(starts) if m[1]=='PerceusProbe_inspect')
function = generated[starts[target].start():starts[target+1].start() if target+1<len(starts) else len(generated)]
kernel = generated[generated.index('#[derive(Clone'):starts[0].start()]+function
prototype, replacements = probe.borrow_projection_receivers(kernel,'crate::Tree::Branch')
assert replacements > 0, kernel
result = {'tast_sha256':hashlib.sha256(tast_path.read_bytes()).hexdigest(),
    'generated_sha256':hashlib.sha256(generated.encode()).hexdigest(),
    'prototype_replacements':replacements,'variants':{}}
for side, code in [('before',kernel),('prototype',prototype)]:
    (BUILD/f'fixture-{side}-kernel.rs').write_text(code.replace('std::rc::Rc','tracked::Rc'))
    source = BUILD/f'fixture-{side}.rs'
    source.write_text('#![allow(warnings)]\n#[path="../tracked_rc.rs"] mod tracked;\n'
        f'include!("fixture-{side}-kernel.rs");\ninclude!("../fixture_harness.rs");\n')
    binary = BUILD/f'fixture-{side}'
    run(['rustc','--edition=2021','-C','opt-level=1',str(source),'-o',str(binary)])
    data, _ = probe.summarize(run([str(binary)]))
    result['variants'][side]=data
for phase in result['variants']['before']['phases']:
    a=result['variants']['before']['phases'][phase]
    b=result['variants']['prototype']['phases'][phase]
    assert a['new']==b['new'] and a['drop_last']==b['drop_last']
    assert a['clone']>b.get('clone',0)
    assert a['clone']-b.get('clone',0)==a['drop_shared']-b.get('drop_shared',0)
result['checks']='Fresh TAST metadata, all pattern outcomes, branch priority, repeated shared inputs, retained child and root, weak references, balanced handle/payload lifetimes.'
(HERE/'fixture.json').write_text(json.dumps(result,indent=2)+'\n')
print(json.dumps({s:d['phases'] for s,d in result['variants'].items()},indent=2),flush=True)
print(result['checks'])
