#include "purescript.h"
namespace { volatile int dummy; auto tco(int n, int acc) -> int { if (n == 0) return acc; return tco(n - 1, acc + n); } }
FOREIGN_BEGIN( Test_TCOFFI )
exports["runTCOFFI"] = [](const boxed& in) -> boxed { int n = unbox<int>(in); int r = tco(n, 0); dummy = r; return r; };
FOREIGN_END
