use std::cell::RefCell;
use std::rc::Rc;

enum Lazy<T> {
    Thunk(Box<dyn Fn() -> T>),
    Value(T),
}

fn force<T: Clone>(l: &Rc<RefCell<Lazy<T>>>) -> T {
    let mut lazy = l.borrow_mut();
    match &*lazy {
        Lazy::Value(v) => v.clone(),
        Lazy::Thunk(f) => {
            let v = f();
            *lazy = Lazy::Value(v.clone());
            v
        }
    }
}

pub fn Test_LazyEvaluationFFI_runLazyEvaluationFFI(mut limit: i64) -> i64 {
    let mut lazy_val = Rc::new(RefCell::new(Lazy::Thunk(Box::new(|| 1))));
    for _ in 0..limit {
        let prev = Rc::clone(&lazy_val);
        lazy_val = Rc::new(RefCell::new(Lazy::Thunk(Box::new(move || {
            force(&prev) + 1
        }))));
    }
    force(&lazy_val)
}