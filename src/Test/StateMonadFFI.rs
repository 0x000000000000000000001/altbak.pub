struct StateResult {
    val: i64,
    state: i64,
}

type StateFn = Box<dyn Fn(i64) -> StateResult>;

fn run_state(s: &StateFn, init: i64) -> StateResult {
    s(init)
}

fn bind_state(s: StateFn, g: Box<dyn Fn(i64) -> StateFn>) -> StateFn {
    Box::new(move |state: i64| {
        let r1 = s(state);
        let g_prime = g(r1.val);
        g_prime(r1.state)
    })
}

fn pure_state(a: i64) -> StateFn {
    Box::new(move |s: i64| StateResult { val: a, state: s })
}

fn get() -> StateFn {
    Box::new(move |s: i64| StateResult { val: s, state: s })
}

fn put(s: i64) -> StateFn {
    Box::new(move |_| StateResult { val: 0, state: s })
}

fn modify(f: Box<dyn Fn(i64) -> i64>) -> StateFn {
    bind_state(
        get(),
        Box::new(move |s: i64| put(f(s)))
    )
}

fn chain_modifications(n: i64) -> StateFn {
    if n == 0 {
        return pure_state(0);
    }
    bind_state(
        modify(Box::new(|x| x + 1)),
        Box::new(move |_| chain_modifications(n - 1))
    )
}

fn run_many_times_state_monad(n: i64, acc: i64) -> i64 {
    if n == 0 {
        return acc;
    }
    let s = chain_modifications(60);
    let res = run_state(&s, 0);
    run_many_times_state_monad(n - 1, acc + res.state)
}

pub fn Test_StateMonadFFI_runStateMonadFFI(limit: i64) -> i64 {
    run_many_times_state_monad(limit, 0)
}
