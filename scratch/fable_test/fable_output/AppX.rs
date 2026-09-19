pub mod PureScript_AppX {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_55d52a3::PureScript_Test_Primes;
    pub fn AppX_main() -> &dyn Any {
        static AppX_main: MutCell<Option<&dyn Any>> = MutCell::new(None);
        AppX_main.get_or_init(||
                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                                                                                         &&&PureScript_Effect::Effect_functorEffect())),
                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_runBench(),
                                                                                                                                         &&&PureScript_Test_Primes::Test_Primes_describe()),
                                                                                                      &&&PureScript_Test_Primes::Test_Primes_act())))
    }
}
