











































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


fn Test_RBTree___purust_rebuild_E(mut __purust_cell: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> { tracked::call("__purust_rebuild_E");
let payload = crate::Tree::E;
if let std::option::Option::Some(slot) = tracked::Rc::get_mut(&mut __purust_cell) { *slot = payload; __purust_cell } else { tracked::Rc::new(payload) }
}
fn Test_RBTree___purust_rebuild_T(a0: crate::Color, a1: tracked::Rc<crate::Tree>, a2: i64, a3: tracked::Rc<crate::Tree>, mut __purust_cell: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> { tracked::call("__purust_rebuild_T");
let payload = crate::Tree::T(a0, a1, a2, a3);
if let std::option::Option::Some(slot) = tracked::Rc::get_mut(&mut __purust_cell) { *slot = payload; __purust_cell } else { tracked::Rc::new(payload) }
}
#[inline]
fn Test_RBTree_balance__purust_child_rebuilds(_purust_guard_arg_0: &crate::Color, _purust_guard_arg_1: &tracked::Rc<crate::Tree>, _purust_guard_arg_2: &i64, _purust_guard_arg_3: &tracked::Rc<crate::Tree>) -> bool { tracked::call("balance__purust_child_rebuilds"); !(matches!(_purust_guard_arg_0, crate::Color::B)) || (if matches!((_purust_guard_arg_1).as_ref(), crate::Tree::T(..)) { if matches!((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R) { if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) })) } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) }) } else { if matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) })) } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) } } } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) } } else { !((matches!((_purust_guard_arg_3).as_ref(), crate::Tree::T(..))) && (matches!((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) || (if matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)) { !(matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) && (!((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)))) } else { !((matches!(((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..))) && (matches!((match ((match (_purust_guard_arg_3).as_ref() { crate::Tree::T(_, _, _, _purust_guard_field) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) }) }) }
#[inline]
fn Test_RBTree_balance__purust_permute_fields__matches(_purust_guard_arg_0: &crate::Color, _purust_guard_arg_1: &tracked::Rc<crate::Tree>, _purust_guard_arg_2: &i64, _purust_guard_arg_3: &tracked::Rc<crate::Tree>) -> bool { tracked::call("balance__purust_permute_fields__matches"); ((((matches!(_purust_guard_arg_0, crate::Color::B)) && (matches!((_purust_guard_arg_1).as_ref(), crate::Tree::T(..)))) && (matches!((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R))) && (matches!(((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref(), crate::Tree::T(..)))) && (matches!((match ((match (_purust_guard_arg_1).as_ref() { crate::Tree::T(_, _purust_guard_field, _, _) => _purust_guard_field, _ => unreachable!() })).as_ref() { crate::Tree::T(_purust_guard_field, _, _, _) => _purust_guard_field, _ => unreachable!() }), crate::Color::R)) }
#[inline]
fn Test_RBTree_balance__purust_permute_fields(_purust_permute_slot: &mut crate::Tree) -> bool { tracked::call("balance__purust_permute_fields");
// purust field permutation: typed fields, exclusive cells, staged scalars
let crate::Tree::T(_purust_permute_0_0, _purust_permute_0_1, _purust_permute_0_2, _purust_permute_0_3) = _purust_permute_slot else { return false; };
if !Test_RBTree_balance__purust_permute_fields__matches(_purust_permute_0_0, _purust_permute_0_1, _purust_permute_0_2, _purust_permute_0_3) { return false; }
{ let std::option::Option::Some(_purust_permute_parent) = tracked::Rc::get_mut(_purust_permute_0_1) else { return false; };
let crate::Tree::T(_purust_permute_1_0, _purust_permute_1_1, _purust_permute_1_2, _purust_permute_1_3) = _purust_permute_parent else { return false; };
{ let std::option::Option::Some(_purust_permute_child) = tracked::Rc::get_mut(_purust_permute_1_1) else { return false; };
let crate::Tree::T(_purust_permute_2_0, _purust_permute_2_1, _purust_permute_2_2, _purust_permute_2_3) = _purust_permute_child else { return false; };
let _purust_permute_value_0 = crate::Color::R; let _purust_permute_value_1 = *_purust_permute_1_2; let _purust_permute_value_2 = crate::Color::B; let _purust_permute_value_3 = *_purust_permute_0_2; let _purust_permute_value_4 = crate::Color::B;
*_purust_permute_0_0 = _purust_permute_value_0; *_purust_permute_0_2 = _purust_permute_value_1; *_purust_permute_1_0 = _purust_permute_value_2; *_purust_permute_1_2 = _purust_permute_value_3; *_purust_permute_2_0 = _purust_permute_value_4;
}
std::mem::swap(_purust_permute_1_1, _purust_permute_1_3);
std::mem::swap(_purust_permute_1_3, _purust_permute_0_3);
}
std::mem::swap(_purust_permute_0_1, _purust_permute_0_3);
true
}




















pub fn Test_RBTree_max(mut purs_local_0: i64, mut purs_local_1: i64) -> i64 { tracked::call("max");
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
/* Typed i64 <- i64 : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Local(...) */purs_local_1.clone()) {
        /* Typed i64 <- i64 : Local(...) */purs_local_0
    } else {
        /* Typed i64 <- i64 : Local(...) */purs_local_1
    }
}

pub fn Test_RBTree_makeBlack(mut purs_local_0: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> { tracked::call("makeBlack");
    // AST: Typed(Abs(..., Typed(Branch(...))))
/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::T(..)) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */{ let _new_field = crate::Color::B; let mut purs_local_0 = purs_local_0; if let std::option::Option::Some(crate::Tree::T(_updated_field, _, _, _)) = tracked::Rc::get_mut(&mut purs_local_0) { *_updated_field = _new_field; purs_local_0 } else { let _rebuilt = crate::Tree::T(crate::Color::B, { match (purs_local_0).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }, { match (purs_local_0).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } }, { match (purs_local_0).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }); let mut _reused = purs_local_0; if let std::option::Option::Some(_slot) = tracked::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { tracked::Rc::new(_rebuilt) } } }
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::E) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}









pub fn Test_RBTree_depth(mut purs_local_0: tracked::Rc<crate::Tree>) -> i64 { tracked::call("depth");
    // AST: Typed(Abs(..., Typed(Branch(...))))
    loop {
        break /* Typed i64 <- i64 : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::E) {
        /* Typed i64 <- i64 : Lit */0
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_0).as_ref(), crate::Tree::T(..)) {
        /* Typed i64 <- i64 : PrimOp(...) */(1 + {
    let mut purs_local_1 = Test_RBTree_depth(/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Accessor(Local(...)) */{ match (purs_local_0).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } });
    {
    let mut purs_local_2 = /* Typed i64 <- i64 : App(Var(...)) */Test_RBTree_depth(/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Accessor(Local(...)) */{ match (purs_local_0).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } });
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

pub fn Test_RBTree_balance(mut purs_local_0: crate::Color, mut purs_local_1: tracked::Rc<crate::Tree>, mut purs_local_2: i64, mut purs_local_3: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> { tracked::call("balance");
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))))))
/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Local(...) */purs_local_0.clone()), crate::Color::B) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_1).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_4 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    let mut purs_local_4 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_5 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_9 = { match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    let mut purs_local_4 = { match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_5 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_9 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    let mut purs_local_4 = { match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_5 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_9 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7))))
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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

