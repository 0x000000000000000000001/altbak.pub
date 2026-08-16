public class Data_Symbol {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object unsafeCoerce = FFI_STUB;
    public static Object unsafeCoerce(Object... args) { return null; }

public static final Object reifySymbol = (java.util.function.Function<Object, Object>) (s_0) -> (java.util.function.Function<Object, Object>) (f_1) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Symbol.unsafeCoerce)).apply((java.util.function.Function<Object, Object>) (dictIsSymbol_2) -> ((java.util.function.Function<Object, Object>) (f_1)).apply(dictIsSymbol_2)))).apply((new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("reflectSymbol", (java.util.function.Function<Object, Object>) (v_2) -> s_0);  return __map; } }).get()))).apply(new Type_Proxy.Proxy());
public static final Object reflectSymbol = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("reflectSymbol");
}
