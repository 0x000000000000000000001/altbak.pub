fn range(start: i64, end: i64) -> Vec<i64> {
    if start <= end {
        (start..=end).collect()
    } else {
        (end..=start).rev().collect()
    }
}

fn filter<A: Clone>(predicate: impl Fn(&A) -> bool, array: &[A]) -> Vec<A> {
    array.iter().filter(|value| predicate(value)).cloned().collect()
}

pub fn Test_ArrayOpsFFI_runArrayOpsFFI(limit: i64) -> i64 {
    let array = range(1, limit);
    let evens = filter(|value| value % 2 == 0, &array);
    evens.iter().fold(0, |acc, value| acc + value)
}
