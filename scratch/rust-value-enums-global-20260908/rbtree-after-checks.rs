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
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::B, { if let crate::Tree::T(_, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }))
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::E)
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}


pub fn Test_RBTree_depth(mut purs_local_0: std::rc::Rc<crate::Tree>) -> i64 {
    // AST: Typed(Typed(Abs(..., Typed(Branch(...)))))
    loop {
        break /* Typed i64 <- i64 : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::E) {
        /* Typed i64 <- i64 : Lit */0
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::T(..)) {
        /* Typed i64 <- i64 : PrimOp(...) */(1 + {
    let mut purs_local_1 = Test_RBTree_depth(/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } });
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
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Local(...) */purs_local_0.clone()), crate::Color::B) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } };
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
        break /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, std::rc::Rc::new(crate::Tree::E), /* Typed i64 <- i64 : Local(...) */purs_local_0, std::rc::Rc::new(crate::Tree::E)))
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() < /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } })
    } else {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }))
    } else {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T({ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }))
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((purs_local_2).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::B, { if let crate::Tree::T(_, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }))
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
use std::sync::atomic::{AtomicUsize, Ordering};
static ALLOCS: AtomicUsize = AtomicUsize::new(0);
static FREES: AtomicUsize = AtomicUsize::new(0);
static BYTES: AtomicUsize = AtomicUsize::new(0);
static SMALL: AtomicUsize = AtomicUsize::new(0);
struct Counting;
unsafe impl GlobalAlloc for Counting {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        ALLOCS.fetch_add(1, Ordering::Relaxed);
        BYTES.fetch_add(layout.size(), Ordering::Relaxed);
        if layout.size() == 24 { SMALL.fetch_add(1, Ordering::Relaxed); }
        mimalloc::MiMalloc.alloc(layout)
    }
    unsafe fn dealloc(&self, ptr: *mut u8, layout: Layout) {
        FREES.fetch_add(1, Ordering::Relaxed);
        mimalloc::MiMalloc.dealloc(ptr, layout)
    }
}
#[global_allocator] static ALLOCATOR: Counting = Counting;
fn red(color: &Color) -> bool { matches!(color, Color::R) }
fn validate(tree: &Tree, min: i64, max: i64) -> (usize, usize) {
    match tree {
        Tree::E => (0, 1),
        Tree::T(color, left, key, right) => {
            assert!(min < *key && *key < max);
            if red(color) {
                for child in [left, right] {
                    if let Tree::T(c, ..) = child.as_ref() { assert!(!red(c)); }
                }
            }
            let (nl, hl) = validate(left, min, *key);
            let (nr, hr) = validate(right, *key, max);
            assert_eq!(hl, hr);
            (1 + nl + nr, hl + if red(color) { 0 } else { 1 })
        }
    }
}
fn checked_root(tree: &Tree, count: usize, max: i64) {
    if let Tree::T(color, ..) = tree { assert!(!red(color)); }
    assert_eq!(validate(tree, 0, max).0, count);
}
fn correctness() {
    // The four rotation shapes, then mixed insertion order and a retained root.
    for keys in [[3, 2, 1], [3, 1, 2], [1, 3, 2], [1, 2, 3]] {
        let mut tree = std::rc::Rc::new(Tree::E);
        for (i, key) in keys.into_iter().enumerate() {
            tree = Test_RBTree_insert(key, tree);
            checked_root(&tree, i + 1, 4);
        }
    }
    let mut keys: Vec<i64> = (1..=1000).collect();
    let mut seed = 37_u64;
    for i in (1..keys.len()).rev() {
        seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
        keys.swap(i, seed as usize % (i + 1));
    }
    let mut tree = std::rc::Rc::new(Tree::E);
    for key in keys { tree = Test_RBTree_insert(key, tree); }
    checked_root(&tree, 1000, 1001);
    let old = tree.clone();
    tree = Test_RBTree_insert(1001, tree);
    checked_root(&tree, 1001, 1002);
    checked_root(&old, 1000, 1001);
    tree = Test_RBTree_insert(501, tree);
    checked_root(&tree, 1001, 1002);
    drop(tree);
    checked_root(&old, 1000, 1001);
}
fn main() {
    correctness();
    for counter in [&ALLOCS, &FREES, &BYTES, &SMALL] { counter.store(0, Ordering::Relaxed); }
    let tree = Test_RBTree_buildTree(std::hint::black_box(100000), std::rc::Rc::new(Tree::E));
    assert_eq!(Test_RBTree_depth(tree), 22);
    let allocs = ALLOCS.load(Ordering::Relaxed);
    let frees = FREES.load(Ordering::Relaxed);
    let bytes = BYTES.load(Ordering::Relaxed);
    let small = SMALL.load(Ordering::Relaxed);
    assert_eq!(allocs, frees);
    println!("{} {} {} {} {}", allocs, frees, bytes, small, std::mem::size_of::<Tree>());
}
