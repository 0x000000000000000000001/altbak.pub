pub mod PureScript_Test_LazyEvaluationFFI {
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
    pub mod Test_LazyEvaluationFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::Func0;
        pub fn buildThunks(depth: i32, initial: Func0<i32>) -> Func0<i32> {
            let result = MutCell::new(initial);
            for forLoopVar in 1_i32..=depth {
                let previous = result.get();
                result.set(Func0::new({
                                          let previous = previous.clone();
                                          move || previous() + 1_i32
                                      }))
            }
            result.get()
        }
        pub fn runLazyEvaluationFFI(input: &dyn Any) -> &dyn Any {
            let result: MutCell<i32> = MutCell::new(0_i32);
            for forLoopVar in 1_i32..=Sharpurs_Prelude::unbox(input) {
                result.set(result.get() +
                               (PureScript_Test_LazyEvaluationFFI::Test_LazyEvaluationFFI_FFI::buildThunks(1000_i32,
                                                                                                           Func0::new(move
                                                                                                                          ||
                                                                                                                          0_i32)))());
            }
            &result.get()
        }
    }
    pub fn Test_LazyEvaluationFFI_runLazyEvaluationFFI() -> &dyn Any {
        static Test_LazyEvaluationFFI_runLazyEvaluationFFI:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluationFFI_runLazyEvaluationFFI.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |arg0|
                                                                                    &PureScript_Test_LazyEvaluationFFI::Test_LazyEvaluationFFI_FFI::runLazyEvaluationFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_LazyEvaluationFFI_describe() -> &dyn Any {
        static Test_LazyEvaluationFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluationFFI_describe.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                         &&&string("Lazy Evaluation FFI (1M Thunks Forced, 1k Depth):")))
    }
    pub fn Test_LazyEvaluationFFI_act() -> &dyn Any {
        static Test_LazyEvaluationFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluationFFI_act.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                          &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                          &&&1000_i32)),
                                                                                    &&&Func1::new(move
                                                                                                      |dummy|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                          &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                             &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_LazyEvaluationFFI::Test_LazyEvaluationFFI_runLazyEvaluationFFI(),
                                                                                                                                                                                                             dummy))))))
    }
}
