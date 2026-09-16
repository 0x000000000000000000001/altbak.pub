
pub fn Test_ListOpsFFICheatcode_runListOpsFFICheatcode(mut limit: i64) -> i64 {
    let mut sum = 0;
    for x in 1..=limit {
        if x % 2 == 0 {
            sum += x;
        }
    }
    sum
}