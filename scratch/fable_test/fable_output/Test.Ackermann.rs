pub mod PureScript_Test_Ackermann {
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
    use crate::module_3ed61c25::PureScript_Effect_Console;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Test_Ackermann_describe() -> &dyn Any {
        static Test_Ackermann_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Ackermann_describe.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                 &&&string("Ackermann (3, 4):")))
    }
    pub fn Test_Ackermann_ackermann_004010() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           PureScript_Test_Ackermann::Test_Ackermann_ackermann_tco(&v,
                                                                                                   v1)
                                   }))
    }
    pub fn Test_Ackermann_ackermann_004010_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_Ackermann_ackermann_004010_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_Ackermann_ackermann_004010_002d1.get_or_init(||
                                                              Lazy(Test_Ackermann_ackermann_004010.clone()))
    }
    pub fn Test_Ackermann_ackermann_tco(v: &dyn Any, v1: &dyn Any)
     -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32, &matchValue) {
            0_i32 =>
            &(Sharpurs_Prelude::unbox(n) + Sharpurs_Prelude::unbox(&&&1_i32)),
            _ =>
            match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                            &matchValue_1) {
                0_i32 =>
                PureScript_Test_Ackermann::Test_Ackermann_ackermann_tco(&&(Sharpurs_Prelude::unbox(m)
                                                                               -
                                                                               Sharpurs_Prelude::unbox(&&&1_i32)),
                                                                        &&1_i32),
                _ => {
                    let m_1 = matchValue;
                    PureScript_Test_Ackermann::Test_Ackermann_ackermann_tco(&&(Sharpurs_Prelude::unbox(&&&m_1)
                                                                                   -
                                                                                   Sharpurs_Prelude::unbox(&&&1_i32)),
                                                                            &PureScript_Test_Ackermann::Test_Ackermann_ackermann_tco(&&m_1,
                                                                                                                                     &&(Sharpurs_Prelude::unbox(&&&matchValue_1)
                                                                                                                                            -
                                                                                                                                            Sharpurs_Prelude::unbox(&&&1_i32))))
                }
            },
        }
    }
    pub fn Test_Ackermann_ackermann() -> &dyn Any {
        static Test_Ackermann_ackermann: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Ackermann_ackermann.get_or_init(||
                                                 Test_Ackermann_ackermann_004010_002d1.Value)
    }
    pub fn Test_Ackermann_act() -> &dyn Any {
        static Test_Ackermann_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Ackermann_act.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                  &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                  &&&3_i32)),
                                                                            &&&Func1::new(move
                                                                                              |dummy|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                  &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                     &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Ackermann::Test_Ackermann_ackermann(),
                                                                                                                                                                                                                                        dummy),
                                                                                                                                                                                                     &&&4_i32))))))
    }
}
