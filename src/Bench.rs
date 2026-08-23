use std::rc::Rc;
use crate::{mk_number, mk_string, UnknownType, Record_a};

pub fn Bench_benchNow() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
        let start = std::time::SystemTime::now();
        let since_the_epoch = start
            .duration_since(std::time::UNIX_EPOCH)
            .expect("Time went backwards");
        let micros = since_the_epoch.as_micros() as f64;
        mk_number(micros)
    })))
}

pub fn Bench_opaque(mut a0: UnknownType) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
        std::hint::black_box(a0.clone())
    })))
}

pub fn Bench_formatNumber(mut n: f64) -> String {
    format!("{:.2}", n)
}
