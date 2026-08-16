public class Control_Monad_Cont {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object withCont = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> (java.util.function.Function<Object, Object>) (k_2) -> ((java.util.function.Function<Object, Object>) (v_1)).apply((new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_3 = ((java.util.function.Function<Object, Object>) (f_0)).apply((java.util.function.Function<Object, Object>) (x_3) -> ((java.util.function.Function<Object, Object>) (k_2)).apply(x_3)); return (java.util.function.Function<Object, Object>) (x_4) -> ((java.util.function.Function<Object, Object>) (__local_var_3)).apply(x_4); } }).get());
public static final Object runCont = (java.util.function.Function<Object, Object>) (cc_0) -> (java.util.function.Function<Object, Object>) (k_1) -> ((java.util.function.Function<Object, Object>) (cc_0)).apply((java.util.function.Function<Object, Object>) (x_2) -> ((java.util.function.Function<Object, Object>) (k_1)).apply(x_2));
public static final Object mapCont = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> (java.util.function.Function<Object, Object>) (k_2) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(((java.util.function.Function<Object, Object>) (v_1)).apply(k_2));
public static final Object cont = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (c_1) -> ((java.util.function.Function<Object, Object>) (f_0)).apply((java.util.function.Function<Object, Object>) (x_2) -> ((java.util.function.Function<Object, Object>) (c_1)).apply(x_2));
}
