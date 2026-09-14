pub fn Test_ArrayOpsFFICheatcode_runArrayOpsFFICheatcode(mut limit: i64) -> i64 {
    let mut sum = 0;
    let step = if limit >= 1 { 1 } else { -1 };
    let mut i = 1;
    loop {
        if i % 2 == 0 {
            sum += i;
        }
        if i == limit {
            break;
        }
        i += step;
    }
    sum
}
