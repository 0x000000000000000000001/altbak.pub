"""Count allocations separately from timings in the two existing Unit prototypes."""
from pathlib import Path
import json, subprocess

here = Path(__file__).resolve().parent
deps = here.parents[1]/'output/purust_output/target/release/deps'
externs = []
for name in ['mimalloc', 'purust_core', 'Purs_Data_Unit']:
    files = list(deps.glob('lib'+name+'-*.rlib'))
    assert len(files) == 1, files
    externs += ['--extern', f'{name}={files[0]}']
instrument = r'''
use std::alloc::{GlobalAlloc, Layout};
use std::sync::atomic::{AtomicU64, Ordering};
static ALLOCS: AtomicU64 = AtomicU64::new(0);
static BYTES: AtomicU64 = AtomicU64::new(0);
static LARGE: AtomicU64 = AtomicU64::new(0);
struct AuditAlloc;
unsafe impl GlobalAlloc for AuditAlloc {
    unsafe fn alloc(&self, l: Layout) -> *mut u8 {
        ALLOCS.fetch_add(1, Ordering::Relaxed);
        BYTES.fetch_add(l.size() as u64, Ordering::Relaxed);
        if l.size() >= 6408 { LARGE.fetch_add(1, Ordering::Relaxed); }
        mimalloc::MiMalloc.alloc(l)
    }
    unsafe fn dealloc(&self, p: *mut u8, l: Layout) { mimalloc::MiMalloc.dealloc(p,l) }
}
#[global_allocator] static GLOBAL: AuditAlloc = AuditAlloc;
'''
main = r'''
fn main() {
    assert_eq!(Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000),0), 1000000);
    ALLOCS.store(0, Ordering::Relaxed);
    BYTES.store(0, Ordering::Relaxed);
    LARGE.store(0, Ordering::Relaxed);
    assert_eq!(Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000),0), 1000000);
    let count=ALLOCS.load(Ordering::Relaxed);
    let bytes=BYTES.load(Ordering::Relaxed);
    let large=LARGE.load(Ordering::Relaxed);
    println!("{count},{bytes},{large}");
}
'''
results = {}
for variant in ['lazy-generated', 'lazy-cached-unit']:
    source = (here/(variant+'.rs')).read_text().split('fn main() {')[0]
    source = source.replace('#[global_allocator] static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;', instrument)
    path = here/(variant+'-counts.rs')
    path.write_text(source+main)
    binary = here/(variant+'-counts')
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', *externs,
                    '-L', f'dependency={deps}', str(path), '-o', str(binary)], check=True)
    values = subprocess.check_output([str(binary)], text=True).strip().split(',')
    results[variant] = dict(zip(['allocations', 'requested_bytes', 'allocations_at_least_6408_bytes'], map(int, values)))
(here/'results-unit-alloc.json').write_text(json.dumps(results, indent=2)+'\n')
print(json.dumps(results, indent=2))
