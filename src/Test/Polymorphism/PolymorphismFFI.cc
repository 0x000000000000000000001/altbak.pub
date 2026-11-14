#include "purescript.h"
#include <functional>

namespace {
    template <class A> using Endo = std::function<A(A)>;
    template <class A>
    struct Monoidish {
        A mempty;
        std::function<Endo<A>(A)> mappend;
    };

    template <class A>
    A polyLoop(const Monoidish<A>& dict, int n, A acc) {
        while (n > 0) {
            acc = dict.mappend(acc)(dict.mempty);
            --n;
        }
        return acc;
    }
}

FOREIGN_BEGIN(Test_PolymorphismFFI)
exports["runPolymorphismFFI"] = [](const boxed& in) -> boxed {
    const Monoidish<int> dict{1, [](int x) -> Endo<int> {
        return [x](int y) { return x + y; };
    }};
    const int result = polyLoop(dict, unbox<int>(in), 0);
    return result;
};
FOREIGN_END
