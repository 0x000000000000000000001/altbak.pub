public class Control_Monad_State {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object withState = Control_Monad_State_Trans.withStateT;
public static final Object runState = (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (x_1) -> ((java.util.function.Function<Object, Object>) (v_0)).apply(x_1);
public static final Object mapState = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> (java.util.function.Function<Object, Object>) (x_2) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(((java.util.function.Function<Object, Object>) (v_1)).apply(x_2));
public static final Object execState = (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (s_1) -> ((Data_Tuple.Tuple) (Object)(((java.util.function.Function<Object, Object>) (v_0)).apply(s_1))).value1;
public static final Object evalState = (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (s_1) -> ((Data_Tuple.Tuple) (Object)(((java.util.function.Function<Object, Object>) (v_0)).apply(s_1))).value0;
}
