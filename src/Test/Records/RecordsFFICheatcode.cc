#include "purescript.h"
namespace {
struct Inner { int e, f; };
struct Middle { int c; Inner d; };
struct Record { int a; Middle b; };
}
FOREIGN_BEGIN( Test_RecordsFFICheatcode )
exports["runRecordsFFICheatcode"] = [](const boxed& in) -> boxed {
    Record record{0, {0, {0, 0}}};
    for (int n = unbox<int>(in); n > 0; --n) {
        record.a += 1;
        record.b.c += 2;
        record.b.d.e += 3;
        record.b.d.f += n % 5;
    }
    return record.b.d.f;
};
FOREIGN_END
