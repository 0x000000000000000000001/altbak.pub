#include "purescript.h"
#include <vector>
FOREIGN_BEGIN( Test_PrimesFFICheatcode )
exports["runPrimesFFICheatcode"] = [](const boxed& in) -> boxed {
    const int n = unbox<int>(in);
    if (n < 2) return 0;
    std::vector<bool> is_prime(n + 1, true);
    int sum = 0;
    for (int p = 2; p <= n; p++) {
        if (is_prime[p]) {
            sum += p;
            for (std::size_t multiple = std::size_t(p) * p;
                 multiple < is_prime.size(); multiple += p) {
                is_prime[multiple] = false;
            }
        }
    }
    return sum;
};
FOREIGN_END
