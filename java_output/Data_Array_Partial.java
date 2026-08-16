public class Data_Array_Partial {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object tail = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> (java.util.function.Function<Object, Object>) (xs_1) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Array.sliceImpl)).apply(1))).apply(((Object[]) xs_1).length))).apply(xs_1);
public static final Object last = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> (java.util.function.Function<Object, Object>) (xs_1) -> ((Object[]) xs_1)[((Integer) ((((Integer) (((Object[]) xs_1).length)) - ((Integer) (1)))))];
public static final Object init = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> (java.util.function.Function<Object, Object>) (xs_1) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Array.sliceImpl)).apply(0))).apply((((Integer) (((Object[]) xs_1).length)) - ((Integer) (1)))))).apply(xs_1);
public static final Object head = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> (java.util.function.Function<Object, Object>) (xs_1) -> ((Object[]) xs_1)[((Integer) (0))];
}
