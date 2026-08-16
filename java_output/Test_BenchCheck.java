public class Test_BenchCheck {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object act = (new java.util.function.Supplier<Object>() { public Object get() { Object t1_0 = ((java.util.function.Supplier) (Object)(Bench.benchNow)).get(); Object t2_1 = ((java.util.function.Supplier) (Object)(Bench.benchNow)).get(); return ((java.util.function.Supplier) (Object)(((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) (arg) -> (java.util.function.Supplier<Object>) () -> { System.out.println(arg); return null; })).apply((((String) ("Delta: ")) + ((String) (((java.util.function.Function<Object, Object>) (Data_Show.showNumberImpl)).apply((((Double) (t2_1)) - ((Double) (t1_0)))))))))).get(); } });
}
