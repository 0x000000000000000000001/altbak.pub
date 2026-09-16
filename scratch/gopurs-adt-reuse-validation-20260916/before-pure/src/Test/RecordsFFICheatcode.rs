pub fn Test_RecordsFFICheatcode_runRecordsFFICheatcode(mut limit: i64) -> i64 {
    let mut f = 0;
    while limit > 0 {
        f += limit % 5;
        limit -= 1;
    }
    f
}
