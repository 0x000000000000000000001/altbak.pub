using System;

namespace Data.String.Common;

public static class FFI {
    public static object _LocaleCompare(object lt, object eq, object gt, object s1, object s2) {
        int result = string.Compare((string)s1, (string)s2, StringComparison.CurrentCulture);
        if (result < 0) return lt;
        if (result > 0) return gt;
        return eq;
    }
    public static object Replace(object arg1, object arg2, object arg3) {
        string search = (string)arg1;
        string replace = (string)arg2;
        string target = (string)arg3;
        if (search == "") return replace + target;
        int pos = target.IndexOf(search, StringComparison.Ordinal);
        if (pos < 0) return target;
        return target.Substring(0, pos) + replace + target.Substring(pos + search.Length);
    }
    public static object ReplaceAll(object arg1, object arg2, object arg3) {
        return ((string)arg3).Replace((string)arg1, (string)arg2);
    }
    public static object Split(object arg1, object arg2) {
        string sep = (string)arg1;
        string s = (string)arg2;
        if (string.IsNullOrEmpty(sep)) {
            if (string.IsNullOrEmpty(s)) return new object[0];
            var arr = new object[s.Length];
            for (int i = 0; i < s.Length; i++) {
                arr[i] = s[i].ToString();
            }
            return arr;
        } else {
            string[] parts = s.Split(new[] { sep }, StringSplitOptions.None);
            object[] res = new object[parts.Length];
            for (int i = 0; i < parts.Length; i++) {
                res[i] = parts[i];
            }
            return res;
        }
    }
    public static object ToLower(object arg1) {
        return ((string)arg1).ToLowerInvariant();
    }
    public static object ToUpper(object arg1) {
        return ((string)arg1).ToUpperInvariant();
    }
    public static object Trim(object arg1) {
        return ((string)arg1).Trim();
    }
    public static object JoinWith(object arg1, object arg2) {
        return string.Join((string)arg1, (object[])arg2);
    }
}
