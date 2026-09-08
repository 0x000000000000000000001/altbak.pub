private static final class Tree {
    final boolean black;
    final Tree left;
    final int value;
    final Tree right;

    Tree(boolean black, Tree left, int value, Tree right) {
        this.black = black;
        this.left = left;
        this.value = value;
        this.right = right;
    }
}

private static boolean red(Tree tree) {
    return tree != null && !tree.black;
}

private static Tree balance(boolean black, Tree a, int x, Tree b) {
    if (black) {
        if (red(a)) {
            if (red(a.left)) {
                Tree child = a.left;
                return new Tree(false,
                    new Tree(true, child.left, child.value, child.right), a.value,
                    new Tree(true, a.right, x, b));
            }
            if (red(a.right)) {
                Tree child = a.right;
                return new Tree(false,
                    new Tree(true, a.left, a.value, child.left), child.value,
                    new Tree(true, child.right, x, b));
            }
        }
        if (red(b)) {
            if (red(b.left)) {
                Tree child = b.left;
                return new Tree(false,
                    new Tree(true, a, x, child.left), child.value,
                    new Tree(true, child.right, b.value, b.right));
            }
            if (red(b.right)) {
                Tree child = b.right;
                return new Tree(false,
                    new Tree(true, a, x, b.left), b.value,
                    new Tree(true, child.left, child.value, child.right));
            }
        }
    }
    return new Tree(black, a, x, b);
}

private static Tree insertRed(int value, Tree tree) {
    if (tree == null) return new Tree(false, null, value, null);
    if (value < tree.value) {
        return balance(tree.black, insertRed(value, tree.left), tree.value, tree.right);
    }
    if (value > tree.value) {
        return balance(tree.black, tree.left, tree.value, insertRed(value, tree.right));
    }
    return new Tree(tree.black, tree.left, tree.value, tree.right);
}

private static Tree insert(int value, Tree tree) {
    Tree result = insertRed(value, tree);
    return new Tree(true, result.left, result.value, result.right);
}

private static int depth(Tree tree) {
    return tree == null ? 0 : 1 + Math.max(depth(tree.left), depth(tree.right));
}

public static final java.util.function.Function<Object, Object> runRBTreeFFI = Bench.nativeBenchmark(input -> {
    Tree tree = null;
    for (int value = ((Number) input).intValue(); value > 0; value--) {
        tree = insert(value, tree);
    }
    return depth(tree);
});
