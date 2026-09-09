using System;

namespace Data.FunctorWithIndex;

public static class FFI {
    public static object[] MapWithIndexArray(Func<long, Func<object, object>> f, object[] xs) {
        var result = new object[xs.Length];
        for (long i = 0; i < xs.Length; i++) {
            result[i] = f(i)(xs[i]);
        }
        return result;
    }
}
