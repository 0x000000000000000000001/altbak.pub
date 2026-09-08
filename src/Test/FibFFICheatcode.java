// Keep the recursive algorithm used by the native Fibonacci baselines.
private static int fib(int n) {
    if (n <= 1) return n;
    return fib(n - 1) + fib(n - 2);
}

public static final java.util.function.Function<Object, Object> runFibFFICheatcode =
    Bench.nativeBenchmark(input -> fib(((Number) input).intValue()));
