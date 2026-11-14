#include "purescript.h"
#include <vector>

namespace {
    std::vector<int> range(int start, int end) {
        std::vector<int> values;
        const int step = start <= end ? 1 : -1;
        for (int value = start;; value += step) {
            values.push_back(value);
            if (value == end) return values;
        }
    }

    template <class A, class Predicate>
    std::vector<A> filter(Predicate predicate, const std::vector<A>& values) {
        std::vector<A> accepted;
        for (const A& value : values) {
            if (predicate(value)) accepted.push_back(value);
        }
        return accepted;
    }

    template <class A, class B, class Combine>
    B foldl(Combine combine, B acc, const std::vector<A>& values) {
        for (const A& value : values) acc = combine(acc, value);
        return acc;
    }
}

FOREIGN_BEGIN(Test_ArrayOpsFFI)
exports["runArrayOpsFFI"] = [](const boxed& in) -> boxed {
    const auto values = range(1, unbox<int>(in));
    const auto evens = filter([](int value) { return value % 2 == 0; }, values);
    return foldl([](int sum, int value) { return sum + value; }, 0, evens);
};
FOREIGN_END
