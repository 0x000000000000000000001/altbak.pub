public class Data_List_Internal {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final class Leaf {
            
            public Leaf() {
                
            }
        }
public static final class Two {
            public final Object value0;
            public final Object value1;
            public final Object value2;
            public Two(Object value0, Object value1, Object value2) {
                this.value0 = value0;
                this.value1 = value1;
                this.value2 = value2;
            }
        }
public static final class Three {
            public final Object value0;
            public final Object value1;
            public final Object value2;
            public final Object value3;
            public final Object value4;
            public Three(Object value0, Object value1, Object value2, Object value3, Object value4) {
                this.value0 = value0;
                this.value1 = value1;
                this.value2 = value2;
                this.value3 = value3;
                this.value4 = value4;
            }
        }
public static final class TwoLeft {
            public final Object value0;
            public final Object value1;
            public TwoLeft(Object value0, Object value1) {
                this.value0 = value0;
                this.value1 = value1;
            }
        }
public static final class TwoRight {
            public final Object value0;
            public final Object value1;
            public TwoRight(Object value0, Object value1) {
                this.value0 = value0;
                this.value1 = value1;
            }
        }
public static final class ThreeLeft {
            public final Object value0;
            public final Object value1;
            public final Object value2;
            public final Object value3;
            public ThreeLeft(Object value0, Object value1, Object value2, Object value3) {
                this.value0 = value0;
                this.value1 = value1;
                this.value2 = value2;
                this.value3 = value3;
            }
        }
public static final class ThreeMiddle {
            public final Object value0;
            public final Object value1;
            public final Object value2;
            public final Object value3;
            public ThreeMiddle(Object value0, Object value1, Object value2, Object value3) {
                this.value0 = value0;
                this.value1 = value1;
                this.value2 = value2;
                this.value3 = value3;
            }
        }
public static final class ThreeRight {
            public final Object value0;
            public final Object value1;
            public final Object value2;
            public final Object value3;
            public ThreeRight(Object value0, Object value1, Object value2, Object value3) {
                this.value0 = value0;
                this.value1 = value1;
                this.value2 = value2;
                this.value3 = value3;
            }
        }
public static final class KickUp {
            public final Object value0;
            public final Object value1;
            public final Object value2;
            public KickUp(Object value0, Object value1, Object value2) {
                this.value0 = value0;
                this.value1 = value1;
                this.value2 = value2;
            }
        }
public static final Object Leaf = new Data_List_Internal.Leaf();
public static final Object Two = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> (java.util.function.Function<Object, Object>) (value2) -> new Data_List_Internal.Two(value0, value1, value2);
public static final Object Three = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> (java.util.function.Function<Object, Object>) (value2) -> (java.util.function.Function<Object, Object>) (value3) -> (java.util.function.Function<Object, Object>) (value4) -> new Data_List_Internal.Three(value0, value1, value2, value3, value4);
public static final Object TwoLeft = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> new Data_List_Internal.TwoLeft(value0, value1);
public static final Object TwoRight = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> new Data_List_Internal.TwoRight(value0, value1);
public static final Object ThreeLeft = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> (java.util.function.Function<Object, Object>) (value2) -> (java.util.function.Function<Object, Object>) (value3) -> new Data_List_Internal.ThreeLeft(value0, value1, value2, value3);
public static final Object ThreeMiddle = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> (java.util.function.Function<Object, Object>) (value2) -> (java.util.function.Function<Object, Object>) (value3) -> new Data_List_Internal.ThreeMiddle(value0, value1, value2, value3);
public static final Object ThreeRight = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> (java.util.function.Function<Object, Object>) (value2) -> (java.util.function.Function<Object, Object>) (value3) -> new Data_List_Internal.ThreeRight(value0, value1, value2, value3);
public static final Object KickUp = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> (java.util.function.Function<Object, Object>) (value2) -> new Data_List_Internal.KickUp(value0, value1, value2);
public static final Object fromZipper = (java.util.function.Function<Object, Object>) (v_0) -> (java.util.function.Function<Object, Object>) (v1_1) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_v_0 = v_0; Object __tco_v1_1 = v1_1; while(true) { final Object __final_v_0 = __tco_v_0; final Object __final_v1_1 = __tco_v1_1; try { return ( ((Boolean) ((__final_v_0 instanceof Data_List_Types.Nil))) ? __final_v1_1 : ( ((Boolean) ((__final_v_0 instanceof Data_List_Types.Cons))) ? ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_0)).value0 instanceof Data_List_Internal.TwoLeft))) ? ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_List_Internal.fromZipper)).apply(((Data_List_Types.Cons) (Object)(__final_v_0)).value1))).apply(new Data_List_Internal.Two(__final_v1_1, ((Data_List_Internal.TwoLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value0, ((Data_List_Internal.TwoLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value1)) : ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_0)).value0 instanceof Data_List_Internal.TwoRight))) ? ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_List_Internal.fromZipper)).apply(((Data_List_Types.Cons) (Object)(__final_v_0)).value1))).apply(new Data_List_Internal.Two(((Data_List_Internal.TwoRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value0, ((Data_List_Internal.TwoRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value1, __final_v1_1)) : ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_0)).value0 instanceof Data_List_Internal.ThreeLeft))) ? ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_List_Internal.fromZipper)).apply(((Data_List_Types.Cons) (Object)(__final_v_0)).value1))).apply(new Data_List_Internal.Three(__final_v1_1, ((Data_List_Internal.ThreeLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value0, ((Data_List_Internal.ThreeLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value1, ((Data_List_Internal.ThreeLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value2, ((Data_List_Internal.ThreeLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value3)) : ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_0)).value0 instanceof Data_List_Internal.ThreeMiddle))) ? ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_List_Internal.fromZipper)).apply(((Data_List_Types.Cons) (Object)(__final_v_0)).value1))).apply(new Data_List_Internal.Three(((Data_List_Internal.ThreeMiddle) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value0, ((Data_List_Internal.ThreeMiddle) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value1, __final_v1_1, ((Data_List_Internal.ThreeMiddle) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value2, ((Data_List_Internal.ThreeMiddle) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value3)) : ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_0)).value0 instanceof Data_List_Internal.ThreeRight))) ? ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_List_Internal.fromZipper)).apply(((Data_List_Types.Cons) (Object)(__final_v_0)).value1))).apply(new Data_List_Internal.Three(((Data_List_Internal.ThreeRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value0, ((Data_List_Internal.ThreeRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value1, ((Data_List_Internal.ThreeRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value2, ((Data_List_Internal.ThreeRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_0)).value0)).value3, __final_v1_1)) : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))))) : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())); } catch (TcoLoop __tco_ex) { __tco_v_0 = __tco_ex.args[0]; __tco_v1_1 = __tco_ex.args[1]; } } } }).get();
public static final Object insertAndLookupBy = (java.util.function.Function<Object, Object>) (comp_0) -> (java.util.function.Function<Object, Object>) (k_1) -> (java.util.function.Function<Object, Object>) (orig_2) -> { Object up_3 = (java.util.function.Function<Object, Object>) (v_4) -> (java.util.function.Function<Object, Object>) (v1_5) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_v_4 = v_4; Object __tco_v1_5 = v1_5; while(true) { final Object __final_v_4 = __tco_v_4; final Object __final_v1_5 = __tco_v1_5; try { return ( ((Boolean) ((__final_v_4 instanceof Data_List_Types.Nil))) ? new Data_List_Internal.Two(((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value0, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value1, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value2) : ( ((Boolean) ((__final_v_4 instanceof Data_List_Types.Cons))) ? ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_4)).value0 instanceof Data_List_Internal.TwoLeft))) ? ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_List_Internal.fromZipper)).apply(((Data_List_Types.Cons) (Object)(__final_v_4)).value1))).apply(new Data_List_Internal.Three(((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value0, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value1, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value2, ((Data_List_Internal.TwoLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value0, ((Data_List_Internal.TwoLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value1)) : ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_4)).value0 instanceof Data_List_Internal.TwoRight))) ? ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_List_Internal.fromZipper)).apply(((Data_List_Types.Cons) (Object)(__final_v_4)).value1))).apply(new Data_List_Internal.Three(((Data_List_Internal.TwoRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value0, ((Data_List_Internal.TwoRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value1, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value0, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value1, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value2)) : ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_4)).value0 instanceof Data_List_Internal.ThreeLeft))) ? (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("up_3", new Object[]{((Data_List_Types.Cons) (Object)(__final_v_4)).value1, new Data_List_Internal.KickUp(new Data_List_Internal.Two(((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value0, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value1, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value2), ((Data_List_Internal.ThreeLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value0, new Data_List_Internal.Two(((Data_List_Internal.ThreeLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value1, ((Data_List_Internal.ThreeLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value2, ((Data_List_Internal.ThreeLeft) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value3))}); } }).get() : ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_4)).value0 instanceof Data_List_Internal.ThreeMiddle))) ? (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("up_3", new Object[]{((Data_List_Types.Cons) (Object)(__final_v_4)).value1, new Data_List_Internal.KickUp(new Data_List_Internal.Two(((Data_List_Internal.ThreeMiddle) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value0, ((Data_List_Internal.ThreeMiddle) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value1, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value0), ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value1, new Data_List_Internal.Two(((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value2, ((Data_List_Internal.ThreeMiddle) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value2, ((Data_List_Internal.ThreeMiddle) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value3))}); } }).get() : ( ((Boolean) ((((Data_List_Types.Cons) (Object)(__final_v_4)).value0 instanceof Data_List_Internal.ThreeRight))) ? (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("up_3", new Object[]{((Data_List_Types.Cons) (Object)(__final_v_4)).value1, new Data_List_Internal.KickUp(new Data_List_Internal.Two(((Data_List_Internal.ThreeRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value0, ((Data_List_Internal.ThreeRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value1, ((Data_List_Internal.ThreeRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value2), ((Data_List_Internal.ThreeRight) (Object)(((Data_List_Types.Cons) (Object)(__final_v_4)).value0)).value3, new Data_List_Internal.Two(((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value0, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value1, ((Data_List_Internal.KickUp) (Object)(__final_v1_5)).value2))}); } }).get() : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()))))) : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())); } catch (TcoLoop __tco_ex) { __tco_v_4 = __tco_ex.args[0]; __tco_v1_5 = __tco_ex.args[1]; } } } }).get(); Object down_4 = (java.util.function.Function<Object, Object>) (v_5) -> (java.util.function.Function<Object, Object>) (v1_6) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_v_5 = v_5; Object __tco_v1_6 = v1_6; while(true) { final Object __final_v_5 = __tco_v_5; final Object __final_v1_6 = __tco_v1_6; try { return ( ((Boolean) ((__final_v1_6 instanceof Data_List_Internal.Leaf))) ? (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("found", false); __map.put("result", ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (up_3)).apply(__final_v_5))).apply(new Data_List_Internal.KickUp(new Data_List_Internal.Leaf(), k_1, new Data_List_Internal.Leaf())));  return __map; } }).get() : (new java.util.function.Supplier<Object>() { public Object get() {  return ( ((Boolean) ((__final_v1_6 instanceof Data_List_Internal.Two))) ? (new java.util.function.Supplier<Object>() { public Object get() { Object v2_7 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (comp_0)).apply(k_1))).apply(((Data_List_Internal.Two) (Object)(__final_v1_6)).value1); return ( ((Boolean) ((v2_7 instanceof Data_Ordering.EQ))) ? (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("found", true); __map.put("result", orig_2);  return __map; } }).get() : ( ((Boolean) ((v2_7 instanceof Data_Ordering.LT))) ? (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("down_4", new Object[]{new Data_List_Types.Cons(new Data_List_Internal.TwoLeft(((Data_List_Internal.Two) (Object)(__final_v1_6)).value1, ((Data_List_Internal.Two) (Object)(__final_v1_6)).value2), __final_v_5), ((Data_List_Internal.Two) (Object)(__final_v1_6)).value0}); } }).get() : (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("down_4", new Object[]{new Data_List_Types.Cons(new Data_List_Internal.TwoRight(((Data_List_Internal.Two) (Object)(__final_v1_6)).value0, ((Data_List_Internal.Two) (Object)(__final_v1_6)).value1), __final_v_5), ((Data_List_Internal.Two) (Object)(__final_v1_6)).value2}); } }).get())); } }).get() : (new java.util.function.Supplier<Object>() { public Object get() {  return ( ((Boolean) ((__final_v1_6 instanceof Data_List_Internal.Three))) ? (new java.util.function.Supplier<Object>() { public Object get() { Object v2_7 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (comp_0)).apply(k_1))).apply(((Data_List_Internal.Three) (Object)(__final_v1_6)).value1); Object v3_8 = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (comp_0)).apply(k_1))).apply(((Data_List_Internal.Three) (Object)(__final_v1_6)).value3); return ( ((Boolean) ((v2_7 instanceof Data_Ordering.EQ))) ? (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("found", true); __map.put("result", orig_2);  return __map; } }).get() : ( ((Boolean) ((v3_8 instanceof Data_Ordering.EQ))) ? (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("found", true); __map.put("result", orig_2);  return __map; } }).get() : ( ((Boolean) ((v2_7 instanceof Data_Ordering.LT))) ? (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("down_4", new Object[]{new Data_List_Types.Cons(new Data_List_Internal.ThreeLeft(((Data_List_Internal.Three) (Object)(__final_v1_6)).value1, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value2, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value3, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value4), __final_v_5), ((Data_List_Internal.Three) (Object)(__final_v1_6)).value0}); } }).get() : ( ((Boolean) ((((Boolean) ((v2_7 instanceof Data_Ordering.GT))) && ((Boolean) ((v3_8 instanceof Data_Ordering.LT)))))) ? (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("down_4", new Object[]{new Data_List_Types.Cons(new Data_List_Internal.ThreeMiddle(((Data_List_Internal.Three) (Object)(__final_v1_6)).value0, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value1, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value3, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value4), __final_v_5), ((Data_List_Internal.Three) (Object)(__final_v1_6)).value2}); } }).get() : (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("down_4", new Object[]{new Data_List_Types.Cons(new Data_List_Internal.ThreeRight(((Data_List_Internal.Three) (Object)(__final_v1_6)).value0, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value1, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value2, ((Data_List_Internal.Three) (Object)(__final_v1_6)).value3), __final_v_5), ((Data_List_Internal.Three) (Object)(__final_v1_6)).value4}); } }).get())))); } }).get() : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get()); } }).get()); } }).get()); } catch (TcoLoop __tco_ex) { __tco_v_5 = __tco_ex.args[0]; __tco_v1_6 = __tco_ex.args[1]; } } } }).get(); return ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (down_4)).apply(new Data_List_Types.Nil()))).apply(orig_2); };
public static final Object emptySet = new Data_List_Internal.Leaf();
}
