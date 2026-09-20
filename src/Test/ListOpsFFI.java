private sealed interface List<A> permits Nil, Cons {}
private record Nil<A>() implements List<A> {}
private record Cons<A>(A value, List<A> tail) implements List<A> {}

private static List<Integer> range(int start, int end) {
    List<Integer> result = new Nil<>();
    for (int current = end; current >= start; current--) {
        result = new Cons<>(current, result);
    }
    return result;
}

private static List<Integer> filterEvens(List<Integer> list) {
    List<Integer> result = new Nil<>();
    while (list instanceof Cons<Integer> cons) {
        if (cons.value() % 2 == 0) result = new Cons<>(cons.value(), result);
        list = cons.tail();
    }
    return result;
}

private static <A, B> B foldl(java.util.function.BiFunction<B, A, B> operation, B acc, List<A> list) {
    while (list instanceof Cons<A> cons) {
        acc = operation.apply(acc, cons.value());
        list = cons.tail();
    }
    return acc;
}

public static final java.util.function.Function<Object, Object> runListOpsFFI = Bench.nativeBenchmark(input ->
    foldl((acc, value) -> acc + value, 0, filterEvens(range(1, ((Number) input).intValue()))));
