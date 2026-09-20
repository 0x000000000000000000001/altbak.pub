#include "purescript.h"
#include <functional>
#include <memory>

namespace {
    template <class A> using Endo = std::shared_ptr<const std::function<A(A)>>;
    template <class A> using ChurchFn = std::function<Endo<A>(Endo<A>)>;
    template <class A> using Church = std::shared_ptr<const ChurchFn<A>>;

    template <class A>
    Church<A> zero() {
        return std::make_shared<ChurchFn<A>>([](Endo<A>) {
            return std::make_shared<std::function<A(A)>>([](A x) { return x; });
        });
    }

    template <class A>
    Church<A> successor(Church<A> previous) {
        return std::make_shared<ChurchFn<A>>([previous](Endo<A> f) {
            return std::make_shared<std::function<A(A)>>([previous, f](A x) {
                return (*f)((*(*previous)(f))(x));
            });
        });
    }

    Church<int> fromInt(int n) {
        return n == 0 ? zero<int>() : successor<int>(fromInt(n - 1));
    }

    template <class A>
    Church<A> multiply(Church<A> m, Church<A> n) {
        return std::make_shared<ChurchFn<A>>([m, n](Endo<A> f) {
            return std::make_shared<std::function<A(A)>>([m, n, f](A x) {
                return (*(*m)((*n)(f)))(x);
            });
        });
    }

    Church<int> square(int n) {
        return multiply<int>(fromInt(n), fromInt(n));
    }

}

FOREIGN_BEGIN(Test_ChurchFFI)
exports["runChurchFFI"] = [](const boxed& in) -> boxed {
    const int n = unbox<int>(in);
    const Church<int> c10k = multiply<int>(square(n), square(n));
    const Church<int> c100k = multiply<int>(c10k, fromInt(n));
    const Endo<int> increment = std::make_shared<std::function<int(int)>>(
        [](int x) { return x + 1; });
    const int result = (*(*c100k)(increment))(0);
    return result;
};
FOREIGN_END
