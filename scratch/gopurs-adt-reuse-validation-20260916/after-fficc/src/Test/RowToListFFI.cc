#include "purescript.h"

namespace {
    template <int Fields>
    struct RecordKeys {
        static int count() { return 1 + RecordKeys<Fields - 1>::count(); }
    };

    template <>
    struct RecordKeys<0> {
        static int count() { return 0; }
    };
}

FOREIGN_BEGIN(Test_RowToListFFI)
exports["runRowToListFFI"] = [](const boxed&) -> boxed {
    // Test.RowToList has five fields; the opaque argument is not a field count.
    const int result = RecordKeys<5>::count();
    return result;
};
FOREIGN_END
