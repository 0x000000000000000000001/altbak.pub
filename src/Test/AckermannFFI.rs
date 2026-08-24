pub fn Test_AckermannFFI_runAckermannFFI(limit: i64) -> i64 {
    fn ack(m: i64, n: i64) -> i64 {
        if m == 0 {
            n + 1
        } else if m > 0 && n == 0 {
            ack(m - 1, 1)
        } else {
            ack(m - 1, ack(m, n - 1))
        }
    }
    ack(3, 4)
}
