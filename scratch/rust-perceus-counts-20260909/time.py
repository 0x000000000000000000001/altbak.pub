"""Time the actual std::rc::Rc kernels, without the counter wrapper."""
from pathlib import Path
import hashlib
import json
import os
import statistics
import subprocess
import probe

HERE, BUILD = probe.HERE, probe.BUILD
DEPS = probe.ROOT/'run/bak/rust/output/purust_output/target/release/deps'
HARNESS = '''
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
mimalloc, = DEPS.glob('libmimalloc-*.rlib')
generated = probe.SOURCE.read_text()
source = probe.kernel(generated)
prototype, n = probe.borrow_projection_receivers(source)
assert n > 0
binaries = {}
BUILD.mkdir(exist_ok=True)
for side, code in [('before', source), ('prototype', prototype)]:
    path = BUILD/f'time-{side}.rs'
    path.write_text('#![allow(warnings)]\n'+code+HARNESS)
    binary = BUILD/f'time-{side}'
    subprocess.run(['rustc','--edition=2021','-C','opt-level=1',str(path),'-o',str(binary),
        '--extern',f'mimalloc={mimalloc}','-L',f'dependency={DEPS}'],check=True)
    binaries[side] = binary
# Optional diagnostics are also compiled before any measurement starts.
if os.environ.get('EMIT_IR') == '1':
    for side in binaries:
        subprocess.run(['rustc','--edition=2021','-C','opt-level=1','-C','debuginfo=line-tables-only',
            '--emit',f'asm={BUILD}/{side}.s,llvm-ir={BUILD}/{side}.ll',str(BUILD/f'time-{side}.rs'),
            '--extern',f'mimalloc={mimalloc}','-L',f'dependency={DEPS}'],check=True)
runs = {side: [] for side in binaries}
for pair in range(3):
    for side in (['before','prototype'] if pair % 2 == 0 else ['prototype','before']):
        values = list(map(int,subprocess.check_output([str(binaries[side])],text=True).split()))
        assert len(values) == 15
        runs[side].append(values)
        print(pair+1,side,statistics.median(values)/1e6,'ms',flush=True)
result = {'method':'Three alternating process pairs; 15 iterations after one warmup per process; O1/mimalloc; native std::rc::Rc; build, depth and destruction included. Isolated kernel, not full runner.',
    'source_sha256':hashlib.sha256(generated.encode()).hexdigest(),
    'rustc':subprocess.check_output(['rustc','--version'],text=True).strip(),
    'runs_ns':runs, 'median_ms':{s:statistics.median(sum(v,[]))/1e6 for s,v in runs.items()}}
(HERE/'timings.json').write_text(json.dumps(result,indent=2)+'\n')
print(result['median_ms'],flush=True)
