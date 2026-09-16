#include "purescript.h"
#include <algorithm>
#include <memory>
#include <utility>

namespace {
enum Color { Red, Black };
struct Node {
    Color color;
    int key;
    std::unique_ptr<Node> left;
    std::unique_ptr<Node> right;
};
using Tree = std::unique_ptr<Node>;

bool red(const Tree& tree) { return tree && tree->color == Red; }

Tree rotate_left(Tree tree) {
    Tree root = std::move(tree->right);
    tree->right = std::move(root->left);
    root->left = std::move(tree);
    return root;
}

Tree rotate_right(Tree tree) {
    Tree root = std::move(tree->left);
    tree->left = std::move(root->right);
    root->right = std::move(tree);
    return root;
}

Tree recolor(Tree tree) {
    tree->color = Red;
    tree->left->color = Black;
    tree->right->color = Black;
    return tree;
}

Tree balance(Tree tree) {
    if (tree->color != Black) return tree;
    // The four Okasaki cases, in the same order as Test.RBTree.balance.
    if (red(tree->left) && red(tree->left->left)) {
        return recolor(rotate_right(std::move(tree)));
    }
    if (red(tree->left) && red(tree->left->right)) {
        tree->left = rotate_left(std::move(tree->left));
        return recolor(rotate_right(std::move(tree)));
    }
    if (red(tree->right) && red(tree->right->left)) {
        tree->right = rotate_right(std::move(tree->right));
        return recolor(rotate_left(std::move(tree)));
    }
    if (red(tree->right) && red(tree->right->right)) {
        return recolor(rotate_left(std::move(tree)));
    }
    return tree;
}

Tree ins(Tree tree, int key) {
    if (!tree) return Tree(new Node{Red, key, nullptr, nullptr});
    if (key < tree->key) tree->left = ins(std::move(tree->left), key);
    else if (key > tree->key) tree->right = ins(std::move(tree->right), key);
    else return tree;
    return balance(std::move(tree));
}

Tree insert(Tree tree, int key) {
    tree = ins(std::move(tree), key);
    tree->color = Black;
    return tree;
}

int depth(const Tree& tree) {
    if (!tree) return 0;
    return 1 + std::max(depth(tree->left), depth(tree->right));
}
}

FOREIGN_BEGIN( Test_RBTreeFFICheatcode )
exports["runRBTreeFFICheatcode"] = [](const boxed& in) -> boxed {
    Tree tree;
    for (int n = unbox<int>(in); n > 0; --n) tree = insert(std::move(tree), n);
    return depth(tree);
};
FOREIGN_END
