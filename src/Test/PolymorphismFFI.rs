trait MyEq {
    fn my_eq(&self, other: &Self) -> bool;
}
impl MyEq for i64 {
    fn my_eq(&self, other: &Self) -> bool { self == other }
}
impl<T: MyEq> MyEq for Vec<T> {
    fn my_eq(&self, other: &Self) -> bool {
        if self.len() != other.len() { return false; }
        for i in 0..self.len() {
            if !self[i].my_eq(&other[i]) { return false; }
        }
        true
    }
}
pub fn Test_PolymorphismFFI_runPolymorphismFFI(mut limit: i64) -> i64 {
    let mut a = vec![vec![1, 2], vec![3, 4]];
    let mut b = vec![vec![1, 2], vec![3, 4]];
    let mut count = 0;
    for _ in 0..limit {
        if a.my_eq(&b) { count += 1; }
    }
    count
}