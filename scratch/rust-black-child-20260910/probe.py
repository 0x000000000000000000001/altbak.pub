"""Isolated experiment: preserve a black parent's fields unless recursion rotates it."""
from pathlib import Path
import argparse
import hashlib
import importlib.util
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE / 'build'
SOURCE = ROOT / 'run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs'
PREVIOUS = HERE.parent / 'rust-perceus-counts-20260909'
spec = importlib.util.spec_from_file_location('previous', PREVIOUS / 'probe.py')
previous = importlib.util.module_from_spec(spec)
spec.loader.exec_module(previous)
SIGNATURE = 'pub fn Test_RBTree_ins(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {'

GUARD = '''
    // Experiment only: the rotation predicate is evaluated AFTER recursion.
    fn probe_rotation(node: &Tree) -> bool {
        fn red(node: &Tree) -> bool { matches!(node, Tree::T(Color::R, ..)) }
        fn red_double(node: &Tree) -> bool {
            match node { Tree::T(Color::R, left, _, right) => red(left) || red(right), _ => false }
        }
        match node { Tree::T(Color::B, left, _, right) => red_double(left) || red_double(right), _ => false }
    }
    let probe_direction = match purs_local_1.as_ref() {
        Tree::T(Color::B, _, key, _) if purs_local_0 < *key => -1,
        Tree::T(Color::B, _, key, _) if purs_local_0 > *key => 1,
        _ => 0,
    };
    if probe_direction != 0 {
        if let Some(probe_slot) = std::rc::Rc::get_mut(&mut purs_local_1) {
            BODY
            if !probe_rotation(probe_slot) {
                return purs_local_1;
            }
            // Keep the existing rotation worker and its sharing fallbacks.
            let Tree::T(color, left, key, right) = probe_slot.__purust_take().unwrap() else { unreachable!() };
            return Test_RBTree_balance__purust_reuse(color, left, key, right, purs_local_1);
        }
    }
'''

REBUILD = '''let Tree::T(color, mut left, key, mut right) = probe_slot.__purust_take().unwrap() else { unreachable!() };
            if probe_direction < 0 { left = Test_RBTree_ins(purs_local_0, left); }
            else { right = Test_RBTree_ins(purs_local_0, right); }
            *probe_slot = Tree::T(color, left, key, right);'''

FIELD = '''let Tree::T(_, left, _, right) = probe_slot else { unreachable!() };
            if probe_direction < 0 {
                let child = std::mem::replace(left, right.clone());
                *left = Test_RBTree_ins(purs_local_0, child);
            } else {
                let child = std::mem::replace(right, left.clone());
                *right = Test_RBTree_ins(purs_local_0, child);
            }'''

def variants():
    original = previous.kernel((BUILD / 'RBTree-original.rs').read_text())
    assert original.count(SIGNATURE) == 1
    return {'before': original,
        'black_rebuild': original.replace(SIGNATURE, SIGNATURE + GUARD.replace('BODY', REBUILD)),
        'black_field': original.replace(SIGNATURE, SIGNATURE + GUARD.replace('BODY', FIELD))}

HARNESS = '''
#[global_allocator] static ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;
fn main() {
    for i in 0..16 {
        let start = std::time::Instant::now();
        let tree = Test_RBTree_buildTree(std::hint::black_box(100000), std::rc::Rc::new(Tree::E));
        assert_eq!(Test_RBTree_depth(tree), 22);
        let elapsed = start.elapsed().as_nanos();
        if i > 0 { println!("{}", elapsed); }
    }
}
'''

def prepare():
    BUILD.mkdir(exist_ok=True)
    if not (BUILD / 'RBTree-original.rs').exists():
        (BUILD / 'RBTree-original.rs').write_bytes(SOURCE.read_bytes())
    deps = SOURCE.parents[2] / 'target/release/deps'
    mimalloc, = deps.glob('libmimalloc-*.rlib')
    for name, code in variants().items():
        path = BUILD / f'time-{name}.rs'
        path.write_text('#![allow(warnings)]\n' + code + HARNESS)
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', '--extern', f'mimalloc={mimalloc}',
            '-L', f'dependency={deps}', str(path), '-o', str(BUILD / f'time-{name}')], check=True)
        print(name, 'compiled', flush=True)
    metadata = {'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'source': str(SOURCE), 'source_sha256': hashlib.sha256(SOURCE.read_bytes()).hexdigest(),
        'profile': 'O1/mimalloc; source Rust patched only in isolated copies',
        'revisions': {str(p): subprocess.check_output(['git', '-C', str(p), 'rev-parse', 'HEAD'], text=True).strip()
            for p in [ROOT, ROOT.parent / 'purust/purust']}}
    (HERE / 'metadata.json').write_text(json.dumps(metadata, indent=2) + '\n')

def validate():
    base = (PREVIOUS / 'count_harness.rs').read_text()
    base = base[base.index('fn validate('):].replace('tracked::Rc', 'std::rc::Rc')
    base = re.sub(r'    tracked::(?:phase\("[^"]*"\)|print_events\(\));\n', '', base)
    base = base.replace('fn main() {', 'fn original_checks() {')
    extra = (HERE.parent / 'rust-child-field-20260909/validation-extra.rs').read_text()
    black = extra[extra.index('fn b13_red_parent_cases()'):extra.index('pub fn extra_checks()')]
    black = black.replace('b13_red_parent_cases', 'probe_black_parent_cases')
    black = black.replace('Tree::T(Color::R, left, 10, right)', 'Tree::T(Color::B, left, 10, right)')
    black = black.replace('Tree::T(Color::R, _, 10, _)', 'Tree::T(Color::B, _, 10, _)')
    black = black.replace('B13 red-parent cases', 'Probe black-parent cases')
    results = {}
    for name, code in variants().items():
        path = BUILD / f'validate-{name}.rs'
        path.write_text('#![allow(warnings)]\n' + code + base + extra + black
            + '\nfn main() { original_checks(); extra_checks(); probe_black_parent_cases(); }\n')
        binary = BUILD / f'validate-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        results[name] = subprocess.check_output([str(binary)], text=True)
        print(name, results[name], flush=True)
    (HERE / 'validation.json').write_text(json.dumps(results, indent=2) + '\n')

def measure():
    runs = {n: [] for n in variants()}
    for block in range(5):
        names = list(runs)
        order = names[block % 3:] + names[:block % 3]
        if block % 2: order.reverse()
        for name in order:
            values = list(map(int, subprocess.check_output([str(BUILD / f'time-{name}')], text=True).split()))
            assert len(values) == 15
            runs[name].append(values)
            print(block + 1, name, statistics.median(values) / 1e6, 'ms', flush=True)
    result = {'method': 'Five rotating/reversed blocks; one warmup + 15 samples per process; native Rc; construction, depth and destruction included; O1/mimalloc; no concurrent build or instrumentation.',
        'runs_ns': runs, 'median_ms': {n: statistics.median(sum(v, [])) / 1e6 for n, v in runs.items()}}
    (HERE / 'timings.json').write_text(json.dumps(result, indent=2) + '\n')
    print(result['median_ms'])

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['prepare', 'validate', 'time'])
    args = parser.parse_args()
    {'prepare': prepare, 'validate': validate, 'time': measure}[args.mode]()
