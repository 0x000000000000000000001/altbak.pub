pub mod PureScript_Test_StateMonadFFICheatcode {
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
    pub mod Test_StateMonadFFICheatcode_FFI {
        use super::*;
        pub fn runStateMonadFFICheatcode(input: &dyn Any) -> &dyn Any {
            let depth: i32 = Sharpurs_Prelude::unbox(input);
            let total: MutCell<i32> = MutCell::new(0_i32);
            for forLoopVar in 1_i32..=20_i32 {
                let state: MutCell<i32> = MutCell::new(0_i32);
                for forLoopVar_1 in 1_i32..=depth {
                    state.set(state.get() + 1_i32);
                }
                total.set(total.get() + state.get())
            }
            &total.get()
        }
    }
    pub fn Test_StateMonadFFICheatcode_runStateMonadFFICheatcode()
     -> &dyn Any {
        static Test_StateMonadFFICheatcode_runStateMonadFFICheatcode:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonadFFICheatcode_runStateMonadFFICheatcode.get_or_init(||
                                                                              &Func1::new(move
                                                                                              |arg0|
                                                                                              &PureScript_Test_StateMonadFFICheatcode::Test_StateMonadFFICheatcode_FFI::runStateMonadFFICheatcode(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_StateMonadFFICheatcode_describe() -> &dyn Any {
        static Test_StateMonadFFICheatcode_describe: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Test_StateMonadFFICheatcode_describe.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                              &&&string("State Monad FFICheatcode (1.2k Binds, 60 Stack Depth):")))
    }
    pub fn Test_StateMonadFFICheatcode_act() -> &dyn Any {
        static Test_StateMonadFFICheatcode_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonadFFICheatcode_act.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                               &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                               &&&60_i32)),
                                                                                         &&&Func1::new(move
                                                                                                           |dummy|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                               &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                  &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonadFFICheatcode::Test_StateMonadFFICheatcode_runStateMonadFFICheatcode(),
                                                                                                                                                                                                                  dummy))))))
    }
}
