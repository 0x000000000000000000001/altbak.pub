pub mod PureScript_Data_Either {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug,)]
    pub enum Data_Either_Either {
        Data_Either_Leftusd_Ctor(&dyn Any),
        Data_Either_Rightusd_Ctor(&dyn Any),
    }
    impl core::fmt::Display for PureScript_Data_Either::Data_Either_Either {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Either_Left() -> &dyn Any {
        static Data_Either_Left: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_Left.get_or_init(||
                                         &Func1::new(move |usd__arg1|
                                                         &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Either_Right() -> &dyn Any {
        static Data_Either_Right: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_Right.get_or_init(||
                                          &Func1::new(move |usd__arg1|
                                                          &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Either_showEither() -> &dyn Any {
        static Data_Either_showEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_showEither.get_or_init(||
                                               &Func1::new(move |dictShow|
                                                               &Func1::new({
                                                                               let dictShow
                                                                                   =
                                                                                   dictShow.clone();
                                                                               move
                                                                                   |dictShow1|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                    &&&add(string("show"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let dictShow1
                                                                                                                                                =
                                                                                                                                                dictShow1.clone();
                                                                                                                                            move
                                                                                                                                                |v|
                                                                                                                                                {
                                                                                                                                                    let matchValue:
                                                                                                                                                            LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                                                    match matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                        PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                        =>
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                            &&&string("(Right ")),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                     &&&dictShow1),
                                                                                                                                                                                                                                                                                                  &&matchValue_1_0)),
                                                                                                                                                                                                                            &&&string(")"))),
                                                                                                                                                        PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)
                                                                                                                                                        =>
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                            &&&string("(Left ")),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                     &&&dictShow),
                                                                                                                                                                                                                                                                                                  &&matchValue_0_0)),
                                                                                                                                                                                                                            &&&string(")"))),
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))
                                                                           })))
    }
    pub fn Data_Either_note_prime() -> &dyn Any {
        static Data_Either_note_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_note_prime.get_or_init(||
                                               &Func1::new(move |f|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe_prime(),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                           |usd__arg1|
                                                                                                                                                                                                                           &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                      f)),
                                                                                                &&&Func1::new(move
                                                                                                                  |usd__arg1_1|
                                                                                                                  &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone()))))))
    }
    pub fn Data_Either_note() -> &dyn Any {
        static Data_Either_note: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_note.get_or_init(||
                                         &Func1::new(move |a|
                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                             &&&LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(a.clone()))),
                                                                                          &&&Func1::new(move
                                                                                                            |usd__arg1|
                                                                                                            &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone()))))))
    }
    pub fn Data_Either_genericEither() -> &dyn Any {
        static Data_Either_genericEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_genericEither.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Genericusd_Dict(),
                                                                                   &&&add(string("to"),
                                                                                          &&Func1::new(move
                                                                                                           |x|
                                                                                                           {
                                                                                                               let matchValue:
                                                                                                                       LrcPtr<Data_Generic_Rep_Sum> =
                                                                                                                   Sharpurs_Prelude::unbox(x);
                                                                                                               match matchValue.as_ref()
                                                                                                                   {
                                                                                                                   Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(matchValue_1_0)
                                                                                                                   =>
                                                                                                                   &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)),
                                                                                                                   Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(matchValue_0_0)
                                                                                                                   =>
                                                                                                                   &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)),
                                                                                                               }
                                                                                                           }),
                                                                                          add(string("from"),
                                                                                              &&Func1::new(move
                                                                                                               |x_1|
                                                                                                               {
                                                                                                                   let matchValue_1:
                                                                                                                           LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                       Sharpurs_Prelude::unbox(x_1);
                                                                                                                   match matchValue_1.as_ref()
                                                                                                                       {
                                                                                                                       PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                                                       =>
                                                                                                                       &LrcPtr::new(Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                                                                           &&matchValue_1_1_0)))),
                                                                                                                       PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                                       =>
                                                                                                                       &LrcPtr::new(Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                                                                           &&matchValue_1_0_0)))),
                                                                                                                   }
                                                                                                               }),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Either_functorEither() -> &dyn Any {
        static Data_Either_functorEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_functorEither.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                   &&&add(string("map"),
                                                                                          &&Func1::new(move
                                                                                                           |f|
                                                                                                           &Func1::new({
                                                                                                                           let f
                                                                                                                               =
                                                                                                                               f.clone();
                                                                                                                           move
                                                                                                                               |m|
                                                                                                                               {
                                                                                                                                   let matchValue:
                                                                                                                                           LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                       Sharpurs_Prelude::unbox(m);
                                                                                                                                   match matchValue.as_ref()
                                                                                                                                       {
                                                                                                                                       PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                       =>
                                                                                                                                       &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                           &&matchValue_1_0))),
                                                                                                                                       PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)
                                                                                                                                       =>
                                                                                                                                       &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)),
                                                                                                                                   }
                                                                                                                               }
                                                                                                                       })),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Either_invariantEither() -> &dyn Any {
        static Data_Either_invariantEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_invariantEither.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                     &&&add(string("imap"),
                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                              &&&PureScript_Data_Either::Data_Either_functorEither()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>())))
    }
    pub fn Data_Either_fromRight_prime() -> &dyn Any {
        static Data_Either_fromRight_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_fromRight_prime.get_or_init(||
                                                    &Func1::new(move |v|
                                                                    &Func1::new({
                                                                                    let v
                                                                                        =
                                                                                        v.clone();
                                                                                    move
                                                                                        |v1|
                                                                                        {
                                                                                            let matchValue =
                                                                                                Sharpurs_Prelude::unbox(&&v);
                                                                                            let matchValue_1:
                                                                                                    LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                Sharpurs_Prelude::unbox(v1);
                                                                                            if let PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                                   =
                                                                                                   matchValue_1.as_ref()
                                                                                               {
                                                                                                &match matchValue_1.as_ref()
                                                                                                     {
                                                                                                     PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                     =>
                                                                                                     x.clone(),
                                                                                                     _
                                                                                                     =>
                                                                                                     unreachable!(),
                                                                                                 }
                                                                                            } else {
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                 &&&PureScript_Data_Unit::Data_Unit_unit())
                                                                                            }
                                                                                        }
                                                                                })))
    }
    pub fn Data_Either_fromRight() -> &dyn Any {
        static Data_Either_fromRight: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_fromRight.get_or_init(||
                                              &Func1::new(move |v|
                                                              &Func1::new({
                                                                              let v
                                                                                  =
                                                                                  v.clone();
                                                                              move
                                                                                  |v1|
                                                                                  {
                                                                                      let matchValue =
                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                      let matchValue_1:
                                                                                              LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                      if let PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                             =
                                                                                             matchValue_1.as_ref()
                                                                                         {
                                                                                          &match matchValue_1.as_ref()
                                                                                               {
                                                                                               PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                               =>
                                                                                               x.clone(),
                                                                                               _
                                                                                               =>
                                                                                               unreachable!(),
                                                                                           }
                                                                                      } else {
                                                                                          &matchValue
                                                                                      }
                                                                                  }
                                                                          })))
    }
    pub fn Data_Either_fromLeft_prime() -> &dyn Any {
        static Data_Either_fromLeft_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_fromLeft_prime.get_or_init(||
                                                   &Func1::new(move |v|
                                                                   &Func1::new({
                                                                                   let v
                                                                                       =
                                                                                       v.clone();
                                                                                   move
                                                                                       |v1|
                                                                                       {
                                                                                           let matchValue =
                                                                                               Sharpurs_Prelude::unbox(&&v);
                                                                                           let matchValue_1:
                                                                                                   LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                           if let PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                  =
                                                                                                  matchValue_1.as_ref()
                                                                                              {
                                                                                               &match matchValue_1.as_ref()
                                                                                                    {
                                                                                                    PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                    =>
                                                                                                    x.clone(),
                                                                                                    _
                                                                                                    =>
                                                                                                    unreachable!(),
                                                                                                }
                                                                                           } else {
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit())
                                                                                           }
                                                                                       }
                                                                               })))
    }
    pub fn Data_Either_fromLeft() -> &dyn Any {
        static Data_Either_fromLeft: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_fromLeft.get_or_init(||
                                             &Func1::new(move |v|
                                                             &Func1::new({
                                                                             let v
                                                                                 =
                                                                                 v.clone();
                                                                             move
                                                                                 |v1|
                                                                                 {
                                                                                     let matchValue =
                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                     let matchValue_1:
                                                                                             LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                         Sharpurs_Prelude::unbox(v1);
                                                                                     if let PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                            =
                                                                                            matchValue_1.as_ref()
                                                                                        {
                                                                                         &match matchValue_1.as_ref()
                                                                                              {
                                                                                              PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                              =>
                                                                                              x.clone(),
                                                                                              _
                                                                                              =>
                                                                                              unreachable!(),
                                                                                          }
                                                                                     } else {
                                                                                         &matchValue
                                                                                     }
                                                                                 }
                                                                         })))
    }
    pub fn Data_Either_extendEither() -> &dyn Any {
        static Data_Either_extendEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_extendEither.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                  &&&add(string("extend"),
                                                                                         &&Func1::new(move
                                                                                                          |v|
                                                                                                          &Func1::new({
                                                                                                                          let v
                                                                                                                              =
                                                                                                                              v.clone();
                                                                                                                          move
                                                                                                                              |v1|
                                                                                                                              {
                                                                                                                                  let matchValue =
                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                  let matchValue_1:
                                                                                                                                          LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                  if let PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                                                         =
                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                     {
                                                                                                                                      &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                             {
                                                                                                                                                                                                                             PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                             _
                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                         }))
                                                                                                                                  } else {
                                                                                                                                      &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                          &&&matchValue_1)))
                                                                                                                                  }
                                                                                                                              }
                                                                                                                      })),
                                                                                         add(string("Functor0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &PureScript_Data_Either::Data_Either_functorEither()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Data_Either_eqEither() -> &dyn Any {
        static Data_Either_eqEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_eqEither.get_or_init(||
                                             &Func1::new(move |dictEq|
                                                             &Func1::new({
                                                                             let dictEq
                                                                                 =
                                                                                 dictEq.clone();
                                                                             move
                                                                                 |dictEq1|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                  &&&add(string("eq"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let dictEq1
                                                                                                                                              =
                                                                                                                                              dictEq1.clone();
                                                                                                                                          move
                                                                                                                                              |x|
                                                                                                                                              &Func1::new({
                                                                                                                                                              let x
                                                                                                                                                                  =
                                                                                                                                                                  x.clone();
                                                                                                                                                              move
                                                                                                                                                                  |y|
                                                                                                                                                                  {
                                                                                                                                                                      let matchValue:
                                                                                                                                                                              LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                      let matchValue_1:
                                                                                                                                                                              LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                                                          Sharpurs_Prelude::unbox(y);
                                                                                                                                                                      if let PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                                             =
                                                                                                                                                                             matchValue.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          if let PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                     &&&dictEq1),
                                                                                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                         PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                      {
                                                                                                                                                                                                                      PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                      _
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                  })
                                                                                                                                                                          } else {
                                                                                                                                                                              &false
                                                                                                                                                                          }
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                     &&&dictEq),
                                                                                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                         PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                      {
                                                                                                                                                                                                                      PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                      _
                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                  })
                                                                                                                                                                          } else {
                                                                                                                                                                              &false
                                                                                                                                                                          }
                                                                                                                                                                      }
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))
                                                                         })))
    }
    pub fn Data_Either_ordEither() -> &dyn Any {
        static Data_Either_ordEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_ordEither.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              {
                                                                  let eqEither1 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_eqEither(),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                 Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                  &Func1::new({
                                                                                  let dictOrd
                                                                                      =
                                                                                      dictOrd.clone();
                                                                                  let eqEither1
                                                                                      =
                                                                                      eqEither1.clone();
                                                                                  move
                                                                                      |dictOrd1|
                                                                                      {
                                                                                          let eqEither2 =
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&eqEither1,
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                         Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                           &&&add(string("compare"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let dictOrd1
                                                                                                                                                       =
                                                                                                                                                       dictOrd1.clone();
                                                                                                                                                   move
                                                                                                                                                       |x|
                                                                                                                                                       &Func1::new({
                                                                                                                                                                       let x
                                                                                                                                                                           =
                                                                                                                                                                           x.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |y|
                                                                                                                                                                           {
                                                                                                                                                                               let matchValue:
                                                                                                                                                                                       LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                               let matchValue_1:
                                                                                                                                                                                       LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(y);
                                                                                                                                                                               if let PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                                                      =
                                                                                                                                                                                      matchValue.as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                   if let PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                                          =
                                                                                                                                                                                          matchValue_1.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                              &&&dictOrd1),
                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                  PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                  _
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                               PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                               _
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                           })
                                                                                                                                                                                   } else {
                                                                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                   }
                                                                                                                                                                               } else {
                                                                                                                                                                                   if let PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_1_0_0)
                                                                                                                                                                                          =
                                                                                                                                                                                          matchValue_1.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                              &&&dictOrd),
                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                  PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                  _
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                               PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                               _
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                           })
                                                                                                                                                                                   } else {
                                                                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                   }
                                                                                                                                                                               }
                                                                                                                                                                           }
                                                                                                                                                                   })
                                                                                                                                               }),
                                                                                                                                  add(string("Eq0"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let eqEither2
                                                                                                                                                           =
                                                                                                                                                           eqEither2.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &eqEither2
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))
                                                                                      }
                                                                              })
                                                              }))
    }
    pub fn Data_Either_eq1Either() -> &dyn Any {
        static Data_Either_eq1Either: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_eq1Either.get_or_init(||
                                              &Func1::new(move |dictEq|
                                                              {
                                                                  let eqEither1 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_eqEither(),
                                                                                                       dictEq);
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                   &&&add(string("eq1"),
                                                                                                          &&Func1::new({
                                                                                                                           let eqEither1
                                                                                                                               =
                                                                                                                               eqEither1.clone();
                                                                                                                           move
                                                                                                                               |dictEq1|
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&eqEither1,
                                                                                                                                                                                                   dictEq1))
                                                                                                                       }),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))
                                                              }))
    }
    pub fn Data_Either_ord1Either() -> &dyn Any {
        static Data_Either_ord1Either: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_ord1Either.get_or_init(||
                                               &Func1::new(move |dictOrd|
                                                               {
                                                                   let ordEither1 =
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_ordEither(),
                                                                                                        dictOrd);
                                                                   let eq1Either1 =
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_eq1Either(),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                  Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                    &&&add(string("compare1"),
                                                                                                           &&Func1::new({
                                                                                                                            let ordEither1
                                                                                                                                =
                                                                                                                                ordEither1.clone();
                                                                                                                            move
                                                                                                                                |dictOrd1|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&ordEither1,
                                                                                                                                                                                                    dictOrd1))
                                                                                                                        }),
                                                                                                           add(string("Eq10"),
                                                                                                               &&Func1::new({
                                                                                                                                let eq1Either1
                                                                                                                                    =
                                                                                                                                    eq1Either1.clone();
                                                                                                                                move
                                                                                                                                    |usd__unused|
                                                                                                                                    &eq1Either1
                                                                                                                            }),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>())))
                                                               }))
    }
    pub fn Data_Either_either() -> &dyn Any {
        static Data_Either_either: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_either.get_or_init(||
                                           &Func1::new(move |v|
                                                           &Func1::new({
                                                                           let v
                                                                               =
                                                                               v.clone();
                                                                           move
                                                                               |v1|
                                                                               &Func1::new({
                                                                                               let v1
                                                                                                   =
                                                                                                   v1.clone();
                                                                                               move
                                                                                                   |v2|
                                                                                                   {
                                                                                                       let matchValue =
                                                                                                           Sharpurs_Prelude::unbox(&&v);
                                                                                                       let matchValue_1 =
                                                                                                           Sharpurs_Prelude::unbox(&&v1);
                                                                                                       let matchValue_2:
                                                                                                               LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                           Sharpurs_Prelude::unbox(v2);
                                                                                                       match matchValue_2.as_ref()
                                                                                                           {
                                                                                                           PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_2_1_0)
                                                                                                           =>
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                            &&matchValue_2_1_0),
                                                                                                           PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_2_0_0)
                                                                                                           =>
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                            &&matchValue_2_0_0),
                                                                                                       }
                                                                                                   }
                                                                                           })
                                                                       })))
    }
    pub fn Data_Either_hush() -> &dyn Any {
        static Data_Either_hush: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_hush.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))),
                                                                          &&&Func1::new(move
                                                                                            |usd__arg1|
                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))))
    }
    pub fn Data_Either_isLeft() -> &dyn Any {
        static Data_Either_isLeft: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_isLeft.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                  &&&true)),
                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                               &&&false)))
    }
    pub fn Data_Either_isRight() -> &dyn Any {
        static Data_Either_isRight: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_isRight.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                   &&&false)),
                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                &&&true)))
    }
    pub fn Data_Either_choose() -> &dyn Any {
        static Data_Either_choose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_choose.get_or_init(||
                                           &Func1::new(move |dictAlt|
                                                           {
                                                               let Functor0 =
                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                           Sharpurs_Prelude::unbox(dictAlt)),
                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                               &Func1::new({
                                                                               let Functor0
                                                                                   =
                                                                                   Functor0.clone();
                                                                               let dictAlt
                                                                                   =
                                                                                   dictAlt.clone();
                                                                               move
                                                                                   |a|
                                                                                   &Func1::new({
                                                                                                   let a
                                                                                                       =
                                                                                                       a.clone();
                                                                                                   move
                                                                                                       |b|
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                              &&&dictAlt),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                    &&&Functor0),
                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                   |usd__arg1|
                                                                                                                                                                                                                                                                   &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                              &&&a)),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                 &&&Functor0),
                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                |usd__arg1_1|
                                                                                                                                                                                                                                &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                           b))
                                                                                               })
                                                                           })
                                                           }))
    }
    pub fn Data_Either_boundedEither() -> &dyn Any {
        static Data_Either_boundedEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_boundedEither.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictBounded|
                                                                  {
                                                                      let ordEither1 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_ordEither(),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                     Sharpurs_Prelude::unbox(dictBounded)),
                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                      &Func1::new({
                                                                                      let dictBounded
                                                                                          =
                                                                                          dictBounded.clone();
                                                                                      let ordEither1
                                                                                          =
                                                                                          ordEither1.clone();
                                                                                      move
                                                                                          |dictBounded1|
                                                                                          {
                                                                                              let ordEither2 =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&ordEither1,
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                                             Sharpurs_Prelude::unbox(dictBounded1)),
                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                                                               &&&add(string("top"),
                                                                                                                                      &&LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                                                                           dictBounded1))),
                                                                                                                                      add(string("bottom"),
                                                                                                                                          &&LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                                                              &&&dictBounded))),
                                                                                                                                          add(string("Ord0"),
                                                                                                                                              &&Func1::new({
                                                                                                                                                               let ordEither2
                                                                                                                                                                   =
                                                                                                                                                                   ordEither2.clone();
                                                                                                                                                               move
                                                                                                                                                                   |usd__unused|
                                                                                                                                                                   &ordEither2
                                                                                                                                                           }),
                                                                                                                                              empty::<string,
                                                                                                                                                      &dyn Any>()))))
                                                                                          }
                                                                                  })
                                                                  }))
    }
    pub fn Data_Either_blush() -> &dyn Any {
        static Data_Either_blush: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_blush.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                              &&&Func1::new(move
                                                                                                                                |usd__arg1|
                                                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                              &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))))
    }
    pub fn Data_Either_applyEither() -> &dyn Any {
        static Data_Either_applyEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_applyEither.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                 &&&add(string("apply"),
                                                                                        &&Func1::new(move
                                                                                                         |v|
                                                                                                         &Func1::new({
                                                                                                                         let v
                                                                                                                             =
                                                                                                                             v.clone();
                                                                                                                         move
                                                                                                                             |v1|
                                                                                                                             {
                                                                                                                                 let matchValue:
                                                                                                                                         LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                     Sharpurs_Prelude::unbox(&&v);
                                                                                                                                 let matchValue_1 =
                                                                                                                                     Sharpurs_Prelude::unbox(v1);
                                                                                                                                 match matchValue.as_ref()
                                                                                                                                     {
                                                                                                                                     PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                     =>
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                            &&&PureScript_Data_Either::Data_Either_functorEither()),
                                                                                                                                                                                                         &&matchValue_1_0),
                                                                                                                                                                      &&&matchValue_1),
                                                                                                                                     PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)
                                                                                                                                     =>
                                                                                                                                     &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)),
                                                                                                                                 }
                                                                                                                             }
                                                                                                                     })),
                                                                                        add(string("Functor0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &PureScript_Data_Either::Data_Either_functorEither()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Either_bindEither() -> &dyn Any {
        static Data_Either_bindEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_bindEither.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                &&&add(string("bind"),
                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |e|
                                                                                                                                                                              &Func1::new({
                                                                                                                                                                                              let e
                                                                                                                                                                                                  =
                                                                                                                                                                                                  e.clone();
                                                                                                                                                                                              move
                                                                                                                                                                                                  |v|
                                                                                                                                                                                                  &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(&e))
                                                                                                                                                                                          }))),
                                                                                                                         &&&Func1::new(move
                                                                                                                                           |a|
                                                                                                                                           &Func1::new({
                                                                                                                                                           let a
                                                                                                                                                               =
                                                                                                                                                               a.clone();
                                                                                                                                                           move
                                                                                                                                                               |f|
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                                                                                                &&&a)
                                                                                                                                                       }))),
                                                                                       add(string("Apply0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Either::Data_Either_applyEither()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Data_Either_semigroupEither() -> &dyn Any {
        static Data_Either_semigroupEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_semigroupEither.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictSemigroup|
                                                                    {
                                                                        let append =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                             dictSemigroup);
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                         &&&add(string("append"),
                                                                                                                &&Func1::new({
                                                                                                                                 let append
                                                                                                                                     =
                                                                                                                                     append.clone();
                                                                                                                                 move
                                                                                                                                     |x|
                                                                                                                                     &Func1::new({
                                                                                                                                                     let x
                                                                                                                                                         =
                                                                                                                                                         x.clone();
                                                                                                                                                     move
                                                                                                                                                         |y|
                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                &&&PureScript_Data_Either::Data_Either_applyEither()),
                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Either::Data_Either_functorEither()),
                                                                                                                                                                                                                                                                                                   &&&append),
                                                                                                                                                                                                                                                                &&&x)),
                                                                                                                                                                                          y)
                                                                                                                                                 })
                                                                                                                             }),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))
                                                                    }))
    }
    pub fn Data_Either_applicativeEither() -> &dyn Any {
        static Data_Either_applicativeEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_applicativeEither.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                       &&&add(string("pure"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__arg1|
                                                                                                               &LrcPtr::new(PureScript_Data_Either::Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone()))),
                                                                                              add(string("Apply0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_Either::Data_Either_applyEither()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Either_monadEither() -> &dyn Any {
        static Data_Either_monadEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_monadEither.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                 &&&add(string("Applicative0"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused|
                                                                                                         &PureScript_Data_Either::Data_Either_applicativeEither()),
                                                                                        add(string("Bind1"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused_1|
                                                                                                             &PureScript_Data_Either::Data_Either_bindEither()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Either_altEither() -> &dyn Any {
        static Data_Either_altEither: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Either_altEither.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                               &&&add(string("alt"),
                                                                                      &&Func1::new(move
                                                                                                       |v|
                                                                                                       &Func1::new({
                                                                                                                       let v
                                                                                                                           =
                                                                                                                           v.clone();
                                                                                                                       move
                                                                                                                           |v1|
                                                                                                                           {
                                                                                                                               let matchValue:
                                                                                                                                       LrcPtr<PureScript_Data_Either::Data_Either_Either> =
                                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                                               if let PureScript_Data_Either::Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)
                                                                                                                                      =
                                                                                                                                      matchValue.as_ref()
                                                                                                                                  {
                                                                                                                                   &Sharpurs_Prelude::unbox(v1)
                                                                                                                               } else {
                                                                                                                                   &matchValue
                                                                                                                               }
                                                                                                                           }
                                                                                                                   })),
                                                                                      add(string("Functor0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Data_Either::Data_Either_functorEither()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
}
