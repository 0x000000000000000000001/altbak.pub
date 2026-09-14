#include "purescript.h"
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_RowToListFFICheatcode )
exports["runRowToListFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); dummy = n; return n;
};
FOREIGN_END
