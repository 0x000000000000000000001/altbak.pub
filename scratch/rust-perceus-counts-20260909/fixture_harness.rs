impl tracked::Kind for Tree {
    fn kind(&self) -> &'static str { match self { Tree::Empty => "empty", Tree::Branch(..) => "node" } }
}
fn leaf(color: Shade, key: i64) -> tracked::Rc<Tree> {
    tracked::Rc::new(Tree::Branch(color, tracked::Rc::new(Tree::Empty), key, tracked::Rc::new(Tree::Empty)))
}
fn node(left_color: Shade, right_color: Shade) -> tracked::Rc<Tree> {
    tracked::Rc::new(Tree::Branch(Shade::Black, leaf(left_color, 1), 2, leaf(right_color, 3)))
}
fn main() {
    tracked::phase("unique");
    assert_eq!(PerceusProbe_inspect(node(Shade::Red, Shade::Black)), 11);
    assert_eq!(PerceusProbe_inspect(node(Shade::Black, Shade::Black)), 22);
    assert_eq!(PerceusProbe_inspect(node(Shade::Black, Shade::Red)), 33);
    assert_eq!(PerceusProbe_inspect(tracked::Rc::new(Tree::Empty)), 33);

    tracked::phase("shared");
    let tree = node(Shade::Red, Shade::Black);
    let weak_root = tracked::Rc::downgrade(&tree);
    let child = match tree.as_ref() { Tree::Branch(_, left, _, _) => left.clone(), _ => unreachable!() };
    let weak_child = tracked::Rc::downgrade(&child);
    assert_eq!(PerceusProbe_inspect(tree.clone()), 11);
    assert_eq!(PerceusProbe_inspect(tree.clone()), 11);
    match tree.as_ref() {
        Tree::Branch(Shade::Black, left, 2, right) => {
            assert!(matches!(left.as_ref(), Tree::Branch(Shade::Red, _, 1, _)));
            assert!(matches!(right.as_ref(), Tree::Branch(Shade::Black, _, 3, _)));
        }
        _ => panic!("Original value changed"),
    }
    drop(tree); assert!(weak_root.upgrade().is_none());
    assert!(weak_child.upgrade().is_some());
    drop(child); assert!(weak_child.upgrade().is_none());
    drop(weak_root); drop(weak_child);

    tracked::phase("weak_only");
    let tree = node(Shade::Black, Shade::Black);
    let weak = tracked::Rc::downgrade(&tree);
    assert_eq!(PerceusProbe_inspect(tree), 22);
    assert!(weak.upgrade().is_none()); drop(weak);
    tracked::print_events();
}
