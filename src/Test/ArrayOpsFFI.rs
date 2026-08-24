pub fn Test_ArrayOpsFFI_runArrayOpsFFI(mut limit: i64) -> i64 {
    let mut arr = Vec::new();
    for i in 1..=limit {
        arr.push(i);
    }
    let mut evens = Vec::new();
    for x in arr.iter() {
        if *x % 2 == 0 { evens.push(*x); }
    }
    let mut sum = 0;
    for x in evens.iter() {
        sum += *x;
    }
    sum
}
