pub fn Test_StateMonadFFI_runStateMonadFFI(mut limit: i64) -> i64 {
    type State<'a> = Box<dyn Fn(i64) -> (i64, i64) + 'a>;
    
    fn pure_state<'a>(a: i64) -> State<'a> {
        Box::new(move |s| (a, s))
    }
    
    fn bind_state<'a, F>(m: State<'a>, k: F) -> State<'a>
    where
        F: Fn(i64) -> State<'a> + 'a,
    {
        Box::new(move |s| {
            let (a, s1) = m(s);
            k(a)(s1)
        })
    }
    
    fn get_state<'a>() -> State<'a> {
        Box::new(|s| (s, s))
    }
    
    fn put_state<'a>(s: i64) -> State<'a> {
        Box::new(move |_| (0, s))
    }

    let mut state = pure_state(0);
    for _ in 0..limit {
        state = bind_state(state, move |_| {
            bind_state(get_state(), move |s| {
                put_state(s + 1)
            })
        });
    }
    
    let (_, final_s) = state(0);
    final_s
}