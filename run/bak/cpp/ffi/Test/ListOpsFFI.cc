#include "purescript.h"
#include <memory>
namespace {
    volatile int dummy;
    struct List { int head; std::shared_ptr<List> tail; };
    auto cons(int h, std::shared_ptr<List> t) -> std::shared_ptr<List> { return std::make_shared<List>(List{h, t}); }
}
FOREIGN_BEGIN( Test_ListOpsFFI )
exports["runListOpsFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in);
    std::shared_ptr<List> l = nullptr;
    for (int i = 0; i < n; i++) l = cons(i, l);
    std::shared_ptr<List> l2 = nullptr;
    auto curr = l; while(curr) { l2 = cons(curr->head * 2, l2); curr = curr->tail; }
    int sum = 0; curr = l2; while(curr) { sum += curr->head; curr = curr->tail; }
    dummy = sum; return sum;
};
FOREIGN_END
