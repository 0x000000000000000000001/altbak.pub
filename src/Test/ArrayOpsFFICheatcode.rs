pub fn Test_ArrayOpsFFICheatcode_runArrayOpsFFICheatcode(mut limit: i64) -> i64 {
    let mut sum = 0;
    for i in 1..=limit {
        if i % 2 == 0 { sum += i; }
    }
    sum
}
