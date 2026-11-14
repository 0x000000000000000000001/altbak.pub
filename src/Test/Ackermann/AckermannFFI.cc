#include "purescript.h"

namespace {
    int ackermann(int m, int n) {
        if (m == 0) return n + 1;
        if (n == 0) return ackermann(m - 1, 1);
        return ackermann(m - 1, ackermann(m, n - 1));
    }
}

FOREIGN_BEGIN(Test_AckermannFFI)
exports["runAckermannFFI"] = [](const boxed& in) -> boxed {
    const int result = ackermann(unbox<int>(in), 4);
    return result;
};
FOREIGN_END
