public class Data_Functor_Contravariant {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object contravariantConst = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("cmap", (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (v1_1) -> v1_1);  return __map; } }).get();
public static final Object cmap = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("cmap");
public static final Object cmapFlipped = (java.util.function.Function<Object, Object>) (dictContravariant_0) -> (java.util.function.Function<Object, Object>) (x_1) -> (java.util.function.Function<Object, Object>) (f_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictContravariant_0).get("cmap"))).apply(f_2))).apply(x_1);
public static final Object coerce = (java.util.function.Function<Object, Object>) (dictContravariant_0) -> (java.util.function.Function<Object, Object>) (dictFunctor_1) -> (java.util.function.Function<Object, Object>) (a_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictFunctor_1).get("map"))).apply(Data_Void.absurd))).apply(((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictContravariant_0).get("cmap"))).apply(Data_Void.absurd))).apply(a_2));
public static final Object imapC = (java.util.function.Function<Object, Object>) (dictContravariant_0) -> (java.util.function.Function<Object, Object>) (v_1) -> (java.util.function.Function<Object, Object>) (f_2) -> ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictContravariant_0).get("cmap"))).apply(f_2);
}
