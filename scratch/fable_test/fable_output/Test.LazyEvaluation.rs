pub mod PureScript_Test_LazyEvaluation {
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
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Test_LazyEvaluation_Lazy() -> &dyn Any {
        static Test_LazyEvaluation_Lazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluation_Lazy.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Test_LazyEvaluation_force() -> &dyn Any {
        static Test_LazyEvaluation_force: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluation_force.get_or_init(||
                                                  &Func1::new(move |v|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(v),
                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit())))
    }
    pub fn Test_LazyEvaluation_describe() -> &dyn Any {
        static Test_LazyEvaluation_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluation_describe.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                      &&&string("Lazy Evaluation (1M Thunks Forced, 1k Depth):")))
    }
    pub fn Test_LazyEvaluation_defer() -> &dyn Any {
        static Test_LazyEvaluation_defer: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluation_defer.get_or_init(||
                                                  &PureScript_Test_LazyEvaluation::Test_LazyEvaluation_Lazy())
    }
    pub fn Test_LazyEvaluation_buildThunks_004016() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Test_LazyEvaluation::Test_LazyEvaluation_buildThunks_tco(&v,
                                                                                                               v1)
                                   }))
    }
    pub fn Test_LazyEvaluation_buildThunks_004016_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_LazyEvaluation_buildThunks_004016_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_LazyEvaluation_buildThunks_004016_002d1.get_or_init(||
                                                                     Lazy(Test_LazyEvaluation_buildThunks_004016.clone()))
    }
    pub fn Test_LazyEvaluation_buildThunks_tco(v: &dyn Any, v1: &dyn Any)
     -> &dyn Any {
        let v = v.clone();
        let v1 = v1.clone();
        '_Test_LazyEvaluation_buildThunks_tco:
            loop  {
                break '_Test_LazyEvaluation_buildThunks_tco
                    ({
                         let matchValue = Sharpurs_Prelude::unbox(&&v);
                         let matchValue_1 = Sharpurs_Prelude::unbox(&&v1);
                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                         &matchValue)
                             {
                             0_i32 => acc.clone(),
                             _ => {
                                 let v_temp =
                                     &(Sharpurs_Prelude::unbox(&&&matchValue)
                                           -
                                           Sharpurs_Prelude::unbox(&&&1_i32));
                                 let v1_temp =
                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_LazyEvaluation::Test_LazyEvaluation_defer(),
                                                                      &&&Func1::new({
                                                                                        let matchValue_1
                                                                                            =
                                                                                            matchValue_1.clone();
                                                                                        move
                                                                                            |v2|
                                                                                            &(Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_LazyEvaluation::Test_LazyEvaluation_force(),
                                                                                                                                                         &&&matchValue_1))
                                                                                                  +
                                                                                                  Sharpurs_Prelude::unbox(&&&1_i32))
                                                                                    }));
                                 v.set(v_temp);
                                 v1.set(v1_temp);
                                 continue
                                     '_Test_LazyEvaluation_buildThunks_tco
                             }
                         }
                     }) ;
            }
    }
    pub fn Test_LazyEvaluation_buildThunks() -> &dyn Any {
        static Test_LazyEvaluation_buildThunks: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluation_buildThunks.get_or_init(||
                                                        Test_LazyEvaluation_buildThunks_004016_002d1.Value)
    }
    pub fn Test_LazyEvaluation_runManyTimes_004020() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Test_LazyEvaluation::Test_LazyEvaluation_runManyTimes_tco(&v,
                                                                                                                v1)
                                   }))
    }
    pub fn Test_LazyEvaluation_runManyTimes_004020_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_LazyEvaluation_runManyTimes_004020_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_LazyEvaluation_runManyTimes_004020_002d1.get_or_init(||
                                                                      Lazy(Test_LazyEvaluation_runManyTimes_004020.clone()))
    }
    pub fn Test_LazyEvaluation_runManyTimes_tco(v: &dyn Any, v1: &dyn Any)
     -> &dyn Any {
        let v = v.clone();
        let v1 = v1.clone();
        '_Test_LazyEvaluation_runManyTimes_tco:
            loop  {
                break '_Test_LazyEvaluation_runManyTimes_tco
                    ({
                         let matchValue = Sharpurs_Prelude::unbox(&&v);
                         let matchValue_1 = Sharpurs_Prelude::unbox(&&v1);
                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                         &matchValue)
                             {
                             0_i32 => acc.clone(),
                             _ => {
                                 let v_temp =
                                     &(Sharpurs_Prelude::unbox(&&&matchValue)
                                           -
                                           Sharpurs_Prelude::unbox(&&&1_i32));
                                 let v1_temp =
                                     &(Sharpurs_Prelude::unbox(&&&matchValue_1)
                                           +
                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_LazyEvaluation::Test_LazyEvaluation_force(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_LazyEvaluation::Test_LazyEvaluation_buildThunks(),
                                                                                                                                                                            &&&1000_i32),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_LazyEvaluation::Test_LazyEvaluation_defer(),
                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                              |v2|
                                                                                                                                                                                              &0_i32))))));
                                 v.set(v_temp);
                                 v1.set(v1_temp);
                                 continue
                                     '_Test_LazyEvaluation_runManyTimes_tco
                             }
                         }
                     }) ;
            }
    }
    pub fn Test_LazyEvaluation_runManyTimes() -> &dyn Any {
        static Test_LazyEvaluation_runManyTimes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluation_runManyTimes.get_or_init(||
                                                         Test_LazyEvaluation_runManyTimes_004020_002d1.Value)
    }
    pub fn Test_LazyEvaluation_act() -> &dyn Any {
        static Test_LazyEvaluation_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_LazyEvaluation_act.get_or_init(||
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
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_LazyEvaluation::Test_LazyEvaluation_runManyTimes(),
                                                                                                                                                                                                                                             dummy),
                                                                                                                                                                                                          &&&0_i32))))))
    }
}
