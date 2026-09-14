#include "purescript.h"
#include <vector>
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_ArrayOpsFFICheatcode )
exports["runArrayOpsFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); std::vector<int> a; a.reserve(n);
    for (int i=0; i<n; i++) a.push_back(i);
    int sum = 0; for(int x : a) sum += x;
    dummy = sum; return sum;
};
FOREIGN_END
