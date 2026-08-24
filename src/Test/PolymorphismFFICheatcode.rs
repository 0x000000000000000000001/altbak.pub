trait Monoidish {
    fn mempty_(&self) -> i64;
    fn mappend_(&self, x: i64, y: i64) -> i64;
}
struct IntMonoidish;
impl Monoidish for IntMonoidish {
    fn mempty_(&self) -> i64 { 1 }
    fn mappend_(&self, x: i64, y: i64) -> i64 { x + y }
}
pub fn Test_PolymorphismFFICheatcode_runPolymorphismFFICheatcode(mut limit: i64) -> i64 {
    let mut acc = 0;
    let m = IntMonoidish;
    while limit > 0 {
        acc = m.mappend_(acc, m.mempty_());
        limit -= 1;
    }
    acc
}
