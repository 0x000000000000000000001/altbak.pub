#include "purescript.h"
#include <memory>
#include <utility>
namespace {
struct Expr {
    int tag;
    int value;
    std::unique_ptr<Expr> left;
    std::unique_ptr<Expr> right;
};
using Tree = std::unique_ptr<Expr>;
Tree val(int value) { return Tree(new Expr{0, value, nullptr, nullptr}); }
Tree binary(int tag, Tree left, Tree right) {
    return Tree(new Expr{tag, 0, std::move(left), std::move(right)});
}
Tree build(int depth) {
    if (depth == 0) return val(1);
    return binary(1, binary(2, val(depth), build(depth - 1)),
                     binary(3, build(depth - 1), val(1)));
}
int eval(const Expr& expr) {
    if (expr.tag == 0) return expr.value;
    if (expr.tag == 1) return eval(*expr.left) + eval(*expr.right);
    if (expr.tag == 2) return eval(*expr.left) * eval(*expr.right);
    return eval(*expr.left) - eval(*expr.right);
}
}
FOREIGN_BEGIN( Test_AstTreeFFICheatcode )
exports["runAstTreeFFICheatcode"] = [](const boxed& in) -> boxed {
    const Tree tree = build(unbox<int>(in));
    return eval(*tree);
};
FOREIGN_END
