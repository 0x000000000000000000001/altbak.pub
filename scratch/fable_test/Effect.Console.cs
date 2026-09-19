using System;

namespace Effect.Console;

public static class FFI {
    public static Func<object> Log(string s) => () => {
        System.Console.WriteLine(s);
        return null;
    };

    public static Func<object> Warn(string s) => () => {
        System.Console.WriteLine("[WARN] " + s);
        return null;
    };

    public static Func<object> Error(string s) => () => {
        System.Console.WriteLine("[ERROR] " + s);
        return null;
    };

    public static Func<object> Info(string s) => () => {
        System.Console.WriteLine("[INFO] " + s);
        return null;
    };

    public static Func<object> Debug(string s) => Log(s);
    public static Func<object> Time(string s) => Log(s);
    public static Func<object> TimeLog(string s) => Log(s);
    public static Func<object> TimeEnd(string s) => Log(s);

    public static void Clear() {}

    public static Func<object> Group(string s) => Log(s);
    public static Func<object> GroupCollapsed(string s) => Log(s);
    public static void GroupEnd() => Clear();
}
