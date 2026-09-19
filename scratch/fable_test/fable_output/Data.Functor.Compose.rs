pub mod PureScript_Data_Functor_Compose {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty as empty_1;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_3ae611ad::PureScript_Data_Functor_App;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Compose_Compose() -> &dyn Any {
        static Data_Functor_Compose_Compose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_Compose.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Data_Functor_Compose_showCompose() -> &dyn Any {
        static Data_Functor_Compose_showCompose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_showCompose.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictShow|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                          &&&add(string("show"),
                                                                                                                 &&Func1::new({
                                                                                                                                  let dictShow
                                                                                                                                      =
                                                                                                                                      dictShow.clone();
                                                                                                                                  move
                                                                                                                                      |v|
                                                                                                                                      {
                                                                                                                                          let fga =
                                                                                                                                              Sharpurs_Prelude::unbox(v);
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                              &&&string("(Compose ")),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                       &&&dictShow),
                                                                                                                                                                                                                                                                                    &&&fga)),
                                                                                                                                                                                                              &&&string(")")))
                                                                                                                                      }
                                                                                                                              }),
                                                                                                                 empty_1::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Functor_Compose_newtypeCompose() -> &dyn Any {
        static Data_Functor_Compose_newtypeCompose: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Compose_newtypeCompose.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                             &&&add(string("Coercible0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &Sharpurs_Prelude::Prim_undefined()),
                                                                                                    empty_1::<string,
                                                                                                              &dyn Any>())))
    }
    pub fn Data_Functor_Compose_functorCompose() -> &dyn Any {
        static Data_Functor_Compose_functorCompose: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Compose_functorCompose.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFunctor|
                                                                            &Func1::new({
                                                                                            let dictFunctor
                                                                                                =
                                                                                                dictFunctor.clone();
                                                                                            move
                                                                                                |dictFunctor1|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                                                 &&&add(string("map"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let dictFunctor1
                                                                                                                                                             =
                                                                                                                                                             dictFunctor1.clone();
                                                                                                                                                         move
                                                                                                                                                             |f|
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
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                         &&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose()),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                               &&&dictFunctor),
                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                  &&&dictFunctor1),
                                                                                                                                                                                                                                                                                                                               &&&matchValue)),
                                                                                                                                                                                                                                                         &&&matchValue_1))
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     }),
                                                                                                                                        empty_1::<string,
                                                                                                                                                  &dyn Any>()))
                                                                                        })))
    }
    pub fn Data_Functor_Compose_eqCompose() -> &dyn Any {
        static Data_Functor_Compose_eqCompose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_eqCompose.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictEq1|
                                                                       &Func1::new({
                                                                                       let dictEq1
                                                                                           =
                                                                                           dictEq1.clone();
                                                                                       move
                                                                                           |dictEq11|
                                                                                           {
                                                                                               let eqApp =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_eqApp(),
                                                                                                                                    dictEq11);
                                                                                               &Func1::new({
                                                                                                               let eqApp
                                                                                                                   =
                                                                                                                   eqApp.clone();
                                                                                                               move
                                                                                                                   |dictEq|
                                                                                                                   {
                                                                                                                       let eqApp1 =
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&eqApp,
                                                                                                                                                            dictEq);
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                                                        &&&add(string("eq"),
                                                                                                                                                               &&Func1::new({
                                                                                                                                                                                let eqApp1
                                                                                                                                                                                    =
                                                                                                                                                                                    eqApp1.clone();
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
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                                                      &&&dictEq1),
                                                                                                                                                                                                                                                                                                                   &&&eqApp1),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_hoistLiftApp(),
                                                                                                                                                                                                                                                                                                                   &&&matchValue)),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_hoistLiftApp(),
                                                                                                                                                                                                                                                                                &&&matchValue_1))
                                                                                                                                                                                                        }
                                                                                                                                                                                                })
                                                                                                                                                                            }),
                                                                                                                                                               empty_1::<string,
                                                                                                                                                                         &dyn Any>()))
                                                                                                                   }
                                                                                                           })
                                                                                           }
                                                                                   })))
    }
    pub fn Data_Functor_Compose_ordCompose() -> &dyn Any {
        static Data_Functor_Compose_ordCompose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_ordCompose.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictOrd1|
                                                                        {
                                                                            let eqCompose1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_eqCompose(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            &Func1::new({
                                                                                            let dictOrd1
                                                                                                =
                                                                                                dictOrd1.clone();
                                                                                            let eqCompose1
                                                                                                =
                                                                                                eqCompose1.clone();
                                                                                            move
                                                                                                |dictOrd11|
                                                                                                {
                                                                                                    let ordApp =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_ordApp(),
                                                                                                                                         dictOrd11);
                                                                                                    let eqCompose2 =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&eqCompose1,
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictOrd11)),
                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                    &Func1::new({
                                                                                                                    let eqCompose2
                                                                                                                        =
                                                                                                                        eqCompose2.clone();
                                                                                                                    let ordApp
                                                                                                                        =
                                                                                                                        ordApp.clone();
                                                                                                                    move
                                                                                                                        |dictOrd|
                                                                                                                        {
                                                                                                                            let ordApp1 =
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&ordApp,
                                                                                                                                                                 dictOrd);
                                                                                                                            let eqCompose3 =
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&eqCompose2,
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                                                             &&&add(string("compare"),
                                                                                                                                                                    &&Func1::new({
                                                                                                                                                                                     let ordApp1
                                                                                                                                                                                         =
                                                                                                                                                                                         ordApp1.clone();
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
                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                                                                                                                                                                                                                                           &&&dictOrd1),
                                                                                                                                                                                                                                                                                                                        &&&ordApp1),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_hoistLiftApp(),
                                                                                                                                                                                                                                                                                                                        &&&matchValue)),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_App::Data_Functor_App_hoistLiftApp(),
                                                                                                                                                                                                                                                                                     &&&matchValue_1))
                                                                                                                                                                                                             }
                                                                                                                                                                                                     })
                                                                                                                                                                                 }),
                                                                                                                                                                    add(string("Eq0"),
                                                                                                                                                                        &&Func1::new({
                                                                                                                                                                                         let eqCompose3
                                                                                                                                                                                             =
                                                                                                                                                                                             eqCompose3.clone();
                                                                                                                                                                                         move
                                                                                                                                                                                             |usd__unused|
                                                                                                                                                                                             &eqCompose3
                                                                                                                                                                                     }),
                                                                                                                                                                        empty_1::<string,
                                                                                                                                                                                  &dyn Any>())))
                                                                                                                        }
                                                                                                                })
                                                                                                }
                                                                                        })
                                                                        }))
    }
    pub fn Data_Functor_Compose_eq1Compose() -> &dyn Any {
        static Data_Functor_Compose_eq1Compose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_eq1Compose.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictEq1|
                                                                        {
                                                                            let eqCompose1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_eqCompose(),
                                                                                                                 dictEq1);
                                                                            &Func1::new({
                                                                                            let eqCompose1
                                                                                                =
                                                                                                eqCompose1.clone();
                                                                                            move
                                                                                                |dictEq11|
                                                                                                {
                                                                                                    let eqCompose2 =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&eqCompose1,
                                                                                                                                         dictEq11);
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                                                     &&&add(string("eq1"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let eqCompose2
                                                                                                                                                                 =
                                                                                                                                                                 eqCompose2.clone();
                                                                                                                                                             move
                                                                                                                                                                 |dictEq|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&eqCompose2,
                                                                                                                                                                                                                                     dictEq))
                                                                                                                                                         }),
                                                                                                                                            empty_1::<string,
                                                                                                                                                      &dyn Any>()))
                                                                                                }
                                                                                        })
                                                                        }))
    }
    pub fn Data_Functor_Compose_ord1Compose() -> &dyn Any {
        static Data_Functor_Compose_ord1Compose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_ord1Compose.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictOrd1|
                                                                         {
                                                                             let ordCompose1 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_ordCompose(),
                                                                                                                  dictOrd1);
                                                                             let eq1Compose1 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_eq1Compose(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             &Func1::new({
                                                                                             let eq1Compose1
                                                                                                 =
                                                                                                 eq1Compose1.clone();
                                                                                             let ordCompose1
                                                                                                 =
                                                                                                 ordCompose1.clone();
                                                                                             move
                                                                                                 |dictOrd11|
                                                                                                 {
                                                                                                     let ordCompose2 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&ordCompose1,
                                                                                                                                          dictOrd11);
                                                                                                     let eq1Compose2 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&eq1Compose1,
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictOrd11)),
                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                                                      &&&add(string("compare1"),
                                                                                                                                             &&Func1::new({
                                                                                                                                                              let ordCompose2
                                                                                                                                                                  =
                                                                                                                                                                  ordCompose2.clone();
                                                                                                                                                              move
                                                                                                                                                                  |dictOrd|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&ordCompose2,
                                                                                                                                                                                                                                      dictOrd))
                                                                                                                                                          }),
                                                                                                                                             add(string("Eq10"),
                                                                                                                                                 &&Func1::new({
                                                                                                                                                                  let eq1Compose2
                                                                                                                                                                      =
                                                                                                                                                                      eq1Compose2.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |usd__unused|
                                                                                                                                                                      &eq1Compose2
                                                                                                                                                              }),
                                                                                                                                                 empty_1::<string,
                                                                                                                                                           &dyn Any>())))
                                                                                                 }
                                                                                         })
                                                                         }))
    }
    pub fn Data_Functor_Compose_bihoistCompose() -> &dyn Any {
        static Data_Functor_Compose_bihoistCompose: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Compose_bihoistCompose.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFunctor|
                                                                            &Func1::new({
                                                                                            let dictFunctor
                                                                                                =
                                                                                                dictFunctor.clone();
                                                                                            move
                                                                                                |natF|
                                                                                                &Func1::new({
                                                                                                                let natF
                                                                                                                    =
                                                                                                                    natF.clone();
                                                                                                                move
                                                                                                                    |natG|
                                                                                                                    &Func1::new({
                                                                                                                                    let natG
                                                                                                                                        =
                                                                                                                                        natG.clone();
                                                                                                                                    move
                                                                                                                                        |v|
                                                                                                                                        {
                                                                                                                                            let matchValue =
                                                                                                                                                Sharpurs_Prelude::unbox(&&natF);
                                                                                                                                            let matchValue_1 =
                                                                                                                                                Sharpurs_Prelude::unbox(&&natG);
                                                                                                                                            let matchValue_2 =
                                                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose(),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                         &&&dictFunctor),
                                                                                                                                                                                                                                                                                      &&&matchValue_1),
                                                                                                                                                                                                                                                   &&&matchValue_2)))
                                                                                                                                        }
                                                                                                                                })
                                                                                                            })
                                                                                        })))
    }
    pub fn Data_Functor_Compose_applyCompose() -> &dyn Any {
        static Data_Functor_Compose_applyCompose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_applyCompose.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictApply|
                                                                          {
                                                                              let Functor0 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                          Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              let functorCompose1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_functorCompose(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              &Func1::new({
                                                                                              let Functor0
                                                                                                  =
                                                                                                  Functor0.clone();
                                                                                              let dictApply
                                                                                                  =
                                                                                                  dictApply.clone();
                                                                                              let functorCompose1
                                                                                                  =
                                                                                                  functorCompose1.clone();
                                                                                              move
                                                                                                  |dictApply1|
                                                                                                  {
                                                                                                      let apply =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                           dictApply1);
                                                                                                      let functorCompose2 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&functorCompose1,
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictApply1)),
                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                                                       &&&add(string("apply"),
                                                                                                                                              &&Func1::new({
                                                                                                                                                               let apply
                                                                                                                                                                   =
                                                                                                                                                                   apply.clone();
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
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose()),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                     &&&dictApply),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                           &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                        &&&apply),
                                                                                                                                                                                                                                                                                                                                     &&&matchValue)),
                                                                                                                                                                                                                                                               &&&matchValue_1))
                                                                                                                                                                                       }
                                                                                                                                                                               })
                                                                                                                                                           }),
                                                                                                                                              add(string("Functor0"),
                                                                                                                                                  &&Func1::new({
                                                                                                                                                                   let functorCompose2
                                                                                                                                                                       =
                                                                                                                                                                       functorCompose2.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |usd__unused|
                                                                                                                                                                       &functorCompose2
                                                                                                                                                               }),
                                                                                                                                                  empty_1::<string,
                                                                                                                                                            &dyn Any>())))
                                                                                                  }
                                                                                          })
                                                                          }))
    }
    pub fn Data_Functor_Compose_applicativeCompose() -> &dyn Any {
        static Data_Functor_Compose_applicativeCompose:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_applicativeCompose.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictApplicative|
                                                                                {
                                                                                    let pure_var =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                         dictApplicative);
                                                                                    let applyCompose1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_applyCompose(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    &Func1::new({
                                                                                                    let applyCompose1
                                                                                                        =
                                                                                                        applyCompose1.clone();
                                                                                                    let pure_var
                                                                                                        =
                                                                                                        pure_var.clone();
                                                                                                    move
                                                                                                        |dictApplicative1|
                                                                                                        {
                                                                                                            let applyCompose2 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&applyCompose1,
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictApplicative1)),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                                                             &&&add(string("pure"),
                                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                         &&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose()),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                            &&&pure_var),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                            dictApplicative1))),
                                                                                                                                                    add(string("Apply0"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let applyCompose2
                                                                                                                                                                             =
                                                                                                                                                                             applyCompose2.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |usd__unused|
                                                                                                                                                                             &applyCompose2
                                                                                                                                                                     }),
                                                                                                                                                        empty_1::<string,
                                                                                                                                                                  &dyn Any>())))
                                                                                                        }
                                                                                                })
                                                                                }))
    }
    pub fn Data_Functor_Compose_altCompose() -> &dyn Any {
        static Data_Functor_Compose_altCompose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_altCompose.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictAlt|
                                                                        {
                                                                            let functorCompose1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_functorCompose(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictAlt)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            &Func1::new({
                                                                                            let dictAlt
                                                                                                =
                                                                                                dictAlt.clone();
                                                                                            let functorCompose1
                                                                                                =
                                                                                                functorCompose1.clone();
                                                                                            move
                                                                                                |dictFunctor|
                                                                                                {
                                                                                                    let functorCompose2 =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&functorCompose1,
                                                                                                                                         dictFunctor);
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
                                                                                                                                                                                     let matchValue =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                     let matchValue_1 =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                         &&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose()),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                                                                                                                                               &&&dictAlt),
                                                                                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                                                                                         &&&matchValue_1))
                                                                                                                                                                                 }
                                                                                                                                                                         })),
                                                                                                                                            add(string("Functor0"),
                                                                                                                                                &&Func1::new({
                                                                                                                                                                 let functorCompose2
                                                                                                                                                                     =
                                                                                                                                                                     functorCompose2.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |usd__unused|
                                                                                                                                                                     &functorCompose2
                                                                                                                                                             }),
                                                                                                                                                empty_1::<string,
                                                                                                                                                          &dyn Any>())))
                                                                                                }
                                                                                        })
                                                                        }))
    }
    pub fn Data_Functor_Compose_plusCompose() -> &dyn Any {
        static Data_Functor_Compose_plusCompose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_plusCompose.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictPlus|
                                                                         {
                                                                             let empty =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                  dictPlus);
                                                                             let altCompose1 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_altCompose(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictPlus)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             &Func1::new({
                                                                                             let altCompose1
                                                                                                 =
                                                                                                 altCompose1.clone();
                                                                                             let empty
                                                                                                 =
                                                                                                 empty.clone();
                                                                                             move
                                                                                                 |dictFunctor|
                                                                                                 {
                                                                                                     let altCompose2 =
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&altCompose1,
                                                                                                                                          dictFunctor);
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                                                                      &&&add(string("empty"),
                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_Compose(),
                                                                                                                                                                               &&&empty),
                                                                                                                                             add(string("Alt0"),
                                                                                                                                                 &&Func1::new({
                                                                                                                                                                  let altCompose2
                                                                                                                                                                      =
                                                                                                                                                                      altCompose2.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |usd__unused|
                                                                                                                                                                      &altCompose2
                                                                                                                                                              }),
                                                                                                                                                 empty_1::<string,
                                                                                                                                                           &dyn Any>())))
                                                                                                 }
                                                                                         })
                                                                         }))
    }
    pub fn Data_Functor_Compose_alternativeCompose() -> &dyn Any {
        static Data_Functor_Compose_alternativeCompose:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Compose_alternativeCompose.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictAlternative|
                                                                                {
                                                                                    let applicativeCompose1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_applicativeCompose(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    let plusCompose1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Compose::Data_Functor_Compose_plusCompose(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    &Func1::new({
                                                                                                    let applicativeCompose1
                                                                                                        =
                                                                                                        applicativeCompose1.clone();
                                                                                                    let plusCompose1
                                                                                                        =
                                                                                                        plusCompose1.clone();
                                                                                                    move
                                                                                                        |dictApplicative|
                                                                                                        {
                                                                                                            let applicativeCompose2 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&applicativeCompose1,
                                                                                                                                                 dictApplicative);
                                                                                                            let plusCompose2 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&plusCompose1,
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                                                                             &&&add(string("Applicative0"),
                                                                                                                                                    &&Func1::new({
                                                                                                                                                                     let applicativeCompose2
                                                                                                                                                                         =
                                                                                                                                                                         applicativeCompose2.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |usd__unused|
                                                                                                                                                                         &applicativeCompose2
                                                                                                                                                                 }),
                                                                                                                                                    add(string("Plus1"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let plusCompose2
                                                                                                                                                                             =
                                                                                                                                                                             plusCompose2.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |usd__unused_1|
                                                                                                                                                                             &plusCompose2
                                                                                                                                                                     }),
                                                                                                                                                        empty_1::<string,
                                                                                                                                                                  &dyn Any>())))
                                                                                                        }
                                                                                                })
                                                                                }))
    }
}
