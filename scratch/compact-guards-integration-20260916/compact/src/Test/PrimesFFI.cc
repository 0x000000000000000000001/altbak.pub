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

    List reverse(List values) {
        List result;
        while (values) {
            result = cons(values->head, result);
            values = values->tail;
        }
        return result;
    }

    List filter(const std::function<bool(int)>& predicate, List values) {
        List accepted;
        while (values) {
            if (predicate(values->head)) accepted = cons(values->head, accepted);
            values = values->tail;
        }
        return reverse(accepted);
    }

    List sieve(List values) {
        if (!values) return List{};
        const int prime = values->head;
        return cons(prime, sieve(filter(
            [prime](int value) { return value % prime != 0; }, values->tail)));
    }

    int sumList(List values) {
        int sum = 0;
        while (values) {
            sum += values->head;
            values = values->tail;
        }
        return sum;
    }
}

FOREIGN_BEGIN(Test_PrimesFFI)
exports["runPrimesFFI"] = [](const boxed& in) -> boxed {
    const int result = sumList(sieve(range(2, unbox<int>(in))));
    return result;
};
FOREIGN_END
