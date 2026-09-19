pub mod PureScript_Test_ChurchFFI {
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
    pub mod Test_ChurchFFI_FFI {
        use super::*;
        use fable_library_rust::Native_::Func2;
        pub fn fromInt(n: i32) -> Func1<Func1<i32, i32>, Func1<i32, i32>> {
            if n == 0_i32 {
                Func1::new(move |_arg: Func1<i32, i32>|
                               Func1::new(move |value: i32| value))
            } else {
                let previous =
                    PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::fromInt(n -
                                                                               1_i32);
                Func1::new({
                               let previous = previous.clone();
                               move |f: Func1<i32, i32>|
                                   Func1::new({
                                                  let f = f.clone();
                                                  let previous =
                                                      previous.clone();
                                                  move |value_1: i32|
                                                      f(previous(f.clone())(value_1))
                                              })
                           })
            }
        }
        pub fn multiply(m: Func2<Func1<i32, i32>, i32, i32>,
                        n: Func2<Func1<i32, i32>, i32, i32>,
                        f: Func1<i32, i32>, value: i32) -> i32 {
            m(Func1::new({
                             let f = f.clone();
                             let n = n.clone();
                             move |a0: i32| n(f.clone(), a0)
                         }), value)
        }
        pub fn square(n: i32) -> Func1<Func1<i32, i32>, Func1<i32, i32>> {
            let m = PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::fromInt(n);
            let n_1 =
                PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::fromInt(n);
            Func1::new({
                           let m = m.clone();
                           let n_1 = n_1.clone();
                           move |f: Func1<i32, i32>|
                               Func1::new({
                                              let f = f.clone();
                                              let m = m.clone();
                                              let n_1 = n_1.clone();
                                              move |value: i32|
                                                  PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::multiply(Func2::new({
                                                                                                                         let m
                                                                                                                             =
                                                                                                                             m.clone();
                                                                                                                         move
                                                                                                                             |b0:
                                                                                                                                  Func1<i32,
                                                                                                                                        i32>,
                                                                                                                              b1:
                                                                                                                                  i32|
                                                                                                                             m(b0)(b1)
                                                                                                                     }),
                                                                                                          Func2::new({
                                                                                                                         let n_1
                                                                                                                             =
                                                                                                                             n_1.clone();
                                                                                                                         move
                                                                                                                             |b0:
                                                                                                                                  Func1<i32,
                                                                                                                                        i32>,
                                                                                                                              b1:
                                                                                                                                  i32|
                                                                                                                             n_1(b0)(b1)
                                                                                                                     }),
                                                                                                          f.clone(),
                                                                                                          value)
                                          })
                       })
        }
        pub fn runChurchFFI(input: &dyn Any) -> &dyn Any {
            let n: i32 = Sharpurs_Prelude::unbox(input);
            &PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::multiply(Func2::new({
                                                                                    let n
                                                                                        =
                                                                                        n.clone();
                                                                                    move
                                                                                        |b0:
                                                                                             Func1<i32,
                                                                                                   i32>,
                                                                                         b1:
                                                                                             i32|
                                                                                        ({
                                                                                             let m =
                                                                                                 PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::square(n);
                                                                                             let n_1 =
                                                                                                 PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::square(n);
                                                                                             Func1::new({
                                                                                                            let m
                                                                                                                =
                                                                                                                m.clone();
                                                                                                            let n_1
                                                                                                                =
                                                                                                                n_1.clone();
                                                                                                            move
                                                                                                                |f:
                                                                                                                     Func1<i32,
                                                                                                                           i32>|
                                                                                                                Func1::new({
                                                                                                                               let f
                                                                                                                                   =
                                                                                                                                   f.clone();
                                                                                                                               let m
                                                                                                                                   =
                                                                                                                                   m.clone();
                                                                                                                               let n_1
                                                                                                                                   =
                                                                                                                                   n_1.clone();
                                                                                                                               move
                                                                                                                                   |value:
                                                                                                                                        i32|
                                                                                                                                   PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::multiply(Func2::new({
                                                                                                                                                                                                          let m
                                                                                                                                                                                                              =
                                                                                                                                                                                                              m.clone();
                                                                                                                                                                                                          move
                                                                                                                                                                                                              |b0:
                                                                                                                                                                                                                   Func1<i32,
                                                                                                                                                                                                                         i32>,
                                                                                                                                                                                                               b1:
                                                                                                                                                                                                                   i32|
                                                                                                                                                                                                              m(b0)(b1)
                                                                                                                                                                                                      }),
                                                                                                                                                                                           Func2::new({
                                                                                                                                                                                                          let n_1
                                                                                                                                                                                                              =
                                                                                                                                                                                                              n_1.clone();
                                                                                                                                                                                                          move
                                                                                                                                                                                                              |b0:
                                                                                                                                                                                                                   Func1<i32,
                                                                                                                                                                                                                         i32>,
                                                                                                                                                                                                               b1:
                                                                                                                                                                                                                   i32|
                                                                                                                                                                                                              n_1(b0)(b1)
                                                                                                                                                                                                      }),
                                                                                                                                                                                           f.clone(),
                                                                                                                                                                                           value)
                                                                                                                           })
                                                                                                        })
                                                                                         })(b0.clone())(b1)
                                                                                }),
                                                                     Func2::new({
                                                                                    let n
                                                                                        =
                                                                                        n.clone();
                                                                                    move
                                                                                        |b0:
                                                                                             Func1<i32,
                                                                                                   i32>,
                                                                                         b1:
                                                                                             i32|
                                                                                        (PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::fromInt(n))(b0)(b1)
                                                                                }),
                                                                     Func1::new(move
                                                                                    |y:
                                                                                         i32|
                                                                                    1_i32
                                                                                        +
                                                                                        y),
                                                                     0_i32)
        }
    }
    pub fn Test_ChurchFFI_runChurchFFI() -> &dyn Any {
        static Test_ChurchFFI_runChurchFFI: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ChurchFFI_runChurchFFI.get_or_init(||
                                                    &Func1::new(move |arg0|
                                                                    &PureScript_Test_ChurchFFI::Test_ChurchFFI_FFI::runChurchFFI(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Test_ChurchFFI_describe() -> &dyn Any {
        static Test_ChurchFFI_describe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ChurchFFI_describe.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Console::Effect_Console_log(),
                                                                                 &&&string("Church Numerals FFI (100k Closure Applications):")))
    }
    pub fn Test_ChurchFFI_act() -> &dyn Any {
        static Test_ChurchFFI_act: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Test_ChurchFFI_act.get_or_init(||
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
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Test_ChurchFFI::Test_ChurchFFI_runChurchFFI(),
                                                                                                                                                                                                     dummy))))))
    }
}
