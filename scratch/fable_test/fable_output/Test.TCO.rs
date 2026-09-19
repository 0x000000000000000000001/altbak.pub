pub mod PureScript_Test_TCO {
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
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Test_TCO_describe() -> &dyn Any {
        static Test_TCO_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_TCO_describe.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                           &&&string("Tail Call Optimization (100k calls):")))
    }
    pub fn Test_TCO_deepTailRec_004010() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Test_TCO::Test_TCO_deepTailRec_tco(&v,
                                                                                         v1)
                                   }))
    }
    pub fn Test_TCO_deepTailRec_004010_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_TCO_deepTailRec_004010_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_TCO_deepTailRec_004010_002d1.get_or_init(||
                                                          Lazy(Test_TCO_deepTailRec_004010.clone()))
    }
    pub fn Test_TCO_deepTailRec_tco(v: &dyn Any, v1: &dyn Any) -> &dyn Any {
        let v = v.clone();
        let v1 = v1.clone();
        '_Test_TCO_deepTailRec_tco:
            loop  {
                break '_Test_TCO_deepTailRec_tco
                    ({
                         let matchValue = Sharpurs_Prelude::unbox(&&v);
                         let matchValue_1 = Sharpurs_Prelude::unbox(&&v1);
                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                         &matchValue)
                             {
                             0_i32 => acc.clone(),
                             _ => {
                                 let n = matchValue;
                                 {
                                     let v_temp =
                                         &(Sharpurs_Prelude::unbox(&&&n) -
                                               Sharpurs_Prelude::unbox(&&&1_i32));
                                     let v1_temp =
                                         &(Sharpurs_Prelude::unbox(&&&matchValue_1)
                                               +
                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                             &&&n),
                                                                                                          &&&3_i32)));
                                     v.set(v_temp);
                                     v1.set(v1_temp);
                                     continue '_Test_TCO_deepTailRec_tco
                                 }
                             }
                         }
                     }) ;
            }
    }
    pub fn Test_TCO_deepTailRec() -> &dyn Any {
        static Test_TCO_deepTailRec: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_TCO_deepTailRec.get_or_init(||
                                             Test_TCO_deepTailRec_004010_002d1.Value)
    }
    pub fn Test_TCO_act() -> &dyn Any {
        static Test_TCO_act: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Test_TCO_act.get_or_init(||
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
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_TCO::Test_TCO_deepTailRec(),
                                                                                                                                                                                                                                  dummy),
                                                                                                                                                                                               &&&0_i32))))))
    }
}
