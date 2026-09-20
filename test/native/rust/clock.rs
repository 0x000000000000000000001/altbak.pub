#![allow(dead_code, unused_imports, unused_mut, non_snake_case)]
extern crate self as purust_core;
#[derive(Clone)]
pub enum Func1<T> { Shared(std::rc::Rc<dyn Fn(T) -> T>) }
#[derive(Clone)]
pub enum Value { Unit, Int(i64), Number(f64), Func1(Func1<Value>) }
impl Value {
    pub fn unwrap_func1(&self) -> std::rc::Rc<dyn Fn(Value) -> Value> {
        match self { Self::Func1(Func1::Shared(f)) => f.clone(), _ => panic!("expected Effect closure") }
    }
    pub fn unwrap_int(&self) -> i64 {
        match self { Self::Int(value) => *value, _ => panic!("expected integer result") }
    }
}
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
    assert_eq!(bench::Bench_formatNumber(0.125), "0.125000");
    let calls = std::rc::Rc::new(std::cell::Cell::new(0));
    let counter = calls.clone();
    let action = Value::Func1(Func1::Shared(std::rc::Rc::new(move |unit| {
        assert!(matches!(unit, Value::Unit));
        counter.set(counter.get() + 1);
        Value::Int(7)
    })));
    let batch = bench::Bench_measureBatch(3, 7, action);
    assert_eq!(calls.get(), 0, "constructing a batch must not execute the Effect");
    let Value::Number(elapsed) = batch.unwrap_func1()(Value::Unit) else { panic!() };
    assert!(elapsed.is_finite() && elapsed >= 0.0);
    assert_eq!(calls.get(), 3, "a batch must execute exactly its requested calls");
    batch.unwrap_func1()(Value::Unit);
    assert_eq!(calls.get(), 6, "reusing an Effect must rerun all calls");
    println!("PASS Rust monotonic fractional clock, numeric Effect batch and opaque barrier contract");
}
