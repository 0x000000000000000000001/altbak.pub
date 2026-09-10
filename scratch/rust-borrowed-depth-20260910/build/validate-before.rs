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
impl Tree { pub fn __purust_take(&mut self) -> std::option::Option<Self> { std::option::Option::Some(std::mem::replace(self, Self::E)) } }


fn Test_RBTree___purust_rebuild_E(mut __purust_cell: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
let payload = crate::Tree::E;
if let std::option::Option::Some(slot) = std::rc::Rc::get_mut(&mut __purust_cell) { *slot = payload; __purust_cell } else { std::rc::Rc::new(payload) }
}
fn Test_RBTree___purust_rebuild_T(a0: crate::Color, a1: std::rc::Rc<crate::Tree>, a2: i64, a3: std::rc::Rc<crate::Tree>, mut __purust_cell: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
let payload = crate::Tree::T(a0, a1, a2, a3);
if let std::option::Option::Some(slot) = std::rc::Rc::get_mut(&mut __purust_cell) { *slot = payload; __purust_cell } else { std::rc::Rc::new(payload) }
}
#[inline]
fn Test_RBTree_balance__purust_child_rebuilds(_purust_guard_arg_0: &crate::Color, _purust_guard_arg_1: &std::rc::Rc<crate::Tree>, _purust_guard_arg_2: &i64, _purust_guard_arg_3: &std::rc::Rc<crate::Tree>) -> bool { !(matches!(_purust_guard_arg_0, crate::Color::B)) || (if matches!((_purust_guard_arg_1).as_ref(), crate::Tree::T(..)) { if matches!((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R) { if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) })) } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) }) } else { if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) })) } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) } } } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) } } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) }) }




















pub fn Test_RBTree_max(mut purs_local_0: i64, mut purs_local_1: i64) -> i64 {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
/* Typed i64 <- i64 : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Local(...) */purs_local_1.clone()) {
        /* Typed i64 <- i64 : Local(...) */purs_local_0
    } else {
        /* Typed i64 <- i64 : Local(...) */purs_local_1
    }
}

pub fn Test_RBTree_makeBlack(mut purs_local_0: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Branch(...))))
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */{ let _new_field = crate::Color::B; let mut purs_local_0 = purs_local_0; if let std::option::Option::Some(crate::Tree::T(_updated_field, _, _, _)) = std::rc::Rc::get_mut(&mut purs_local_0) { *_updated_field = _new_field; purs_local_0 } else { let _rebuilt = crate::Tree::T(crate::Color::B, { if let crate::Tree::T(_, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }); let mut _reused = purs_local_0; if let std::option::Option::Some(_slot) = std::rc::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { std::rc::Rc::new(_rebuilt) } } }
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_0
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}






