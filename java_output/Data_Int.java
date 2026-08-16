public class Data_Int {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    // FFI provided by ../javapurs/javapurs-integers/src/Data/Int.java
    public static Object fromNumberImpl = (java.util.function.Function<Object, Object>) (just) -> (java.util.function.Function<Object, Object>) (nothing) -> (java.util.function.Function<Object, Object>) (n) -> {
        Double d = (Double) n;
        if (d.intValue() == d) {
            return ((java.util.function.Function<Object, Object>) just).apply(d.intValue());
        }
        return nothing;
    };

    public static Object toNumber = (java.util.function.Function<Object, Object>) (n) -> {
        return ((Integer) n).doubleValue();
    };

    public static Object fromStringAsImpl = (java.util.function.Function<Object, Object>) (just) -> (java.util.function.Function<Object, Object>) (nothing) -> (java.util.function.Function<Object, Object>) (radixObj) -> {
        int radix = (Integer) radixObj;
        return (java.util.function.Function<Object, Object>) (s) -> {
            try {
                int i = Integer.parseInt((String) s, radix);
                return ((java.util.function.Function<Object, Object>) just).apply(i);
            } catch (NumberFormatException e) {
                return nothing;
            }
        };
    };

    public static Object toStringAs = (java.util.function.Function<Object, Object>) (radixObj) -> (java.util.function.Function<Object, Object>) (i) -> {
        int radix = (Integer) radixObj;
        return Integer.toString((Integer) i, radix);
    };

    public static Object quot = (java.util.function.Function<Object, Object>) (x) -> (java.util.function.Function<Object, Object>) (y) -> {
        return (Integer) x / (Integer) y;
    };

    public static Object rem = (java.util.function.Function<Object, Object>) (x) -> (java.util.function.Function<Object, Object>) (y) -> {
        return (Integer) x % (Integer) y;
    };

    public static Object pow = (java.util.function.Function<Object, Object>) (x) -> (java.util.function.Function<Object, Object>) (y) -> {
        return (int) Math.pow((Integer) x, (Integer) y);
    };


public static final class Even {
            
            public Even() {
                
            }
        }
public static final class Odd {
            
            public Odd() {
                
            }
        }
public static final Object Even = new Data_Int.Even();
public static final Object Odd = new Data_Int.Odd();
public static final Object showParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", (java.util.function.Function<Object, Object>) (v_0) -> ( ((Boolean) ((v_0 instanceof Data_Int.Even))) ? "Even" : ( ((Boolean) ((v_0 instanceof Data_Int.Odd))) ? "Odd" : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())));  return __map; } }).get();
public static final Object radix = (java.util.function.Function<Object, Object>) (n_0) -> ( ((Boolean) ((((Boolean) ((((Integer) (n_0)) >= ((Integer) (2))))) && ((Boolean) ((((Integer) (n_0)) <= ((Integer) (36)))))))) ? new Data_Maybe.Just(n_0) : new Data_Maybe.Nothing());
public static final Object odd = (java.util.function.Function<Object, Object>) (x_0) -> !(((Boolean) (java.util.Objects.equals((((Integer) (x_0)) & ((Integer) (1))), 0))));
public static final Object octal = 8;
public static final Object hexadecimal = 16;
public static final Object fromStringAs = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Int.fromStringAsImpl)).apply(Data_Maybe.Just))).apply(new Data_Maybe.Nothing());
public static final Object fromString = ((java.util.function.Function<Object, Object>) (Data_Int.fromStringAs)).apply(10);
public static final Object fromNumber = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Int.fromNumberImpl)).apply(Data_Maybe.Just))).apply(new Data_Maybe.Nothing());
public static final Object unsafeClamp = (java.util.function.Function<Object, Object>) (x_0) -> ( ((Boolean) (!(((Boolean) (((java.util.function.Function<Object, Object>) (Data_Number.isFinite)).apply(x_0)))))) ? 0 : ( ((Boolean) (!(((Boolean) ((((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Ord.ordNumberImpl)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT()))).apply(x_0))).apply(((java.util.function.Function<Object, Object>) (Data_Int.toNumber)).apply(Data_Bounded.topInt)) instanceof Data_Ordering.LT)))))) ? Data_Bounded.topInt : ( ((Boolean) (!(((Boolean) ((((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_Ord.ordNumberImpl)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT()))).apply(x_0))).apply(((java.util.function.Function<Object, Object>) (Data_Int.toNumber)).apply(Data_Bounded.bottomInt)) instanceof Data_Ordering.GT)))))) ? Data_Bounded.bottomInt : (new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_1 = ((java.util.function.Function<Object, Object>) (Data_Int.fromNumber)).apply(x_0); return ( ((Boolean) ((__local_var_1 instanceof Data_Maybe.Nothing))) ? 0 : ( ((Boolean) ((__local_var_1 instanceof Data_Maybe.Just))) ? ((Data_Maybe.Just) (Object)(__local_var_1)).value0 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())); } }).get())));
public static final Object round = (java.util.function.Function<Object, Object>) (x_0) -> ((java.util.function.Function<Object, Object>) (Data_Int.unsafeClamp)).apply(((java.util.function.Function<Object, Object>) (Data_Number.round)).apply(x_0));
public static final Object trunc = (java.util.function.Function<Object, Object>) (x_0) -> ((java.util.function.Function<Object, Object>) (Data_Int.unsafeClamp)).apply(((java.util.function.Function<Object, Object>) (Data_Number.trunc)).apply(x_0));
public static final Object floor = (java.util.function.Function<Object, Object>) (x_0) -> ((java.util.function.Function<Object, Object>) (Data_Int.unsafeClamp)).apply(((java.util.function.Function<Object, Object>) (Data_Number.floor)).apply(x_0));
public static final Object even = (java.util.function.Function<Object, Object>) (x_0) -> java.util.Objects.equals((((Integer) (x_0)) & ((Integer) (1))), 0);
public static final Object parity = (java.util.function.Function<Object, Object>) (n_0) -> ( ((Boolean) (java.util.Objects.equals((((Integer) (n_0)) & ((Integer) (1))), 0))) ? new Data_Int.Even() : new Data_Int.Odd());
public static final Object eqParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("eq", (java.util.function.Function<Object, Object>) (x_0) -> (java.util.function.Function<Object, Object>) (y_1) -> ( ((Boolean) ((x_0 instanceof Data_Int.Even))) ? (y_1 instanceof Data_Int.Even) : (((Boolean) ((x_0 instanceof Data_Int.Odd))) && ((Boolean) ((y_1 instanceof Data_Int.Odd))))));  return __map; } }).get();
public static final Object ordParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("compare", (java.util.function.Function<Object, Object>) (x_0) -> (java.util.function.Function<Object, Object>) (y_1) -> ( ((Boolean) ((x_0 instanceof Data_Int.Even))) ? ( ((Boolean) ((y_1 instanceof Data_Int.Even))) ? new Data_Ordering.EQ() : new Data_Ordering.LT()) : ( ((Boolean) ((y_1 instanceof Data_Int.Even))) ? new Data_Ordering.GT() : ( ((Boolean) ((((Boolean) ((x_0 instanceof Data_Int.Odd))) && ((Boolean) ((y_1 instanceof Data_Int.Odd)))))) ? new Data_Ordering.EQ() : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())))); __map.put("Eq0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Data_Int.eqParity);  return __map; } }).get();
public static final Object semiringParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("zero", new Data_Int.Even()); __map.put("add", (java.util.function.Function<Object, Object>) (x_0) -> (java.util.function.Function<Object, Object>) (y_1) -> ( ((Boolean) (( ((Boolean) ((x_0 instanceof Data_Int.Even))) ? (y_1 instanceof Data_Int.Even) : (((Boolean) ((x_0 instanceof Data_Int.Odd))) && ((Boolean) ((y_1 instanceof Data_Int.Odd))))))) ? new Data_Int.Even() : new Data_Int.Odd())); __map.put("one", new Data_Int.Odd()); __map.put("mul", (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (v1_1) -> ( ((Boolean) ((((Boolean) ((v_0 instanceof Data_Int.Odd))) && ((Boolean) ((v1_1 instanceof Data_Int.Odd)))))) ? new Data_Int.Odd() : new Data_Int.Even()));  return __map; } }).get();
public static final Object ringParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("sub", (java.util.function.Function<Object, Object>) (x_0) -> (java.util.function.Function<Object, Object>) (y_1) -> ( ((Boolean) (( ((Boolean) ((x_0 instanceof Data_Int.Even))) ? (y_1 instanceof Data_Int.Even) : (((Boolean) ((x_0 instanceof Data_Int.Odd))) && ((Boolean) ((y_1 instanceof Data_Int.Odd))))))) ? new Data_Int.Even() : new Data_Int.Odd())); __map.put("Semiring0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Data_Int.semiringParity);  return __map; } }).get();
public static final Object divisionRingParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("recip", (java.util.function.Function<Object, Object>) (x_0) -> x_0); __map.put("Ring0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Data_Int.ringParity);  return __map; } }).get();
public static final Object decimal = 10;
public static final Object commutativeRingParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("Ring0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Data_Int.ringParity);  return __map; } }).get();
public static final Object euclideanRingParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("degree", (java.util.function.Function<Object, Object>) (v_0) -> ( ((Boolean) ((v_0 instanceof Data_Int.Even))) ? 0 : ( ((Boolean) ((v_0 instanceof Data_Int.Odd))) ? 1 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))); __map.put("div", (java.util.function.Function<Object, Object>) (x_0) -> (java.util.function.Function<Object, Object>) (v_1) -> x_0); __map.put("mod", (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (v1_1) -> new Data_Int.Even()); __map.put("CommutativeRing0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Data_Int.commutativeRingParity);  return __map; } }).get();
public static final Object ceil = (java.util.function.Function<Object, Object>) (x_0) -> ((java.util.function.Function<Object, Object>) (Data_Int.unsafeClamp)).apply(((java.util.function.Function<Object, Object>) (Data_Number.ceil)).apply(x_0));
public static final Object boundedParity = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("bottom", new Data_Int.Even()); __map.put("top", new Data_Int.Odd()); __map.put("Ord0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Data_Int.ordParity);  return __map; } }).get();
public static final Object binary = 2;
public static final Object base36 = 36;
}
