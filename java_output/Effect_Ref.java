public class Effect_Ref {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object _new = FFI_STUB;
    public static Object _new(Object... args) { return null; }
    public static Object modifyImpl = FFI_STUB;
    public static Object modifyImpl(Object... args) { return null; }
    public static Object newWithSelf = FFI_STUB;
    public static Object newWithSelf(Object... args) { return null; }
    public static Object read = FFI_STUB;
    public static Object read(Object... args) { return null; }
    public static Object write = FFI_STUB;
    public static Object write(Object... args) { return null; }

public static final Object $new = Effect_Ref._new;
public static final Object modifyprime = Effect_Ref.modifyImpl;
public static final Object modify = (java.util.function.Function<Object, Object>) (f_0) -> ((java.util.function.Function<Object, Object>) (Effect_Ref.modifyImpl)).apply((java.util.function.Function<Object, Object>) (s_1) -> { Object s_prime__2 = ((java.util.function.Function<Object, Object>) (f_0)).apply(s_1); return (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("state", s_prime__2); __map.put("value", s_prime__2);  return __map; } }).get(); });
public static final Object modify_ = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (s_1) -> { Object __local_var_2 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Effect_Ref.modifyImpl)).apply((java.util.function.Function<Object, Object>) (s_2) -> { Object s_prime__3 = ((java.util.function.Function<Object, Object>) (f_0)).apply(s_2); return (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("state", s_prime__3); __map.put("value", s_prime__3);  return __map; } }).get(); }))).apply(s_1); return (new java.util.function.Supplier<Object>() { public Object get() { Object a_prime__3 = ((java.util.function.Supplier) (Object)(__local_var_2)).get(); return Data_Unit.unit; } }); };
}
