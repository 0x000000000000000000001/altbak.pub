"""Time the original generated kernel and the one-line prototype, with O1/mimalloc."""
from pathlib import Path
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
DEPS = HERE.parents[1] / 'run/bak/rust/output/purust_output/target/release/deps'
BUILD = HERE / 'build'
BUILD.mkdir(exist_ok=True)
mimalloc = list(DEPS.glob('libmimalloc-*.rlib'))
assert len(mimalloc) == 1
harness = '''
#[global_allocator] static ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;
fn main() {
    for i in 0..16 {
        let start = std::time::Instant::now();
        let tree = Test_RBTree_buildTree(std::hint::black_box(100000), std::rc::Rc::new(Tree::E));
        assert_eq!(Test_RBTree_depth(tree), 22);
        let ns = start.elapsed().as_nanos();
        if i > 0 { println!("{}", ns); }
    }
}
'''
binaries = {}
for side, filename in [('before', 'RBTree-before.rs'), ('prototype', 'RBTree-prototype.rs')]:
    generated = (HERE / filename).read_text()
    starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', generated, re.M))
    functions = {m[1]: generated[m.start():starts[i+1].start() if i+1 < len(starts) else len(generated)]
                 for i, m in enumerate(starts)}
    names = ['max', 'makeBlack', 'depth', 'balance', 'ins', 'insert', 'buildTree']
    names += [n.removeprefix('Test_RBTree_') for n in functions if '__purust_rebuild_' in n or n.endswith('__purust_reuse')]
    source = BUILD / f'time-{side}.rs'
    source.write_text('#![allow(warnings)]\n' + generated[generated.index('#[derive(Clone'):starts[0].start()]
                      + '\n'.join(functions['Test_RBTree_' + n] for n in names) + harness)
    binaries[side] = BUILD / f'time-{side}'
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', '--extern', f'mimalloc={mimalloc[0]}',
                    '-L', f'dependency={DEPS}', str(source), '-o', str(binaries[side])], check=True)
runs = {side: [] for side in binaries}
for pair in range(3):
    for side in (['before', 'prototype'] if pair % 2 == 0 else ['prototype', 'before']):
        values = list(map(int, subprocess.check_output([str(binaries[side])], text=True).split()))
        assert len(values) == 15
        runs[side].append(values)
        print(pair + 1, side, statistics.median(values) / 1e6, 'ms', flush=True)
result = {'method': '3 alternating process pairs, 15 measurements after one warmup per process; O1 and original mimalloc; generated standalone kernel, no allocation instrumentation',
          'runs_ns': runs, 'median_ms': {side: statistics.median(sum(values, [])) / 1e6 for side, values in runs.items()}}
(HERE / 'prototype-time.json').write_text(json.dumps(result, indent=2) + '\n')
print(result['median_ms'])
