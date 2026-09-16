public static final java.util.function.Function<Object, Object> runTCOFFICheatcode = Bench.nativeBenchmark(input -> {
    int n = ((Number) input).intValue();
    int acc = 0;
    while (n > 0) {
        acc += n % 3;
        n--;
    }
    return acc;
});
