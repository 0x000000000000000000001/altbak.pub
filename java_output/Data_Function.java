public class Data_Function {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object on = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (g_1) -> (java.util.function.Function<Object, Object>) (x_2) -> (java.util.function.Function<Object, Object>) (y_3) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (f_0)).apply(((java.util.function.Function<Object, Object>) (g_1)).apply(x_2)))).apply(((java.util.function.Function<Object, Object>) (g_1)).apply(y_3));
public static final Object flip = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (b_1) -> (java.util.function.Function<Object, Object>) (a_2) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (f_0)).apply(a_2))).apply(b_1);
public static final Object $const = (java.util.function.Function<Object, Object>) (a_0) -> (java.util.function.Function<Object, Object>) (v_1) -> a_0;
public static final Object applyN = (java.util.function.Function<Object, Object>) (f_0) -> { Object go__go_1 = (java.util.function.Function<Object, Object>) (n_2) -> (java.util.function.Function<Object, Object>) (acc_3) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_n_2 = n_2; Object __tco_acc_3 = acc_3; while(true) { final Object __final_n_2 = __tco_n_2; final Object __final_acc_3 = __tco_acc_3; try { return ( ((Boolean) ((((Integer) (__final_n_2)) <= ((Integer) (0))))) ? __final_acc_3 : (new java.util.function.Supplier<Object>() { public Object get() { throw new TcoLoop("go__go_1", new Object[]{(((Integer) (__final_n_2)) - ((Integer) (1))), ((java.util.function.Function<Object, Object>) (f_0)).apply(__final_acc_3)}); } }).get()); } catch (TcoLoop __tco_ex) { __tco_n_2 = __tco_ex.args[0]; __tco_acc_3 = __tco_ex.args[1]; } } } }).get(); return go__go_1; };
public static final Object applyFlipped = (java.util.function.Function<Object, Object>) (x_0) -> (java.util.function.Function<Object, Object>) (f_1) -> ((java.util.function.Function<Object, Object>) (f_1)).apply(x_0);
public static final Object apply = (java.util.function.Function<Object, Object>) (f_0) -> (java.util.function.Function<Object, Object>) (x_1) -> ((java.util.function.Function<Object, Object>) (f_0)).apply(x_1);
}
