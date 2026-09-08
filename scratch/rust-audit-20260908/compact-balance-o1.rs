#![allow(warnings)]
#[derive(Clone)]
pub enum Color {
    R,
    B
}

#[derive(Clone)]
pub enum Tree {
    E,
    T(std::rc::Rc<crate::Color>, std::rc::Rc<crate::Tree>, i64, std::rc::Rc<crate::Tree>)
}


pub fn Test_RBTree_max(mut purs_local_0: i64, mut purs_local_1: i64) -> i64 {
    // AST: Typed(Typed(Abs(..., Typed(Abs(..., Typed(Branch(...)))))))
/* Typed i64 <- i64 : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Local(...) */purs_local_1.clone()) {
        /* Typed i64 <- i64 : Local(...) */purs_local_0
    } else {
        /* Typed i64 <- i64 : Local(...) */purs_local_1
    }
}


pub fn Test_RBTree_makeBlack(mut purs_local_0: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Typed(Abs(..., Typed(Branch(...)))))
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0.clone()).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(std::rc::Rc::new(crate::Color::B), { if let crate::Tree::T(_, ref f, ..) = (purs_local_0.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_0.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }))
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::E)
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}


pub fn Test_RBTree_depth(mut purs_local_0: std::rc::Rc<crate::Tree>) -> i64 {
    // AST: Typed(Typed(Abs(..., Typed(Branch(...)))))
    loop {
        break /* Typed i64 <- i64 : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0.clone()).as_ref(), crate::Tree::E) {
        /* Typed i64 <- i64 : Lit */0
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0.clone()).as_ref(), crate::Tree::T(..)) {
        /* Typed i64 <- i64 : PrimOp(...) */(1 + {
    let mut purs_local_1 = Test_RBTree_depth(/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_0.clone()).as_ref() { f.clone() } else { unreachable!() } });
    /* Typed i64 <- i64 : Let(...) */{
    let mut purs_local_2 = /* Typed i64 <- i64 : App(Var(...)) */Test_RBTree_depth(/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } });
    /* Typed i64 <- i64 : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_1.clone() > /* Typed i64 <- i64 : Local(...) */purs_local_2.clone()) {
        /* Typed i64 <- i64 : Local(...) */purs_local_1
    } else {
        /* Typed i64 <- i64 : Local(...) */purs_local_2
    }
}
})
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    };
    }
}


// Same four Okasaki cases and persistent Rc sharing, with borrowed projections.
fn rb_node(color: std::rc::Rc<crate::Color>, a: std::rc::Rc<crate::Tree>, x: i64, b: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    std::rc::Rc::new(crate::Tree::T(color, a, x, b))
}
pub fn Test_RBTree_balance(color: std::rc::Rc<crate::Color>, left: std::rc::Rc<crate::Tree>, key: i64, right: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    if matches!(color.as_ref(), Color::B) {
        if let Tree::T(lc, ll, lk, lr) = left.as_ref() {
            if matches!(lc.as_ref(), Color::R) {
                if let Tree::T(llc, a, x, b) = ll.as_ref() {
                    if matches!(llc.as_ref(), Color::R) {
                        return rb_node(std::rc::Rc::new(crate::Color::R),
                            rb_node(std::rc::Rc::new(crate::Color::B), a.clone(), *x, b.clone()), *lk,
                            rb_node(std::rc::Rc::new(crate::Color::B), lr.clone(), key, right));
                    }
                }
                if let Tree::T(lrc, b, y, c) = lr.as_ref() {
                    if matches!(lrc.as_ref(), Color::R) {
                        return rb_node(std::rc::Rc::new(crate::Color::R),
                            rb_node(std::rc::Rc::new(crate::Color::B), ll.clone(), *lk, b.clone()), *y,
                            rb_node(std::rc::Rc::new(crate::Color::B), c.clone(), key, right));
                    }
                }
            }
        }
        if let Tree::T(rc, rl, rk, rr) = right.as_ref() {
            if matches!(rc.as_ref(), Color::R) {
                if let Tree::T(rlc, b, y, c) = rl.as_ref() {
                    if matches!(rlc.as_ref(), Color::R) {
                        return rb_node(std::rc::Rc::new(crate::Color::R),
                            rb_node(std::rc::Rc::new(crate::Color::B), left, key, b.clone()), *y,
                            rb_node(std::rc::Rc::new(crate::Color::B), c.clone(), *rk, rr.clone()));
                    }
                }
                if let Tree::T(rrc, c, z, d) = rr.as_ref() {
                    if matches!(rrc.as_ref(), Color::R) {
                        return rb_node(std::rc::Rc::new(crate::Color::R),
                            rb_node(std::rc::Rc::new(crate::Color::B), left, key, rl.clone()), *rk,
                            rb_node(std::rc::Rc::new(crate::Color::B), c.clone(), *z, d.clone()));
                    }
                }
            }
        }
    }
    rb_node(color, left, key, right)
}

