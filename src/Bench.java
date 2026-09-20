public static final java.util.function.Supplier<Object> benchNow = () -> {
    return (double) (System.nanoTime() / 1000.0);
};

private static volatile Object benchmarkInput;
private static volatile Object benchmarkResult;
public static final java.util.function.Function<Object, Object> opaque = (a) ->
    (java.util.function.Supplier<Object>) () -> {
        benchmarkInput = a;
        return benchmarkInput;
    };

public static final java.util.function.Function<Object, Object> formatNumber = (nObj) -> {
    double n = ((Number) nObj).doubleValue();
    return String.format(java.util.Locale.US, "%.6f", n);
};

// The common numeric harness protects compiled and FFI kernels equally.
public static java.util.function.Function<Object, Object> nativeBenchmark(
        java.util.function.Function<Object, Object> action) {
    return action;
}

public static final java.util.function.Function<Object, Object> measureBatch = iterations ->
    (java.util.function.Function<Object, Object>) expected ->
    (java.util.function.Function<Object, Object>) action ->
    (java.util.function.Supplier<Object>) () -> {
        java.util.function.Supplier<?> act = (java.util.function.Supplier<?>) action;
        int count = ((Number) iterations).intValue();
        Object result = null;
        long start = System.nanoTime();
        for (int i = 0; i < count; i++) {
            result = act.get();
            benchmarkResult = result;
        }
        double elapsed = (System.nanoTime() - start) / 1000.0;
        if (((Number) result).intValue() != ((Number) expected).intValue())
            throw new IllegalStateException("Unstable benchmark result");
        return elapsed / count;
    };
