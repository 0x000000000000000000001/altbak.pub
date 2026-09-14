#include "purescript.h"
#include <functional>
namespace {
    volatile int dummy;
    using Lazy = std::function<int()>;
    Lazy add(Lazy a, Lazy b) { return [a, b]() { return a() + b(); }; }
}
FOREIGN_BEGIN( Test_LazyEvaluationFFI )
exports["runLazyEvaluationFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); Lazy l = [](){ return 0; };
    for(int i=0; i<n; i++) l = add(l, [](){ return 1; });
    int r = l(); dummy = r; return r;
};
FOREIGN_END
