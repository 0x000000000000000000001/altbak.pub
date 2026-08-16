public class Data_Number_Format {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object toExponentialNative = FFI_STUB;
    public static Object toExponentialNative(Object... args) { return null; }
    public static Object toFixedNative = FFI_STUB;
    public static Object toFixedNative(Object... args) { return null; }
    public static Object toPrecisionNative = FFI_STUB;
    public static Object toPrecisionNative(Object... args) { return null; }
    public static Object toString = FFI_STUB;
    public static Object toString(Object... args) { return null; }

public static final class Precision {
            public final Object value0;
            public Precision(Object value0) {
                this.value0 = value0;
            }
        }
public static final class Fixed {
            public final Object value0;
            public Fixed(Object value0) {
                this.value0 = value0;
            }
        }
public static final class Exponential {
            public final Object value0;
            public Exponential(Object value0) {
                this.value0 = value0;
            }
        }
public static final Object Precision = (java.util.function.Function<Object, Object>) (value0) -> new Data_Number_Format.Precision(value0);
public static final Object Fixed = (java.util.function.Function<Object, Object>) (value0) -> new Data_Number_Format.Fixed(value0);
public static final Object Exponential = (java.util.function.Function<Object, Object>) (value0) -> new Data_Number_Format.Exponential(value0);
public static final Object toStringWith = (java.util.function.Function<Object, Object>) (v_0) -> ( ((Boolean) ((v_0 instanceof Data_Number_Format.Precision))) ? ((java.util.function.Function<Object, Object>) (Data_Number_Format.toPrecisionNative)).apply(((Data_Number_Format.Precision) (Object)(v_0)).value0) : ( ((Boolean) ((v_0 instanceof Data_Number_Format.Fixed))) ? ((java.util.function.Function<Object, Object>) (Data_Number_Format.toFixedNative)).apply(((Data_Number_Format.Fixed) (Object)(v_0)).value0) : ( ((Boolean) ((v_0 instanceof Data_Number_Format.Exponential))) ? ((java.util.function.Function<Object, Object>) (Data_Number_Format.toExponentialNative)).apply(((Data_Number_Format.Exponential) (Object)(v_0)).value0) : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())));
public static final Object precision = (java.util.function.Function<Object, Object>) (x_0) -> new Data_Number_Format.Precision((new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_1 = (new java.util.function.Supplier<Object>() { public Object get() { Object v_1 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Ord.ordIntImpl)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT()))).apply(1))).apply(x_0); return ( ((Boolean) ((v_1 instanceof Data_Ordering.LT))) ? x_0 : ( ((Boolean) ((v_1 instanceof Data_Ordering.EQ))) ? 1 : ( ((Boolean) ((v_1 instanceof Data_Ordering.GT))) ? 1 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))); } }).get(); Object v_2 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Ord.ordIntImpl)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT()))).apply(21))).apply(__local_var_1); return ( ((Boolean) ((v_2 instanceof Data_Ordering.LT))) ? 21 : ( ((Boolean) ((v_2 instanceof Data_Ordering.EQ))) ? 21 : ( ((Boolean) ((v_2 instanceof Data_Ordering.GT))) ? __local_var_1 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))); } }).get());
public static final Object fixed = (java.util.function.Function<Object, Object>) (x_0) -> new Data_Number_Format.Fixed((new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_1 = (new java.util.function.Supplier<Object>() { public Object get() { Object v_1 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Ord.ordIntImpl)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT()))).apply(0))).apply(x_0); return ( ((Boolean) ((v_1 instanceof Data_Ordering.LT))) ? x_0 : ( ((Boolean) ((v_1 instanceof Data_Ordering.EQ))) ? 0 : ( ((Boolean) ((v_1 instanceof Data_Ordering.GT))) ? 0 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))); } }).get(); Object v_2 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Ord.ordIntImpl)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT()))).apply(20))).apply(__local_var_1); return ( ((Boolean) ((v_2 instanceof Data_Ordering.LT))) ? 20 : ( ((Boolean) ((v_2 instanceof Data_Ordering.EQ))) ? 20 : ( ((Boolean) ((v_2 instanceof Data_Ordering.GT))) ? __local_var_1 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))); } }).get());
public static final Object exponential = (java.util.function.Function<Object, Object>) (x_0) -> new Data_Number_Format.Exponential((new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_1 = (new java.util.function.Supplier<Object>() { public Object get() { Object v_1 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Ord.ordIntImpl)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT()))).apply(0))).apply(x_0); return ( ((Boolean) ((v_1 instanceof Data_Ordering.LT))) ? x_0 : ( ((Boolean) ((v_1 instanceof Data_Ordering.EQ))) ? 0 : ( ((Boolean) ((v_1 instanceof Data_Ordering.GT))) ? 0 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))); } }).get(); Object v_2 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Ord.ordIntImpl)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT()))).apply(20))).apply(__local_var_1); return ( ((Boolean) ((v_2 instanceof Data_Ordering.LT))) ? 20 : ( ((Boolean) ((v_2 instanceof Data_Ordering.EQ))) ? 20 : ( ((Boolean) ((v_2 instanceof Data_Ordering.GT))) ? __local_var_1 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))); } }).get());
}
