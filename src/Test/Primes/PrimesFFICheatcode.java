public static final java.util.function.Function<Object, Object> runPrimesFFICheatcode = Bench.nativeBenchmark(input -> {
    int n = ((Number) input).intValue();
    if (n < 2) return 0;
    boolean[] composite = new boolean[n + 1];
    for (int p = 2; p <= n / p; p++) {
        if (!composite[p]) {
            for (int multiple = p * p; multiple <= n; multiple += p) {
                composite[multiple] = true;
            }
        }
    }
    int sum = 0;
    for (int p = 2; p <= n; p++) {
        if (!composite[p]) sum += p;
    }
    return sum;
});
