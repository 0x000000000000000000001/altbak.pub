#include "purescript.h"
namespace { volatile int dummy; auto ack(int m, int n) -> int { if (m == 0) return n + 1; if (n == 0) return ack(m - 1, 1); return ack(m - 1, ack(m, n - 1)); } }
FOREIGN_BEGIN( Test_AckermannFFICheatcode )
exports["runAckermannFFICheatcode"] = [](const boxed& in) -> boxed { int r = ack(3, 4); dummy = r; return r; };
FOREIGN_END
