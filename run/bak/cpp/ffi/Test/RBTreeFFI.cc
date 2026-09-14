#include "purescript.h"
#include <memory>
#include <algorithm>
namespace {
    volatile int dummy;
    struct Tree { bool black; std::shared_ptr<Tree> left; int val; std::shared_ptr<Tree> right; };
    auto red(std::shared_ptr<Tree> t) -> bool { return t && !t->black; }
    auto balance(bool b, std::shared_ptr<Tree> l, int v, std::shared_ptr<Tree> r) -> std::shared_ptr<Tree> {
        if (b) {
            if (red(l)) {
                if (red(l->left)) return std::make_shared<Tree>(Tree{false, std::make_shared<Tree>(Tree{true, l->left->left, l->left->val, l->left->right}), l->val, std::make_shared<Tree>(Tree{true, l->right, v, r})});
                if (red(l->right)) return std::make_shared<Tree>(Tree{false, std::make_shared<Tree>(Tree{true, l->left, l->val, l->right->left}), l->right->val, std::make_shared<Tree>(Tree{true, l->right->right, v, r})});
            }
            if (red(r)) {
                if (red(r->left)) return std::make_shared<Tree>(Tree{false, std::make_shared<Tree>(Tree{true, l, v, r->left->left}), r->left->val, std::make_shared<Tree>(Tree{true, r->left->right, r->val, r->right})});
                if (red(r->right)) return std::make_shared<Tree>(Tree{false, std::make_shared<Tree>(Tree{true, l, v, r->left}), r->val, std::make_shared<Tree>(Tree{true, r->right->left, r->right->val, r->right->right})});
            }
        }
        return std::make_shared<Tree>(Tree{b, l, v, r});
    }
    auto insRed(int v, std::shared_ptr<Tree> t) -> std::shared_ptr<Tree> {
        if(!t) return std::make_shared<Tree>(Tree{false, nullptr, v, nullptr});
        if(v < t->val) return balance(t->black, insRed(v, t->left), t->val, t->right);
        if(v > t->val) return balance(t->black, t->left, t->val, insRed(v, t->right));
        return std::make_shared<Tree>(Tree{t->black, t->left, t->val, t->right});
    }
    auto insert(int v, std::shared_ptr<Tree> t) -> std::shared_ptr<Tree> {
        auto res = insRed(v, t);
        return std::make_shared<Tree>(Tree{true, res->left, res->val, res->right});
    }
    auto depth(std::shared_ptr<Tree> t) -> int { return !t ? 0 : 1 + std::max(depth(t->left), depth(t->right)); }
}
FOREIGN_BEGIN( Test_RBTreeFFI )
exports["runRBTreeFFI"] = [](const boxed& in) -> boxed { 
    int n = unbox<int>(in); std::shared_ptr<Tree> t = nullptr;
    for(int i=n; i>0; i--) t = insert(i, t);
    int r = depth(t); dummy = r; return r; 
};
FOREIGN_END
