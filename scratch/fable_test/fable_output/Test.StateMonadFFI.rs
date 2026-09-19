pub mod PureScript_Test_StateMonadFFI {
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
    pub mod Test_StateMonadFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::Func2;
        use fable_library_rust::Native_::LrcPtr;
        #[derive(Clone, Debug,)]
        pub enum State_2<s: Clone + 'static, a: Clone + 'static> {
            State(Func1<s, LrcPtr<(a, s)>>),
        }
        impl <s: Clone + 'static, a: Clone + 'static> core::fmt::Display for
         PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<s, a>
         {
            fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
                write!(f, "{}", core::any::type_name::<Self>())
            }
        }
        pub fn runState<a: Clone + 'static, b: Clone +
                        'static>(_arg:
                                     LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<a,
                                                                                                           b>>,
                                 s: a) -> LrcPtr<(b, a)> {
            (match _arg.as_ref() {
                 PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2::State(x)
                 => x.clone(),
             })(s)
        }
        pub fn bind<a: Clone + 'static, b: Clone + 'static, c: Clone +
                    'static>(_arg:
                                 LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<a,
                                                                                                       b>>,
                             g:
                                 Func1<b,
                                       LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<a,
                                                                                                             c>>>)
         ->
             LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<a,
                                                                                   c>> {
            LrcPtr::new(PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2::State(Func1::new({
                                                                                                             let _arg
                                                                                                                 =
                                                                                                                 _arg.clone();
                                                                                                             let g
                                                                                                                 =
                                                                                                                 g.clone();
                                                                                                             move
                                                                                                                 |s:
                                                                                                                      a|
                                                                                                                 {
                                                                                                                     let patternInput:
                                                                                                                             LrcPtr<(b,
                                                                                                                                     a)> =
                                                                                                                         (match _arg.as_ref()
                                                                                                                              {
                                                                                                                              PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2::State(x)
                                                                                                                              =>
                                                                                                                              x.clone(),
                                                                                                                          })(s);
                                                                                                                     PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::runState(g(patternInput.0.clone()),
                                                                                                                                                                                     patternInput.1.clone())
                                                                                                                 }
                                                                                                         })))
        }
        pub fn pure_<a: Clone + 'static, b: Clone + 'static>(a: a)
         ->
             LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<b,
                                                                                   a>> {
            LrcPtr::new(PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2::State(Func1::new({
                                                                                                             let a
                                                                                                                 =
                                                                                                                 a.clone();
                                                                                                             move
                                                                                                                 |s:
                                                                                                                      b|
                                                                                                                 LrcPtr::new((a.clone(),
                                                                                                                              s))
                                                                                                         })))
        }
        pub fn get<a: Clone + 'static>()
         ->
             LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<a,
                                                                                   a>> {
            LrcPtr::new(PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2::State(Func1::new(move
                                                                                                             |s:
                                                                                                                  a|
                                                                                                             LrcPtr::new((s.clone(),
                                                                                                                          s)))))
        }
        pub fn put<a: Clone + 'static>(s: a)
         ->
             LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<a,
                                                                                   ()>> {
            LrcPtr::new(PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2::State(Func1::new({
                                                                                                             let s
                                                                                                                 =
                                                                                                                 s.clone();
                                                                                                             move
                                                                                                                 |_arg:
                                                                                                                      a|
                                                                                                                 LrcPtr::new(((),
                                                                                                                              s.clone()))
                                                                                                         })))
        }
        pub fn modify<a: Clone + 'static>(f: Func1<a, a>)
         ->
             LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<a,
                                                                                   ()>> {
            PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::bind(PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::get(),
                                                                        Func1::new({
                                                                                       let f
                                                                                           =
                                                                                           f.clone();
                                                                                       move
                                                                                           |s:
                                                                                                a|
                                                                                           PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::put(f(s))
                                                                                   }))
        }
        pub fn chain(n: i32)
         ->
             LrcPtr<PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::State_2<i32,
                                                                                   ()>> {
            if n == 0_i32 {
                PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::pure_()
            } else {
                PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::bind(PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::modify(Func1::new(move
                                                                                                                                                         |x:
                                                                                                                                                              i32|
                                                                                                                                                         x
                                                                                                                                                             +
                                                                                                                                                             1_i32)),
                                                                            Func1::new({
                                                                                           let n
                                                                                               =
                                                                                               n.clone();
                                                                                           move
                                                                                               |_arg:
                                                                                                    ()|
                                                                                               PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::chain(n
                                                                                                                                                                -
                                                                                                                                                                1_i32)
                                                                                       }))
            }
        }
        pub fn runStateMonadFFI(n: &dyn Any) -> &dyn Any {
            let r#loop =
                Func2::new({
                               let n = n.clone();
                               move |i: i32, acc: i32|
                                   {
                                       let i: MutCell<i32> = MutCell::new(i);
                                       let acc: MutCell<i32> =
                                           MutCell::new(acc);
                                       '_loop:
                                           loop  {
                                               break '_loop
                                                   (if i.get() == 0_i32 {
                                                        acc.get()
                                                    } else {
                                                        let i_temp: i32 =
                                                            i.get() - 1_i32;
                                                        let acc_temp: i32 =
                                                            acc.get() +
                                                                (PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::runState(PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::chain(Sharpurs_Prelude::unbox(&n)),
                                                                                                                                 0_i32)).1.clone();
                                                        i.set(i_temp);
                                                        acc.set(acc_temp);
                                                        continue '_loop
                                                    }) ;
                                           }
                                   }
                           });
            &r#loop(20_i32, 0_i32)
        }
    }
    pub fn Test_StateMonadFFI_runStateMonadFFI() -> &dyn Any {
        static Test_StateMonadFFI_runStateMonadFFI: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Test_StateMonadFFI_runStateMonadFFI.get_or_init(||
                                                            &Func1::new(move
                                                                            |arg0|
                                                                            &PureScript_Test_StateMonadFFI::Test_StateMonadFFI_FFI::runStateMonadFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_StateMonadFFI_describe() -> &dyn Any {
        static Test_StateMonadFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonadFFI_describe.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                     &&&string("State Monad FFI (1.2k Binds, 60 Stack Depth):")))
    }
    pub fn Test_StateMonadFFI_act() -> &dyn Any {
        static Test_StateMonadFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_StateMonadFFI_act.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                      &&&PureScript_Effect::Effect_bindEffect()),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Bench::Bench_opaque(),
                                                                                                                                                      &&&60_i32)),
                                                                                &&&Func1::new(move
                                                                                                  |dummy|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                      &&&PureScript_Effect::Effect_applicativeEffect()),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                         &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_StateMonadFFI::Test_StateMonadFFI_runStateMonadFFI(),
                                                                                                                                                                                                         dummy))))))
    }
}
