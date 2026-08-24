use std::rc::Rc;

#[derive(Clone, Copy, PartialEq)]
enum Color {
    Red,
    Black,
}

enum Tree {
    Leaf,
    Node(Color, Rc<Tree>, i64, Rc<Tree>),
}

fn alloc(c: Color, l: Rc<Tree>, v: i64, r: Rc<Tree>) -> Rc<Tree> {
    Rc::new(Tree::Node(c, l, v, r))
}

fn balance(c: Color, a: Rc<Tree>, x: i64, b: Rc<Tree>) -> Rc<Tree> {
    if c == Color::Black {
        if let Tree::Node(Color::Red, al, a_val, ar) = &*a {
            if let Tree::Node(Color::Red, all, all_val, alr) = &**al {
                return alloc(Color::Red, alloc(Color::Black, all.clone(), *all_val, alr.clone()), *a_val, alloc(Color::Black, ar.clone(), x, b.clone()));
            }
            if let Tree::Node(Color::Red, arl, arl_val, arr) = &**ar {
                return alloc(Color::Red, alloc(Color::Black, al.clone(), *a_val, arl.clone()), *arl_val, alloc(Color::Black, arr.clone(), x, b.clone()));
            }
        }
        if let Tree::Node(Color::Red, bl, b_val, br) = &*b {
            if let Tree::Node(Color::Red, bll, bll_val, blr) = &**bl {
                return alloc(Color::Red, alloc(Color::Black, a.clone(), x, bll.clone()), *bll_val, alloc(Color::Black, blr.clone(), *b_val, br.clone()));
            }
            if let Tree::Node(Color::Red, brl, brl_val, brr) = &**br {
                return alloc(Color::Red, alloc(Color::Black, a.clone(), x, bl.clone()), *b_val, alloc(Color::Black, brl.clone(), *brl_val, brr.clone()));
            }
        }
    }
    alloc(c, a, x, b)
}

fn ins(x: i64, t: &Rc<Tree>) -> Rc<Tree> {
    match &**t {
        Tree::Leaf => alloc(Color::Red, Rc::new(Tree::Leaf), x, Rc::new(Tree::Leaf)),
        Tree::Node(c, l, y, r) => {
            if x < *y {
                balance(*c, ins(x, l), *y, r.clone())
            } else if x > *y {
                balance(*c, l.clone(), *y, ins(x, r))
            } else {
                t.clone()
            }
        }
    }
}

fn insert(x: i64, t: &Rc<Tree>) -> Rc<Tree> {
    let res = ins(x, t);
    if let Tree::Node(_, l, y, r) = &*res {
        alloc(Color::Black, l.clone(), *y, r.clone())
    } else {
        Rc::new(Tree::Leaf)
    }
}

fn build_tree(n: i64, mut acc: Rc<Tree>) -> Rc<Tree> {
    let mut i = n;
    while i > 0 {
        acc = insert(i, &acc);
        i -= 1;
    }
    acc
}

fn depth(t: &Rc<Tree>) -> i64 {
    match &**t {
        Tree::Leaf => 0,
        Tree::Node(_, l, _, r) => {
            let ld = depth(l);
            let rd = depth(r);
            if ld > rd {
                1 + ld
            } else {
                1 + rd
            }
        }
    }
}

pub fn Test_RBTreeFFI_runRBTreeFFI(mut limit: i64) -> i64 {
    let t = build_tree(limit, Rc::new(Tree::Leaf));
    depth(&t)
}