public class Data_String_Unsafe {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    // FFI provided by ../javapurs/javapurs-strings/src/Data/String/Unsafe.java
    public static Object charAt = (java.util.function.Function<Object, Object>) (i) -> (java.util.function.Function<Object, Object>) (s) -> {
        int idx = (Integer) i;
        String str = (String) s;
        if (idx >= 0 && idx < str.length()) return String.valueOf(str.charAt(idx));
        throw new RuntimeException("Data.String.Unsafe.charAt: Invalid index.");
    };

    public static Object $char = (java.util.function.Function<Object, Object>) (s) -> {
        String str = (String) s;
        if (str.length() == 1) return String.valueOf(str.charAt(0));
        throw new RuntimeException("Data.String.Unsafe.char: Expected string of length 1.");
    };



}
