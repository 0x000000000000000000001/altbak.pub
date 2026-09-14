pub fn check_structure() {
    let input = range_list(1, 8);
    let output = filter_evens(&input);
    assert_eq!(foldl(|a,b| a*10+b, 0, &input), 12345678);
    assert_eq!(foldl(|a,b| a*10+b, 0, &output), 8642);
}
