#![allow(non_snake_case)]
use purust_core::*;
use perceus_ptr::PerceusPtr;
use std::hint::black_box;
use call_fields::{record_step, scalar_step, Test_Records_updateRec};

#[cfg(not(counted))]
#[global_allocator]
static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;

#[cfg(counted)]
mod counts {
    use std::alloc::{GlobalAlloc, Layout};
    use std::sync::atomic::{AtomicBool, AtomicUsize, Ordering};
    pub static ON: AtomicBool = AtomicBool::new(false);
    pub static ALLOC: AtomicUsize = AtomicUsize::new(0);
    pub static FREE: AtomicUsize = AtomicUsize::new(0);
    pub static MATERIALIZE: AtomicUsize = AtomicUsize::new(0);
    pub struct Counter;
    unsafe impl GlobalAlloc for Counter {
        unsafe fn alloc(&self, l: Layout) -> *mut u8 {
            if ON.load(Ordering::Relaxed) { ALLOC.fetch_add(1, Ordering::Relaxed); }
            mimalloc::MiMalloc.alloc(l)
        }
        unsafe fn dealloc(&self, p: *mut u8, l: Layout) {
            if ON.load(Ordering::Relaxed) { FREE.fetch_add(1, Ordering::Relaxed); }
            mimalloc::MiMalloc.dealloc(p, l)
        }
    }
}
#[cfg(counted)]
#[global_allocator]
static GLOBAL: counts::Counter = counts::Counter;

include!("../records-tast-opportunity-20260916/kernel-scalar.rs");

fn initial(v: [i64; 4], extra: bool) -> Value {
    let b = Value::Record_c_d(PerceusPtr::new(Record_c_d {
        c: Some(mk_int(v[1])),
        d: Some(Value::Record_e_f(PerceusPtr::new(Record_e_f {
            e: Some(mk_int(v[2])), f: Some(mk_int(v[3])),
        }))),
    }));
    if extra {
        Value::Record_a_b_c(PerceusPtr::new(Record_a_b_c {
            a: Some(mk_int(v[0])), b: Some(b), c: Some(mk_int(12345)),
        }))
    } else {
        Value::Record_a_b(PerceusPtr::new(Record_a_b { a: Some(mk_int(v[0])), b: Some(b) }))
    }
}

#[inline]
fn values(r: &Value) -> [i64; 4] {
    [r.__purust_borrow_a().unwrap_int(),
     r.__purust_borrow_b().__purust_borrow_c().unwrap_int(),
     r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_e().unwrap_int(),
     r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_f().unwrap_int()]
}

#[inline]
fn materialize(mut r: Value, f: [i64; 4]) -> Value {
    #[cfg(counted)]
    counts::MATERIALIZE.fetch_add(1, std::sync::atomic::Ordering::Relaxed);
    r.set_a(mk_int(f[0]));
    let mut b = r.get_b(); r.set_b(Value::Unit);
    b.set_c(mk_int(f[1]));
    let mut d = b.get_d(); b.set_d(Value::Unit);
    d.set_e(mk_int(f[2])); d.set_f(mk_int(f[3]));
    b.set_d(d); r.set_b(b); r
}

type Kernel = fn(i64, Value, i64, &mut Vec<Value>) -> Value;

#[inline(never)]
fn run_record(mut n: i64, mut r: Value, stride: i64, snapshots: &mut Vec<Value>) -> Value {
    let mut i = 0;
    while n != 0 {
        r = record_step(n, r);
        n -= 1; i += 1;
        if stride > 0 && i % stride == 0 { snapshots.push(r.clone()); }
    }
    r
}

#[inline(never)]
fn run_materialized(mut n: i64, mut r: Value, stride: i64, snapshots: &mut Vec<Value>) -> Value {
    let mut i = 0;
    while n != 0 {
        let fields = scalar_step(n, values(&r));
        r = materialize(r, fields);
        n -= 1; i += 1;
        if stride > 0 && i % stride == 0 { snapshots.push(r.clone()); }
    }
    r
}

#[inline(never)]
fn run_scalar(mut n: i64, mut r: Value, stride: i64, snapshots: &mut Vec<Value>) -> Value {
    if n == 0 { return r; }
    let mut fields = values(&r);
    let mut dirty = false;
    let mut i = 0;
    while n != 0 {
        fields = scalar_step(n, fields);
        dirty = true;
        n -= 1; i += 1;
        if stride > 0 && i % stride == 0 {
            r = materialize(r, fields);
            dirty = false;
            snapshots.push(r.clone());
        }
    }
    if dirty { materialize(r, fields) } else { r }
}

#[inline(never)]
fn run_original(n: i64, r: Value, stride: i64, _: &mut Vec<Value>) -> Value {
    assert_eq!(stride, 0); Test_Records_updateRec(n, r)
}
#[inline(never)]
fn run_local(n: i64, r: Value, stride: i64, _: &mut Vec<Value>) -> Value {
    assert_eq!(stride, 0); scalar_replacement(n, r)
}