pub fn Test_RBTree_ins(mut purs_local_0: i64, mut purs_local_1: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> { tracked::call("ins");
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
    loop {
        break /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::E) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */tracked::Rc::new(crate::Tree::T(crate::Color::R, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1.clone(), /* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1))
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() < /* Typed i64 <- i64 : Accessor(Local(...)) */{ match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } }) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */if !(/* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ match (purs_local_1).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::B)) { { /* purust child call: retained fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_child_slot) = tracked::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_child_field_0, _purust_child_field_1, _purust_child_field_2, _purust_child_field_3) = _purust_child_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_child_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_child_field_2).clone(); let mut _owned_purs_local_1_1 = std::mem::replace(_purust_child_field_1, (*_purust_child_field_3).clone()); let _purust_new_child = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1); *_purust_child_field_1 = _purust_new_child; purs_local_1 } else { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = tracked::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3) }, _ => unreachable!() } } } } else { { /* purust child call: post-call fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_post_slot) = tracked::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) = _purust_post_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_post_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_post_field_2).clone(); let mut _owned_purs_local_1_1 = std::mem::replace(_purust_post_field_1, (*_purust_post_field_3).clone()); let _purust_post_child = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1); *_purust_post_field_1 = _purust_post_child; if Test_RBTree_balance__purust_child_rebuilds(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) { purs_local_1 } else if Test_RBTree_balance__purust_permute_fields(_purust_post_slot) { purs_local_1 } else { let std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) = _purust_post_slot.__purust_take() else { unreachable!() }; Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1) } } else { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = tracked::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1), /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3) }, _ => unreachable!() } } } }
    } else {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() > /* Typed i64 <- i64 : Accessor(Local(...)) */{ match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } }) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */if !(/* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Accessor(Local(...)) */{ match (purs_local_1).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::B)) { { /* purust child call: retained fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_child_slot) = tracked::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_child_field_0, _purust_child_field_1, _purust_child_field_2, _purust_child_field_3) = _purust_child_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_child_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_child_field_2).clone(); let mut _owned_purs_local_1_3 = std::mem::replace(_purust_child_field_3, (*_purust_child_field_1).clone()); let _purust_new_child = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *_purust_child_field_3 = _purust_new_child; purs_local_1 } else { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = tracked::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)) }, _ => unreachable!() } } } } else { { /* purust child call: post-call fields */ let mut purs_local_1 = purs_local_1; if let std::option::Option::Some(_purust_post_slot) = tracked::Rc::get_mut(&mut purs_local_1) { let crate::Tree::T(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) = _purust_post_slot else { unreachable!() }; let mut _owned_purs_local_1_0 = (*_purust_post_field_0).clone(); let mut _owned_purs_local_1_2 = (*_purust_post_field_2).clone(); let mut _owned_purs_local_1_3 = std::mem::replace(_purust_post_field_3, (*_purust_post_field_1).clone()); let _purust_post_child = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *_purust_post_field_3 = _purust_post_child; if Test_RBTree_balance__purust_child_rebuilds(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) { purs_local_1 } else if Test_RBTree_balance__purust_permute_fields(_purust_post_slot) { purs_local_1 } else { let std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) = _purust_post_slot.__purust_take() else { unreachable!() }; Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3, purs_local_1) } } else { let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => Test_RBTree_balance__purust_reuse(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3), purs_local_1), std::option::Option::None => { let crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3) = tracked::Rc::unwrap_or_clone(purs_local_1) else { unreachable!() }; Test_RBTree_balance(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3)) }, _ => unreachable!() } } } }
    } else {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */{ let mut purs_local_1 = purs_local_1; let _taken = tracked::Rc::get_mut(&mut purs_local_1).and_then(|node| node.__purust_take()); match _taken { std::option::Option::Some(crate::Tree::T(mut _owned_purs_local_1_0, mut _owned_purs_local_1_1, mut _owned_purs_local_1_2, mut _owned_purs_local_1_3)) => { let _rebuilt = crate::Tree::T(/* Typed crate::Color <- crate::Color : Local(...) */_owned_purs_local_1_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1, /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3); *tracked::Rc::get_mut(&mut purs_local_1).unwrap() = _rebuilt; purs_local_1 }, std::option::Option::None => { let _rebuilt = crate::Tree::T({ match (purs_local_1).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }, { match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }, { match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } }, { match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }); let mut _reused = purs_local_1; if let std::option::Option::Some(_slot) = tracked::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { tracked::Rc::new(_rebuilt) } }, _ => unreachable!() } }
    }
    }
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    };
    }
}

