use std::collections::BTreeSet;

pub fn Test_RBTreeFFICheatcode_runRBTreeFFICheatcode(mut limit: i64) -> i64 {
    let mut set = BTreeSet::new();
    for i in 1..=limit {
        set.insert(i);
    }
    let mut count = 0;
    for i in 1..=limit {
        if set.contains(&i) {
            count += 1;
        }
    }
    count
}