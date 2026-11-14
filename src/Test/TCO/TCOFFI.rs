pub fn Test_TCOFFI_runTCOFFI(mut limit: i64) -> i64 {
    let mut n = limit;
    let mut acc = 0;
    while n > 0 {
        acc += n % 3;
        n -= 1;
    }
    acc
}
