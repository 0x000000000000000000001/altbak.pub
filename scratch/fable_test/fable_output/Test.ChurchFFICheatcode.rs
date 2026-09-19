pub mod PureScript_Test_ChurchFFICheatcode {
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
    pub mod Test_ChurchFFICheatcode_FFI {
        use super::*;
        pub fn runChurchFFICheatcode(input: &dyn Any) -> &dyn Any {
            let count: i32 = Sharpurs_Prelude::unbox(input).pow(5_i32);
            let result: MutCell<i32> = MutCell::new(0_i32);
            for forLoopVar in 1_i32..=count {
                result.set(result.get() + 1_i32);
            }
            &result.get()
        }
    }
    pub fn Test_ChurchFFICheatcode_runChurchFFICheatcode() -> &dyn Any {
        static Test_ChurchFFICheatcode_runChurchFFICheatcode:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ChurchFFICheatcode_runChurchFFICheatcode.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |arg0|
                                                                                      &PureScript_Test_ChurchFFICheatcode::Test_ChurchFFICheatcode_FFI::runChurchFFICheatcode(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_ChurchFFICheatcode_describe() -> &dyn Any {
        static Test_ChurchFFICheatcode_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ChurchFFICheatcode_describe.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                          &&&string("Church Numerals FFICheatcode (100k Closure Applications):")))
    }
    pub fn Test_ChurchFFICheatcode_act() -> &dyn Any {
        static Test_ChurchFFICheatcode_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ChurchFFICheatcode_act.get_or_init(||
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
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ChurchFFICheatcode::Test_ChurchFFICheatcode_runChurchFFICheatcode(),
                                                                                                                                                                                                              dummy))))))
    }
}
