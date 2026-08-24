use std::rc::Rc;

type Lazy<'a> = Rc<dyn Fn() -> i64 + 'a>;

fn force(l: &Lazy) -> i64 {
    l()
}

fn defer<'a, F: Fn() -> i64 + 'a>(f: F) -> Lazy<'a> {
    Rc::new(f)
}

fn build_thunks<'a>(depth: i64, acc: Lazy<'a>) -> Lazy<'a> {
    if depth == 0 {
        return acc;
    }
    let prev = acc.clone();
    build_thunks(depth - 1, defer(move || force(&prev) + 1))
}

pub fn Test_LazyEvaluationFFI_runLazyEvaluationFFI(mut limit: i64) -> i64 {
    let mut acc = 0;
    for _ in 0..limit {
        let t = build_thunks(1000, defer(|| 0));
        acc += force(&t);
    }
    acc
}
