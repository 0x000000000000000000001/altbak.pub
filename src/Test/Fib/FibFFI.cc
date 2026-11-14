#include "purescript.h"
namespace { auto fib(int n) -> int { if(n<=1) return n; return fib(n-1) + fib(n-2); } }
FOREIGN_BEGIN( Test_FibFFI )
exports["runFibFFI"] = [](const boxed& in) -> boxed { int n = unbox<int>(in); int r = fib(n); return r; };
FOREIGN_END