fn kernel(name: &str) -> Kernel {
    match name {
        "original" => run_original, "record" => run_record,
        "materialized" => run_materialized, "scalar" => run_scalar,
        "local" => run_local, _ => panic!("unknown kernel"),
    }
}

fn oracle(mut fields: [i64; 4], mut n: i64, stride: i64) -> ([i64; 4], Vec<[i64; 4]>) {
    let mut snaps = Vec::new(); let mut i = 0;
    while n != 0 {
        fields[0] = fields[0].wrapping_add(1);
        fields[1] = fields[1].wrapping_add(2);
        fields[2] = fields[2].wrapping_add(3);
        fields[3] = fields[3].wrapping_add(n.checked_rem_euclid(5).unwrap_or(0));
        n -= 1; i += 1;
        if stride > 0 && i % stride == 0 { snaps.push(fields); }
    }
    (fields, snaps)
}

fn validate() {
    let mut cases = 0; let mut checked_snapshots = 0;
    for seed_fields in [[0,0,0,0], [3,19,-100,7], [-5,999,12,-50], [i64::MAX-2,i64::MAX-2,i64::MAX-2,i64::MAX-2]] {
        for n in [0,1,2,3,4,5,6,31,100,10000] {
            for stride in [0,1,100] {
                let (expected, expected_snapshots) = oracle(seed_fields, n, stride);
                for extra in [false, true] {
                    let seed = initial(seed_fields, extra);
                    let child_b = seed.get_b(); let child_d = child_b.get_d();
                    for name in ["original", "record", "materialized", "scalar", "local"] {
                        if stride != 0 && (name == "original" || name == "local") { continue; }
                        let mut snapshots = Vec::new();
                        let result = kernel(name)(n, seed.clone(), stride, &mut snapshots);
                        assert_eq!(values(&result), expected, "{name}, n={n}, stride={stride}");
                        assert_eq!(values(&seed), seed_fields);
                        assert_eq!(snapshots.len(), expected_snapshots.len());
                        for (actual, expected) in snapshots.iter().zip(&expected_snapshots) {
                            assert_eq!(values(actual), *expected);
                            if extra { assert_eq!(actual.get_c().unwrap_int(), 12345); }
                            checked_snapshots += 1;
                        }
                        if extra { assert_eq!(result.get_c().unwrap_int(), 12345); }
                        assert_eq!(child_b.get_c().unwrap_int(), seed_fields[1]);
                        assert_eq!(child_d.get_e().unwrap_int(), seed_fields[2]);
                        assert_eq!(child_d.get_f().unwrap_int(), seed_fields[3]);
                        cases += 1;
                    }
                }
            }
        }
    }
    println!("validated {cases} cases and {checked_snapshots} retained snapshots, all four fields, old root/child aliases, extra fields and existing i64 release wrapping");
}

#[cfg(counted)]
fn count(name: &str, stride: i64) {
    use std::sync::atomic::Ordering::Relaxed;
    let seed = initial([0,0,0,0], false);
    let input = seed.clone(); let mut snapshots = Vec::new();
    call_fields::reset_counts();
    counts::ALLOC.store(0, Relaxed); counts::FREE.store(0, Relaxed); counts::MATERIALIZE.store(0, Relaxed);
    counts::ON.store(true, Relaxed);
    let result = kernel(name)(black_box(10000), input, stride, &mut snapshots);
    let fields = values(&result); let snapshots_len = snapshots.len();
    drop(result); drop(snapshots); counts::ON.store(false, Relaxed);
    assert_eq!(fields, [10000,20000,30000,20000]);
    println!("{{\"variant\":\"{name}\",\"stride\":{stride},\"calls\":{:?},\"materializations\":{},\"allocations\":{},\"deallocations\":{},\"snapshots\":{snapshots_len}}}",
        call_fields::call_counts(), counts::MATERIALIZE.load(Relaxed), counts::ALLOC.load(Relaxed), counts::FREE.load(Relaxed));
}

fn main() {
    let args: Vec<String> = std::env::args().collect();
    if args.len() == 1 || args[1] == "validate" { validate(); return; }
    let name = &args[2]; let stride: i64 = args[3].parse().unwrap();
    #[cfg(counted)]
    if args[1] == "count" { count(name, stride); return; }
    assert_eq!(args[1], "time");
    let fun = kernel(name); let seed = initial([0,0,0,0], false);
    for _ in 0..3 {
        let mut snaps = Vec::new();
        drop(black_box(fun(black_box(10000), black_box(seed.clone()), black_box(stride), &mut snaps)));
        drop(black_box(snaps));
    }
    for _ in 0..10 {
        let mut snaps = Vec::new();
        let start = std::time::Instant::now();
        let result = fun(black_box(10000), black_box(seed.clone()), black_box(stride), &mut snaps);
        let fields = black_box(values(&result)); black_box(&snaps);
        drop(result); drop(snaps);
        let ns = start.elapsed().as_nanos();
        assert_eq!(fields, [10000,20000,30000,20000]);
        println!("{ns}");
    }
}
