// Compile the exact FP kernels in separate namespaces to exercise their private
// polymorphic helpers with nonnumeric values, outside the benchmark harness.
#include "purescript.h"
#include <functional>
#include <memory>
#include <string>
#include <vector>
#include <stdexcept>
#include <iostream>

void require(bool condition, const char* message) {
    if (!condition) throw std::runtime_error(message);
}

namespace list_checks {
#include "../src/Test/ListOpsFFI.cc"
void check() {
    const auto words = cons(std::string("a"), cons(std::string("b"), List<std::string>{}));
    require(foldl<std::string, std::string>([](std::string a, std::string b) { return a + b; }, "!", words) == "!ab",
        "generic list/fold");
    require(foldl<int, std::string>([](std::string a, int b) { return a + std::to_string(b); }, "", filterEvens(range(1, 6))) == "642",
        "reversed even filter");
}
}
namespace prime_checks {
#include "../src/Test/PrimesFFI.cc"
void check() {
    const auto words = cons(std::string("a"), cons(std::string("skip"), cons(std::string("b"), List<std::string>{})));
    const auto selected = filter<std::string>([](std::string value) { return value != "skip"; }, words);
    require(selected->head == "a" && selected->tail->head == "b" && !selected->tail->tail, "generic filter/reverse");
}
}
namespace church_checks {
#include "../src/Test/ChurchFFI.cc"
void check() {
    const auto two = successor<std::string>(successor<std::string>(zero<std::string>()));
    int calls = 0;
    const Endo<std::string> append = std::make_shared<std::function<std::string(std::string)>>(
        [&calls](std::string value) { ++calls; return value + "!"; });
    const auto applied = (*multiply<std::string>(two, two))(append);
    require(calls == 0, "Church application must wait for its value");
    require((*applied)("s") == "s!!!!" && calls == 4, "generic Church");
}
}
namespace polymorphism_checks {
#include "../src/Test/PolymorphismFFI.cc"
void check() {
    const Monoidish<std::string> dictionary{"x", [](std::string left) -> Endo<std::string> {
        return [left](std::string right) { return left + right; };
    }};
    require(polyLoop(dictionary, 3, std::string("s")) == "sxxx", "generic dictionary");
}
}
namespace state_checks {
#include "../src/Test/StateMonadFFI.cc"
void check() {
    const auto action = bindState<std::string, std::string, int>(get<std::string>(), [](std::string initial) {
        return bindState<std::string, Unit, int>(put<std::string>(initial + "!"), [initial](Unit) {
            return pureState<std::string, int>(int(initial.size()));
        });
    });
    const auto result = (*action)("abc");
    require(result.val == 3 && result.state == "abc!", "generic State");
}
}
namespace lazy_checks {
#include "../src/Test/LazyEvaluationFFI.cc"
void check() {
    int calls = 0;
    const auto value = defer<std::string>([&calls]() { ++calls; return std::string("word"); });
    require(force<std::string>(value) == "word" && force<std::string>(value) == "word" && calls == 2,
        "non-memoizing generic Lazy");
}
}
namespace array_checks {
#include "../src/Test/ArrayOpsFFI.cc"
void check() {
    const std::vector<std::string> words{"a", "skip", "b"};
    const auto selected = filter([](std::string value) { return value != "skip"; }, words);
    require(foldl([](std::string a, std::string b) { return a + b; }, std::string("!"), selected) == "!ab",
        "generic array filter/fold");
    require(words.size() == 3 && words[1] == "skip", "array filter must preserve input");
}
}
namespace row_checks {
#include "../src/Test/RowToListFFI.cc"
void check() {
    const RowNil empty{};
    const RowCons<std::string, RowNil> single{"one", empty};
    const RowCons<bool, decltype(single)> pair{true, single};
    require(keys(keysNil(), empty) == 0 && keys(keysCons<std::string>(keysNil()), single) == 1,
        "empty and single rows");
    using Pair = RowCons<bool, RowCons<std::string, RowNil>>;
    require(keys(Instances<Pair>::dictionary(), Pair{true, single}) == 2, "typed heterogeneous row");
}
}

int main() {
    list_checks::check();
    prime_checks::check();
    church_checks::check();
    polymorphism_checks::check();
    state_checks::check();
    lazy_checks::check();
    array_checks::check();
    row_checks::check();
    std::cout << "C++ generic FP structures, non-memoizing Lazy and typed rows passed\n";
}
