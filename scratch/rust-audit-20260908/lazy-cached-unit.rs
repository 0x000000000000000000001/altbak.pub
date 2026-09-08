#![allow(warnings)]
use purust_core::*;

#[global_allocator] static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;

thread_local! { static AUDIT_UNIT: UnknownType = Purs_Data_Unit::Data_Unit_unit(); }
fn Data_Unit_unit() -> UnknownType { AUDIT_UNIT.with(Clone::clone) }
pub fn Test_LazyEvaluation_buildThunks(mut a0: i64, mut a1: purust_core::Func1<crate::UnknownType, i64>, mut a2: crate::UnknownType) -> i64 {
    // AST: Typed(Typed(Abs(..., Typed(Abs(..., Typed(Branch(...)))))))
    loop {
        break (/* Typed Abs */purust_core::Func3::Shared(std::rc::Rc::new({ let _f = (purust_core::Func1::Static(|mut _a0: i64| -> purust_core::Func2<purust_core::Func1<crate::UnknownType, i64>, crate::UnknownType, i64> {
    let mut purs_local_0 = _a0;
    /* Typed Abs */purust_core::Func2::Shared(std::rc::Rc::new({ let _f = (purust_core::Func1::Shared(std::rc::Rc::new(move |mut _a0: purust_core::Func1<crate::UnknownType, i64>| -> purust_core::Func1<crate::UnknownType, i64> {
    let mut purs_local_1 = _a0;
    /* Typed purust_core::Func1<crate::UnknownType, i64> <- purust_core::Func1<crate::UnknownType, i64> : Branch(...) */if (/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == 0) {
        /* Typed purust_core::Func1<crate::UnknownType, i64> <- purust_core::Func1<crate::UnknownType, i64> : Local(...) */purs_local_1
    } else {
        /* Typed purust_core::Func1<crate::UnknownType, i64> <- purust_core::Func1<crate::UnknownType, i64> : App(Var(...)) */{
        let mut eval_arg_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let mut eval_arg_1 = /* Typed Abs */purust_core::Func1::Shared(std::rc::Rc::new(move |mut _a0: crate::UnknownType| -> i64 {
    drop(_a0);
    /* Typed i64 <- i64 : PrimOp(...) */((purs_local_1.clone())(/* Typed crate::UnknownType <- crate::UnknownType : Var(...) */Data_Unit_unit()) + 1)
}));
    purust_core::Func1::Shared(std::rc::Rc::new(move |mut eta_0: crate::UnknownType| -> i64 {
    let mut eval_arg_0 = eval_arg_0.clone();
    let mut eval_arg_1 = eval_arg_1.clone();
    Test_LazyEvaluation_buildThunks(eval_arg_0, eval_arg_1, eta_0.clone())
}))
}
    }
}))).clone(); move |mut _a0: purust_core::Func1<crate::UnknownType, i64>, mut _a1: crate::UnknownType| -> i64 { ((_f)(_a0.clone()))(_a1.clone()) } }))
} as fn(i64) -> purust_core::Func2<purust_core::Func1<crate::UnknownType, i64>, crate::UnknownType, i64>)).clone(); move |mut _a0: i64, mut _a1: purust_core::Func1<crate::UnknownType, i64>, mut _a2: crate::UnknownType| -> i64 { ((_f)(_a0.clone()))(_a1.clone(), _a2.clone()) } })))(a0.clone(), a1.clone(), a2.clone());
    }
}

pub fn Test_LazyEvaluation_runManyTimes(mut purs_local_0: i64, mut purs_local_1: i64) -> i64 {
    // AST: Typed(Typed(Abs(..., Typed(Abs(..., Typed(Branch(...)))))))
    loop {
        break /* Typed i64 <- i64 : Branch(...) */if (/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == 0) {
        /* Typed i64 <- i64 : Local(...) */purs_local_1
    } else {
        /* Typed i64 <- i64 : App(Var(...)) */{
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0 - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed i64 <- i64 : PrimOp(...) */(purs_local_1 + Test_LazyEvaluation_buildThunks(/* Typed i64 <- i64 : Lit */1000, /* Typed Abs */purust_core::Func1::Static(|mut _a0: crate::UnknownType| -> i64 {
    drop(_a0);
    /* Typed i64 <- i64 : Lit */0
} as fn(crate::UnknownType) -> i64), /* Typed crate::UnknownType <- crate::UnknownType : Var(...) */Data_Unit_unit()));
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        continue;
    }
    };
    }
}


fn main() {
    eprintln!("layout: Value={} Record_a={}", std::mem::size_of::<Value>(), std::mem::size_of::<Record_a>());
    for n in [0,1,10,1000] {
        let calls = std::rc::Rc::new(std::cell::Cell::new(0));
        let counter = calls.clone();
        let f = purust_core::Func1::Shared(std::rc::Rc::new(move |_| { counter.set(counter.get()+1); 7 }));
        assert_eq!(Test_LazyEvaluation_buildThunks(n, f, Data_Unit_unit()), n+7);
        assert_eq!(calls.get(), 1);
    }
    assert_eq!(Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000),0),1000000);
    for _ in 0..3 {
        let start = std::time::Instant::now();
        let value = std::hint::black_box(Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000),0));
        let us = start.elapsed().as_secs_f64()*1e6;
        assert_eq!(value,1000000);
        println!("{:.3}",us);
    }
}
