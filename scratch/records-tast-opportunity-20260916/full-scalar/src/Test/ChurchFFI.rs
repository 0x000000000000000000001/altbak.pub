use std::sync::Arc;

type IntFn = Arc<dyn Fn(i64) -> i64>;
type Church = Arc<dyn Fn(IntFn) -> IntFn>;

fn zero() -> Church {
    Arc::new(|_f| Arc::new(|x| x))
}

fn succ(n: Church) -> Church {
    Arc::new(move |f| {
        let n = n.clone();
        Arc::new(move |x| f(n(f.clone())(x)))
    })
}

fn add_c(m: Church, n: Church) -> Church {
    Arc::new(move |f| {
        let m = m.clone();
        let n = n.clone();
        Arc::new(move |x| m(f.clone())(n(f.clone())(x)))
    })
}

fn mul_c(m: Church, n: Church) -> Church {
    Arc::new(move |f| {
        let m = m.clone();
        let n = n.clone();
        Arc::new(move |x| m(n(f.clone()))(x))
    })
}

fn from_int(n: i64) -> Church {
    if n == 0 {
        zero()
    } else {
        succ(from_int(n - 1))
    }
}

fn to_int(n: Church) -> i64 {
    let f: IntFn = Arc::new(|x| x + 1);
    n(f)(0)
}

fn c10(limit: i64) -> Church {
    from_int(limit) // Note: using limit here to match benchmark behavior
}

fn c100(limit: i64) -> Church {
    mul_c(c10(limit), c10(limit))
}

fn c10k(limit: i64) -> Church {
    mul_c(c100(limit), c100(limit))
}

fn c100k(limit: i64) -> Church {
    mul_c(c10k(limit), c10(limit))
}

pub fn Test_ChurchFFI_runChurchFFI(mut limit: i64) -> i64 {
    to_int(c100k(limit))
}
