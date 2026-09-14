// Include in the module that defines Tree, Color and probe_rotate_ll.
// Entry point: validation_extra(). This file intentionally has no main().

struct ValidationLlFixture {
    root: std::rc::Rc<Tree>,
    root_address: usize,
    parent_address: usize,
    grandchild_address: usize,
    boundary_addresses: [usize; 4],
    boundary_watches: [std::rc::Weak<Tree>; 4],
    boundary_handles: Vec<std::rc::Rc<Tree>>,
}

fn validation_leaf(key: i64) -> std::rc::Rc<Tree> {
    std::rc::Rc::new(Tree::T(
        Color::B,
        std::rc::Rc::new(Tree::E),
        key,
        std::rc::Rc::new(Tree::E),
    ))
}

fn validation_ll_fixture(shared_boundaries: bool) -> ValidationLlFixture {
    let boundaries = [
        validation_leaf(10),
        validation_leaf(30),
        validation_leaf(50),
        validation_leaf(70),
    ];
    let boundary_addresses = boundaries.each_ref().map(|p| std::rc::Rc::as_ptr(p) as usize);
    let boundary_watches = boundaries.each_ref().map(std::rc::Rc::downgrade);
    let boundary_handles = if shared_boundaries {
        boundaries.to_vec()
    } else {
        Vec::new()
    };
    let [a, b, c, d] = boundaries;
    let grandchild = std::rc::Rc::new(Tree::T(Color::R, a, 20, b));
    let grandchild_address = std::rc::Rc::as_ptr(&grandchild) as usize;
    let parent = std::rc::Rc::new(Tree::T(Color::R, grandchild, 40, c));
    let parent_address = std::rc::Rc::as_ptr(&parent) as usize;
    let root = std::rc::Rc::new(Tree::T(Color::B, parent, 60, d));
    let root_address = std::rc::Rc::as_ptr(&root) as usize;
    ValidationLlFixture {
        root,
        root_address,
        parent_address,
        grandchild_address,
        boundary_addresses,
        boundary_watches,
        boundary_handles,
    }
}

fn validation_left(tree: &Tree) -> &std::rc::Rc<Tree> {
    match tree {
        Tree::T(_, left, _, _) => left,
        Tree::E => panic!("expected a nonempty validation tree"),
    }
}

fn validation_snapshot(tree: &Tree) -> String {
    match tree {
        Tree::E => "E".to_owned(),
        Tree::T(color, left, key, right) => {
            let color = match color {
                Color::B => "B",
                Color::R => "R",
            };
            format!(
                "{color}({:p}:{},{key},{:p}:{})",
                std::rc::Rc::as_ptr(left),
                validation_snapshot(left),
                std::rc::Rc::as_ptr(right),
                validation_snapshot(right),
            )
        }
    }
}

fn validation_assert_leaf(tree: &Tree, expected_key: i64) {
    match tree {
        Tree::T(Color::B, left, key, right) => {
            assert_eq!(*key, expected_key);
            assert!(matches!(left.as_ref(), Tree::E));
            assert!(matches!(right.as_ref(), Tree::E));
        }
        _ => panic!("rotation changed a boundary subtree"),
    }
}

fn validation_assert_rotated(fixture: &ValidationLlFixture, shared_boundaries: bool) {
    assert_eq!(std::rc::Rc::as_ptr(&fixture.root) as usize, fixture.root_address);
    let (left, right) = match fixture.root.as_ref() {
        Tree::T(Color::R, left, 40, right) => (left, right),
        _ => panic!("LL rotation must produce a red root with key 40"),
    };
    assert_eq!(std::rc::Rc::as_ptr(left) as usize, fixture.grandchild_address);
    assert_eq!(std::rc::Rc::as_ptr(right) as usize, fixture.parent_address);
    assert_eq!(std::rc::Rc::strong_count(left), 1);
    assert_eq!(std::rc::Rc::strong_count(right), 1);
    let (a, b) = match left.as_ref() {
        Tree::T(Color::B, a, 20, b) => (a, b),
        _ => panic!("old LL cell must become the black left child"),
    };
    let (c, d) = match right.as_ref() {
        Tree::T(Color::B, c, 60, d) => (c, d),
        _ => panic!("old L cell must become the black right child"),
    };
    for (index, boundary) in [a, b, c, d].into_iter().enumerate() {
        assert_eq!(std::rc::Rc::as_ptr(boundary) as usize, fixture.boundary_addresses[index]);
        assert_eq!(std::rc::Rc::strong_count(boundary), if shared_boundaries { 2 } else { 1 });
        validation_assert_leaf(boundary, 10 + 20 * index as i64);
    }
}

fn validation_check_released(fixture: ValidationLlFixture) {
    let ValidationLlFixture {
        root,
        boundary_handles,
        boundary_watches,
        ..
    } = fixture;
    drop(root);
    drop(boundary_handles);
    for watch in boundary_watches {
        assert!(watch.upgrade().is_none(), "rotation retained a boundary subtree after all owners dropped");
    }
}

