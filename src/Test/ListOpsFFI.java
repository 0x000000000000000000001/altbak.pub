private sealed interface IntList permits Nil, Cons {}
private enum Nil implements IntList { INSTANCE }
private record Cons(int value, IntList tail) implements IntList {}

private static IntList range(int start, int end) {
    IntList result = Nil.INSTANCE;
    for (int current = end; current >= start; current--) {
        result = new Cons(current, result);
    }
    return result;
}

private static IntList filterEvens(IntList list) {
    IntList result = Nil.INSTANCE;
    while (list instanceof Cons cons) {
        if (cons.value() % 2 == 0) result = new Cons(cons.value(), result);
        list = cons.tail();
    }
    return result;
}

private static int foldl(java.util.function.IntBinaryOperator operation, int acc, IntList list) {
    while (list instanceof Cons cons) {
        acc = operation.applyAsInt(acc, cons.value());
        list = cons.tail();
    }
    return acc;
}

public static final java.util.function.Function<Object, Object> runListOpsFFI = Bench.nativeBenchmark(input ->
    foldl((acc, value) -> acc + value, 0, filterEvens(range(1, ((Number) input).intValue()))));
