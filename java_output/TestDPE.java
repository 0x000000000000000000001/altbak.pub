public class TestDPE {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object mempty_ = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("mempty_");
public static final Object mappend_ = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("mappend_");
public static final Object polyLoop = (java.util.function.Function<Object, Object>) (dictMonoidish_0) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_dictMonoidish_0 = dictMonoidish_0; while(true) { final Object __final_dictMonoidish_0 = __tco_dictMonoidish_0; try { return (new java.util.function.Supplier<Object>() { public Object get() { Object mempty_1_1 = ((java.util.LinkedHashMap<String, Object>) __final_dictMonoidish_0).get("mempty_"); return (java.util.function.Function<Object, Object>) (v_2) -> (java.util.function.Function<Object, Object>) (v1_3) -> ( ((Boolean) (java.util.Objects.equals(v_2, 0))) ? v1_3 : ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (TestDPE.polyLoop)).apply(dictMonoidish_0))).apply((((Integer) (v_2)) - ((Integer) (1)))))).apply(((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictMonoidish_0).get("mappend_"))).apply(v1_3))).apply(mempty_1_1))); } }).get(); } catch (TcoLoop __tco_ex) { __tco_dictMonoidish_0 = __tco_ex.args[0]; } } } }).get();
public static final Object intMonoidish = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("mempty_", 0); __map.put("mappend_", (java.util.function.Function<Object, Object>) (a_0) -> (java.util.function.Function<Object, Object>) (b_1) -> (((Integer) (a_0)) + ((Integer) (b_1))));  return __map; } }).get();
public static final Object test = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (TestDPE.polyLoop)).apply(TestDPE.intMonoidish))).apply(9999999))).apply(0);
}
