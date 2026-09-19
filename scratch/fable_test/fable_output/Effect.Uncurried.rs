pub mod PureScript_Effect_Uncurried {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Effect_Uncurried_FFI {
        use super::*;
        use fable_library_rust::Exception_::try_catch;
        use fable_library_rust::Native_::LrcPtr;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::System::Exception;
        use fable_library_rust::System::InvalidCastException;
        pub fn mkEffectFn1() -> &dyn Any {
            static mkEffectFn1: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            mkEffectFn1.get_or_init(||
                                        &Func1::new(move |f|
                                                        &Func1::new({
                                                                        let f
                                                                            =
                                                                            f.clone();
                                                                        move
                                                                            |a|
                                                                            Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                               a),
                                                                                                             &defaultOf())
                                                                    })))
        }
        pub fn mkEffectFn2() -> &dyn Any {
            static mkEffectFn2: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            mkEffectFn2.get_or_init(||
                                        &Func1::new(move |f|
                                                        &Func1::new({
                                                                        let f
                                                                            =
                                                                            f.clone();
                                                                        move
                                                                            |a|
                                                                            &Func1::new({
                                                                                            let a
                                                                                                =
                                                                                                a.clone();
                                                                                            move
                                                                                                |b|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                     &a),
                                                                                                                                                                   b),
                                                                                                                                 &defaultOf())
                                                                                        })
                                                                    })))
        }
        pub fn mkEffectFn3() -> &dyn Any {
            static mkEffectFn3: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            mkEffectFn3.get_or_init(||
                                        &Func1::new(move |f|
                                                        &Func1::new({
                                                                        let f
                                                                            =
                                                                            f.clone();
                                                                        move
                                                                            |a|
                                                                            &Func1::new({
                                                                                            let a
                                                                                                =
                                                                                                a.clone();
                                                                                            move
                                                                                                |b|
                                                                                                &Func1::new({
                                                                                                                let b
                                                                                                                    =
                                                                                                                    b.clone();
                                                                                                                move
                                                                                                                    |c|
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                                                                           &a),
                                                                                                                                                                                                                         &b),
                                                                                                                                                                                       c),
                                                                                                                                                     &defaultOf())
                                                                                                            })
                                                                                        })
                                                                    })))
        }
        pub fn mkEffectFn4() -> &dyn Any {
            static mkEffectFn4: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            mkEffectFn4.get_or_init(||
                                        &Func1::new(move |f|
                                                        &Func1::new({
                                                                        let f
                                                                            =
                                                                            f.clone();
                                                                        move
                                                                            |a|
                                                                            &Func1::new({
                                                                                            let a
                                                                                                =
                                                                                                a.clone();
                                                                                            move
                                                                                                |b|
                                                                                                &Func1::new({
                                                                                                                let b
                                                                                                                    =
                                                                                                                    b.clone();
                                                                                                                move
                                                                                                                    |c|
                                                                                                                    &Func1::new({
                                                                                                                                    let c
                                                                                                                                        =
                                                                                                                                        c.clone();
                                                                                                                                    move
                                                                                                                                        |d|
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(&f,
                                                                                                                                                                                                                                                                                                                 &a),
                                                                                                                                                                                                                                                                               &b),
                                                                                                                                                                                                                                             &c),
                                                                                                                                                                                                           d),
                                                                                                                                                                         &defaultOf())
                                                                                                                                })
                                                                                                            })
                                                                                        })
                                                                    })))
        }
        pub fn mkEffectFn5<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn mkEffectFn6<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn mkEffectFn7<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn mkEffectFn8<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn mkEffectFn9<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn mkEffectFn10<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn runEffectFn1() -> &dyn Any {
            static runEffectFn1: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            runEffectFn1.get_or_init(||
                                         &Func1::new(move |eff|
                                                         &Func1::new({
                                                                         let eff
                                                                             =
                                                                             eff.clone();
                                                                         move
                                                                             |a|
                                                                             &Func1::new({
                                                                                             let a
                                                                                                 =
                                                                                                 a.clone();
                                                                                             move
                                                                                                 |_arg|
                                                                                                 eff(a)
                                                                                         })
                                                                     })))
        }
        pub fn runEffectFn2() -> &dyn Any {
            static runEffectFn2: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            runEffectFn2.get_or_init(||
                                         &Func1::new(move |eff|
                                                         &Func1::new({
                                                                         let eff
                                                                             =
                                                                             eff.clone();
                                                                         move
                                                                             |a|
                                                                             &Func1::new({
                                                                                             let a
                                                                                                 =
                                                                                                 a.clone();
                                                                                             move
                                                                                                 |b|
                                                                                                 &Func1::new({
                                                                                                                 let b
                                                                                                                     =
                                                                                                                     b.clone();
                                                                                                                 move
                                                                                                                     |_arg|
                                                                                                                     try_catch(||
                                                                                                                                   eff(a)(b),
                                                                                                                               |matchValue:
                                                                                                                                    LrcPtr<Exception>|
                                                                                                                                   if let Some(matchValue)
                                                                                                                                          =
                                                                                                                                          (matchValue
                                                                                                                                               as
                                                                                                                                               &dyn Any).downcast_ref::<LrcPtr<InvalidCastException>>()
                                                                                                                                      {
                                                                                                                                       eff(a)(b)
                                                                                                                                   } else {
                                                                                                                                       panic!("{}",
                                                                                                                                              matchValue.get_Message(),)
                                                                                                                                   })
                                                                                                             })
                                                                                         })
                                                                     })))
        }
        pub fn runEffectFn3() -> &dyn Any {
            static runEffectFn3: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            runEffectFn3.get_or_init(||
                                         &Func1::new(move |eff|
                                                         &Func1::new({
                                                                         let eff
                                                                             =
                                                                             eff.clone();
                                                                         move
                                                                             |a|
                                                                             &Func1::new({
                                                                                             let a
                                                                                                 =
                                                                                                 a.clone();
                                                                                             move
                                                                                                 |b|
                                                                                                 &Func1::new({
                                                                                                                 let b
                                                                                                                     =
                                                                                                                     b.clone();
                                                                                                                 move
                                                                                                                     |c|
                                                                                                                     &Func1::new({
                                                                                                                                     let c
                                                                                                                                         =
                                                                                                                                         c.clone();
                                                                                                                                     move
                                                                                                                                         |_arg|
                                                                                                                                         eff(a)(b)(c)
                                                                                                                                 })
                                                                                                             })
                                                                                         })
                                                                     })))
        }
        pub fn runEffectFn4() -> &dyn Any {
            static runEffectFn4: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            runEffectFn4.get_or_init(||
                                         &Func1::new(move |eff|
                                                         &Func1::new({
                                                                         let eff
                                                                             =
                                                                             eff.clone();
                                                                         move
                                                                             |a|
                                                                             &Func1::new({
                                                                                             let a
                                                                                                 =
                                                                                                 a.clone();
                                                                                             move
                                                                                                 |b|
                                                                                                 &Func1::new({
                                                                                                                 let b
                                                                                                                     =
                                                                                                                     b.clone();
                                                                                                                 move
                                                                                                                     |c|
                                                                                                                     &Func1::new({
                                                                                                                                     let c
                                                                                                                                         =
                                                                                                                                         c.clone();
                                                                                                                                     move
                                                                                                                                         |d|
                                                                                                                                         &Func1::new({
                                                                                                                                                         let d
                                                                                                                                                             =
                                                                                                                                                             d.clone();
                                                                                                                                                         move
                                                                                                                                                             |_arg|
                                                                                                                                                             eff(a)(b)(c)(d)
                                                                                                                                                     })
                                                                                                                                 })
                                                                                                             })
                                                                                         })
                                                                     })))
        }
        pub fn runEffectFn5<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn runEffectFn6<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn runEffectFn7<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn runEffectFn8<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn runEffectFn9<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
        pub fn runEffectFn10<a: Clone + 'static>(_arg: a) -> &dyn Any {
            Sharpurs_Prelude::undefined()
        }
    }
    pub fn Effect_Uncurried_mkEffectFn1() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn1.get_or_init(||
                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn1())
    }
    pub fn Effect_Uncurried_mkEffectFn10() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn10.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn10(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_mkEffectFn2() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn2.get_or_init(||
                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn2())
    }
    pub fn Effect_Uncurried_mkEffectFn3() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn3.get_or_init(||
                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn3())
    }
    pub fn Effect_Uncurried_mkEffectFn4() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn4.get_or_init(||
                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn4())
    }
    pub fn Effect_Uncurried_mkEffectFn5() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn5.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn5(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_mkEffectFn6() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn6.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn6(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_mkEffectFn7() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn7.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn7(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_mkEffectFn8() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn8.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn8(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_mkEffectFn9() -> &dyn Any {
        static Effect_Uncurried_mkEffectFn9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_mkEffectFn9.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::mkEffectFn9(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_runEffectFn1() -> &dyn Any {
        static Effect_Uncurried_runEffectFn1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn1.get_or_init(||
                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn1())
    }
    pub fn Effect_Uncurried_runEffectFn10() -> &dyn Any {
        static Effect_Uncurried_runEffectFn10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn10.get_or_init(||
                                                       &Func1::new(move |arg0|
                                                                       &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn10(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_runEffectFn2() -> &dyn Any {
        static Effect_Uncurried_runEffectFn2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn2.get_or_init(||
                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn2())
    }
    pub fn Effect_Uncurried_runEffectFn3() -> &dyn Any {
        static Effect_Uncurried_runEffectFn3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn3.get_or_init(||
                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn3())
    }
    pub fn Effect_Uncurried_runEffectFn4() -> &dyn Any {
        static Effect_Uncurried_runEffectFn4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn4.get_or_init(||
                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn4())
    }
    pub fn Effect_Uncurried_runEffectFn5() -> &dyn Any {
        static Effect_Uncurried_runEffectFn5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn5.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn5(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_runEffectFn6() -> &dyn Any {
        static Effect_Uncurried_runEffectFn6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn6.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn6(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_runEffectFn7() -> &dyn Any {
        static Effect_Uncurried_runEffectFn7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn7.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn7(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_runEffectFn8() -> &dyn Any {
        static Effect_Uncurried_runEffectFn8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn8.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn8(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_runEffectFn9() -> &dyn Any {
        static Effect_Uncurried_runEffectFn9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_runEffectFn9.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Effect_Uncurried::Effect_Uncurried_FFI::runEffectFn9(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Uncurried_semigroupEffectFn9() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn9: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn9.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn9(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                        let a
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            a.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                            let b
                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                b.clone();
                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                |c|
                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                let c
                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                    c.clone();
                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                    |d|
                                                                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                                                                    let d
                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                        d.clone();
                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                        |e|
                                                                                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                                                                                        let e
                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                            e.clone();
                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                            |f|
                                                                                                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                                                                                                            let f
                                                                                                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                                                                                                f.clone();
                                                                                                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                                                                                                |g|
                                                                                                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                let g
                                                                                                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                                                                                                    g.clone();
                                                                                                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                                                                                                    |h|
                                                                                                                                                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                    let h
                                                                                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                                                                                        h.clone();
                                                                                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                                                                                        |i|
                                                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn9(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&g),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&h),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               i)),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn9(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&f2),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&g),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&h),
                                                                                                                                                                                                                                                                                                                                                                                                                                                            i))
                                                                                                                                                                                                                                                                                                                                                                                })
                                                                                                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                                                                                })
                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn8() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn8: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn8.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn8(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                        let a
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            a.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                            let b
                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                b.clone();
                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                |c|
                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                let c
                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                    c.clone();
                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                    |d|
                                                                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                                                                    let d
                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                        d.clone();
                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                        |e|
                                                                                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                                                                                        let e
                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                            e.clone();
                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                            |f|
                                                                                                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                                                                                                            let f
                                                                                                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                                                                                                f.clone();
                                                                                                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                                                                                                |g|
                                                                                                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                let g
                                                                                                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                                                                                                    g.clone();
                                                                                                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                                                                                                    |h|
                                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn8(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&g),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                           h)),
                                                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn8(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&f2),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&g),
                                                                                                                                                                                                                                                                                                                                                                                                                                        h))
                                                                                                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                                                                                })
                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn7() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn7: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn7.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn7(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                        let a
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            a.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                            let b
                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                b.clone();
                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                |c|
                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                let c
                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                    c.clone();
                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                    |d|
                                                                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                                                                    let d
                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                        d.clone();
                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                        |e|
                                                                                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                                                                                        let e
                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                            e.clone();
                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                            |f|
                                                                                                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                                                                                                            let f
                                                                                                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                                                                                                f.clone();
                                                                                                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                                                                                                |g|
                                                                                                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn7(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       g)),
                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn7(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&f2),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                    g))
                                                                                                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                                                                                })
                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn6() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn6: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn6.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn6(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                        let a
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            a.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                            let b
                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                b.clone();
                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                |c|
                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                let c
                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                    c.clone();
                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                    |d|
                                                                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                                                                    let d
                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                        d.clone();
                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                        |e|
                                                                                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                                                                                        let e
                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                            e.clone();
                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                            |f|
                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn6(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                   f)),
                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn6(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&f2),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                f))
                                                                                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                                                                                })
                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn5() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn5: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn5.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn5(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                        let a
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            a.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                            let b
                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                b.clone();
                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                |c|
                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                let c
                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                    c.clone();
                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                    |d|
                                                                                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                                                                                    let d
                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                        d.clone();
                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                        |e|
                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                               &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn5(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                               e)),
                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn5(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&f2),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                               &&&d),
                                                                                                                                                                                                                                                                                                                                                                            e))
                                                                                                                                                                                                                                                                                                })
                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn4() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn4: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn4.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn4(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                        let a
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            a.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                            let b
                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                b.clone();
                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                |c|
                                                                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                                                                let c
                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                    c.clone();
                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                    |d|
                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                           &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn4(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                              &&&c),
                                                                                                                                                                                                                                                                                                                                                                                           d)),
                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn4(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&f2),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                              &&&b),
                                                                                                                                                                                                                                                                                                                                                                                           &&&c),
                                                                                                                                                                                                                                                                                                                                                        d))
                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn3() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn3: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn3.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn3(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                        let a
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            a.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                                                                                            let b
                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                b.clone();
                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                |c|
                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                       &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn3(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&b),
                                                                                                                                                                                                                                                                                                                                                                       c)),
                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn3(),
                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&f2),
                                                                                                                                                                                                                                                                                                                                                                                                          &&&a),
                                                                                                                                                                                                                                                                                                                                                                       &&&b),
                                                                                                                                                                                                                                                                                                                                    c))
                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn2() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn2: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn2.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn2(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                        let a
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            a.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |b|
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                   &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn2(),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                      &&&a),
                                                                                                                                                                                                                                                                                                                                                   b)),
                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn2(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&f2),
                                                                                                                                                                                                                                                                                                                                                   &&&a),
                                                                                                                                                                                                                                                                                                                b))
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn10() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn10: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn10.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictSemigroup|
                                                                             {
                                                                                 let semigroupEffect =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                      dictSemigroup);
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                  &&&add(string("append"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffect
                                                                                                                                              =
                                                                                                                                              semigroupEffect.clone();
                                                                                                                                          move
                                                                                                                                              |f1|
                                                                                                                                              &Func1::new({
                                                                                                                                                              let f1
                                                                                                                                                                  =
                                                                                                                                                                  f1.clone();
                                                                                                                                                              move
                                                                                                                                                                  |f2|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn10(),
                                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                                     let f2
                                                                                                                                                                                                                         =
                                                                                                                                                                                                                         f2.clone();
                                                                                                                                                                                                                     move
                                                                                                                                                                                                                         |a|
                                                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                                                         let a
                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                             a.clone();
                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                             |b|
                                                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                                                             let b
                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                 b.clone();
                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                 |c|
                                                                                                                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                                                                                                                 let c
                                                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                                                     c.clone();
                                                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                                                     |d|
                                                                                                                                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                                                                                                                                     let d
                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                         d.clone();
                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                         |e|
                                                                                                                                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                                                                                                                                         let e
                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                             e.clone();
                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                             |f|
                                                                                                                                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                                                                                                                                             let f
                                                                                                                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                                                                                                                 f.clone();
                                                                                                                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                                                                                                                 |g|
                                                                                                                                                                                                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                 let g
                                                                                                                                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                                                                                                                                     g.clone();
                                                                                                                                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                                                                                                                                     |h|
                                                                                                                                                                                                                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                     let h
                                                                                                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                                                                                                         h.clone();
                                                                                                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                                                                                                         |i|
                                                                                                                                                                                                                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                         let i
                                                                                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                                                                                             i.clone();
                                                                                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                                                                                             |j|
                                                                                                                                                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&semigroupEffect),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn10(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&f1),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&g),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&h),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&i),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    j)),
                                                                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn10(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&f2),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &&&a),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&b),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&c),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&d),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &&&e),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&f),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&g),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&h),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&i),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 j))
                                                                                                                                                                                                                                                                                                                                                                                                     })
                                                                                                                                                                                                                                                                                                                                                                                 })
                                                                                                                                                                                                                                                                                                                                                             })
                                                                                                                                                                                                                                                                                                                                         })
                                                                                                                                                                                                                                                                                                                     })
                                                                                                                                                                                                                                                                                                 })
                                                                                                                                                                                                                                                                             })
                                                                                                                                                                                                                                                         })
                                                                                                                                                                                                                                     })
                                                                                                                                                                                                                 }))
                                                                                                                                                          })
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))
                                                                             }))
    }
    pub fn Effect_Uncurried_semigroupEffectFn1() -> &dyn Any {
        static Effect_Uncurried_semigroupEffectFn1: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Uncurried_semigroupEffectFn1.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemigroup|
                                                                            {
                                                                                let semigroupEffect =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_semigroupEffect(),
                                                                                                                     dictSemigroup);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                 &&&add(string("append"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let semigroupEffect
                                                                                                                                             =
                                                                                                                                             semigroupEffect.clone();
                                                                                                                                         move
                                                                                                                                             |f1|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let f1
                                                                                                                                                                 =
                                                                                                                                                                 f1.clone();
                                                                                                                                                             move
                                                                                                                                                                 |f2|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn1(),
                                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                                    let f2
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        f2.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |a|
                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                               &&&semigroupEffect),
                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn1(),
                                                                                                                                                                                                                                                                                                                                                                  &&&f1),
                                                                                                                                                                                                                                                                                                                               a)),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_runEffectFn1(),
                                                                                                                                                                                                                                                                                                                               &&&f2),
                                                                                                                                                                                                                                                                                            a))
                                                                                                                                                                                                                }))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>()))
                                                                            }))
    }
    pub fn Effect_Uncurried_monoidEffectFn9() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn9.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn91 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn9(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn9(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                             |v2|
                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                             |v3|
                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                             |v4|
                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                             |v5|
                                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                                             |v6|
                                                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                                                             |v7|
                                                                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                                                                             |v8|
                                                                                                                                                                                                                                                                                                             &mempty))))))))
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn91
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn91.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn91
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Uncurried_monoidEffectFn8() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn8.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn81 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn8(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn8(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                             |v2|
                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                             |v3|
                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                             |v4|
                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                             |v5|
                                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                                             |v6|
                                                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                                                             |v7|
                                                                                                                                                                                                                                                                                             &mempty)))))))
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn81
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn81.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn81
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Uncurried_monoidEffectFn7() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn7.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn71 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn7(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn7(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                             |v2|
                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                             |v3|
                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                             |v4|
                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                             |v5|
                                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                                             |v6|
                                                                                                                                                                                                                                                                             &mempty))))))
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn71
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn71.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn71
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Uncurried_monoidEffectFn6() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn6.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn61 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn6(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn6(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                             |v2|
                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                             |v3|
                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                             |v4|
                                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                                             |v5|
                                                                                                                                                                                                                                                             &mempty)))))
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn61
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn61.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn61
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Uncurried_monoidEffectFn5() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn5.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn51 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn5(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn5(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                             |v2|
                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                             |v3|
                                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                                             |v4|
                                                                                                                                                                                                                                             &mempty))))
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn51
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn51.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn51
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Uncurried_monoidEffectFn4() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn4.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn41 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn4(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn4(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                             |v2|
                                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                                             |v3|
                                                                                                                                                                                                                             &mempty)))
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn41
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn41.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn41
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Uncurried_monoidEffectFn3() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn3.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn31 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn3(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn3(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                                             |v2|
                                                                                                                                                                                                             &mempty))
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn31
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn31.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn31
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Uncurried_monoidEffectFn2() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn2.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn21 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn2(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn2(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &Func1::new(move
                                                                                                                                                                                             |v1|
                                                                                                                                                                                             &mempty)
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn21
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn21.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn21
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Effect_Uncurried_monoidEffectFn10() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn10.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonoid|
                                                                          {
                                                                              let mempty =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                      dictMonoid));
                                                                              let semigroupEffectFn101 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn10(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                               &&&add(string("mempty"),
                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn10(),
                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                          let mempty
                                                                                                                                                                              =
                                                                                                                                                                              mempty.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v|
                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                              |v1|
                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                              |v2|
                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                              |v3|
                                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                                              |v4|
                                                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                                                              |v5|
                                                                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                                                                              |v6|
                                                                                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                                                                                              |v7|
                                                                                                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                                                                                                              |v8|
                                                                                                                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                                                                                                                              |v9|
                                                                                                                                                                                                                                                                                                                              &mempty)))))))))
                                                                                                                                                                      })),
                                                                                                                      add(string("Semigroup0"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let semigroupEffectFn101
                                                                                                                                               =
                                                                                                                                               semigroupEffectFn101.clone();
                                                                                                                                           move
                                                                                                                                               |usd__unused|
                                                                                                                                               &semigroupEffectFn101
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>())))
                                                                          }))
    }
    pub fn Effect_Uncurried_monoidEffectFn1() -> &dyn Any {
        static Effect_Uncurried_monoidEffectFn1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Uncurried_monoidEffectFn1.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonoid|
                                                                         {
                                                                             let mempty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect::Effect_monoidEffect(),
                                                                                                                                                     dictMonoid));
                                                                             let semigroupEffectFn11 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_semigroupEffectFn1(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                              &&&add(string("mempty"),
                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Effect_Uncurried::Effect_Uncurried_mkEffectFn1(),
                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                         let mempty
                                                                                                                                                                             =
                                                                                                                                                                             mempty.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |v|
                                                                                                                                                                             &mempty
                                                                                                                                                                     })),
                                                                                                                     add(string("Semigroup0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let semigroupEffectFn11
                                                                                                                                              =
                                                                                                                                              semigroupEffectFn11.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &semigroupEffectFn11
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
}
