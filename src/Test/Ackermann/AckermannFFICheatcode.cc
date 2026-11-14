#include "purescript.h"
namespace {
int ack(int m, int n) {
    if (m == 0) return n + 1;
    if (n == 0) return ack(m - 1, 1);
    return ack(m - 1, ack(m, n - 1));
}
}
FOREIGN_BEGIN( Test_AckermannFFICheatcode )
exports["runAckermannFFICheatcode"] = [](const boxed& in) -> boxed {
    return ack(unbox<int>(in), 4);
};
FOREIGN_END
