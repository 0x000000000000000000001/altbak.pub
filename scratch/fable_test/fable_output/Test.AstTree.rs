pub mod PureScript_Test_AstTree {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    #[derive(Clone, Debug,)]
    pub enum Test_AstTree_Expr {
        Test_AstTree_Valusd_Ctor(&dyn Any),
        Test_AstTree_Addusd_Ctor(&dyn Any, &dyn Any),
        Test_AstTree_Mulusd_Ctor(&dyn Any, &dyn Any),
        Test_AstTree_Subusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for PureScript_Test_AstTree::Test_AstTree_Expr {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Test_AstTree_Val() -> &dyn Any {
        static Test_AstTree_Val: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTree_Val.get_or_init(||
                                         &Func1::new(move |usd__arg1|
                                                         &LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Valusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Test_AstTree_Add() -> &dyn Any {
        static Test_AstTree_Add: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTree_Add.get_or_init(||
                                         &Func1::new(move |usd__arg1|
                                                         Func1::new({
                                                                        let usd__arg1
                                                                            =
                                                                            usd__arg1.clone();
                                                                        move
                                                                            |usd__arg2|
                                                                            &LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Addusd_Ctor(usd__arg1,
                                                                                                                                                              usd__arg2.clone()))
                                                                    })))
    }
    pub fn Test_AstTree_Mul() -> &dyn Any {
        static Test_AstTree_Mul: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTree_Mul.get_or_init(||
                                         &Func1::new(move |usd__arg1|
                                                         Func1::new({
                                                                        let usd__arg1
                                                                            =
                                                                            usd__arg1.clone();
                                                                        move
                                                                            |usd__arg2|
                                                                            &LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Mulusd_Ctor(usd__arg1,
                                                                                                                                                              usd__arg2.clone()))
                                                                    })))
    }
    pub fn Test_AstTree_Sub() -> &dyn Any {
        static Test_AstTree_Sub: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTree_Sub.get_or_init(||
                                         &Func1::new(move |usd__arg1|
                                                         Func1::new({
                                                                        let usd__arg1
                                                                            =
                                                                            usd__arg1.clone();
                                                                        move
                                                                            |usd__arg2|
                                                                            &LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Subusd_Ctor(usd__arg1,
                                                                                                                                                              usd__arg2.clone()))
                                                                    })))
    }
    pub fn Test_AstTree_eval_004022() -> &dyn Any {
        &Func1::new(move |v|
                        PureScript_Test_AstTree::Test_AstTree_eval_tco(v))
    }
    pub fn Test_AstTree_eval_004022_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_AstTree_eval_004022_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_AstTree_eval_004022_002d1.get_or_init(||
                                                       Lazy(Test_AstTree_eval_004022.clone()))
    }
    pub fn Test_AstTree_eval_tco(v: &dyn Any) -> &dyn Any {
        let matchValue: LrcPtr<PureScript_Test_AstTree::Test_AstTree_Expr> =
            Sharpurs_Prelude::unbox(v);
        match matchValue.as_ref() {
            PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Addusd_Ctor(matchValue_1_0,
                                                                                 matchValue_1_1)
            =>
            &(Sharpurs_Prelude::unbox(&&PureScript_Test_AstTree::Test_AstTree_eval_tco(&matchValue_1_0))
                  +
                  Sharpurs_Prelude::unbox(&&PureScript_Test_AstTree::Test_AstTree_eval_tco(&matchValue_1_1))),
            PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Mulusd_Ctor(matchValue_2_0,
                                                                                 matchValue_2_1)
            =>
            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                &&PureScript_Test_AstTree::Test_AstTree_eval_tco(&matchValue_2_0)),
                                             &&PureScript_Test_AstTree::Test_AstTree_eval_tco(&matchValue_2_1)),
            PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Subusd_Ctor(matchValue_3_0,
                                                                                 matchValue_3_1)
            =>
            &(Sharpurs_Prelude::unbox(&&PureScript_Test_AstTree::Test_AstTree_eval_tco(&matchValue_3_0))
                  -
                  Sharpurs_Prelude::unbox(&&PureScript_Test_AstTree::Test_AstTree_eval_tco(&matchValue_3_1))),
            PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Valusd_Ctor(matchValue_0_0)
            => matchValue_0_0,
        }
    }
    pub fn Test_AstTree_eval() -> &dyn Any {
        static Test_AstTree_eval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTree_eval.get_or_init(|| Test_AstTree_eval_004022_002d1.Value)
    }
    pub fn Test_AstTree_describe() -> &dyn Any {
        static Test_AstTree_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTree_describe.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                               &&&string("AST Evaluation:")))
    }
    pub fn Test_AstTree_buildTree_004028() -> &dyn Any {
        &Func1::new(move |v|
                        PureScript_Test_AstTree::Test_AstTree_buildTree_tco(v))
    }
    pub fn Test_AstTree_buildTree_004028_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_AstTree_buildTree_004028_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_AstTree_buildTree_004028_002d1.get_or_init(||
                                                            Lazy(Test_AstTree_buildTree_004028.clone()))
    }
    pub fn Test_AstTree_buildTree_tco(v: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32, &matchValue) {
            0_i32 =>
            &LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Valusd_Ctor(&1_i32)),
            _ => {
                let n = matchValue;
                &LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Addusd_Ctor(&LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Mulusd_Ctor(&LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Valusd_Ctor(&n)),
                                                                                                                                                                                    PureScript_Test_AstTree::Test_AstTree_buildTree_tco(&&(Sharpurs_Prelude::unbox(&&&n)
                                                                                                                                                                                                                                               -
                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&&1_i32))))),
                                                                                                  &LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Subusd_Ctor(PureScript_Test_AstTree::Test_AstTree_buildTree_tco(&&(Sharpurs_Prelude::unbox(&&&n)
                                                                                                                                                                                                                                               -
                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&&1_i32))),
                                                                                                                                                                                    &LrcPtr::new(PureScript_Test_AstTree::Test_AstTree_Expr::Test_AstTree_Valusd_Ctor(&1_i32))))))
            }
        }
    }
    pub fn Test_AstTree_buildTree() -> &dyn Any {
        static Test_AstTree_buildTree: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTree_buildTree.get_or_init(||
                                               Test_AstTree_buildTree_004028_002d1.Value)
    }
    pub fn Test_AstTree_act() -> &dyn Any {
        static Test_AstTree_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTree_act.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                &&&3_i32)),
                                                                          &&&Func1::new(move
                                                                                            |dummy|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                   &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_AstTree::Test_AstTree_eval(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_AstTree::Test_AstTree_buildTree(),
                                                                                                                                                                                                                                      dummy)))))))
    }
}
