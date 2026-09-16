#include "purescript.h"
#include <functional>
#include <memory>

namespace {
    struct StateResult { int val, state; };
    using StateFn = std::function<StateResult(int)>;
    using State = std::shared_ptr<const StateFn>;

    State bindState(State action, std::function<State(int)> next) {
        return std::make_shared<StateFn>([action, next](int state) {
            const auto result = (*action)(state);
            return (*next(result.val))(result.state);
        });
    }

    State pureState(int value) {
        return std::make_shared<StateFn>([value](int state) {
            return StateResult{value, state};
        });
    }

    State get() {
        return std::make_shared<StateFn>([](int state) {
            return StateResult{state, state};
        });
    }

    State put(int nextState) {
        return std::make_shared<StateFn>([nextState](int) {
            return StateResult{0, nextState};
        });
    }

    State modify(std::function<int(int)> f) {
        return bindState(get(), [f](int state) { return put(f(state)); });
    }

    State chainModifications(int depth) {
        if (depth == 0) return pureState(0);
        return bindState(modify([](int x) { return x + 1; }),
            [depth](int) { return chainModifications(depth - 1); });
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
