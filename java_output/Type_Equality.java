public class Type_Equality {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object refl = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("proof", (java.util.function.Function<Object, Object>) (a_0) -> a_0); __map.put("Coercible0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> null /* TODO: PrimUndefined */);  return __map; } }).get();
public static final Object proof = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("proof");
public static final Object to = (java.util.function.Function<Object, Object>) (dictTypeEquals_0) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictTypeEquals_0).get("proof"))).apply((java.util.function.Function<Object, Object>) (a_1) -> a_1);
public static final Object from = (java.util.function.Function<Object, Object>) (dictTypeEquals_0) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictTypeEquals_0).get("proof"))).apply((java.util.function.Function<Object, Object>) (a_1) -> a_1);
}
