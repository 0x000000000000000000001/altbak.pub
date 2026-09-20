pub fn check_structure() {
    let action: StateFn<String, usize> = bind_state(get(), Box::new(|value: String| {
        let length = value.len();
        bind_state(put(value + "!"), Box::new(move |_| pure_state(length)))
    }));
    let result = run_state(&action, String::from("ab"));
    assert_eq!((result.val, result.state), (2, String::from("ab!")));
    let again = run_state(&action, String::from("z"));
    assert_eq!((again.val, again.state), (1, String::from("z!")));
    let modified = run_state(&modify(Box::new(|value: String| value + "?")), String::from("a"));
    assert_eq!((modified.val, modified.state), ((), String::from("a?")));
}