pub fn Test_RBTree_insert(mut purs_local_0: i64, mut purs_local_1: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> { tracked::call("insert");
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Let(...))))))
{
    let mut purs_local_2 = Test_RBTree_ins(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1);
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((purs_local_2).as_ref(), crate::Tree::T(..)) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : CtorSaturated(...) */{ let _new_field = crate::Color::B; let mut purs_local_2 = purs_local_2; if let std::option::Option::Some(crate::Tree::T(_updated_field, _, _, _)) = tracked::Rc::get_mut(&mut purs_local_2) { *_updated_field = _new_field; purs_local_2 } else { let _rebuilt = crate::Tree::T(crate::Color::B, { match (purs_local_2).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }, { match (purs_local_2).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } }, { match (purs_local_2).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }); let mut _reused = purs_local_2; if let std::option::Option::Some(_slot) = tracked::Rc::get_mut(&mut _reused) { *_slot = _rebuilt; _reused } else { tracked::Rc::new(_rebuilt) } } }
    } else if /* OpIsTag Debug: Test_RBTree_E -> Other */ matches!((purs_local_2).as_ref(), crate::Tree::E) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_2
    } else {
        unimplemented!() /* Unsupported Expr: Fail */
    }
}
}

pub fn Test_RBTree_buildTree(mut purs_local_0: i64, mut purs_local_1: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> { tracked::call("buildTree");
    // AST: Typed(Abs(..., Typed(Abs(..., Typed(Branch(...))))))
    loop {
        break /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if (/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == 0) {
        /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1
    } else {
        {
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree_insert(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1);
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        continue;
    }
    };
    }
}
































fn Test_RBTree_balance__purust_reuse(mut purs_local_0: crate::Color, mut purs_local_1: tracked::Rc<crate::Tree>, mut purs_local_2: i64, mut purs_local_3: tracked::Rc<crate::Tree>, mut __purust_cell: tracked::Rc<crate::Tree>) -> tracked::Rc<crate::Tree> { tracked::call("balance__purust_reuse");
    // AST: Typed(Abs(..., Typed(Branch(...))))
/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Branch(...) */if /* OpIsTag Debug: Test_RBTree_B -> Other */ matches!((/* Typed crate::Color <- crate::Color : Local(...) */purs_local_0.clone()), crate::Color::B) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_1).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_1).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_4 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3;
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    let mut purs_local_4 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_5 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_9 = { match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    let mut purs_local_4 = { match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_5 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_9 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_1_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_1_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    let mut purs_local_4 = { match (purs_local_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_5 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3;
    {
    let mut purs_local_8 = { match (purs_local_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_9 = { match ({ match (purs_local_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!((/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
        if /* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) {
        if /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R) {
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
    let mut purs_local_5 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3;
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_1).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }, std::option::Option::None => {
    let mut purs_local_5 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5)), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}
}, _ => unreachable!() } }
}
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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
    } else if (/* OpIsTag Debug: Test_RBTree_T -> Func */ matches!(({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref(), crate::Tree::T(..)) && /* OpIsTag Debug: Test_RBTree_R -> Other */ matches!(({ match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(ref f, ..) => f.clone(), _ => unreachable!() } }), crate::Color::R)) {
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
    let mut purs_local_6 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = /* Typed i64 <- i64 : Local(...) */_owned_purs_local_3_2;
    {
    let mut purs_local_10 = { match (/* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */_owned_purs_local_3_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : App(Var(...)) */Test_RBTree___purust_rebuild_T(crate::Color::R, Test_RBTree___purust_rebuild_T(crate::Color::B, purs_local_4, purs_local_8, purs_local_5, /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */purs_local_3), purs_local_9, tracked::Rc::new(crate::Tree::T(crate::Color::B, purs_local_6, purs_local_10, purs_local_7)), /* Typed tracked::Rc<crate::Tree> <- tracked::Rc<crate::Tree> : Local(...) */__purust_cell)
}
}
}
}
}, _ => unreachable!() } }
}, std::option::Option::None => {
    let mut purs_local_5 = { match (purs_local_3).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_6 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_7 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_8 = /* Typed i64 <- i64 : Local(...) */purs_local_2;
    {
    let mut purs_local_9 = { match (purs_local_3).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
    {
    let mut purs_local_10 = { match ({ match (purs_local_3).as_ref() { crate::Tree::T(_, _, _, ref f, ..) => f.clone(), _ => unreachable!() } }).as_ref() { crate::Tree::T(_, _, ref f, ..) => f.clone(), _ => unreachable!() } };
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





