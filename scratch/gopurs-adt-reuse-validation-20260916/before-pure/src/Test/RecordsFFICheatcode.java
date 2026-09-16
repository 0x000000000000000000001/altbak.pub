private static final class Inner {
    int e;
    int f;
}

private static final class Middle {
    int c;
    final Inner d = new Inner();
}

private static final class DeepRecord {
    int a;
    final Middle b = new Middle();
}

public static final java.util.function.Function<Object, Object> runRecordsFFICheatcode = Bench.nativeBenchmark(input -> {
    int n = ((Number) input).intValue();
    DeepRecord record = new DeepRecord();
    while (n > 0) {
        record.a++;
        record.b.c += 2;
        record.b.d.e += 3;
        record.b.d.f += n % 5;
        n--;
    }
    return record.b.d.f;
});
