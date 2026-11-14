#include "purescript.h"
FOREIGN_BEGIN( Test_TCOFFICheatcode )
exports["runTCOFFICheatcode"] = [](const boxed& in) -> boxed {
    int acc = 0;
    for (int n = unbox<int>(in); n > 0; --n) acc += n % 3;
    return acc;
};
FOREIGN_END
