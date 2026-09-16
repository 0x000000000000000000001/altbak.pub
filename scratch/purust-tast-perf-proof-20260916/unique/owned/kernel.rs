// Handwritten alternative lowering of the same Okasaki algorithm.
// Compile with --cfg rc_owner for Rc, otherwise Box. No unsafe code.
#[cfg(rc_owner)]
type Owner<T> = std::rc::Rc<T>;
#[cfg(not(rc_owner))]
type Owner<T> = Box<T>;
type Link = Option<Owner<Node>>;
#[derive(Clone, Copy, PartialEq, Eq)]
enum Color { R, B }
#[derive(Clone)]
struct Node { color: Color, left: Link, key: i64, right: Link }
#[cfg(rc_owner)]
#[inline(always)]
fn own(node: &mut Owner<Node>) -> &mut Node { std::rc::Rc::make_mut(node) }
#[cfg(not(rc_owner))]
#[inline(always)]
fn own(node: &mut Owner<Node>) -> &mut Node { node.as_mut() }
#[inline(always)]
fn red(link: &Link) -> bool { link.as_ref().is_some_and(|node| node.color == Color::R) }
fn rotate_right(link: &mut Link) {
    let mut old = link.take().unwrap();
    let mut promoted = own(&mut old).left.take().unwrap();
    own(&mut old).left = own(&mut promoted).right.take();
    own(&mut promoted).right = Some(old);
    *link = Some(promoted);
}
fn rotate_left(link: &mut Link) {
    let mut old = link.take().unwrap();
    let mut promoted = own(&mut old).right.take().unwrap();
    own(&mut old).right = own(&mut promoted).left.take();
    own(&mut promoted).left = Some(old);
    *link = Some(promoted);
}
fn balance(link: &mut Link) {
    let root = link.as_ref().unwrap();
    if root.color != Color::B { return; }
    let direction = if red(&root.left) {
        let left = root.left.as_ref().unwrap();
        if red(&left.left) { 1 } else if red(&left.right) { 2 } else { 0 }
    } else { 0 };
    let direction = if direction == 0 && red(&root.right) {
        let right = root.right.as_ref().unwrap();
        if red(&right.left) { 3 } else if red(&right.right) { 4 } else { 0 }
    } else { direction };
    match direction {
        1 => rotate_right(link),
        2 => { rotate_left(&mut own(link.as_mut().unwrap()).left); rotate_right(link); },
        3 => { rotate_right(&mut own(link.as_mut().unwrap()).right); rotate_left(link); },
        4 => rotate_left(link),
        _ => return,
    }
    let root = own(link.as_mut().unwrap());
    root.color = Color::R;
    own(root.left.as_mut().unwrap()).color = Color::B;
    own(root.right.as_mut().unwrap()).color = Color::B;
}
fn ins(key: i64, link: &mut Link) {
    let Some(node) = link.as_mut() else {
        *link = Some(Owner::new(Node { color: Color::R, left: None, key, right: None }));
        return;
    };
    let node = own(node);
    if key < node.key { ins(key, &mut node.left); }
    else if key > node.key { ins(key, &mut node.right); }
    else { return; }
    balance(link);
}
fn insert(key: i64, link: &mut Link) {
    ins(key, link);
    own(link.as_mut().unwrap()).color = Color::B;
}
fn depth(link: &Link) -> i64 {
    match link {
        None => 0,
        Some(node) => 1 + depth(&node.left).max(depth(&node.right)),
    }
}
#[inline(never)]
pub fn run(n: i64) -> i64 {
    let mut tree = None;
    for key in (1..=n).rev() { insert(key, &mut tree); }
    let tree = std::hint::black_box(tree);
    let answer = depth(&tree);
    drop(tree);
    answer
}
