#include "purescript.h"
#include <vector>
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_ArrayOpsFFI )
exports["runArrayOpsFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); std::vector<int> a;
    for (int i=0; i<n; i++) { std::vector<int> n_a = a; n_a.push_back(i); a = n_a; }
    int sum = 0; for(int x : a) sum += x;
    dummy = sum; return sum;
};
FOREIGN_END
