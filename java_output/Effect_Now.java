public class Effect_Now {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object getTimezoneOffset = FFI_STUB;
    public static Object getTimezoneOffset(Object... args) { return null; }
    public static Object now = FFI_STUB;
    public static Object now(Object... args) { return null; }

public static final Object nowTime = (new java.util.function.Supplier<Object>() { public Object get() { Object a_prime__0 = ((java.util.function.Supplier) (Object)(Effect_Now.now)).get(); return ((Data_DateTime.DateTime) (Object)(((java.util.function.Function<Object, Object>) (Data_DateTime_Instant.toDateTime)).apply(a_prime__0))).value1; } });
public static final Object nowDateTime = (new java.util.function.Supplier<Object>() { public Object get() { Object a_prime__0 = ((java.util.function.Supplier) (Object)(Effect_Now.now)).get(); return ((java.util.function.Function<Object, Object>) (Data_DateTime_Instant.toDateTime)).apply(a_prime__0); } });
public static final Object nowDate = (new java.util.function.Supplier<Object>() { public Object get() { Object a_prime__0 = ((java.util.function.Supplier) (Object)(Effect_Now.now)).get(); return ((Data_DateTime.DateTime) (Object)(((java.util.function.Function<Object, Object>) (Data_DateTime_Instant.toDateTime)).apply(a_prime__0))).value0; } });
}
