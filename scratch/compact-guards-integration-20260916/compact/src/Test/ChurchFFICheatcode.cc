#include "purescript.h"
FOREIGN_BEGIN( Test_ChurchFFICheatcode )
exports["runChurchFFICheatcode"] = [](const boxed& in) -> boxed {
    // c100k represents n^5 applications of the increment function.
    const int n = unbox<int>(in);
    const int squared = n * n;
    const int applications = squared * squared * n;
    int value = 0;
    for (int i = 0; i < applications; ++i) ++value;
    return value;
};
FOREIGN_END
