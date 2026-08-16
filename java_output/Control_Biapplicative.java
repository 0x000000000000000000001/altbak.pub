public class Control_Biapplicative {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object bipure = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("bipure");
public static final Object biapplicativeTuple = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("bipure", Data_Tuple.Tuple); __map.put("Biapply0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Control_Biapply.biapplyTuple);  return __map; } }).get();
}
