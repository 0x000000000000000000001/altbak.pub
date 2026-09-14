#include "purescript.h"
#include <vector>
namespace { volatile int dummy; }
FOREIGN_BEGIN( Test_PrimesFFICheatcode )
exports["runPrimesFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in);
    std::vector<bool> is_prime(n + 1, true);
    int count = 0;
    for (int p = 2; p <= n; p++) {
        if (is_prime[p]) {
            count++;
            for (int i = p * 2; i <= n; i += p) is_prime[i] = false;
        }
    }
    dummy = count; return count;
};
FOREIGN_END
