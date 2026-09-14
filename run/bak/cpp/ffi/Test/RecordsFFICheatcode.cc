#include "purescript.h"
namespace {
    volatile int dummy;
    struct Point { int x, y, z; };
}
FOREIGN_BEGIN( Test_RecordsFFICheatcode )
exports["runRecordsFFICheatcode"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); Point p{0, 0, 0};
    for(int i=0; i<n; i++) { p.x+=1; p.y+=2; p.z+=3; }
    int r = p.x + p.y + p.z; dummy = r; return r;
};
FOREIGN_END