fn validation_guard_matrix() -> usize {
    let mut checked = 0;
    // Guard values: 0 = unique, 1 = extra strong owner, 2 = weak owner.
    // Every combination covers overlapping guards as well as single guards.
    for shared_boundaries in [false, true] {
        for root_guard in 0..3 {
            for parent_guard in 0..3 {
                for grandchild_guard in 0..3 {
                    let mut fixture = validation_ll_fixture(shared_boundaries);
                    let mut strong_guards = Vec::new();
                    let mut weak_guards = Vec::new();
                    for (depth, guard) in [root_guard, parent_guard, grandchild_guard].into_iter().enumerate() {
                        let pointer = match depth {
                            0 => &fixture.root,
                            1 => validation_left(&fixture.root),
                            2 => validation_left(validation_left(&fixture.root)),
                            _ => unreachable!(),
                        };
                        match guard {
                            0 => (),
                            1 => strong_guards.push(std::rc::Rc::clone(pointer)),
                            2 => weak_guards.push(std::rc::Rc::downgrade(pointer)),
                            _ => unreachable!(),
                        }
                    }
                    let before = validation_snapshot(&fixture.root);
                    let strong_before: Vec<_> = strong_guards.iter().map(|p| validation_snapshot(p)).collect();
                    let weak_before: Vec<_> = weak_guards.iter().map(|p| validation_snapshot(&p.upgrade().unwrap())).collect();
                    let slot = std::rc::Rc::get_mut(&mut fixture.root);
                    assert_eq!(slot.is_some(), root_guard == 0, "Rc::get_mut must exclude shared or weak root owners");
                    let changed = match slot {
                        Some(slot) => probe_rotate_ll(slot),
                        None => false,
                    };
                    let expected = root_guard == 0 && parent_guard == 0 && grandchild_guard == 0;
                    assert_eq!(changed, expected, "guard matrix: root={root_guard}, L={parent_guard}, LL={grandchild_guard}, shared boundaries={shared_boundaries}");
                    if expected {
                        validation_assert_rotated(&fixture, shared_boundaries);
                        let after = validation_snapshot(&fixture.root);
                        assert!(!probe_rotate_ll(std::rc::Rc::get_mut(&mut fixture.root).unwrap()), "the completed rotation must not match a second time");
                        assert_eq!(validation_snapshot(&fixture.root), after);
                    } else {
                        assert_eq!(validation_snapshot(&fixture.root), before, "a rejected guard must leave all fields and edges unchanged");
                    }
                    for (guard, snapshot) in strong_guards.iter().zip(&strong_before) {
                        assert_eq!(validation_snapshot(guard), *snapshot, "rotation changed a persistent strong alias");
                    }
                    for (guard, snapshot) in weak_guards.iter().zip(&weak_before) {
                        assert_eq!(validation_snapshot(&guard.upgrade().unwrap()), *snapshot, "rotation invalidated or changed a weak alias");
                    }
                    drop(strong_guards);
                    drop(weak_guards);
                    validation_check_released(fixture);
                    checked += 1;
                }
            }
        }
    }
    checked
}

fn validation_missing_patterns() -> usize {
    let mut checked = 0;
    for absent in 0..6 {
        let mut fixture = validation_ll_fixture(false);
        let root = std::rc::Rc::get_mut(&mut fixture.root).unwrap();
        match absent {
            0 => *root = Tree::E,
            1 => match root {
                Tree::T(color, _, _, _) => *color = Color::R,
                _ => unreachable!(),
            },
            2 | 3 => match root {
                Tree::T(_, left, _, _) => {
                    let parent = std::rc::Rc::get_mut(left).unwrap();
                    match absent {
                        2 => *parent = Tree::E,
                        3 => match parent {
                            Tree::T(color, _, _, _) => *color = Color::B,
                            _ => unreachable!(),
                        },
                        _ => unreachable!(),
                    }
                }
                _ => unreachable!(),
            },
            4 | 5 => match root {
                Tree::T(_, left, _, _) => match std::rc::Rc::get_mut(left).unwrap() {
                    Tree::T(_, left_left, _, _) => {
                        let grandchild = std::rc::Rc::get_mut(left_left).unwrap();
                        match absent {
                            4 => *grandchild = Tree::E,
                            5 => match grandchild {
                                Tree::T(color, _, _, _) => *color = Color::B,
                                _ => unreachable!(),
                            },
                            _ => unreachable!(),
                        }
                    }
                    _ => unreachable!(),
                },
                _ => unreachable!(),
            },
            _ => unreachable!(),
        }
        let before = validation_snapshot(root);
        assert!(!probe_rotate_ll(root), "absent LL pattern {absent} must return false");
        assert_eq!(validation_snapshot(&fixture.root), before, "absent LL pattern {absent} changed a field or edge");
        validation_check_released(fixture);
        checked += 1;
    }
    checked
}

fn validation_extra() {
    let guarded = validation_guard_matrix();
    let absent = validation_missing_patterns();
    assert_eq!(guarded, 54);
    assert_eq!(absent, 6);
    println!("LL extra validation: {guarded} guard/ownership cases and {absent} absent patterns passed");
}
