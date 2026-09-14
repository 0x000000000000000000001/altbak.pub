// Structural checks for the handwritten Okasaki implementation. This is a
// separate translation unit: including the kernel exposes its private helpers.
#include "../src/Test/RBTreeFFICheatcode.cc"
#include <cstdint>
#include <iostream>
#include <limits>
#include <stdexcept>
#include <string>
#include <vector>

namespace {
void require(bool condition, const char* message) {
    if (!condition) throw std::runtime_error(message);
}

struct Invariants {
    int nodes;
    int black_height;
    int height;
};

Invariants verify(const Tree& tree, long long lower, long long upper) {
    if (!tree) return Invariants{0, 1, 0};
    require(lower < tree->key && tree->key < upper, "keys are not strictly ordered");
    if (tree->color == Red) {
        require(!red(tree->left) && !red(tree->right), "red node has a red child");
    }
    const Invariants left = verify(tree->left, lower, tree->key);
    const Invariants right = verify(tree->right, tree->key, upper);
    require(left.black_height == right.black_height, "black heights differ");
    return Invariants{
        1 + left.nodes + right.nodes,
        left.black_height + (tree->color == Black ? 1 : 0),
        1 + std::max(left.height, right.height)
    };
}

void verify_root(const Tree& tree, int expected_nodes) {
    require(!tree || tree->color == Black, "root is not black");
    const Invariants found = verify(tree,
        std::numeric_limits<long long>::min(), std::numeric_limits<long long>::max());
    require(found.nodes == expected_nodes, "unexpected node count");
    require(found.height == depth(tree), "depth disagrees with structural traversal");
}

std::string shape(const Tree& tree) {
    if (!tree) return "E";
    return std::string(tree->color == Red ? "R(" : "B(") +
        std::to_string(tree->key) + "," + shape(tree->left) + "," +
        shape(tree->right) + ")";
}

void check_rotations() {
    const int orders[4][3] = {
        {3, 2, 1}, // LL
        {3, 1, 2}, // LR
        {1, 3, 2}, // RL
        {1, 2, 3}  // RR
    };
    for (const auto& order : orders) {
        Tree tree;
        int count = 0;
        for (int key : order) {
            tree = insert(std::move(tree), key);
            verify_root(tree, ++count);
        }
        require(shape(tree) == "B(2,B(1,E,E),B(3,E,E))",
            "rotation does not produce the expected Okasaki tree");
    }
}

void check_mixed_and_duplicates() {
    std::vector<int> keys;
    for (int key = 1; key <= 1000; ++key) keys.push_back(key);
    // Explicit xorshift/Fisher-Yates keeps the sequence identical across
    // standard-library implementations, without a random-device dependency.
    std::uint32_t state = 20260914;
    for (std::size_t remaining = keys.size(); remaining > 1; --remaining) {
        state ^= state << 13;
        state ^= state >> 17;
        state ^= state << 5;
        std::swap(keys[remaining - 1], keys[state % remaining]);
    }
    Tree tree;
    int count = 0;
    for (int key : keys) {
        tree = insert(std::move(tree), key);
        verify_root(tree, ++count);
    }
    const std::string before = shape(tree);
    const Node* original_root = tree.get();
    for (int key : keys) {
        tree = insert(std::move(tree), key);
        verify_root(tree, count);
    }
    require(tree.get() == original_root, "duplicate insertion replaced the root");
    require(shape(tree) == before, "duplicate insertion changed tree shape or colors");
}

void check_descending() {
    for (int size : {0, 1, 2, 3, 7, 31, 127, 1000, 100000}) {
        Tree tree;
        for (int key = size; key > 0; --key) tree = insert(std::move(tree), key);
        verify_root(tree, size);
        if (size == 100000) {
            require(depth(tree) == 22, "100000 descending insertions must have depth 22");
        }
    }
}
}

int main() {
    try {
        check_rotations();
        check_mixed_and_duplicates();
        check_descending();
        std::cout << "C++ handwritten RBTree: LL/LR/RL/RR, ordering, colors, black heights, "
                     "counts, deterministic mixed insertions, duplicates, and depth(100000)=22 passed\n";
        return 0;
    } catch (const std::exception& error) {
        std::cerr << "C++ handwritten RBTree check failed: " << error.what() << '\n';
        return 1;
    }
}
