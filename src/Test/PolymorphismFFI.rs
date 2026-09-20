struct Monoidish<A> {
    mempty_: A,
    mappend_: Box<dyn Fn(A) -> Box<dyn Fn(A) -> A>>,
}

fn poly_loop<A: Clone>(dict: &Monoidish<A>, mut limit: i64, mut acc: A) -> A {
    while limit > 0 {
        let f = (dict.mappend_)(acc);
        acc = f(dict.mempty_.clone());
        limit -= 1;
    }
    acc
}

pub fn Test_PolymorphismFFI_runPolymorphismFFI(limit: i64) -> i64 {
    let int_monoidish = Monoidish {
        mempty_: 1,
        mappend_: Box::new(|x| Box::new(move |y| x + y)),
    };
    poly_loop(&int_monoidish, limit, 0)
}
