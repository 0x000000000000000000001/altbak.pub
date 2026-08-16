public class Data_String_Regex {
    public static final Object FFI_STUB = new java.util.function.Function<Object, Object>() {
        public Object apply(Object arg) { return this; }
    };
    // FFI provided by ../javapurs/javapurs-strings/src/Data/String/Regex.java
    public static Object showRegexImpl = (java.util.function.Function<Object, Object>) (r) -> {
        return r.toString();
    };

    public static Object regexImpl = (java.util.function.Function<Object, Object>) (left) -> (java.util.function.Function<Object, Object>) (right) -> (java.util.function.Function<Object, Object>) (s1) -> (java.util.function.Function<Object, Object>) (s2) -> {
        try {
            String pattern = (String) s1;
            String flags = (String) s2;
            int f = 0;
            if (flags.contains("i")) f |= java.util.regex.Pattern.CASE_INSENSITIVE;
            if (flags.contains("m")) f |= java.util.regex.Pattern.MULTILINE;
            if (flags.contains("s")) f |= java.util.regex.Pattern.DOTALL;
            java.util.regex.Pattern p = java.util.regex.Pattern.compile(pattern, f);
            java.util.LinkedHashMap<String, Object> map = new java.util.LinkedHashMap<>();
            map.put("pattern", p);
            map.put("flags", flags);
            return ((java.util.function.Function<Object, Object>) right).apply(map);
        } catch (Exception e) {
            return ((java.util.function.Function<Object, Object>) left).apply(e.getMessage());
        }
    };

    public static Object source = (java.util.function.Function<Object, Object>) (r) -> {
        java.util.regex.Pattern p = (java.util.regex.Pattern) ((java.util.LinkedHashMap<String, Object>) r).get("pattern");
        return p.pattern();
    };

    public static Object flagsImpl = (java.util.function.Function<Object, Object>) (r) -> {
        String flags = (String) ((java.util.LinkedHashMap<String, Object>) r).get("flags");
        java.util.LinkedHashMap<String, Object> map = new java.util.LinkedHashMap<>();
        map.put("multiline", flags.contains("m"));
        map.put("ignoreCase", flags.contains("i"));
        map.put("global", flags.contains("g"));
        map.put("dotAll", flags.contains("s"));
        map.put("sticky", flags.contains("y"));
        map.put("unicode", flags.contains("u"));
        return map;
    };

    public static Object test = (java.util.function.Function<Object, Object>) (r) -> (java.util.function.Function<Object, Object>) (s) -> {
        java.util.regex.Pattern p = (java.util.regex.Pattern) ((java.util.LinkedHashMap<String, Object>) r).get("pattern");
        return p.matcher((String) s).find();
    };

    public static Object _match = (java.util.function.Function<Object, Object>) (just) -> (java.util.function.Function<Object, Object>) (nothing) -> (java.util.function.Function<Object, Object>) (r) -> (java.util.function.Function<Object, Object>) (s) -> {
        java.util.regex.Pattern p = (java.util.regex.Pattern) ((java.util.LinkedHashMap<String, Object>) r).get("pattern");
        String flags = (String) ((java.util.LinkedHashMap<String, Object>) r).get("flags");
        java.util.regex.Matcher m = p.matcher((String) s);
        if (flags.contains("g")) {
            java.util.List<Object> list = new java.util.ArrayList<>();
            while (m.find()) {
                list.add(((java.util.function.Function<Object, Object>) just).apply(m.group()));
            }
            if (list.isEmpty()) return nothing;
            return ((java.util.function.Function<Object, Object>) just).apply(list.toArray(new Object[0]));
        } else {
            if (m.find()) {
                Object[] list = new Object[m.groupCount() + 1];
                for (int i = 0; i <= m.groupCount(); i++) {
                    String group = m.group(i);
                    list[i] = group == null ? nothing : ((java.util.function.Function<Object, Object>) just).apply(group);
                }
                return ((java.util.function.Function<Object, Object>) just).apply(list);
            }
            return nothing;
        }
    };

    public static Object replace = (java.util.function.Function<Object, Object>) (r) -> (java.util.function.Function<Object, Object>) (s1) -> (java.util.function.Function<Object, Object>) (s2) -> {
        java.util.regex.Pattern p = (java.util.regex.Pattern) ((java.util.LinkedHashMap<String, Object>) r).get("pattern");
        String flags = (String) ((java.util.LinkedHashMap<String, Object>) r).get("flags");
        java.util.regex.Matcher m = p.matcher((String) s2);
        String replacement = (String) s1;
        if (flags.contains("g")) {
            return m.replaceAll(replacement);
        } else {
            return m.replaceFirst(replacement);
        }
    };

    public static Object _replaceBy = (java.util.function.Function<Object, Object>) (just) -> (java.util.function.Function<Object, Object>) (nothing) -> (java.util.function.Function<Object, Object>) (r) -> (java.util.function.Function<Object, Object>) (f) -> (java.util.function.Function<Object, Object>) (s) -> {
        java.util.regex.Pattern p = (java.util.regex.Pattern) ((java.util.LinkedHashMap<String, Object>) r).get("pattern");
        String flags = (String) ((java.util.LinkedHashMap<String, Object>) r).get("flags");
        java.util.regex.Matcher m = p.matcher((String) s);
        StringBuffer sb = new StringBuffer();
        boolean global = flags.contains("g");
        while (m.find()) {
            Object[] groups = new Object[m.groupCount()];
            for (int i = 1; i <= m.groupCount(); i++) {
                String g = m.group(i);
                groups[i - 1] = g == null ? nothing : ((java.util.function.Function<Object, Object>) just).apply(g);
            }
            String replacement = (String) ((java.util.function.Function<Object, Object>) ((java.util.function.Function<Object, Object>) f).apply(m.group())).apply(groups);
            m.appendReplacement(sb, java.util.regex.Matcher.quoteReplacement(replacement));
            if (!global) break;
        }
        m.appendTail(sb);
        return sb.toString();
    };

    public static Object _search = (java.util.function.Function<Object, Object>) (just) -> (java.util.function.Function<Object, Object>) (nothing) -> (java.util.function.Function<Object, Object>) (r) -> (java.util.function.Function<Object, Object>) (s) -> {
        java.util.regex.Pattern p = (java.util.regex.Pattern) ((java.util.LinkedHashMap<String, Object>) r).get("pattern");
        java.util.regex.Matcher m = p.matcher((String) s);
        if (m.find()) {
            return ((java.util.function.Function<Object, Object>) just).apply(m.start());
        }
        return nothing;
    };

    public static Object split = (java.util.function.Function<Object, Object>) (r) -> (java.util.function.Function<Object, Object>) (s) -> {
        java.util.regex.Pattern p = (java.util.regex.Pattern) ((java.util.LinkedHashMap<String, Object>) r).get("pattern");
        return p.split((String) s, -1);
    };


public static final Object showRegex = (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("show", Data_String_Regex.showRegexImpl);  return __map; } }).get();
public static final Object search = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_Regex._search)).apply(Data_Maybe.Just))).apply(new Data_Maybe.Nothing());
public static final Object replaceprime = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_Regex._replaceBy)).apply(Data_Maybe.Just))).apply(new Data_Maybe.Nothing());
public static final Object renderFlags = (java.util.function.Function<Object, Object>) (v_0) -> (((String) ((((String) ((((String) ((((String) ((((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) v_0).get("global"))) ? "g" : ""))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) v_0).get("ignoreCase"))) ? "i" : "")))))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) v_0).get("multiline"))) ? "m" : "")))))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) v_0).get("dotAll"))) ? "s" : "")))))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) v_0).get("sticky"))) ? "y" : "")))))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) v_0).get("unicode"))) ? "u" : ""))));
public static final Object regex = (java.util.function.Function<Object, Object>) (s_0) -> (java.util.function.Function<Object, Object>) (f_1) -> ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_Regex.regexImpl)).apply(Data_Either.Left))).apply(Data_Either.Right))).apply(s_0))).apply((((String) ((((String) ((((String) ((((String) ((((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) f_1).get("global"))) ? "g" : ""))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) f_1).get("ignoreCase"))) ? "i" : "")))))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) f_1).get("multiline"))) ? "m" : "")))))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) f_1).get("dotAll"))) ? "s" : "")))))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) f_1).get("sticky"))) ? "y" : "")))))) + ((String) (( ((Boolean) (((java.util.LinkedHashMap<String, Object>) f_1).get("unicode"))) ? "u" : "")))));
public static final Object parseFlags = (java.util.function.Function<Object, Object>) (s_0) -> (new java.util.function.Supplier<Object>() { public Object get() { java.util.Map<String, Object> __map = new java.util.LinkedHashMap<>(); __map.put("global", ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_CodeUnits.contains)).apply("g"))).apply(s_0)); __map.put("ignoreCase", ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_CodeUnits.contains)).apply("i"))).apply(s_0)); __map.put("multiline", ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_CodeUnits.contains)).apply("m"))).apply(s_0)); __map.put("dotAll", ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_CodeUnits.contains)).apply("s"))).apply(s_0)); __map.put("sticky", ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_CodeUnits.contains)).apply("y"))).apply(s_0)); __map.put("unicode", ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_CodeUnits.contains)).apply("u"))).apply(s_0));  return __map; } }).get();
public static final Object match = ((java.util.function.Function<Object, Object>) (((java.util.function.Function<Object, Object>) (Data_String_Regex._match)).apply(Data_Maybe.Just))).apply(new Data_Maybe.Nothing());
public static final Object flags = (java.util.function.Function<Object, Object>) (x_0) -> ((java.util.function.Function<Object, Object>) (Data_String_Regex.flagsImpl)).apply(x_0);
}
