using System;

namespace Effect;

public static class FFI {
    public static Func<object> PureE(object a) => () => a;
    
    public static Func<object> BindE(Func<object> a, Func<object, Func<object>> f) => () => {
        var resA = a();
        return f(resA)();
    };
    
    public static Func<object> UntilE(Func<bool> f) => () => {
        while (!f()) {}
        return null;
    };
    
    public static Func<object> WhileE(Func<bool> f, Func<object> a) => () => {
        while (f()) {
            a();
        }
        return null;
    };
    
    public static Func<object> ForE(long lo, long hi, Func<long, Func<object>> f) => () => {
        for (long i = lo; i < hi; i++) {
            f(i)();
        }
        return null;
    };
    
    public static Func<object> ForeachE(object[] as_, Func<object, Func<object>> f) => () => {
        foreach (var v in as_) {
            f(v)();
        }
        return null;
    };
}
