#include "purescript.h"
#include <chrono>
#include <iostream>
#include <iomanip>
#include <sstream>

FOREIGN_BEGIN( Bench )

exports["benchNow"] = []() -> boxed {
    auto now = std::chrono::high_resolution_clock::now().time_since_epoch();
    double ms = std::chrono::duration_cast<std::chrono::nanoseconds>(now).count() / 1e6;
    return ms;
};

exports["formatNumber"] = [](const boxed& n) -> boxed {
    double val = unbox<double>(n);
    std::ostringstream ss;
    ss << std::fixed << std::setprecision(2) << val;
    return ss.str();
};

exports["opaque"] = [](const boxed& a) -> boxed {
    return [=]() -> boxed {
        return a;
    };
};

FOREIGN_END
