use std::rc::Rc;

type Lazy<'a, A> = Rc<dyn Fn() -> A + 'a>;

fn force<A>(l: &Lazy<'_, A>) -> A {
    l()
}

fn defer<'a, A, F: Fn() -> A + 'a>(f: F) -> Lazy<'a, A> {
    Rc::new(f)
}

fn build_thunks<'a>(depth: i64, acc: Lazy<'a, i64>) -> Lazy<'a, i64> {
    if depth == 0 {
        return acc;
    }
    let prev = acc.clone();
    build_thunks(depth - 1, defer(move || force(&prev) + 1))
}

pub fn Test_LazyEvaluationFFI_runLazyEvaluationFFI(limit: i64) -> i64 {
    let mut acc = 0;
    for _ in 0..limit {
        let t = build_thunks(1000, defer(|| 0));
        acc += force(&t);
    }
    acc
}
