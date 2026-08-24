
#[derive(Clone)]
enum Expr {
    Val(i64),
    Add(Box<Expr>, Box<Expr>),
    Mul(Box<Expr>, Box<Expr>),
    Sub(Box<Expr>, Box<Expr>),
}

fn eval_ast(e: &Expr) -> i64 {
    match e {
        Expr::Val(v) => *v,
        Expr::Add(l, r) => eval_ast(l) + eval_ast(r),
        Expr::Mul(l, r) => eval_ast(l) * eval_ast(r),
        Expr::Sub(l, r) => eval_ast(l) - eval_ast(r),
    }
}

fn build_tree_ast(n: i64) -> Box<Expr> {
    if n == 0 {
        return Box::new(Expr::Val(1));
    }
    let val_n = Box::new(Expr::Val(n));
    let val_1 = Box::new(Expr::Val(1));
    let left_tree = build_tree_ast(n - 1);
    let right_tree = build_tree_ast(n - 1);
    let mul_node = Box::new(Expr::Mul(val_n, left_tree));
    let sub_node = Box::new(Expr::Sub(right_tree, val_1));
    Box::new(Expr::Add(mul_node, sub_node))
}

pub fn Test_AstTreeFFICheatcode_runAstTreeFFICheatcode(mut limit: i64) -> i64 {
    let t = build_tree_ast(limit);
    eval_ast(&t)
}