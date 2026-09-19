pub mod PureScript_Test_ArrayOpsFFI {
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
    pub mod Test_ArrayOpsFFI_FFI {
        use super::*;
        use fable_library_rust::Array_::filter;
        use fable_library_rust::Array_::fold;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IEnumerable_1;
        use fable_library_rust::Native_::Func2;
        use fable_library_rust::Native_::LrcPtr;
        use fable_library_rust::Range_::rangeNumeric;
        use fable_library_rust::Seq_::toArray;
        pub fn runArrayOpsFFI(n: &dyn Any) -> &dyn Any {
            let finish: i32 = Sharpurs_Prelude::unbox(n);
            &fold(Func2::new(move |x: i32, y: i32| x + y), 0_i32,
                  filter(Func1::new(move |value: i32| value % 2_i32 == 0_i32),
                         toArray(rangeNumeric(1_i32,
                                              if finish >= 1_i32 {
                                                  1_i32
                                              } else { -1_i32 }, finish))))
        }
    }
    pub fn Test_ArrayOpsFFI_runArrayOpsFFI() -> &dyn Any {
        static Test_ArrayOpsFFI_runArrayOpsFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOpsFFI_runArrayOpsFFI.get_or_init(||
                                                        &Func1::new(move
                                                                        |arg0|
                                                                        &PureScript_Test_ArrayOpsFFI::Test_ArrayOpsFFI_FFI::runArrayOpsFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_ArrayOpsFFI_describe() -> &dyn Any {
        static Test_ArrayOpsFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOpsFFI_describe.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                   &&&string("Array Processing FFI (900 elements):")))
    }
    pub fn Test_ArrayOpsFFI_act() -> &dyn Any {
        static Test_ArrayOpsFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ArrayOpsFFI_act.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                    &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                    &&&900_i32)),
                                                                              &&&Func1::new(move
                                                                                                |dummy|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                    &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                       &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ArrayOpsFFI::Test_ArrayOpsFFI_runArrayOpsFFI(),
                                                                                                                                                                                                       dummy))))))
    }
}
