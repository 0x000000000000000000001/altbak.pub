public class Test_Fib {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object fib = (java.util.function.Function<Object, Object>) (v_0) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_v_0 = v_0; while(true) { final Object __final_v_0 = __tco_v_0; try { return ( ((Boolean) (java.util.Objects.equals(__final_v_0, 0))) ? 0 : ( ((Boolean) (java.util.Objects.equals(__final_v_0, 1))) ? 1 : (((Integer) (((java.util.function.Function<Object, Object>) (Test_Fib.fib)).apply((((Integer) (__final_v_0)) - ((Integer) (1)))))) + ((Integer) (((java.util.function.Function<Object, Object>) (Test_Fib.fib)).apply((((Integer) (__final_v_0)) - ((Integer) (2))))))))); } catch (TcoLoop __tco_ex) { __tco_v_0 = __tco_ex.args[0]; } } } }).get();
public static final Object describe = ((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) (arg) -> (java.util.function.Supplier<Object>) () -> { System.out.println(arg); return null; })).apply("Fibonacci:");
public static final Object act = (new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_0 = ((java.util.function.Function<Object, Object>) (Bench.opaque)).apply(10); Object dummy_1 = ((java.util.function.Supplier) (Object)(__local_var_0)).get(); return ((java.util.function.Supplier) (Object)(((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) (arg) -> (java.util.function.Supplier<Object>) () -> { System.out.println(arg); return null; })).apply(((java.util.function.Function<Object, Object>) (Data_Show.showIntImpl)).apply(((java.util.function.Function<Object, Object>) (Test_Fib.fib)).apply(dummy_1))))).get(); } });
}
