pub fn Test_PrimesFFI_runPrimesFFI(mut limit: i64) -> i64 {
    if limit < 2 { return 0; }
    let mut sieve = vec![true; (limit + 1) as usize];
    let mut sum = 0;
    for p in 2..=(limit as usize) {
        if sieve[p] {
            let mut i = p * p;
            while i <= limit as usize {
                sieve[i] = false;
                i += p;
            }
        }
    }
    for p in 2..=(limit as usize) {
        if sieve[p] { sum += p as i64; }
    }
    sum
}
