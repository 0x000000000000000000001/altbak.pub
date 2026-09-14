// Independent result oracles for the two native C++ benchmark columns.
// Link against the real pscpp runtime and all 28 FFI translation units.
#include "purescript.h"
#include <chrono>
#include <iostream>
#include <stdexcept>
#include <string>
#include <thread>

namespace Bench {
    using namespace purescript;
    DEFINE_FOREIGN_DICTIONARY_AND_ACCESSOR()
}

#define DECLARE_FOREIGN(NAME) \
    namespace Test_##NAME { \
        using namespace purescript; \
        DEFINE_FOREIGN_DICTIONARY_AND_ACCESSOR() \
    }
#define DECLARE_PAIR(NAME) DECLARE_FOREIGN(NAME##FFI) DECLARE_FOREIGN(NAME##FFICheatcode)
DECLARE_PAIR(AstTree)
DECLARE_PAIR(Fib)
DECLARE_PAIR(ListOps)
DECLARE_PAIR(TCO)
DECLARE_PAIR(Records)
DECLARE_PAIR(Ackermann)
DECLARE_PAIR(Church)
DECLARE_PAIR(Primes)
DECLARE_PAIR(RBTree)
DECLARE_PAIR(Polymorphism)
DECLARE_PAIR(StateMonad)
DECLARE_PAIR(LazyEvaluation)
DECLARE_PAIR(ArrayOps)
DECLARE_PAIR(RowToList)

namespace {
int checks = 0;
void check(const char* name, purescript::dict_t& foreign, int input, int expected) {
    const int actual = purescript::unbox<int>(foreign[name](purescript::boxed(input)));
    if (actual != expected) {
        throw std::runtime_error(std::string(name) + "(" + std::to_string(input) +
            "): expected " + std::to_string(expected) + ", got " + std::to_string(actual));
    }
    ++checks;
}
#define CHECK_PAIR(NAME, INPUT, EXPECTED) do { \
    check("run" #NAME "FFI", Test_##NAME##FFI::foreign(), INPUT, EXPECTED); \
    check("run" #NAME "FFICheatcode", Test_##NAME##FFICheatcode::foreign(), INPUT, EXPECTED); \
} while (false)

int astValue(int n) {
    int result = 1;
    for (int i = 1; i <= n; ++i) result = (i + 1) * result - 1;
    return result;
}
int fibonacci(int n) {
    int a = 0, b = 1;
    for (int i = 0; i < n; ++i) { const int next = a + b; a = b; b = next; }
    return a;
}
int sumPrimes(int n) {
    int total = 0;
    for (int candidate = 2; candidate <= n; ++candidate) {
        bool prime = true;
        for (int divisor = 2; divisor * divisor <= candidate; ++divisor)
            if (candidate % divisor == 0) { prime = false; break; }
        if (prime) total += candidate;
    }
    return total;
}
}

int main() {
    try {
        // A duration in milliseconds or nanoseconds must fail this unit check.
        auto& bench = Bench::foreign();
        using Clock = std::chrono::steady_clock;
        const auto before = Clock::now();
        const double start = purescript::unbox<double>(bench["benchNow"]());
        std::this_thread::sleep_for(std::chrono::milliseconds(20));
        const double finish = purescript::unbox<double>(bench["benchNow"]());
        const double wallUs = std::chrono::duration<double, std::micro>(Clock::now() - before).count();
        if (!(finish - start > wallUs * 0.5 && finish - start < wallUs * 2.0))
            throw std::runtime_error("benchNow must return monotonic microseconds");
        if (purescript::unbox<int>(bench["opaque"](purescript::boxed(123))()) != 123 ||
            purescript::unbox<double>(bench["opaque"](purescript::boxed(1.25))()) != 1.25 ||
            purescript::unbox<std::string>(bench["opaque"](purescript::boxed("ok"))()) != "ok")
            throw std::runtime_error("opaque changed its input");
        std::cout << "C++ clock unit and opaque input checks passed\n";
        for (int n : {0, 1, 2, 3, 5, 8}) CHECK_PAIR(AstTree, n, astValue(n));
        for (int n : {0, 1, 2, 3, 10, 15}) CHECK_PAIR(Fib, n, fibonacci(n));
        for (int n : {0, 1, 2, 3, 7, 32, 899, 900}) {
            const int k = n / 2;
            CHECK_PAIR(ListOps, n, k * (k + 1));
            CHECK_PAIR(ArrayOps, n, k * (k + 1));
        }
        for (int n : {-1, -2, -3, -8}) {
            const int k = (-n) / 2;
            CHECK_PAIR(ArrayOps, n, -k * (k + 1));
        }
        for (int n : {0, 1, 2, 3, 4, 5, 100000}) {
            const int r = n % 3;
            CHECK_PAIR(TCO, n, 3 * (n / 3) + r * (r + 1) / 2);
        }
        for (int n : {0, 1, 4, 5, 6, 19, 10000}) {
            const int r = n % 5;
            CHECK_PAIR(Records, n, 10 * (n / 5) + r * (r + 1) / 2);
        }
        const int ackermann[] = {5, 6, 11, 125}; // A(m, 4), m = 0..3.
        for (int m = 0; m < 4; ++m) CHECK_PAIR(Ackermann, m, ackermann[m]);
        for (int n : {0, 1, 2, 3, 5, 10}) CHECK_PAIR(Church, n, n * n * n * n * n);
        for (int n : {0, 1, 2, 3, 10, 31, 500}) CHECK_PAIR(Primes, n, sumPrimes(n));
        // Perfect Okasaki trees at 2^k-1 sorted insertions; canonical benchmark depth is 22.
        for (int k = 0; k <= 8; ++k) CHECK_PAIR(RBTree, (1 << k) - 1, k);
        CHECK_PAIR(RBTree, 2, 2);
        CHECK_PAIR(RBTree, 100000, 22);
        for (int n : {0, 1, 7, 31, 10000000}) CHECK_PAIR(Polymorphism, n, n);
        for (int depth : {0, 1, 2, 7, 60}) CHECK_PAIR(StateMonad, depth, 20 * depth);
        for (int depth : {0, 1, 2, 7, 1000}) CHECK_PAIR(LazyEvaluation, depth, 1000 * depth);
        for (int dummy : {0, 1, 10000}) CHECK_PAIR(RowToList, dummy, 5);
        std::cout << checks << " native C++ result checks passed\n";
    } catch (const std::exception& error) {
        std::cerr << error.what() << '\n';
        return 1;
    }
}
