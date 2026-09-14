pub fn check_structure() {
    use std::cell::Cell;
    for operation in 0..3 {
        let applications = Arc::new(Cell::new(0));
        let counter = applications.clone();
        let number: Church = Arc::new(move |f| {
            counter.set(counter.get() + 1);
            f
        });
        let result = match operation {
            0 => succ(number),
            1 => add_c(number.clone(), number),
            _ => mul_c(number.clone(), number),
        };
        let applied = result(Arc::new(|x| x+1));
        assert_eq!(applications.get(), 0, "numeral applied before final argument");
        assert_eq!(applied(2), if operation == 2 {3} else {4});
        assert_eq!(applications.get(), if operation == 0 {1} else {2});
        applied(3);
        assert_eq!(applications.get(), if operation == 0 {2} else {4});
    }
}
