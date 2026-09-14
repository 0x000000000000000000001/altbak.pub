#![allow(warnings)]
pub use purust_core::*;
use perceus_ptr::PerceusPtr;
use std::hint::black_box;
use std::rc::Rc;
use std::cell::Cell;
use std::time::Instant;
include!("build/records.rs");

fn initial() -> Value {
    Value::Record_a_b(PerceusPtr::new(Record_a_b {
        a: Some(mk_int(2)), b: Some(Value::Record_c_d(PerceusPtr::new(Record_c_d {
            c: Some(mk_int(3)), d: Some(Value::Record_e_f(PerceusPtr::new(Record_e_f {
                e: Some(mk_int(5)), f: Some(mk_int(17)),
            }))),
        }))),
    }))
}

// This changes the count in this private runtime copy only. It models the
// existing immortal state, not a realizable workload with four billion owners.
unsafe fn force_sticky(value: &Value) {
    match value {
        Value::Record_a_b(cell) => {
            force_sticky(cell.b.as_ref().unwrap());
            cell.set_count(PerceusPtr::<Record_a_b>::STICKY_COUNT);
        }
        Value::Record_c_d(cell) => {
            force_sticky(cell.d.as_ref().unwrap());
            cell.set_count(PerceusPtr::<Record_c_d>::STICKY_COUNT);
        }
        Value::Record_e_f(cell) => cell.set_count(PerceusPtr::<Record_e_f>::STICKY_COUNT),
        _ => panic!("unexpected generated record shape"),
    }
}

fn values(value: &Value) -> [i64; 4] {
    [value.__purust_borrow_a().unwrap_int(),
     value.__purust_borrow_b().__purust_borrow_c().unwrap_int(),
     value.__purust_borrow_b().__purust_borrow_d().__purust_borrow_e().unwrap_int(),
     value.__purust_borrow_b().__purust_borrow_d().__purust_borrow_f().unwrap_int()]
}

#[inline(never)]
fn read_owned(value: &Value) -> i64 {
    let owned = black_box(value.clone());
    owned.get_b().get_d().get_f().unwrap_int()
}
#[inline(never)]
fn read_borrowed(value: &Value) -> i64 {
    black_box(value).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f().unwrap_int()
}

fn workload(scenario: &str, sticky: bool, n: i64, timing: bool) -> (i64, u128) {
    let root = initial();
    if sticky { unsafe { force_sticky(&root) }; }
    let old = (scenario == "update-retained").then(|| root.clone());
    let callback = if scenario == "closure-read" {
        let capture = root.clone();
        Some(Func1::Shared(Rc::new(move |_: i64| read_owned(black_box(&capture)))))
    } else { None };
    let start = timing.then(Instant::now);
    let checksum = match scenario {
        "read-owned" => {
            let mut sum = 0;
            for _ in 0..black_box(n) { sum += read_owned(black_box(&root)); }
            black_box(sum)
        }
        "read-borrowed" => {
            let mut sum = 0;
            for _ in 0..black_box(n) { sum += read_borrowed(black_box(&root)); }
            black_box(sum)
        }
        "closure-read" => {
            let mut sum = 0;
            for i in 0..black_box(n) { sum += callback.as_ref().unwrap()(black_box(i)); }
            black_box(sum)
        }
        "update-unique" | "update-retained" => {
            let result = Test_Records_updateRec(black_box(n), black_box(root));
            let result_values = values(black_box(&result));
            let remainder_sum = (n / 5) * 10 + ((n % 5) * (n % 5 + 1) / 2);
            assert_eq!(result_values, [2 + n, 3 + 2*n, 5 + 3*n, 17 + remainder_sum]);
            if let Some(ref old) = old { assert_eq!(values(old), [2, 3, 5, 17]); }
            let sum = result_values.into_iter().sum();
            drop(result);
            drop(old);
            let elapsed = start.map_or(0, |start| start.elapsed().as_nanos());
            return (black_box(sum), elapsed);
        }
        _ => panic!("unknown scenario"),
    };
    assert_eq!(values(&root), [2, 3, 5, 17]);
    drop(callback); drop(root); drop(old);
    let elapsed = start.map_or(0, |start| start.elapsed().as_nanos());
    (checksum, elapsed)
}

