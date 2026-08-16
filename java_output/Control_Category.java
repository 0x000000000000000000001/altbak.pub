public class Control_Category {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object identity = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("identity");
public static final Object categoryFn = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("identity", (java.util.function.Function<Object, Object>) (x_0) -> x_0); __map.put("Semigroupoid0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Control_Semigroupoid.semigroupoidFn);  return __map; } }).get();
}
