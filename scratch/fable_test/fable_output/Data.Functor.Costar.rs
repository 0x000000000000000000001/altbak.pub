pub mod PureScript_Data_Functor_Costar {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_83e4823d::PureScript_Data_Distributive;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_cf56105e::PureScript_Data_Functor_Contravariant;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_6da15eb3::PureScript_Data_Profunctor_Closed;
    use crate::module_15730292::PureScript_Data_Profunctor_Strong;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_Functor_Costar_Costar() -> &dyn Any {
        static Data_Functor_Costar_Costar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_Costar.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Functor_Costar_semigroupoidCostar() -> &dyn Any {
        static Data_Functor_Costar_semigroupoidCostar:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_semigroupoidCostar.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictExtend|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_Semigroupoidusd_Dict(),
                                                                                                                &&&add(string("compose"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let dictExtend
                                                                                                                                            =
                                                                                                                                            dictExtend.clone();
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
                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_composeCoKleisliFlipped(),
                                                                                                                                                                                                                                                                                                              &&&dictExtend),
                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                        &&&matchValue_1))
                                                                                                                                                                }
                                                                                                                                                        })
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Functor_Costar_profunctorCostar() -> &dyn Any {
        static Data_Functor_Costar_profunctorCostar: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Costar_profunctorCostar.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictFunctor|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_Profunctorusd_Dict(),
                                                                                                              &&&add(string("dimap"),
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
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                   &&&dictFunctor),
                                                                                                                                                                                                                                                                                                                                &&&matchValue)),
                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                &&&matchValue_2),
                                                                                                                                                                                                                                                                                             &&&matchValue_1)))
                                                                                                                                                                                  }
                                                                                                                                                                          })
                                                                                                                                                      })
                                                                                                                                  }),
                                                                                                                     empty::<string,
                                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Functor_Costar_strongCostar() -> &dyn Any {
        static Data_Functor_Costar_strongCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_strongCostar.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictComonad|
                                                                         {
                                                                             let Extend0 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                         Sharpurs_Prelude::unbox(dictComonad)),
                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                                             let Functor0 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                         Sharpurs_Prelude::unbox(&&Extend0)),
                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                                             let profunctorCostar1 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_profunctorCostar(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(&&Extend0)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Strong::Data_Profunctor_Strong_Strongusd_Dict(),
                                                                                                              &&&add(string("first"),
                                                                                                                     &&Func1::new({
                                                                                                                                      let Functor0
                                                                                                                                          =
                                                                                                                                          Functor0.clone();
                                                                                                                                      let dictComonad
                                                                                                                                          =
                                                                                                                                          dictComonad.clone();
                                                                                                                                      move
                                                                                                                                          |v|
                                                                                                                                          {
                                                                                                                                              let f =
                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                 let f
                                                                                                                                                                                                     =
                                                                                                                                                                                                     f.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |x|
                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Tuple::Data_Tuple_fst()),
                                                                                                                                                                                                                                                                                                                                 x)),
                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_snd(),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                                                                                                                                                                    &&&dictComonad),
                                                                                                                                                                                                                                                                                                                                 x))))
                                                                                                                                                                                             }))
                                                                                                                                          }
                                                                                                                                  }),
                                                                                                                     add(string("second"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let Functor0
                                                                                                                                              =
                                                                                                                                              Functor0.clone();
                                                                                                                                          let dictComonad
                                                                                                                                              =
                                                                                                                                              dictComonad.clone();
                                                                                                                                          move
                                                                                                                                              |v_1|
                                                                                                                                              {
                                                                                                                                                  let f_1 =
                                                                                                                                                      Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                     let f_1
                                                                                                                                                                                                         =
                                                                                                                                                                                                         f_1.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |x_1|
                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_fst(),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                                                                                                                                                                        &&&dictComonad),
                                                                                                                                                                                                                                                                                                                                     x_1)),
                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                           &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                                                                                                                                                                                                                     x_1))))
                                                                                                                                                                                                 }))
                                                                                                                                              }
                                                                                                                                      }),
                                                                                                                         add(string("Profunctor0"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let profunctorCostar1
                                                                                                                                                  =
                                                                                                                                                  profunctorCostar1.clone();
                                                                                                                                              move
                                                                                                                                                  |usd__unused|
                                                                                                                                                  &profunctorCostar1
                                                                                                                                          }),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>()))))
                                                                         }))
    }
    pub fn Data_Functor_Costar_newtypeCostar() -> &dyn Any {
        static Data_Functor_Costar_newtypeCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_newtypeCostar.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                           &&&add(string("Coercible0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &Sharpurs_Prelude::Prim_undefined()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Functor_Costar_hoistCostar() -> &dyn Any {
        static Data_Functor_Costar_hoistCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_hoistCostar.get_or_init(||
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
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor::Data_Profunctor_lcmap(),
                                                                                                                                                                                                                                          &&&PureScript_Data_Profunctor::Data_Profunctor_profunctorFn()),
                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                    &&&matchValue_1))
                                                                                            }
                                                                                    })))
    }
    pub fn Data_Functor_Costar_functorCostar() -> &dyn Any {
        static Data_Functor_Costar_functorCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_functorCostar.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                           &&&add(string("map"),
                                                                                                  &&Func1::new(move
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
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                               &&&matchValue_1))
                                                                                                                                       }
                                                                                                                               })),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Functor_Costar_invariantCostar() -> &dyn Any {
        static Data_Functor_Costar_invariantCostar: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Costar_invariantCostar.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                             &&&add(string("imap"),
                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                                      &&&PureScript_Data_Functor_Costar::Data_Functor_Costar_functorCostar()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Functor_Costar_distributiveCostar_004023() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_Distributiveusd_Dict(),
                                         &&&add(string("distribute"),
                                                &&Func1::new(move
                                                                 |dictFunctor|
                                                                 &Func1::new({
                                                                                 let dictFunctor
                                                                                     =
                                                                                     dictFunctor.clone();
                                                                                 move
                                                                                     |f|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                      &&&Func1::new({
                                                                                                                                        let f
                                                                                                                                            =
                                                                                                                                            f.clone();
                                                                                                                                        move
                                                                                                                                            |a|
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                   &&&dictFunctor),
                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                  let a
                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                      a.clone();
                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                      |v|
                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(v),
                                                                                                                                                                                                                                                                       &&&a)
                                                                                                                                                                                                                              })),
                                                                                                                                                                             &&&f)
                                                                                                                                    }))
                                                                             })),
                                                add(string("collect"),
                                                    &&Func1::new({
                                                                     let Data_Functor_Costar_distributiveCostar_004023_002d1
                                                                         =
                                                                         Data_Functor_Costar_distributiveCostar_004023_002d1.clone();
                                                                     move
                                                                         |dictFunctor_1|
                                                                         &Func1::new({
                                                                                         let Data_Functor_Costar_distributiveCostar_004023_002d1
                                                                                             =
                                                                                             Data_Functor_Costar_distributiveCostar_004023_002d1.clone();
                                                                                         let dictFunctor_1
                                                                                             =
                                                                                             dictFunctor_1.clone();
                                                                                         move
                                                                                             |f_1|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_distribute(),
                                                                                                                                                                                                                                       &&&Data_Functor_Costar_distributiveCostar_004023_002d1.Value),
                                                                                                                                                                                                    &&&dictFunctor_1)),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                    &&&dictFunctor_1),
                                                                                                                                                                 f_1))
                                                                                     })
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Functor_Costar::Data_Functor_Costar_functorCostar()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Functor_Costar_distributiveCostar_004023_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Functor_Costar_distributiveCostar_004023_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Functor_Costar_distributiveCostar_004023_002d1.get_or_init(||
                                                                            Lazy(Data_Functor_Costar_distributiveCostar_004023.clone()))
    }
    pub fn Data_Functor_Costar_distributiveCostar() -> &dyn Any {
        static Data_Functor_Costar_distributiveCostar:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_distributiveCostar.get_or_init(||
                                                               Data_Functor_Costar_distributiveCostar_004023_002d1.Value)
    }
    pub fn Data_Functor_Costar_closedCostar() -> &dyn Any {
        static Data_Functor_Costar_closedCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_closedCostar.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictFunctor|
                                                                         {
                                                                             let profunctorCostar1 =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_profunctorCostar(),
                                                                                                                  dictFunctor);
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Closed::Data_Profunctor_Closed_Closedusd_Dict(),
                                                                                                              &&&add(string("closed"),
                                                                                                                     &&Func1::new({
                                                                                                                                      let dictFunctor
                                                                                                                                          =
                                                                                                                                          dictFunctor.clone();
                                                                                                                                      move
                                                                                                                                          |v|
                                                                                                                                          {
                                                                                                                                              let f =
                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                               &&&Func1::new({
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
                                                                                                                                                                                                                         |x|
                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                   &&&dictFunctor),
                                                                                                                                                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                                                                                                                                                  let x
                                                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                                                      x.clone();
                                                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                                                      |v1|
                                                                                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                          v1),
                                                                                                                                                                                                                                                                                                                                                                                       &&&x)
                                                                                                                                                                                                                                                                                                                                              })),
                                                                                                                                                                                                                                                                                             &&&g))
                                                                                                                                                                                                                 })
                                                                                                                                                                                             }))
                                                                                                                                          }
                                                                                                                                  }),
                                                                                                                     add(string("Profunctor0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let profunctorCostar1
                                                                                                                                              =
                                                                                                                                              profunctorCostar1.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &profunctorCostar1
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
                                                                         }))
    }
    pub fn Data_Functor_Costar_categoryCostar() -> &dyn Any {
        static Data_Functor_Costar_categoryCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_categoryCostar.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictComonad|
                                                                           {
                                                                               let semigroupoidCostar1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_semigroupoidCostar(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictComonad)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_Categoryusd_Dict(),
                                                                                                                &&&add(string("identity"),
                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                            dictComonad)),
                                                                                                                       add(string("Semigroupoid0"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let semigroupoidCostar1
                                                                                                                                                =
                                                                                                                                                semigroupoidCostar1.clone();
                                                                                                                                            move
                                                                                                                                                |usd__unused|
                                                                                                                                                &semigroupoidCostar1
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>())))
                                                                           }))
    }
    pub fn Data_Functor_Costar_bifunctorCostar() -> &dyn Any {
        static Data_Functor_Costar_bifunctorCostar: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Costar_bifunctorCostar.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictContravariant|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                                             &&&add(string("bimap"),
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
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Contravariant::Data_Functor_Contravariant_cmap(),
                                                                                                                                                                                                                                                                                                                                                                  &&&dictContravariant),
                                                                                                                                                                                                                                                                                                                               &&&matchValue)),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                               &&&matchValue_2),
                                                                                                                                                                                                                                                                                            &&&matchValue_1)))
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Functor_Costar_applyCostar() -> &dyn Any {
        static Data_Functor_Costar_applyCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_applyCostar.get_or_init(||
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
                                                                                                                                         let matchValue =
                                                                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                                                                         let matchValue_1 =
                                                                                                                                             Sharpurs_Prelude::unbox(v1);
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                          &&&Func1::new({
                                                                                                                                                                                            let matchValue_1
                                                                                                                                                                                                =
                                                                                                                                                                                                matchValue_1.clone();
                                                                                                                                                                                            move
                                                                                                                                                                                                |a|
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                    a),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                    a))
                                                                                                                                                                                        }))
                                                                                                                                     }
                                                                                                                             })),
                                                                                                add(string("Functor0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &PureScript_Data_Functor_Costar::Data_Functor_Costar_functorCostar()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Functor_Costar_bindCostar() -> &dyn Any {
        static Data_Functor_Costar_bindCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_bindCostar.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                        &&&add(string("bind"),
                                                                                               &&Func1::new(move
                                                                                                                |v|
                                                                                                                &Func1::new({
                                                                                                                                let v
                                                                                                                                    =
                                                                                                                                    v.clone();
                                                                                                                                move
                                                                                                                                    |f|
                                                                                                                                    {
                                                                                                                                        let matchValue =
                                                                                                                                            Sharpurs_Prelude::unbox(&&v);
                                                                                                                                        let matchValue_1 =
                                                                                                                                            Sharpurs_Prelude::unbox(f);
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                                         &&&Func1::new({
                                                                                                                                                                                           let matchValue_1
                                                                                                                                                                                               =
                                                                                                                                                                                               matchValue_1.clone();
                                                                                                                                                                                           move
                                                                                                                                                                                               |x|
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                 x))),
                                                                                                                                                                                                                                x)
                                                                                                                                                                                       }))
                                                                                                                                    }
                                                                                                                            })),
                                                                                               add(string("Apply0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_Functor_Costar::Data_Functor_Costar_applyCostar()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Functor_Costar_applicativeCostar() -> &dyn Any {
        static Data_Functor_Costar_applicativeCostar:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_applicativeCostar.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                               &&&add(string("pure"),
                                                                                                      &&Func1::new(move
                                                                                                                       |a|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Costar::Data_Functor_Costar_Costar(),
                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                          let a
                                                                                                                                                                              =
                                                                                                                                                                              a.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v|
                                                                                                                                                                              &a
                                                                                                                                                                      }))),
                                                                                                      add(string("Apply0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Functor_Costar::Data_Functor_Costar_applyCostar()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Functor_Costar_monadCostar() -> &dyn Any {
        static Data_Functor_Costar_monadCostar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Costar_monadCostar.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                         &&&add(string("Applicative0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Functor_Costar::Data_Functor_Costar_applicativeCostar()),
                                                                                                add(string("Bind1"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused_1|
                                                                                                                     &PureScript_Data_Functor_Costar::Data_Functor_Costar_bindCostar()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
}
