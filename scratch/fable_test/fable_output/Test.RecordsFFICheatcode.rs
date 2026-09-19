pub mod PureScript_Test_RecordsFFICheatcode {
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
    pub mod Test_RecordsFFICheatcode_FFI {
        use super::*;
        use fable_library_rust::Exception_::finally;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IEnumerable_1;
        use fable_library_rust::Interfaces_::System::Collections::Generic::IEnumerator_1;
        use fable_library_rust::Interfaces_::System::IDisposable;
        use fable_library_rust::Native_::LrcPtr;
        use fable_library_rust::Range_::rangeNumeric;
        #[derive(Clone, Debug, Default, PartialEq, PartialOrd, Hash, Eq,
                 Ord,)]
        pub struct Inner {
            pub e: MutCell<i32>,
            pub f: MutCell<i32>,
        }
        impl core::fmt::Display for
         PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::Inner
         {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        #[derive(Clone, Debug, Default, PartialEq, PartialOrd, Hash, Eq,
                 Ord,)]
        pub struct Middle {
            pub c: MutCell<i32>,
            pub d: LrcPtr<PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::Inner>,
        }
        impl core::fmt::Display for
         PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::Middle
         {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        #[derive(Clone, Debug, Default, PartialEq, PartialOrd, Hash, Eq,
                 Ord,)]
        pub struct DeepRecord {
            pub a: MutCell<i32>,
            pub b: LrcPtr<PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::Middle>,
        }
        impl core::fmt::Display for
         PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::DeepRecord
         {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn runRecordsFFICheatcode(n: &dyn Any) -> &dyn Any {
            let record:
                    LrcPtr<PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::DeepRecord> =
                LrcPtr::new(PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::DeepRecord{a:
                                                                                                              MutCell::new(0_i32),
                                                                                                          b:
                                                                                                              LrcPtr::new(PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::Middle{c:
                                                                                                                                                                                                        MutCell::new(0_i32),
                                                                                                                                                                                                    d:
                                                                                                                                                                                                        LrcPtr::new(PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::Inner{e:
                                                                                                                                                                                                                                                                                                 MutCell::new(0_i32),
                                                                                                                                                                                                                                                                                             f:
                                                                                                                                                                                                                                                                                                 MutCell::new(0_i32),}),}),});
            {
                let inputSequence: LrcPtr<dyn IEnumerable_1<i32>> =
                    rangeNumeric(Sharpurs_Prelude::unbox(n), -1_i32, 1_i32);
                let enumerator: LrcPtr<dyn IEnumerator_1<i32>> =
                    IEnumerable_1::GetEnumerator(inputSequence.as_ref());
                {
                    finally(||
                                if ((&enumerator) as
                                        &dyn Any).is::<LrcPtr<dyn IDisposable>>()
                                   {
                                    (&enumerator).Dispose();
                                });
                    while IEnumerator_1::MoveNext(enumerator.as_ref()) {
                        let remaining: i32 =
                            IEnumerator_1::get_Current(enumerator.as_ref());
                        record.a.set(record.a.get() + 1_i32);
                        (record.b).c.set((record.b).c.get() + 2_i32);
                        ((record.b).d).e.set(((record.b).d).e.get() + 3_i32);
                        ((record.b).d).f.set(((record.b).d).f.get() +
                                                 remaining % 5_i32)
                    }
                }
            }
            &((record.b).d).f.get()
        }
    }
    pub fn Test_RecordsFFICheatcode_runRecordsFFICheatcode() -> &dyn Any {
        static Test_RecordsFFICheatcode_runRecordsFFICheatcode:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RecordsFFICheatcode_runRecordsFFICheatcode.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |arg0|
                                                                                        &PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_FFI::runRecordsFFICheatcode(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_RecordsFFICheatcode_describe() -> &dyn Any {
        static Test_RecordsFFICheatcode_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RecordsFFICheatcode_describe.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                           &&&string("Deep Record Updates FFICheatcode (10k iterations):")))
    }
    pub fn Test_RecordsFFICheatcode_act() -> &dyn Any {
        static Test_RecordsFFICheatcode_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_RecordsFFICheatcode_act.get_or_init(||
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
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_RecordsFFICheatcode::Test_RecordsFFICheatcode_runRecordsFFICheatcode(),
                                                                                                                                                                                                               dummy))))))
    }
}
