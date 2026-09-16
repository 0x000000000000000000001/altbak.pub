    private static java.util.function.IntSupplier buildThunks(int depth, java.util.function.IntSupplier initial) {
        java.util.function.IntSupplier result = initial;
        for (int i = 0; i < depth; i++) {
            java.util.function.IntSupplier previous = result;
            result = () -> previous.getAsInt() + 1;
        }
        return result;
    }

    public static final java.util.function.Function<Object, Object> runLazyEvaluationFFI = Bench.nativeBenchmark(input -> {
        int count = (Integer) input;
        int result = 0;
        for (int i = 0; i < count; i++) {
            result += buildThunks(1000, () -> 0).getAsInt();
        }
        return result;
    });
