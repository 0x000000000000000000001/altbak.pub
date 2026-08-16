public class Data_String_Regex_Unsafe {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object unsafeRegex = (java.util.function.Function<Object, Object>) (s_0) -> (java.util.function.Function<Object, Object>) (f_1) -> { Object __local_var_2 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_Regex.regex)).apply(s_0))).apply(f_1); return ( ((Boolean) ((__local_var_2 instanceof Data_Either.Left))) ? ((java.util.function.Function<Object, Object>) (Partial._crashWith)).apply(((Data_Either.Left) (Object)(__local_var_2)).value0) : ( ((Boolean) ((__local_var_2 instanceof Data_Either.Right))) ? ((Data_Either.Right) (Object)(__local_var_2)).value0 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())); };
}
