pub mod PureScript_Test_RecordsFFI {
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
    pub mod Test_RecordsFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::LrcPtr;
        #[derive(Clone, Debug, Default, PartialEq, PartialOrd, Hash, Eq,
                 Ord,)]
        pub struct Inner {
            pub e: i32,
            pub f: i32,
        }
        impl core::fmt::Display for
         PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::Inner {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        #[derive(Clone, Debug, Default, PartialEq, PartialOrd, Hash, Eq,
                 Ord,)]
        pub struct Middle {
            pub c: i32,
            pub d: LrcPtr<PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::Inner>,
        }
        impl core::fmt::Display for
         PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::Middle {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        #[derive(Clone, Debug, Default, PartialEq, PartialOrd, Hash, Eq,
                 Ord,)]
        pub struct DeepRecord {
            pub a: i32,
            pub b: LrcPtr<PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::Middle>,
        }
        impl core::fmt::Display for
         PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::DeepRecord {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn update(n: i32,
                      r:
                          LrcPtr<PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::DeepRecord>)
         ->
             LrcPtr<PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::DeepRecord> {
            let n: MutCell<i32> = MutCell::new(n);
            let r:
                    MutCell<LrcPtr<PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::DeepRecord>> =
                MutCell::new(r.clone());
            '_update:
                loop  {
                    break '_update
                        (if n.get() == 0_i32 {
                             r.get()
                         } else {
                             let n_temp: i32 = n.get() - 1_i32;
                             let r_temp:
                                     LrcPtr<PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::DeepRecord> =
                                 LrcPtr::new(PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::DeepRecord{a:
                                                                                                             r.a
                                                                                                                 +
                                                                                                                 1_i32,
                                                                                                         b:
                                                                                                             LrcPtr::new(PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::Middle{c:
                                                                                                                                                                                     (r.b).c
                                                                                                                                                                                         +
                                                                                                                                                                                         2_i32,
                                                                                                                                                                                 d:
                                                                                                                                                                                     LrcPtr::new(PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::Inner{e:
                                                                                                                                                                                                                                                            ((r.b).d).e
                                                                                                                                                                                                                                                                +
                                                                                                                                                                                                                                                                3_i32,
                                                                                                                                                                                                                                                        f:
                                                                                                                                                                                                                                                            ((r.b).d).f
                                                                                                                                                                                                                                                                +
                                                                                                                                                                                                                                                                n.get()
                                                                                                                                                                                                                                                                    %
                                                                                                                                                                                                                                                                    5_i32,}),}),});
                             n.set(n_temp);
                             r.set(r_temp);
                             continue '_update
                         }) ;
                }
        }
        pub fn runRecordsFFI(n: &dyn Any) -> &dyn Any {
            let initial:
                    LrcPtr<PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::DeepRecord> =
                LrcPtr::new(PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::DeepRecord{a:
                                                                                            0_i32,
                                                                                        b:
                                                                                            LrcPtr::new(PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::Middle{c:
                                                                                                                                                                    0_i32,
                                                                                                                                                                d:
                                                                                                                                                                    LrcPtr::new(PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::Inner{e:
                                                                                                                                                                                                                                           0_i32,
                                                                                                                                                                                                                                       f:
                                                                                                                                                                                                                                           0_i32,}),}),});
            &(((PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::update(Sharpurs_Prelude::unbox(n),
                                                                        initial)).b).d).f
        }
    }
    pub fn Test_RecordsFFI_runRecordsFFI() -> &dyn Any {
        static Test_RecordsFFI_runRecordsFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RecordsFFI_runRecordsFFI.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Test_RecordsFFI::Test_RecordsFFI_FFI::runRecordsFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_RecordsFFI_describe() -> &dyn Any {
        static Test_RecordsFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RecordsFFI_describe.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                  &&&string("Deep Record Updates FFI (10k iterations):")))
    }
    pub fn Test_RecordsFFI_act() -> &dyn Any {
        static Test_RecordsFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RecordsFFI_act.get_or_init(||
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
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RecordsFFI::Test_RecordsFFI_runRecordsFFI(),
                                                                                                                                                                                                      dummy))))))
    }
}
