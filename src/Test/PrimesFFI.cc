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

    template <class A>
    List<A> reverse(List<A> values) {
        List<A> result;
        while (values) {
            result = cons(values->head, result);
            values = values->tail;
        }
        return result;
    }

    template <class A>
    List<A> filter(const std::function<bool(A)>& predicate, List<A> values) {
        List<A> accepted;
        while (values) {
            if (predicate(values->head)) accepted = cons(values->head, accepted);
            values = values->tail;
        }
        return reverse(accepted);
    }

    List<int> sieve(List<int> values) {
        if (!values) return List<int>{};
        const int prime = values->head;
        return cons(prime, sieve(filter<int>(
            [prime](int value) { return value % prime != 0; }, values->tail)));
    }

    int sumList(List<int> values) {
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
