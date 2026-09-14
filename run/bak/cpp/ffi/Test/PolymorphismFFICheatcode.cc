#include "purescript.h"
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_PolymorphismFFICheatcode )
exports["runPolymorphismFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); long long sum = 0;
    for(int i=0; i<n; i++) sum += i;
    dummy = sum; return (int)sum;
};
FOREIGN_END
