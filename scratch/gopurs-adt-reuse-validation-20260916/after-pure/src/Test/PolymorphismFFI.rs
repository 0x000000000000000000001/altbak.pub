struct Dict {
    mempty_: i64,
    mappend_: Box<dyn Fn(i64) -> Box<dyn Fn(i64) -> i64>>,
}
pub fn Test_PolymorphismFFI_runPolymorphismFFI(mut limit: i64) -> i64 {
    let dict = Dict {
        mempty_: 1,
        mappend_: Box::new(|x| Box::new(move |y| x + y)),
    };
    let mut acc = 0;
    while limit > 0 {
        let f = (dict.mappend_)(acc);
        acc = f(dict.mempty_);
        limit -= 1;
    }
    acc
}
