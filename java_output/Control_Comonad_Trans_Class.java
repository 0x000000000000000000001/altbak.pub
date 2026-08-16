public class Control_Comonad_Trans_Class {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object lower = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("lower");
public static final Object comonadTransIdentityT = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("lower", (java.util.function.Function<Object, Object>) (dictComonad_0) -> Control_Monad_Identity_Trans.runIdentityT);  return __map; } }).get();
}
