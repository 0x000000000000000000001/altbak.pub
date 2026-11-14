use std::rc::Rc;

type Step<A> = Rc<dyn Fn(A) -> A>;
type Church<A> = Rc<dyn Fn(Step<A>) -> Step<A>>;

fn zero<A: 'static>() -> Church<A> {
    Rc::new(|_f| Rc::new(|x| x))
}

fn succ<A: 'static>(n: Church<A>) -> Church<A> {
    Rc::new(move |f| {
        let n = n.clone();
        Rc::new(move |x| f(n(f.clone())(x)))
    })
}

fn add_c<A: 'static>(m: Church<A>, n: Church<A>) -> Church<A> {
    Rc::new(move |f| {
        let m = m.clone();
        let n = n.clone();
        Rc::new(move |x| m(f.clone())(n(f.clone())(x)))
    })
}

fn mul_c<A: 'static>(m: Church<A>, n: Church<A>) -> Church<A> {
    Rc::new(move |f| {
        let m = m.clone();
        let n = n.clone();
        Rc::new(move |x| m(n(f.clone()))(x))
    })
}

fn from_int(n: i64) -> Church<i64> {
    if n == 0 {
        zero()
    } else {
        succ(from_int(n - 1))
    }
}

fn to_int(n: Church<i64>) -> i64 {
    let f: Step<i64> = Rc::new(|x| x + 1);
    n(f)(0)
}

fn c10(limit: i64) -> Church<i64> {
    from_int(limit)
}

fn c100(limit: i64) -> Church<i64> {
    mul_c(c10(limit), c10(limit))
}

fn c10k(limit: i64) -> Church<i64> {
    mul_c(c100(limit), c100(limit))
}

fn c100k(limit: i64) -> Church<i64> {
    mul_c(c10k(limit), c10(limit))
}

pub fn Test_ChurchFFI_runChurchFFI(limit: i64) -> i64 {
    to_int(c100k(limit))
}
