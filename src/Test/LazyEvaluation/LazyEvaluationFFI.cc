#include "purescript.h"
#include <functional>
#include <memory>

namespace {
    template <class A> using LazyFn = std::function<A()>;
    template <class A> using Lazy = std::shared_ptr<const LazyFn<A>>;

    template <class A>
    Lazy<A> defer(std::function<A()> thunk) {
        return std::make_shared<LazyFn<A>>(std::move(thunk));
    }

    template <class A>
    A force(const Lazy<A>& value) {
        return (*value)();
    }

    Lazy<int> buildThunks(int depth, Lazy<int> acc) {
        while (depth > 0) {
            const Lazy<int> previous = acc;
            acc = defer<int>([previous]() { return force<int>(previous) + 1; });
            --depth;
        }
        return acc;
    }
}

FOREIGN_BEGIN(Test_LazyEvaluationFFI)
exports["runLazyEvaluationFFI"] = [](const boxed& in) -> boxed {
    const int repetitions = unbox<int>(in);
    int result = 0;
    for (int repetition = 0; repetition < repetitions; ++repetition) {
        result += force<int>(buildThunks(1000, defer<int>([]() { return 0; })));
    }
    return result;
};
FOREIGN_END
