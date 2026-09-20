private static final class List<A> {
    final A head;
    final List<A> tail;

    List(A head, List<A> tail) {
        this.head = head;
        this.tail = tail;
    }
}

private static List<Integer> range(int start, int end) {
    List<Integer> result = null;
    for (int current = end; current >= start; current--) {
        result = new List<>(current, result);
    }
    return result;
}

private static <A> List<A> reverse(List<A> list) {
    List<A> result = null;
    for (; list != null; list = list.tail) {
        result = new List<>(list.head, result);
    }
    return result;
}

private static <A> List<A> filter(java.util.function.Predicate<A> predicate, List<A> list) {
    List<A> result = null;
    for (; list != null; list = list.tail) {
        if (predicate.test(list.head)) result = new List<>(list.head, result);
    }
    return reverse(result);
}

private static List<Integer> sieve(List<Integer> list) {
    if (list == null) return null;
    int prime = list.head;
    return new List<>(prime, sieve(filter(x -> x % prime != 0, list.tail)));
}

public static final java.util.function.Function<Object, Object> runPrimesFFI = Bench.nativeBenchmark(input -> {
    List<Integer> primes = sieve(range(2, ((Number) input).intValue()));
    int sum = 0;
    for (; primes != null; primes = primes.tail) sum += primes.head;
    return sum;
});
