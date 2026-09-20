#include "purescript.h"
#include <functional>
#include <memory>

namespace {
    struct Unit {};
    template <class S, class A> struct StateResult { A val; S state; };
    template <class S, class A> using StateFn = std::function<StateResult<S, A>(S)>;
    template <class S, class A> using State = std::shared_ptr<const StateFn<S, A>>;

    template <class S, class A, class B>
    State<S, B> bindState(State<S, A> action, std::function<State<S, B>(A)> next) {
        return std::make_shared<StateFn<S, B>>([action, next](S state) {
            const auto result = (*action)(state);
            return (*next(result.val))(result.state);
        });
    }

    template <class S, class A>
    State<S, A> pureState(A value) {
        return std::make_shared<StateFn<S, A>>([value](S state) {
            return StateResult<S, A>{value, state};
        });
    }

    template <class S>
    State<S, S> get() {
        return std::make_shared<StateFn<S, S>>([](S state) {
            return StateResult<S, S>{state, state};
        });
    }

    template <class S>
    State<S, Unit> put(S nextState) {
        return std::make_shared<StateFn<S, Unit>>([nextState](S) {
            return StateResult<S, Unit>{Unit{}, nextState};
        });
    }

    template <class S>
    State<S, Unit> modify(std::function<S(S)> f) {
        return bindState<S, S, Unit>(get<S>(), [f](S state) { return put<S>(f(state)); });
    }

    State<int, Unit> chainModifications(int depth) {
        if (depth == 0) return pureState<int, Unit>(Unit{});
        return bindState<int, Unit, Unit>(modify<int>([](int x) { return x + 1; }),
            [depth](Unit) { return chainModifications(depth - 1); });
    }
}

FOREIGN_BEGIN(Test_StateMonadFFI)
exports["runStateMonadFFI"] = [](const boxed& in) -> boxed {
    // This FFI wrapper passes the chain depth (60), not the repetition count.
    const int depth = unbox<int>(in);
    int result = 0;
    for (int repetition = 0; repetition < 20; ++repetition) {
        result += (*chainModifications(depth))(0).state;
    }
    return result;
};
FOREIGN_END
