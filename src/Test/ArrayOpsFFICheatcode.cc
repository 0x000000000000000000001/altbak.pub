#include "purescript.h"
#include <algorithm>
#include <numeric>
#include <vector>
FOREIGN_BEGIN( Test_ArrayOpsFFICheatcode )
exports["runArrayOpsFFICheatcode"] = [](const boxed& in) -> boxed {
    const int n = unbox<int>(in);
    std::vector<int> values;
    if (n > 0) values.reserve(n);
    for (int i = 1; i <= n; ++i) values.push_back(i);
    values.erase(std::remove_if(values.begin(), values.end(),
        [](int value) { return value % 2 != 0; }), values.end());
    return std::accumulate(values.begin(), values.end(), 0);
};
FOREIGN_END
