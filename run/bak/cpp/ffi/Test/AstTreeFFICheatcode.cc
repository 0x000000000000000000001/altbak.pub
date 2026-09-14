#include "purescript.h"
namespace {
    volatile int dummy;
    struct Expr { int tag; int val; Expr* left; Expr* right; };
    auto Val(int v) -> Expr* { return new Expr{0, v, nullptr, nullptr}; }
    auto Add(Expr* l, Expr* r) -> Expr* { return new Expr{1, 0, l, r}; }
    auto Mul(Expr* l, Expr* r) -> Expr* { return new Expr{2, 0, l, r}; }
    auto Sub(Expr* l, Expr* r) -> Expr* { return new Expr{3, 0, l, r}; }
    auto build(int d) -> Expr* {
        if(d == 0) return Val(1);
        return Add(Mul(Val(d), build(d-1)), Sub(build(d-1), Val(1)));
    }
    auto eval(Expr* e) -> int {
        if(e->tag == 0) return e->val;
        if(e->tag == 1) return eval(e->left) + eval(e->right);
        if(e->tag == 2) return eval(e->left) * eval(e->right);
        return eval(e->left) - eval(e->right);
    }
}
FOREIGN_BEGIN( Test_AstTreeFFICheatcode )
exports["runAstTreeFFICheatcode"] = [](const boxed& in) -> boxed { int r = eval(build(unbox<int>(in))); dummy = r; return r; };
FOREIGN_END
