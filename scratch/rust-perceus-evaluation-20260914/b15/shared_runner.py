"""Prepare comparable retained-root workloads; clocks run only with `time`."""
from pathlib import Path
import argparse
import hashlib
import json
import statistics
import subprocess

import probe

HERE = Path(__file__).resolve().parent
BUILD = HERE / 'build/shared'
NAMES = ['before', 'borrowed', 'consuming', 'hybrid']

MAIN = r'''
#[global_allocator]
static ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;

fn shared_workload() -> i64 {
    let n = std::hint::black_box(100000_i64);
    let repetitions = std::hint::black_box(20_i64);
    let tree = Test_RBTree_buildTree(n, std::rc::Rc::new(Tree::E));
    let mut result = 0_i64;
    for _ in 0..repetitions {
        result += std::hint::black_box(Test_RBTree_depth(std::hint::black_box(tree.clone())));
    }
    // All retained ownership and final destruction are inside each sample.
    drop(tree);
    std::hint::black_box(result)
}

fn main() {
    match std::env::args().nth(1).as_deref() {
        Some("smoke") => {
            assert_eq!(shared_workload(), 440);
            println!("100000 nodes, 20 retained-root reads, depth sum 440; final owner released");
        }
        Some("time") => {
            assert_eq!(shared_workload(), 440); // unmeasured process warm-up
            for _ in 0..7 {
                let start = std::time::Instant::now();
                let result = shared_workload();
                let elapsed = start.elapsed().as_nanos();
                assert_eq!(result, 440);
                println!("{elapsed}");
            }
        }
        _ => panic!("Pass smoke (no clocks) or time explicitly"),
    }
}
'''


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def build():
    (BUILD / 'src/bin').mkdir(parents=True, exist_ok=True)
    sources = probe.variants()
    for name in NAMES:
        code = '#![allow(warnings)]\n' + probe.extractor.kernel(sources[name]) + MAIN
        (BUILD / f'src/bin/{name}.rs').write_text(code)
    (BUILD / 'Cargo.toml').write_text('''[workspace]
[package]
name = "b15_retained_root_probe"
version = "0.1.0"
edition = "2021"
[profile.release]
opt-level = 1
debug = true
[dependencies]
mimalloc = "0.1.32"
''')
    subprocess.run(['cargo', 'build', '--release', '--offline'], cwd=BUILD, check=True)
    outputs = {}
    for name in NAMES:
        outputs[name] = subprocess.check_output([str(BUILD / f'target/release/{name}'), 'smoke'], text=True).strip()
        print(name, outputs[name], flush=True)
    metadata = {
        'source_sha256': sha(probe.BUILD / 'RBTree-original.rs'),
        'workload': 'Build 100000 nodes with generated constructor; 20 depth calls each consume clone of retained root; final drop inside sample; black_box size/count/owner/result.',
        'profile': 'release opt-level=1 debug=true mimalloc; identical harness and Cargo package across four exact kernels',
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'cargo_lock_sha256': sha(BUILD / 'Cargo.lock'),
        'variants': {name: {
            'source_sha256': sha(BUILD / f'src/bin/{name}.rs'),
            'binary_sha256': sha(BUILD / f'target/release/{name}'),
            'smoke': outputs[name],
        } for name in NAMES},
        'timings_run': False,
    }
    (HERE / 'shared-build.json').write_text(json.dumps(metadata, indent=2) + '\n')


def measure():
    metadata = json.loads((HERE / 'shared-build.json').read_text())
    for name in NAMES:
        assert sha(BUILD / f'target/release/{name}') == metadata['variants'][name]['binary_sha256']
    rows = {name: [] for name in NAMES}
    order_log = []
    for block in range(8):
        offset = block % len(NAMES)
        order = NAMES[offset:] + NAMES[:offset]
        if block // len(NAMES):
            order = list(reversed(order))
        order_log.append(order)
        for name in order:
            output = subprocess.check_output([str(BUILD / f'target/release/{name}'), 'time'], text=True)
            values = list(map(int, output.split()))
            assert len(values) == 7
            rows[name].append(values)
            print(block + 1, name, statistics.median(values) / 1e6, 'ms', flush=True)
    result = {
        'method': 'Eight blocks: four rotations then four reversed rotations; separate process per variant, one unmeasured warmup then seven samples; construction + 20 retained-root reads + final destruction; no instrumentation.',
        'order': order_log,
        'runs_ns': rows,
        'process_medians_ms': {name: [statistics.median(row) / 1e6 for row in runs] for name, runs in rows.items()},
        'median_ms': {name: statistics.median([statistics.median(row) for row in runs]) / 1e6 for name, runs in rows.items()},
        'build': metadata,
    }
    (HERE / 'shared-timings.json').write_text(json.dumps(result, indent=2) + '\n')
    print(json.dumps(result['median_ms'], indent=2))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['build', 'time'])
    args = parser.parse_args()
    {'build': build, 'time': measure}[args.mode]()
