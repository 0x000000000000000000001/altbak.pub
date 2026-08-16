public class Data_Array_ST_Partial {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object peekImpl = FFI_STUB;
    public static Object peekImpl(Object... args) { return null; }
    public static Object pokeImpl = FFI_STUB;
    public static Object pokeImpl(Object... args) { return null; }

public static final Object poke = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> ((java.util.function.Function<Object, Object>) (Control_Monad_ST_Uncurried.runSTFn3)).apply(Data_Array_ST_Partial.pokeImpl);
public static final Object peek = (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> ((java.util.function.Function<Object, Object>) (Control_Monad_ST_Uncurried.runSTFn2)).apply(Data_Array_ST_Partial.peekImpl);
}
