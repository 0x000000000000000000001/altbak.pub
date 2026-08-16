public class Control_Monad_RWS {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object withRWS = Control_Monad_RWS_Trans.withRWST;
public static final Object rws = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (r_1) -> (java.util.function.Function<Object, Object>) (s_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (f_0)).apply(r_1))).apply(s_2);
public static final Object runRWS = (java.util.function.Function<Object, Object>) (m_0) -> (java.util.function.Function<Object, Object>) (r_1) -> (java.util.function.Function<Object, Object>) (s_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (m_0)).apply(r_1))).apply(s_2);
public static final Object mapRWS = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> (java.util.function.Function<Object, Object>) (r_2) -> (java.util.function.Function<Object, Object>) (s_3) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (v_1)).apply(r_2))).apply(s_3));
public static final Object execRWS = (java.util.function.Function<Object, Object>) (m_0) -> (java.util.function.Function<Object, Object>) (r_1) -> (java.util.function.Function<Object, Object>) (s_2) -> { Object __local_var_3 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (m_0)).apply(r_1))).apply(s_2); return new Data_Tuple.Tuple(((Control_Monad_RWS_Trans.RWSResult) (Object)(__local_var_3)).value0, ((Control_Monad_RWS_Trans.RWSResult) (Object)(__local_var_3)).value2); };
public static final Object evalRWS = (java.util.function.Function<Object, Object>) (m_0) -> (java.util.function.Function<Object, Object>) (r_1) -> (java.util.function.Function<Object, Object>) (s_2) -> { Object __local_var_3 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (m_0)).apply(r_1))).apply(s_2); return new Data_Tuple.Tuple(((Control_Monad_RWS_Trans.RWSResult) (Object)(__local_var_3)).value1, ((Control_Monad_RWS_Trans.RWSResult) (Object)(__local_var_3)).value2); };
}