pub fn Test_RBTree_depth(mut purs_local_0: std::rc::Rc<crate::Tree>) -> i64 {
    // AST: Typed(Abs(..., Typed(Branch(...))))
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
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))))))
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Local(...) */purs_local_0.clone()), crate::Color::B) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => { let mut _owned_purs_local_1_1 = _owned_purs_local_1_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_1_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_1_0, mut _owned__owned_purs_local_1_1_1, mut _owned__owned_purs_local_1_1_2, mut _owned__owned_purs_local_1_1_3)) => {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_1_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_1_3;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1)))
}
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
}, _ => unreachable!() } }
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1;
    { let mut _owned_purs_local_1_3 = _owned_purs_local_1_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_1_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_3_0, mut _owned__owned_purs_local_1_3_1, mut _owned__owned_purs_local_1_3_2, mut _owned__owned_purs_local_1_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
        { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1;
    { let mut _owned_purs_local_1_3 = _owned_purs_local_1_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_1_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_3_0, mut _owned__owned_purs_local_1_3_1, mut _owned__owned_purs_local_1_3_2, mut _owned__owned_purs_local_1_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
}, _ => unreachable!() } }
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
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
    loop {
        break /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */std::rc::Rc::new(crate::Tree::T(crate::Color::R, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1.clone(), /* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1))
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() < /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */if !(/* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::B)) { { /* purust child call: retained fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_child_slot) = std::rc::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_child_field_0, _purust_child_field_1, _purust_child_field_2, _purust_child_field_3) = _purust_child_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_child_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_child_field_2).clone(); let mut _owned_purs_local_1_1 = std::mem::replace(_purust_child_field_1, (*_purust_child_field_3).clone()); let _purust_new_child = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1); *_purust_child_field_1 = _purust_new_child; purs_local_1 } else { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = std::rc::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3) }, _ => unreachable!() } } } } else { { /* purust child call: post-call fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_post_slot) = std::rc::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) = _purust_post_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_post_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_post_field_2).clone(); let mut _owned_purs_local_1_1 = std::mem::replace(_purust_post_field_1, (*_purust_post_field_3).clone()); let _purust_post_child = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1); *_purust_post_field_1 = _purust_post_child; if Test_RBTree_balance__purust_child_rebuilds(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) { purs_local_1 } else { let std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) = _purust_post_slot.__purust_take() else { unreachable!() }; Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1) } } else { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = std::rc::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3) }, _ => unreachable!() } } } }
    } else {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */if !(/* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::B)) { { /* purust child call: retained fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_child_slot) = std::rc::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_child_field_0, _purust_child_field_1, _purust_child_field_2, _purust_child_field_3) = _purust_child_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_child_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_child_field_2).clone(); let mut _owned_purs_local_1_3 = std::mem::replace(_purust_child_field_3, (*_purust_child_field_1).clone()); let _purust_new_child = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *_purust_child_field_3 = _purust_new_child; purs_local_1 } else { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = std::rc::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)) }, _ => unreachable!() } } } } else { { /* purust child call: post-call fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_post_slot) = std::rc::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) = _purust_post_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_post_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_post_field_2).clone(); let mut _owned_purs_local_1_3 = std::mem::replace(_purust_post_field_3, (*_purust_post_field_1).clone()); let _purust_post_child = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *_purust_post_field_3 = _purust_post_child; if Test_RBTree_balance__purust_child_rebuilds(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) { purs_local_1 } else { let std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) = _purust_post_slot.__purust_take() else { unreachable!() }; Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1) } } else { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = std::rc::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)) }, _ => unreachable!() } } } }
    } else {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */{ let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => { let _rebuilt = crate::Tree::T(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *std::rc::Rc::get_mut(&mut purs_local_1).unwrap() = _rebuilt; purs_local_1 }, std::option::Option::None => { let _rebuilt = crate::Tree::T({ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }); let mut _reused = purs_local_1; if let std::option::Option::Some(_slot) = std::rc::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { std::rc::Rc::new(_rebuilt) } }, _ => unreachable!() } }
    }
    }
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    };
    }
}

pub fn Test_RBTree_insert(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Let(...))))))
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Let(...) */{
    let mut purs_local_2 = Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1);
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((purs_local_2).as_ref(), crate::Tree::T(..)) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : CtorSaturated(...) */{ let _new_field = crate::Color::B; let mut purs_local_2 = purs_local_2; if let std::option::Option::Some(crate::Tree::T(_updated_field, _, _, _)) = std::rc::Rc::get_mut(&mut purs_local_2) { *_updated_field = _new_field; purs_local_2 } else { let _rebuilt = crate::Tree::T(crate::Color::B, { if let crate::Tree::T(_, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }); let mut _reused = purs_local_2; if let std::option::Option::Some(_slot) = std::rc::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { std::rc::Rc::new(_rebuilt) } } }
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((purs_local_2).as_ref(), crate::Tree::E) {
        /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_2
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}
}

