public class Effect_Class {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };


public static final Object monadEffectEffect = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("liftEffect", (java.util.function.Function<Object, Object>) (x_0) -> x_0); __map.put("Monad0", (java.util.function.Function<Object, Object>) (_dollar___unused_0) -> Effect.monadEffect);  return __map; } }).get();
public static final Object liftEffect = (java.util.function.Function<Object, Object>) (dict_0) -> ((java.util.LinkedHashMap<String, Object>) dict_0).get("liftEffect");
}
