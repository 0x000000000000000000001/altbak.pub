public static final java.util.function.Supplier<Object> benchNow = () -> {
    return (double) (System.nanoTime() / 1000.0);
};

public static final java.util.function.Function<Object, Object> opaque = (a) -> {
    return (java.util.function.Supplier<Object>) () -> a;
};

public static final java.util.function.Function<Object, Object> formatNumber = (nObj) -> {
    double n = ((Number) nObj).doubleValue();
    return String.format(java.util.Locale.US, "%.2f", n);
};
