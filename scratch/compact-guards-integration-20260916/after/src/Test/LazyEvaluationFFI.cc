#include "purescript.h"
#include <functional>
#include <memory>

namespace {
    using LazyFn = std::function<int()>;
    using Lazy = std::shared_ptr<const LazyFn>;

    Lazy defer(std::function<int()> thunk) {
        return std::make_shared<LazyFn>(std::move(thunk));
    }

    int force(const Lazy& value) {
        return (*value)();
    }

    Lazy buildThunks(int depth, Lazy acc) {
        while (depth > 0) {
            const Lazy previous = acc;
            acc = defer([previous]() { return force(previous) + 1; });
            --depth;
        }
        return acc;
    }
}

FOREIGN_BEGIN(Test_LazyEvaluationFFI)
exports["runLazyEvaluationFFI"] = [](const boxed& in) -> boxed {
    const int depth = unbox<int>(in);
    int result = 0;
    for (int repetition = 0; repetition < 1000; ++repetition) {
        result += force(buildThunks(depth, defer([]() { return 0; })));
    }
    return result;
};
FOREIGN_END
