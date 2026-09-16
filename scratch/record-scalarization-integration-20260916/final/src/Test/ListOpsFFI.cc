#include "purescript.h"
#include <functional>
#include <memory>

namespace {
    struct Node;
    using List = std::shared_ptr<const Node>;
    struct Node { int head; List tail; };

    List cons(int head, List tail) {
        return std::make_shared<Node>(Node{head, tail});
    }

    List range(int start, int end) {
        List result;
        for (int value = end; value >= start; --value) {
            result = cons(value, result);
        }
        return result;
    }

    List filterEvens(List values) {
        List evens;
        while (values) {
            if (values->head % 2 == 0) evens = cons(values->head, evens);
            values = values->tail;
        }
        return evens;
    }

    int foldl(const std::function<int(int, int)>& combine, int acc, List values) {
        while (values) {
            acc = combine(acc, values->head);
            values = values->tail;
        }
        return acc;
    }
}

FOREIGN_BEGIN(Test_ListOpsFFI)
exports["runListOpsFFI"] = [](const boxed& in) -> boxed {
    const auto values = range(1, unbox<int>(in));
    const auto evens = filterEvens(values);
    const int result = foldl([](int sum, int value) { return sum + value; }, 0, evens);
    return result;
};
FOREIGN_END
