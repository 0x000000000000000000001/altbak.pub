#include "purescript.h"
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_ChurchFFICheatcode )
exports["runChurchFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in);
    long long r = 1;
    for (int i = 0; i < 5; i++) r *= n;
    dummy = r; return (int)r; 
};
FOREIGN_END
