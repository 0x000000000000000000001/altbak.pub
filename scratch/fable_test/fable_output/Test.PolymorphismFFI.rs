pub mod PureScript_Test_PolymorphismFFI {
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
    pub mod Test_PolymorphismFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::Func2;
        use fable_library_rust::Native_::LrcPtr;
        #[derive(Clone, Debug,)]
        pub struct Monoidish {
            pub mempty: i32,
            pub mappend: Func2<i32, i32, i32>,
        }
        impl core::fmt::Display for
         PureScript_Test_PolymorphismFFI::Test_PolymorphismFFI_FFI::Monoidish
         {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn polyLoop(dictionary:
                            LrcPtr<PureScript_Test_PolymorphismFFI::Test_PolymorphismFFI_FFI::Monoidish>,
                        count: i32, acc: i32) -> i32 {
            let dictionary:
                    MutCell<LrcPtr<PureScript_Test_PolymorphismFFI::Test_PolymorphismFFI_FFI::Monoidish>> =
                MutCell::new(dictionary.clone());
            let count: MutCell<i32> = MutCell::new(count);
            let acc: MutCell<i32> = MutCell::new(acc);
            '_polyLoop:
                loop  {
                    break '_polyLoop
                        (if count.get() == 0_i32 {
                             acc.get()
                         } else {
                             let dictionary_temp:
                                     LrcPtr<PureScript_Test_PolymorphismFFI::Test_PolymorphismFFI_FFI::Monoidish> =
                                 dictionary.get();
                             let count_temp: i32 = count.get() - 1_i32;
                             let acc_temp: i32 =
                                 (dictionary.mappend)(acc.get(),
                                                      dictionary.mempty);
                             dictionary.set(dictionary_temp);
                             count.set(count_temp);
                             acc.set(acc_temp);
                             continue '_polyLoop
                         }) ;
                }
        }
        pub fn runPolymorphismFFI(n: &dyn Any) -> &dyn Any {
            &PureScript_Test_PolymorphismFFI::Test_PolymorphismFFI_FFI::polyLoop(LrcPtr::new(PureScript_Test_PolymorphismFFI::Test_PolymorphismFFI_FFI::Monoidish{mempty:
                                                                                                                                                                      1_i32,
                                                                                                                                                                  mappend:
                                                                                                                                                                      Func2::new(move
                                                                                                                                                                                     |x:
                                                                                                                                                                                          i32,
                                                                                                                                                                                      y:
                                                                                                                                                                                          i32|
                                                                                                                                                                                     x
                                                                                                                                                                                         +
                                                                                                                                                                                         y),}),
                                                                                 Sharpurs_Prelude::unbox(n),
                                                                                 0_i32)
        }
    }
    pub fn Test_PolymorphismFFI_runPolymorphismFFI() -> &dyn Any {
        static Test_PolymorphismFFI_runPolymorphismFFI:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PolymorphismFFI_runPolymorphismFFI.get_or_init(||
                                                                &Func1::new(move
                                                                                |arg0|
                                                                                &PureScript_Test_PolymorphismFFI::Test_PolymorphismFFI_FFI::runPolymorphismFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_PolymorphismFFI_describe() -> &dyn Any {
        static Test_PolymorphismFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PolymorphismFFI_describe.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                       &&&string("Polymorphism FFI (10M Type Class Dict Lookups):")))
    }
    pub fn Test_PolymorphismFFI_act() -> &dyn Any {
        static Test_PolymorphismFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PolymorphismFFI_act.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                        &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                        &&&10000000_i32)),
                                                                                  &&&Func1::new(move
                                                                                                    |dummy|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                        &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                           &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_PolymorphismFFI::Test_PolymorphismFFI_runPolymorphismFFI(),
                                                                                                                                                                                                           dummy))))))
    }
}
