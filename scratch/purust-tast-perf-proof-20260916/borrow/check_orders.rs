fn validate_order(tree: &Tree, lo: i64, hi: i64, out: &mut Vec<i64>) -> (usize, usize) {
    match tree {
        Tree::E => (1, 0),
        Tree::T(color, left, key, right) => {
            assert!(lo < *key && *key < hi);
            if matches!(color, Color::R) {
                for child in [left, right] {
                    if let Tree::T(c, ..) = child.as_ref() { assert!(!matches!(c, Color::R)); }
                }
            }
            let (lh, ld) = validate_order(left, lo, *key, out);
            out.push(*key);
            let (rh, rd) = validate_order(right, *key, hi, out);
            assert_eq!(lh, rh);
            (lh + usize::from(matches!(color, Color::B)), 1 + ld.max(rd))
        }
    }
}
fn main() {
    let n = 100000_i64;
    let ascending: Vec<_> = (1..=n).collect();
    let mut shuffled = ascending.clone();
    let mut state = 0x6a09e667f3bcc909_u64;
    for i in (1..shuffled.len()).rev() {
        state ^= state << 13; state ^= state >> 7; state ^= state << 17;
        shuffled.swap(i, state as usize % (i + 1));
    }
    let duplicates: Vec<_> = (0..n).map(|k| (k * 73) % 997 + 1).collect();
    for (name, order) in [("ascending", ascending), ("shuffled", shuffled), ("duplicates", duplicates)] {
        let mut tree = std::rc::Rc::new(Tree::E);
        for key in &order { tree = Test_RBTree_insert(*key, tree); }
        if let Tree::T(color, ..) = tree.as_ref() { assert!(matches!(color, Color::B)); }
        let mut got = Vec::new();
        let (black_height, independent_depth) = validate_order(&tree, i64::MIN, i64::MAX, &mut got);
        let mut expected = order; expected.sort(); expected.dedup();
        assert_eq!(got, expected);
        let depth = Test_RBTree_depth(tree);
        assert_eq!(depth as usize, independent_depth);
        let hash = got.iter().fold(0xcbf29ce484222325_u64, |h, &k| (h ^ k as u64).wrapping_mul(0x100000001b3));
        println!("{name}\t{}\t{black_height}\t{depth}\t{hash}", got.len());
    }
}
