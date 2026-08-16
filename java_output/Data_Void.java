public class Data_Void {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object absurd = (java.util.function.Function<Object, Object>) (a_0) -> { Object spin_1 = (java.util.function.Function<Object, Object>) (v_2) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_v_2 = v_2; while(true) { final Object __final_v_2 = __tco_v_2; try { return (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("spin_1", new Object[]{__final_v_2}); } }).get(); } catch (TcoLoop __tco_ex) { __tco_v_2 = __tco_ex.args[0]; } } } }).get(); return ((java.util.function.Function<Object, Object>) (spin_1)).apply(a_0); };
}
