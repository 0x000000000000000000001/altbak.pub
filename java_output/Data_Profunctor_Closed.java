public class Data_Profunctor_Closed {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object closedFunction = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("closed", ((java.util.LinkedHashMap<String, Object>) Control_Semigroupoid.semigroupoidFn).get("compose")); __map.put("Profunctor0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Data_Profunctor.profunctorFn);  return __map; } }).get();
public static final Object closed = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("closed");
}
