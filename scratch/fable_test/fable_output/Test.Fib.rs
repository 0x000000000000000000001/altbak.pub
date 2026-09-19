pub mod PureScript_Test_Fib {
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
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Test_Fib_fib_00408() -> &dyn Any {
        &Func1::new(move |v| PureScript_Test_Fib::Test_Fib_fib_tco(v))
    }
    pub fn Test_Fib_fib_00408_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_Fib_fib_00408_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_Fib_fib_00408_002d1.get_or_init(||
                                                 Lazy(Test_Fib_fib_00408.clone()))
    }
    pub fn Test_Fib_fib_tco(v: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32, &matchValue) {
            0_i32 => &0_i32,
            _ =>
            match &Sharpurs_Prelude::_007cLitInt_007c__007c(1_i32,
                                                            &matchValue) {
                0_i32 => &1_i32,
                _ => {
                    let n = matchValue;
                    &(Sharpurs_Prelude::unbox(&&PureScript_Test_Fib::Test_Fib_fib_tco(&&(Sharpurs_Prelude::unbox(&&&n)
                                                                                             -
                                                                                             Sharpurs_Prelude::unbox(&&&1_i32))))
                          +
                          Sharpurs_Prelude::unbox(&&PureScript_Test_Fib::Test_Fib_fib_tco(&&(Sharpurs_Prelude::unbox(&&&n)
                                                                                                 -
                                                                                                 Sharpurs_Prelude::unbox(&&&2_i32)))))
                }
            },
        }
    }
    pub fn Test_Fib_fib() -> &dyn Any {
        static Test_Fib_fib: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Test_Fib_fib.get_or_init(|| Test_Fib_fib_00408_002d1.Value)
    }
    pub fn Test_Fib_describe() -> &dyn Any {
        static Test_Fib_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Fib_describe.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                           &&&string("Fibonacci:")))
    }
    pub fn Test_Fib_act() -> &dyn Any {
        static Test_Fib_act: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Test_Fib_act.get_or_init(||
                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                            &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                            &&&10_i32)),
                                                                      &&&Func1::new(move
                                                                                        |dummy|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                            &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                               &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Fib::Test_Fib_fib(),
                                                                                                                                                                                               dummy))))))
    }
}
