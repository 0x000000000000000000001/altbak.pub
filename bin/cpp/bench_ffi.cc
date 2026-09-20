#include "purescript.h"
#include <chrono>
#include <iomanip>
#include <sstream>

namespace {
// Treat the boxed input (including primitive fields) as unknown at every call.
// The runner also compiles separate translation units without LTO.
#if defined(__clang__) || defined(__GNUC__)
__attribute__((noinline))
void opaque_input(purescript::boxed& value) {
    asm volatile("" : "+m"(value) : : "memory");
}
#else
#error "The C++ benchmark requires a compiler supporting the opaque input barrier"
#endif
}

FOREIGN_BEGIN(Bench)
exports["benchNow"] = []() -> boxed {
    using Clock = std::chrono::steady_clock;
    static const auto origin = Clock::now();
    return std::chrono::duration<double, std::micro>(Clock::now() - origin).count();
};
exports["formatNumber"] = [](const boxed& value) -> boxed {
    std::ostringstream text;
    text << std::fixed << std::setprecision(6) << unbox<double>(value);
    return text.str();
};
exports["opaque"] = [](const boxed& value) -> boxed {
    return [value]() -> boxed {
        boxed input = value;
        opaque_input(input);
        return input;
    };
};
exports["measureBatch"] = [](const boxed& count) -> boxed {
    return [count](const boxed& expected) -> boxed {
        return [count, expected](const boxed& act) -> boxed {
            return [count, expected, act]() -> boxed {
                using Clock = std::chrono::steady_clock;
                boxed result = 0;
                const auto start = Clock::now();
                for (int i = 0; i < unbox<int>(count); ++i) {
                    result = act();
                    opaque_input(result);
                }
                const double elapsed = std::chrono::duration<double, std::micro>(Clock::now() - start).count();
                if (unbox<int>(result) != unbox<int>(expected))
                    throw std::runtime_error("Unstable benchmark result");
                return elapsed;
            };
        };
    };
};
FOREIGN_END
