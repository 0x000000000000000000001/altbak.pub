    public static final java.util.function.Function<Object, Object> runLazyEvaluationFFICheatcode = Bench.nativeBenchmark(input -> {
        int count = (Integer) input;
        int result = 0;
        for (int i = 0; i < count; i++) {
            for (int depth = 0; depth < 1000; depth++) result++;
        }
        return result;
    });
