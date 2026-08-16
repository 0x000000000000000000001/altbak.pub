public class Effect_Console {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    public static Object clear = FFI_STUB;
    public static Object clear(Object... args) { return null; }
    public static Object debug = FFI_STUB;
    public static Object debug(Object... args) { return null; }
    public static Object error = FFI_STUB;
    public static Object error(Object... args) { return null; }
    public static Object group = FFI_STUB;
    public static Object group(Object... args) { return null; }
    public static Object groupCollapsed = FFI_STUB;
    public static Object groupCollapsed(Object... args) { return null; }
    public static Object groupEnd = FFI_STUB;
    public static Object groupEnd(Object... args) { return null; }
    public static Object info = FFI_STUB;
    public static Object info(Object... args) { return null; }
    public static Object log = FFI_STUB;
    public static Object log(Object... args) { return null; }
    public static Object time = FFI_STUB;
    public static Object time(Object... args) { return null; }
    public static Object timeEnd = FFI_STUB;
    public static Object timeEnd(Object... args) { return null; }
    public static Object timeLog = FFI_STUB;
    public static Object timeLog(Object... args) { return null; }
    public static Object warn = FFI_STUB;
    public static Object warn(Object... args) { return null; }

public static final Object warnShow = (java.util.function.Function<Object, Object>) (dictShow_0) -> (java.util.function.Function<Object, Object>) (a_1) -> ((java.util.function.Function<Object, Object>) (Effect_Console.warn)).apply(((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(a_1));
public static final Object logShow = (java.util.function.Function<Object, Object>) (dictShow_0) -> (java.util.function.Function<Object, Object>) (a_1) -> ((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) (arg) -> (java.util.function.Supplier<Object>) () -> { System.out.println(arg); return null; })).apply(((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(a_1));
public static final Object infoShow = (java.util.function.Function<Object, Object>) (dictShow_0) -> (java.util.function.Function<Object, Object>) (a_1) -> ((java.util.function.Function<Object, Object>) (Effect_Console.info)).apply(((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(a_1));
public static final Object grouped = (java.util.function.Function<Object, Object>) (name_0) -> (java.util.function.Function<Object, Object>) (inner_1) -> (new java.util.function.Supplier<Object>() { public Object get() { Object __local_var_2 = ((java.util.function.Function<Object, Object>) (Effect_Console.group)).apply(name_0); Object _dollar___unused_3 = ((java.util.function.Supplier) (Object)(__local_var_2)).get(); Object result_4 = ((java.util.function.Supplier) (Object)(inner_1)).get(); Object _dollar___unused_5 = ((java.util.function.Supplier) (Object)(Effect_Console.groupEnd)).get(); return result_4; } });
public static final Object errorShow = (java.util.function.Function<Object, Object>) (dictShow_0) -> (java.util.function.Function<Object, Object>) (a_1) -> ((java.util.function.Function<Object, Object>) (Effect_Console.error)).apply(((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(a_1));
public static final Object debugShow = (java.util.function.Function<Object, Object>) (dictShow_0) -> (java.util.function.Function<Object, Object>) (a_1) -> ((java.util.function.Function<Object, Object>) (Effect_Console.debug)).apply(((java.util.function.Function<Object, Object>) (((java.util.LinkedHashMap<String, Object>) dictShow_0).get("show"))).apply(a_1));
}
