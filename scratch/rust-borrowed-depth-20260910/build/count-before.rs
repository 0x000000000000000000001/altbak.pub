#![allow(warnings)]
#[path="/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rust-perceus-counts-20260909/tracked_rc.rs"] mod tracked;











































#[derive(Clone, Copy)]
pub enum Color {
    R,
    B
}

#[derive(Clone)]
pub enum Tree {
    E,
    T(crate::Color, tracked::Rc<crate::Tree>, i64, tracked::Rc<crate::Tree>)
}
impl Tree { pub fn __purust_take(&mut self) -> std::option::Option<Self> { std::option::Option::Some(std::mem::replace(self, Self::E)) } }


fn Test_RBTree___purust_rebuild_E(mut __purust_cell: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> {
let payload = crate::Tree::E;
if let std::option::Option::Some(slot) = tracked::Rc::get_mut(&mut __purust_cell) { *slot = payload; __purust_cell } else { tracked::Rc::new(payload) }
}
fn Test_RBTree___purust_rebuild_T(a0: crate::Color, a1: tracked::Rc<crate::Tree>, a2: i64, a3: tracked::Rc<crate::Tree>, mut __purust_cell: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> {
let payload = crate::Tree::T(a0, a1, a2, a3);
if let std::option::Option::Some(slot) = tracked::Rc::get_mut(&mut __purust_cell) { *slot = payload; __purust_cell } else { tracked::Rc::new(payload) }
}
#[inline]
fn Test_RBTree_balance__purust_child_rebuilds(_purust_guard_arg_0: &crate::Color, _purust_guard_arg_1: &tracked::Rc<crate::Tree>, _purust_guard_arg_2: &i64, _purust_guard_arg_3: &tracked::Rc<crate::Tree>) -> bool { !(matches!(_purust_guard_arg_0, crate::Color::B)) || (if matches!((_purust_guard_arg_1).as_ref(), crate::Tree::T(..)) { if matches!((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R) { if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) })) } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) }) } else { if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) })) } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) } } } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) } } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) }) }




















pub fn Test_RBTree_max(mut purs_local_0: i64, mut purs_local_1: i64) -> i64 {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
/* Typed i64 <- i64 : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Local(...) */purs_local_1.clone()) {
        /* Typed i64 <- i64 : Local(...) */purs_local_0
    } else {
        /* Typed i64 <- i64 : Local(...) */purs_local_1
    }
}

pub fn Test_RBTree_makeBlack(mut purs_local_0: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Branch(...))))
/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::T(..)) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */{ let _new_field = crate::Color::B; let mut purs_local_0 = purs_local_0; if let std::option::Option::Some(crate::Tree::T(_updated_field, _, _, _)) = tracked::Rc::get_mut(&mut purs_local_0) { *_updated_field = _new_field; purs_local_0 } else { let _rebuilt = crate::Tree::T(crate::Color::B, { if let crate::Tree::T(_, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } }); let mut _reused = purs_local_0; if let std::option::Option::Some(_slot) = tracked::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { tracked::Rc::new(_rebuilt) } } }
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::E) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}






pub fn Test_RBTree_depth(mut purs_local_0: tracked::Rc<crate::Tree>) -> i64 {
    // AST: Typed(Abs(..., Typed(Branch(...))))
    loop {
        break /* Typed i64 <- i64 : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::E) {
        /* Typed i64 <- i64 : Lit */0
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::T(..)) {
        /* Typed i64 <- i64 : PrimOp(...) */(1 + {
    let mut purs_local_1 = Test_RBTree_depth(/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } });
    /* Typed i64 <- i64 : Let(...) */{
    let mut purs_local_2 = /* Typed i64 <- i64 : App(Var(...)) */Test_RBTree_depth(/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_0).as_ref() { f.clone() } else { unreachable!() } });
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

pub fn Test_RBTree_balance(mut purs_local_0: crate::Color, mut purs_local_1: tracked::Rc<crate::Tree>, mut purs_local_2: i64, mut purs_local_3: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))))))
/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Local(...) */purs_local_0.clone()), crate::Color::B) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => { let mut _owned_purs_local_1_1 = _owned_purs_local_1_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_1_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_1_0, mut _owned__owned_purs_local_1_1_1, mut _owned__owned_purs_local_1_1_2, mut _owned__owned_purs_local_1_1_3)) => {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_1_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_1_3;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1)))
}
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1;
    { let mut _owned_purs_local_1_3 = _owned_purs_local_1_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_1_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_3_0, mut _owned__owned_purs_local_1_3_1, mut _owned__owned_purs_local_1_3_2, mut _owned__owned_purs_local_1_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1;
    { let mut _owned_purs_local_1_3 = _owned_purs_local_1_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_1_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_3_0, mut _owned__owned_purs_local_1_3_1, mut _owned__owned_purs_local_1_3_2, mut _owned__owned_purs_local_1_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1)))
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3)))
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(purs_local_6, purs_local_4, purs_local_7, purs_local_5))
}
}
}
}
    }
}

