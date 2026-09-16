    public static final java.util.function.Function<Object, Object> runArrayOpsFFICheatcode = Bench.nativeBenchmark(input -> {
        int end = (Integer) input;
        int direction = end >= 1 ? 1 : -1;
        int result = 0;
        int value = 1;
        while (true) {
            if (value % 2 == 0) result += value;
            if (value == end) break;
            value += direction;
        }
        return result;
    });
