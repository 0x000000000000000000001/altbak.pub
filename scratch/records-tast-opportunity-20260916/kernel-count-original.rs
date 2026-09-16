pub fn Test_Records_updateRec(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
    loop {
        break /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if (/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == 0) {
        /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1
    } else {
        {
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed crate::UnknownType <- crate::UnknownType : Update */{
 OP_READS.fetch_add(1,Ordering::Relaxed);
    let _record_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_a()).unwrap_int() + 1));
 OP_READS.fetch_add(2,Ordering::Relaxed);
    let _record_child_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_c()).unwrap_int() + 2));
 OP_READS.fetch_add(3,Ordering::Relaxed);
    let _record_child_1_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_e()).unwrap_int() + 3));
 OP_READS.fetch_add(3,Ordering::Relaxed);
    let _record_child_1_update_1 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f()).unwrap_int() + { let _mod_l: i64 = /* Typed i64 <- i64 : Local(...) */purs_local_0; let _mod_r: i64 = /* Typed i64 <- i64 : Lit */5; _mod_l.checked_rem_euclid(_mod_r).unwrap_or(0_i64) }));
    let mut _base = purs_local_1;
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    _base.set_a(_record_update_0);
 OP_GETS.fetch_add(1,Ordering::Relaxed);
    let mut _record_child = _base.get_b();
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    _base.set_b(purust_core::Value::Unit);
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    _record_child.set_c(_record_child_update_0);
 OP_GETS.fetch_add(1,Ordering::Relaxed);
    let mut _record_child_1 = _record_child.get_d();
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    _record_child.set_d(purust_core::Value::Unit);
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    _record_child_1.set_e(_record_child_1_update_0);
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    _record_child_1.set_f(_record_child_1_update_1);
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    _record_child.set_d(_record_child_1);
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    _base.set_b(_record_child);
    _base
};
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        continue;
    }
    };
    }
}

