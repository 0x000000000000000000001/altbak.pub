#include "purescript.h"
#include <functional>

namespace {
    using IntFn = std::function<int(int)>;
    struct Monoidish {
        int mempty;
        std::function<IntFn(int)> mappend;
    };

    int polyLoop(const Monoidish& dict, int n, int acc) {
        while (n > 0) {
            acc = dict.mappend(acc)(dict.mempty);
            --n;
        }
        return acc;
    }
}

FOREIGN_BEGIN(Test_PolymorphismFFI)
exports["runPolymorphismFFI"] = [](const boxed& in) -> boxed {
    const Monoidish dict{1, [](int x) -> IntFn {
        return [x](int y) { return x + y; };
    }};
    const int result = polyLoop(dict, unbox<int>(in), 0);
    return result;
};
FOREIGN_END
