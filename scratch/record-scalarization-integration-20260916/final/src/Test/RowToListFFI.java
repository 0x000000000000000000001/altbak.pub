    @FunctionalInterface
    private interface RecordKeys {
        int keysImpl(Object proxy);
    }

    private static RecordKeys keysCons(RecordKeys tail) {
        return proxy -> 1 + tail.keysImpl(null);
    }

    public static final java.util.function.Function<Object, Object> runRowToListFFI = Bench.nativeBenchmark(input -> {
        // The record's type has five fields; values and the opaque input are unused by keysImpl.
        RecordKeys dictionary = proxy -> 0;
        dictionary = keysCons(keysCons(keysCons(keysCons(keysCons(dictionary)))));
        return dictionary.keysImpl(null);
    });
