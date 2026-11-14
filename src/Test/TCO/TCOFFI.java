// Java has no tail-call elimination; this loop is the tail-recursive state transition.
private static int deepTailRec(int n, int acc) {
    while (n > 0) {
        acc += n % 3;
        n--;
    }
    return acc;
}

public static final java.util.function.Function<Object, Object> runTCOFFI =
    Bench.nativeBenchmark(input -> deepTailRec(((Number) input).intValue(), 0));
