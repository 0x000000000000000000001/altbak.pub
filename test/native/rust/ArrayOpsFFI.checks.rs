pub fn check_structure() {
    let values = vec![String::from("a"), String::from("bb"), String::from("c")];
    let mut selected = filter(|word| word.len() == 1, &values);
    assert_eq!(selected, ["a", "c"]);
    assert_eq!(selected.iter().fold(4_usize, |count, word| count + word.len()), 6);
    selected[0].push_str(" changed");
    assert_eq!(values, ["a", "bb", "c"], "filter must preserve the input array");
    assert!(filter(|_: &String| false, &values).is_empty());
}
