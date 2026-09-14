#include "purescript.h"
#include <functional>
namespace {
    volatile int dummy;
    struct Show { std::function<int(int)> show; };
    auto showInt = Show{[](int x) { return x; }};
    int print(Show dict, int x) { return dict.show(x); }
}
FOREIGN_BEGIN( Test_PolymorphismFFI )
exports["runPolymorphismFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); int sum = 0;
    for(int i=0; i<n; i++) sum += print(showInt, i);
    dummy = sum; return sum;
};
FOREIGN_END
