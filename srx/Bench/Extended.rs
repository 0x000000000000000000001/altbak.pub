pub fn Bench_Extended_consumeResult(result: String) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        std::hint::black_box(&result);
        crate::Value::Unit
    })))
}
