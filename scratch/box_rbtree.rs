use std::time::Instant;

#[derive(Clone, PartialEq, Debug)]
enum Color {
    R,
    B,
}

#[derive(Clone)]
enum Tree {
    E,
    T(Color, Box<Tree>, i64, Box<Tree>),
}

fn make_black(t: Box<Tree>) -> Box<Tree> {
    if let Tree::T(_, l, x, r) = *t {
        Box::new(Tree::T(Color::B, l, x, r))
    } else {
        t
    }
}

fn balance(c: Color, a: Box<Tree>, x: i64, b: Box<Tree>) -> Box<Tree> {
    if c == Color::B {
        if let Tree::T(Color::R, ref ll, lx, ref lr) = *a {
            if let Tree::T(Color::R, ref lll, llx, ref llr) = **ll {
                return Box::new(Tree::T(Color::R, Box::new(Tree::T(Color::B, lll.clone(), llx, llr.clone())), lx, Box::new(Tree::T(Color::B, lr.clone(), x, b))));
            }
            if let Tree::T(Color::R, ref lrl, lrx, ref lrr) = **lr {
                return Box::new(Tree::T(Color::R, Box::new(Tree::T(Color::B, ll.clone(), lx, lrl.clone())), lrx, Box::new(Tree::T(Color::B, lrr.clone(), x, b))));
            }
        }
        if let Tree::T(Color::R, ref rl, rx, ref rr) = *b {
            if let Tree::T(Color::R, ref rll, rlx, ref rlr) = **rl {
                return Box::new(Tree::T(Color::R, Box::new(Tree::T(Color::B, a, x, rll.clone())), rlx, Box::new(Tree::T(Color::B, rlr.clone(), rx, rr.clone()))));
            }
            if let Tree::T(Color::R, ref rrl, rrx, ref rrr) = **rr {
                return Box::new(Tree::T(Color::R, Box::new(Tree::T(Color::B, a, x, rl.clone())), rx, Box::new(Tree::T(Color::B, rrl.clone(), rrx, rrr.clone()))));
            }
        }
    }
    Box::new(Tree::T(c, a, x, b))
}

fn ins(x: i64, t: Box<Tree>) -> Box<Tree> {
    match *t {
        Tree::E => Box::new(Tree::T(Color::R, Box::new(Tree::E), x, Box::new(Tree::E))),
        Tree::T(c, a, y, b) => {
            if x < y {
                balance(c, ins(x, a), y, b)
            } else if x > y {
                balance(c, a, y, ins(x, b))
            } else {
                Box::new(Tree::T(c, a, y, b))
            }
        }
    }
}

fn main() {
    let mut t = Box::new(Tree::E);
    let start = Instant::now();
    for i in 0..100_000 {
        t = make_black(ins(i, t));
    }
    let duration = start.elapsed();
    println!("Insertion 100k (Box Native): {:?}", duration);
}
