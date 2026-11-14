#include "purescript.h"
FOREIGN_BEGIN( Test_LazyEvaluationFFICheatcode )
exports["runLazyEvaluationFFICheatcode"] = [](const boxed& in) -> boxed {
    // This FFI takes the chain depth; each of the 1000 chains starts at zero.
    const int depth = unbox<int>(in);
    int total = 0;
    for (int repetition = 0; repetition < 1000; ++repetition) {
        int value = 0;
        for (int step = 0; step < depth; ++step) ++value;
        total += value;
    }
    return total;
};
FOREIGN_END
