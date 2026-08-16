public class Control_Monad_Except {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object withExcept = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> ( ((Boolean) ((v_1 instanceof Data_Either.Right))) ? new Data_Either.Right(((Data_Either.Right) (Object)(v_1)).value0) : ( ((Boolean) ((v_1 instanceof Data_Either.Left))) ? new Data_Either.Left(((java.util.function.Function<Object, Object>) (f_0)).apply(((Data_Either.Left) (Object)(v_1)).value0)) : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()));
public static final Object runExcept = (java.util.function.Function<Object, Object>) (x_0) -> x_0;
public static final Object mapExcept = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (v_1) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(v_1);
}
