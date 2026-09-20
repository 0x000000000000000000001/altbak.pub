#include "purescript.h"
#include <functional>
#include <string>

namespace {
    struct RowNil {};
    template <class Head, class Tail> struct RowCons { Head head; Tail tail; };
    template <class Row> struct RecordKeys { std::function<int()> keysImpl; };

    RecordKeys<RowNil> keysNil() {
        return RecordKeys<RowNil>{[]() { return 0; }};
    }

    template <class Head, class Tail>
    RecordKeys<RowCons<Head, Tail>> keysCons(RecordKeys<Tail> tail) {
        return RecordKeys<RowCons<Head, Tail>>{[tail]() { return 1 + tail.keysImpl(); }};
    }

    // Native C++ representation of type-directed instance resolution:
    // the dictionary shape is selected from the heterogeneous record type.
    template <class Row> struct Instances;
    template <> struct Instances<RowNil> {
        static RecordKeys<RowNil> dictionary() { return keysNil(); }
    };
    template <class Head, class Tail> struct Instances<RowCons<Head, Tail>> {
        static RecordKeys<RowCons<Head, Tail>> dictionary() {
            return keysCons<Head>(Instances<Tail>::dictionary());
        }
    };

    template <class Row>
    int keys(const RecordKeys<Row>& dictionary, const Row&) {
        return dictionary.keysImpl();
    }
}

FOREIGN_BEGIN(Test_RowToListFFI)
exports["runRowToListFFI"] = [](const boxed&) -> boxed {
    const RowCons<int, RowCons<std::string, RowCons<bool, RowCons<double, RowCons<std::string, RowNil>>>>>
        record{1, {"two", {true, {4.0, {"five", {}}}}}};
    return keys(Instances<decltype(record)>::dictionary(), record);
};
FOREIGN_END
