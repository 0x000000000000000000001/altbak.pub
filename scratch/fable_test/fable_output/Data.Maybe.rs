pub mod PureScript_Data_Maybe {
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
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_NoArguments;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Sum;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug,)]
    pub enum Data_Maybe_Maybe {
        Data_Maybe_Nothingusd_Ctor,
        Data_Maybe_Justusd_Ctor(&dyn Any),
    }
    impl core::fmt::Display for PureScript_Data_Maybe::Data_Maybe_Maybe {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Maybe_identity() -> &dyn Any {
        static Data_Maybe_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_identity.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                             &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_Maybe_Nothing() -> &dyn Any {
        static Data_Maybe_Nothing: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Nothing.get_or_init(||
                                           &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))
    }
    pub fn Data_Maybe_Just() -> &dyn Any {
        static Data_Maybe_Just: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Just.get_or_init(||
                                        &Func1::new(move |usd__arg1|
                                                        &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Maybe_showMaybe() -> &dyn Any {
        static Data_Maybe_showMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_showMaybe.get_or_init(||
                                             &Func1::new(move |dictShow|
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                              &&&add(string("show"),
                                                                                                     &&Func1::new({
                                                                                                                      let dictShow
                                                                                                                          =
                                                                                                                          dictShow.clone();
                                                                                                                      move
                                                                                                                          |v|
                                                                                                                          {
                                                                                                                              let matchValue:
                                                                                                                                      LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                              match matchValue.as_ref()
                                                                                                                                  {
                                                                                                                                  PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                  =>
                                                                                                                                  &string("Nothing"),
                                                                                                                                  PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                  =>
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                      &&&string("(Just ")),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                               &&&dictShow),
                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                      &&&string(")"))),
                                                                                                                              }
                                                                                                                          }
                                                                                                                  }),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Maybe_semigroupMaybe() -> &dyn Any {
        static Data_Maybe_semigroupMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_semigroupMaybe.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictSemigroup|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                   &&&add(string("append"),
                                                                                                          &&Func1::new({
                                                                                                                           let dictSemigroup
                                                                                                                               =
                                                                                                                               dictSemigroup.clone();
                                                                                                                           move
                                                                                                                               |v|
                                                                                                                               &Func1::new({
                                                                                                                                               let v
                                                                                                                                                   =
                                                                                                                                                   v.clone();
                                                                                                                                               move
                                                                                                                                                   |v1|
                                                                                                                                                   {
                                                                                                                                                       let matchValue:
                                                                                                                                                               LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                                           Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                       let matchValue_1:
                                                                                                                                                               LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                                           Sharpurs_Prelude::unbox(v1);
                                                                                                                                                       if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                              =
                                                                                                                                                              matchValue.as_ref()
                                                                                                                                                          {
                                                                                                                                                           if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                  =
                                                                                                                                                                  matchValue_1.as_ref()
                                                                                                                                                              {
                                                                                                                                                               &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                    &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                        PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                     PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                     _
                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                                                                                 })))
                                                                                                                                                           } else {
                                                                                                                                                               &matchValue
                                                                                                                                                           }
                                                                                                                                                       } else {
                                                                                                                                                           &matchValue_1
                                                                                                                                                       }
                                                                                                                                                   }
                                                                                                                                           })
                                                                                                                       }),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Maybe_optional() -> &dyn Any {
        static Data_Maybe_optional: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_optional.get_or_init(||
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
                                                                                    |dictApplicative|
                                                                                    &Func1::new({
                                                                                                    let dictApplicative
                                                                                                        =
                                                                                                        dictApplicative.clone();
                                                                                                    move
                                                                                                        |a|
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                               &&&dictAlt),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                     &&&Functor0),
                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                    |usd__arg1|
                                                                                                                                                                                                                                                                    &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                               a)),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                               &&&dictApplicative),
                                                                                                                                                                            &&&LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
                                                                                                })
                                                                            })
                                                            }))
    }
    pub fn Data_Maybe_monoidMaybe() -> &dyn Any {
        static Data_Maybe_monoidMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_monoidMaybe.get_or_init(||
                                               &Func1::new(move
                                                               |dictSemigroup|
                                                               {
                                                                   let semigroupMaybe1 =
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_semigroupMaybe(),
                                                                                                        dictSemigroup);
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                    &&&add(string("mempty"),
                                                                                                           &&LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                           add(string("Semigroup0"),
                                                                                                               &&Func1::new({
                                                                                                                                let semigroupMaybe1
                                                                                                                                    =
                                                                                                                                    semigroupMaybe1.clone();
                                                                                                                                move
                                                                                                                                    |usd__unused|
                                                                                                                                    &semigroupMaybe1
                                                                                                                            }),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>())))
                                                               }))
    }
    pub fn Data_Maybe_maybe_prime() -> &dyn Any {
        static Data_Maybe_maybe_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_maybe_prime.get_or_init(||
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
                                                                                                                   LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                               Sharpurs_Prelude::unbox(v2);
                                                                                                           match matchValue_2.as_ref()
                                                                                                               {
                                                                                                               PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_2_1_0)
                                                                                                               =>
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                &&matchValue_2_1_0),
                                                                                                               _
                                                                                                               =>
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                           }
                                                                                                       }
                                                                                               })
                                                                           })))
    }
    pub fn Data_Maybe_maybe() -> &dyn Any {
        static Data_Maybe_maybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_maybe.get_or_init(||
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
                                                                                                             LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                         Sharpurs_Prelude::unbox(v2);
                                                                                                     match matchValue_2.as_ref()
                                                                                                         {
                                                                                                         PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_2_1_0)
                                                                                                         =>
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                          &&matchValue_2_1_0),
                                                                                                         _
                                                                                                         =>
                                                                                                         &matchValue,
                                                                                                     }
                                                                                                 }
                                                                                         })
                                                                     })))
    }
    pub fn Data_Maybe_isNothing() -> &dyn Any {
        static Data_Maybe_isNothing: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_isNothing.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                 &&&true),
                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                 &&&false)))
    }
    pub fn Data_Maybe_isJust() -> &dyn Any {
        static Data_Maybe_isJust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_isJust.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                              &&&false),
                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                              &&&true)))
    }
    pub fn Data_Maybe_genericMaybe() -> &dyn Any {
        static Data_Maybe_genericMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_genericMaybe.get_or_init(||
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
                                                                                                                 &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)),
                                                                                                                 _
                                                                                                                 =>
                                                                                                                 &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                             }
                                                                                                         }),
                                                                                        add(string("from"),
                                                                                            &&Func1::new(move
                                                                                                             |x_1|
                                                                                                             {
                                                                                                                 let matchValue_1:
                                                                                                                         LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                     Sharpurs_Prelude::unbox(x_1);
                                                                                                                 match matchValue_1.as_ref()
                                                                                                                     {
                                                                                                                     PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                     =>
                                                                                                                     &LrcPtr::new(Data_Generic_Rep_Sum::Data_Generic_Rep_Inrusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                                                                         &&matchValue_1_1_0)))),
                                                                                                                     _
                                                                                                                     =>
                                                                                                                     &LrcPtr::new(Data_Generic_Rep_Sum::Data_Generic_Rep_Inlusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                                                                                      &&&LrcPtr::new(Data_Generic_Rep_NoArguments::Data_Generic_Rep_NoArgumentsusd_Ctor)))),
                                                                                                                 }
                                                                                                             }),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Maybe_functorMaybe() -> &dyn Any {
        static Data_Maybe_functorMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_functorMaybe.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                 &&&add(string("map"),
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
                                                                                                                                         LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                     Sharpurs_Prelude::unbox(v1);
                                                                                                                                 if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                        =
                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                    {
                                                                                                                                     &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                           PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                           _
                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                                                       })))
                                                                                                                                 } else {
                                                                                                                                     &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                                 }
                                                                                                                             }
                                                                                                                     })),
                                                                                        empty::<string,
                                                                                                &dyn Any>())))
    }
    pub fn Data_Maybe_invariantMaybe() -> &dyn Any {
        static Data_Maybe_invariantMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_invariantMaybe.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                   &&&add(string("imap"),
                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                            &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Maybe_fromMaybe_prime() -> &dyn Any {
        static Data_Maybe_fromMaybe_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_fromMaybe_prime.get_or_init(||
                                                   &Func1::new(move |a|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe_prime(),
                                                                                                                                       a),
                                                                                                    &&&PureScript_Data_Maybe::Data_Maybe_identity())))
    }
    pub fn Data_Maybe_fromMaybe() -> &dyn Any {
        static Data_Maybe_fromMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_fromMaybe.get_or_init(||
                                             &Func1::new(move |a|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                 a),
                                                                                              &&&PureScript_Data_Maybe::Data_Maybe_identity())))
    }
    pub fn Data_Maybe_fromJust() -> &dyn Any {
        static Data_Maybe_fromJust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_fromJust.get_or_init(||
                                            &Func1::new(move |usd__unused|
                                                            &Func1::new(move
                                                                            |v|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&Func1::new({
                                                                                                                               let v
                                                                                                                                   =
                                                                                                                                   v.clone();
                                                                                                                               move
                                                                                                                                   |usd__unused_1|
                                                                                                                                   {
                                                                                                                                       let matchValue:
                                                                                                                                               LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                           Sharpurs_Prelude::unbox(&&v);
                                                                                                                                       if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                              =
                                                                                                                                              matchValue.as_ref()
                                                                                                                                          {
                                                                                                                                           &match matchValue.as_ref()
                                                                                                                                                {
                                                                                                                                                PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                =>
                                                                                                                                                x.clone(),
                                                                                                                                                _
                                                                                                                                                =>
                                                                                                                                                unreachable!(),
                                                                                                                                            }
                                                                                                                                       } else {
                                                                                                                                           panic!("{}",
                                                                                                                                                  string("Match failure"),)
                                                                                                                                       }
                                                                                                                                   }
                                                                                                                           }),
                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()))))
    }
    pub fn Data_Maybe_extendMaybe() -> &dyn Any {
        static Data_Maybe_extendMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_extendMaybe.get_or_init(||
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
                                                                                                                                        LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                    Sharpurs_Prelude::unbox(v1);
                                                                                                                                if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                       =
                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                    &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                                } else {
                                                                                                                                    &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                   &&&matchValue_1)))
                                                                                                                                }
                                                                                                                            }
                                                                                                                    })),
                                                                                       add(string("Functor0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Data_Maybe_eqMaybe() -> &dyn Any {
        static Data_Maybe_eqMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_eqMaybe.get_or_init(||
                                           &Func1::new(move |dictEq|
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                            &&&add(string("eq"),
                                                                                                   &&Func1::new({
                                                                                                                    let dictEq
                                                                                                                        =
                                                                                                                        dictEq.clone();
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
                                                                                                                                                        LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                                    Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                let matchValue_1:
                                                                                                                                                        LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                                    Sharpurs_Prelude::unbox(y);
                                                                                                                                                if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                       =
                                                                                                                                                       matchValue.as_ref()
                                                                                                                                                   {
                                                                                                                                                    if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                                           =
                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                       {
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                               &&&dictEq),
                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                   PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                               }),
                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
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
                                                                                                                                                    if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                           =
                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                       {
                                                                                                                                                        &true
                                                                                                                                                    } else {
                                                                                                                                                        &false
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                    })
                                                                                                                }),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Maybe_ordMaybe() -> &dyn Any {
        static Data_Maybe_ordMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_ordMaybe.get_or_init(||
                                            &Func1::new(move |dictOrd|
                                                            {
                                                                let eqMaybe1 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_eqMaybe(),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                               Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                 &&&add(string("compare"),
                                                                                                        &&Func1::new({
                                                                                                                         let dictOrd
                                                                                                                             =
                                                                                                                             dictOrd.clone();
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
                                                                                                                                                             LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                                         Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                     let matchValue_1:
                                                                                                                                                             LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                                         Sharpurs_Prelude::unbox(y);
                                                                                                                                                     if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                            =
                                                                                                                                                            matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                         if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                =
                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                            {
                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                    &&&dictOrd),
                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                    }),
                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                     PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
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
                                                                                                                                                         if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                            {
                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                         } else {
                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                         }
                                                                                                                                                     }
                                                                                                                                                 }
                                                                                                                                         })
                                                                                                                     }),
                                                                                                        add(string("Eq0"),
                                                                                                            &&Func1::new({
                                                                                                                             let eqMaybe1
                                                                                                                                 =
                                                                                                                                 eqMaybe1.clone();
                                                                                                                             move
                                                                                                                                 |usd__unused|
                                                                                                                                 &eqMaybe1
                                                                                                                         }),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>())))
                                                            }))
    }
    pub fn Data_Maybe_eq1Maybe() -> &dyn Any {
        static Data_Maybe_eq1Maybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_eq1Maybe.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                             &&&add(string("eq1"),
                                                                                    &&Func1::new(move
                                                                                                     |dictEq|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_eqMaybe(),
                                                                                                                                                                         dictEq))),
                                                                                    empty::<string,
                                                                                            &dyn Any>())))
    }
    pub fn Data_Maybe_ord1Maybe() -> &dyn Any {
        static Data_Maybe_ord1Maybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_ord1Maybe.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                              &&&add(string("compare1"),
                                                                                     &&Func1::new(move
                                                                                                      |dictOrd|
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_ordMaybe(),
                                                                                                                                                                          dictOrd))),
                                                                                     add(string("Eq10"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Data_Maybe::Data_Maybe_eq1Maybe()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Maybe_boundedMaybe() -> &dyn Any {
        static Data_Maybe_boundedMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_boundedMaybe.get_or_init(||
                                                &Func1::new(move |dictBounded|
                                                                {
                                                                    let ordMaybe1 =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_ordMaybe(),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                   Sharpurs_Prelude::unbox(dictBounded)),
                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                                     &&&add(string("top"),
                                                                                                            &&LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                                            dictBounded))),
                                                                                                            add(string("bottom"),
                                                                                                                &&LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                add(string("Ord0"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let ordMaybe1
                                                                                                                                         =
                                                                                                                                         ordMaybe1.clone();
                                                                                                                                     move
                                                                                                                                         |usd__unused|
                                                                                                                                         &ordMaybe1
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>()))))
                                                                }))
    }
    pub fn Data_Maybe_applyMaybe() -> &dyn Any {
        static Data_Maybe_applyMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_applyMaybe.get_or_init(||
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
                                                                                                                                       LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                                               let matchValue_1 =
                                                                                                                                   Sharpurs_Prelude::unbox(v1);
                                                                                                                               match matchValue.as_ref()
                                                                                                                                   {
                                                                                                                                   PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                   =>
                                                                                                                                   &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                   PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                   =>
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                          &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          }),
                                                                                                                                                                    &&&matchValue_1),
                                                                                                                               }
                                                                                                                           }
                                                                                                                   })),
                                                                                      add(string("Functor0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Maybe_bindMaybe() -> &dyn Any {
        static Data_Maybe_bindMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_bindMaybe.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                              &&&add(string("bind"),
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
                                                                                                                                      LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                                                              let matchValue_1 =
                                                                                                                                  Sharpurs_Prelude::unbox(v1);
                                                                                                                              match matchValue.as_ref()
                                                                                                                                  {
                                                                                                                                  PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                  =>
                                                                                                                                  &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                  PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                  =>
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                   &&&match matchValue.as_ref()
                                                                                                                                                                          {
                                                                                                                                                                          PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                          =>
                                                                                                                                                                          x.clone(),
                                                                                                                                                                          _
                                                                                                                                                                          =>
                                                                                                                                                                          unreachable!(),
                                                                                                                                                                      }),
                                                                                                                              }
                                                                                                                          }
                                                                                                                  })),
                                                                                     add(string("Apply0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Maybe_semiringMaybe() -> &dyn Any {
        static Data_Maybe_semiringMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_semiringMaybe.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictSemiring|
                                                                 {
                                                                     let mul =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                          dictSemiring);
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                                      &&&add(string("zero"),
                                                                                                             &&LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                             add(string("one"),
                                                                                                                 &&LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                                                                                 dictSemiring))),
                                                                                                                 add(string("add"),
                                                                                                                     &&Func1::new({
                                                                                                                                      let dictSemiring
                                                                                                                                          =
                                                                                                                                          dictSemiring.clone();
                                                                                                                                      move
                                                                                                                                          |v|
                                                                                                                                          &Func1::new({
                                                                                                                                                          let v
                                                                                                                                                              =
                                                                                                                                                              v.clone();
                                                                                                                                                          move
                                                                                                                                                              |v1|
                                                                                                                                                              {
                                                                                                                                                                  let matchValue:
                                                                                                                                                                          LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                  let matchValue_1:
                                                                                                                                                                          LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                  if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                         =
                                                                                                                                                                         matchValue.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                               &&&dictSemiring),
                                                                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                   PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                                                                            })))
                                                                                                                                                                      } else {
                                                                                                                                                                          &matchValue
                                                                                                                                                                      }
                                                                                                                                                                  } else {
                                                                                                                                                                      &matchValue_1
                                                                                                                                                                  }
                                                                                                                                                              }
                                                                                                                                                      })
                                                                                                                                  }),
                                                                                                                     add(string("mul"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let mul
                                                                                                                                              =
                                                                                                                                              mul.clone();
                                                                                                                                          move
                                                                                                                                              |x_2|
                                                                                                                                              &Func1::new({
                                                                                                                                                              let x_2
                                                                                                                                                                  =
                                                                                                                                                                  x_2.clone();
                                                                                                                                                              move
                                                                                                                                                                  |y_2|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                            &&&mul),
                                                                                                                                                                                                                                                                         &&&x_2)),
                                                                                                                                                                                                   y_2)
                                                                                                                                                          })
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))))
                                                                 }))
    }
    pub fn Data_Maybe_applicativeMaybe() -> &dyn Any {
        static Data_Maybe_applicativeMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_applicativeMaybe.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                     &&&add(string("pure"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__arg1|
                                                                                                             &LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone()))),
                                                                                            add(string("Apply0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Maybe_monadMaybe() -> &dyn Any {
        static Data_Maybe_monadMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_monadMaybe.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                               &&&add(string("Applicative0"),
                                                                                      &&Func1::new(move
                                                                                                       |usd__unused|
                                                                                                       &PureScript_Data_Maybe::Data_Maybe_applicativeMaybe()),
                                                                                      add(string("Bind1"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused_1|
                                                                                                           &PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Maybe_altMaybe() -> &dyn Any {
        static Data_Maybe_altMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_altMaybe.get_or_init(||
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
                                                                                                                                     LrcPtr<PureScript_Data_Maybe::Data_Maybe_Maybe> =
                                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                                             if let PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
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
                                                                                                         &PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                        empty::<string,
                                                                                                &dyn Any>()))))
    }
    pub fn Data_Maybe_plusMaybe() -> &dyn Any {
        static Data_Maybe_plusMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_plusMaybe.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                              &&&add(string("empty"),
                                                                                     &&LrcPtr::new(PureScript_Data_Maybe::Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                     add(string("Alt0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Data_Maybe::Data_Maybe_altMaybe()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Maybe_alternativeMaybe() -> &dyn Any {
        static Data_Maybe_alternativeMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_alternativeMaybe.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                     &&&add(string("Applicative0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &PureScript_Data_Maybe::Data_Maybe_applicativeMaybe()),
                                                                                            add(string("Plus1"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused_1|
                                                                                                                 &PureScript_Data_Maybe::Data_Maybe_plusMaybe()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
}
