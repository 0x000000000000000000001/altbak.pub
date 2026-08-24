pub fn Test_PolymorphismFFICheatcode_runPolymorphismFFICheatcode(mut limit: i64) -> i64 {
    let a = vec![vec![1, 2], vec![3, 4]];
    let b = vec![vec![1, 2], vec![3, 4]];
    let mut count = 0;
    for _ in 0..limit {
        if a == b { count += 1; }
    }
    count
}