use std::hint::black_box;
use std::time::{Duration, Instant};

use Purs_BenchPurust::*;

#[global_allocator]
static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;

struct Case {
    name: &'static str,
    act: fn(i64) -> i64,
    input: i64,
    expected: i64,
}

const CASES: [Case; 14] = [
    Case { name: "AST Evaluation:", act: BenchPurust_runAstTree, input: 3, expected: 7 },
    Case { name: "Fibonacci:", act: BenchPurust_runFib, input: 10, expected: 55 },
    Case { name: "List Processing (900 elements):", act: BenchPurust_runListOps, input: 900, expected: 202950 },
    Case { name: "Tail Call Optimization (100k calls):", act: BenchPurust_runTCO, input: 100000, expected: 100000 },
    Case { name: "Deep Record Updates (10k iterations):", act: BenchPurust_runRecords, input: 10000, expected: 20000 },
    Case { name: "Ackermann (3, 4):", act: BenchPurust_runAckermann, input: 3, expected: 125 },
    Case { name: "Church Numerals (100k Closure Applications):", act: BenchPurust_runChurch, input: 10, expected: 100000 },
    Case { name: "Prime Sieve (sum primes up to 500):", act: BenchPurust_runPrimes, input: 500, expected: 21536 },
    Case { name: "Red-Black Tree (100k Worst-Case Insertions):", act: BenchPurust_runRBTree, input: 100000, expected: 22 },
    Case { name: "Polymorphism (10M Type Class Dict Lookups):", act: BenchPurust_runPolymorphism, input: 10000000, expected: 10000000 },
    Case { name: "State Monad (1.2k Binds, 60 Stack Depth):", act: BenchPurust_runStateMonad, input: 20, expected: 1200 },
    Case { name: "Lazy Evaluation (1M Thunks Forced, 1k Depth):", act: BenchPurust_runLazyEvaluation, input: 1000, expected: 1000000 },
    Case { name: "Array Processing (900 elements):", act: BenchPurust_runArrayOps, input: 900, expected: 202950 },
    Case { name: "RowToList (Keys Count):", act: BenchPurust_runRowToList, input: 10000, expected: 5 },
];

fn check_result(case: &Case, result: i64) {
    assert_eq!(result, case.expected, "Incorrect result: {}", case.name);
}

fn warmup(case: &Case) -> i64 {
    let result = black_box((black_box(case.act))(black_box(case.input)));
    check_result(case, result);
    result
}

#[inline(never)]
fn batch(act: fn(i64) -> i64, input: i64, iterations: u32) -> (Duration, i64) {
    let act = black_box(act);
    let mut result = 0;
    let start = Instant::now();
    for _ in 0..iterations {
        result = black_box(act(black_box(input)));
    }
    (start.elapsed(), result)
}

fn bench(case: &Case) -> f64 {
    println!("--------------------------------------------------\n\n(Test)\n{}\n\n(Output & Warm-up)", case.name);
    println!("{}", warmup(case));
    warmup(case);
    warmup(case);

    let mut iterations = 1;
    loop {
        let (elapsed, result) = batch(case.act, case.input, iterations);
        check_result(case, result);
        if elapsed >= Duration::from_millis(10) || iterations >= 16777216 {
            break;
        }
        iterations *= 2;
    }

    let mut best = Duration::MAX;
    for _ in 0..10 {
        let (elapsed, result) = batch(case.act, case.input, iterations);
        check_result(case, result);
        best = best.min(elapsed);
    }
    let us = best.as_secs_f64() * 1_000_000.0 / f64::from(iterations);
    println!("\n(Execution time - best of 10)\n\n{us:.6} μs\n\nBatch iterations: {iterations}\n");
    us
}

fn main() {
    let args: Vec<_> = std::env::args().skip(1).collect();
    if args == ["--check-only"] {
        for case in &CASES {
            println!("{} {}", case.name, warmup(case));
        }
        return;
    }
    if !args.is_empty() {
        eprintln!("Usage: bench_purust [--check-only]");
        std::process::exit(2);
    }

    println!("Global warm-up in progress...");
    for _ in 0..3 {
        for case in &CASES {
            warmup(case);
        }
    }
    let total: f64 = CASES.iter().map(bench).sum();
    println!("\n==================================================\n\nTotal exec time: {:.6} ms\n", total / 1000.0);
}
