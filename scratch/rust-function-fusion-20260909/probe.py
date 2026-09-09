"""Isolate the Church function producer and fuse only its immediate application."""
from pathlib import Path
import argparse
import hashlib
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE/'build'
GENERATED = ROOT/'run/bak/rust/output/purust_output'
SOURCE = GENERATED/'Purs_Test_Church/src/lib.rs'
DEPS = GENERATED/'target/release/deps'

def functions_in(source):
    starts = list(re.finditer(r'^pub fn (Test_Church_\w+)\(', source, re.M))
    return {m[1]: source[m.start():starts[i+1].start() if i+1 < len(starts) else len(source)] for i,m in enumerate(starts)}

def variants(integrated=False):
    path = BUILD/'Church-before.rs'
    source = path.read_text() if path.exists() else SOURCE.read_text()
    functions = functions_in(source)
    names = ['zeroC','fromInt','c100','c10k','c100k']
    before = '\n'.join(functions['Test_Church_'+name] for name in names)
    original = functions['Test_Church_fromInt']
    start = original.index('{')+1
    end = original.rfind('}')
    prototype = original[:start]+'''
    if a0 >= 0 {
        let mut n = a0;
        let mut result = a2;
        while n > 0 { result = a1(result); n -= 1; }
        result
    } else {'''+original[start:end]+'}\n}\n'
    result = {'before': before, 'prototype': before.replace(original, prototype)}
    if integrated:
        generated = functions_in(SOURCE.read_text())
        assert '_function_count' in generated['Test_Church_fromInt']
        result['integrated'] = '\n'.join(generated['Test_Church_'+name] for name in names)
    return result

HEADER = '#![allow(warnings)]\npub use purust_core::*;\n'
TIME = '''
#[global_allocator] static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;
fn main() {
    for i in 0..21 {
        let start = std::time::Instant::now();
        let value = Test_Church_c100k(std::hint::black_box(10), Func1::Static(|x| x+1), 0);
        assert_eq!(value,100000);
        if i>0 { println!("{}",start.elapsed().as_nanos()); }
    }
}
'''
COUNT = '''
use std::alloc::{GlobalAlloc, Layout};
use std::sync::atomic::{AtomicUsize, Ordering};
static ALLOCS: AtomicUsize=AtomicUsize::new(0);
static FREES: AtomicUsize=AtomicUsize::new(0);
struct Counter;
unsafe impl GlobalAlloc for Counter {
 unsafe fn alloc(&self,l:Layout)->*mut u8 {ALLOCS.fetch_add(1,Ordering::Relaxed);mimalloc::MiMalloc.alloc(l)}
 unsafe fn dealloc(&self,p:*mut u8,l:Layout){FREES.fetch_add(1,Ordering::Relaxed);mimalloc::MiMalloc.dealloc(p,l)}
}
#[global_allocator] static GLOBAL:Counter=Counter;
fn main() {
 for n in [0_i64,1,2,3,10] {
  ALLOCS.store(0,Ordering::Relaxed);FREES.store(0,Ordering::Relaxed);
  let value=Test_Church_c100k(std::hint::black_box(n),Func1::Static(|x|x+1),0);
  let allocations=ALLOCS.load(Ordering::Relaxed);let frees=FREES.load(Ordering::Relaxed);
  assert_eq!(value,n.pow(5));assert_eq!(allocations,frees);
  println!("{n} {value} {allocations} {frees}");
 }
}
'''
CHECKS = '''
fn main() {
 for n in [0_i64,1,2,3,10,30] {
  let calls=std::rc::Rc::new(std::cell::RefCell::new(Vec::new()));
  let observed=calls.clone();
  let f=Func1::Shared(std::rc::Rc::new(move |x| {observed.borrow_mut().push(x); 2*x+1}));
  let value=Test_Church_fromInt(n,f,0);
  let mut expected=0;let mut trace=Vec::new();
  for _ in 0..n {trace.push(expected);expected=2*expected+1;}
  assert_eq!(value,expected);assert_eq!(*calls.borrow(),trace);
  assert_eq!(std::rc::Rc::strong_count(&calls),1);
 }
 // A panic stops at precisely the same callback, and drops its captures.
 std::panic::set_hook(Box::new(|_| {}));
 let calls=std::rc::Rc::new(std::cell::Cell::new(0));let observed=calls.clone();
 let f=Func1::Shared(std::rc::Rc::new(move |x| {observed.set(observed.get()+1);if x==3 {panic!("stop")};x+1}));
 assert!(std::panic::catch_unwind(std::panic::AssertUnwindSafe(||Test_Church_fromInt(10,f,0))).is_err());
 assert_eq!(calls.get(),4);assert_eq!(std::rc::Rc::strong_count(&calls),1);
 println!("callback order, zero applications, capture lifetime and panic checked");
}
'''

def compile_variants(mode, integrated):
    core, = DEPS.glob('libpurust_core-*.rlib')
    allocator, = DEPS.glob('libmimalloc-*.rlib')
    harness = {'time': TIME, 'count': COUNT, 'check': CHECKS}[mode]
    for name, code in variants(integrated).items():
        path = BUILD/f'{mode}-{name}.rs'
        path.write_text(HEADER+code+harness)
        subprocess.run(['rustc','--edition=2021','-C','opt-level=1',str(path),'-o',str(BUILD/f'{mode}-{name}'),
            '--extern',f'purust_core={core}','--extern',f'mimalloc={allocator}','-L',f'dependency={DEPS}'],check=True)

if __name__=='__main__':
    parser=argparse.ArgumentParser()
    parser.add_argument('mode',choices=['time','count','check'])
    parser.add_argument('--integrated', action='store_true')
    args=parser.parse_args()
    BUILD.mkdir(exist_ok=True)
    if not (BUILD/'Church-before.rs').exists(): (BUILD/'Church-before.rs').write_text(SOURCE.read_text())
    compile_variants(args.mode, args.integrated)
    if args.mode=='time':
        runs={name:[] for name in variants(args.integrated)}
        for pair in range(5):
            for name in (list(runs) if pair%2==0 else list(reversed(runs))):
                samples=list(map(int,subprocess.check_output([str(BUILD/f'time-{name}')],text=True).split()))
                assert len(samples)==20
                runs[name].append(samples)
                print(pair+1,name,statistics.median(samples)/1e6,'ms',flush=True)
        result={'method':'Five alternating pairs, twenty samples after one warmup; O1/mimalloc; no instrumentation or compilation during timing.',
                'runs_ns':runs,'median_ms':{name:statistics.median(sum(v,[]))/1e6 for name,v in runs.items()}}
    else:
        result={name:subprocess.check_output([str(BUILD/f'{args.mode}-{name}')],text=True) for name in variants(args.integrated)}
    suffix = '-integrated' if args.integrated else ''
    (HERE/f'{args.mode}{suffix}.json').write_text(json.dumps(result,indent=2)+'\n')
    print(result.get('median_ms',result),flush=True)
