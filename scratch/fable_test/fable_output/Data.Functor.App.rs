pub mod PureScript_Data_Functor_App {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_2a7662d2::PureScript_Unsafe_Coerce;
    pub fn Data_Functor_App_App() -> &dyn Any {
        static Data_Functor_App_App: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_App.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Functor_App_showApp() -> &dyn Any {
        static Data_Functor_App_showApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_showApp.get_or_init(||
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
                                                                                                                                  let fa =
                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                      &&&string("(App ")),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                               &&&dictShow),
                                                                                                                                                                                                                                                                            &&&fa)),
                                                                                                                                                                                                      &&&string(")")))
                                                                                                                              }
                                                                                                                      }),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Functor_App_semigroupApp() -> &dyn Any {
        static Data_Functor_App_semigroupApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_semigroupApp.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictApply|
                                                                      &Func1::new({
                                                                                      let dictApply
                                                                                          =
                                                                                          dictApply.clone();
                                                                                      move
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
                                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_App(),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_lift2(),
                                                                                                                                                                                                                                                                                                                                                                &&&dictApply),
                                                                                                                                                                                                                                                                                                                             &&&append),
                                                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                                                       &&&matchValue_1))
                                                                                                                                                                               }
                                                                                                                                                                       })
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>()))
                                                                                          }
                                                                                  })))
    }
    pub fn Data_Functor_App_plusApp() -> &dyn Any {
        static Data_Functor_App_plusApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_plusApp.get_or_init(||
                                                 &Func1::new(move |dictPlus|
                                                                 dictPlus.clone()))
    }
    pub fn Data_Functor_App_newtypeApp() -> &dyn Any {
        static Data_Functor_App_newtypeApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_newtypeApp.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                     &&&add(string("Coercible0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &Sharpurs_Prelude::Prim_undefined()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>())))
    }
    pub fn Data_Functor_App_monoidApp() -> &dyn Any {
        static Data_Functor_App_monoidApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_monoidApp.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictApplicative|
                                                                   {
                                                                       let semigroupApp1 =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_semigroupApp(),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                      Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                       &Func1::new({
                                                                                       let dictApplicative
                                                                                           =
                                                                                           dictApplicative.clone();
                                                                                       let semigroupApp1
                                                                                           =
                                                                                           semigroupApp1.clone();
                                                                                       move
                                                                                           |dictMonoid|
                                                                                           {
                                                                                               let semigroupApp2 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&semigroupApp1,
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                                &&&add(string("mempty"),
                                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_App(),
                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                               &&&dictApplicative),
                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                               dictMonoid))),
                                                                                                                                       add(string("Semigroup0"),
                                                                                                                                           &&Func1::new({
                                                                                                                                                            let semigroupApp2
                                                                                                                                                                =
                                                                                                                                                                semigroupApp2.clone();
                                                                                                                                                            move
                                                                                                                                                                |usd__unused|
                                                                                                                                                                &semigroupApp2
                                                                                                                                                        }),
                                                                                                                                           empty::<string,
                                                                                                                                                   &dyn Any>())))
                                                                                           }
                                                                                   })
                                                                   }))
    }
    pub fn Data_Functor_App_monadPlusApp() -> &dyn Any {
        static Data_Functor_App_monadPlusApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_monadPlusApp.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonadPlus|
                                                                      dictMonadPlus.clone()))
    }
    pub fn Data_Functor_App_monadApp() -> &dyn Any {
        static Data_Functor_App_monadApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_monadApp.get_or_init(||
                                                  &Func1::new(move |dictMonad|
                                                                  dictMonad.clone()))
    }
    pub fn Data_Functor_App_lazyApp() -> &dyn Any {
        static Data_Functor_App_lazyApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_lazyApp.get_or_init(||
                                                 &Func1::new(move |dictLazy|
                                                                 dictLazy.clone()))
    }
    pub fn Data_Functor_App_hoistLowerApp() -> &dyn Any {
        static Data_Functor_App_hoistLowerApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_hoistLowerApp.get_or_init(||
                                                       &PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce())
    }
    pub fn Data_Functor_App_hoistLiftApp() -> &dyn Any {
        static Data_Functor_App_hoistLiftApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_hoistLiftApp.get_or_init(||
                                                      &PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce())
    }
    pub fn Data_Functor_App_hoistApp() -> &dyn Any {
        static Data_Functor_App_hoistApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_hoistApp.get_or_init(||
                                                  &Func1::new(move |f|
                                                                  &Func1::new({
                                                                                  let f
                                                                                      =
                                                                                      f.clone();
                                                                                  move
                                                                                      |v|
                                                                                      {
                                                                                          let matchValue =
                                                                                              Sharpurs_Prelude::unbox(&&f);
                                                                                          let matchValue_1 =
                                                                                              Sharpurs_Prelude::unbox(v);
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_App(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                              &&&matchValue_1))
                                                                                      }
                                                                              })))
    }
    pub fn Data_Functor_App_functorApp() -> &dyn Any {
        static Data_Functor_App_functorApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_functorApp.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictFunctor|
                                                                    dictFunctor.clone()))
    }
    pub fn Data_Functor_App_extendApp() -> &dyn Any {
        static Data_Functor_App_extendApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_extendApp.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictExtend|
                                                                   dictExtend.clone()))
    }
    pub fn Data_Functor_App_eqApp() -> &dyn Any {
        static Data_Functor_App_eqApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_eqApp.get_or_init(||
                                               &Func1::new(move |dictEq1|
                                                               &Func1::new({
                                                                               let dictEq1
                                                                                   =
                                                                                   dictEq1.clone();
                                                                               move
                                                                                   |dictEq|
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
                                                                                                                                                                        let matchValue =
                                                                                                                                                                            Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                        let matchValue_1 =
                                                                                                                                                                            Sharpurs_Prelude::unbox(y);
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                  &&&dictEq1),
                                                                                                                                                                                                                                                                               &&&dictEq),
                                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                                         &&&matchValue_1)
                                                                                                                                                                    }
                                                                                                                                                            })
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))
                                                                           })))
    }
    pub fn Data_Functor_App_ordApp() -> &dyn Any {
        static Data_Functor_App_ordApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_ordApp.get_or_init(||
                                                &Func1::new(move |dictOrd1|
                                                                {
                                                                    let eqApp1 =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_eqApp(),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                   Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                    &Func1::new({
                                                                                    let dictOrd1
                                                                                        =
                                                                                        dictOrd1.clone();
                                                                                    let eqApp1
                                                                                        =
                                                                                        eqApp1.clone();
                                                                                    move
                                                                                        |dictOrd|
                                                                                        {
                                                                                            let eqApp2 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&eqApp1,
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
                                                                                                                                                                                 let matchValue =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                                 let matchValue_1 =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(y);
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                                                                                                                                                                                                           &&&dictOrd1),
                                                                                                                                                                                                                                                                                        &&&dictOrd),
                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                  &&&matchValue_1)
                                                                                                                                                                             }
                                                                                                                                                                     })
                                                                                                                                                 }),
                                                                                                                                    add(string("Eq0"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let eqApp2
                                                                                                                                                             =
                                                                                                                                                             eqApp2.clone();
                                                                                                                                                         move
                                                                                                                                                             |usd__unused|
                                                                                                                                                             &eqApp2
                                                                                                                                                     }),
                                                                                                                                        empty::<string,
                                                                                                                                                &dyn Any>())))
                                                                                        }
                                                                                })
                                                                }))
    }
    pub fn Data_Functor_App_eq1App() -> &dyn Any {
        static Data_Functor_App_eq1App: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_eq1App.get_or_init(||
                                                &Func1::new(move |dictEq1|
                                                                {
                                                                    let eqApp1 =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_eqApp(),
                                                                                                         dictEq1);
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                     &&&add(string("eq1"),
                                                                                                            &&Func1::new({
                                                                                                                             let eqApp1
                                                                                                                                 =
                                                                                                                                 eqApp1.clone();
                                                                                                                             move
                                                                                                                                 |dictEq|
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&eqApp1,
                                                                                                                                                                                                     dictEq))
                                                                                                                         }),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>()))
                                                                }))
    }
    pub fn Data_Functor_App_ord1App() -> &dyn Any {
        static Data_Functor_App_ord1App: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_ord1App.get_or_init(||
                                                 &Func1::new(move |dictOrd1|
                                                                 {
                                                                     let ordApp1 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_ordApp(),
                                                                                                          dictOrd1);
                                                                     let eq1App1 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_eq1App(),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                    Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                      &&&add(string("compare1"),
                                                                                                             &&Func1::new({
                                                                                                                              let ordApp1
                                                                                                                                  =
                                                                                                                                  ordApp1.clone();
                                                                                                                              move
                                                                                                                                  |dictOrd|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&ordApp1,
                                                                                                                                                                                                      dictOrd))
                                                                                                                          }),
                                                                                                             add(string("Eq10"),
                                                                                                                 &&Func1::new({
                                                                                                                                  let eq1App1
                                                                                                                                      =
                                                                                                                                      eq1App1.clone();
                                                                                                                                  move
                                                                                                                                      |usd__unused|
                                                                                                                                      &eq1App1
                                                                                                                              }),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>())))
                                                                 }))
    }
    pub fn Data_Functor_App_comonadApp() -> &dyn Any {
        static Data_Functor_App_comonadApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_comonadApp.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictComonad|
                                                                    dictComonad.clone()))
    }
    pub fn Data_Functor_App_bindApp() -> &dyn Any {
        static Data_Functor_App_bindApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_bindApp.get_or_init(||
                                                 &Func1::new(move |dictBind|
                                                                 dictBind.clone()))
    }
    pub fn Data_Functor_App_applyApp() -> &dyn Any {
        static Data_Functor_App_applyApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_applyApp.get_or_init(||
                                                  &Func1::new(move |dictApply|
                                                                  dictApply.clone()))
    }
    pub fn Data_Functor_App_applicativeApp() -> &dyn Any {
        static Data_Functor_App_applicativeApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_applicativeApp.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictApplicative|
                                                                        dictApplicative.clone()))
    }
    pub fn Data_Functor_App_alternativeApp() -> &dyn Any {
        static Data_Functor_App_alternativeApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_alternativeApp.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictAlternative|
                                                                        dictAlternative.clone()))
    }
    pub fn Data_Functor_App_altApp() -> &dyn Any {
        static Data_Functor_App_altApp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_App_altApp.get_or_init(||
                                                &Func1::new(move |dictAlt|
                                                                dictAlt.clone()))
    }
}
