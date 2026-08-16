public class TestLet {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    // FFI provided by src/TestLet.java
public static final java.util.function.Function<Object, Object> opaque = (xObj) -> xObj;


public static final Object addOne = (java.util.function.Function<Object, Object>) (x_0) -> { Object y_1 = ((java.util.function.Function<Object, Object>) (TestLet.opaque)).apply(x_0); Object z_2 = (((Integer) (y_1)) + ((Integer) (y_1))); return (((Integer) (z_2)) + ((Integer) (z_2))); };
}
