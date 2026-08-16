public class Effect_Exception_Unsafe {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object unsafeThrowException = (java.util.function.Function<Object, Object>) (x_0) -> ((java.util.function.Function<Object, Object>) (Effect_Unsafe.unsafePerformEffect)).apply(((java.util.function.Function<Object, Object>) (Effect_Exception.throwException)).apply(x_0));
public static final Object unsafeThrow = (java.util.function.Function<Object, Object>) (x_0) -> ((java.util.function.Function<Object, Object>) (Effect_Unsafe.unsafePerformEffect)).apply(((java.util.function.Function<Object, Object>) (Effect_Exception.throwException)).apply(((java.util.function.Function<Object, Object>) (Effect_Exception.error)).apply(x_0)));
}
