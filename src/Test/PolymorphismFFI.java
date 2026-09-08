    private record Monoidish(int mempty, java.util.function.IntFunction<java.util.function.IntUnaryOperator> mappend) {}

    private static int polyLoop(Monoidish dictionary, int count, int initial) {
        int result = initial;
        while (count > 0) {
            result = dictionary.mappend().apply(result).applyAsInt(dictionary.mempty());
            count--;
        }
        return result;
    }

    public static final java.util.function.Function<Object, Object> runPolymorphismFFI = Bench.nativeBenchmark(input ->
        polyLoop(new Monoidish(1, left -> right -> left + right), (Integer) input, 0));
