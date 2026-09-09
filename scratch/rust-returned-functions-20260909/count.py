"""Count allocations in unchanged extracts of the before/after generated Rust.
Requires the worktree runner dependencies; runs independently of timing.
"""
from pathlib import Path
import subprocess,json
r=Path(__file__).resolve().parent
build=r/'build';build.mkdir(exist_ok=True)
deps=r.parents[1]/'run/bak/rust/output/purust_output/target/release/deps'
extern=[]
for name in ['purust_core','Purs_Data_Unit','mimalloc']:
    paths=list(deps.glob('lib'+name+'-*.rlib'));assert len(paths)==1
    extern+=['--extern',f'{name}={paths[0]}']
allocator='''
use std::alloc::{GlobalAlloc, Layout};
use std::sync::atomic::{AtomicUsize, Ordering};
static ALLOCS: AtomicUsize = AtomicUsize::new(0);
static FREES: AtomicUsize = AtomicUsize::new(0);
static BYTES: AtomicUsize = AtomicUsize::new(0);
struct Count;
unsafe impl GlobalAlloc for Count {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        ALLOCS.fetch_add(1, Ordering::Relaxed);
        BYTES.fetch_add(layout.size(), Ordering::Relaxed);
        mimalloc::MiMalloc.alloc(layout)
    }
    unsafe fn dealloc(&self, p: *mut u8, layout: Layout) {
        FREES.fetch_add(1, Ordering::Relaxed);
        mimalloc::MiMalloc.dealloc(p, layout)
    }
}
#[global_allocator] static GLOBAL: Count = Count;
'''
main='''
fn main() {
    // The same captured input thunk remains callable across independent builds.
    for n in [0,1,2,17,1000] {
        let count=std::rc::Rc::new(std::cell::Cell::new(0));
        let counter=count.clone();
        let f=purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
            counter.set(counter.get()+1); -7
        }));
        for expected_calls in 1..=3 {
            assert_eq!(Test_LazyEvaluation_buildThunks(n, f.clone(), ()), n-7);
            assert_eq!(count.get(), expected_calls);
        }
    }
    for n in [0,1,2,10] {
        assert_eq!(Test_LazyEvaluation_runManyTimes(n, 17), 1000*n+17);
    }
    let a=ALLOCS.load(Ordering::Relaxed);
    let b=BYTES.load(Ordering::Relaxed);
    let f=FREES.load(Ordering::Relaxed);
    let result=Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000), 0);
    let allocations=ALLOCS.load(Ordering::Relaxed)-a;
    let bytes=BYTES.load(Ordering::Relaxed)-b;
    let frees=FREES.load(Ordering::Relaxed)-f;
    assert_eq!(result, 1_000_000);
    assert_eq!(allocations, frees);
    println!("{} {} {} {}", result, allocations, bytes, frees);
}
'''
results={}
for name in ['before','after']:
    source=(r/f'LazyEvaluation-{name}.rs').read_text()
    start=source.index('pub fn Test_LazyEvaluation_buildThunks(')
    end=source.index('pub fn Test_LazyEvaluation_act(')
    s='#![allow(warnings)]\nuse purust_core::*;\nuse Purs_Data_Unit::Data_Unit_unit;\n'+allocator+source[start:end]
    src=build/(name+'-count.rs');src.write_text(s+main)
    binary=build/(name+'-count')
    subprocess.run(['rustc','--edition=2021','-C','opt-level=1',*extern,'-L',f'dependency={deps}',str(src),'-o',str(binary)],check=True)
    values=subprocess.check_output([str(binary)],text=True).split()
    results[name]=dict(zip(['result','allocations','bytes_requested','deallocations'],map(int,values)))
(r/'allocations.json').write_text(json.dumps(results,indent=2)+'\n')
print(json.dumps(results,indent=2))
