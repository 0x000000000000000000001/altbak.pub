public class Test_AstTree {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final class Val {
            public final Object value0;
            public Val(Object value0) {
                this.value0 = value0;
            }
        }
public static final class Add {
            public final Object value0;
            public final Object value1;
            public Add(Object value0, Object value1) {
                this.value0 = value0;
                this.value1 = value1;
            }
        }
public static final class Mul {
            public final Object value0;
            public final Object value1;
            public Mul(Object value0, Object value1) {
                this.value0 = value0;
                this.value1 = value1;
            }
        }
public static final class Sub {
            public final Object value0;
            public final Object value1;
            public Sub(Object value0, Object value1) {
                this.value0 = value0;
                this.value1 = value1;
            }
        }
public static final Object Val = (java.util.function.Function<Object, Object>) (value0) -> new Test_AstTree.Val(value0);
public static final Object Add = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> new Test_AstTree.Add(value0, value1);
public static final Object Mul = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> new Test_AstTree.Mul(value0, value1);
public static final Object Sub = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> new Test_AstTree.Sub(value0, value1);
public static final Object eval = (java.util.function.Function<Object, Object>) (v_0) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_v_0 = v_0; while(true) { final Object __final_v_0 = __tco_v_0; try { return ( ((Boolean) ((__final_v_0 instanceof Test_AstTree.Val))) ? ((Test_AstTree.Val) (Object)(__final_v_0)).value0 : ( ((Boolean) ((__final_v_0 instanceof Test_AstTree.Add))) ? (((Integer) (((java.util.function.Function<Object, Object>) (Test_AstTree.eval)).apply(((Test_AstTree.Add) (Object)(__final_v_0)).value0))) + ((Integer) (((java.util.function.Function<Object, Object>) (Test_AstTree.eval)).apply(((Test_AstTree.Add) (Object)(__final_v_0)).value1)))) : ( ((Boolean) ((__final_v_0 instanceof Test_AstTree.Mul))) ? (((Integer) (((java.util.function.Function<Object, Object>) (Test_AstTree.eval)).apply(((Test_AstTree.Mul) (Object)(__final_v_0)).value0))) * ((Integer) (((java.util.function.Function<Object, Object>) (Test_AstTree.eval)).apply(((Test_AstTree.Mul) (Object)(__final_v_0)).value1)))) : ( ((Boolean) ((__final_v_0 instanceof Test_AstTree.Sub))) ? (((Integer) (((java.util.function.Function<Object, Object>) (Test_AstTree.eval)).apply(((Test_AstTree.Sub) (Object)(__final_v_0)).value0))) - ((Integer) (((java.util.function.Function<Object, Object>) (Test_AstTree.eval)).apply(((Test_AstTree.Sub) (Object)(__final_v_0)).value1)))) : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())))); } catch (TcoLoop __tco_ex) { __tco_v_0 = __tco_ex.args[0]; } } } }).get();
public static final Object describe = ((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) (arg) -> (java.util.function.Supplier<Object>) () -> { System.out.println(arg); return null; })).apply("AST Evaluation:");
public static final Object buildTree = (java.util.function.Function<Object, Object>) (v_0) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __tco_v_0 = v_0; while(true) { final Object __final_v_0 = __tco_v_0; try { return ( ((Boolean) (java.util.Objects.equals(__final_v_0, 0))) ? new Test_AstTree.Val(1) : new Test_AstTree.Add(new Test_AstTree.Mul(new Test_AstTree.Val(__final_v_0), ((java.util.function.Function<Object, Object>) (Test_AstTree.buildTree)).apply((((Integer) (__final_v_0)) - ((Integer) (1))))), new Test_AstTree.Sub(((java.util.function.Function<Object, Object>) (Test_AstTree.buildTree)).apply((((Integer) (__final_v_0)) - ((Integer) (1)))), new Test_AstTree.Val(1)))); } catch (TcoLoop __tco_ex) { __tco_v_0 = __tco_ex.args[0]; } } } }).get();
public static final Object act = (new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_0 = ((java.util.function.Function<Object, Object>) (Bench.opaque)).apply(3); Object dummy_1 = ((java.util.function.Supplier) (Object)(__local_var_0)).get(); return ((java.util.function.Supplier) (Object)(((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) (arg) -> (java.util.function.Supplier<Object>) () -> { System.out.println(arg); return null; })).apply(((java.util.function.Function<Object, Object>) (Data_Show.showIntImpl)).apply(((java.util.function.Function<Object, Object>) (Test_AstTree.eval)).apply(((java.util.function.Function<Object, Object>) (Test_AstTree.buildTree)).apply(dummy_1)))))).get(); } });
}
