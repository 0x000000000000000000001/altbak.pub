#![allow(warnings)]
use purust_core::*;
use perceus_ptr::PerceusPtr;
use std::hint::black_box;
#[global_allocator] static GLOBAL:mimalloc::MiMalloc=mimalloc::MiMalloc;
include!("kernel-original.rs");
include!("kernel-scalar.rs");
fn initial(v:[i64;4])->Value {
 Value::Record_a_b(PerceusPtr::new(Record_a_b{a:Some(mk_int(v[0])),b:Some(Value::Record_c_d(PerceusPtr::new(Record_c_d{c:Some(mk_int(v[1])),d:Some(Value::Record_e_f(PerceusPtr::new(Record_e_f{e:Some(mk_int(v[2])),f:Some(mk_int(v[3]))})))})))}))
}
fn values(r:&Value)->[i64;4]{[r.__purust_borrow_a().unwrap_int(),r.__purust_borrow_b().__purust_borrow_c().unwrap_int(),r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_e().unwrap_int(),r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_f().unwrap_int()]}
fn main(){
 let args:Vec<String>=std::env::args().collect();
 let fun:fn(i64,Value)->Value=if args[1]=="original" {Test_Records_updateRec}else{scalar_replacement};
 let seed=initial([0,0,0,0]);
 for _ in 0..3 {drop(black_box(fun(black_box(10000),black_box(seed.clone()))));}
 for _ in 0..15 {let t=std::time::Instant::now();let r=fun(black_box(10000),black_box(seed.clone()));let v=black_box(values(&r));drop(r);let ns=t.elapsed().as_nanos();assert_eq!(v,[10000,20000,30000,20000]);println!("{}",ns);}
}
