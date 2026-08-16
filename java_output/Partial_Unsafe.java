public class Partial_Unsafe {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    // FFI provided by ../javapurs/javapurs-partial/src/Partial/Unsafe.java
    public static Object _unsafePartial = (java.util.function.Function<Object, Object>) (f) -> {
        return ((java.util.function.Function<Object, Object>) f).apply(null /* TODO: PrimUndefined */);
    };


public static final Object unsafePartial = Partial_Unsafe._unsafePartial;
public static final Object unsafeCrashWith = (java.util.function.Function<Object, Object>) (msg_0) -> ((java.util.function.Function<Object, Object>) (Partial._crashWith)).apply(msg_0);
}
