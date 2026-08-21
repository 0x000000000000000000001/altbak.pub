use perceus_ptr::PerceusPtr;
use std::time::Instant;

#[derive(Clone, PartialEq, Debug)]
enum Color {
    R,
    B,
}

#[derive(Clone)]
enum Tree {
    E,
    T(Color, PerceusPtr<Tree>, i64, PerceusPtr<Tree>),
}

fn make_black(mut t: PerceusPtr<Tree>) -> PerceusPtr<Tree> {
    if !t.is_unique() {
        t = PerceusPtr::new((*t).clone());
    }
    let mut_t = unsafe { PerceusPtr::force_mut(&mut t) };
    if let Tree::T(ref mut c, _, _, _) = mut_t {
        *c = Color::B;
    }
    t
}

// In FBIP, balance takes ownership of the node to reuse its memory.
fn balance(mut reuse: PerceusPtr<Tree>, c: Color, a: PerceusPtr<Tree>, x: i64, b: PerceusPtr<Tree>) -> PerceusPtr<Tree> {
    // If it's Black, we check for red-red violations
    if c == Color::B {
        if let Tree::T(Color::R, ref ll, lx, ref lr) = *a {
            if let Tree::T(Color::R, ref lll, llx, ref llr) = **ll {
                // T R (T R a x b) y c
                let mut new_left = PerceusPtr::new(Tree::T(Color::B, lll.clone(), llx, llr.clone())); // ideally reuse ll
                let mut new_right = PerceusPtr::new(Tree::T(Color::B, lr.clone(), x, b.clone())); // ideally reuse a
                if !reuse.is_unique() { reuse = PerceusPtr::new((*reuse).clone()); }
                let mut_t = unsafe { PerceusPtr::force_mut(&mut reuse) };
                *mut_t = Tree::T(Color::R, new_left, lx, new_right);
                return reuse;
            }
            if let Tree::T(Color::R, ref lrl, lrx, ref lrr) = **lr {
                // T R a x (T R b y c)
                let mut new_left = PerceusPtr::new(Tree::T(Color::B, ll.clone(), lx, lrl.clone())); // reuse lr
                let mut new_right = PerceusPtr::new(Tree::T(Color::B, lrr.clone(), x, b.clone())); // reuse a
                if !reuse.is_unique() { reuse = PerceusPtr::new((*reuse).clone()); }
                let mut_t = unsafe { PerceusPtr::force_mut(&mut reuse) };
                *mut_t = Tree::T(Color::R, new_left, lrx, new_right);
                return reuse;
            }
        }
        if let Tree::T(Color::R, ref rl, rx, ref rr) = *b {
            if let Tree::T(Color::R, ref rll, rlx, ref rlr) = **rl {
                // a x (T R (T R b y c) z d)
                let mut new_left = PerceusPtr::new(Tree::T(Color::B, a.clone(), x, rll.clone())); // reuse rl
                let mut new_right = PerceusPtr::new(Tree::T(Color::B, rlr.clone(), rx, rr.clone())); // reuse b
                if !reuse.is_unique() { reuse = PerceusPtr::new((*reuse).clone()); }
                let mut_t = unsafe { PerceusPtr::force_mut(&mut reuse) };
                *mut_t = Tree::T(Color::R, new_left, rlx, new_right);
                return reuse;
            }
            if let Tree::T(Color::R, ref rrl, rrx, ref rrr) = **rr {
                // a x (T R b y (T R c z d))
                let mut new_left = PerceusPtr::new(Tree::T(Color::B, a.clone(), x, rl.clone())); // reuse b
                let mut new_right = PerceusPtr::new(Tree::T(Color::B, rrl.clone(), rrx, rrr.clone())); // reuse rr
                if !reuse.is_unique() { reuse = PerceusPtr::new((*reuse).clone()); }
                let mut_t = unsafe { PerceusPtr::force_mut(&mut reuse) };
                *mut_t = Tree::T(Color::R, new_left, rx, new_right);
                return reuse;
            }
        }
    }
    
    // No rotation needed, reuse the parent node!
    if !reuse.is_unique() { reuse = PerceusPtr::new((*reuse).clone()); }
    let mut_t = unsafe { PerceusPtr::force_mut(&mut reuse) };
    *mut_t = Tree::T(c, a, x, b);
    reuse
}

fn ins(x: i64, mut t: PerceusPtr<Tree>) -> PerceusPtr<Tree> {
    if !t.is_unique() {
        t = PerceusPtr::new((*t).clone());
    }
    let mut_t = unsafe { PerceusPtr::force_mut(&mut t) };
    
    // We clone the components because we are about to overwrite *mut_t in balance
    let res = match mut_t {
        Tree::E => {
            *mut_t = Tree::T(Color::R, PerceusPtr::new(Tree::E), x, PerceusPtr::new(Tree::E));
            t
        }
        Tree::T(c, ref mut a, y, ref mut b) => {
            let cx = c.clone();
            let yx = *y;
            if x < *y {
                let new_a = ins(x, a.clone());
                let bx = b.clone();
                // Pass t as the reuse node for balance
                balance(t, cx, new_a, yx, bx)
            } else if x > *y {
                let ax = a.clone();
                let new_b = ins(x, b.clone());
                balance(t, cx, ax, yx, new_b)
            } else {
                t
            }
        }
    };
    res
}

fn main() {
    let mut t = PerceusPtr::new(Tree::E);
    let start = Instant::now();
    for i in 0..100_000 {
        t = make_black(ins(i, t));
    }
    let duration = start.elapsed();
    println!("Insertion 100k (FBIP RBTree Native): {:?}", duration);
}
