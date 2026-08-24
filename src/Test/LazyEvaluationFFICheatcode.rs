pub fn Test_LazyEvaluationFFICheatcode_runLazyEvaluationFFICheatcode(mut limit: i64) -> i64 {
    let mut val = 1;
    for _ in 0..limit {
        val += 1;
    }
    val
}