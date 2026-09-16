#[inline(never)]
pub fn scalar_replacement(mut n: i64, mut r: Value) -> Value {
    if n == 0 { return r; }
 OP_READS.fetch_add(1,Ordering::Relaxed);
    let mut a = r.__purust_borrow_a().unwrap_int();
 OP_READS.fetch_add(2,Ordering::Relaxed);
    let mut c = r.__purust_borrow_b().__purust_borrow_c().unwrap_int();
 OP_READS.fetch_add(3,Ordering::Relaxed);
    let mut e = r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_e().unwrap_int();
 OP_READS.fetch_add(3,Ordering::Relaxed);
    let mut f = r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_f().unwrap_int();
    while n != 0 {
        a += 1; c += 2; e += 3;
        f += n.checked_rem_euclid(5).unwrap_or(0);
        n -= 1;
    }
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    r.set_a(mk_int(a));
 OP_GETS.fetch_add(1,Ordering::Relaxed);
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    let mut b = r.get_b(); r.set_b(Value::Unit);
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    b.set_c(mk_int(c));
 OP_GETS.fetch_add(1,Ordering::Relaxed);
 OP_SETS.fetch_add(1,Ordering::Relaxed);
    let mut d = b.get_d(); b.set_d(Value::Unit);
 OP_SETS.fetch_add(2,Ordering::Relaxed);
    d.set_e(mk_int(e)); d.set_f(mk_int(f));
 OP_SETS.fetch_add(2,Ordering::Relaxed);
    b.set_d(d); r.set_b(b); r
}
