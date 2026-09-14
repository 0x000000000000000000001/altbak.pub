#include "purescript.h"
#include <functional>
namespace {
    volatile int dummy;
    using ChurchFn = std::function<int(int)>;
    using Church = std::function<ChurchFn(ChurchFn)>;
    Church fromInt(int n) {
        if (n == 0) return [](ChurchFn) { return [](int x) { return x; }; };
        Church prev = fromInt(n - 1);
        return [prev](ChurchFn f) { return [prev, f](int x) { return f(prev(f)(x)); }; };
    }
    Church multiply(Church m, Church n) { return [m, n](ChurchFn f) { return [m, n, f](int x) { return m(n(f))(x); }; }; }
    Church square(int n) { return multiply(fromInt(n), fromInt(n)); }
}
FOREIGN_BEGIN( Test_ChurchFFI )
exports["runChurchFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in);
    Church p4 = multiply(square(n), square(n));
    Church p5 = multiply(p4, fromInt(n));
    int r = p5([](int x) { return x + 1; })(0);
    dummy = r; return r; 
};
FOREIGN_END
