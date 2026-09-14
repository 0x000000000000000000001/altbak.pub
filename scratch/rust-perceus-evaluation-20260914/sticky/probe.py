#!/usr/bin/env python3
"""Snapshot real generated Records/runtime; force saturation only in scratch.

build/check/count never time workloads. `time --timing-authorized` is separate.
"""
from pathlib import Path
import argparse
import hashlib
import json
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[2]
GENERATED = ROOT / 'run/bak/rust/output/purust_output'
RUNTIME = ROOT.parent / 'purust/purust/tests/runtime/perceus_ptr/src'
BUILD = HERE / 'build'

METRICS = r'''
#[cfg(sticky_counts)]
pub mod sticky_metrics {
    use std::sync::atomic::{AtomicU64, Ordering::Relaxed};
    pub static ALLOCS: AtomicU64 = AtomicU64::new(0);
    pub static FREES: AtomicU64 = AtomicU64::new(0);
    pub static ALLOC_BYTES: AtomicU64 = AtomicU64::new(0);
    pub static FREE_BYTES: AtomicU64 = AtomicU64::new(0);
    pub static CLONES: AtomicU64 = AtomicU64::new(0);
    pub static DROPS: AtomicU64 = AtomicU64::new(0);
    pub static WRITES: AtomicU64 = AtomicU64::new(0);
    pub static STICKY_CLONES: AtomicU64 = AtomicU64::new(0);
    pub static STICKY_DROPS: AtomicU64 = AtomicU64::new(0);
    pub static COW: AtomicU64 = AtomicU64::new(0);
    pub static MAX_NORMAL_OWNERS: AtomicU64 = AtomicU64::new(1);
    pub fn hit(counter: &AtomicU64) { counter.fetch_add(1, Relaxed); }
    pub fn bytes(counter: &AtomicU64, n: usize) { counter.fetch_add(n as u64, Relaxed); }
    pub fn owners(n: u32) { MAX_NORMAL_OWNERS.fetch_max(n as u64, Relaxed); }
    pub fn snapshot() -> [u64; 11] {
        [&ALLOCS, &FREES, &ALLOC_BYTES, &FREE_BYTES, &CLONES, &DROPS,
            &WRITES, &STICKY_CLONES, &STICKY_DROPS, &COW, &MAX_NORMAL_OWNERS]
            .map(|counter| counter.load(Relaxed))
    }
}
'''

def replace_once(text, old, new):
    assert text.count(old) == 1, old
    return text.replace(old, new)

def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()

def run(command, **kwargs):
    result = subprocess.run(list(map(str, command)), text=True, capture_output=True, **kwargs)
    if result.returncode:
        raise RuntimeError(f'{command}\n{result.stdout}\n{result.stderr}')
    return result.stdout

