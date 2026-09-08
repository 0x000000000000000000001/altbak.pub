@FunctionalInterface
private interface Church {
    java.util.function.IntUnaryOperator apply(java.util.function.IntUnaryOperator f);
}

private static Church fromInt(int n) {
    if (n == 0) return f -> x -> x;
    Church previous = fromInt(n - 1);
    return f -> x -> f.applyAsInt(previous.apply(f).applyAsInt(x));
}

private static Church multiply(Church m, Church n) {
    return f -> x -> m.apply(n.apply(f)).applyAsInt(x);
}

private static Church square(int n) {
    return multiply(fromInt(n), fromInt(n));
}

public static final java.util.function.Function<Object, Object> runChurchFFI = Bench.nativeBenchmark(input -> {
    int n = ((Number) input).intValue();
    Church fourthPower = multiply(square(n), square(n));
    Church fifthPower = multiply(fourthPower, fromInt(n));
    return fifthPower.apply(x -> x + 1).applyAsInt(0);
});
