pub fn Test_StateMonadFFICheatcode_runStateMonadFFICheatcode(mut limit: i64) -> i64 {
    let mut s = 0;
    for _ in 0..(limit * 20) {
        s += 1;
    }
    s
}
