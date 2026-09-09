// Include after the extracted generated Tree/Color and insertion functions.
// Uses native std::rc::Rc; contains no timing or allocation instrumentation.
fn b13_keys(tree: &Tree, out: &mut Vec<i64>) {
    if let Tree::T(_, left, key, right) = tree {
        b13_keys(left, out);
        out.push(*key);
        b13_keys(right, out);
    }
}

fn b13_shape(tree: &Tree) -> String {
    match tree {
        Tree::E => "E".into(),
        Tree::T(color, left, key, right) => format!("({} {} {} {})",
            if matches!(color, Color::R) { "R" } else { "B" },
            b13_shape(left), key, b13_shape(right)),
    }
}

fn b13_validate(tree: &Tree, low: Option<i64>, high: Option<i64>) -> (usize, usize) {
    match tree {
        Tree::E => (0, 1),
        Tree::T(color, left, key, right) => {
            assert!(low.is_none_or(|bound| bound < *key), "BST lower bound");
            assert!(high.is_none_or(|bound| *key < bound), "BST upper bound");
            if matches!(color, Color::R) {
                assert!(!matches!(left.as_ref(), Tree::T(Color::R, ..)), "red left child");
                assert!(!matches!(right.as_ref(), Tree::T(Color::R, ..)), "red right child");
            }
            let (ln, lh) = b13_validate(left, low, Some(*key));
            let (rn, rh) = b13_validate(right, Some(*key), high);
            assert_eq!(lh, rh, "black height at key {key}");
            (1 + ln + rn, lh + usize::from(matches!(color, Color::B)))
        }
    }
}

fn b13_check(tree: &Tree, expected: &std::collections::BTreeSet<i64>, black_root: bool) {
    if black_root {
        assert!(!matches!(tree, Tree::T(Color::R, ..)), "nonblack root");
    }
    assert_eq!(b13_validate(tree, None, None).0, expected.len());
    let mut actual = Vec::new();
    b13_keys(tree, &mut actual);
    assert_eq!(actual, expected.iter().copied().collect::<Vec<_>>());
}

fn b13_collect_weaks(tree: &std::rc::Rc<Tree>, out: &mut Vec<std::rc::Weak<Tree>>) {
    out.push(std::rc::Rc::downgrade(tree));
    if let Tree::T(_, left, _, right) = tree.as_ref() {
        b13_collect_weaks(left, out);
        b13_collect_weaks(right, out);
    }
}

fn b13_black_leaf(key: i64) -> std::rc::Rc<Tree> {
    std::rc::Rc::new(Tree::T(Color::B,
        std::rc::Rc::new(Tree::E), key, std::rc::Rc::new(Tree::E)))
}

// Neither identity nor survival of a weak-only allocation is imposed: both may
// change under legal copy-on-write implementations. Any surviving weak must still
// observe its old value, and all weak handles must expire when all owners drop.
fn b13_red_parent_cases() {
    let mut cases = 0;
    let mut unique_cases = 0;
    let mut unique_same_address = 0;
    for key in [3, 7, 13, 17] {
        for parent_shared in [false, true] {
            for parent_weak in [false, true] {
                for left_shared in [false, true] {
                    for right_shared in [false, true] {
                        for children_weak in [false, true] {
                            let left = b13_black_leaf(5);
                            let right = b13_black_leaf(15);
                            let left_old = left_shared.then(|| left.clone());
                            let right_old = right_shared.then(|| right.clone());
                            let left_shape = b13_shape(&left);
                            let right_shape = b13_shape(&right);
                            let left_weak = children_weak.then(|| std::rc::Rc::downgrade(&left));
                            let right_weak = children_weak.then(|| std::rc::Rc::downgrade(&right));
                            let root = std::rc::Rc::new(Tree::T(Color::R, left, 10, right));
                            let root_shape = b13_shape(&root);
                            let old = parent_shared.then(|| root.clone());
                            let weak = parent_weak.then(|| std::rc::Rc::downgrade(&root));
                            let address = std::rc::Rc::as_ptr(&root);
                            let root = Test_RBTree_ins(key, root);
                            let expected = std::collections::BTreeSet::from([5, 10, 15, key]);
                            b13_check(&root, &expected, false);
                            assert!(matches!(root.as_ref(), Tree::T(Color::R, _, 10, _)));
                            if !parent_shared && !parent_weak {
                                unique_cases += 1;
                                unique_same_address += usize::from(address == std::rc::Rc::as_ptr(&root));
                            }
                            if let Some(ref old) = old { assert_eq!(b13_shape(old), root_shape); }
                            if let Some(ref old) = left_old { assert_eq!(b13_shape(old), left_shape); }
                            if let Some(ref old) = right_old { assert_eq!(b13_shape(old), right_shape); }
                            for (weak, shape) in [(&weak, &root_shape), (&left_weak, &left_shape), (&right_weak, &right_shape)] {
                                if let Some(node) = weak.as_ref().and_then(|w| w.upgrade()) {
                                    assert_eq!(&b13_shape(&node), shape, "weak observed mutation");
                                }
                            }
                            let mut all_weak = Vec::new();
                            b13_collect_weaks(&root, &mut all_weak);
                            if let Some(ref node) = old { b13_collect_weaks(node, &mut all_weak); }
                            if let Some(ref node) = left_old { b13_collect_weaks(node, &mut all_weak); }
                            if let Some(ref node) = right_old { b13_collect_weaks(node, &mut all_weak); }
                            drop(root); drop(old); drop(left_old); drop(right_old);
                            assert!(all_weak.iter().all(|w| w.upgrade().is_none()), "node retained after final drop");
                            for weak in [&weak, &left_weak, &right_weak] {
                                assert!(weak.as_ref().is_none_or(|w| w.upgrade().is_none()));
                            }
                            cases += 1;
                        }
                    }
                }
            }
        }
    }
    println!("B13 red-parent cases: {cases}; unique parent address retained: {unique_same_address}/{unique_cases} (observation only)");
}

