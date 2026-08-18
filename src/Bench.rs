

pub fn Bench_benchNow() -> UnknownType {
    let start = std::time::SystemTime::now();
    let since_the_epoch = start
        .duration_since(std::time::UNIX_EPOCH)
        .expect("Time went backwards");
    let micros = since_the_epoch.as_micros() as f64;
    mk_number(micros)
}

pub fn Bench_opaque(mut a0: UnknownType) -> UnknownType {
    std::hint::black_box(a0)
}

pub fn Bench_formatNumber(mut nObj: UnknownType) -> UnknownType {
    let n = nObj.init_number.unwrap();
    mk_string(&format!("{:.2}", n))
}
