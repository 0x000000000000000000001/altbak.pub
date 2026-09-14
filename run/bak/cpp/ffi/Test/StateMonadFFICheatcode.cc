#include "purescript.h"
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_StateMonadFFICheatcode )
exports["runStateMonadFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); int state = 0;
    for(int i=0; i<n; i++) state++;
    dummy = state; return state;
};
FOREIGN_END
