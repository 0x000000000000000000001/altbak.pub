    private record Monoidish<A>(A mempty, java.util.function.Function<A, java.util.function.Function<A, A>> mappend) {}

    private static <A> A polyLoop(Monoidish<A> dictionary, int count, A initial) {
        A result = initial;
        while (count > 0) {
            result = dictionary.mappend().apply(result).apply(dictionary.mempty());
            count--;
        }
        return result;
    }

    public static final java.util.function.Function<Object, Object> runPolymorphismFFI = Bench.nativeBenchmark(input ->
        polyLoop(new Monoidish<>(1, left -> right -> left + right), (Integer) input, 0));
