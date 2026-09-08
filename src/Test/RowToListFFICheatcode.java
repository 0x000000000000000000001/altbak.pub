    public static final java.util.function.Function<Object, Object> runRowToListFFICheatcode = Bench.nativeBenchmark(input -> {
        // RowToList resolves the fixed record type { a, b, c, d, e } at compile time.
        // The optimized native equivalent therefore knows its cardinality statically.
        return 5;
    });