pub fn Test_RBTree_buildTree(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
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





























fn Test_RBTree_balance__purust_reuse(mut purs_local_0: crate::Color, mut purs_local_1: std::rc::Rc<crate::Tree>, mut purs_local_2: i64, mut purs_local_3: std::rc::Rc<crate::Tree>, mut __purust_cell: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Branch(...))))
/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Local(...) */purs_local_0.clone()), crate::Color::B) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => { let mut _owned_purs_local_1_1 = _owned_purs_local_1_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_1_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_1_0, mut _owned__owned_purs_local_1_1_1, mut _owned__owned_purs_local_1_1_2, mut _owned__owned_purs_local_1_1_3)) => {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_1_1;
    {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_1_3;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1;
    { let mut _owned_purs_local_1_3 = _owned_purs_local_1_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_1_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_3_0, mut _owned__owned_purs_local_1_3_1, mut _owned__owned_purs_local_1_3_2, mut _owned__owned_purs_local_1_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = std::rc::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1;
    { let mut _owned_purs_local_1_3 = _owned_purs_local_1_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_1_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_3_0, mut _owned__owned_purs_local_1_3_1, mut _owned__owned_purs_local_1_3_2, mut _owned__owned_purs_local_1_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = std::rc::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = std::rc::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, std::rc::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed std::rc::Rc<crate::Tree> <- std::rc::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
}





fn validate(tree: &Tree, low: i64, high: i64) -> (usize, usize) {
    match tree {
        Tree::E => (0, 1),
        Tree::T(color, left, key, right) => {
            assert!(low < *key && *key < high);
            if matches!(color, Color::R) {
                for child in [left, right] {
                    if let Tree::T(c, ..) = child.as_ref() { assert!(!matches!(c, Color::R)); }
                }
            }
            let (nl, hl) = validate(left, low, *key);
            let (nr, hr) = validate(right, *key, high);
            assert_eq!(hl, hr);
            (1 + nl + nr, hl + usize::from(matches!(color, Color::B)))
        }
    }
}
fn check(tree: &Tree, count: usize) {
    if let Tree::T(color, ..) = tree { assert!(matches!(color, Color::B)); }
    assert_eq!(validate(tree, i64::MIN, i64::MAX).0, count);
}
fn keys(tree: &Tree) -> Vec<i64> {
    match tree {
        Tree::E => Vec::new(),
        Tree::T(_, left, key, right) => {
            let mut result = keys(left); result.push(*key); result.extend(keys(right)); result
        }
    }
}
fn original_checks() {
    assert_eq!(std::mem::size_of::<std::rc::Rc<Tree>>(), std::mem::size_of::<std::rc::Rc<Tree>>());
    assert_eq!(std::mem::size_of::<Tree>(), 32);
    for order in [[3,2,1], [3,1,2], [1,3,2], [1,2,3]] {
        let mut tree = std::rc::Rc::new(Tree::E);
        for (i, key) in order.into_iter().enumerate() {
            tree = Test_RBTree_insert(key, tree); check(&tree, i+1);
        }
    }
    // Verify get_mut's weak-reference exclusion and unwrap's distinct rule.
    let mut tree = std::rc::Rc::new(Tree::T(Color::B,
        std::rc::Rc::new(Tree::E), 1, std::rc::Rc::new(Tree::E)));
    let weak = std::rc::Rc::downgrade(&tree);
    assert!(std::rc::Rc::get_mut(&mut tree).is_none());
    tree = Test_RBTree_insert(2, tree); check(&tree, 2);
    assert!(weak.upgrade().is_none()); drop(tree); drop(weak);

    let tree = Test_RBTree_buildTree(std::hint::black_box(100000), std::rc::Rc::new(Tree::E));
    check(&tree, 100000);
    assert_eq!(Test_RBTree_depth(tree), 22);

    let mut tree = std::rc::Rc::new(Tree::E);
    let mut versions = Vec::new();
    let mut expected = Vec::new();
    // A permutation covers both descent directions and all retained versions.
    for n in 0..200 {
        versions.push((tree.clone(), expected.clone()));
        let key = (n * 73) % 200 + 1;
        tree = Test_RBTree_insert(key, tree);
        expected.push(key); expected.sort(); check(&tree, expected.len());
        assert_eq!(keys(&tree), expected);
        for (old, keys_before) in &versions { check(old, keys_before.len()); assert_eq!(&keys(old), keys_before); }
    }
    assert!(Test_RBTree_depth(tree.clone()) > 0);
    for (old, expected) in &versions { assert_eq!(&keys(old), expected); }
    drop(tree); drop(versions);
}
// Include after the extracted generated Tree/Color and insertion functions.
// Uses native std::rc::Rc; contains no timing or allocation instrumentation.
fn b13_keys(tree: &Tree, out: &mut Vec<i64>) {
    if let Tree::T(_, left, key, right) = tree {
        b13_keys(left, out);
        out.push(*key);
        b13_keys(right, out);
    }
}

fn b13_shape(tree: &Tree) -> String {
    match tree {
        Tree::E => "E".into(),
        Tree::T(color, left, key, right) => format!("({} {} {} {})",
            if matches!(color, Color::R) { "R" } else { "B" },
            b13_shape(left), key, b13_shape(right)),
    }
}

fn b13_validate(tree: &Tree, low: Option<i64>, high: Option<i64>) -> (usize, usize) {
    match tree {
        Tree::E => (0, 1),
        Tree::T(color, left, key, right) => {
            assert!(low.is_none_or(|bound| bound < *key), "BST lower bound");
            assert!(high.is_none_or(|bound| *key < bound), "BST upper bound");
            if matches!(color, Color::R) {
                assert!(!matches!(left.as_ref(), Tree::T(Color::R, ..)), "red left child");
                assert!(!matches!(right.as_ref(), Tree::T(Color::R, ..)), "red right child");
            }
            let (ln, lh) = b13_validate(left, low, Some(*key));
            let (rn, rh) = b13_validate(right, Some(*key), high);
            assert_eq!(lh, rh, "black height at key {key}");
            (1 + ln + rn, lh + usize::from(matches!(color, Color::B)))
        }
    }
}

fn b13_check(tree: &Tree, expected: &std::collections::BTreeSet<i64>, black_root: bool) {
    if black_root {
        assert!(!matches!(tree, Tree::T(Color::R, ..)), "nonblack root");
    }
    assert_eq!(b13_validate(tree, None, None).0, expected.len());
    let mut actual = Vec::new();
    b13_keys(tree, &mut actual);
    assert_eq!(actual, expected.iter().copied().collect::<Vec<_>>());
}

fn b13_collect_weaks(tree: &std::rc::Rc<Tree>, out: &mut Vec<std::rc::Weak<Tree>>) {
    out.push(std::rc::Rc::downgrade(tree));
    if let Tree::T(_, left, _, right) = tree.as_ref() {
        b13_collect_weaks(left, out);
        b13_collect_weaks(right, out);
    }
}

fn b13_black_leaf(key: i64) -> std::rc::Rc<Tree> {
    std::rc::Rc::new(Tree::T(Color::B,
        std::rc::Rc::new(Tree::E), key, std::rc::Rc::new(Tree::E)))
}

// Neither identity nor survival of a weak-only allocation is imposed: both may
// change under legal copy-on-write implementations. Any surviving weak must still
// observe its old value, and all weak handles must expire when all owners drop.
fn b13_red_parent_cases() {
    let mut cases = 0;
    let mut unique_cases = 0;
    let mut unique_same_address = 0;
    for key in [3, 7, 13, 17] {
        for parent_shared in [false, true] {
            for parent_weak in [false, true] {
                for left_shared in [false, true] {
                    for right_shared in [false, true] {
                        for children_weak in [false, true] {
                            let left = b13_black_leaf(5);
                            let right = b13_black_leaf(15);
                            let left_old = left_shared.then(|| left.clone());
                            let right_old = right_shared.then(|| right.clone());
                            let left_shape = b13_shape(&left);
                            let right_shape = b13_shape(&right);
                            let left_weak = children_weak.then(|| std::rc::Rc::downgrade(&left));
                            let right_weak = children_weak.then(|| std::rc::Rc::downgrade(&right));
                            let root = std::rc::Rc::new(Tree::T(Color::R, left, 10, right));
                            let root_shape = b13_shape(&root);
                            let old = parent_shared.then(|| root.clone());
                            let weak = parent_weak.then(|| std::rc::Rc::downgrade(&root));
                            let address = std::rc::Rc::as_ptr(&root);
                            let root = Test_RBTree_ins(key, root);
                            let expected = std::collections::BTreeSet::from([5, 10, 15, key]);
                            b13_check(&root, &expected, false);
                            assert!(matches!(root.as_ref(), Tree::T(Color::R, _, 10, _)));
                            if !parent_shared && !parent_weak {
                                unique_cases += 1;
                                unique_same_address += usize::from(address == std::rc::Rc::as_ptr(&root));
                            }
                            if let Some(ref old) = old { assert_eq!(b13_shape(old), root_shape); }
                            if let Some(ref old) = left_old { assert_eq!(b13_shape(old), left_shape); }
                            if let Some(ref old) = right_old { assert_eq!(b13_shape(old), right_shape); }
                            for (weak, shape) in [(&weak, &root_shape), (&left_weak, &left_shape), (&right_weak, &right_shape)] {
                                if let Some(node) = weak.as_ref().and_then(|w| w.upgrade()) {
                                    assert_eq!(&b13_shape(&node), shape, "weak observed mutation");
                                }
                            }
                            let mut all_weak = Vec::new();
                            b13_collect_weaks(&root, &mut all_weak);
                            if let Some(ref node) = old { b13_collect_weaks(node, &mut all_weak); }
                            if let Some(ref node) = left_old { b13_collect_weaks(node, &mut all_weak); }
                            if let Some(ref node) = right_old { b13_collect_weaks(node, &mut all_weak); }
                            drop(root); drop(old); drop(left_old); drop(right_old);
                            assert!(all_weak.iter().all(|w| w.upgrade().is_none()), "node retained after final drop");
                            for weak in [&weak, &left_weak, &right_weak] {
                                assert!(weak.as_ref().is_none_or(|w| w.upgrade().is_none()));
                            }
                            cases += 1;
                        }
                    }
                }
            }
        }
    }
    println!("B13 red-parent cases: {cases}; unique parent address retained: {unique_same_address}/{unique_cases} (observation only)");
}

pub fn extra_checks() {
    b13_red_parent_cases();
    // Retain exactly 200 previous versions; compare each against BTreeSet after
    // every operation. Include the i64 boundaries without sentinel assumptions.
    let mut tree = std::rc::Rc::new(Tree::E);
    let mut expected = std::collections::BTreeSet::new();
    let mut versions = Vec::new();
    for n in 0..200 {
        versions.push((tree.clone(), expected.clone(), b13_shape(&tree)));
        let key = match n { 0 => i64::MIN, 1 => i64::MAX, _ => (n * 73 % 197) - 100 };
        tree = Test_RBTree_insert(key, tree);
        expected.insert(key);
        b13_check(&tree, &expected, true);
        for (old, old_expected, shape) in &versions {
            b13_check(old, old_expected, true);
            assert_eq!(&b13_shape(old), shape);
        }
    }
    let mut all_weak = Vec::new();
    b13_collect_weaks(&tree, &mut all_weak);
    for (old, _, _) in &versions { b13_collect_weaks(old, &mut all_weak); }
    drop(tree); drop(versions);
    assert!(all_weak.iter().all(|w| w.upgrade().is_none()));

    // Predominantly unique roots, periodically retained historical versions,
    // independent child references, root weak handles, and duplicate insertions.
    let mut tree = std::rc::Rc::new(Tree::E);
    let mut expected = std::collections::BTreeSet::new();
    let mut versions = Vec::new();
    let mut children = Vec::new();
    let mut state = 0xB13_u64;
    for n in 0..512 {
        if n % 17 == 0 { versions.push((tree.clone(), expected.clone(), b13_shape(&tree))); }
        if n % 7 == 0 {
            if let Tree::T(_, left, _, right) = tree.as_ref() {
                let child = if n % 2 == 0 { left } else { right };
                children.push((child.clone(), b13_shape(child)));
            }
        }
        let weak = (n % 23 == 0).then(|| (std::rc::Rc::downgrade(&tree), b13_shape(&tree)));
        state = state.wrapping_mul(6364136223846793005).wrapping_add(1);
        let key = ((state >> 32) % 173) as i64 - 86;
        tree = Test_RBTree_insert(key, tree);
        expected.insert(key);
        b13_check(&tree, &expected, true);
        for (old, old_expected, shape) in &versions {
            b13_check(old, old_expected, true);
            assert_eq!(&b13_shape(old), shape);
        }
        for (child, shape) in &children { assert_eq!(&b13_shape(child), shape); }
        if let Some((weak, shape)) = weak {
            if let Some(old) = weak.upgrade() { assert_eq!(b13_shape(&old), shape); }
        }
    }
    let mut all_weak = Vec::new();
    b13_collect_weaks(&tree, &mut all_weak);
    for (old, _, _) in &versions { b13_collect_weaks(old, &mut all_weak); }
    for (child, _) in &children { b13_collect_weaks(child, &mut all_weak); }
    drop(tree); drop(versions); drop(children);
    assert!(all_weak.iter().all(|w| w.upgrade().is_none()));
    println!("B13 extra checks passed: BTreeSet, RB invariants, 200 persistent versions, 512 mixed-sharing insertions, independent children, weak values and final lifetimes");
}
fn depth_checks() {
    assert_eq!(Test_RBTree_depth(std::rc::Rc::new(Tree::E)), 0);
    // The public owned boundary must release its owner before it returns.
    let tree = Test_RBTree_buildTree(100000, std::rc::Rc::new(Tree::E));
    let weak = std::rc::Rc::downgrade(&tree);
    assert_eq!(Test_RBTree_depth(tree), 22);
    assert!(weak.upgrade().is_none(), "Unique owner is destroyed inside depth");

    let tree = Test_RBTree_buildTree(1000, std::rc::Rc::new(Tree::E));
    let before = b13_shape(&tree);
    let depth = Test_RBTree_depth(tree.clone());
    let weak = std::rc::Rc::downgrade(&tree);
    for _ in 0..20 {
        assert_eq!(Test_RBTree_depth(tree.clone()), depth);
        assert_eq!(b13_shape(&tree), before);
        assert_eq!(std::rc::Rc::strong_count(&tree), 1, "A read must not retain its root");
        assert_eq!(b13_shape(&weak.upgrade().unwrap()), before);
    }
    let changed = Test_RBTree_insert(1001, tree.clone());
    assert_eq!(b13_shape(&tree), before);
    assert!(Test_RBTree_depth(changed) >= depth);
    drop(tree);
    assert!(weak.upgrade().is_none());

    // Repeated edges and an independently retained child are legal read-only
    // inputs even when the artificial shape is not an ordered red-black tree.
    for root_shared in [false, true] {
        for child_shared in [false, true] {
            let child = b13_black_leaf(7);
            let saved_child = child_shared.then(|| child.clone());
            let child_weak = std::rc::Rc::downgrade(&child);
            let root = std::rc::Rc::new(Tree::T(Color::B, child.clone(), 10, child));
            let saved_root = root_shared.then(|| root.clone());
            let root_weak = std::rc::Rc::downgrade(&root);
            assert_eq!(Test_RBTree_depth(root), 2);
            assert_eq!(root_weak.upgrade().is_some(), root_shared);
            assert_eq!(child_weak.upgrade().is_some(), root_shared || child_shared);
            if let Some(ref root) = saved_root { assert_eq!(Test_RBTree_depth(root.clone()), 2); }
            if let Some(ref child) = saved_child { assert_eq!(Test_RBTree_depth(child.clone()), 1); }
            drop(saved_root); drop(saved_child);
            assert!(root_weak.upgrade().is_none());
            assert!(child_weak.upgrade().is_none());
        }
    }
    println!("Depth checks passed: 100k tree, owned destruction before return, repeated reads, later insertion, shared edges and independent child/Weak lifetimes");
}

fn main() { original_checks(); extra_checks(); depth_checks(); }
