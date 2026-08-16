public class Control_Bind {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    // FFI provided by ../javapurs/javapurs-prelude/src/Control/Bind.java
    public static Object arrayBind = (java.util.function.Function<Object, Object>) (arrObj) -> (java.util.function.Function<Object, Object>) (f) -> {
        Object[] arr = (Object[]) arrObj;
        java.util.List<Object> result = new java.util.ArrayList<>();
        for (Object item : arr) {
            Object[] mapped = (Object[]) ((java.util.function.Function<Object, Object>) f).apply(item);
            for (Object mappedItem : mapped) {
                result.add(mappedItem);
            }
        }
        return result.toArray(new Object[0]);
    };


public static final Object discard = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("discard");
public static final Object bindProxy = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("bind", (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (v1_1) -> new Type_Proxy.Proxy()); __map.put("Apply0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Control_Apply.applyProxy);  return __map; } }).get();
public static final Object bindFn = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("bind", (java.util.function.Function<Object, Object>) (m_0) -> (java.util.function.Function<Object, Object>) (f_1) -> (java.util.function.Function<Object, Object>) (x_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (f_1)).apply(((java.util.function.Function<Object, Object>) (m_0)).apply(x_2)))).apply(x_2)); __map.put("Apply0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Control_Apply.applyFn);  return __map; } }).get();
public static final Object bindArray = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("bind", Control_Bind.arrayBind); __map.put("Apply0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Control_Apply.applyArray);  return __map; } }).get();
public static final Object bind = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("bind");
public static final Object bindFlipped = (java.util.function.Function<Object, Object>) (dictBind_0) -> (java.util.function.Function<Object, Object>) (b_1) -> (java.util.function.Function<Object, Object>) (a_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictBind_0).get("bind"))).apply(a_2))).apply(b_1);
public static final Object composeKleisliFlipped = (java.util.function.Function<Object, Object>) (dictBind_0) -> (java.util.function.Function<Object, Object>) (f_1) -> (java.util.function.Function<Object, Object>) (g_2) -> (java.util.function.Function<Object, Object>) (a_3) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictBind_0).get("bind"))).apply(((java.util.function.Function<Object, Object>) (g_2)).apply(a_3)))).apply(f_1);
public static final Object composeKleisli = (java.util.function.Function<Object, Object>) (dictBind_0) -> (java.util.function.Function<Object, Object>) (f_1) -> (java.util.function.Function<Object, Object>) (g_2) -> (java.util.function.Function<Object, Object>) (a_3) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictBind_0).get("bind"))).apply(((java.util.function.Function<Object, Object>) (f_1)).apply(a_3)))).apply(g_2);
public static final Object discardProxy = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("discard", (java.util.function.Function<Object, Object>) (dictBind_0) -> ((java.util.LinkedHashMap<String, Object>) dictBind_0).get("bind"));  return __map; } }).get();
public static final Object discardUnit = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("discard", (java.util.function.Function<Object, Object>) (dictBind_0) -> ((java.util.LinkedHashMap<String, Object>) dictBind_0).get("bind"));  return __map; } }).get();
public static final Object ifM = (java.util.function.Function<Object, Object>) (dictBind_0) -> (java.util.function.Function<Object, Object>) (cond_1) -> (java.util.function.Function<Object, Object>) (t_2) -> (java.util.function.Function<Object, Object>) (f_3) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictBind_0).get("bind"))).apply(cond_1))).apply((java.util.function.Function<Object, Object>) (cond_prime__4) -> ( ((Boolean) (cond_prime__4)) ? t_2 : f_3));
public static final Object join = (java.util.function.Function<Object, Object>) (dictBind_0) -> (java.util.function.Function<Object, Object>) (m_1) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictBind_0).get("bind"))).apply(m_1))).apply((java.util.function.Function<Object, Object>) (x_2) -> x_2);
}
