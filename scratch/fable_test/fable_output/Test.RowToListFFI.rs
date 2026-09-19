pub mod PureScript_Test_RowToListFFI {
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
    pub mod Test_RowToListFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::Func0;
        use fable_library_rust::Native_::LrcPtr;
        #[derive(Clone, Debug,)]
        pub struct RecordKeys {
            pub keysImpl: Func0<i32>,
        }
        impl core::fmt::Display for
         PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::RecordKeys {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn keysCons(tail:
                            LrcPtr<PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::RecordKeys>)
         ->
             LrcPtr<PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::RecordKeys> {
            LrcPtr::new(PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::RecordKeys{keysImpl:
                                                                                            Func0::new({
                                                                                                           let tail
                                                                                                               =
                                                                                                               tail.clone();
                                                                                                           move
                                                                                                               ||
                                                                                                               1_i32
                                                                                                                   +
                                                                                                                   (tail.keysImpl)()
                                                                                                       }),})
        }
        pub fn runRowToListFFI(_input: &dyn Any) -> &dyn Any {
            &((PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::keysCons(PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::keysCons(PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::keysCons(PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::keysCons(PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::keysCons(LrcPtr::new(PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::RecordKeys{keysImpl:
                                                                                                                                                                                                                                                                                                                                                                                                                     Func0::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                    ||
                                                                                                                                                                                                                                                                                                                                                                                                                                    0_i32),}))))))).keysImpl)()
        }
    }
    pub fn Test_RowToListFFI_runRowToListFFI() -> &dyn Any {
        static Test_RowToListFFI_runRowToListFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToListFFI_runRowToListFFI.get_or_init(||
                                                          &Func1::new(move
                                                                          |arg0|
                                                                          &PureScript_Test_RowToListFFI::Test_RowToListFFI_FFI::runRowToListFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_RowToListFFI_describe() -> &dyn Any {
        static Test_RowToListFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToListFFI_describe.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                    &&&string("RowToList FFI (Keys Count):")))
    }
    pub fn Test_RowToListFFI_act() -> &dyn Any {
        static Test_RowToListFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RowToListFFI_act.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                     &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                     &&&0_i32)),
                                                                               &&&Func1::new(move
                                                                                                 |dummy|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                     &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                        &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RowToListFFI::Test_RowToListFFI_runRowToListFFI(),
                                                                                                                                                                                                        dummy))))))
    }
}
