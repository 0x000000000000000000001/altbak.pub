#include "purescript.h"
#include <vector>
FOREIGN_BEGIN( Test_ListOpsFFICheatcode )
exports["runListOpsFFICheatcode"] = [](const boxed& in) -> boxed {
    const int n = unbox<int>(in);
    std::vector<int> values;
    if (n > 0) values.reserve(n);
    for (int i = 1; i <= n; ++i) values.push_back(i);
    // filterEvens prepends retained elements; sum that descending sequence.
    int sum = 0;
    for (auto it = values.rbegin(); it != values.rend(); ++it) {
        if (*it % 2 == 0) sum += *it;
    }
    return sum;
};
FOREIGN_END
