pub mod PureScript_Test_FibFFI {
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
    pub mod Test_FibFFI_FFI {
        use super::*;
        pub fn fib(n: i32) -> i32 {
            if n < 2_i32 {
                n
            } else {
                PureScript_Test_FibFFI::Test_FibFFI_FFI::fib(n - 1_i32) +
                    PureScript_Test_FibFFI::Test_FibFFI_FFI::fib(n - 2_i32)
            }
        }
        pub fn runFibFFI(n: &dyn Any) -> &dyn Any {
            &PureScript_Test_FibFFI::Test_FibFFI_FFI::fib(Sharpurs_Prelude::unbox(n))
        }
    }
    pub fn Test_FibFFI_runFibFFI() -> &dyn Any {
        static Test_FibFFI_runFibFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_FibFFI_runFibFFI.get_or_init(||
                                              &Func1::new(move |arg0|
                                                              &PureScript_Test_FibFFI::Test_FibFFI_FFI::runFibFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_FibFFI_describe() -> &dyn Any {
        static Test_FibFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_FibFFI_describe.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                              &&&string("Fibonacci FFI:")))
    }
    pub fn Test_FibFFI_act() -> &dyn Any {
        static Test_FibFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_FibFFI_act.get_or_init(||
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
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_FibFFI::Test_FibFFI_runFibFFI(),
                                                                                                                                                                                                  dummy))))))
    }
}
