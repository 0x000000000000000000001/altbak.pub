public class Control_Monad_Reader {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object withReader = Control_Monad_Reader_Trans.withReaderT;
public static final Object runReader = (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (x_1) -> ((java.util.function.Function<Object, Object>) (v_0)).apply(x_1);
public static final Object mapReader = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> (java.util.function.Function<Object, Object>) (x_2) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(((java.util.function.Function<Object, Object>) (v_1)).apply(x_2));
}
