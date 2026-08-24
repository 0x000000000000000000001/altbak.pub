pub fn Test_PrimesFFICheatcode_runPrimesFFICheatcode(mut limit: i64) -> i64 {
    if limit < 2 { return 0; }
    let mut sieve = vec![true; (limit + 1) as usize];
    let mut count = 0;
    for p in 2..=(limit as usize) {
        if sieve[p] {
            count += 1;
            let mut i = p * p;
            while i <= limit as usize {
                sieve[i] = false;
                i += p;
            }
        }
    }
    count
}