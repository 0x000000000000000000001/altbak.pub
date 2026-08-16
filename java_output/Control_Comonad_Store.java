public class Control_Comonad_Store {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object store = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (x_1) -> new Data_Tuple.Tuple(f_0, x_1);
public static final Object runStore = (java.util.function.Function<Object, Object>) (v_0) -> { Object __local_var_1 = (new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_1 = new Data_Tuple.Tuple(((Data_Tuple.Tuple) (Object)(v_0)).value1, ((Data_Tuple.Tuple) (Object)(v_0)).value0); return new Data_Tuple.Tuple(((Data_Tuple.Tuple) (Object)(__local_var_1)).value0, ((Data_Tuple.Tuple) (Object)(__local_var_1)).value1); } }).get(); return new Data_Tuple.Tuple(((Data_Tuple.Tuple) (Object)(__local_var_1)).value1, ((Data_Tuple.Tuple) (Object)(__local_var_1)).value0); };
}
