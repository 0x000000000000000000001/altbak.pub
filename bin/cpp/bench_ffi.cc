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
    text << std::fixed << std::setprecision(2) << unbox<double>(value);
    return text.str();
};
exports["opaque"] = [](const boxed& value) -> boxed {
    return [value]() -> boxed {
        boxed input = value;
        opaque_input(input);
        return input;
    };
};
FOREIGN_END
