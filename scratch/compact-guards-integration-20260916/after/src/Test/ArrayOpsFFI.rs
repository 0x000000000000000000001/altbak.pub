pub fn Test_ArrayOpsFFI_runArrayOpsFFI(mut limit: i64) -> i64 {
    let mut arr = Vec::new();
    let step = if limit >= 1 { 1 } else { -1 };
    let mut i = 1;
    loop {
        arr.push(i);
        if i == limit {
            break;
        }
        i += step;
    }
    let mut evens = Vec::new();
    for x in arr.iter() {
        if *x % 2 == 0 {
            evens.push(*x);
        }
    }
    let mut sum = 0;
    for x in evens.iter() {
        sum += *x;
    }
    sum
}
