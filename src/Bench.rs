use std::rc::Rc;
use crate::{mk_number, mk_string, UnknownType, Record_a};

pub fn Bench_benchNow() -> UnknownType {
    perceus_ptr::PerceusPtr::new(Record_a {
        call: Some(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            let start = std::time::SystemTime::now();
            let since_the_epoch = start
                .duration_since(std::time::UNIX_EPOCH)
                .expect("Time went backwards");
            let micros = since_the_epoch.as_micros() as f64;
            mk_number(micros)
        })),
        ..Default::default()
    })
}

pub fn Bench_opaque(mut a0: UnknownType) -> UnknownType {
    perceus_ptr::PerceusPtr::new(Record_a {
        call: Some(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            std::hint::black_box(a0.clone())
        })),
        ..Default::default()
    })
}

pub fn Bench_formatNumber(mut nObj: UnknownType) -> UnknownType {
    let n = nObj.init_number.unwrap();
    mk_string(&format!("{:.2}", n))
}
