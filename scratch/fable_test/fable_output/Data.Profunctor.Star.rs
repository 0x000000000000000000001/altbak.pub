pub mod PureScript_Data_Profunctor_Star {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty as empty_1;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_53a2e11::PureScript_Control_MonadPlus;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_83e4823d::PureScript_Data_Distributive;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_8c6f434a::PureScript_Data_Profunctor_Choice;
    use crate::module_6da15eb3::PureScript_Data_Profunctor_Closed;
    use crate::module_15730292::PureScript_Data_Profunctor_Strong;
    use crate::module_2db53acf::PureScript_Data_Profunctor;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_Profunctor_Star_Star() -> &dyn Any {
        static Data_Profunctor_Star_Star: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_Star.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  x.clone()))
    }
    pub fn Data_Profunctor_Star_semigroupoidStar() -> &dyn Any {
        static Data_Profunctor_Star_semigroupoidStar:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_semigroupoidStar.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictBind|
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_Semigroupoidusd_Dict(),
                                                                                                               &&&add(string("compose"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let dictBind
                                                                                                                                           =
                                                                                                                                           dictBind.clone();
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
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                                      let matchValue_1
                                                                                                                                                                                                                          =
                                                                                                                                                                                                                          matchValue_1.clone();
                                                                                                                                                                                                                      move
                                                                                                                                                                                                                          |x|
                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                 &&&dictBind),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                 x)),
                                                                                                                                                                                                                                                           &&&matchValue)
                                                                                                                                                                                                                  }))
                                                                                                                                                               }
                                                                                                                                                       })
                                                                                                                                   }),
                                                                                                                      empty_1::<string,
                                                                                                                                &dyn Any>()))))
    }
    pub fn Data_Profunctor_Star_profunctorStar() -> &dyn Any {
        static Data_Profunctor_Star_profunctorStar: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Profunctor_Star_profunctorStar.get_or_init(||
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
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                               &&&matchValue_2),
                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                  &&&dictFunctor),
                                                                                                                                                                                                                                                                                                                               &&&matchValue_1))))
                                                                                                                                                                                 }
                                                                                                                                                                         })
                                                                                                                                                     })
                                                                                                                                 }),
                                                                                                                    empty_1::<string,
                                                                                                                              &dyn Any>()))))
    }
    pub fn Data_Profunctor_Star_strongStar() -> &dyn Any {
        static Data_Profunctor_Star_strongStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_strongStar.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictFunctor|
                                                                        {
                                                                            let profunctorStar1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_profunctorStar(),
                                                                                                                 dictFunctor);
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Strong::Data_Profunctor_Strong_Strongusd_Dict(),
                                                                                                             &&&add(string("first"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictFunctor
                                                                                                                                         =
                                                                                                                                         dictFunctor.clone();
                                                                                                                                     move
                                                                                                                                         |v|
                                                                                                                                         {
                                                                                                                                             let f =
                                                                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                let f
                                                                                                                                                                                                    =
                                                                                                                                                                                                    f.clone();
                                                                                                                                                                                                move
                                                                                                                                                                                                    |v1|
                                                                                                                                                                                                    {
                                                                                                                                                                                                        let matchValue:
                                                                                                                                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                            Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                               &&&dictFunctor),
                                                                                                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                                                                                                              |v2|
                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(v2.clone(),
                                                                                                                                                                                                                                                                                                                                                      &match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                                       })))),
                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               }))
                                                                                                                                                                                                    }
                                                                                                                                                                                            }))
                                                                                                                                         }
                                                                                                                                 }),
                                                                                                                    add(string("second"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let dictFunctor
                                                                                                                                             =
                                                                                                                                             dictFunctor.clone();
                                                                                                                                         move
                                                                                                                                             |v_1|
                                                                                                                                             {
                                                                                                                                                 let f_1 =
                                                                                                                                                     Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                    let f_1
                                                                                                                                                                                                        =
                                                                                                                                                                                                        f_1.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |v1_1|
                                                                                                                                                                                                        {
                                                                                                                                                                                                            let matchValue_1:
                                                                                                                                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                   &&&dictFunctor),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                     |usd__arg1|
                                                                                                                                                                                                                                                                                                                                     Func1::new({
                                                                                                                                                                                                                                                                                                                                                    let usd__arg1
                                                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                                                        usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                                                        |usd__arg2|
                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                                                usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                                                                                                                                   &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                      })),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                                                                                &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                   }))
                                                                                                                                                                                                        }
                                                                                                                                                                                                }))
                                                                                                                                             }
                                                                                                                                     }),
                                                                                                                        add(string("Profunctor0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let profunctorStar1
                                                                                                                                                 =
                                                                                                                                                 profunctorStar1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &profunctorStar1
                                                                                                                                         }),
                                                                                                                            empty_1::<string,
                                                                                                                                      &dyn Any>()))))
                                                                        }))
    }
    pub fn Data_Profunctor_Star_newtypeStar() -> &dyn Any {
        static Data_Profunctor_Star_newtypeStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_newtypeStar.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                          &&&add(string("Coercible0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &Sharpurs_Prelude::Prim_undefined()),
                                                                                                 empty_1::<string,
                                                                                                           &dyn Any>())))
    }
    pub fn Data_Profunctor_Star_invariantStar() -> &dyn Any {
        static Data_Profunctor_Star_invariantStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_invariantStar.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictInvariant|
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                                            &&&add(string("imap"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let dictInvariant
                                                                                                                                        =
                                                                                                                                        dictInvariant.clone();
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
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imap(),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictInvariant),
                                                                                                                                                                                                                                                                                                                                                                 &&&matchValue),
                                                                                                                                                                                                                                                                                                                              &&&matchValue_1)),
                                                                                                                                                                                                                                                        &&&matchValue_2))
                                                                                                                                                                                }
                                                                                                                                                                        })
                                                                                                                                                    })
                                                                                                                                }),
                                                                                                                   empty_1::<string,
                                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Profunctor_Star_hoistStar() -> &dyn Any {
        static Data_Profunctor_Star_hoistStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_hoistStar.get_or_init(||
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
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                      &&&matchValue),
                                                                                                                                                                   &&&matchValue_1))
                                                                                           }
                                                                                   })))
    }
    pub fn Data_Profunctor_Star_functorStar() -> &dyn Any {
        static Data_Profunctor_Star_functorStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_functorStar.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictFunctor|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                          &&&add(string("map"),
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
                                                                                                                                                          {
                                                                                                                                                              let matchValue =
                                                                                                                                                                  Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                              let matchValue_1 =
                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                           &&&dictFunctor),
                                                                                                                                                                                                                                                                                                        &&&matchValue)),
                                                                                                                                                                                                                                  &&&matchValue_1))
                                                                                                                                                          }
                                                                                                                                                  })
                                                                                                                              }),
                                                                                                                 empty_1::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Profunctor_Star_distributiveStar_004024() -> &dyn Any {
        &Func1::new(move |dictDistributive|
                        PureScript_Data_Profunctor_Star::Data_Profunctor_Star_distributiveStar_tco(dictDistributive))
    }
    pub fn Data_Profunctor_Star_distributiveStar_004024_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Profunctor_Star_distributiveStar_004024_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Profunctor_Star_distributiveStar_004024_002d1.get_or_init(||
                                                                           Lazy(Data_Profunctor_Star_distributiveStar_004024.clone()))
    }
    pub fn Data_Profunctor_Star_distributiveStar_tco(dictDistributive:
                                                         &dyn Any)
     -> &dyn Any {
        let functorStar1 =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_functorStar(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                       Sharpurs_Prelude::unbox(dictDistributive)),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_Distributiveusd_Dict(),
                                         &&&add(string("distribute"),
                                                &&Func1::new({
                                                                 let dictDistributive
                                                                     =
                                                                     dictDistributive.clone();
                                                                 move
                                                                     |dictFunctor|
                                                                     &Func1::new({
                                                                                     let dictFunctor
                                                                                         =
                                                                                         dictFunctor.clone();
                                                                                     move
                                                                                         |f|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                          &&&Func1::new({
                                                                                                                                            let f
                                                                                                                                                =
                                                                                                                                                f.clone();
                                                                                                                                            move
                                                                                                                                                |a|
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_collect(),
                                                                                                                                                                                                                                                                                          &&&dictDistributive),
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
                                                                                 })
                                                             }),
                                                add(string("collect"),
                                                    &&Func1::new({
                                                                     let dictDistributive
                                                                         =
                                                                         dictDistributive.clone();
                                                                     move
                                                                         |dictFunctor_1|
                                                                         &Func1::new({
                                                                                         let dictFunctor_1
                                                                                             =
                                                                                             dictFunctor_1.clone();
                                                                                         move
                                                                                             |f_1|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_distribute(),
                                                                                                                                                                                                                                       &&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_distributiveStar_tco(&&dictDistributive)),
                                                                                                                                                                                                    &&&dictFunctor_1)),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                    &&&dictFunctor_1),
                                                                                                                                                                 f_1))
                                                                                     })
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new({
                                                                         let functorStar1
                                                                             =
                                                                             functorStar1.clone();
                                                                         move
                                                                             |usd__unused|
                                                                             &functorStar1
                                                                     }),
                                                        empty_1::<string,
                                                                  &dyn Any>()))))
    }
    pub fn Data_Profunctor_Star_distributiveStar() -> &dyn Any {
        static Data_Profunctor_Star_distributiveStar:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_distributiveStar.get_or_init(||
                                                              Data_Profunctor_Star_distributiveStar_004024_002d1.Value)
    }
    pub fn Data_Profunctor_Star_closedStar() -> &dyn Any {
        static Data_Profunctor_Star_closedStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_closedStar.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictDistributive|
                                                                        {
                                                                            let profunctorStar1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_profunctorStar(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictDistributive)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Closed::Data_Profunctor_Closed_Closedusd_Dict(),
                                                                                                             &&&add(string("closed"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let dictDistributive
                                                                                                                                         =
                                                                                                                                         dictDistributive.clone();
                                                                                                                                     move
                                                                                                                                         |v|
                                                                                                                                         {
                                                                                                                                             let f =
                                                                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                let f
                                                                                                                                                                                                    =
                                                                                                                                                                                                    f.clone();
                                                                                                                                                                                                move
                                                                                                                                                                                                    |g|
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Distributive::Data_Distributive_distribute(),
                                                                                                                                                                                                                                                                                                           &&&dictDistributive),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_Functor::Data_Functor_functorFn()),
                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                           &&&f),
                                                                                                                                                                                                                                                                        g))
                                                                                                                                                                                            }))
                                                                                                                                         }
                                                                                                                                 }),
                                                                                                                    add(string("Profunctor0"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let profunctorStar1
                                                                                                                                             =
                                                                                                                                             profunctorStar1.clone();
                                                                                                                                         move
                                                                                                                                             |usd__unused|
                                                                                                                                             &profunctorStar1
                                                                                                                                     }),
                                                                                                                        empty_1::<string,
                                                                                                                                  &dyn Any>())))
                                                                        }))
    }
    pub fn Data_Profunctor_Star_choiceStar() -> &dyn Any {
        static Data_Profunctor_Star_choiceStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_choiceStar.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictApplicative|
                                                                        {
                                                                            let Apply0 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                        Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                            let Functor0 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                        Sharpurs_Prelude::unbox(&&Apply0)),
                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                            let pure_var =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                 dictApplicative);
                                                                            let pure1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                 dictApplicative);
                                                                            let profunctorStar1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_profunctorStar(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(&&Apply0)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Choice::Data_Profunctor_Choice_Choiceusd_Dict(),
                                                                                                             &&&add(string("left"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let Functor0
                                                                                                                                         =
                                                                                                                                         Functor0.clone();
                                                                                                                                     let pure_var
                                                                                                                                         =
                                                                                                                                         pure_var.clone();
                                                                                                                                     move
                                                                                                                                         |v|
                                                                                                                                         {
                                                                                                                                             let f =
                                                                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                 &&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star()),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                &&&Functor0),
                                                                                                                                                                                                                                                                                                                                                             &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                               |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                               &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))))),
                                                                                                                                                                                                                                                                                       &&&f)),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                       &&&pure_var),
                                                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                                                      |usd__arg1_1|
                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_1.clone()))))))
                                                                                                                                         }
                                                                                                                                 }),
                                                                                                                    add(string("right"),
                                                                                                                        &&Func1::new({
                                                                                                                                         let Functor0
                                                                                                                                             =
                                                                                                                                             Functor0.clone();
                                                                                                                                         let pure1
                                                                                                                                             =
                                                                                                                                             pure1.clone();
                                                                                                                                         move
                                                                                                                                             |v_1|
                                                                                                                                             {
                                                                                                                                                 let f_1 =
                                                                                                                                                     Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                     &&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star()),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                              &&&pure1),
                                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                                             |usd__arg1_2|
                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_2.clone()))))),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                 &&&Functor0),
                                                                                                                                                                                                                                                                                                                              &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                |usd__arg1_3|
                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_3.clone()))))),
                                                                                                                                                                                                                                                        &&&f_1)))
                                                                                                                                             }
                                                                                                                                     }),
                                                                                                                        add(string("Profunctor0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let profunctorStar1
                                                                                                                                                 =
                                                                                                                                                 profunctorStar1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &profunctorStar1
                                                                                                                                         }),
                                                                                                                            empty_1::<string,
                                                                                                                                      &dyn Any>()))))
                                                                        }))
    }
    pub fn Data_Profunctor_Star_categoryStar() -> &dyn Any {
        static Data_Profunctor_Star_categoryStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_categoryStar.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonad|
                                                                          {
                                                                              let semigroupoidStar1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_semigroupoidStar(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                             Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_Categoryusd_Dict(),
                                                                                                               &&&add(string("identity"),
                                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                      add(string("Semigroupoid0"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let semigroupoidStar1
                                                                                                                                               =
                                                                                                                                               semigroupoidStar1.clone();
                                                                                                                                           move
                                                                                                                                               |usd__unused|
                                                                                                                                               &semigroupoidStar1
                                                                                                                                       }),
                                                                                                                          empty_1::<string,
                                                                                                                                    &dyn Any>())))
                                                                          }))
    }
    pub fn Data_Profunctor_Star_applyStar() -> &dyn Any {
        static Data_Profunctor_Star_applyStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_applyStar.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictApply|
                                                                       {
                                                                           let functorStar1 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_functorStar(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                          Sharpurs_Prelude::unbox(dictApply)),
                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                            &&&add(string("apply"),
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
                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                                                   let matchValue_1
                                                                                                                                                                                                                       =
                                                                                                                                                                                                                       matchValue_1.clone();
                                                                                                                                                                                                                   move
                                                                                                                                                                                                                       |a|
                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                              &&&dictApply),
                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                              a)),
                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                           a))
                                                                                                                                                                                                               }))
                                                                                                                                                            }
                                                                                                                                                    })
                                                                                                                                }),
                                                                                                                   add(string("Functor0"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let functorStar1
                                                                                                                                            =
                                                                                                                                            functorStar1.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused|
                                                                                                                                            &functorStar1
                                                                                                                                    }),
                                                                                                                       empty_1::<string,
                                                                                                                                 &dyn Any>())))
                                                                       }))
    }
    pub fn Data_Profunctor_Star_bindStar() -> &dyn Any {
        static Data_Profunctor_Star_bindStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_bindStar.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictBind|
                                                                      {
                                                                          let applyStar1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_applyStar(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                         Sharpurs_Prelude::unbox(dictBind)),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                           &&&add(string("bind"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let dictBind
                                                                                                                                       =
                                                                                                                                       dictBind.clone();
                                                                                                                                   move
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
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                                  let matchValue_1
                                                                                                                                                                                                                      =
                                                                                                                                                                                                                      matchValue_1.clone();
                                                                                                                                                                                                                  move
                                                                                                                                                                                                                      |x|
                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                             &&&dictBind),
                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                             x)),
                                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                                         let x
                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                             x.clone();
                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                             |a|
                                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                                                            a)),
                                                                                                                                                                                                                                                                                                              &&&x)
                                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                                              }))
                                                                                                                                                           }
                                                                                                                                                   })
                                                                                                                               }),
                                                                                                                  add(string("Apply0"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let applyStar1
                                                                                                                                           =
                                                                                                                                           applyStar1.clone();
                                                                                                                                       move
                                                                                                                                           |usd__unused|
                                                                                                                                           &applyStar1
                                                                                                                                   }),
                                                                                                                      empty_1::<string,
                                                                                                                                &dyn Any>())))
                                                                      }))
    }
    pub fn Data_Profunctor_Star_applicativeStar() -> &dyn Any {
        static Data_Profunctor_Star_applicativeStar: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Profunctor_Star_applicativeStar.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictApplicative|
                                                                             {
                                                                                 let applyStar1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_applyStar(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                                  &&&add(string("pure"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let dictApplicative
                                                                                                                                              =
                                                                                                                                              dictApplicative.clone();
                                                                                                                                          move
                                                                                                                                              |a|
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                 let a
                                                                                                                                                                                                     =
                                                                                                                                                                                                     a.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |v|
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                         &&&dictApplicative),
                                                                                                                                                                                                                                      &&&a)
                                                                                                                                                                                             }))
                                                                                                                                      }),
                                                                                                                         add(string("Apply0"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let applyStar1
                                                                                                                                                  =
                                                                                                                                                  applyStar1.clone();
                                                                                                                                              move
                                                                                                                                                  |usd__unused|
                                                                                                                                                  &applyStar1
                                                                                                                                          }),
                                                                                                                             empty_1::<string,
                                                                                                                                       &dyn Any>())))
                                                                             }))
    }
    pub fn Data_Profunctor_Star_monadStar() -> &dyn Any {
        static Data_Profunctor_Star_monadStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_monadStar.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonad|
                                                                       {
                                                                           let applicativeStar1 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_applicativeStar(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                           let bindStar1 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_bindStar(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                            &&&add(string("Applicative0"),
                                                                                                                   &&Func1::new({
                                                                                                                                    let applicativeStar1
                                                                                                                                        =
                                                                                                                                        applicativeStar1.clone();
                                                                                                                                    move
                                                                                                                                        |usd__unused|
                                                                                                                                        &applicativeStar1
                                                                                                                                }),
                                                                                                                   add(string("Bind1"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let bindStar1
                                                                                                                                            =
                                                                                                                                            bindStar1.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused_1|
                                                                                                                                            &bindStar1
                                                                                                                                    }),
                                                                                                                       empty_1::<string,
                                                                                                                                 &dyn Any>())))
                                                                       }))
    }
    pub fn Data_Profunctor_Star_altStar() -> &dyn Any {
        static Data_Profunctor_Star_altStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_altStar.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictAlt|
                                                                     {
                                                                         let functorStar1 =
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_functorStar(),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                        Sharpurs_Prelude::unbox(dictAlt)),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                                          &&&add(string("alt"),
                                                                                                                 &&Func1::new({
                                                                                                                                  let dictAlt
                                                                                                                                      =
                                                                                                                                      dictAlt.clone();
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
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                 let matchValue_1
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     matchValue_1.clone();
                                                                                                                                                                                                                 move
                                                                                                                                                                                                                     |a|
                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                                                                                                                                            &&&dictAlt),
                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                            a)),
                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                         a))
                                                                                                                                                                                                             }))
                                                                                                                                                          }
                                                                                                                                                  })
                                                                                                                              }),
                                                                                                                 add(string("Functor0"),
                                                                                                                     &&Func1::new({
                                                                                                                                      let functorStar1
                                                                                                                                          =
                                                                                                                                          functorStar1.clone();
                                                                                                                                      move
                                                                                                                                          |usd__unused|
                                                                                                                                          &functorStar1
                                                                                                                                  }),
                                                                                                                     empty_1::<string,
                                                                                                                               &dyn Any>())))
                                                                     }))
    }
    pub fn Data_Profunctor_Star_plusStar() -> &dyn Any {
        static Data_Profunctor_Star_plusStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_plusStar.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictPlus|
                                                                      {
                                                                          let empty =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                               dictPlus);
                                                                          let altStar1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_altStar(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                         Sharpurs_Prelude::unbox(dictPlus)),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                                           &&&add(string("empty"),
                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_Star(),
                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                      let empty
                                                                                                                                                                          =
                                                                                                                                                                          empty.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |v|
                                                                                                                                                                          &empty
                                                                                                                                                                  })),
                                                                                                                  add(string("Alt0"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let altStar1
                                                                                                                                           =
                                                                                                                                           altStar1.clone();
                                                                                                                                       move
                                                                                                                                           |usd__unused|
                                                                                                                                           &altStar1
                                                                                                                                   }),
                                                                                                                      empty_1::<string,
                                                                                                                                &dyn Any>())))
                                                                      }))
    }
    pub fn Data_Profunctor_Star_alternativeStar() -> &dyn Any {
        static Data_Profunctor_Star_alternativeStar: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Profunctor_Star_alternativeStar.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictAlternative|
                                                                             {
                                                                                 let applicativeStar1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_applicativeStar(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 let plusStar1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_plusStar(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                                                  &&&add(string("Applicative0"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let applicativeStar1
                                                                                                                                              =
                                                                                                                                              applicativeStar1.clone();
                                                                                                                                          move
                                                                                                                                              |usd__unused|
                                                                                                                                              &applicativeStar1
                                                                                                                                      }),
                                                                                                                         add(string("Plus1"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let plusStar1
                                                                                                                                                  =
                                                                                                                                                  plusStar1.clone();
                                                                                                                                              move
                                                                                                                                                  |usd__unused_1|
                                                                                                                                                  &plusStar1
                                                                                                                                          }),
                                                                                                                             empty_1::<string,
                                                                                                                                       &dyn Any>())))
                                                                             }))
    }
    pub fn Data_Profunctor_Star_monadPlusStar() -> &dyn Any {
        static Data_Profunctor_Star_monadPlusStar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Profunctor_Star_monadPlusStar.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictMonadPlus|
                                                                           {
                                                                               let monadStar1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_monadStar(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonadPlus)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                               let alternativeStar1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Profunctor_Star::Data_Profunctor_Star_alternativeStar(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alternative1"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictMonadPlus)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_MonadPlus::Control_MonadPlus_MonadPlususd_Dict(),
                                                                                                                &&&add(string("Monad0"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let monadStar1
                                                                                                                                            =
                                                                                                                                            monadStar1.clone();
                                                                                                                                        move
                                                                                                                                            |usd__unused|
                                                                                                                                            &monadStar1
                                                                                                                                    }),
                                                                                                                       add(string("Alternative1"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let alternativeStar1
                                                                                                                                                =
                                                                                                                                                alternativeStar1.clone();
                                                                                                                                            move
                                                                                                                                                |usd__unused_1|
                                                                                                                                                &alternativeStar1
                                                                                                                                        }),
                                                                                                                           empty_1::<string,
                                                                                                                                     &dyn Any>())))
                                                                           }))
    }
}
