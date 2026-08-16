public class Effect_Unsafe {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object unsafePerformEffect = FFI_STUB;
    public static Object unsafePerformEffect(Object... args) { return null; }


}
