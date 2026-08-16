public class Control_Monad_State_Class {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object state = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("state");
public static final Object put = (java.util.function.Function<Object, Object>) (dictMonadState_0) -> (java.util.function.Function<Object, Object>) (s_1) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictMonadState_0).get("state"))).apply((java.util.function.Function<Object, Object>) (v_2) -> new Data_Tuple.Tuple(Data_Unit.unit, s_1));
public static final Object modify_ = (java.util.function.Function<Object, Object>) (dictMonadState_0) -> (java.util.function.Function<Object, Object>) (f_1) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictMonadState_0).get("state"))).apply((java.util.function.Function<Object, Object>) (s_2) -> new Data_Tuple.Tuple(Data_Unit.unit, ((java.util.function.Function<Object, Object>) (f_1)).apply(s_2)));
public static final Object modify = (java.util.function.Function<Object, Object>) (dictMonadState_0) -> (java.util.function.Function<Object, Object>) (f_1) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictMonadState_0).get("state"))).apply((java.util.function.Function<Object, Object>) (s_2) -> { Object s_prime__3 = ((java.util.function.Function<Object, Object>) (f_1)).apply(s_2); return new Data_Tuple.Tuple(s_prime__3, s_prime__3); });
public static final Object gets = (java.util.function.Function<Object, Object>) (dictMonadState_0) -> (java.util.function.Function<Object, Object>) (f_1) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictMonadState_0).get("state"))).apply((java.util.function.Function<Object, Object>) (s_2) -> new Data_Tuple.Tuple(((java.util.function.Function<Object, Object>) (f_1)).apply(s_2), s_2));
public static final Object get = (java.util.function.Function<Object, Object>) (dictMonadState_0) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictMonadState_0).get("state"))).apply((java.util.function.Function<Object, Object>) (s_1) -> new Data_Tuple.Tuple(s_1, s_1));
}
