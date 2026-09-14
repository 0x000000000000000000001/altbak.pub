#![allow(dead_code, unused_imports, unused_mut, non_snake_case)]
extern crate self as purust_core;
#[derive(Clone)]
pub enum Func1<T> { Shared(std::rc::Rc<dyn Fn(T) -> T>) }
#[derive(Clone)]
pub enum Value { Number(f64), Func1(Func1<Value>) }
pub type UnknownType = Value;
pub fn mk_number(n: f64) -> Value { Value::Number(n) }
pub fn mk_string(_: String) -> Value { Value::Number(0.0) }
mod bench { include!("BENCH_PATH"); }
fn main() {
    let Value::Func1(Func1::Shared(now)) = bench::Bench_benchNow() else { panic!() };
    let mut previous = 0.0;
    let mut fractional = false;
    for _ in 0..100 {
        let Value::Number(value) = now(Value::Number(0.0)) else { panic!() };
        assert!(value.is_finite() && value >= previous);
        fractional |= value.fract() != 0.0;
        previous = value;
    }
    assert!(fractional, "clock truncated all sub-microsecond values");
    let Value::Func1(Func1::Shared(opaque)) = bench::Bench_opaque(Value::Number(123.0)) else { panic!() };
    let Value::Number(value) = opaque(Value::Number(0.0)) else { panic!() };
    assert_eq!(value, 123.0);
    assert_eq!(bench::Bench_formatNumber(0.125), "0.12");
    println!("PASS Rust monotonic fractional clock and opaque barrier contract");
}
