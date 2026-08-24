pub fn Test_ArrayOpsFFICheatcode_runArrayOpsFFICheatcode(mut limit: i64) -> i64 {
    let mut sum = 0;
    for i in 0..limit {
        sum += i * 2;
    }
    sum
}