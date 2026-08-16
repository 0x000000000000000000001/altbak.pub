public class Data_Profunctor_Split {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final class SplitF {
            public final Object value0;
            public final Object value1;
            public final Object value2;
            public SplitF(Object value0, Object value1, Object value2) {
                this.value0 = value0;
                this.value1 = value1;
                this.value2 = value2;
            }
        }
public static final Object SplitF = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> (java.util.function.Function<Object, Object>) (value2) -> new Data_Profunctor_Split.SplitF(value0, value1, value2);
public static final Object unSplit = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (f_0)).apply(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value0))).apply(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value1))).apply(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value2);
public static final Object split = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (g_1) -> (java.util.function.Function<Object, Object>) (fx_2) -> new Data_Profunctor_Split.SplitF(f_0, g_1, fx_2);
public static final Object profunctorSplit = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("dimap", (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (g_1) -> (java.util.function.Function<Object, Object>) (v_2) -> new Data_Profunctor_Split.SplitF((java.util.function.Function<Object, Object>) (x_3) -> ((java.util.function.Function<Object, Object>) (((Data_Profunctor_Split.SplitF) (Object)(v_2)).value0)).apply(((java.util.function.Function<Object, Object>) (f_0)).apply(x_3)), (java.util.function.Function<Object, Object>) (x_3) -> ((java.util.function.Function<Object, Object>) (g_1)).apply(((java.util.function.Function<Object, Object>) (((Data_Profunctor_Split.SplitF) (Object)(v_2)).value1)).apply(x_3)), ((Data_Profunctor_Split.SplitF) (Object)(v_2)).value2));  return __map; } }).get();
public static final Object lowerSplit = (java.util.function.Function<Object, Object>) (dictInvariant_0) -> (java.util.function.Function<Object, Object>) (v_1) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictInvariant_0).get("imap"))).apply(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value1))).apply(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value0))).apply(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value2);
public static final Object liftSplit = (java.util.function.Function<Object, Object>) (fx_0) -> new Data_Profunctor_Split.SplitF((java.util.function.Function<Object, Object>) (x_1) -> x_1, (java.util.function.Function<Object, Object>) (x_1) -> x_1, fx_0);
public static final Object hoistSplit = (java.util.function.Function<Object, Object>) (nat_0) -> (java.util.function.Function<Object, Object>) (v_1) -> new Data_Profunctor_Split.SplitF(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value0, ((Data_Profunctor_Split.SplitF) (Object)(v_1)).value1, ((java.util.function.Function<Object, Object>) (nat_0)).apply(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value2));
public static final Object functorSplit = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("map", (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> new Data_Profunctor_Split.SplitF(((Data_Profunctor_Split.SplitF) (Object)(v_1)).value0, (java.util.function.Function<Object, Object>) (x_2) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(((java.util.function.Function<Object, Object>) (((Data_Profunctor_Split.SplitF) (Object)(v_1)).value1)).apply(x_2)), ((Data_Profunctor_Split.SplitF) (Object)(v_1)).value2));  return __map; } }).get();
}
