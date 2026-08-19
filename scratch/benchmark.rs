use std::rc::Rc;
use std::time::Instant;

// Simulate purust's current approach
pub type UnknownType = Box<RecordA>; // Using Box instead of PerceusPtr for simplicity, same dynamic allocation cost

#[derive(Clone, Default)]
pub struct RecordA {
    pub tag: &'static str,
    pub call: Option<Rc<dyn Fn(UnknownType) -> UnknownType>>,
    pub init_number: Option<f64>,
    pub field1: Option<UnknownType>,
    pub field2: Option<UnknownType>,
    pub field3: Option<UnknownType>,
    pub field4: Option<UnknownType>,
    pub field5: Option<UnknownType>,
    pub field6: Option<UnknownType>,
    pub field7: Option<UnknownType>,
    pub field8: Option<UnknownType>,
    pub field9: Option<UnknownType>,
    pub field10: Option<UnknownType>,
    // padding out to a large size to simulate the 278 fields
    pub p1: [u64; 32],
    pub p2: [u64; 32],
    pub p3: [u64; 32],
    pub p4: [u64; 32],
    pub p5: [u64; 32],
    pub p6: [u64; 32],
    pub p7: [u64; 32],
    pub p8: [u64; 26],
}

pub fn mk_number_dynamic(val: f64) -> UnknownType {
    Box::new(RecordA {
        init_number: Some(val),
        ..Default::default()
    })
}

// Simulate a function call in purust
pub fn call_dynamic(f: UnknownType, arg: UnknownType) -> UnknownType {
    (f.call.as_ref().unwrap())(arg)
}

fn bench_dynamic_fib(n: f64) -> UnknownType {
    if n <= 1.0 {
        return mk_number_dynamic(1.0);
    }
    
    // In purust, everything is an UnknownType and mathematical operations unbox and rebox
    let a = bench_dynamic_fib(n - 1.0);
    let b = bench_dynamic_fib(n - 2.0);
    
    mk_number_dynamic(a.init_number.unwrap() + b.init_number.unwrap())
}


// Proposed approach: Native Unboxed Types
fn bench_native_fib(n: f64) -> f64 {
    if n <= 1.0 {
        return 1.0;
    }
    bench_native_fib(n - 1.0) + bench_native_fib(n - 2.0)
}


fn main() {
    let fib_n = 30.0; // Reasonable size to see difference
    
    println!("Benchmarking Dynamic 'God Struct' vs Native Typed approach...");
    
    // 1. Benchmark Dynamic Approach
    let start_dyn = Instant::now();
    let res_dyn = bench_dynamic_fib(fib_n);
    let dur_dyn = start_dyn.elapsed();
    println!("[Dynamic/Record_a] Fib({}): {} in {:?}", fib_n, res_dyn.init_number.unwrap(), dur_dyn);
    
    // 2. Benchmark Native Approach
    let start_nat = Instant::now();
    let res_nat = bench_native_fib(fib_n);
    let dur_nat = start_nat.elapsed();
    println!("[Native/Typed]      Fib({}): {} in {:?}", fib_n, res_nat, dur_nat);
    
    let ratio = dur_dyn.as_nanos() as f64 / dur_nat.as_nanos() as f64;
    println!("Native is {:.2}x faster for recursive calls.", ratio);
}
