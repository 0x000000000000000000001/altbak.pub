use crate::{mk_number, mk_string, UnknownType};
use std::rc::Rc;

pub fn Bench_benchNow() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(
        move |mut _u: UnknownType| -> UnknownType {
            static EPOCH: std::sync::OnceLock<std::time::Instant> = std::sync::OnceLock::new();
            let elapsed = EPOCH.get_or_init(std::time::Instant::now).elapsed();
            mk_number(elapsed.as_secs_f64() * 1_000_000.0)
        },
    )))
}

pub fn Bench_opaque(mut a0: UnknownType) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(
        move |mut _u: UnknownType| -> UnknownType { std::hint::black_box(a0.clone()) },
    )))
}

pub fn Bench_formatNumber(mut n: f64) -> String {
    format!("{:.6}", n)
}

pub fn Bench_measureBatch(iterations: i64, expected: i64, action: UnknownType) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |_u: UnknownType| {
        let act = std::hint::black_box(action.clone()).unwrap_func1();
        let mut result = 0;
        let start = std::time::Instant::now();
        for _ in 0..iterations {
            result = std::hint::black_box(act(crate::Value::Unit)).unwrap_int();
        }
        let elapsed = start.elapsed().as_secs_f64() * 1_000_000.0;
        assert_eq!(result, expected, "Unstable benchmark result");
        mk_number(elapsed)
    })))
}
