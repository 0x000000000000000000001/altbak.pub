
pub fn Test_TCOFFICheatcode_runTCOFFICheatcode(mut limit: i64) -> i64 {
    let mut acc = 0;
    let mut n = limit;
    while n > 0 {
        acc += n;
        n -= 1;
    }
    acc
}