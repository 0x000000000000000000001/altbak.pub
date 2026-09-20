pub fn check_structure() {
    use std::cell::Cell;
    for operation in 0..3 {
        let applications = Rc::new(Cell::new(0));
        let counter = applications.clone();
        let number: Church<i64> = Rc::new(move |f| {
            counter.set(counter.get() + 1);
            f
        });
        let result = match operation {
            0 => succ(number),
            1 => add_c(number.clone(), number),
            _ => mul_c(number.clone(), number),
        };
        let applied = result(Rc::new(|x| x+1));
        assert_eq!(applications.get(), 0, "numeral applied before final argument");
        assert_eq!(applied(2), if operation == 2 {3} else {4});
        assert_eq!(applications.get(), if operation == 0 {1} else {2});
        applied(3);
        assert_eq!(applications.get(), if operation == 0 {2} else {4});
    }
    let two = succ(succ(zero::<String>()));
    let append: Step<String> = Rc::new(|value| value + "x");
    let four = mul_c(two.clone(), two);
    assert_eq!(four(append)(String::from("!")), "!xxxx");
}
