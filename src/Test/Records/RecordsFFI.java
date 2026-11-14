private record Inner(int e, int f) {}
private record Middle(int c, Inner d) {}
private record DeepRecord(int a, Middle b) {}

private static DeepRecord updateRec(int n, DeepRecord record) {
    while (n > 0) {
        record = new DeepRecord(record.a() + 1,
            new Middle(record.b().c() + 2,
                new Inner(record.b().d().e() + 3, record.b().d().f() + n % 5)));
        n--;
    }
    return record;
}

public static final java.util.function.Function<Object, Object> runRecordsFFI = Bench.nativeBenchmark(input ->
    updateRec(((Number) input).intValue(), new DeepRecord(0, new Middle(0, new Inner(0, 0)))).b().d().f());
