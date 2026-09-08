    public static final java.util.function.Function<Object, Object> runStateMonadFFICheatcode = Bench.nativeBenchmark(input -> {
        // The shared FFI wrapper supplies depth 60; keep the PureScript test's 20 repetitions.
        int depth = (Integer) input;
        int result = 0;
        for (int i = 0; i < 20; i++) {
            int state = 0;
            for (int j = 0; j < depth; j++) state++;
            result += state;
        }
        return result;
    });
