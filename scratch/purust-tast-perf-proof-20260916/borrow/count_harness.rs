impl tracked::Kind for Tree {
    fn kind(&self) -> &'static str { match self { Tree::E => "empty", Tree::T(..) => "node" } }
}
fn validate(tree: &Tree, low: i64, high: i64) -> (usize, usize) {
    match tree {
        Tree::E => (0, 1),
        Tree::T(color, left, key, right) => {
            assert!(low < *key && *key < high);
            if matches!(color, Color::R) {
                for child in [left, right] {
                    if let Tree::T(c, ..) = child.as_ref() { assert!(!matches!(c, Color::R)); }
                }
            }
            let (nl, hl) = validate(left, low, *key);
            let (nr, hr) = validate(right, *key, high);
            assert_eq!(hl, hr);
            (1 + nl + nr, hl + usize::from(matches!(color, Color::B)))
        }
    }
}
fn check(tree: &Tree, count: usize) {
    if let Tree::T(color, ..) = tree { assert!(matches!(color, Color::B)); }
    assert_eq!(validate(tree, i64::MIN, i64::MAX).0, count);
}
fn keys(tree: &Tree) -> Vec<i64> {
    match tree {
        Tree::E => Vec::new(),
        Tree::T(_, left, key, right) => {
            let mut result = keys(left); result.push(*key); result.extend(keys(right)); result
        }
    }
}
fn main() {
    assert_eq!(std::mem::size_of::<tracked::Rc<Tree>>(), std::mem::size_of::<std::rc::Rc<Tree>>());
    assert_eq!(std::mem::size_of::<Tree>(), 32);
    for order in [[3,2,1], [3,1,2], [1,3,2], [1,2,3]] {
        let mut tree = tracked::Rc::new(Tree::E);
        for (i, key) in order.into_iter().enumerate() {
            tree = Test_RBTree_insert(key, tree); check(&tree, i+1);
        }
    }
    // Verify get_mut's weak-reference exclusion and unwrap's distinct rule.
    let mut tree = tracked::Rc::new(Tree::T(Color::B,
        tracked::Rc::new(Tree::E), 1, tracked::Rc::new(Tree::E)));
    let weak = tracked::Rc::downgrade(&tree);
    assert!(tracked::Rc::get_mut(&mut tree).is_none());
    tree = Test_RBTree_insert(2, tree); check(&tree, 2);
    assert!(weak.upgrade().is_none()); drop(tree); drop(weak);

    tracked::phase("unique_build");
    let tree = Test_RBTree_buildTree(std::hint::black_box(100000), tracked::Rc::new(Tree::E));
    check(&tree, 100000);
    tracked::phase("unique_depth_and_drop");
    assert_eq!(Test_RBTree_depth(tree), 22);

    tracked::phase("persistent_build");
    let mut tree = tracked::Rc::new(Tree::E);
    let mut versions = Vec::new();
    let mut expected = Vec::new();
    // A permutation covers both descent directions and all retained versions.
    for n in 0..200 {
        versions.push((tree.clone(), expected.clone()));
        let key = (n * 73) % 200 + 1;
        tree = Test_RBTree_insert(key, tree);
        expected.push(key); expected.sort(); check(&tree, expected.len());
        assert_eq!(keys(&tree), expected);
        for (old, keys_before) in &versions { check(old, keys_before.len()); assert_eq!(&keys(old), keys_before); }
    }
    tracked::phase("persistent_depth");
    assert!(Test_RBTree_depth(tree.clone()) > 0);
    for (old, expected) in &versions { assert_eq!(&keys(old), expected); }
    tracked::phase("persistent_drop");
    drop(tree); drop(versions);
    tracked::phase("done");
    tracked::print_events();
}
