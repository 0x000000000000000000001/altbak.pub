
pub fn Test_RecordsFFICheatcode_runRecordsFFICheatcode(mut limit: i64) -> i64 {
    let mut x = 0;
    let mut y = 0;
    let mut z = 0;
    for _ in 0..limit {
        x += 1;
        y += 2;
        z += 3;
    }
    x + y + z
}