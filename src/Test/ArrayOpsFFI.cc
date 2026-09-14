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

    std::vector<int> filterEvens(const std::vector<int>& values) {
        std::vector<int> evens;
        for (int value : values) {
            if (value % 2 == 0) evens.push_back(value);
        }
        return evens;
    }
}

FOREIGN_BEGIN(Test_ArrayOpsFFI)
exports["runArrayOpsFFI"] = [](const boxed& in) -> boxed {
    const auto values = range(1, unbox<int>(in));
    const auto evens = filterEvens(values);
    int sum = 0;
    for (int value : evens) sum += value;
    return sum;
};
FOREIGN_END
