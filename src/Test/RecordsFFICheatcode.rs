pub fn Test_RecordsFFICheatcode_runRecordsFFICheatcode(mut limit: i64) -> i64 {
    let mut a = 0;
    let mut c = 0;
    let mut e = 0;
    let mut f = 0;
    while limit > 0 {
        a += 1;
        c += 2;
        e += 3;
        f += limit % 5;
        limit -= 1;
    }
    f
}
