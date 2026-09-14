#include "purescript.h"
#include <memory>
namespace {
    volatile int dummy;
    struct List { int head; std::shared_ptr<List> tail; };
    auto cons(int h, std::shared_ptr<List> t) -> std::shared_ptr<List> { return std::make_shared<List>(List{h, t}); }
    auto filter(int p, std::shared_ptr<List> l) -> std::shared_ptr<List> {
        if (!l) return nullptr;
        if (l->head % p != 0) return cons(l->head, filter(p, l->tail));
        return filter(p, l->tail);
    }
    auto sieve(std::shared_ptr<List> l) -> std::shared_ptr<List> {
        if (!l) return nullptr;
        return cons(l->head, sieve(filter(l->head, l->tail)));
    }
}
FOREIGN_BEGIN( Test_PrimesFFI )
exports["runPrimesFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in);
    std::shared_ptr<List> l = nullptr;
    for(int i = n; i >= 2; i--) l = cons(i, l);
    auto p = sieve(l);
    int count = 0; auto curr = p; while(curr) { count++; curr = curr->tail; }
    dummy = count; return count;
};
FOREIGN_END
