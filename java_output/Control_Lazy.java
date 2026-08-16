public class Control_Lazy {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object lazyUnit = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("defer", (java.util.function.Function<Object, Object>) (v_0) -> Data_Unit.unit);  return __map; } }).get();
public static final Object lazyFn = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("defer", (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (x_1) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (f_0)).apply(Data_Unit.unit))).apply(x_1));  return __map; } }).get();
public static final Object defer = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("defer");
public static final Object fix = (java.util.function.Function<Object, Object>) (dictLazy_0) -> (java.util.function.Function<Object, Object>) (f_1) -> (new java.util.function.Supplier<Object>() { public Object get() { class LetRecScope { Object go__go_2; LetRecScope() { go__go_2 = ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictLazy_0).get("defer"))).apply((java.util.function.Function<Object, Object>) (v_3) -> ((java.util.function.Function<Object, Object>) (f_1)).apply(go__go_2)); } } LetRecScope _scope = new LetRecScope(); Object go__go_2 = _scope.go__go_2; return go__go_2; } }).get();
}
