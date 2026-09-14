#include "purescript.h"
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_TCOFFICheatcode )
exports["runTCOFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); int acc = 0; 
    while (n > 0) { acc += n; n--; } 
    dummy = acc; return acc; 
};
FOREIGN_END
