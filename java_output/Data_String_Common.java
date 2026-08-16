public class Data_String_Common {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    // FFI provided by ../javapurs/javapurs-strings/src/Data/String/Common.java
    public static Object _localeCompare = (java.util.function.Function<Object, Object>) (lt) -> (java.util.function.Function<Object, Object>) (eq) -> (java.util.function.Function<Object, Object>) (gt) -> (java.util.function.Function<Object, Object>) (s1) -> (java.util.function.Function<Object, Object>) (s2) -> {
        int result = ((String) s1).compareTo((String) s2);
        return result < 0 ? lt : result > 0 ? gt : eq;
    };

    public static Object replace = (java.util.function.Function<Object, Object>) (s1) -> (java.util.function.Function<Object, Object>) (s2) -> (java.util.function.Function<Object, Object>) (s3) -> {
        String target = (String) s1;
        String replacement = (String) s2;
        String str = (String) s3;
        int index = str.indexOf(target);
        if (index < 0) return str;
        return str.substring(0, index) + replacement + str.substring(index + target.length());
    };

    public static Object replaceAll = (java.util.function.Function<Object, Object>) (s1) -> (java.util.function.Function<Object, Object>) (s2) -> (java.util.function.Function<Object, Object>) (s3) -> {
        return ((String) s3).replace((String) s1, (String) s2);
    };

    public static Object split = (java.util.function.Function<Object, Object>) (sep) -> (java.util.function.Function<Object, Object>) (s) -> {
        String separator = (String) sep;
        String str = (String) s;
        if (str.isEmpty() && separator.isEmpty()) return new Object[0];
        if (separator.isEmpty()) {
            Object[] arr = new Object[str.length()];
            for (int i = 0; i < str.length(); i++) {
                arr[i] = String.valueOf(str.charAt(i));
            }
            return arr;
        }
        java.util.List<Object> list = new java.util.ArrayList<>();
        int start = 0;
        int index = str.indexOf(separator);
        while (index >= 0) {
            list.add(str.substring(start, index));
            start = index + separator.length();
            index = str.indexOf(separator, start);
        }
        list.add(str.substring(start));
        return list.toArray(new Object[0]);
    };

    public static Object toLower = (java.util.function.Function<Object, Object>) (s) -> ((String) s).toLowerCase();
    public static Object toUpper = (java.util.function.Function<Object, Object>) (s) -> ((String) s).toUpperCase();
    public static Object trim = (java.util.function.Function<Object, Object>) (s) -> ((String) s).trim();
    public static Object joinWith = (java.util.function.Function<Object, Object>) (s) -> (java.util.function.Function<Object, Object>) (xs) -> {
        String separator = (String) s;
        Object[] arr = (Object[]) xs;
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < arr.length; i++) {
            sb.append((String) arr[i]);
            if (i < arr.length - 1) sb.append(separator);
        }
        return sb.toString();
    };


public static final Object $null = (java.util.function.Function<Object, Object>) (s_0) -> java.util.Objects.equals(s_0, "");
public static final Object localeCompare = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_Common._localeCompare)).apply(new Data_Ordering.LT()))).apply(new Data_Ordering.EQ()))).apply(new Data_Ordering.GT());
}
