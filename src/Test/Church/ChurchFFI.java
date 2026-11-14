@FunctionalInterface
private interface Church<A> {
    java.util.function.Function<A, A> apply(java.util.function.Function<A, A> f);
}

private static <A> Church<A> zero() { return f -> x -> x; }
private static <A> Church<A> successor(Church<A> previous) {
    return f -> x -> f.apply(previous.apply(f).apply(x));
}

private static Church<Integer> fromInt(int n) {
    return n == 0 ? zero() : successor(fromInt(n - 1));
}

private static <A> Church<A> multiply(Church<A> m, Church<A> n) {
    return f -> x -> m.apply(n.apply(f)).apply(x);
}

private static Church<Integer> square(int n) {
    return multiply(fromInt(n), fromInt(n));
}

public static final java.util.function.Function<Object, Object> runChurchFFI = Bench.nativeBenchmark(input -> {
    int n = ((Number) input).intValue();
    Church<Integer> fourthPower = multiply(square(n), square(n));
    Church<Integer> fifthPower = multiply(fourthPower, fromInt(n));
    return fifthPower.apply(x -> x + 1).apply(0);
});
