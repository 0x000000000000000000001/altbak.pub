    private interface Monoidish {
        int mempty();
        int mappend(int left, int right);
    }

    private static final class IntMonoidish implements Monoidish {
        public int mempty() { return 1; }
        public int mappend(int left, int right) { return left + right; }
    }

    public static final java.util.function.Function<Object, Object> runPolymorphismFFICheatcode = Bench.nativeBenchmark(input -> {
        int count = (Integer) input;
        int result = 0;
        Monoidish dictionary = new IntMonoidish();
        for (int i = 0; i < count; i++) {
            result = dictionary.mappend(result, dictionary.mempty());
        }
        return result;
    });
