#include "purescript.h"
#include <functional>
namespace {
    volatile int dummy;
    using State = std::function<std::pair<int, int>(int)>;
    State stateBind(State m, std::function<State(int)> f) {
        return [m, f](int s) { auto p = m(s); return f(p.second)(p.first); };
    }
    State pure(int v) { return [v](int s) { return std::make_pair(s, v); }; }
    State get() { return [](int s) { return std::make_pair(s, s); }; }
    State put(int s) { return [s](int) { return std::make_pair(s, 0); }; }
}
FOREIGN_BEGIN( Test_StateMonadFFI )
exports["runStateMonadFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in);
    State m = pure(0);
    for(int i=0; i<n; i++) m = stateBind(m, [](int) { return stateBind(get(), [](int s) { return put(s + 1); }); });
    int r = m(0).first; dummy = r; return r;
};
FOREIGN_END
