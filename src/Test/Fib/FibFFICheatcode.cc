#include "purescript.h"
FOREIGN_BEGIN( Test_FibFFICheatcode )
exports["runFibFFICheatcode"] = [](const boxed& in) -> boxed {
    const int n = unbox<int>(in);
    if (n == 0) return 0;
    int previous = 0;
    int current = 1;
    for (int i = 1; i < n; ++i) {
        const int next = previous + current;
        previous = current;
        current = next;
    }
    return current;
};
FOREIGN_END
