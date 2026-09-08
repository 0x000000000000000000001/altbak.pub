private static int fib(int n) {
    if (n == 0) return 0;
    if (n == 1) return 1;
    return fib(n - 1) + fib(n - 2);
}

public static final java.util.function.Function<Object, Object> runFibFFI =
    Bench.nativeBenchmark(input -> fib(((Number) input).intValue()));
