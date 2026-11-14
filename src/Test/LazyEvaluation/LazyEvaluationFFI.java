    @FunctionalInterface
    private interface Lazy<A> { A force(); }

    private static <A> Lazy<A> defer(java.util.function.Supplier<A> thunk) {
        return () -> thunk.get();
    }

    private static <A> A force(Lazy<A> value) { return value.force(); }

    private static Lazy<Integer> buildThunks(int depth, Lazy<Integer> initial) {
        Lazy<Integer> result = initial;
        for (int i = 0; i < depth; i++) {
            Lazy<Integer> previous = result;
            result = defer(() -> force(previous) + 1);
        }
        return result;
    }

    public static final java.util.function.Function<Object, Object> runLazyEvaluationFFI = Bench.nativeBenchmark(input -> {
        int count = (Integer) input;
        int result = 0;
        for (int i = 0; i < count; i++) {
            result += force(buildThunks(1000, defer(() -> 0)));
        }
        return result;
    });
