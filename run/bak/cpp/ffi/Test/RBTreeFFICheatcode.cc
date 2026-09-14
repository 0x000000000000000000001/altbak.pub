#include "purescript.h"
#include <set>
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_RBTreeFFICheatcode )
exports["runRBTreeFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); std::set<int> s;
    for(int i=n; i>0; i--) s.insert(i);
    dummy = s.size(); return (int)s.size(); 
};
FOREIGN_END
