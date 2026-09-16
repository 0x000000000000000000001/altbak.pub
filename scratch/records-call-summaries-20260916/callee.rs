#![allow(warnings)]
use purust_core::*;

// The baseline below is extracted byte-for-byte from generated output.
include!("kernel-original.rs");

#[cfg(counted)]
use std::sync::atomic::{AtomicUsize, Ordering};
#[cfg(counted)]
static RECORD_CALLS: AtomicUsize = AtomicUsize::new(0);
#[cfg(counted)]
static SCALAR_CALLS: AtomicUsize = AtomicUsize::new(0);

#[cfg(counted)]
pub fn call_counts() -> [usize; 2] {
    [RECORD_CALLS.load(Ordering::Relaxed), SCALAR_CALLS.load(Ordering::Relaxed)]
}

#[cfg(counted)]
pub fn reset_counts() {
    RECORD_CALLS.store(0, Ordering::Relaxed);
    SCALAR_CALLS.store(0, Ordering::Relaxed);
}

// Exact generated update expression; only parameter identifiers are renamed.
#[inline(never)]
pub fn record_step(n: i64, r: Value) -> Value {
    #[cfg(counted)]
    RECORD_CALLS.fetch_add(1, Ordering::Relaxed);
{
    let _record_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&r).__purust_borrow_a()).unwrap_int() + 1));
    let _record_child_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&r).__purust_borrow_b().__purust_borrow_c()).unwrap_int() + 2));
    let _record_child_1_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&r).__purust_borrow_b().__purust_borrow_d().__purust_borrow_e()).unwrap_int() + 3));
    let _record_child_1_update_1 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&r).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f()).unwrap_int() + { let _mod_l: i64 = /* Typed i64 <- i64 : Local(...) */n; let _mod_r: i64 = /* Typed i64 <- i64 : Lit */5; _mod_l.checked_rem_euclid(_mod_r).unwrap_or(0_i64) }));
    let mut _base = r;
    _base.set_a(_record_update_0);
    let mut _record_child = _base.get_b();
    _base.set_b(purust_core::Value::Unit);
    _record_child.set_c(_record_child_update_0);
    let mut _record_child_1 = _record_child.get_d();
    _record_child.set_d(purust_core::Value::Unit);
    _record_child_1.set_e(_record_child_1_update_0);
    _record_child_1.set_f(_record_child_1_update_1);
    _record_child.set_d(_record_child_1);
    _base.set_b(_record_child);
    _base
}
}

// Four scalar fields follow the same update arithmetic as record_step.
// This function is compiled in the same separate crate as record_step.
#[inline(never)]
pub fn scalar_step(n: i64, fields: [i64; 4]) -> [i64; 4] {
    #[cfg(counted)]
    SCALAR_CALLS.fetch_add(1, Ordering::Relaxed);
    [
        fields[0] + 1,
        fields[1] + 2,
        fields[2] + 3,
        fields[3] + n.checked_rem_euclid(5).unwrap_or(0_i64),
    ]
}
