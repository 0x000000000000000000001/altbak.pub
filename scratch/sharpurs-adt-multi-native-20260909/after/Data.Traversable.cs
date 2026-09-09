using System;

namespace Data.Traversable;

public static class FFI {
    public static object TraverseArrayImpl(
        Func<object, Func<object, object>> apply,
        Func<object, Func<object, object>> mapFn,
        Func<object, object> pure,
        Func<object, object> f,
        object[] arrayVal) 
    {
        Func<object, object> array1 = a => new object[] { a };
        Func<object, Func<object, object>> array2 = a => b => new object[] { a, b };
        Func<object, Func<object, Func<object, object>>> array3 = a => b => c => new object[] { a, b, c };
        
        Func<object, Func<object, object>> concat2 = xsVal => ysVal => {
            var xs = (object[])xsVal;
            var ys = (object[])ysVal;
            var res = new object[xs.Length + ys.Length];
            System.Array.Copy(xs, 0, res, 0, xs.Length);
            System.Array.Copy(ys, 0, res, xs.Length, ys.Length);
            return (object)res;
        };

        Func<long, long, object> goFn = null;
        goFn = (bot, top) => {
            long diff = top - bot;
            if (diff == 0) return pure(System.Array.Empty<object>());
            if (diff == 1) return mapFn(array1)(f(arrayVal[bot]));
            if (diff == 2) return apply(mapFn(array2)(f(arrayVal[bot])))(f(arrayVal[bot + 1]));
            if (diff == 3) return apply(apply(mapFn(array3)(f(arrayVal[bot])))(f(arrayVal[bot + 1])))(f(arrayVal[bot + 2]));
            
            long pivot = bot + (diff / 4) * 2;
            return apply(mapFn(concat2)(goFn(bot, pivot)))(goFn(pivot, top));
        };
        
        return goFn(0, arrayVal.Length);
    }
}
