#![allow(warnings)]
use purust_core::*;
use Purs_Data_Unit::Data_Unit_unit;
#[global_allocator] static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;
pub fn Test_LazyEvaluation_buildThunks(mut n: i64, mut acc: purust_core::Func1<crate::UnknownType, i64>, unit: crate::UnknownType) -> i64 {
    while n != 0 {
        let previous = acc;
        acc = purust_core::Func1::Shared(std::rc::Rc::new(move |_| previous(Data_Unit_unit()) + 1));
        n -= 1;
    }
    acc(unit)
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
