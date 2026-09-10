fn depth_checks() {
    assert_eq!(Test_RBTree_depth(std::rc::Rc::new(Tree::E)), 0);
    // The public owned boundary must release its owner before it returns.
    let tree = Test_RBTree_buildTree(100000, std::rc::Rc::new(Tree::E));
    let weak = std::rc::Rc::downgrade(&tree);
    assert_eq!(Test_RBTree_depth(tree), 22);
    assert!(weak.upgrade().is_none(), "Unique owner is destroyed inside depth");

    let tree = Test_RBTree_buildTree(1000, std::rc::Rc::new(Tree::E));
    let before = b13_shape(&tree);
    let depth = Test_RBTree_depth(tree.clone());
    let weak = std::rc::Rc::downgrade(&tree);
    for _ in 0..20 {
        assert_eq!(Test_RBTree_depth(tree.clone()), depth);
        assert_eq!(b13_shape(&tree), before);
        assert_eq!(std::rc::Rc::strong_count(&tree), 1, "A read must not retain its root");
        assert_eq!(b13_shape(&weak.upgrade().unwrap()), before);
    }
    let changed = Test_RBTree_insert(1001, tree.clone());
    assert_eq!(b13_shape(&tree), before);
    assert!(Test_RBTree_depth(changed) >= depth);
    drop(tree);
    assert!(weak.upgrade().is_none());

    // Repeated edges and an independently retained child are legal read-only
    // inputs even when the artificial shape is not an ordered red-black tree.
    for root_shared in [false, true] {
        for child_shared in [false, true] {
            let child = b13_black_leaf(7);
            let saved_child = child_shared.then(|| child.clone());
            let child_weak = std::rc::Rc::downgrade(&child);
            let root = std::rc::Rc::new(Tree::T(Color::B, child.clone(), 10, child));
            let saved_root = root_shared.then(|| root.clone());
            let root_weak = std::rc::Rc::downgrade(&root);
            assert_eq!(Test_RBTree_depth(root), 2);
            assert_eq!(root_weak.upgrade().is_some(), root_shared);
            assert_eq!(child_weak.upgrade().is_some(), root_shared || child_shared);
            if let Some(ref root) = saved_root { assert_eq!(Test_RBTree_depth(root.clone()), 2); }
            if let Some(ref child) = saved_child { assert_eq!(Test_RBTree_depth(child.clone()), 1); }
            drop(saved_root); drop(saved_child);
            assert!(root_weak.upgrade().is_none());
            assert!(child_weak.upgrade().is_none());
        }
    }
    println!("Depth checks passed: 100k tree, owned destruction before return, repeated reads, later insertion, shared edges and independent child/Weak lifetimes");
}