def prepare():
    BUILD.mkdir(exist_ok=True)
    original = (RUNTIME / 'local.rs').read_text()
    local = replace_once(original,
        '#[cfg(test)]\n    pub unsafe fn set_count',
        '#[cfg(any(test, sticky_probe))]\n    pub unsafe fn set_count')
    local = replace_once(local, 'std::ptr::write(ptr, PerceusBox { count: 1, data });',
        '''std::ptr::write(ptr, PerceusBox { count: 1, data });
            #[cfg(sticky_counts)] {
                sticky_metrics::hit(&sticky_metrics::ALLOCS);
                sticky_metrics::bytes(&sticky_metrics::ALLOC_BYTES, layout.size());
            }''')
    local = replace_once(local, 'if !this.is_unique() {',
        '''if !this.is_unique() {
            #[cfg(sticky_counts)] sticky_metrics::hit(&sticky_metrics::COW);''')
    local = replace_once(local, 'fn clone(&self) -> Self {',
        '''fn clone(&self) -> Self {
        #[cfg(sticky_counts)] sticky_metrics::hit(&sticky_metrics::CLONES);''')
    local = replace_once(local, '(*b).count = count.saturating_add(1);',
        '''(*b).count = count.saturating_add(1);
                #[cfg(sticky_counts)] {
                    sticky_metrics::hit(&sticky_metrics::WRITES);
                    sticky_metrics::owners((*b).count);
                }''')
    local = replace_once(local, 'PerceusPtr { ptr: self.ptr }',
        '''#[cfg(sticky_counts)] if self.count() == Self::STICKY_COUNT {
            sticky_metrics::hit(&sticky_metrics::STICKY_CLONES);
        }
        PerceusPtr { ptr: self.ptr }''')
    local = replace_once(local, 'impl<T> Drop for PerceusPtr<T> {\n    fn drop(&mut self) {',
        '''impl<T> Drop for PerceusPtr<T> {
    fn drop(&mut self) {
        #[cfg(sticky_counts)] {
            sticky_metrics::hit(&sticky_metrics::DROPS);
            if self.count() == Self::STICKY_COUNT { sticky_metrics::hit(&sticky_metrics::STICKY_DROPS); }
        }''')
    local = replace_once(local, '(*b).count = new_count;',
        '''(*b).count = new_count;
                #[cfg(sticky_counts)] sticky_metrics::hit(&sticky_metrics::WRITES);''')
    local = replace_once(local, 'dealloc(b as *mut u8, layout);',
        '''#[cfg(sticky_counts)] {
                        sticky_metrics::hit(&sticky_metrics::FREES);
                        sticky_metrics::bytes(&sticky_metrics::FREE_BYTES, layout.size());
                    }
                    dealloc(b as *mut u8, layout);''')
    (BUILD / 'local.rs').write_text(local + METRICS)
    (BUILD / 'lib.rs').write_bytes((RUNTIME / 'lib.rs').read_bytes())
    (BUILD / 'threaded.rs').write_bytes((RUNTIME / 'threaded.rs').read_bytes())
    core = GENERATED / 'purust_core/src/lib.rs'
    record = GENERATED / 'Purs_Test_Records/src/lib.rs'
    source = record.read_text()
    kernel = source[source.index('pub fn Test_Records_updateRec('):source.index('pub fn Test_Records_initial(')]
    (BUILD / 'records.rs').write_text(kernel)
    (BUILD / 'purust_core.rs').write_bytes(core.read_bytes())
    sources = [RUNTIME / name for name in ['lib.rs', 'local.rs', 'threaded.rs']] + [core, record]
    (HERE / 'metadata.json').write_text(json.dumps({
        'inputs': {str(p): sha(p) for p in sources},
        'generated_records_kernel_sha256': sha(BUILD / 'records.rs'),
        'instrumented_runtime_sha256': sha(BUILD / 'local.rs'),
        'scope': 'Exact generated update kernel and purust_core; only private runtime copy exposes test count setter and optional counters.',
        'artificial_saturation': 'Set initial three PerceusPtr cells directly to u32::MAX. Never lower threshold; never claim natural saturation.',
        'threaded': 'Source inspected only: Arc delegates clone/drop/make_mut; STICKY_COUNT is a constant, not a sticky execution branch.',
        'rustc': run(['rustc', '--version']).strip(),
        'profile': {'opt_level': 1, 'overflow_checks': True, 'allocator': 'Rust default System; full benchmark runner uses mimalloc'},
    }, indent=2) + '\n')

def build():
    prepare()
    deps = GENERATED / 'target/release/deps'
    regex = list(deps.glob('libfancy_regex-*.rlib'))
    assert len(regex) == 1, regex
    for variant in ['counts', 'plain']:
        directory = BUILD / variant
        directory.mkdir(exist_ok=True)
        flags = ['--edition=2021', '-C', 'opt-level=1', '-C', 'overflow-checks=yes', '--cfg', 'sticky_probe']
        if variant == 'counts': flags += ['--cfg', 'sticky_counts']
        runtime = directory / 'libperceus_ptr.rlib'
        core = directory / 'libpurust_core.rlib'
        run(['rustc', *flags, '--crate-type=rlib', '--crate-name=perceus_ptr', BUILD / 'lib.rs', '-o', runtime])
        run(['rustc', *flags, '--crate-type=rlib', '--crate-name=purust_core', BUILD / 'purust_core.rs',
            '-o', core, '-L', f'dependency={directory}', '-L', f'dependency={deps}',
            '--extern', f'perceus_ptr={runtime}', '--extern', f'fancy_regex={regex[0]}'])
        run(['rustc', *flags, HERE / 'harness.rs', '-o', directory / 'probe',
            '-L', f'dependency={directory}', '-L', f'dependency={deps}',
            '--extern', f'perceus_ptr={runtime}', '--extern', f'purust_core={core}'])
        print(f'Compiled {variant}; no timing run.', flush=True)
    metadata = json.loads((HERE / 'metadata.json').read_text())
    metadata['frozen_sources'] = {str(path.relative_to(HERE)): sha(path) for path in
        [HERE / 'harness.rs', HERE / 'probe.py', BUILD / 'records.rs', BUILD / 'purust_core.rs',
         BUILD / 'lib.rs', BUILD / 'local.rs', BUILD / 'threaded.rs']}
    metadata['binaries'] = {variant: {'path': str(BUILD / variant / 'probe'),
        'sha256': sha(BUILD / variant / 'probe')} for variant in ['counts', 'plain']}
    metadata['timing_protocol'] = {
        'warmup': '10000 iterations, same scenario and saturation mode, outside the timer',
        'setup': 'Fresh input construction, saturation and closure capture outside timer',
        'segment': 'Workload loop, checksum/result observation, validation, and destruction of all final owners',
        'forced_retention': 'Three cells deliberately retained per forced invocation; warmup retains its own three cells',
    }
    (HERE / 'metadata.json').write_text(json.dumps(metadata, indent=2) + '\n')

