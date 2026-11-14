#include "purescript.h"

namespace {
    int deepTailRec(int n, int acc) {
        if (n == 0) return acc;
        return deepTailRec(n - 1, acc + n % 3);
    }
}

FOREIGN_BEGIN(Test_TCOFFI)
exports["runTCOFFI"] = [](const boxed& in) -> boxed {
    const int result = deepTailRec(unbox<int>(in), 0);
    return result;
};
FOREIGN_END
