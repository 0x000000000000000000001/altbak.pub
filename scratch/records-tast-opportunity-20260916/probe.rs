#![allow(warnings)]
use purust_core::*;
use perceus_ptr::PerceusPtr;
use std::alloc::{GlobalAlloc, Layout, System};
use std::sync::atomic::{AtomicBool, AtomicUsize, Ordering};
use std::hint::black_box;
static ON: AtomicBool=AtomicBool::new(false);
static ALLOCS: AtomicUsize=AtomicUsize::new(0);
static DEALLOCS: AtomicUsize=AtomicUsize::new(0);
static BYTES: AtomicUsize=AtomicUsize::new(0);
struct Counter;
unsafe impl GlobalAlloc for Counter {
 unsafe fn alloc(&self, l: Layout)->*mut u8 {if ON.load(Ordering::Relaxed){ALLOCS.fetch_add(1,Ordering::Relaxed); BYTES.fetch_add(l.size(),Ordering::Relaxed);}System.alloc(l)}
 unsafe fn dealloc(&self,p:*mut u8,l:Layout){if ON.load(Ordering::Relaxed){DEALLOCS.fetch_add(1,Ordering::Relaxed);}System.dealloc(p,l)}
}
#[global_allocator] static GLOBAL: Counter=Counter;
include!("kernel-original.rs");
include!("kernel-scalar.rs");
fn initial(v:[i64;4])->Value {
 Value::Record_a_b(PerceusPtr::new(Record_a_b{a:Some(mk_int(v[0])),b:Some(Value::Record_c_d(PerceusPtr::new(Record_c_d{c:Some(mk_int(v[1])),d:Some(Value::Record_e_f(PerceusPtr::new(Record_e_f{e:Some(mk_int(v[2])),f:Some(mk_int(v[3]))})))})))}))
}
fn values(r:&Value)->[i64;4]{[r.__purust_borrow_a().unwrap_int(),r.__purust_borrow_b().__purust_borrow_c().unwrap_int(),r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_e().unwrap_int(),r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_f().unwrap_int()]}
fn main(){
 let args:Vec<String>=std::env::args().collect();
 if args.len()>1 && args[1]=="time" {
   let fun:fn(i64,Value)->Value=if args[2]=="original" {Test_Records_updateRec}else{scalar_replacement};
   let seed=initial([0,0,0,0]);
   for _ in 0..3 {drop(black_box(fun(black_box(10000),black_box(seed.clone()))));}
   for _ in 0..15 {let t=std::time::Instant::now();let r=fun(black_box(10000),black_box(seed.clone()));black_box(values(&r));drop(r);println!("{}",t.elapsed().as_nanos());}
   return;
 }
 for start in [[0,0,0,0],[3,19,-100,7],[-5,999,12,-50]]{
  for n in [0,1,2,3,4,5,6,31,100,10000]{
   let seed=initial(start);let snap=seed.clone();
   let a=Test_Records_updateRec(n,seed.clone());let b=scalar_replacement(n,seed.clone());
   assert_eq!(values(&a),values(&b));assert_eq!(values(&seed),start);assert_eq!(values(&snap),start);
  }
 }
 println!("all 30 cases and retained snapshots match");
 println!("Value={} Option<Value>={} Record_a_b={}",std::mem::size_of::<Value>(),std::mem::size_of::<Option<Value>>(),std::mem::size_of::<Record_a_b>());
 for (name,fun) in [("original",Test_Records_updateRec as fn(i64,Value)->Value),("scalar",scalar_replacement as fn(i64,Value)->Value)] {
  for n in [0,1,10000] {
   let seed=initial([0,0,0,0]);let arg=seed.clone();
   ALLOCS.store(0,Ordering::Relaxed);DEALLOCS.store(0,Ordering::Relaxed);BYTES.store(0,Ordering::Relaxed);ON.store(true,Ordering::Relaxed);
   let r=fun(black_box(n),black_box(arg)); let vals=values(&r);drop(r);
   ON.store(false,Ordering::Relaxed);
   println!("{} n={} allocations={} deallocations={} bytes={} result={:?}",name,n,ALLOCS.load(Ordering::Relaxed),DEALLOCS.load(Ordering::Relaxed),BYTES.load(Ordering::Relaxed),vals);
  }
 }
}
