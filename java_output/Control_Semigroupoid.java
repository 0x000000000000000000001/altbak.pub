public class Control_Semigroupoid {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object semigroupoidFn = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("compose", (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (g_1) -> (java.util.function.Function<Object, Object>) (x_2) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(((java.util.function.Function<Object, Object>) (g_1)).apply(x_2)));  return __map; } }).get();
public static final Object compose = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("compose");
public static final Object composeFlipped = (java.util.function.Function<Object, Object>) (dictSemigroupoid_0) -> (java.util.function.Function<Object, Object>) (f_1) -> (java.util.function.Function<Object, Object>) (g_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictSemigroupoid_0).get("compose"))).apply(g_2))).apply(f_1);
}
