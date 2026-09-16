pub fn Test_ChurchFFICheatcode_runChurchFFICheatcode(mut limit: i64) -> i64 {
    let mut acc = 0;
    for _ in 0..(limit * limit * limit * limit * limit) {
        acc += 1;
    }
    acc
}