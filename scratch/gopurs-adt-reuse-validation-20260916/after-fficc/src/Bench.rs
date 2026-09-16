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
    format!("{:.2}", n)
}
