public class Control_Comonad_Env {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object withEnv = Control_Comonad_Env_Trans.withEnvT;
public static final Object runEnv = (java.util.function.Function<Object, Object>) (v_0) -> new Data_Tuple.Tuple(((Data_Tuple.Tuple) (Object)(v_0)).value0, ((Data_Tuple.Tuple) (Object)(v_0)).value1);
public static final Object mapEnv = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> new Data_Tuple.Tuple(((Data_Tuple.Tuple) (Object)(v_1)).value0, ((java.util.function.Function<Object, Object>) (f_0)).apply(((Data_Tuple.Tuple) (Object)(v_1)).value1));
public static final Object env = (java.util.function.Function<Object, Object>) (e_0) -> (java.util.function.Function<Object, Object>) (a_1) -> new Data_Tuple.Tuple(e_0, a_1);
}
