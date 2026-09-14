#include "purescript.h"
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_FibFFICheatcode )
exports["runFibFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); 
    int a = 0, b = 1;
    for (int i = 0; i < n; i++) { int temp = a; a = b; b = temp + b; }
    dummy = a; return a; 
};
FOREIGN_END
