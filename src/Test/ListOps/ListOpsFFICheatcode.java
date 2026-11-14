public static final java.util.function.Function<Object, Object> runListOpsFFICheatcode = Bench.nativeBenchmark(input -> {
    int limit = ((Number) input).intValue();
    int sum = 0;
    for (int value = 1; value <= limit; value++) {
        if (value % 2 == 0) sum += value;
    }
    return sum;
});
