pub fn Bench_Extended_consumeResult(expected: String, result: String) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        std::hint::black_box(&result);
        assert_eq!(result, expected, "Unstable extended benchmark result");
        crate::Value::Unit
    })))
}
