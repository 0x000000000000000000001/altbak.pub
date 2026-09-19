pub mod PureScript_Data_Functor_Clown {
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
    use crate::module_2f2d2d01::PureScript_Control_Biapplicative;
    use crate::module_c52ab4fd::PureScript_Control_Biapply;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_cf56105e::PureScript_Data_Functor_Contravariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Clown_Clown() -> &dyn Any {
        static Data_Functor_Clown_Clown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_Clown.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Functor_Clown_showClown() -> &dyn Any {
        static Data_Functor_Clown_showClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_showClown.get_or_init(||
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
                                                                                                                                      let x =
                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                          &&&string("(Clown ")),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                   &&&dictShow),
                                                                                                                                                                                                                                                                                &&&x)),
                                                                                                                                                                                                          &&&string(")")))
                                                                                                                                  }
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))
    }
    pub fn Data_Functor_Clown_profunctorClown() -> &dyn Any {
        static Data_Functor_Clown_profunctorClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_profunctorClown.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictContravariant|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_Profunctorusd_Dict(),
                                                                                                            &&&add(string("dimap"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let dictContravariant
                                                                                                                                        =
                                                                                                                                        dictContravariant.clone();
                                                                                                                                    move
                                                                                                                                        |f|
                                                                                                                                        &Func1::new({
                                                                                                                                                        let f
                                                                                                                                                            =
                                                                                                                                                            f.clone();
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
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                    let matchValue_1 =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                    let matchValue_2 =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_Clown(),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Contravariant::Data_Functor_Contravariant_cmap(),
                                                                                                                                                                                                                                                                                                                              &&&dictContravariant),
                                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                                        &&&matchValue_2))
                                                                                                                                                                                }
                                                                                                                                                                        })
                                                                                                                                                    })
                                                                                                                                }),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Functor_Clown_ordClown() -> &dyn Any {
        static Data_Functor_Clown_ordClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_ordClown.get_or_init(||
                                                    &Func1::new(move |dictOrd|
                                                                    dictOrd.clone()))
    }
    pub fn Data_Functor_Clown_newtypeClown() -> &dyn Any {
        static Data_Functor_Clown_newtypeClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_newtypeClown.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                         &&&add(string("Coercible0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &Sharpurs_Prelude::Prim_undefined()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Functor_Clown_hoistClown() -> &dyn Any {
        static Data_Functor_Clown_hoistClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_hoistClown.get_or_init(||
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
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_Clown(),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                  &&&matchValue_1))
                                                                                          }
                                                                                  })))
    }
    pub fn Data_Functor_Clown_functorClown() -> &dyn Any {
        static Data_Functor_Clown_functorClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_functorClown.get_or_init(||
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
                                                                                                                                         let matchValue_1 =
                                                                                                                                             Sharpurs_Prelude::unbox(v1);
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_Clown(),
                                                                                                                                                                          &&&matchValue_1)
                                                                                                                                     }
                                                                                                                             })),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Functor_Clown_eqClown() -> &dyn Any {
        static Data_Functor_Clown_eqClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_eqClown.get_or_init(||
                                                   &Func1::new(move |dictEq|
                                                                   dictEq.clone()))
    }
    pub fn Data_Functor_Clown_bifunctorClown() -> &dyn Any {
        static Data_Functor_Clown_bifunctorClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_bifunctorClown.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictFunctor|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                                           &&&add(string("bimap"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let dictFunctor
                                                                                                                                       =
                                                                                                                                       dictFunctor.clone();
                                                                                                                                   move
                                                                                                                                       |f|
                                                                                                                                       &Func1::new({
                                                                                                                                                       let f
                                                                                                                                                           =
                                                                                                                                                           f.clone();
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
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                   let matchValue_1 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                   let matchValue_2 =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_Clown(),
                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                             &&&dictFunctor),
                                                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                                                       &&&matchValue_2))
                                                                                                                                                                               }
                                                                                                                                                                       })
                                                                                                                                                   })
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Functor_Clown_biapplyClown() -> &dyn Any {
        static Data_Functor_Clown_biapplyClown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_biapplyClown.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictApply|
                                                                        {
                                                                            let bifunctorClown1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_bifunctorClown(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_Biapplyusd_Dict(),
                                                                                                             &&&add(string("biapply"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictApply
                                                                                                                                         =
                                                                                                                                         dictApply.clone();
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
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_Clown(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                           &&&dictApply),
                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                     &&&matchValue_1))
                                                                                                                                                             }
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    add(string("Bifunctor0"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let bifunctorClown1
                                                                                                                                             =
                                                                                                                                             bifunctorClown1.clone();
                                                                                                                                         move
                                                                                                                                             |usd__unused|
                                                                                                                                             &bifunctorClown1
                                                                                                                                     }),
                                                                                                                        empty::<string,
                                                                                                                                &dyn Any>())))
                                                                        }))
    }
    pub fn Data_Functor_Clown_biapplicativeClown() -> &dyn Any {
        static Data_Functor_Clown_biapplicativeClown:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Clown_biapplicativeClown.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictApplicative|
                                                                              {
                                                                                  let biapplyClown1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_biapplyClown(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapplicative::Control_Biapplicative_Biapplicativeusd_Dict(),
                                                                                                                   &&&add(string("bipure"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictApplicative
                                                                                                                                               =
                                                                                                                                               dictApplicative.clone();
                                                                                                                                           move
                                                                                                                                               |a|
                                                                                                                                               &Func1::new({
                                                                                                                                                               let a
                                                                                                                                                                   =
                                                                                                                                                                   a.clone();
                                                                                                                                                               move
                                                                                                                                                                   |v|
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Clown::Data_Functor_Clown_Clown(),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                          &&&dictApplicative),
                                                                                                                                                                                                                                       &&&a))
                                                                                                                                                           })
                                                                                                                                       }),
                                                                                                                          add(string("Biapply0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let biapplyClown1
                                                                                                                                                   =
                                                                                                                                                   biapplyClown1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &biapplyClown1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
}
