#![allow(warnings)]
#[derive(Clone, Copy)]
pub enum Color {
    R,
    B
}

#[derive(Clone)]
pub enum Tree {
    E,
    T(crate::Color, std::rc::Rc<crate::Tree>, i64, std::rc::Rc<crate::Tree>)
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
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::B, { if let crate::Tree::T(_, ref f, ..) = (purs_local_0.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_0.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }))
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


pub fn Test_RBTree_balance(mut purs_local_0: crate::Color, mut purs_local_1: std::rc::Rc<crate::Tree>, mut purs_local_2: i64, mut purs_local_3: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Typed(Abs(..., Typed(Abs(..., Typed(Abs(..., Typed(Abs(..., Typed(Branch(...)))))))))))
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Local(...) */purs_local_0.clone()).as_ref(), crate::Color::B) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1.clone()).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3.clone()).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3.clone()).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3.clone()).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3.clone()).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3.clone()).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3.clone()).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3.clone()).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
}


pub fn Test_RBTree_ins(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Typed(Abs(..., Typed(Abs(..., Typed(Branch(...)))))))
    loop {
        break /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1.clone()).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::E), /* Typed i64 <- i64 : Local(...) */purs_local_0, std::rc::Rc::new(crate::Tree::E)))
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1.clone()).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() < /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }), /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } })
    } else {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1.clone()).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }))
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
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::B, { if let crate::Tree::T(_, ref f, ..) = (purs_local_2.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_2.clone()).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }))
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


impl Color { fn as_ref(&self) -> &Self { self } }

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
