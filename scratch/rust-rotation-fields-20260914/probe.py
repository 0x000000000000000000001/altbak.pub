"""Isolated LL rotation experiment; production sources are read-only inputs."""
from pathlib import Path
import argparse
import hashlib
import importlib.util
import json
import re
import statistics
import subprocess
import sys

sys.dont_write_bytecode = True
HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE / 'build'
GENERATED = BUILD / 'generated'
SOURCE = GENERATED / 'Purs_Test_RBTree/src/lib.rs'
PREVIOUS = ROOT / 'scratch/rust-perceus-counts-20260909'
NAMES = ['before', 'rebuild', 'fields']
TAIL = 'else { let std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) = _purust_post_slot.__purust_take()'


def digest(data):
    return hashlib.sha256(data).hexdigest()


def kernel(source):
    starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', source, re.M))
    ranges = [(source.index('#[derive(Clone'), starts[0].start())]
    names = {'max', 'makeBlack', 'depth', 'balance', 'ins', 'insert', 'buildTree'}
    for index, match in enumerate(starts):
        name = match[1].removeprefix('Test_RBTree_')
        if name in names or '__purust_' in name:
            ranges.append((match.start(), starts[index + 1].start() if index + 1 < len(starts) else len(source)))
    return '\n'.join(source[a:b] for a, b in ranges)


def variants(full=False):
    original = (BUILD / 'RBTree-original.rs').read_text()
    if not full:
        original = kernel(original)
    assert original.count(TAIL) == 2
    result = {'before': original}
    for name, filename in [('rebuild', 'rebuild.rs'), ('fields', 'rotation.rs')]:
        patched = original.replace(TAIL,
            'else if probe_rotate_ll(_purust_post_slot) { purs_local_1 } ' + TAIL)
        result[name] = patched + '\n' + (HERE / filename).read_text()
    return result


TIMING = '''
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
    snapshot = BUILD / 'RBTree-original.rs'
    if snapshot.exists():
        assert snapshot.read_bytes() == SOURCE.read_bytes(), 'Generated input changed'
    else:
        snapshot.write_bytes(SOURCE.read_bytes())
    (BUILD / 'src/bin').mkdir(parents=True, exist_ok=True)
    for name, code in variants().items():
        (BUILD / f'{name}.rs').write_text(code)
        (BUILD / f'src/bin/{name}.rs').write_text('#![allow(warnings)]\n' + code + TIMING)
    (BUILD / 'Cargo.toml').write_text('''[workspace]
