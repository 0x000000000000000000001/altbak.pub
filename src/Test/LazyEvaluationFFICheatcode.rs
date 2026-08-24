pub fn Test_LazyEvaluationFFICheatcode_runLazyEvaluationFFICheatcode(mut limit: i64) -> i64 {
    let mut val = 0;
    for _ in 0..limit {
        val += 1000;
    }
    val
}
