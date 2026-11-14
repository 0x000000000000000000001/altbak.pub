#include "purescript.h"
#include <tuple>
FOREIGN_BEGIN( Test_RowToListFFICheatcode )
exports["runRowToListFFICheatcode"] = [](const boxed&) -> boxed {
    // Native heterogeneous product corresponding to the fixed record's fields.
    const auto record = std::make_tuple(1, "two", true, 4.0, "five");
    return static_cast<int>(std::tuple_size<decltype(record)>::value);
};
FOREIGN_END
