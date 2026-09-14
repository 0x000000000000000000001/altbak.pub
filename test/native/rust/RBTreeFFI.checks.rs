pub fn check_structure() {
    fn insert(x: i64, tree: &Tree) -> Tree {
        let n = ins(x, tree);
        Some(node(1, n.left.clone(), n.value, n.right.clone()))
    }
    fn inspect(tree: &Tree, low: i64, high: i64) -> (i64, usize) {
        let Some(n) = tree else { return (1, 0); };
        assert!(low < n.value && n.value < high);
        if n.color == 0 {
            assert!(n.left.as_ref().map_or(true, |v| v.color == 1));
            assert!(n.right.as_ref().map_or(true, |v| v.color == 1));
        }
        let (lh, lc) = inspect(&n.left, low, n.value);
        let (rh, rc) = inspect(&n.right, n.value, high);
        assert_eq!(lh, rh);
        (lh + i64::from(n.color == 1), lc+rc+1)
    }
    for order in [ (1..=100).collect::<Vec<_>>(), (1..=100).rev().collect(),
                   (1..=50).flat_map(|x| [x,101-x]).collect() ] {
        let mut tree = None;
        let mut versions = Vec::new();
        for (index, key) in order.into_iter().enumerate() {
            tree = insert(key, &tree);
            assert_eq!(tree.as_ref().unwrap().color, 1);
            assert_eq!(inspect(&tree, i64::MIN, i64::MAX).1, index+1);
            versions.push(tree.clone());
        }
        for (index, version) in versions.iter().enumerate() {
            assert_eq!(inspect(version, i64::MIN, i64::MAX).1, index+1);
        }
        drop(versions);
        let old = tree.as_ref().unwrap();
        let weak = Rc::downgrade(old);
        let duplicate = ins(old.value, &tree);
        assert!(match (&old.left, &duplicate.left) {
            (Some(a), Some(b)) => Rc::ptr_eq(a,b), (None,None)=>true, _=>false
        });
        assert!(match (&old.right, &duplicate.right) {
            (Some(a), Some(b)) => Rc::ptr_eq(a,b), (None,None)=>true, _=>false
        });
        let newer = insert(101, &tree);
        assert_eq!(inspect(&tree, i64::MIN, i64::MAX).1,100);
        assert_eq!(inspect(&newer, i64::MIN, i64::MAX).1,101);
        assert!(weak.upgrade().is_some());
        drop(tree);
        assert!(weak.upgrade().is_none(), "old root must be released");
    }
}
