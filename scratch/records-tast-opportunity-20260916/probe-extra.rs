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
 for fun in [Test_Records_updateRec as fn(i64,Value)->Value,scalar_replacement as fn(i64,Value)->Value]{
  let original=initial([4,8,12,16]);
  let b=original.get_b(); let d=b.get_d();
  let record=Value::Record_a_b_c(PerceusPtr::new(Record_a_b_c{a:Some(mk_int(4)),b:Some(b.clone()),c:Some(mk_int(12345))}));
  let snap=record.clone(); let result=fun(10000,record);
  assert_eq!(values(&result),[10004,20008,30012,20016]);assert_eq!(result.get_c().unwrap_int(),12345);
  assert_eq!(values(&snap),[4,8,12,16]);assert_eq!(snap.get_c().unwrap_int(),12345);
  assert_eq!(b.get_c().unwrap_int(),8);assert_eq!(d.get_e().unwrap_int(),12);assert_eq!(d.get_f().unwrap_int(),16);
 }
 println!("retained child aliases and unrelated root field c preserved in both variants");
}
