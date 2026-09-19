pub mod PureScript_Test_Records {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_4da9e9e9::PureScript_Bench;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Test_Records_updateRec_00408() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Test_Records::Test_Records_updateRec_tco(&v,
                                                                                               v1)
                                   }))
    }
    pub fn Test_Records_updateRec_00408_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_Records_updateRec_00408_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_Records_updateRec_00408_002d1.get_or_init(||
                                                           Lazy(Test_Records_updateRec_00408.clone()))
    }
    pub fn Test_Records_updateRec_tco(v: &dyn Any, v1: &dyn Any) -> &dyn Any {
        let v = v.clone();
        let v1 = v1.clone();
        '_Test_Records_updateRec_tco:
            loop  {
                break '_Test_Records_updateRec_tco
                    ({
                         let matchValue = Sharpurs_Prelude::unbox(&&v);
                         let matchValue_1 = Sharpurs_Prelude::unbox(&&v1);
                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                         &matchValue)
                             {
                             0_i32 => r.clone(),
                             _ => {
                                 let r_1 = matchValue_1;
                                 let n = matchValue;
                                 {
                                     let v_temp =
                                         &(Sharpurs_Prelude::unbox(&&&n) -
                                               Sharpurs_Prelude::unbox(&&&1_i32));
                                     let v1_temp =
                                         {
                                             let v2 = &r_1;
                                             &add(string("a"),
                                                  &&(Sharpurs_Prelude::unbox(&&find(string("a"),
                                                                                    Sharpurs_Prelude::unbox(&&r_1)))
                                                         +
                                                         Sharpurs_Prelude::unbox(&&&1_i32)),
                                                  add(string("b"),
                                                      &{
                                                           let v3 =
                                                               find(string("b"),
                                                                    Sharpurs_Prelude::unbox(&&r_1));
                                                           &add(string("c"),
                                                                &&(Sharpurs_Prelude::unbox(&&find(string("c"),
                                                                                                  Sharpurs_Prelude::unbox(&find(string("b"),
                                                                                                                                Sharpurs_Prelude::unbox(&&r_1)))))
                                                                       +
                                                                       Sharpurs_Prelude::unbox(&&&2_i32)),
                                                                add(string("d"),
                                                                    &{
                                                                         let v4 =
                                                                             find(string("d"),
                                                                                  Sharpurs_Prelude::unbox(&find(string("b"),
                                                                                                                Sharpurs_Prelude::unbox(&&r_1))));
                                                                         &add(string("e"),
                                                                              &&(Sharpurs_Prelude::unbox(&&find(string("e"),
                                                                                                                Sharpurs_Prelude::unbox(&find(string("d"),
                                                                                                                                              Sharpurs_Prelude::unbox(&find(string("b"),
                                                                                                                                                                            Sharpurs_Prelude::unbox(&&r_1)))))))
                                                                                     +
                                                                                     Sharpurs_Prelude::unbox(&&&3_i32)),
                                                                              add(string("f"),
                                                                                  &&(Sharpurs_Prelude::unbox(&&find(string("f"),
                                                                                                                    Sharpurs_Prelude::unbox(&find(string("d"),
                                                                                                                                                  Sharpurs_Prelude::unbox(&find(string("b"),
                                                                                                                                                                                Sharpurs_Prelude::unbox(&&r_1)))))))
                                                                                         +
                                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                                          &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                       &&&n),
                                                                                                                                                    &&&5_i32))),
                                                                                  Sharpurs_Prelude::unbox(&&v4)))
                                                                     },
                                                                    Sharpurs_Prelude::unbox(&&v3)))
                                                       },
                                                      Sharpurs_Prelude::unbox(&&v2)))
                                         };
                                     v.set(v_temp);
                                     v1.set(v1_temp);
                                     continue '_Test_Records_updateRec_tco
                                 }
                             }
                         }
                     }) ;
            }
    }
    pub fn Test_Records_updateRec() -> &dyn Any {
        static Test_Records_updateRec: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Records_updateRec.get_or_init(||
                                               Test_Records_updateRec_00408_002d1.Value)
    }
    pub fn Test_Records_initial() -> &dyn Any {
        static Test_Records_initial: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Records_initial.get_or_init(||
                                             &add(string("a"), &&0_i32,
                                                  add(string("b"),
                                                      &&add(string("c"),
                                                            &&0_i32,
                                                            add(string("d"),
                                                                &&add(string("e"),
                                                                      &&0_i32,
                                                                      add(string("f"),
                                                                          &&0_i32,
                                                                          empty::<string,
                                                                                  &dyn Any>())),
                                                                empty::<string,
                                                                        &dyn Any>())),
                                                      empty::<string,
                                                              &dyn Any>())))
    }
    pub fn Test_Records_describe() -> &dyn Any {
        static Test_Records_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Records_describe.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                               &&&string("Deep Record Updates (10k iterations):")))
    }
    pub fn Test_Records_act() -> &dyn Any {
        static Test_Records_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Records_act.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                &&&10000_i32)),
                                                                          &&&Func1::new(move
                                                                                            |dummy|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                   &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                &&find(string("f"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(&find(string("d"),
                                                                                                                                                                                                     Sharpurs_Prelude::unbox(&find(string("b"),
                                                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Records::Test_Records_updateRec(),
                                                                                                                                                                                                                                                                                                                                dummy),
                                                                                                                                                                                                                                                                                             &&&PureScript_Test_Records::Test_Records_initial()))))))))))))
    }
}
