using System;

namespace Data.Foldable;

public static class FFI {
    public static object FoldrArray(Func<object, Func<object, object>> f, object init, object[] xs) {
        var acc = init;
        for (long i = xs.Length - 1; i >= 0; i--) {
            acc = f(xs[i])(acc);
        }
        return acc;
    }

    public static object FoldlArray(Func<object, Func<object, object>> f, object init, object[] xs) {
        var acc = init;
        for (long i = 0; i < xs.Length; i++) {
            acc = f(acc)(xs[i]);
        }
        return acc;
    }
}
