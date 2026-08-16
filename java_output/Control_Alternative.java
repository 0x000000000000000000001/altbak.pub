public class Control_Alternative {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object guard = (java.util.function.Function<Object, Object>) (dictAlternative_0) -> { Object Applicative0_1 = ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictAlternative_0).get("Applicative0"))).apply(null /* TODO: PrimUndefined */); Object empty_2 = ((java.util.LinkedHashMap<String, Object>) ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictAlternative_0).get("Plus1"))).apply(null /* TODO: PrimUndefined */)).get("empty"); return (java.util.function.Function<Object, Object>) (v_3) -> ( ((Boolean) (v_3)) ? ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) Applicative0_1).get("pure"))).apply(Data_Unit.unit) : empty_2); };
public static final Object alternativeArray = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("Applicative0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Control_Applicative.applicativeArray); __map.put("Plus1", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Control_Plus.plusArray);  return __map; } }).get();
}
