#include "purescript.h"
#include <memory>
namespace {
    volatile int dummy;
    struct Expr { int tag; int val; std::shared_ptr<Expr> left, right; };
    auto Val(int v) -> std::shared_ptr<Expr> { return std::make_shared<Expr>(Expr{0, v, nullptr, nullptr}); }
    auto Add(std::shared_ptr<Expr> l, std::shared_ptr<Expr> r) -> std::shared_ptr<Expr> { return std::make_shared<Expr>(Expr{1, 0, l, r}); }
    auto Mul(std::shared_ptr<Expr> l, std::shared_ptr<Expr> r) -> std::shared_ptr<Expr> { return std::make_shared<Expr>(Expr{2, 0, l, r}); }
    auto Sub(std::shared_ptr<Expr> l, std::shared_ptr<Expr> r) -> std::shared_ptr<Expr> { return std::make_shared<Expr>(Expr{3, 0, l, r}); }
    auto build(int d) -> std::shared_ptr<Expr> {
        if(d == 0) return Val(1);
        return Add(Mul(Val(d), build(d-1)), Sub(build(d-1), Val(1)));
    }
    auto eval(std::shared_ptr<Expr> e) -> int {
        if(e->tag == 0) return e->val;
        if(e->tag == 1) return eval(e->left) + eval(e->right);
        if(e->tag == 2) return eval(e->left) * eval(e->right);
        return eval(e->left) - eval(e->right);
    }
}
FOREIGN_BEGIN( Test_AstTreeFFI )
exports["runAstTreeFFI"] = [](const boxed& in) -> boxed { int r = eval(build(unbox<int>(in))); dummy = r; return r; };
FOREIGN_END