pub fn Test_RBTree_ins(mut purs_local_0: i64, mut purs_local_1: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
    loop {
        break /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::E) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1.clone(), /* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1))
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() < /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */if !(/* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::B)) { { /* purust child call: retained fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_child_slot) = tracked::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_child_field_0, _purust_child_field_1, _purust_child_field_2, _purust_child_field_3) = _purust_child_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_child_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_child_field_2).clone(); let mut _owned_purs_local_1_1 = std::mem::replace(_purust_child_field_1, (*_purust_child_field_3).clone()); let _purust_new_child = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1); *_purust_child_field_1 = _purust_new_child; purs_local_1 } else { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = tracked::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3) }, _ => unreachable!() } } } } else { { /* purust child call: post-call fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_post_slot) = tracked::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) = _purust_post_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_post_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_post_field_2).clone(); let mut _owned_purs_local_1_1 = std::mem::replace(_purust_post_field_1, (*_purust_post_field_3).clone()); let _purust_post_child = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1); *_purust_post_field_1 = _purust_post_child; if Test_RBTree_balance__purust_child_rebuilds(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) { purs_local_1 } else { let std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) = _purust_post_slot.__purust_take() else { unreachable!() }; Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1) } } else { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = tracked::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3) }, _ => unreachable!() } } } }
    } else {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Accessor(Local(...)) */{ if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */if !(/* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::B)) { { /* purust child call: retained fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_child_slot) = tracked::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_child_field_0, _purust_child_field_1, _purust_child_field_2, _purust_child_field_3) = _purust_child_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_child_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_child_field_2).clone(); let mut _owned_purs_local_1_3 = std::mem::replace(_purust_child_field_3, (*_purust_child_field_1).clone()); let _purust_new_child = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *_purust_child_field_3 = _purust_new_child; purs_local_1 } else { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = tracked::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)) }, _ => unreachable!() } } } } else { { /* purust child call: post-call fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_post_slot) = tracked::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) = _purust_post_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_post_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_post_field_2).clone(); let mut _owned_purs_local_1_3 = std::mem::replace(_purust_post_field_3, (*_purust_post_field_1).clone()); let _purust_post_child = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *_purust_post_field_3 = _purust_post_child; if Test_RBTree_balance__purust_child_rebuilds(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) { purs_local_1 } else { let std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) = _purust_post_slot.__purust_take() else { unreachable!() }; Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1) } } else { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = tracked::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)) }, _ => unreachable!() } } } }
    } else {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */{ let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => { let _rebuilt = crate::Tree::T(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *tracked::Rc::get_mut(&mut purs_local_1).unwrap() = _rebuilt; purs_local_1 }, std::option::Option::None => { let _rebuilt = crate::Tree::T({ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }); let mut _reused = purs_local_1; if let std::option::Option::Some(_slot) = tracked::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { tracked::Rc::new(_rebuilt) } }, _ => unreachable!() } }
    }
    }
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    };
    }
}

pub fn Test_RBTree_insert(mut purs_local_0: i64, mut purs_local_1: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Let(...))))))
/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Let(...) */{
    let mut purs_local_2 = Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1);
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((purs_local_2).as_ref(), crate::Tree::T(..)) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */{ let _new_field = crate::Color::B; let mut purs_local_2 = purs_local_2; if let std::option::Option::Some(crate::Tree::T(_updated_field, _, _, _)) = tracked::Rc::get_mut(&mut purs_local_2) { *_updated_field = _new_field; purs_local_2 } else { let _rebuilt = crate::Tree::T(crate::Color::B, { if let crate::Tree::T(_, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }, { if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_2).as_ref() { f.clone() } else { unreachable!() } }); let mut _reused = purs_local_2; if let std::option::Option::Some(_slot) = tracked::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { tracked::Rc::new(_rebuilt) } } }
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((purs_local_2).as_ref(), crate::Tree::E) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_2
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}
}

pub fn Test_RBTree_buildTree(mut purs_local_0: i64, mut purs_local_1: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
    loop {
        break /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if (/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == 0) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1
    } else {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */{
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_insert(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1);
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        continue;
    }
    };
    }
}





























