using System;

namespace Effect.Exception;

public static class FFI {
    public static System.Exception Error(string msg) => new System.Exception(msg);
    public static string Message(System.Exception e) => e.Message;
    public static string Name(System.Exception e) => e.GetType().Name;
    public static string ShowErrorImpl(System.Exception e) => e.ToString();
    public static object StackImpl(Func<string, object> just, object nothing, System.Exception e) => e.StackTrace != null ? just(e.StackTrace) : nothing;
    public static Func<object> ThrowException(System.Exception e) => () => throw e;
    public static Func<object> CatchException(Func<System.Exception, Func<object>> c, Func<object> t) => () => {
        try {
            return t();
        } catch (System.Exception e) {
            return c(e)();
        }
    };
    public static System.Exception ErrorWithCause(string msg, System.Exception cause) => new System.Exception(msg, cause);
    public static System.Exception ErrorWithName(string name, string msg) => new System.Exception(msg);
}
