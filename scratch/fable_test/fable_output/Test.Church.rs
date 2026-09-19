pub mod PureScript_Test_Church {
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
    pub fn Test_Church_zeroC() -> &dyn Any {
        static Test_Church_zeroC: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_zeroC.get_or_init(||
                                          &Func1::new(move |v|
                                                          &Func1::new(move |x|
                                                                          x.clone())))
    }
    pub fn Test_Church_toInt() -> &dyn Any {
        static Test_Church_toInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_toInt.get_or_init(||
                                          &Func1::new(move |n|
                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(n,
                                                                                                                              &&&Func1::new(move
                                                                                                                                                |x|
                                                                                                                                                &(Sharpurs_Prelude::unbox(x)
                                                                                                                                                      +
                                                                                                                                                      Sharpurs_Prelude::unbox(&&&1_i32)))),
                                                                                           &&&0_i32)))
    }
    pub fn Test_Church_succC() -> &dyn Any {
        static Test_Church_succC: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_succC.get_or_init(||
                                          &Func1::new(move |n|
                                                          &Func1::new({
                                                                          let n
                                                                              =
                                                                              n.clone();
                                                                          move
                                                                              |f|
                                                                              &Func1::new({
                                                                                              let f
                                                                                                  =
                                                                                                  f.clone();
                                                                                              move
                                                                                                  |x|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&n,
                                                                                                                                                                                                         &&&f),
                                                                                                                                                                      x))
                                                                                          })
                                                                      })))
    }
    pub fn Test_Church_mulC() -> &dyn Any {
        static Test_Church_mulC: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_mulC.get_or_init(||
                                         &Func1::new(move |m|
                                                         &Func1::new({
                                                                         let m
                                                                             =
                                                                             m.clone();
                                                                         move
                                                                             |n|
                                                                             &Func1::new({
                                                                                             let n
                                                                                                 =
                                                                                                 n.clone();
                                                                                             move
                                                                                                 |f|
                                                                                                 &Func1::new({
                                                                                                                 let f
                                                                                                                     =
                                                                                                                     f.clone();
                                                                                                                 move
                                                                                                                     |x|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&m,
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&n,
                                                                                                                                                                                                                            &&&f)),
                                                                                                                                                      x)
                                                                                                             })
                                                                                         })
                                                                     })))
    }
    pub fn Test_Church_fromInt_004016() -> &dyn Any {
        &Func1::new(move |v|
                        PureScript_Test_Church::Test_Church_fromInt_tco(v))
    }
    pub fn Test_Church_fromInt_004016_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Test_Church_fromInt_004016_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Test_Church_fromInt_004016_002d1.get_or_init(||
                                                         Lazy(Test_Church_fromInt_004016.clone()))
    }
    pub fn Test_Church_fromInt_tco(v: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32, &matchValue) {
            0_i32 => &PureScript_Test_Church::Test_Church_zeroC(),
            _ =>
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_succC(),
                                             &&PureScript_Test_Church::Test_Church_fromInt_tco(&&(Sharpurs_Prelude::unbox(&&&matchValue)
                                                                                                      -
                                                                                                      Sharpurs_Prelude::unbox(&&&1_i32)))),
        }
    }
    pub fn Test_Church_fromInt() -> &dyn Any {
        static Test_Church_fromInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_fromInt.get_or_init(||
                                            Test_Church_fromInt_004016_002d1.Value)
    }
    pub fn Test_Church_describe() -> &dyn Any {
        static Test_Church_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_describe.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                              &&&string("Church Numerals (100k Closure Applications):")))
    }
    pub fn Test_Church_c10() -> &dyn Any {
        static Test_Church_c10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_c10.get_or_init(||
                                        &Func1::new(move |n|
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_fromInt(),
                                                                                         n)))
    }
    pub fn Test_Church_c100() -> &dyn Any {
        static Test_Church_c100: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_c100.get_or_init(||
                                         &Func1::new(move |n|
                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_mulC(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_c10(),
                                                                                                                                                                n)),
                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_c10(),
                                                                                                                             n))))
    }
    pub fn Test_Church_c10k() -> &dyn Any {
        static Test_Church_c10k: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_c10k.get_or_init(||
                                         &Func1::new(move |n|
                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_mulC(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_c100(),
                                                                                                                                                                n)),
                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_c100(),
                                                                                                                             n))))
    }
    pub fn Test_Church_c100k() -> &dyn Any {
        static Test_Church_c100k: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_c100k.get_or_init(||
                                          &Func1::new(move |n|
                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_mulC(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_c10k(),
                                                                                                                                                                 n)),
                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_c10(),
                                                                                                                              n))))
    }
    pub fn Test_Church_addC() -> &dyn Any {
        static Test_Church_addC: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_addC.get_or_init(||
                                         &Func1::new(move |m|
                                                         &Func1::new({
                                                                         let m
                                                                             =
                                                                             m.clone();
                                                                         move
                                                                             |n|
                                                                             &Func1::new({
                                                                                             let n
                                                                                                 =
                                                                                                 n.clone();
                                                                                             move
                                                                                                 |f|
                                                                                                 &Func1::new({
                                                                                                                 let f
                                                                                                                     =
                                                                                                                     f.clone();
                                                                                                                 move
                                                                                                                     |x|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&m,
                                                                                                                                                                                         &&&f),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&n,
                                                                                                                                                                                                                            &&&f),
                                                                                                                                                                                         x))
                                                                                                             })
                                                                                         })
                                                                     })))
    }
    pub fn Test_Church_act() -> &dyn Any {
        static Test_Church_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_Church_act.get_or_init(||
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
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_toInt(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_Church::Test_Church_c100k(),
                                                                                                                                                                                                                                     dummy)))))))
    }
}
