"""Measure an isolated accumulator change against the exact generated module.

Requires the existing release dependencies from bin/rust/run. No generated
workspace source is changed; outputs stay under this scratch directory.
"""
from pathlib import Path
import difflib
import hashlib
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
GENERATED = HERE.parents[1] / 'run/bak/rust/output/purust_output'
BUILD = HERE / 'build'
BUILD.mkdir(exist_ok=True)
DEPS = GENERATED / 'target/release/deps'
before = (GENERATED / 'Purs_Test_Polymorphism/src/lib.rs').read_text()
start = before.index('        fn purs_local_2_rec_0_impl(')
end = before.index('        crate::Value::Func2(', start)
kernel = before[start:end].strip()
native = kernel.replace('mut purs_local_4: crate::UnknownType) -> crate::UnknownType',
                        'mut purs_local_4: i64) -> i64')
old_add = 'crate::mk_int(/* Typed crate::UnknownType <- i64 : PrimOp(...) */((purs_local_4.clone()).unwrap_int() + 1))'
assert native.count(old_add) == 1
native = native.replace(old_add, '(purs_local_4 + 1)')
after = before[:start] + '        ' + native + '\n' + before[end:]
old_call = 'purs_local_2_rec_0_impl(purs_local_2_rec_0, purs_local_3, purs_local_4)'
assert after.count(old_call) == 1
after = after.replace(old_call,
    'crate::mk_int(purs_local_2_rec_0_impl(purs_local_2_rec_0, purs_local_3, purs_local_4.unwrap_int()))')
(HERE / 'Polymorphism-before.rs').write_text(before)
(HERE / 'Polymorphism-prototype.rs').write_text(after)
(HERE / 'prototype.diff').write_text(''.join(difflib.unified_diff(
    before.splitlines(True), after.splitlines(True), fromfile='before.rs', tofile='prototype.rs')))

def dependency(name):
    paths = list(DEPS.glob(f'lib{name}-*.rlib'))
    assert len(paths) == 1, (name, paths)
    return paths[0]

def run(args):
    proc = subprocess.run([str(a) for a in args], capture_output=True, text=True)
    if proc.returncode:
        raise RuntimeError(f'{args[0]} failed:\n{proc.stdout}\n{proc.stderr}')
    return proc.stdout

manifest = (GENERATED / 'Purs_Test_Polymorphism/Cargo.toml').read_text().split('[dependencies]\n')[1]
deps = [name.replace('-', '_') for name in re.findall(r'^([\w-]+) =', manifest, re.M)]
externs = [arg for name in deps for arg in ['--extern', f'{name}={dependency(name)}']]
common = ['rustc', '--edition=2021', '-C', 'opt-level=1', '-C', 'debuginfo=2', '-L', f'dependency={DEPS}']
for side, source, probe in [('before', before, kernel), ('native', after, native)]:
    crate = f'poly_{side}'
    path = BUILD / f'{crate}.rs'
    probe = probe.replace('fn purs_local_2_rec_0_impl(', 'pub fn kernel(')
    path.write_text(source + '\n' + probe + '\n')
    run(common + ['--crate-type=rlib', '--crate-name', crate, path,
                  '--emit', f'link={BUILD}/lib{crate}.rlib,asm={BUILD}/{crate}.s'] + externs)
    call = ('poly::kernel(Value::Unit, n, mk_int(initial)).unwrap_int()' if side == 'before'
            else 'poly::kernel(Value::Unit, n, initial)')
    harness = '''#![allow(warnings)]
use purust_core::*;
use std::hint::black_box;
#[global_allocator] static ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;
fn main() {
    for n in [0_i64, 1, 2, 7, 1000, 100000] {
        for initial in [-42_i64, 0, 17] {
            let n = black_box(n);
            let initial = black_box(initial);
            let result = CALL;
            assert_eq!(result, initial + n);
        }
    }
    let effect = poly::Test_Polymorphism_act().unwrap_func1();
    let execute = || {
        let result = black_box(effect(Value::Unit));
        assert_eq!(result.unwrap_string(), "10000000");
    };
    execute();
    let mut samples = Vec::new();
    for _ in 0..10 {
        let start = std::time::Instant::now();
        for _ in 0..BATCH { execute(); }
        samples.push(start.elapsed().as_secs_f64() * 1e6 / BATCH as f64);
    }
    println!("{:?}", samples);
}
'''.replace('CALL', call).replace('BATCH', '1' if side == 'before' else '1024')
    path = BUILD / f'{side}.rs'
    path.write_text(harness)
    run(common + [path, '-o', BUILD / side,
        '--extern', f'poly={BUILD}/lib{crate}.rlib',
        '--extern', f'purust_core={dependency("purust_core")}',
        '--extern', f'mimalloc={dependency("mimalloc")}'])
    print(f'{side}: compiled', flush=True)

samples = {side: [] for side in ['before', 'native']}
for pair in range(7):
    for side in (['before', 'native'] if pair % 2 == 0 else ['native', 'before']):
        timings = json.loads(run([BUILD / side]))
        samples[side].append(timings)
        print(f'{pair + 1} {side}: best {min(timings):.3f} us', flush=True)
results = {
    'method': '7 alternating pairs of processes, warmup + best of 10; O1, mimalloc; complete generated act with original effects/adapters; batches of 1 before and 1024 native (normalized per call); 18 parameterized kernel assertions per process',
    'rustc': run(['rustc', '--version']).strip(),
    'sha256_before': hashlib.sha256(before.encode()).hexdigest(),
    'samples_us': samples,
    'median_best_us': {side: statistics.median(map(min, values)) for side, values in samples.items()},
}
(HERE / 'results.json').write_text(json.dumps(results, indent=2) + '\n')
print(json.dumps(results['median_best_us']))
