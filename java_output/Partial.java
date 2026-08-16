public class Partial {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    // FFI provided by ../javapurs/javapurs-partial/src/Partial.java
    public static Object _crashWith = (java.util.function.Function<Object, Object>) (msg) -> {
        throw new RuntimeException((String) msg);
    };


public static final Object crashWith = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Partial._crashWith;
public static final Object crash = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> ((java.util.function.Function<Object, Object>) (Partial._crashWith)).apply("Partial.crash: partial function");
}
