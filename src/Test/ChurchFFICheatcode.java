public static final java.util.function.Function<Object, Object> runChurchFFICheatcode = Bench.nativeBenchmark(input -> {
    int n = ((Number) input).intValue();
    int count = n * n * n * n * n;
    int result = 0;
    for (int i = 0; i < count; i++) result++;
    return result;
});
