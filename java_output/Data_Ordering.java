public class Data_Ordering {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final class LT {
            
            public LT() {
                
            }
        }
public static final class GT {
            
            public GT() {
                
            }
        }
public static final class EQ {
            
            public EQ() {
                
            }
        }
public static final Object LT = new Data_Ordering.LT();
public static final Object GT = new Data_Ordering.GT();
public static final Object EQ = new Data_Ordering.EQ();
public static final Object showOrdering = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", (java.util.function.Function<Object, Object>) (v_0) -> ( ((Boolean) ((v_0 instanceof Data_Ordering.LT))) ? "LT" : ( ((Boolean) ((v_0 instanceof Data_Ordering.GT))) ? "GT" : ( ((Boolean) ((v_0 instanceof Data_Ordering.EQ))) ? "EQ" : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))));  return __map; } }).get();
public static final Object semigroupOrdering = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("append", (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (v1_1) -> ( ((Boolean) ((v_0 instanceof Data_Ordering.LT))) ? new Data_Ordering.LT() : ( ((Boolean) ((v_0 instanceof Data_Ordering.GT))) ? new Data_Ordering.GT() : ( ((Boolean) ((v_0 instanceof Data_Ordering.EQ))) ? v1_1 : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))));  return __map; } }).get();
public static final Object invert = (java.util.function.Function<Object, Object>) (v_0) -> ( ((Boolean) ((v_0 instanceof Data_Ordering.GT))) ? new Data_Ordering.LT() : ( ((Boolean) ((v_0 instanceof Data_Ordering.EQ))) ? new Data_Ordering.EQ() : ( ((Boolean) ((v_0 instanceof Data_Ordering.LT))) ? new Data_Ordering.GT() : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())));
public static final Object eqOrdering = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("eq", (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (v1_1) -> ( ((Boolean) ((v_0 instanceof Data_Ordering.LT))) ? (v1_1 instanceof Data_Ordering.LT) : ( ((Boolean) ((v_0 instanceof Data_Ordering.GT))) ? (v1_1 instanceof Data_Ordering.GT) : (((Boolean) ((v_0 instanceof Data_Ordering.EQ))) && ((Boolean) ((v1_1 instanceof Data_Ordering.EQ)))))));  return __map; } }).get();
}
