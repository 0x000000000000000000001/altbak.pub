"""Compare allocations in the saved, actually generated LazyEvaluation kernels."""
from pathlib import Path
import json
import re
import subprocess

HERE = Path(__file__).resolve().parent
DEPS = HERE.parents[1] / 'run/bak/rust/output/purust_output/target/release/deps'
BUILD = HERE / 'build'
BUILD.mkdir(exist_ok=True)
extern = []
for name in ['purust_core', 'Purs_Data_Unit', 'mimalloc']:
    paths = list(DEPS.glob('lib' + name + '-*.rlib'))
    assert len(paths) == 1, (name, paths)
    extern += ['--extern', f'{name}={paths[0]}']

allocator = '''
use std::alloc::{GlobalAlloc, Layout};
use std::sync::atomic::{AtomicUsize, Ordering};
static ALLOCS: AtomicUsize = AtomicUsize::new(0);
static FREES: AtomicUsize = AtomicUsize::new(0);
static BYTES: AtomicUsize = AtomicUsize::new(0);
struct Count;
unsafe impl GlobalAlloc for Count {
    unsafe fn alloc(&self, l: Layout) -> *mut u8 {
        ALLOCS.fetch_add(1, Ordering::Relaxed);
        BYTES.fetch_add(l.size(), Ordering::Relaxed);
        mimalloc::MiMalloc.alloc(l)
    }
    unsafe fn dealloc(&self, p: *mut u8, l: Layout) {
        FREES.fetch_add(1, Ordering::Relaxed);
        mimalloc::MiMalloc.dealloc(p, l);
    }
}
#[global_allocator] static GLOBAL: Count = Count;
'''
checks = '''
fn main() {
    for depth in [0, 1, 2, 17, 1000] {
        let calls = std::rc::Rc::new(std::cell::Cell::new(0));
        let observer = calls.clone();
        let seed = purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
            observer.set(observer.get() + 1); -7
        }));
        for count in 1..=3 {
            assert_eq!(Test_LazyEvaluation_buildThunks(depth, seed.clone(), ()), depth - 7);
            assert_eq!(calls.get(), count);
        }
    }
    for n in [0, 1, 2, 10] {
        assert_eq!(Test_LazyEvaluation_runManyTimes(n, 17), 1000 * n + 17);
    }
    let a = ALLOCS.load(Ordering::Relaxed);
    let f = FREES.load(Ordering::Relaxed);
    let b = BYTES.load(Ordering::Relaxed);
    let result = std::hint::black_box(Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000), 0));
    let a = ALLOCS.load(Ordering::Relaxed) - a;
    let f = FREES.load(Ordering::Relaxed) - f;
    let b = BYTES.load(Ordering::Relaxed) - b;
    assert_eq!(result, 1_000_000);
    assert_eq!(a, f);
    println!("{a} {f} {b}");
}
'''

def kernel(source):
    chunks = []
    pattern = re.compile(r'^(?:pub )?fn Test_LazyEvaluation_(?:buildThunks(?:__purust_strict_thunk_\d+)?|runManyTimes)\(', re.M)
    cursor = 0
    while match := pattern.search(source, cursor):
        opening = source.index('{', match.end())
        cursor = opening + 1
        depth = 1
        while depth:
            depth += (source[cursor] == '{') - (source[cursor] == '}')
            cursor += 1
        chunks.append(source[match.start():cursor])
    assert len(chunks) in [2, 3]
    return '\n\n'.join(chunks)

results = {}
for side in ['before', 'after']:
    code = kernel((HERE / f'LazyEvaluation-{side}.rs').read_text())
    (BUILD / f'{side}-kernel.rs').write_text(code)
    src = BUILD / f'{side}-count.rs'
    binary = BUILD / f'{side}-count'
    src.write_text('#![allow(warnings)]\nuse purust_core::*;\nuse Purs_Data_Unit::Data_Unit_unit;\n' + allocator + code + checks)
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', *extern,
                    '-L', f'dependency={DEPS}', str(src), '-o', str(binary)], check=True)
    values = list(map(int, subprocess.check_output([str(binary)], text=True).split()))
    results[side] = dict(zip(['allocations', 'deallocations', 'bytes_requested'], values))
(HERE / 'allocations.json').write_text(json.dumps(results, indent=2) + '\n')
print(json.dumps(results, indent=2))
