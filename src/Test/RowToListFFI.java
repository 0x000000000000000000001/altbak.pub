    private record RowNil() {}
    private record RowCons<Head, Tail>(Head head, Tail tail) {}
    @FunctionalInterface
    private interface RecordKeys<Row> {
        int keysImpl(Object proxy);
    }

    private static RecordKeys<RowNil> keysNil() {
        return proxy -> 0;
    }

    private static <Head, Tail> RecordKeys<RowCons<Head, Tail>> keysCons(RecordKeys<Tail> tail) {
        return proxy -> 1 + tail.keysImpl(null);
    }

    private static <Row> int keys(RecordKeys<Row> dictionary, Row record) {
        return dictionary.keysImpl(null);
    }

    public static final java.util.function.Function<Object, Object> runRowToListFFI = Bench.nativeBenchmark(input -> {
        // The heterogeneous row type must agree with its recursive dictionary.
        var record = new RowCons<>(1, new RowCons<>("two", new RowCons<>(true,
            new RowCons<>(4.0, new RowCons<>("five", new RowNil())))));
        return keys(keysCons(keysCons(keysCons(keysCons(keysCons(keysNil()))))), record);
    });
