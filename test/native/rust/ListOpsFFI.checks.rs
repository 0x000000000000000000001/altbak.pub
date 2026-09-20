pub fn check_structure() {
    let input = range_list(1, 8);
    let output = filter_evens(&input);
    assert_eq!(foldl(|a,b| a*10+b, 0, &input), 12345678);
    assert_eq!(foldl(|a,b| a*10+b, 0, &output), 8642);
    let words = List::Cons(String::from("a"), Box::new(List::Cons(
        String::from("bc"), Box::new(List::Nil))));
    assert_eq!(foldl(|count, word| count + word.len(), 0_usize, &words), 3);
    assert_eq!(foldl(|text, value| text + &value.to_string(), String::new(), &output), "8642");
}