#[derive(Clone)]
struct Tracked { value: i64, drops: Rc<Cell<usize>> }
impl Drop for Tracked { fn drop(&mut self) { self.drops.set(self.drops.get()+1); } }

fn lifecycle() {
    let dropped = Rc::new(Cell::new(0));
    let mut normal = PerceusPtr::new(Tracked { value: 11, drops: dropped.clone() });
    let alias = normal.clone();
    assert!(!normal.is_unique()); assert_eq!(normal.count(), 2);
    drop(alias); assert!(normal.is_unique()); assert_eq!(normal.count(), 1);
    PerceusPtr::make_mut(&mut normal).value = 12;
    drop(normal); assert_eq!(dropped.get(), 1);
    let mut sticky = PerceusPtr::new(Tracked { value: 21, drops: dropped.clone() });
    unsafe { sticky.set_count(PerceusPtr::<Tracked>::STICKY_COUNT); }
    let alias = sticky.clone();
    drop(alias); assert!(!sticky.is_unique());
    assert_eq!(sticky.count(), u32::MAX);
    let retained = sticky.clone();
    PerceusPtr::make_mut(&mut sticky).value = 22;
    assert_eq!(retained.value, 21); assert_eq!(sticky.value, 22);
    assert_eq!(retained.count(), u32::MAX); assert!(sticky.is_unique());
    drop(sticky); assert_eq!(dropped.get(), 2);
    drop(retained); assert_eq!(dropped.get(), 2, "immortal original payload is never dropped");
    println!("{{\"normal_returns_to_unique_after_two_owners\":true,\"normal_drop_runs\":true,\"sticky_count\":4294967295,\"sticky_stays_nonunique\":true,\"make_mut_copies_and_preserves_old\":true,\"sticky_payload_not_dropped\":true}}");
}

fn main() {
    let args: Vec<String> = std::env::args().collect();
    if args[1] == "lifecycle" { lifecycle(); return; }
    let timing = args[1] == "time";
    let scenario = &args[2];
    let sticky = &args[3];
    let n: i64 = args[4].parse().unwrap();
    if timing {
        // Same scenario and mode, with no Instant and no elapsed sample.
        // Forced-sticky warm-up intentionally retains its own three cells.
        black_box(workload(scenario, sticky == "forced-sticky", 10000, false));
    }
    let (checksum, elapsed) = workload(scenario, sticky == "forced-sticky", n, timing);
    #[cfg(sticky_counts)] {
        assert!(!timing, "timings must use the uninstrumented executable");
        let m = perceus_ptr::sticky_metrics::snapshot();
        println!(concat!("{{\"scenario\":\"{}\",\"sticky\":\"{}\",\"n\":{},\"checksum\":{},",
            "\"counts\":{{\"allocations\":{},\"frees\":{},\"allocated_bytes\":{},\"freed_bytes\":{},",
            "\"clone_calls\":{},\"drop_calls\":{},\"refcount_writes\":{},\"sticky_clones\":{},",
            "\"sticky_drops\":{},\"make_mut_copies\":{},\"max_normal_owners\":{}}}}}"),
            scenario, sticky, n, checksum, m[0],m[1],m[2],m[3],m[4],m[5],m[6],m[7],m[8],m[9],m[10]);
    }
    #[cfg(not(sticky_counts))]
    println!("{{\"scenario\":\"{}\",\"sticky\":\"{}\",\"n\":{},\"checksum\":{},\"elapsed_ns\":{}}}",
        scenario, sticky, n, checksum, elapsed);
}
