pub fn check_structure() {
    let calls = std::cell::Cell::new(0);
    let thunk = defer(|| { calls.set(calls.get() + 1); String::from("value") });
    assert_eq!(calls.get(), 0, "defer must not execute its thunk");
    assert_eq!(force(&thunk), "value");
    assert_eq!(force(&thunk), "value");
    assert_eq!(calls.get(), 2, "force must not memoize the thunk");
}
