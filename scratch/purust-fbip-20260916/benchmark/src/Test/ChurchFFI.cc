#include "purescript.h"
#include <functional>
#include <memory>

namespace {
    using IntFn = std::shared_ptr<const std::function<int(int)>>;
    using Church = std::shared_ptr<const std::function<IntFn(IntFn)>>;

    Church fromInt(int n) {
        if (n == 0) {
            return std::make_shared<std::function<IntFn(IntFn)>>([](IntFn) {
                return std::make_shared<std::function<int(int)>>([](int x) { return x; });
            });
        }
        const Church previous = fromInt(n - 1);
        return std::make_shared<std::function<IntFn(IntFn)>>([previous](IntFn f) {
            return std::make_shared<std::function<int(int)>>([previous, f](int x) {
                return (*f)((*(*previous)(f))(x));
            });
        });
    }

    Church multiply(Church m, Church n) {
        return std::make_shared<std::function<IntFn(IntFn)>>([m, n](IntFn f) {
            return std::make_shared<std::function<int(int)>>([m, n, f](int x) {
                return (*(*m)((*n)(f)))(x);
            });
        });
    }

    Church square(int n) {
        return multiply(fromInt(n), fromInt(n));
    }

}

FOREIGN_BEGIN(Test_ChurchFFI)
exports["runChurchFFI"] = [](const boxed& in) -> boxed {
    const int n = unbox<int>(in);
    const Church c10k = multiply(square(n), square(n));
    const Church c100k = multiply(c10k, fromInt(n));
    const IntFn increment = std::make_shared<std::function<int(int)>>(
        [](int x) { return x + 1; });
    const int result = (*(*c100k)(increment))(0);
    return result;
};
FOREIGN_END
