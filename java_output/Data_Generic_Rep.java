public class Data_Generic_Rep {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final class Inl {
            public final Object value0;
            public Inl(Object value0) {
                this.value0 = value0;
            }
        }
public static final class Inr {
            public final Object value0;
            public Inr(Object value0) {
                this.value0 = value0;
            }
        }
public static final class Product {
            public final Object value0;
            public final Object value1;
            public Product(Object value0, Object value1) {
                this.value0 = value0;
                this.value1 = value1;
            }
        }
public static final class NoArguments {
            
            public NoArguments() {
                
            }
        }
public static final Object Inl = (java.util.function.Function<Object, Object>) (value0) -> new Data_Generic_Rep.Inl(value0);
public static final Object Inr = (java.util.function.Function<Object, Object>) (value0) -> new Data_Generic_Rep.Inr(value0);
public static final Object Product = (java.util.function.Function<Object, Object>) (value0) -> (java.util.function.Function<Object, Object>) (value1) -> new Data_Generic_Rep.Product(value0, value1);
public static final Object NoArguments = new Data_Generic_Rep.NoArguments();
public static final Object Constructor = (java.util.function.Function<Object, Object>) (x_0) -> x_0;
public static final Object Argument = (java.util.function.Function<Object, Object>) (x_0) -> x_0;
public static final Object to = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("to");
public static final Object showSum = (java.util.function.Function<Object, Object>) (dictShow_0) -> (java.util.function.Function<Object, Object>) (dictShow1_1) -> (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", (java.util.function.Function<Object, Object>) (v_2) -> ( ((Boolean) ((v_2 instanceof Data_Generic_Rep.Inl))) ? (((String) ((((String) ("(Inl ")) + ((String) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(((Data_Generic_Rep.Inl) (Object)(v_2)).value0)))))) + ((String) (")"))) : ( ((Boolean) ((v_2 instanceof Data_Generic_Rep.Inr))) ? (((String) ((((String) ("(Inr ")) + ((String) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow1_1).get("show"))).apply(((Data_Generic_Rep.Inr) (Object)(v_2)).value0)))))) + ((String) (")"))) : (new java.util.function.Supplier<Object>() { public Object get() { throw new RuntimeException("Failed pattern match"); } }).get())));  return __map; } }).get();
public static final Object showProduct = (java.util.function.Function<Object, Object>) (dictShow_0) -> (java.util.function.Function<Object, Object>) (dictShow1_1) -> (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", (java.util.function.Function<Object, Object>) (v_2) -> (((String) ((((String) ((((String) ((((String) ("(Product ")) + ((String) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(((Data_Generic_Rep.Product) (Object)(v_2)).value0)))))) + ((String) (" "))))) + ((String) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow1_1).get("show"))).apply(((Data_Generic_Rep.Product) (Object)(v_2)).value1)))))) + ((String) (")"))));  return __map; } }).get();
public static final Object showNoArguments = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", (java.util.function.Function<Object, Object>) (v_0) -> "NoArguments");  return __map; } }).get();
public static final Object showConstructor = (java.util.function.Function<Object, Object>) (dictIsSymbol_0) -> (java.util.function.Function<Object, Object>) (dictShow_1) -> (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", (java.util.function.Function<Object, Object>) (v_2) -> (((String) ((((String) ((((String) ((((String) ("(Constructor @")) + ((String) (((java.util.function.Function<Object, Object>) (Data_Show.showStringImpl)).apply(((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictIsSymbol_0).get("reflectSymbol"))).apply(new Type_Proxy.Proxy()))))))) + ((String) (" "))))) + ((String) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_1).get("show"))).apply(v_2)))))) + ((String) (")"))));  return __map; } }).get();
public static final Object showArgument = (java.util.function.Function<Object, Object>) (dictShow_0) -> (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", (java.util.function.Function<Object, Object>) (v_1) -> (((String) ((((String) ("(Argument ")) + ((String) (((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(v_1)))))) + ((String) (")"))));  return __map; } }).get();
public static final Object repOf = (java.util.function.Function<Object, Object>) (dictGeneric_0) -> (java.util.function.Function<Object, Object>) (v_1) -> new Type_Proxy.Proxy();
public static final Object from = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("from");
}
