#[derive(Clone)]
struct Node {
    color: u8,
    left: Option<Box<Node>>,
    value: i64,
    right: Option<Box<Node>>,
}

impl Node {
    fn new(color: u8, left: Option<Box<Node>>, value: i64, right: Option<Box<Node>>) -> Option<Box<Node>> {
        Some(Box::new(Node { color, left, value, right }))
    }
}

fn balance(c: u8, a: Option<Box<Node>>, x: i64, b: Option<Box<Node>>) -> Option<Box<Node>> {
    if c == 1 { // black
        if let Some(ref a_node) = a {
            if a_node.color == 0 { // red
                if let Some(ref al) = a_node.left {
                    if al.color == 0 {
                        return Node::new(0, 
                            Node::new(1, al.left.clone(), al.value, al.right.clone()), 
                            a_node.value, 
                            Node::new(1, a_node.right.clone(), x, b)
                        );
                    }
                }
                if let Some(ref ar) = a_node.right {
                    if ar.color == 0 {
                        return Node::new(0, 
                            Node::new(1, a_node.left.clone(), a_node.value, ar.left.clone()), 
                            ar.value, 
                            Node::new(1, ar.right.clone(), x, b)
                        );
                    }
                }
            }
        }
        if let Some(ref b_node) = b {
            if b_node.color == 0 { // red
                if let Some(ref bl) = b_node.left {
                    if bl.color == 0 {
                        return Node::new(0, 
                            Node::new(1, a, x, bl.left.clone()), 
                            bl.value, 
                            Node::new(1, bl.right.clone(), b_node.value, b_node.right.clone())
                        );
                    }
                }
                if let Some(ref br) = b_node.right {
                    if br.color == 0 {
                        return Node::new(0, 
                            Node::new(1, a, x, b_node.left.clone()), 
                            b_node.value, 
                            Node::new(1, br.left.clone(), br.value, br.right.clone())
                        );
                    }
                }
            }
        }
    }
    Node::new(c, a, x, b)
}

fn ins(x: i64, t: Option<Box<Node>>) -> Option<Box<Node>> {
    match t {
        None => Node::new(0, None, x, None),
        Some(mut node) => {
            if x < node.value {
                balance(node.color, ins(x, node.left.take()), node.value, node.right.take())
            } else if x > node.value {
                balance(node.color, node.left.take(), node.value, ins(x, node.right.take()))
            } else {
                Some(node)
            }
        }
    }
}

fn depth(t: Option<Box<Node>>) -> i64 {
    match t {
        None => 0,
        Some(node) => {
            let l = depth(node.left);
            let r = depth(node.right);
            1 + std::cmp::max(l, r)
        }
    }
}

pub fn Test_RBTreeFFI_runRBTreeFFI(limit: i64) -> i64 {
    let mut acc = None;
    let mut i = limit;
    while i > 0 {
        let mut res = ins(i, acc).unwrap();
        res.color = 1; // root is always black
        acc = Some(res);
        i -= 1;
    }
    depth(acc)
}
