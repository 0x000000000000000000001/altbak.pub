public class Data_Monoid_Alternate {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object Alternate = (java.util.function.Function<Object, Object>) (x_0) -> x_0;
public static final Object showAlternate = (java.util.function.Function<Object, Object>) (dictShow_0) -> (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", (java.util.function.Function<Object, Object>) (v_1) -> (((String) ((((String) ("(Alternate ")) + ((String) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(v_1)))))) + ((String) (")"))));  return __map; } }).get();
public static final Object semigroupAlternate = (java.util.function.Function<Object, Object>) (dictAlt_0) -> (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("append", (java.util.function.Function<Object, Object>) (v_1) -> (java.util.function.Function<Object, Object>) (v1_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictAlt_0).get("alt"))).apply(v_1))).apply(v1_2));  return __map; } }).get();
public static final Object plusAlternate = (java.util.function.Function<Object, Object>) (dictPlus_0) -> dictPlus_0;
public static final Object ordAlternate = (java.util.function.Function<Object, Object>) (dictOrd_0) -> dictOrd_0;
public static final Object ord1Alternate = (java.util.function.Function<Object, Object>) (dictOrd1_0) -> dictOrd1_0;
public static final Object newtypeAlternate = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("Coercible0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> null /* TODO: PrimUndefined */);  return __map; } }).get();
public static final Object monoidAlternate = (java.util.function.Function<Object, Object>) (dictPlus_0) -> { Object semigroupAlternate1_1 = (new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_1 = ((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictPlus_0).get("Alt0"))).apply(null /* TODO: PrimUndefined */); return (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("append", (java.util.function.Function<Object, Object>) (v_2) -> (java.util.function.Function<Object, Object>) (v1_3) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) __local_var_1).get("alt"))).apply(v_2))).apply(v1_3));  return __map; } }).get(); } }).get(); return (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("mempty", ((java.util.LinkedHashMap<String, Object>) dictPlus_0).get("empty")); __map.put("Semigroup0", (java.util.function.Function<Object, Object>) (_dollar___unused_2) -> semigroupAlternate1_1);  return __map; } }).get(); };
public static final Object monadAlternate = (java.util.function.Function<Object, Object>) (dictMonad_0) -> dictMonad_0;
public static final Object functorAlternate = (java.util.function.Function<Object, Object>) (dictFunctor_0) -> dictFunctor_0;
public static final Object extendAlternate = (java.util.function.Function<Object, Object>) (dictExtend_0) -> dictExtend_0;
public static final Object eqAlternate = (java.util.function.Function<Object, Object>) (dictEq_0) -> dictEq_0;
public static final Object eq1Alternate = (java.util.function.Function<Object, Object>) (dictEq1_0) -> dictEq1_0;
public static final Object comonadAlternate = (java.util.function.Function<Object, Object>) (dictComonad_0) -> dictComonad_0;
public static final Object boundedAlternate = (java.util.function.Function<Object, Object>) (dictBounded_0) -> dictBounded_0;
public static final Object bindAlternate = (java.util.function.Function<Object, Object>) (dictBind_0) -> dictBind_0;
public static final Object applyAlternate = (java.util.function.Function<Object, Object>) (dictApply_0) -> dictApply_0;
public static final Object applicativeAlternate = (java.util.function.Function<Object, Object>) (dictApplicative_0) -> dictApplicative_0;
public static final Object alternativeAlternate = (java.util.function.Function<Object, Object>) (dictAlternative_0) -> dictAlternative_0;
public static final Object altAlternate = (java.util.function.Function<Object, Object>) (dictAlt_0) -> dictAlt_0;
}
