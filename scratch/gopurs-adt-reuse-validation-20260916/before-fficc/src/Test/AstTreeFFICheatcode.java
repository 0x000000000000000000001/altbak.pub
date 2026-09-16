private static final class Expr {
    final int tag;
    final int value;
    final Expr left;
    final Expr right;

    Expr(int tag, int value, Expr left, Expr right) {
        this.tag = tag;
        this.value = value;
        this.left = left;
        this.right = right;
    }
}

private static Expr buildTree(int depth) {
    if (depth == 0) return new Expr(0, 1, null, null);
    return new Expr(1, 0,
        new Expr(2, 0, new Expr(0, depth, null, null), buildTree(depth - 1)),
        new Expr(3, 0, buildTree(depth - 1), new Expr(0, 1, null, null)));
}

private static int eval(Expr expr) {
    return switch (expr.tag) {
        case 0 -> expr.value;
        case 1 -> eval(expr.left) + eval(expr.right);
        case 2 -> eval(expr.left) * eval(expr.right);
        case 3 -> eval(expr.left) - eval(expr.right);
        default -> throw new IllegalStateException("Unknown expression tag");
    };
}

public static final java.util.function.Function<Object, Object> runAstTreeFFICheatcode =
    Bench.nativeBenchmark(input -> eval(buildTree(((Number) input).intValue())));
