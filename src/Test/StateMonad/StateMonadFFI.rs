struct StateResult<S, A> {
    val: A,
    state: S,
}

type StateFn<S, A> = Box<dyn Fn(S) -> StateResult<S, A>>;

fn run_state<S, A>(s: &StateFn<S, A>, init: S) -> StateResult<S, A> {
    s(init)
}

fn bind_state<S: 'static, A: 'static, B: 'static>(
    s: StateFn<S, A>, g: Box<dyn Fn(A) -> StateFn<S, B>>,
) -> StateFn<S, B> {
    Box::new(move |state: S| {
        let r1 = s(state);
        let g_prime = g(r1.val);
        g_prime(r1.state)
    })
}

fn pure_state<S: 'static, A: Clone + 'static>(a: A) -> StateFn<S, A> {
    Box::new(move |s| StateResult { val: a.clone(), state: s })
}

fn get<S: Clone + 'static>() -> StateFn<S, S> {
    Box::new(move |s: S| StateResult { val: s.clone(), state: s })
}

fn put<S: Clone + 'static>(s: S) -> StateFn<S, ()> {
    Box::new(move |_| StateResult { val: (), state: s.clone() })
}

fn modify<S: Clone + 'static>(f: Box<dyn Fn(S) -> S>) -> StateFn<S, ()> {
    bind_state(get(), Box::new(move |s| put(f(s))))
}

fn chain_modifications(n: i64) -> StateFn<i64, ()> {
    if n == 0 {
        return pure_state(());
    }
    bind_state(
        modify(Box::new(|x| x + 1)),
        Box::new(move |_| chain_modifications(n - 1)),
    )
}

fn run_many_times_state_monad(n: i64, depth: i64, acc: i64) -> i64 {
    if n == 0 {
        return acc;
    }
    let s = chain_modifications(depth);
    let res = run_state(&s, 0);
    run_many_times_state_monad(n - 1, depth, acc + res.state)
}

pub fn Test_StateMonadFFI_runStateMonadFFI(limit: i64) -> i64 {
    run_many_times_state_monad(20, limit, 0)
}