pub fn Test_RBTree_ins(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Typed(Abs(..., Typed(Abs(..., Typed(Branch(...)))))))
    loop {
        break /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1.clone()).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(std::rc::Rc::new(crate::Color::R), std::rc::Rc::new(crate::Tree::E), /* Typed i64 <- i64 : Local(...) */purs_local_0, std::rc::Rc::new(crate::Tree::E)))
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1.clone()).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() < /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_balance(/* Typed std::rc::Rc<crate::Color> <- std::rc::Rc<crate::Color> : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }), /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } })
    } else {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_balance(/* Typed std::rc::Rc<crate::Color> <- std::rc::Rc<crate::Color> : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }))
    } else {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T({ if let crate::Tree::T(ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }))
    }
    }
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    };
    }
}


pub fn Test_RBTree_insert(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Typed(Abs(..., Typed(Abs(..., Typed(Let(...)))))))
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Let(...) */{
    let mut purs_local_2 = Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1);
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((purs_local_2.clone()).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(std::rc::Rc::new(crate::Color::B), { if let crate::Tree::T(_, ref f, ..) = (purs_local_2.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_2.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }))
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((purs_local_2).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::E)
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}
}


pub fn Test_RBTree_buildTree(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Typed(Abs(..., Typed(Abs(..., Typed(Branch(...)))))))
    loop {
        break /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if (/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == 0) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1
    } else {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */{
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_insert(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1);
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        continue;
    }
    };
    }
}


use std::alloc::{GlobalAlloc, Layout};
use std::sync::atomic::{AtomicU64, Ordering};
static ALLOCS: AtomicU64 = AtomicU64::new(0);
static BYTES: AtomicU64 = AtomicU64::new(0);
static SMALL: AtomicU64 = AtomicU64::new(0);
struct Alloc;
unsafe impl GlobalAlloc for Alloc {
    unsafe fn alloc(&self, l: Layout) -> *mut u8 {
        #[cfg(audit_allocs)] {
            ALLOCS.fetch_add(1, Ordering::Relaxed);
            BYTES.fetch_add(l.size() as u64, Ordering::Relaxed);
            if l.size() == 24 { SMALL.fetch_add(1, Ordering::Relaxed); }
        }
        mimalloc::MiMalloc.alloc(l)
    }
    unsafe fn dealloc(&self, p: *mut u8, l: Layout) { mimalloc::MiMalloc.dealloc(p,l) }
}
#[global_allocator] static GLOBAL: Alloc = Alloc;

fn validate(t: &Tree, min: i64, max: i64) -> (usize, usize) {
    match t {
        Tree::E => (0, 1),
        Tree::T(c, a, x, b) => {
            assert!(min < *x && *x < max);
            if matches!(c.as_ref(), Color::R) {
                for child in [a, b] {
                    if let Tree::T(cc, ..) = child.as_ref() { assert!(matches!(cc.as_ref(), Color::B)); }
                }
            }
            let (na, ha) = validate(a.as_ref(), min, *x);
            let (nb, hb) = validate(b.as_ref(), *x, max);
            assert_eq!(ha, hb);
            (1 + na + nb, ha + if matches!(c.as_ref(), Color::B) {1} else {0})
        }
    }
}
fn correctness() {
    let mut keys: Vec<i64> = (1..=1000).collect();
    let mut seed = 37_u64;
    for i in (1..keys.len()).rev() {
        seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
        keys.swap(i, seed as usize % (i+1));
    }
    let mut tree = std::rc::Rc::new(Tree::E);
    for x in keys { tree = Test_RBTree_insert(x, tree); }
    assert_eq!(validate(tree.as_ref(), 0, 1001).0, 1000);
    let old = tree.clone();
    tree = Test_RBTree_insert(1001, tree);
    assert_eq!(validate(old.as_ref(), 0, 1001).0, 1000);
    assert_eq!(validate(tree.as_ref(), 0, 1002).0, 1001);
    tree = Test_RBTree_insert(501, tree);
    assert_eq!(validate(tree.as_ref(), 0, 1002).0, 1001);
}
fn run(n: i64) -> i64 {
    let tree = Test_RBTree_buildTree(std::hint::black_box(n), std::rc::Rc::new(Tree::E));
    std::hint::black_box(Test_RBTree_depth(tree))
}
fn main() {
    correctness();
    let repeats: usize = std::env::args().nth(1).unwrap_or("5".into()).parse().unwrap();
    assert_eq!(run(100000), 22);
    for _ in 0..repeats {
        ALLOCS.store(0, Ordering::Relaxed); BYTES.store(0, Ordering::Relaxed); SMALL.store(0, Ordering::Relaxed);
        let start = std::time::Instant::now();
        let value = run(100000);
        let elapsed = start.elapsed().as_secs_f64() * 1e6;
        assert_eq!(value, 22);
        println!("{:.3},{},{},{},{},{}", elapsed, ALLOCS.load(Ordering::Relaxed), BYTES.load(Ordering::Relaxed), SMALL.load(Ordering::Relaxed), std::mem::size_of::<Tree>(), std::mem::size_of::<Color>());
    }
}
