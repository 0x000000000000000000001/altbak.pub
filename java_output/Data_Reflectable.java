public class Data_Reflectable {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object unsafeCoerce = FFI_STUB;
    public static Object unsafeCoerce(Object... args) { return null; }

public static final Object reifiableString = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>();  return __map; } }).get();
public static final Object reifiableOrdering = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>();  return __map; } }).get();
public static final Object reifiableInt = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>();  return __map; } }).get();
public static final Object reifiableBoolean = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>();  return __map; } }).get();
public static final Object reifyType = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> (java.util.function.Function<Object, Object>) (s_1) -> (java.util.function.Function<Object, Object>) (f_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Reflectable.unsafeCoerce)).apply((java.util.function.Function<Object, Object>) (dictReflectable_3) -> ((java.util.function.Function<Object, Object>) (f_2)).apply(dictReflectable_3)))).apply((new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("reflectType", (java.util.function.Function<Object, Object>) (v_3) -> s_1);  return __map; } }).get()))).apply(new Type_Proxy.Proxy());
public static final Object reflectType = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("reflectType");
}
