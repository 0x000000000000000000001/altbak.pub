#include "purescript.h"
#include <memory>

namespace {
    struct Inner { int e, f; };
    struct Middle { int c; std::shared_ptr<const Inner> d; };
    struct Record { int a; std::shared_ptr<const Middle> b; };
    using RecordPtr = std::shared_ptr<const Record>;

    RecordPtr initial() {
        const auto d = std::make_shared<Inner>(Inner{0, 0});
        const auto b = std::make_shared<Middle>(Middle{0, d});
        return std::make_shared<Record>(Record{0, b});
    }

    RecordPtr updateRec(int n, RecordPtr record) {
        while (n > 0) {
            const auto d = std::make_shared<Inner>(Inner{
                record->b->d->e + 3, record->b->d->f + n % 5});
            const auto b = std::make_shared<Middle>(Middle{record->b->c + 2, d});
            record = std::make_shared<Record>(Record{record->a + 1, b});
            --n;
        }
        return record;
    }
}

FOREIGN_BEGIN(Test_RecordsFFI)
exports["runRecordsFFI"] = [](const boxed& in) -> boxed {
    const auto result = updateRec(unbox<int>(in), initial());
    const int value = result->b->d->f;
    return value;
};
FOREIGN_END
