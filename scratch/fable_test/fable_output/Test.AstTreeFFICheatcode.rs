pub mod PureScript_Test_AstTreeFFICheatcode {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Test_AstTreeFFICheatcode_FFI {
        use super::*;
        use fable_library_rust::Native_::LrcPtr;
        #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
        pub enum Expr {
            Val(i32),
            Add(LrcPtr<PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr>,
                LrcPtr<PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr>),
            Mul(LrcPtr<PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr>,
                LrcPtr<PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr>),
            Sub(LrcPtr<PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr>,
                LrcPtr<PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr>),
        }
        impl core::fmt::Display for
         PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr
         {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn build(n: i32)
         ->
             LrcPtr<PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr> {
            if n == 0_i32 {
                LrcPtr::new(PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Val(1_i32))
            } else {
                LrcPtr::new(PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Add(LrcPtr::new(PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Mul(LrcPtr::new(PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Val(n)),
                                                                                                                                                                                                  PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::build(n
                                                                                                                                                                                                                                                                               -
                                                                                                                                                                                                                                                                               1_i32))),
                                                                                                         LrcPtr::new(PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Sub(PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::build(n
                                                                                                                                                                                                                                                                               -
                                                                                                                                                                                                                                                                               1_i32),
                                                                                                                                                                                                  LrcPtr::new(PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Val(1_i32))))))
            }
        }
        pub fn eval(_arg:
                        LrcPtr<PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr>)
         -> i32 {
            match _arg.as_ref() {
                PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Add(_arg_1_0,
                                                                                             _arg_1_1)
                =>
                PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::eval(_arg_1_0.clone())
                    +
                    PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::eval(_arg_1_1.clone()),
                PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Mul(_arg_2_0,
                                                                                             _arg_2_1)
                =>
                PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::eval(_arg_2_0.clone())
                    *
                    PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::eval(_arg_2_1.clone()),
                PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Sub(_arg_3_0,
                                                                                             _arg_3_1)
                =>
                PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::eval(_arg_3_0.clone())
                    -
                    PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::eval(_arg_3_1.clone()),
                PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::Expr::Val(_arg_0_0)
                => _arg_0_0.clone(),
            }
        }
        pub fn runAstTreeFFICheatcode(n: &dyn Any) -> &dyn Any {
            &PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::eval(PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::build(Sharpurs_Prelude::unbox(n)))
        }
    }
    pub fn Test_AstTreeFFICheatcode_runAstTreeFFICheatcode() -> &dyn Any {
        static Test_AstTreeFFICheatcode_runAstTreeFFICheatcode:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTreeFFICheatcode_runAstTreeFFICheatcode.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |arg0|
                                                                                        &PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_FFI::runAstTreeFFICheatcode(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_AstTreeFFICheatcode_describe() -> &dyn Any {
        static Test_AstTreeFFICheatcode_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTreeFFICheatcode_describe.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                           &&&string("AST Evaluation FFICheatcode:")))
    }
    pub fn Test_AstTreeFFICheatcode_act() -> &dyn Any {
        static Test_AstTreeFFICheatcode_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_AstTreeFFICheatcode_act.get_or_init(||
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
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_AstTreeFFICheatcode::Test_AstTreeFFICheatcode_runAstTreeFFICheatcode(),
                                                                                                                                                                                                               dummy))))))
    }
}
