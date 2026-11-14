private sealed interface Expr permits Val, Add, Mul, Sub {}
private record Val(int value) implements Expr {}
private record Add(Expr left, Expr right) implements Expr {}
private record Mul(Expr left, Expr right) implements Expr {}
private record Sub(Expr left, Expr right) implements Expr {}

private static Expr buildTree(int depth) {
    if (depth == 0) return new Val(1);
    return new Add(
        new Mul(new Val(depth), buildTree(depth - 1)),
        new Sub(buildTree(depth - 1), new Val(1)));
}

private static int eval(Expr expr) {
    if (expr instanceof Val value) return value.value();
    if (expr instanceof Add add) return eval(add.left()) + eval(add.right());
    if (expr instanceof Mul mul) return eval(mul.left()) * eval(mul.right());
    Sub sub = (Sub) expr;
    return eval(sub.left()) - eval(sub.right());
}

public static final java.util.function.Function<Object, Object> runAstTreeFFI =
    Bench.nativeBenchmark(input -> eval(buildTree(((Number) input).intValue())));
