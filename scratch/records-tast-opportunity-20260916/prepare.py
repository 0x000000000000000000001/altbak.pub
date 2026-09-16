from pathlib import Path
import hashlib, json, subprocess
base=Path(__file__).resolve().parent
src=base.parents[1]/'output/purust_output/Purs_Test_Records/src/lib.rs'
text=src.read_text()
kernel=text[text.index('pub fn Test_Records_updateRec'):text.index('pub fn Test_Records_initial')]
(base/'kernel-original.rs').write_text(kernel)
variant='''#[inline(never)]
pub fn scalar_replacement(mut n: i64, mut r: Value) -> Value {
    if n == 0 { return r; }
    let mut a = r.__purust_borrow_a().unwrap_int();
    let mut c = r.__purust_borrow_b().__purust_borrow_c().unwrap_int();
    let mut e = r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_e().unwrap_int();
    let mut f = r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_f().unwrap_int();
    while n != 0 {
        a += 1; c += 2; e += 3;
        f += n.checked_rem_euclid(5).unwrap_or(0);
        n -= 1;
    }
    r.set_a(mk_int(a));
    let mut b = r.get_b(); r.set_b(Value::Unit);
    b.set_c(mk_int(c));
    let mut d = b.get_d(); b.set_d(Value::Unit);
    d.set_e(mk_int(e)); d.set_f(mk_int(f));
    b.set_d(d); r.set_b(b); r
}
'''
(base/'kernel-scalar.rs').write_text(variant)
harness='''#![allow(warnings)]
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
'''
(base/'probe.rs').write_text(harness)
deps=base.parents[1]/'output/purust_output/target/release/deps'
cmd=['rustc','--edition=2021','-C','opt-level=3',str(base/'probe.rs'),'-L',f'dependency={deps}','--extern',f'purust_core={next(deps.glob("libpurust_core-*.rlib"))}','--extern',f'perceus_ptr={next(deps.glob("libperceus_ptr-*.rlib"))}','-o',str(base/'probe')]
subprocess.run(cmd,check=True,timeout=60)
result=subprocess.check_output([str(base/'probe')],text=True,timeout=60)
(base/'results.txt').write_text(result)
(base/'manifest.json').write_text(json.dumps({'source':str(src),'sha256':hashlib.sha256(src.read_bytes()).hexdigest(),'command':cmd},indent=2)+'\n')
print(result)
import re
for name in ['original','scalar']:
 lines=[]
 for line in (base/f'kernel-{name}.rs').read_text().splitlines():
  for counter,pattern in [('READS',r'\.__purust_borrow_[a-f]\('),('GETS',r'\.get_[bd]\('),('SETS',r'\.set_[a-f]\(')]:
   n=len(re.findall(pattern,line))
   if n: lines.append(f' OP_{counter}.fetch_add({n},Ordering::Relaxed);')
  lines.append(line)
 (base/f'kernel-count-{name}.rs').write_text('\n'.join(lines)+'\n')
h=harness.replace('include!("kernel-original.rs");','include!("kernel-count-original.rs");').replace('include!("kernel-scalar.rs");','include!("kernel-count-scalar.rs");')
h=h.replace('struct Counter;','static OP_READS:AtomicUsize=AtomicUsize::new(0);\nstatic OP_GETS:AtomicUsize=AtomicUsize::new(0);\nstatic OP_SETS:AtomicUsize=AtomicUsize::new(0);\nstruct Counter;')
h=h.replace('ALLOCS.store(0,Ordering::Relaxed);','OP_READS.store(0,Ordering::Relaxed);OP_GETS.store(0,Ordering::Relaxed);OP_SETS.store(0,Ordering::Relaxed);ALLOCS.store(0,Ordering::Relaxed);',1)
h=h.replace('println!("{} n={} allocations=', 'println!("operations {} n={} reads={} cloning_getters={} setters={}",name,n,OP_READS.load(Ordering::Relaxed),OP_GETS.load(Ordering::Relaxed),OP_SETS.load(Ordering::Relaxed));\n   println!("{} n={} allocations=')
(base/'probe-count.rs').write_text(h)
countcmd=[str(base/'probe-count.rs') if a==str(base/'probe.rs') else str(base/'probe-count') if a==str(base/'probe') else a for a in cmd]
subprocess.run(countcmd,check=True,timeout=60)
result=subprocess.check_output([str(base/'probe-count')],text=True,timeout=60)
(base/'operation-counts.txt').write_text(result)
print(result)
# Separate timing binary: native mimalloc, no counter types, checks, or atomic state.
timeh='''#![allow(warnings)]
use purust_core::*;
use perceus_ptr::PerceusPtr;
use std::hint::black_box;
#[global_allocator] static GLOBAL:mimalloc::MiMalloc=mimalloc::MiMalloc;
include!("kernel-original.rs");
include!("kernel-scalar.rs");
'''+harness[harness.index('fn initial('):harness.index('fn main()')]+'''fn main(){
 let args:Vec<String>=std::env::args().collect();
 let fun:fn(i64,Value)->Value=if args[1]=="original" {Test_Records_updateRec}else{scalar_replacement};
 let seed=initial([0,0,0,0]);
 for _ in 0..3 {drop(black_box(fun(black_box(10000),black_box(seed.clone()))));}
 for _ in 0..15 {let t=std::time::Instant::now();let r=fun(black_box(10000),black_box(seed.clone()));let v=black_box(values(&r));drop(r);let ns=t.elapsed().as_nanos();assert_eq!(v,[10000,20000,30000,20000]);println!("{}",ns);}
}
'''
(base/'probe-time.rs').write_text(timeh)
timecmd=[str(base/'probe-time.rs') if a==str(base/'probe.rs') else str(base/'probe-time') if a==str(base/'probe') else a for a in cmd]
timecmd+=['--extern',f'mimalloc={next(deps.glob("libmimalloc-*.rlib"))}','-L',f'native={next((deps.parent/"build").glob("libmimalloc-sys-*/out/libmimalloc.a")).parent}']
subprocess.run(timecmd,check=True,timeout=60)
# Retained subrecords and extra fields at root: supplementary semantic checks.
extra=harness[:harness.index('fn main()')]+'''fn main(){
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
'''
(base/'probe-extra.rs').write_text(extra)
extracmd=[str(base/'probe-extra.rs') if a==str(base/'probe.rs') else str(base/'probe-extra') if a==str(base/'probe') else a for a in cmd]
subprocess.run(extracmd,check=True,timeout=60)
result=subprocess.check_output([str(base/'probe-extra')],text=True,timeout=60)
(base/'extra-checks.txt').write_text(result)
print(result)
