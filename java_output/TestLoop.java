public class TestLoop {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object loop = (java.util.function.Function<Object, Object>) (v_0) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_v_0 = v_0; while(true) { final Object __final_v_0 = __tco_v_0; try { return ( ((Boolean) (java.util.Objects.equals(__final_v_0, 0))) ? 0 : ((java.util.function.Function<Object, Object>) (TestLoop.loop)).apply((((Integer) (__final_v_0)) - ((Integer) (1))))); } catch (TcoLoop __tco_ex) { __tco_v_0 = __tco_ex.args[0]; } } } }).get();
}
