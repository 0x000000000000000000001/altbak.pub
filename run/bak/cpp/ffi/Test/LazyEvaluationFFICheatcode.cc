#include "purescript.h"
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_LazyEvaluationFFICheatcode )
exports["runLazyEvaluationFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); int acc = 0;
    for(int i=0; i<n; i++) acc += 1;
    dummy = acc; return acc;
};
FOREIGN_END