SCENARIOS = ['read-owned', 'read-borrowed', 'closure-read', 'update-unique', 'update-retained']

def checks():
    rows = []
    for scenario in SCENARIOS:
        for sticky in ['normal', 'forced-sticky']:
            result = json.loads(run([BUILD / 'counts/probe', 'count', scenario, sticky, '10000']))
            plain = json.loads(run([BUILD / 'plain/probe', 'check', scenario, sticky, '10000']))
            assert plain['checksum'] == result['checksum'] and plain['elapsed_ns'] == 0
            rows.append(result)
            assert result['checksum'] == (170000 if scenario.startswith('read') or scenario == 'closure-read' else 80027)
            counts = result['counts']
            assert counts['max_normal_owners'] < 100
            if sticky == 'normal':
                assert counts['allocations'] == counts['frees']
                assert counts['sticky_clones'] == counts['sticky_drops'] == 0
            else:
                assert counts['allocations'] - counts['frees'] == 3
    lifecycle = json.loads(run([BUILD / 'counts/probe', 'lifecycle']))
    (HERE / 'counts.json').write_text(json.dumps({'rows': rows, 'lifecycle': lifecycle,
        'counter_scope': 'PerceusPtr cells only; closure Rc and other allocations excluded.',
        'plain_executable_same_checksums': True, 'no_timings': True}, indent=2) + '\n')
    for row in rows: print(json.dumps(row), flush=True)
    print(json.dumps(lifecycle), flush=True)

def times(authorized):
    if not authorized: raise SystemExit('Timing requires explicit --timing-authorized after coordinator approval.')
    metadata = json.loads((HERE / 'metadata.json').read_text())
    for relative, expected in metadata['frozen_sources'].items():
        assert sha(HERE / relative) == expected, f'Source changed after build: {relative}'
    for binary in metadata['binaries'].values():
        assert sha(Path(binary['path'])) == binary['sha256'], f'Binary changed: {binary["path"]}'
    rows = []
    for scenario in SCENARIOS:
        for pair in range(7):
            for sticky in (['normal', 'forced-sticky'] if pair % 2 == 0 else ['forced-sticky', 'normal']):
                row = json.loads(run([BUILD / 'plain/probe', 'time', scenario, sticky, '1000000']))
                row['pair'] = pair
                rows.append(row)
    summary = {scenario: {sticky: statistics.median(row['elapsed_ns'] for row in rows
        if row['scenario'] == scenario and row['sticky'] == sticky)
        for sticky in ['normal', 'forced-sticky']} for scenario in SCENARIOS}
    (HERE / 'timings.json').write_text(json.dumps({'rows': rows, 'summary_ns': summary,
        'scope': 'Synthetic forced saturation; exact Records kernel for update scenarios. Not a full suite benchmark.'}, indent=2) + '\n')
    print(json.dumps(summary, indent=2))

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('command', choices=['build', 'check', 'count', 'time'])
    parser.add_argument('--timing-authorized', action='store_true')
    args = parser.parse_args()
    if args.command == 'build': build()
    elif args.command in ['check', 'count']: checks()
    else: times(args.timing_authorized)
