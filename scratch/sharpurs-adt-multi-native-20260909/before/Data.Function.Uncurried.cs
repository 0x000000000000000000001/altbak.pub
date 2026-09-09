using System;

namespace Data.Function.Uncurried;

public static class FFI {
    private static object Identity(object fn) => fn;

    public static object RunFn2(object fn) => Identity(fn);
    public static object RunFn3(object fn) => Identity(fn);
    public static object RunFn4(object fn) => Identity(fn);
    public static object RunFn5(object fn) => Identity(fn);
    public static object RunFn6(object fn) => Identity(fn);
    public static object RunFn7(object fn) => Identity(fn);
    public static object RunFn8(object fn) => Identity(fn);
    public static object RunFn9(object fn) => Identity(fn);
    public static object RunFn10(object fn) => Identity(fn);

    public static object MkFn2(object fn) => Identity(fn);
    public static object MkFn3(object fn) => Identity(fn);
    public static object MkFn4(object fn) => Identity(fn);
    public static object MkFn5(object fn) => Identity(fn);
    public static object MkFn6(object fn) => Identity(fn);
    public static object MkFn7(object fn) => Identity(fn);
    public static object MkFn8(object fn) => Identity(fn);
    public static object MkFn9(object fn) => Identity(fn);
    public static object MkFn10(object fn) => Identity(fn);

    public static object MkFn0(object fn) => Identity(fn);
    public static object RunFn0(Func<object, object> f) => f(null);
}
