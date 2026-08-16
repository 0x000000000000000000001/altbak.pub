public class Test_Tiny {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final class Circle {
            public final Object value0;
            public Circle(Object value0) {
                this.value0 = value0;
            }
        }
public static final class Rect {
            public final Object value0;
            public final Object value1;
            public Rect(Object value0, Object value1) {
                this.value0 = value0;
                this.value1 = value1;
            }
        }
public static final Object Circle = (java.util.function.Function<Object, Object>) (value0) -> new Test_Tiny.Circle(value0);
public static final Object Rect = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> new Test_Tiny.Rect(value0, value1);
public static final Object area = (java.util.function.Function<Object, Object>) (v_0) -> ( ((Boolean) ((v_0 instanceof Test_Tiny.Circle))) ? (((Integer) (((Test_Tiny.Circle) (Object)(v_0)).value0)) * ((Integer) (((Test_Tiny.Circle) (Object)(v_0)).value0))) : ( ((Boolean) ((v_0 instanceof Test_Tiny.Rect))) ? (((Integer) (((Test_Tiny.Rect) (Object)(v_0)).value0)) * ((Integer) (((Test_Tiny.Rect) (Object)(v_0)).value1))) : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()));
}
