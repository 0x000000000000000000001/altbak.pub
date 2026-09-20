pub fn check_structure() {
    fn words(values: &[&str]) -> List<String> {
        values.iter().rev().fold(List::Nil, |tail, word|
            List::Cons(String::from(*word), Box::new(tail)))
    }
    fn collect(mut list: List<String>) -> Vec<String> {
        let mut values = Vec::new();
        while let List::Cons(value, tail) = list {
            values.push(value);
            list = *tail;
        }
        values
    }
    assert_eq!(collect(filter(&|word: &String| word.len() == 1,
        words(&["a", "bb", "c"]))), ["a", "c"]);
    assert_eq!(collect(reverse(words(&["a", "bb", "c"]))), ["c", "bb", "a"]);
    assert!(collect(filter(&|_: &String| false, words(&["a"]))).is_empty());
}
