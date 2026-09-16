private static final class Tree {
    boolean black;
    Tree left;
    final int value;
    Tree right;

    Tree(int value) {
        this.value = value;
    }
}

private static boolean red(Tree tree) {
    return tree != null && !tree.black;
}

// The same Okasaki rotations as the functional implementation, reusing nodes.
private static Tree balance(Tree tree) {
    if (!tree.black) return tree;
    Tree left = tree.left;
    Tree right = tree.right;
    if (red(left)) {
        if (red(left.left)) {
            tree.left = left.right;
            left.right = tree;
            left.left.black = true;
            left.black = false;
            return left;
        }
        if (red(left.right)) {
            Tree middle = left.right;
            left.right = middle.left;
            tree.left = middle.right;
            middle.left = left;
            middle.right = tree;
            left.black = true;
            middle.black = false;
            return middle;
        }
    }
    if (red(right)) {
        if (red(right.left)) {
            Tree middle = right.left;
            tree.right = middle.left;
            right.left = middle.right;
            middle.left = tree;
            middle.right = right;
            right.black = true;
            middle.black = false;
            return middle;
        }
        if (red(right.right)) {
            tree.right = right.left;
            right.left = tree;
            right.right.black = true;
            right.black = false;
            return right;
        }
    }
    return tree;
}

private static Tree insertRed(int value, Tree tree) {
    if (tree == null) return new Tree(value);
    if (value < tree.value) tree.left = insertRed(value, tree.left);
    else if (value > tree.value) tree.right = insertRed(value, tree.right);
    else return tree;
    return balance(tree);
}

private static Tree insert(int value, Tree tree) {
    Tree result = insertRed(value, tree);
    result.black = true;
    return result;
}

private static int depth(Tree tree) {
    return tree == null ? 0 : 1 + Math.max(depth(tree.left), depth(tree.right));
}

public static final java.util.function.Function<Object, Object> runRBTreeFFICheatcode = Bench.nativeBenchmark(input -> {
    Tree tree = null;
    for (int value = ((Number) input).intValue(); value > 0; value--) {
        tree = insert(value, tree);
    }
    return depth(tree);
});
