
pub fn Test_TCOFFI_runTCOFFI(mut limit: i64) -> i64 {
    fn sum_tco(n: i64, acc: i64) -> i64 {
        let mut curr_n = n;
        let mut curr_acc = acc;
        loop {
            if curr_n == 0 {
                return curr_acc;
            }
            curr_acc += curr_n;
            curr_n -= 1;
        }
    }
    sum_tco(limit, 0)
}