pub fn extra_checks() {
    b13_red_parent_cases();
    // Retain exactly 200 previous versions; compare each against BTreeSet after
    // every operation. Include the i64 boundaries without sentinel assumptions.
    let mut tree = std::rc::Rc::new(Tree::E);
    let mut expected = std::collections::BTreeSet::new();
    let mut versions = Vec::new();
    for n in 0..200 {
        versions.push((tree.clone(), expected.clone(), b13_shape(&tree)));
        let key = match n { 0 => i64::MIN, 1 => i64::MAX, _ => (n * 73 % 197) - 100 };
        tree = Test_RBTree_insert(key, tree);
        expected.insert(key);
        b13_check(&tree, &expected, true);
        for (old, old_expected, shape) in &versions {
            b13_check(old, old_expected, true);
            assert_eq!(&b13_shape(old), shape);
        }
    }
    let mut all_weak = Vec::new();
    b13_collect_weaks(&tree, &mut all_weak);
    for (old, _, _) in &versions { b13_collect_weaks(old, &mut all_weak); }
    drop(tree); drop(versions);
    assert!(all_weak.iter().all(|w| w.upgrade().is_none()));

    // Predominantly unique roots, periodically retained historical versions,
    // independent child references, root weak handles, and duplicate insertions.
    let mut tree = std::rc::Rc::new(Tree::E);
    let mut expected = std::collections::BTreeSet::new();
    let mut versions = Vec::new();
    let mut children = Vec::new();
    let mut state = 0xB13_u64;
    for n in 0..512 {
        if n % 17 == 0 { versions.push((tree.clone(), expected.clone(), b13_shape(&tree))); }
        if n % 7 == 0 {
            if let Tree::T(_, left, _, right) = tree.as_ref() {
                let child = if n % 2 == 0 { left } else { right };
                children.push((child.clone(), b13_shape(child)));
            }
        }
        let weak = (n % 23 == 0).then(|| (std::rc::Rc::downgrade(&tree), b13_shape(&tree)));
        state = state.wrapping_mul(6364136223846793005).wrapping_add(1);
        let key = ((state >> 32) % 173) as i64 - 86;
        tree = Test_RBTree_insert(key, tree);
        expected.insert(key);
        b13_check(&tree, &expected, true);
        for (old, old_expected, shape) in &versions {
            b13_check(old, old_expected, true);
            assert_eq!(&b13_shape(old), shape);
        }
        for (child, shape) in &children { assert_eq!(&b13_shape(child), shape); }
        if let Some((weak, shape)) = weak {
            if let Some(old) = weak.upgrade() { assert_eq!(b13_shape(&old), shape); }
        }
    }
    let mut all_weak = Vec::new();
    b13_collect_weaks(&tree, &mut all_weak);
    for (old, _, _) in &versions { b13_collect_weaks(old, &mut all_weak); }
    for (child, _) in &children { b13_collect_weaks(child, &mut all_weak); }
    drop(tree); drop(versions); drop(children);
    assert!(all_weak.iter().all(|w| w.upgrade().is_none()));
    println!("B13 extra checks passed: BTreeSet, RB invariants, 200 persistent versions, 512 mixed-sharing insertions, independent children, weak values and final lifetimes");
}
