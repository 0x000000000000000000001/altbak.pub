    private static int[] rangeFromOne(int end) {
        int size = Math.toIntExact(Math.abs((long) end - 1) + 1);
        int[] result = new int[size];
        int direction = end >= 1 ? 1 : -1;
        for (int i = 0; i < size; i++) result[i] = 1 + direction * i;
        return result;
    }

    public static final java.util.function.Function<Object, Object> runArrayOpsFFI = Bench.nativeBenchmark(input -> {
        int[] values = rangeFromOne((Integer) input);
        int[] evens = java.util.Arrays.stream(values).filter(value -> value % 2 == 0).toArray();
        return java.util.Arrays.stream(evens).reduce(0, (sum, value) -> sum + value);
    });