fn Test_RBTree_balance__purust_reuse(mut purs_local_0: crate::Color, mut purs_local_1: tracked::Rc<crate::Tree>, mut purs_local_2: i64, mut purs_local_3: tracked::Rc<crate::Tree>, mut __purust_cell: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> {
    // AST: Typed(Abs(..., Typed(Branch(...))))
/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Local(...) */purs_local_0.clone()), crate::Color::B) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => { let mut _owned_purs_local_1_1 = _owned_purs_local_1_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_1_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_1_0, mut _owned__owned_purs_local_1_1_1, mut _owned__owned_purs_local_1_1_2, mut _owned__owned_purs_local_1_1_3)) => {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_1_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_1_3;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_4 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_5 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1;
    { let mut _owned_purs_local_1_3 = _owned_purs_local_1_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_1_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_3_0, mut _owned__owned_purs_local_1_3_1, mut _owned__owned_purs_local_1_3_2, mut _owned__owned_purs_local_1_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1;
    { let mut _owned_purs_local_1_3 = _owned_purs_local_1_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_1_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_1_3_0, mut _owned__owned_purs_local_1_3_1, mut _owned__owned_purs_local_1_3_2, mut _owned__owned_purs_local_1_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_1_3_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_1_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { if let crate::Tree::T(_, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_1).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}
}, _ => unreachable!() } }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => { let mut _owned_purs_local_3_1 = _owned_purs_local_3_1; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_1_0, mut _owned__owned_purs_local_3_1_1, mut _owned__owned_purs_local_3_1_2, mut _owned__owned_purs_local_3_1_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_1;
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_1_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_5 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_6 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ if let crate::Tree::T(ref f, ..) = ({ if let crate::Tree::T(_, _, _, ref f, ..) = (purs_local_3).as_ref() { f.clone() } else { unreachable!() } }).as_ref() { f.clone() } else { unreachable!() } }), crate::Color::R)) {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    { let mut purs_local_3 = purs_local_3; let _taken = tracked::Rc::get_mut(&mut purs_local_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_3_0, mut _owned_purs_local_3_1, mut _owned_purs_local_3_2, mut _owned_purs_local_3_3)) => {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1;
    { let mut _owned_purs_local_3_3 = _owned_purs_local_3_3; let _taken = tracked::Rc::get_mut(&mut _owned_purs_local_3_3).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned__owned_purs_local_3_3_0, mut _owned__owned_purs_local_3_3_1, mut _owned__owned_purs_local_3_3_2, mut _owned__owned_purs_local_3_3_3)) => {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_1;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned__owned_purs_local_3_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned__owned_purs_local_3_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, std::option::Option::None => {
    let mut purs_local_6 = { if let crate::Tree::T(_, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_7 = { if let crate::Tree::T(_, _, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { if let crate::Tree::T(_, _, ref f, ..) = (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { f.clone() } else { unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
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
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
    } else {
        {
    let mut purs_local_4 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1;
    {
    let mut purs_local_5 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_6 = /* Typed crate::Color <- crate::Color : Local(...) */purs_local_0;
    {
    let mut purs_local_7 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(purs_local_6, purs_local_4, purs_local_7, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
    }
}





impl tracked::Kind for Tree {
    fn kind(&self) -> &'static str { match self { Tree::E => "empty", Tree::T(..) => "node" } }
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
fn main() {
    assert_eq!(std::mem::size_of::<tracked::Rc<Tree>>(), std::mem::size_of::<std::rc::Rc<Tree>>());
    assert_eq!(std::mem::size_of::<Tree>(), 32);
    for order in [[3,2,1], [3,1,2], [1,3,2], [1,2,3]] {
        let mut tree = tracked::Rc::new(Tree::E);
        for (i, key) in order.into_iter().enumerate() {
            tree = Test_RBTree_insert(key, tree); check(&tree, i+1);
        }
    }
    // Verify get_mut's weak-reference exclusion and unwrap's distinct rule.
    let mut tree = tracked::Rc::new(Tree::T(Color::B,
        tracked::Rc::new(Tree::E), 1, tracked::Rc::new(Tree::E)));
    let weak = tracked::Rc::downgrade(&tree);
    assert!(tracked::Rc::get_mut(&mut tree).is_none());
    tree = Test_RBTree_insert(2, tree); check(&tree, 2);
    assert!(weak.upgrade().is_none()); drop(tree); drop(weak);

    tracked::phase("unique_build");
    let tree = Test_RBTree_buildTree(std::hint::black_box(100000), tracked::Rc::new(Tree::E));
    check(&tree, 100000);
    tracked::phase("unique_depth_and_drop");
    assert_eq!(Test_RBTree_depth(tree), 22);

    tracked::phase("persistent_build");
    let mut tree = tracked::Rc::new(Tree::E);
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
    tracked::phase("persistent_depth");
    assert!(Test_RBTree_depth(tree.clone()) > 0);
    for (old, expected) in &versions { assert_eq!(&keys(old), expected); }
    tracked::phase("persistent_drop");
    drop(tree); drop(versions);
    tracked::phase("done");
    tracked::print_events();
}
