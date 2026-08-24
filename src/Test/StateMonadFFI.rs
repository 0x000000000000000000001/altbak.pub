pub fn Test_StateMonadFFI_runStateMonadFFI(mut limit: i64) -> i64 {
    let mut s = 0;
    for _ in 0..(limit * 20) {
        s += 1;
    }
    s
}
