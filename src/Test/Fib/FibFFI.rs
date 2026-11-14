
pub fn Test_FibFFI_runFibFFI(mut limit: i64) -> i64 {
    fn fib(n: i64) -> i64 {
        if n == 0 { 0 } else if n == 1 { 1 } else { fib(n - 1) + fib(n - 2) }
    }
    fib(limit)
}