pub mod PureScript_Test_PrimesFFI {
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
    pub mod Test_PrimesFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::LrcPtr;
        #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
        pub enum List {
            Nil,
            Cons(i32,
                 LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List>),
        }
        impl core::fmt::Display for
         PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn range(s: i32, e: i32)
         -> LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List> {
            if s > e {
                LrcPtr::new(PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Nil)
            } else {
                LrcPtr::new(PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Cons(s,
                                                                                      PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::range(s
                                                                                                                                               +
                                                                                                                                               1_i32,
                                                                                                                                           e)))
            }
        }
        pub fn filter(p: Func1<i32, bool>,
                      _arg:
                          LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List>)
         -> LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List> {
            match _arg.as_ref() {
                PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Cons(_arg_1_0,
                                                                          _arg_1_1)
                => {
                    let xs:
                            LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List> =
                        _arg_1_1.clone();
                    let x: i32 = _arg_1_0.clone();
                    if p(x) {
                        LrcPtr::new(PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Cons(x,
                                                                                              PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::filter(p.clone(),
                                                                                                                                                    xs.clone())))
                    } else {
                        PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::filter(p.clone(),
                                                                              xs)
                    }
                }
                _ =>
                LrcPtr::new(PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Nil),
            }
        }
        pub fn sieve(_arg:
                         LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List>)
         -> LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List> {
            match _arg.as_ref() {
                PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Cons(_arg_1_0,
                                                                          _arg_1_1)
                => {
                    let p: i32 = _arg_1_0.clone();
                    LrcPtr::new(PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Cons(p,
                                                                                          PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::sieve(PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::filter(Func1::new({
                                                                                                                                                                                                                    let p
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        p.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |x:
                                                                                                                                                                                                                             i32|
                                                                                                                                                                                                                        x
                                                                                                                                                                                                                            %
                                                                                                                                                                                                                            p
                                                                                                                                                                                                                            !=
                                                                                                                                                                                                                            0_i32
                                                                                                                                                                                                                }),
                                                                                                                                                                                                     _arg_1_1.clone()))))
                }
                _ =>
                LrcPtr::new(PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Nil),
            }
        }
        pub fn sum(acc: i32,
                   _arg:
                       LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List>)
         -> i32 {
            let acc: MutCell<i32> = MutCell::new(acc);
            let _arg:
                    MutCell<LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List>> =
                MutCell::new(_arg.clone());
            '_sum:
                loop  {
                    break '_sum
                        (match _arg.get().as_ref() {
                             PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List::Cons(_arg_1_0,
                                                                                       _arg_1_1)
                             => {
                                 let acc_temp: i32 =
                                     acc.get() + _arg_1_0.clone();
                                 let _arg_temp:
                                         LrcPtr<PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::List> =
                                     _arg_1_1.clone();
                                 acc.set(acc_temp);
                                 _arg.set(_arg_temp);
                                 continue '_sum
                             }
                             _ => acc.get(),
                         }) ;
                }
        }
        pub fn runPrimesFFI(n: &dyn Any) -> &dyn Any {
            &PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::sum(0_i32,
                                                                PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::sieve(PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::range(2_i32,
                                                                                                                                                                          Sharpurs_Prelude::unbox(n))))
        }
    }
    pub fn Test_PrimesFFI_runPrimesFFI() -> &dyn Any {
        static Test_PrimesFFI_runPrimesFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PrimesFFI_runPrimesFFI.get_or_init(||
                                                    &Func1::new(move |arg0|
                                                                    &PureScript_Test_PrimesFFI::Test_PrimesFFI_FFI::runPrimesFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_PrimesFFI_describe() -> &dyn Any {
        static Test_PrimesFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PrimesFFI_describe.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                 &&&string("Prime Sieve FFI (sum primes up to 500):")))
    }
    pub fn Test_PrimesFFI_act() -> &dyn Any {
        static Test_PrimesFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_PrimesFFI_act.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                  &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                  &&&500_i32)),
                                                                            &&&Func1::new(move
                                                                                              |dummy|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                  &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                     &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_PrimesFFI::Test_PrimesFFI_runPrimesFFI(),
                                                                                                                                                                                                     dummy))))))
    }
}
