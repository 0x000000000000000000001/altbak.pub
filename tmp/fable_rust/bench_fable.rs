// Append this harness after Fable's generated Runner.rs at the crate root.
// Runner.rs has inner attributes and must not be expanded with include!.
// The build driver installs mimalloc, matching the Purust reference harness.
mod altbak_fable_benchmark {
    use super::BenchFable;
    use std::hint::black_box;
    use std::time::{Duration, Instant};

    struct Case {
        key: &'static str,
        name: &'static str,
        act: fn(i32) -> i32,
        input: i32,
        expected: i32,
    }

    const CASES: [Case; 14] = [
        Case { key: "AstTree", name: "AST Evaluation:", act: BenchFable::runAstTree, input: 3, expected: 7 },
        Case { key: "Fib", name: "Fibonacci:", act: BenchFable::runFib, input: 10, expected: 55 },
        Case { key: "ListOps", name: "List Processing (900 elements):", act: BenchFable::runListOps, input: 900, expected: 202950 },
        Case { key: "TCO", name: "Tail Call Optimization (100k calls):", act: BenchFable::runTCO, input: 100000, expected: 100000 },
        Case { key: "Records", name: "Deep Record Updates (10k iterations):", act: BenchFable::runRecords, input: 10000, expected: 20000 },
        Case { key: "Ackermann", name: "Ackermann (3, 4):", act: BenchFable::runAckermann, input: 3, expected: 125 },
        Case { key: "Church", name: "Church Numerals (100k Closure Applications):", act: BenchFable::runChurch, input: 10, expected: 100000 },
        Case { key: "Primes", name: "Prime Sieve (sum primes up to 500):", act: BenchFable::runPrimes, input: 500, expected: 21536 },
        Case { key: "RBTree", name: "Red-Black Tree (100k Worst-Case Insertions):", act: BenchFable::runRBTree, input: 100000, expected: 22 },
        Case { key: "Polymorphism", name: "Polymorphism (10M Type Class Dict Lookups):", act: BenchFable::runPolymorphism, input: 10000000, expected: 10000000 },
        Case { key: "StateMonad", name: "State Monad (1.2k Binds, 60 Stack Depth):", act: BenchFable::runStateMonad, input: 20, expected: 1200 },
        Case { key: "LazyEvaluation", name: "Lazy Evaluation (1M Thunks Forced, 1k Depth):", act: BenchFable::runLazyEvaluation, input: 1000, expected: 1000000 },
        Case { key: "ArrayOps", name: "Array Processing (900 elements):", act: BenchFable::runArrayOps, input: 900, expected: 202950 },
        Case { key: "RowToList", name: "RowToList (Keys Count):", act: BenchFable::runRowToList, input: 10000, expected: 5 },
    ];

    fn check_result(case: &Case, result: i32) {
        assert_eq!(result, case.expected, "Incorrect result: {}", case.name);
    }

    fn warmup(case: &Case) -> i32 {
        let result = black_box((black_box(case.act))(black_box(case.input)));
        check_result(case, result);
        result
    }

    #[inline(never)]
    fn batch(act: fn(i32) -> i32, input: i32, iterations: u32) -> (Duration, i32) {
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

    fn usage() {
        eprintln!("Usage: bench_fable [--case NAME] [--check-only]");
        eprintln!("Cases: {}", CASES.iter().map(|case| case.key).collect::<Vec<_>>().join(", "));
    }

    fn parse_args() -> Result<(Option<&'static Case>, bool), String> {
        let mut selected = None;
        let mut check_only = false;
        let mut args = std::env::args().skip(1);
        while let Some(arg) = args.next() {
            match arg.as_str() {
                "--case" if selected.is_none() => {
                    let key = args.next().ok_or("--case requires a case name")?;
                    selected = Some(CASES.iter().find(|case| case.key == key)
                        .ok_or_else(|| format!("Unknown case: {key}"))?);
                }
                "--check-only" if !check_only => check_only = true,
                _ => return Err(format!("Unknown or repeated argument: {arg}")),
            }
        }
        Ok((selected, check_only))
    }

    pub fn main() {
        let (selected, check_only) = parse_args().unwrap_or_else(|error| {
            eprintln!("{error}");
            usage();
            std::process::exit(2);
        });
        let cases: Vec<&Case> = match selected {
            Some(case) => vec![case],
            None => CASES.iter().collect(),
        };
        if check_only {
            for case in &cases {
                println!("{} {}", case.name, warmup(case));
            }
            return;
        }

        println!("Global warm-up in progress...");
        for _ in 0..3 {
            for case in &cases {
                warmup(case);
            }
        }
        let total: f64 = cases.into_iter().map(bench).sum();
        println!("\n==================================================\n\nTotal exec time: {:.6} ms\n", total / 1000.0);
    }
}

fn main() {
    std::thread::Builder::new().name("benchmark".into())
        .stack_size(1024 * 1024 * 1024)
        .spawn(altbak_fable_benchmark::main).expect("benchmark thread")
        .join().expect("benchmark completed");
}
