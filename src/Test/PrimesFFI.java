private static final class IntList {
    final int head;
    final IntList tail;

    IntList(int head, IntList tail) {
        this.head = head;
        this.tail = tail;
    }
}

private static IntList range(int start, int end) {
    IntList result = null;
    for (int current = end; current >= start; current--) {
        result = new IntList(current, result);
    }
    return result;
}

private static IntList reverse(IntList list) {
    IntList result = null;
    for (; list != null; list = list.tail) {
        result = new IntList(list.head, result);
    }
    return result;
}

private static IntList filter(java.util.function.IntPredicate predicate, IntList list) {
    IntList result = null;
    for (; list != null; list = list.tail) {
        if (predicate.test(list.head)) result = new IntList(list.head, result);
    }
    return reverse(result);
}

private static IntList sieve(IntList list) {
    if (list == null) return null;
    int prime = list.head;
    return new IntList(prime, sieve(filter(x -> x % prime != 0, list.tail)));
}

public static final java.util.function.Function<Object, Object> runPrimesFFI = Bench.nativeBenchmark(input -> {
    IntList primes = sieve(range(2, ((Number) input).intValue()));
    int sum = 0;
    for (; primes != null; primes = primes.tail) sum += primes.head;
    return sum;
});
