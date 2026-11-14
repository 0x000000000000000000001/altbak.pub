#include "purescript.h"
FOREIGN_BEGIN( Test_StateMonadFFICheatcode )
exports["runStateMonadFFICheatcode"] = [](const boxed& in) -> boxed {
    // This FFI takes the chain depth; every run starts from state zero.
    const int depth = unbox<int>(in);
    int total = 0;
    for (int repetition = 0; repetition < 20; ++repetition) {
        int state = 0;
        for (int step = 0; step < depth; ++step) ++state;
        total += state;
    }
    return total;
};
FOREIGN_END
