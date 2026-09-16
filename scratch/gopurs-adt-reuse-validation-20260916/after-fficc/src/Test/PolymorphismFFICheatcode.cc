#include "purescript.h"
FOREIGN_BEGIN( Test_PolymorphismFFICheatcode )
exports["runPolymorphismFFICheatcode"] = [](const boxed& in) -> boxed {
    const int n = unbox<int>(in);
    int sum = 0;
    // Monoidish Int: mempty_ is 1 and mappend_ is integer addition.
    for (int i = 0; i < n; ++i) sum += 1;
    return sum;
};
FOREIGN_END
