pub mod PureScript_Data_Functor_Flip {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_2f2d2d01::PureScript_Control_Biapplicative;
    use crate::module_c52ab4fd::PureScript_Control_Biapply;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_cf56105e::PureScript_Data_Functor_Contravariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Flip_Flip() -> &dyn Any {
        static Data_Functor_Flip_Flip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_Flip.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Functor_Flip_showFlip() -> &dyn Any {
        static Data_Functor_Flip_showFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_showFlip.get_or_init(||
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
                                                                                                                                    let x =
                                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                        &&&string("(Flip ")),
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
    pub fn Data_Functor_Flip_semigroupoidFlip() -> &dyn Any {
        static Data_Functor_Flip_semigroupoidFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_semigroupoidFlip.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictSemigroupoid|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_Semigroupoidusd_Dict(),
                                                                                                            &&&add(string("compose"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let dictSemigroupoid
                                                                                                                                        =
                                                                                                                                        dictSemigroupoid.clone();
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
                                                                                                                                                                                                                                    &&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip()),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                          &&&dictSemigroupoid),
                                                                                                                                                                                                                                                                       &&&matchValue_1),
                                                                                                                                                                                                                                    &&&matchValue))
                                                                                                                                                            }
                                                                                                                                                    })
                                                                                                                                }),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Functor_Flip_ordFlip() -> &dyn Any {
        static Data_Functor_Flip_ordFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_ordFlip.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  dictOrd.clone()))
    }
    pub fn Data_Functor_Flip_newtypeFlip() -> &dyn Any {
        static Data_Functor_Flip_newtypeFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_newtypeFlip.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                       &&&add(string("Coercible0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &Sharpurs_Prelude::Prim_undefined()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Functor_Flip_functorFlip() -> &dyn Any {
        static Data_Functor_Flip_functorFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_functorFlip.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictBifunctor|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                       &&&add(string("map"),
                                                                                                              &&Func1::new({
                                                                                                                               let dictBifunctor
                                                                                                                                   =
                                                                                                                                   dictBifunctor.clone();
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
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip(),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_lmap(),
                                                                                                                                                                                                                                                                                                     &&&dictBifunctor),
                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                               &&&matchValue_1))
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                           }),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Functor_Flip_eqFlip() -> &dyn Any {
        static Data_Functor_Flip_eqFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_eqFlip.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 dictEq.clone()))
    }
    pub fn Data_Functor_Flip_contravariantFlip() -> &dyn Any {
        static Data_Functor_Flip_contravariantFlip: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Flip_contravariantFlip.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictProfunctor|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Contravariant::Data_Functor_Contravariant_Contravariantusd_Dict(),
                                                                                                             &&&add(string("cmap"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictProfunctor
                                                                                                                                         =
                                                                                                                                         dictProfunctor.clone();
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
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_lcmap(),
                                                                                                                                                                                                                                                                                                           &&&dictProfunctor),
                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                     &&&matchValue_1))
                                                                                                                                                             }
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Functor_Flip_categoryFlip() -> &dyn Any {
        static Data_Functor_Flip_categoryFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_categoryFlip.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictCategory|
                                                                       {
                                                                           let semigroupoidFlip1 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_semigroupoidFlip(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroupoid0"),
                                                                                                                                                          Sharpurs_Prelude::unbox(dictCategory)),
                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_Categoryusd_Dict(),
                                                                                                            &&&add(string("identity"),
                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip(),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                        dictCategory)),
                                                                                                                   add(string("Semigroupoid0"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let semigroupoidFlip1
                                                                                                                                            =
                                                                                                                                            semigroupoidFlip1.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused|
                                                                                                                                            &semigroupoidFlip1
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>())))
                                                                       }))
    }
    pub fn Data_Functor_Flip_bifunctorFlip() -> &dyn Any {
        static Data_Functor_Flip_bifunctorFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_bifunctorFlip.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictBifunctor|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                                         &&&add(string("bimap"),
                                                                                                                &&Func1::new({
                                                                                                                                 let dictBifunctor
                                                                                                                                     =
                                                                                                                                     dictBifunctor.clone();
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
                                                                                                                                                                             |v|
                                                                                                                                                                             {
                                                                                                                                                                                 let matchValue =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                 let matchValue_1 =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(&&g);
                                                                                                                                                                                 let matchValue_2 =
                                                                                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip(),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                              &&&dictBifunctor),
                                                                                                                                                                                                                                                                                                                           &&&matchValue_1),
                                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                                     &&&matchValue_2))
                                                                                                                                                                             }
                                                                                                                                                                     })
                                                                                                                                                 })
                                                                                                                             }),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Functor_Flip_biapplyFlip() -> &dyn Any {
        static Data_Functor_Flip_biapplyFlip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Flip_biapplyFlip.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictBiapply|
                                                                      {
                                                                          let bifunctorFlip1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_bifunctorFlip(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bifunctor0"),
                                                                                                                                                         Sharpurs_Prelude::unbox(dictBiapply)),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_Biapplyusd_Dict(),
                                                                                                           &&&add(string("biapply"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let dictBiapply
                                                                                                                                       =
                                                                                                                                       dictBiapply.clone();
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
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip(),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapply::Control_Biapply_biapply(),
                                                                                                                                                                                                                                                                                                         &&&dictBiapply),
                                                                                                                                                                                                                                                                      &&&matchValue),
                                                                                                                                                                                                                                   &&&matchValue_1))
                                                                                                                                                           }
                                                                                                                                                   })
                                                                                                                               }),
                                                                                                                  add(string("Bifunctor0"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let bifunctorFlip1
                                                                                                                                           =
                                                                                                                                           bifunctorFlip1.clone();
                                                                                                                                       move
                                                                                                                                           |usd__unused|
                                                                                                                                           &bifunctorFlip1
                                                                                                                                   }),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>())))
                                                                      }))
    }
    pub fn Data_Functor_Flip_biapplicativeFlip() -> &dyn Any {
        static Data_Functor_Flip_biapplicativeFlip: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Flip_biapplicativeFlip.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictBiapplicative|
                                                                            {
                                                                                let biapplyFlip1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_biapplyFlip(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Biapply0"),
                                                                                                                                                               Sharpurs_Prelude::unbox(dictBiapplicative)),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapplicative::Control_Biapplicative_Biapplicativeusd_Dict(),
                                                                                                                 &&&add(string("bipure"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let dictBiapplicative
                                                                                                                                             =
                                                                                                                                             dictBiapplicative.clone();
                                                                                                                                         move
                                                                                                                                             |a|
                                                                                                                                             &Func1::new({
                                                                                                                                                             let a
                                                                                                                                                                 =
                                                                                                                                                                 a.clone();
                                                                                                                                                             move
                                                                                                                                                                 |b|
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Flip::Data_Functor_Flip_Flip(),
                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Biapplicative::Control_Biapplicative_bipure(),
                                                                                                                                                                                                                                                                                                           &&&dictBiapplicative),
                                                                                                                                                                                                                                                                        b),
                                                                                                                                                                                                                                     &&&a))
                                                                                                                                                         })
                                                                                                                                     }),
                                                                                                                        add(string("Biapply0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let biapplyFlip1
                                                                                                                                                 =
                                                                                                                                                 biapplyFlip1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &biapplyFlip1
                                                                                                                                         }),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>())))
                                                                            }))
    }
}
