use std::rc::Rc;

type Tree = Option<Rc<Node>>;

#[derive(Clone)]
struct Node {
    color: u8,
    left: Tree,
    value: i64,
    right: Tree,
}

fn node(color: u8, left: Tree, value: i64, right: Tree) -> Rc<Node> {
    Rc::new(Node { color, left, value, right })
}

fn balance(c: u8, a: Tree, x: i64, b: Tree) -> Rc<Node> {
    if c == 1 { // black
        if let Some(a_node) = &a {
            if a_node.color == 0 { // red
                if let Some(al) = &a_node.left {
                    if al.color == 0 {
                        let left = node(1, al.left.clone(), al.value, al.right.clone());
                        let right = node(1, a_node.right.clone(), x, b);
                        return node(0, Some(left), a_node.value, Some(right));
                    }
                }
                if let Some(ar) = &a_node.right {
                    if ar.color == 0 {
                        let left = node(1, a_node.left.clone(), a_node.value, ar.left.clone());
                        let right = node(1, ar.right.clone(), x, b);
                        return node(0, Some(left), ar.value, Some(right));
                    }
                }
            }
        }
        if let Some(b_node) = &b {
            if b_node.color == 0 {
                if let Some(bl) = &b_node.left {
                    if bl.color == 0 {
                        let left = node(1, a, x, bl.left.clone());
                        let right = node(1, bl.right.clone(), b_node.value, b_node.right.clone());
                        return node(0, Some(left), bl.value, Some(right));
                    }
                }
                if let Some(br) = &b_node.right {
                    if br.color == 0 {
                        let left = node(1, a, x, b_node.left.clone());
                        let right = node(1, br.left.clone(), br.value, br.right.clone());
                        return node(0, Some(left), b_node.value, Some(right));
                    }
                }
            }
        }
    }
    node(c, a, x, b)
}

fn ins(x: i64, t: &Tree) -> Rc<Node> {
    let Some(node) = t else {
        return self::node(0, None, x, None);
    };
    if x < node.value {
        let left = ins(x, &node.left);
        balance(node.color, Some(left), node.value, node.right.clone())
    } else if x > node.value {
        let right = ins(x, &node.right);
        balance(node.color, node.left.clone(), node.value, Some(right))
    } else {
        node.clone()
    }
}

fn depth(t: &Tree) -> i64 {
    let Some(node) = t else { return 0; };
    let l = depth(&node.left);
    let r = depth(&node.right);
    1 + std::cmp::max(l, r)
}

pub fn Test_RBTreeFFICheatcode_runRBTreeFFICheatcode(limit: i64) -> i64 {
    let mut acc = None;
    let mut i = limit;
    while i > 0 {
        let mut res = ins(i, &acc);
        Rc::make_mut(&mut res).color = 1; // root is always black
        acc = Some(res);
        i -= 1;
    }
    depth(&acc)
}
