#include "purescript.h"
#include <cstdlib>
#include <cstring>

FOREIGN_BEGIN(Main)
exports["smokeRequested"] = []() -> boxed {
    const char* mode = std::getenv("ALTBAK_CPP_SMOKE");
    return mode != nullptr && std::strcmp(mode, "1") == 0;
};
FOREIGN_END
