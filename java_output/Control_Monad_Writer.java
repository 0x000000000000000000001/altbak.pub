public class Control_Monad_Writer {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object writer = (java.util.function.Function<Object, Object>) (x_0) -> x_0;
public static final Object runWriter = (java.util.function.Function<Object, Object>) (x_0) -> x_0;
public static final Object mapWriter = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(v_1);
public static final Object execWriter = (java.util.function.Function<Object, Object>) (m_0) -> ((Data_Tuple.Tuple) (Object)(m_0)).value1;
}