[package]
name = "rotation_fields_probe"
version = "0.1.0"
edition = "2021"
[profile.release]
opt-level = 1
debug = true
[dependencies]
mimalloc = "0.1.32"
''')
    subprocess.run(['cargo', 'build', '--release', '--offline'], cwd=BUILD, check=True)
    metadata = {
        'source': str(SOURCE), 'source_sha256': digest(snapshot.read_bytes()),
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'profile': 'O1/debug=true/mimalloc; isolated generated kernel; construction, depth and destruction included',
        'baseline_readme': {'RBTree_ms': 11.631, 'suite_ms': 12.56, 'native_RBTree_ms': 36.070, 'native_suite_ms': 36.13},
        'variants': {name: {'source_sha256': digest(code.encode()),
            'binary_sha256': digest((BUILD / f'target/release/{name}').read_bytes())}
            for name, code in variants().items()},
        'backend_revision': subprocess.check_output(['git', '-C', str(ROOT.parent / 'purust/purust'), 'rev-parse', 'HEAD'], text=True).strip(),
    }
    (HERE / 'metadata.json').write_text(json.dumps(metadata, indent=2) + '\n')
    print('Three isolated timing binaries built; no timings yet.', flush=True)


def validate():
    base = (PREVIOUS / 'count_harness.rs').read_text()
    base = base[base.index('fn validate('):].replace('tracked::Rc', 'std::rc::Rc')
    base = re.sub(r'    tracked::(?:phase\("[^"]*"\)|print_events\(\));\n', '', base)
    base = base.replace('fn main() {', 'fn original_checks() {')
    extra = (ROOT / 'scratch/rust-child-field-20260909/validation-extra.rs').read_text()
    results = {}
    for name, code in variants().items():
        ll = '' if name == 'before' else (HERE / 'validation-extra.rs').read_text()
        call = '' if name == 'before' else 'validation_extra();'
        source = '#![allow(warnings)]\n' + code + base + extra + ll
        source += '\nfn main() { original_checks(); extra_checks(); ' + call + ' println!("all checks passed"); }\n'
        path = BUILD / f'validate-{name}.rs'
        path.write_text(source)
        binary = BUILD / f'validate-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        results[name] = subprocess.check_output([str(binary)], text=True)
        print(name, results[name], flush=True)
    (HERE / 'validation.json').write_text(json.dumps(results, indent=2) + '\n')


COUNTERS = '''
thread_local! { static PROBE_COUNTS: std::cell::Cell<[u64;4]> = const { std::cell::Cell::new([0;4]) }; }
fn probe_event(i: usize) { PROBE_COUNTS.with(|s| { let mut a=s.get(); a[i]+=1; s.set(a); }); }
fn probe_reset() { PROBE_COUNTS.with(|s|s.set([0;4])); }
fn probe_print() { PROBE_COUNTS.with(|s| { print!("PROBE"); for n in s.get() { print!(" {n}"); } println!(); }); }
'''


def count():
    spec = importlib.util.spec_from_file_location('previous', PREVIOUS / 'probe.py')
    previous = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(previous)
    harness = (PREVIOUS / 'count_harness.rs').read_text()
    harness = harness.replace('tracked::phase("unique_build");', 'tracked::phase("unique_build"); probe_reset();')
    harness = harness.replace('tracked::phase("unique_depth_and_drop");', 'probe_print(); tracked::phase("unique_depth_and_drop");')
    results = {}
    for name, code in variants().items():
        take = 'pub fn __purust_take(&mut self) -> std::option::Option<Self> {'
        rebuild = 'let payload = crate::Tree::T(a0, a1, a2, a3);'
        assert code.count(take) == code.count(rebuild) == 1
        code = code.replace(take, take + ' probe_event(0);')
        code = code.replace(rebuild, 'probe_event(1); ' + rebuild)
        code = code.replace('// PROBE_SUCCESS', 'probe_event(2);')
        code = code.replace('// PROBE_REBUILD', 'probe_event(3);')
        code = code.replace('std::rc::Rc', 'tracked::Rc')
        path = BUILD / f'count-{name}.rs'
        path.write_text('#![allow(warnings)]\n#[path="' + str(PREVIOUS / 'tracked_rc.rs') + '"] mod tracked;\n' + code + COUNTERS + harness)
        binary = BUILD / f'count-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        output = subprocess.check_output([str(binary)], text=True)
        (BUILD / f'events-{name}.tsv').write_text(output)
        line, = [line for line in output.splitlines() if line.startswith('PROBE ')]
        summary, _ = previous.summarize('\n'.join(line for line in output.splitlines() if not line.startswith('PROBE ')))
        results[name] = {'unique_build': dict(zip(['take', 'rebuild_helper', 'prototype_rotations', 'rebuild_direct'], map(int, line.split()[1:]))), 'phases': summary['phases']}
        print(name, json.dumps(results[name]['unique_build']), flush=True)
        assert summary['phases']['unique_build']['new'] == 100001
        assert results[name]['unique_build']['prototype_rotations'] == (0 if name == 'before' else 99978)
        assert results[name]['unique_build']['take'] == (0 if name == 'fields' else 299934)
    (HERE / 'counts.json').write_text(json.dumps(results, indent=2) + '\n')


def measure():
    runs = {name: [] for name in NAMES}
    order_log = []
    for block in range(5):
        order = NAMES[block % 3:] + NAMES[:block % 3]
        if block % 2: order = order[::-1]
        order_log.append(order)
        for name in order:
            values = list(map(int, subprocess.check_output([str(BUILD / f'target/release/{name}')], text=True).split()))
            assert len(values) == 15
            runs[name].append(values)
            print(block + 1, name, statistics.median(values) / 1e6, 'ms', flush=True)
    medians = {name: statistics.median([statistics.median(values) for values in rows]) / 1e6 for name, rows in runs.items()}
    result = {'method': 'Five rotating/reversed blocks; one warmup plus 15 samples per process; median of process medians; native Rc; O1/mimalloc; construction, depth, destruction included; no simultaneous compilation or instrumentation by this task',
        'order': order_log, 'runs_ns': runs, 'median_ms': medians}
    (HERE / 'timings.json').write_text(json.dumps(result, indent=2) + '\n')
    print(json.dumps(medians, indent=2))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['prepare', 'validate', 'count', 'time'])
    args = parser.parse_args()
    {'prepare': prepare, 'validate': validate, 'count': count, 'time': measure}[args.mode]()
