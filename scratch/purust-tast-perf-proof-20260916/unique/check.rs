#![allow(warnings)]
#![recursion_limit="512"]
include!("kernel_unique.rs");

fn validate(tree: &Tree, low: i64, high: i64, actual: &mut Vec<i64>) -> (usize, usize, i64) {
    match tree {
        Tree::E => (0, 1, 0),
        Tree::T(color, left, key, right) => {
            assert!(low < *key && *key < high);
            if matches!(color, Color::R) {
                for child in [left, right] {
                    if let Tree::T(c, ..) = child.as_ref() { assert!(!matches!(c, Color::R)); }
                }
            }
            let (nl, hl, dl) = validate(left, low, *key, actual);
            actual.push(*key);
            let (nr, hr, dr) = validate(right, *key, high, actual);
            assert_eq!(hl, hr, "black height mismatch at key {key}");
            assert_eq!(std::rc::Rc::strong_count(left), if matches!(left.as_ref(), Tree::E) { std::rc::Rc::strong_count(left) } else { 1 });
            assert_eq!(std::rc::Rc::strong_count(right), if matches!(right.as_ref(), Tree::E) { std::rc::Rc::strong_count(right) } else { 1 });
            (1 + nl + nr, hl + usize::from(matches!(color, Color::B)), 1 + dl.max(dr))
        }
    }
}
fn check(tree: &Tree, expected: &[i64]) -> i64 {
    if let Tree::T(color, ..) = tree { assert!(matches!(color, Color::B)); }
    let mut keys = Vec::new();
    let (count, _, depth) = validate(tree, i64::MIN, i64::MAX, &mut keys);
    assert_eq!(count, expected.len());
    assert_eq!(keys, expected);
    depth
}
fn test(label: &str, order: Vec<i64>) {
    let mut expected = order.clone(); expected.sort_unstable(); expected.dedup();
    let mut tree = std::rc::Rc::new(Tree::E);
    UNIQUE_CALLS.store(0, std::sync::atomic::Ordering::Relaxed);
    EMPTY_CHECKS.store(0, std::sync::atomic::Ordering::Relaxed);
    for key in order {
        tree = Test_RBTree_insert(key, tree);
    }
    let expected_depth = check(&tree, &expected);
    assert_eq!(Test_RBTree_depth(tree), expected_depth);
    println!("{label}\tkeys={}\tdepth={expected_depth}\tmutation_calls={}\tempty_dynamic_checks={}", expected.len(),
             UNIQUE_CALLS.load(std::sync::atomic::Ordering::Relaxed), EMPTY_CHECKS.load(std::sync::atomic::Ordering::Relaxed));
}
fn main() {
    for order in [[3,2,1], [3,1,2], [1,3,2], [1,2,3]] { test("rotation", order.to_vec()); }
    test("descending_100k", (1..=100000).rev().collect());
    test("ascending_100k", (1..=100000).collect());
    let mut shuffled: Vec<i64> = (1..=100000).collect();
    let mut rng = 0x71e9af148634u64;
    for i in (1..shuffled.len()).rev() {
        rng ^= rng << 13; rng ^= rng >> 7; rng ^= rng << 17;
        shuffled.swap(i, (rng as usize) % (i+1));
    }
    test("shuffled_100k", shuffled);
    test("duplicates_100k", (0..100000).map(|n| (n * 73) % 4096 + 1).collect());
    let tree = Test_RBTree_buildTree(100000, std::rc::Rc::new(Tree::E));
    let expected: Vec<i64> = (1..=100000).collect();
    assert_eq!(check(&tree, &expected), 22);
    assert_eq!(Test_RBTree_depth(tree), 22);
    println!("original_buildTree_100k\tok");
}
