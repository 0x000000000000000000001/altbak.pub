public class Test_RowToList {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object keysNil = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("keysImpl", (java.util.function.Function<Object, Object>) (v_0) -> 0);  return __map; } }).get();
public static final Object keysImpl = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("keysImpl");
public static final Object keysCons = (java.util.function.Function<Object, Object>) (dictRecordKeys_0) -> (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("keysImpl", (java.util.function.Function<Object, Object>) (v_1) -> (((Integer) (1)) + ((Integer) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictRecordKeys_0).get("keysImpl"))).apply(new Type_Proxy.Proxy())))));  return __map; } }).get();
public static final Object keys = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> (java.util.function.Function<Object, Object>) (dictRecordKeys_1) -> (java.util.function.Function<Object, Object>) (v_2) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictRecordKeys_1).get("keysImpl"))).apply(new Type_Proxy.Proxy());
public static final Object describe = ((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) (arg) -> (java.util.function.Supplier<Object>) () -> { System.out.println(arg); return null; })).apply("RowToList (Keys Count):");
public static final Object act = (new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_0 = ((java.util.function.Function<Object, Object>) (Bench.opaque)).apply(10000); Object _dollar___unused_1 = ((java.util.function.Supplier) (Object)(__local_var_0)).get(); return ((java.util.function.Supplier) (Object)(((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) (arg) -> (java.util.function.Supplier<Object>) () -> { System.out.println(arg); return null; })).apply(((java.util.function.Function<Object, Object>) (Data_Show.showIntImpl)).apply(5)))).get(); } });
}
