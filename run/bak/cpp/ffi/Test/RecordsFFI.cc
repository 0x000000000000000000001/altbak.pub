#include "purescript.h"
#include <memory>
namespace {
    volatile int dummy;
    struct Point { int x, y, z; };
    auto update(std::shared_ptr<Point> p) -> std::shared_ptr<Point> { return std::make_shared<Point>(Point{p->x + 1, p->y + 2, p->z + 3}); }
}
FOREIGN_BEGIN( Test_RecordsFFI )
exports["runRecordsFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in);
    auto p = std::make_shared<Point>(Point{0, 0, 0});
    for (int i = 0; i < n; i++) p = update(p);
    int r = p->x + p->y + p->z; dummy = r; return r;
};
FOREIGN_END
