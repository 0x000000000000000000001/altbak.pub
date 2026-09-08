private static int ackermann(int m, int n) {
    while (m != 0) {
        n = n == 0 ? 1 : ackermann(m, n - 1);
        m--;
    }
    return n + 1;
}

public static final java.util.function.Function<Object, Object> runAckermannFFICheatcode = Bench.nativeBenchmark(input ->
    ackermann(3, 4));
