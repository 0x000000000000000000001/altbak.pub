from pathlib import Path
import json,re,subprocess,statistics,sys
HERE=Path(__file__).resolve().parent
BUILD=HERE/'build'
BUILD.mkdir(exist_ok=True)
ROOT=HERE.parents[1]
DEPS=ROOT/'run/bak/rust/output/purust_output/target/release/deps'
source=(HERE/'generated-before.rs').read_text()
original=source[source.index('pub fn Test_Records_updateRec('):source.index('pub fn Test_Records_describe(')]
initial=original[original.index('pub fn Test_Records_initial('):]
HEADER='#![allow(warnings)]\npub use purust_core::*;\nuse perceus_ptr::PerceusPtr;\n'
STAGED='''
pub fn Test_Records_updateRec(mut n:i64,mut r:UnknownType)->UnknownType {
 while n!=0 {
  // Evaluate every RHS in source order, then move the old root at last use.
  let a=mk_int(r.get_a().unwrap_int()+1);
  let b={
   let mut b=r.get_b();
   b.set_c(mk_int(r.get_b().get_c().unwrap_int()+2));
   b.set_d({
    let mut d=r.get_b().get_d();
    d.set_e(mk_int(r.get_b().get_d().get_e().unwrap_int()+3));
    d.set_f(mk_int(r.get_b().get_d().get_f().unwrap_int()+n.checked_rem_euclid(5).unwrap_or(0)));
    d
   });
   b
  };
  let mut base=r;
  base.set_a(a);base.set_b(b);
  r=base;n-=1;
 }
 r
}
'''
PATH='''
pub fn Test_Records_updateRec(mut n:i64,mut r:UnknownType)->UnknownType {
 while n!=0 {
  // These scalar RHSs neither retain the source nor invoke a callback.
  let a=mk_int(r.get_a().unwrap_int()+1);
  let c=mk_int(r.get_b().get_c().unwrap_int()+2);
  let e=mk_int(r.get_b().get_d().get_e().unwrap_int()+3);
  let f=mk_int(r.get_b().get_d().get_f().unwrap_int()+n.checked_rem_euclid(5).unwrap_or(0));
  // Keep exactly the current Value/Option/PerceusPtr representation.
  let Value::Record_a_b(root)= &mut r else {unreachable!()};
  let root=PerceusPtr::make_mut(root);
  root.a=Some(a);
  let Value::Record_c_d(middle)=root.b.as_mut().unwrap() else {unreachable!()};
  let middle=PerceusPtr::make_mut(middle);
  middle.c=Some(c);
  let Value::Record_e_f(leaf)=middle.d.as_mut().unwrap() else {unreachable!()};
  let leaf=PerceusPtr::make_mut(leaf);
  leaf.e=Some(e);leaf.f=Some(f);
  n-=1;
 }
 r
}
'''
HELPERS='''
fn values(r:&UnknownType)->[i64;4] {
 [r.get_a().unwrap_int(),r.get_b().get_c().unwrap_int(),r.get_b().get_d().get_e().unwrap_int(),r.get_b().get_d().get_f().unwrap_int()]
}
fn expected(n:i64, start:[i64;4])->[i64;4] {
 [start[0]+n,start[1]+2*n,start[2]+3*n,start[3]+(1..=n).map(|i|i%5).sum::<i64>()]
}
'''
TIME='''
#[global_allocator] static GLOBAL:mimalloc::MiMalloc=mimalloc::MiMalloc;
fn main(){
 for i in 0..21 {
  let n=std::hint::black_box(10000_i64);
  let start=std::time::Instant::now();
  let r=Test_Records_updateRec(n,Test_Records_initial());
  let found=std::hint::black_box(values(&r));drop(r);
  let elapsed=start.elapsed().as_nanos();
  assert_eq!(found,expected(n,[0;4]));
  if i>0 {println!("{elapsed}");}
 }
}
'''
COUNT='''
use std::alloc::{GlobalAlloc,Layout};
use std::sync::atomic::{AtomicUsize,Ordering};
static ALLOCS:AtomicUsize=AtomicUsize::new(0);static FREES:AtomicUsize=AtomicUsize::new(0);
struct Counter;
unsafe impl GlobalAlloc for Counter {
 unsafe fn alloc(&self,l:Layout)->*mut u8 {ALLOCS.fetch_add(1,Ordering::Relaxed);mimalloc::MiMalloc.alloc(l)}
 unsafe fn dealloc(&self,p:*mut u8,l:Layout){FREES.fetch_add(1,Ordering::Relaxed);mimalloc::MiMalloc.dealloc(p,l)}
}
#[global_allocator] static GLOBAL:Counter=Counter;
fn main(){
 for n in [0_i64,1,2,10,10000] {
  ALLOCS.store(0,Ordering::Relaxed);FREES.store(0,Ordering::Relaxed);
  let r=Test_Records_updateRec(std::hint::black_box(n),Test_Records_initial());
  let found=values(&r);drop(r);
  let allocations=ALLOCS.load(Ordering::Relaxed);let frees=FREES.load(Ordering::Relaxed);
  assert_eq!(allocations,frees);assert_eq!(found,expected(n,[0;4]));
  println!("{n} {allocations} {frees}");
 }
}
'''
CHECK='''
fn main(){
 for seed in [0_i64,1,7,50] {
  for n in [0_i64,1,2,5,10,10000] {
   let r=Test_Records_updateRec(seed,Test_Records_initial());
   let old=values(&r);
   let kept=r.clone();
   let old_b=r.get_b();let old_d=old_b.get_d();
   let result=Test_Records_updateRec(n,r);
   assert_eq!(values(&result),expected(n,old));assert_eq!(values(&kept),old);
   assert_eq!(old_b.get_c().unwrap_int(),old[1]);assert_eq!(old_d.get_e().unwrap_int(),old[2]);
   assert_eq!(old_d.get_f().unwrap_int(),old[3]);
   // Sharing a nested node alone must also retain its original version.
   let r=Test_Records_updateRec(seed,Test_Records_initial());
   let kept_leaf=r.get_b().get_d();
   let result=Test_Records_updateRec(n,r);
   assert_eq!(values(&result),expected(n,old));
   assert_eq!(kept_leaf.get_e().unwrap_int(),old[2]);
   assert_eq!(kept_leaf.get_f().unwrap_int(),old[3]);
  }
 }
 println!("24 seed/count combinations: all four fields, retained roots, nested and leaf-only sharing passed");
}
'''
variants={'before':original,'root_move':STAGED+initial,'owned_path':PATH+initial}
mode=sys.argv[1]
for name,code in variants.items():
 path=BUILD/f'{mode}-{name}.rs';path.write_text(HEADER+code+HELPERS+{'time':TIME,'count':COUNT,'check':CHECK}[mode])
 deps=[]
 for lib in ['purust_core','perceus_ptr','mimalloc']:
  matches=list(DEPS.glob(f'lib{lib}-*.rlib'));assert len(matches)==1,(lib,matches)
  deps+=['--extern',f'{lib}={matches[0]}']
 subprocess.run(['rustc','--edition=2021','-C','opt-level=1',str(path),'-o',str(BUILD/f'{mode}-{name}'),'-L',f'dependency={DEPS}',*deps],check=True)
if mode=='time':
 runs={name:[] for name in variants}
 for pair in range(5):
  for name in (list(variants) if pair%2==0 else list(reversed(variants))):
   samples=list(map(int,subprocess.check_output([str(BUILD/f'time-{name}')],text=True).split()))
   assert len(samples)==20;runs[name].append(samples)
   print(pair+1,name,statistics.median(samples)/1e6,flush=True)
 result={'method':'Five alternating triplets, twenty timed samples after one warmup; O1, mimalloc; all binaries compiled first; full four-field result observed, destruction included; no instrumentation in timed binaries.', 'samples_ns':runs,'medians_ms':{n:statistics.median(sum(v,[]))/1e6 for n,v in runs.items()}}
else:
 result={n:subprocess.check_output([str(BUILD/f'{mode}-{n}')],text=True) for n in variants}
(HERE/f'{mode}.json').write_text(json.dumps(result,indent=2)+'\n')
print(result.get('medians_ms',result))
