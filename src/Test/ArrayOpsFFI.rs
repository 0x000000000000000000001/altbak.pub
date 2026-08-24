pub fn Test_ArrayOpsFFI_runArrayOpsFFI(mut limit: i64) -> i64 {
    let mut arr = Vec::new();
    for i in 0..limit {
        arr.push(i);
    }
    let mut arr2 = Vec::new();
    for x in &arr {
        arr2.push(*x * 2);
    }
    let mut sum = 0;
    for x in &arr2 {
        sum += *x;
    }
    sum
}