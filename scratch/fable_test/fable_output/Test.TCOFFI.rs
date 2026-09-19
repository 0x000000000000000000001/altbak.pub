pub mod PureScript_Test_TCOFFI {
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
    pub mod Test_TCOFFI_FFI {
        use super::*;
        pub fn r#loop(n: i32, acc: i32) -> i32 {
            let n: MutCell<i32> = MutCell::new(n);
            let acc: MutCell<i32> = MutCell::new(acc);
            '_loop:
                loop  {
                    break '_loop
                        (if n.get() == 0_i32 {
                             acc.get()
                         } else {
                             let n_temp: i32 = n.get() - 1_i32;
                             let acc_temp: i32 = acc.get() + n.get() % 3_i32;
                             n.set(n_temp);
                             acc.set(acc_temp);
                             continue '_loop
                         }) ;
                }
        }
        pub fn runTCOFFI(n: &dyn Any) -> &dyn Any {
            &PureScript_Test_TCOFFI::Test_TCOFFI_FFI::r#loop(Sharpurs_Prelude::unbox(n),
                                                             0_i32)
        }
    }
    pub fn Test_TCOFFI_runTCOFFI() -> &dyn Any {
        static Test_TCOFFI_runTCOFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_TCOFFI_runTCOFFI.get_or_init(||
                                              &Func1::new(move |arg0|
                                                              &PureScript_Test_TCOFFI::Test_TCOFFI_FFI::runTCOFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_TCOFFI_describe() -> &dyn Any {
        static Test_TCOFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_TCOFFI_describe.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                              &&&string("Tail Call Optimization FFI (100k calls):")))
    }
    pub fn Test_TCOFFI_act() -> &dyn Any {
        static Test_TCOFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_TCOFFI_act.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                               &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                               &&&100000_i32)),
                                                                         &&&Func1::new(move
                                                                                           |dummy|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                               &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                  &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_TCOFFI::Test_TCOFFI_runTCOFFI(),
                                                                                                                                                                                                  dummy))))))
    }
}
