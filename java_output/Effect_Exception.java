public class Effect_Exception {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object catchException = FFI_STUB;
    public static Object catchException(Object... args) { return null; }
    public static Object error = FFI_STUB;
    public static Object error(Object... args) { return null; }
    public static Object errorWithCause = FFI_STUB;
    public static Object errorWithCause(Object... args) { return null; }
    public static Object errorWithName = FFI_STUB;
    public static Object errorWithName(Object... args) { return null; }
    public static Object message = FFI_STUB;
    public static Object message(Object... args) { return null; }
    public static Object name = FFI_STUB;
    public static Object name(Object... args) { return null; }
    public static Object showErrorImpl = FFI_STUB;
    public static Object showErrorImpl(Object... args) { return null; }
    public static Object stackImpl = FFI_STUB;
    public static Object stackImpl(Object... args) { return null; }
    public static Object throwException = FFI_STUB;
    public static Object throwException(Object... args) { return null; }

public static final Object $try = (java.util.function.Function<Object, Object>) (action_0) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Effect_Exception.catchException)).apply((java.util.function.Function<Object, Object>) (x_1) -> (new java.util.function.Supplier<Object>() { public Object get() { return new Data_Either.Left(x_1); } })))).apply((new java.util.function.Supplier<Object>() { public Object get() { Object a_prime__1 = ((java.util.function.Supplier) (Object)(action_0)).get(); return new Data_Either.Right(a_prime__1); } }));
public static final Object $throw = (java.util.function.Function<Object, Object>) (x_0) -> ((java.util.function.Function<Object, Object>) (Effect_Exception.throwException)).apply(((java.util.function.Function<Object, Object>) (Effect_Exception.error)).apply(x_0));
public static final Object stack = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Effect_Exception.stackImpl)).apply(Data_Maybe.Just))).apply(new Data_Maybe.Nothing());
public static final Object showError = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", Effect_Exception.showErrorImpl);  return __map; } }).get();
}
