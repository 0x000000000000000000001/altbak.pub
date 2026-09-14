#include "purescript.h"
#include <vector>
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_ListOpsFFICheatcode )
exports["runListOpsFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in);
    std::vector<int> v; v.reserve(n);
    for (int i = 0; i < n; i++) v.push_back(n - 1 - i);
    for (int i = 0; i < n; i++) v[i] *= 2;
    int sum = 0; for(int i = n - 1; i >= 0; i--) sum += v[i];
    dummy = sum; return sum;
};
FOREIGN_END
