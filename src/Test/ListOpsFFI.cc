#include "purescript.h"
#include <functional>
#include <memory>

namespace {
    template <class A> struct Node;
    template <class A> using List = std::shared_ptr<const Node<A>>;
    template <class A> struct Node { A head; List<A> tail; };

    template <class A>
    List<A> cons(A head, List<A> tail) {
        return std::make_shared<Node<A>>(Node<A>{head, tail});
    }

    List<int> range(int start, int end) {
        List<int> result;
        for (int value = end; value >= start; --value) {
            result = cons(value, result);
        }
        return result;
    }

    List<int> filterEvens(List<int> values) {
        List<int> evens;
        while (values) {
            if (values->head % 2 == 0) evens = cons(values->head, evens);
            values = values->tail;
        }
        return evens;
    }

    template <class A, class B>
    B foldl(const std::function<B(B, A)>& combine, B acc, List<A> values) {
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
    const int result = foldl<int, int>([](int sum, int value) { return sum + value; }, 0, evens);
    return result;
};
FOREIGN_END
