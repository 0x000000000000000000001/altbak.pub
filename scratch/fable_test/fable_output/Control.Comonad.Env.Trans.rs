pub mod PureScript_Control_Comonad_Env_Trans {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_35294a53::PureScript_Control_Comonad_Trans_Class;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_9201da02::PureScript_Data_FoldableWithIndex;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f8b1f47e::PureScript_Data_FunctorWithIndex;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_829cacf6::PureScript_Data_TraversableWithIndex;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Env_Trans_EnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_EnvT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_EnvT.get_or_init(||
                                                       &Func1::new(move |x|
                                                                       x.clone()))
    }
    pub fn Control_Comonad_Env_Trans_withEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_withEnvT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_withEnvT.get_or_init(||
                                                           &Func1::new(move
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
                                                                                                   let matchValue_1:
                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                       &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
                                                                                                                                    &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                  }),
                                                                                                                                                                                              &match matchValue_1.as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                               })))
                                                                                               }
                                                                                       })))
    }
    pub fn Control_Comonad_Env_Trans_runEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_runEnvT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_runEnvT.get_or_init(||
                                                          &Func1::new(move |v|
                                                                          &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Comonad_Env_Trans_newtypeEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_newtypeEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_newtypeEnvT.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                               &&&add(string("Coercible0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &Sharpurs_Prelude::Prim_undefined()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))
    }
    pub fn Control_Comonad_Env_Trans_mapEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_mapEnvT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_mapEnvT.get_or_init(||
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
                                                                                                  let matchValue_1:
                                                                                                          LrcPtr<Data_Tuple_Tuple> =
                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                      &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
                                                                                                                                   &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                  =>
                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                              },
                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                 }))))
                                                                                              }
                                                                                      })))
    }
    pub fn Control_Comonad_Env_Trans_functorEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_functorEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_functorEnvT.get_or_init(||
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
                                                                                                                                                                   let matchValue_1:
                                                                                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                       &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
                                                                                                                                                                                                    &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                               },
                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                     &&&dictFunctor),
                                                                                                                                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                  }))))
                                                                                                                                                               }
                                                                                                                                                       })
                                                                                                                                   }),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
    }
    pub fn Control_Comonad_Env_Trans_functorWithIndexEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_functorWithIndexEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_functorWithIndexEnvT.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictFunctorWithIndex|
                                                                                       {
                                                                                           let functorEnvT1 =
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_functorEnvT(),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                          Sharpurs_Prelude::unbox(dictFunctorWithIndex)),
                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_FunctorWithIndexusd_Dict(),
                                                                                                                            &&&add(string("mapWithIndex"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let dictFunctorWithIndex
                                                                                                                                                        =
                                                                                                                                                        dictFunctorWithIndex.clone();
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
                                                                                                                                                                                let matchValue_1:
                                                                                                                                                                                        LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                    &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
                                                                                                                                                                                                                 &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                            },
                                                                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FunctorWithIndex::Data_FunctorWithIndex_mapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                  &&&dictFunctorWithIndex),
                                                                                                                                                                                                                                                                                                                                               &&&matchValue),
                                                                                                                                                                                                                                                                                                            &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                               }))))
                                                                                                                                                                            }
                                                                                                                                                                    })
                                                                                                                                                }),
                                                                                                                                   add(string("Functor0"),
                                                                                                                                       &&Func1::new({
                                                                                                                                                        let functorEnvT1
                                                                                                                                                            =
                                                                                                                                                            functorEnvT1.clone();
                                                                                                                                                        move
                                                                                                                                                            |usd__unused|
                                                                                                                                                            &functorEnvT1
                                                                                                                                                    }),
                                                                                                                                       empty::<string,
                                                                                                                                               &dyn Any>())))
                                                                                       }))
    }
    pub fn Control_Comonad_Env_Trans_foldableEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_foldableEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_foldableEnvT.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictFoldable|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                                                                                                &&&add(string("foldl"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let dictFoldable
                                                                                                                                            =
                                                                                                                                            dictFoldable.clone();
                                                                                                                                        move
                                                                                                                                            |r#fn|
                                                                                                                                            &Func1::new({
                                                                                                                                                            let r#fn
                                                                                                                                                                =
                                                                                                                                                                r#fn.clone();
                                                                                                                                                            move
                                                                                                                                                                |a|
                                                                                                                                                                &Func1::new({
                                                                                                                                                                                let a
                                                                                                                                                                                    =
                                                                                                                                                                                    a.clone();
                                                                                                                                                                                move
                                                                                                                                                                                    |v|
                                                                                                                                                                                    {
                                                                                                                                                                                        let matchValue =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&r#fn);
                                                                                                                                                                                        let matchValue_1 =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&a);
                                                                                                                                                                                        let matchValue_2:
                                                                                                                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                                                                  &&&dictFoldable),
                                                                                                                                                                                                                                                                                               &&&matchValue),
                                                                                                                                                                                                                                                            &&&matchValue_1),
                                                                                                                                                                                                                         &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                            })
                                                                                                                                                                                    }
                                                                                                                                                                            })
                                                                                                                                                        })
                                                                                                                                    }),
                                                                                                                       add(string("foldr"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let dictFoldable
                                                                                                                                                =
                                                                                                                                                dictFoldable.clone();
                                                                                                                                            move
                                                                                                                                                |fn_1|
                                                                                                                                                &Func1::new({
                                                                                                                                                                let fn_1
                                                                                                                                                                    =
                                                                                                                                                                    fn_1.clone();
                                                                                                                                                                move
                                                                                                                                                                    |a_1|
                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                    let a_1
                                                                                                                                                                                        =
                                                                                                                                                                                        a_1.clone();
                                                                                                                                                                                    move
                                                                                                                                                                                        |v_1|
                                                                                                                                                                                        {
                                                                                                                                                                                            let matchValue_4 =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&fn_1);
                                                                                                                                                                                            let matchValue_5 =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&a_1);
                                                                                                                                                                                            let matchValue_6:
                                                                                                                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                                                                                      &&&dictFoldable),
                                                                                                                                                                                                                                                                                                   &&&matchValue_4),
                                                                                                                                                                                                                                                                &&&matchValue_5),
                                                                                                                                                                                                                             &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                })
                                                                                                                                                                                        }
                                                                                                                                                                                })
                                                                                                                                                            })
                                                                                                                                        }),
                                                                                                                           add(string("foldMap"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let dictFoldable
                                                                                                                                                    =
                                                                                                                                                    dictFoldable.clone();
                                                                                                                                                move
                                                                                                                                                    |dictMonoid|
                                                                                                                                                    &Func1::new({
                                                                                                                                                                    let dictMonoid
                                                                                                                                                                        =
                                                                                                                                                                        dictMonoid.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |fn_2|
                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                        let fn_2
                                                                                                                                                                                            =
                                                                                                                                                                                            fn_2.clone();
                                                                                                                                                                                        move
                                                                                                                                                                                            |v_2|
                                                                                                                                                                                            {
                                                                                                                                                                                                let matchValue_8 =
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&fn_2);
                                                                                                                                                                                                let matchValue_9:
                                                                                                                                                                                                        LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMap(),
                                                                                                                                                                                                                                                                                                                                          &&&dictFoldable),
                                                                                                                                                                                                                                                                                                       &&&dictMonoid),
                                                                                                                                                                                                                                                                    &&&matchValue_8),
                                                                                                                                                                                                                                 &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    })
                                                                                                                                                                                            }
                                                                                                                                                                                    })
                                                                                                                                                                })
                                                                                                                                            }),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))))))
    }
    pub fn Control_Comonad_Env_Trans_foldableWithIndexEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_foldableWithIndexEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_foldableWithIndexEnvT.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictFoldableWithIndex|
                                                                                        {
                                                                                            let foldableEnvT1 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_foldableEnvT(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                           Sharpurs_Prelude::unbox(dictFoldableWithIndex)),
                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_FoldableWithIndexusd_Dict(),
                                                                                                                             &&&add(string("foldlWithIndex"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let dictFoldableWithIndex
                                                                                                                                                         =
                                                                                                                                                         dictFoldableWithIndex.clone();
                                                                                                                                                     move
                                                                                                                                                         |f|
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
                                                                                                                                                                                                 |v|
                                                                                                                                                                                                 {
                                                                                                                                                                                                     let matchValue =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                                     let matchValue_1 =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&a);
                                                                                                                                                                                                     let matchValue_2:
                                                                                                                                                                                                             LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldlWithIndex(),
                                                                                                                                                                                                                                                                                                                                               &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                                                                                                                         &&&matchValue_1),
                                                                                                                                                                                                                                      &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                         })
                                                                                                                                                                                                 }
                                                                                                                                                                                         })
                                                                                                                                                                     })
                                                                                                                                                 }),
                                                                                                                                    add(string("foldrWithIndex"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let dictFoldableWithIndex
                                                                                                                                                             =
                                                                                                                                                             dictFoldableWithIndex.clone();
                                                                                                                                                         move
                                                                                                                                                             |f_1|
                                                                                                                                                             &Func1::new({
                                                                                                                                                                             let f_1
                                                                                                                                                                                 =
                                                                                                                                                                                 f_1.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |a_1|
                                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                                 let a_1
                                                                                                                                                                                                     =
                                                                                                                                                                                                     a_1.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |v_1|
                                                                                                                                                                                                     {
                                                                                                                                                                                                         let matchValue_4 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&f_1);
                                                                                                                                                                                                         let matchValue_5 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&a_1);
                                                                                                                                                                                                         let matchValue_6:
                                                                                                                                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldrWithIndex(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                &&&matchValue_4),
                                                                                                                                                                                                                                                                             &&&matchValue_5),
                                                                                                                                                                                                                                          &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                             })
                                                                                                                                                                                                     }
                                                                                                                                                                                             })
                                                                                                                                                                         })
                                                                                                                                                     }),
                                                                                                                                        add(string("foldMapWithIndex"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let dictFoldableWithIndex
                                                                                                                                                                 =
                                                                                                                                                                 dictFoldableWithIndex.clone();
                                                                                                                                                             move
                                                                                                                                                                 |dictMonoid|
                                                                                                                                                                 &Func1::new({
                                                                                                                                                                                 let dictMonoid
                                                                                                                                                                                     =
                                                                                                                                                                                     dictMonoid.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |f_2|
                                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                                     let f_2
                                                                                                                                                                                                         =
                                                                                                                                                                                                         f_2.clone();
                                                                                                                                                                                                     move
                                                                                                                                                                                                         |v_2|
                                                                                                                                                                                                         {
                                                                                                                                                                                                             let matchValue_8 =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&f_2);
                                                                                                                                                                                                             let matchValue_9:
                                                                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(v_2);
                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_FoldableWithIndex::Data_FoldableWithIndex_foldMapWithIndex(),
                                                                                                                                                                                                                                                                                                                                                       &&&dictFoldableWithIndex),
                                                                                                                                                                                                                                                                                                                    &&&dictMonoid),
                                                                                                                                                                                                                                                                                 &&&matchValue_8),
                                                                                                                                                                                                                                              &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                 })
                                                                                                                                                                                                         }
                                                                                                                                                                                                 })
                                                                                                                                                                             })
                                                                                                                                                         }),
                                                                                                                                            add(string("Foldable0"),
                                                                                                                                                &&Func1::new({
                                                                                                                                                                 let foldableEnvT1
                                                                                                                                                                     =
                                                                                                                                                                     foldableEnvT1.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |usd__unused|
                                                                                                                                                                     &foldableEnvT1
                                                                                                                                                             }),
                                                                                                                                                empty::<string,
                                                                                                                                                        &dyn Any>())))))
                                                                                        }))
    }
    pub fn Control_Comonad_Env_Trans_traversableEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_traversableEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_traversableEnvT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictTraversable|
                                                                                  {
                                                                                      let functorEnvT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_functorEnvT(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      let foldableEnvT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_foldableEnvT(),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable1"),
                                                                                                                                                                     Sharpurs_Prelude::unbox(dictTraversable)),
                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                                                                                                       &&&add(string("sequence"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let dictTraversable
                                                                                                                                                   =
                                                                                                                                                   dictTraversable.clone();
                                                                                                                                               move
                                                                                                                                                   |dictApplicative|
                                                                                                                                                   {
                                                                                                                                                       let Functor0 =
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                       &Func1::new({
                                                                                                                                                                       let Functor0
                                                                                                                                                                           =
                                                                                                                                                                           Functor0.clone();
                                                                                                                                                                       let dictApplicative
                                                                                                                                                                           =
                                                                                                                                                                           dictApplicative.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |v|
                                                                                                                                                                           {
                                                                                                                                                                               let matchValue:
                                                                                                                                                                                       LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(v);
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                      &&&Functor0),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Functor::Data_Functor_functorFn()),
                                                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
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
                                                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                            }))),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                                                                                                         &&&dictTraversable),
                                                                                                                                                                                                                                                                                      &&&dictApplicative),
                                                                                                                                                                                                                                                   &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                      }))
                                                                                                                                                                           }
                                                                                                                                                                   })
                                                                                                                                                   }
                                                                                                                                           }),
                                                                                                                              add(string("traverse"),
                                                                                                                                  &&Func1::new({
                                                                                                                                                   let dictTraversable
                                                                                                                                                       =
                                                                                                                                                       dictTraversable.clone();
                                                                                                                                                   move
                                                                                                                                                       |dictApplicative_1|
                                                                                                                                                       {
                                                                                                                                                           let Functor0_1 =
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictApplicative_1)),
                                                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                           &Func1::new({
                                                                                                                                                                           let Functor0_1
                                                                                                                                                                               =
                                                                                                                                                                               Functor0_1.clone();
                                                                                                                                                                           let dictApplicative_1
                                                                                                                                                                               =
                                                                                                                                                                               dictApplicative_1.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |f|
                                                                                                                                                                               &Func1::new({
                                                                                                                                                                                               let f
                                                                                                                                                                                                   =
                                                                                                                                                                                                   f.clone();
                                                                                                                                                                                               move
                                                                                                                                                                                                   |v_1|
                                                                                                                                                                                                   {
                                                                                                                                                                                                       let matchValue_1 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                                       let matchValue_2:
                                                                                                                                                                                                               LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                              &&&Functor0_1),
                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Functor::Data_Functor_functorFn()),
                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                   |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                   Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                  let usd__arg1_1
                                                                                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                                                                                      usd__arg1_1.clone();
                                                                                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                                                                                      |usd__arg2_1|
                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                                                                                                                                              usd__arg2_1.clone()))
                                                                                                                                                                                                                                                                                                                                                                              })),
                                                                                                                                                                                                                                                                                                                                                 &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                    }))),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                                                                                                                    &&&dictTraversable),
                                                                                                                                                                                                                                                                                                                                                 &&&dictApplicative_1),
                                                                                                                                                                                                                                                                                                              &&&matchValue_1),
                                                                                                                                                                                                                                                                           &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                             x)
                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                              }))
                                                                                                                                                                                                   }
                                                                                                                                                                                           })
                                                                                                                                                                       })
                                                                                                                                                       }
                                                                                                                                               }),
                                                                                                                                  add(string("Functor0"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let functorEnvT1
                                                                                                                                                           =
                                                                                                                                                           functorEnvT1.clone();
                                                                                                                                                       move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &functorEnvT1
                                                                                                                                                   }),
                                                                                                                                      add(string("Foldable1"),
                                                                                                                                          &&Func1::new({
                                                                                                                                                           let foldableEnvT1
                                                                                                                                                               =
                                                                                                                                                               foldableEnvT1.clone();
                                                                                                                                                           move
                                                                                                                                                               |usd__unused_1|
                                                                                                                                                               &foldableEnvT1
                                                                                                                                                       }),
                                                                                                                                          empty::<string,
                                                                                                                                                  &dyn Any>())))))
                                                                                  }))
    }
    pub fn Control_Comonad_Env_Trans_traversableWithIndexEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_traversableWithIndexEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_traversableWithIndexEnvT.get_or_init(||
                                                                           &Func1::new(move
                                                                                           |dictTraversableWithIndex|
                                                                                           {
                                                                                               let functorWithIndexEnvT1 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_functorWithIndexEnvT(),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FunctorWithIndex0"),
                                                                                                                                                                              Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                               let foldableWithIndexEnvT1 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_foldableWithIndexEnvT(),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("FoldableWithIndex1"),
                                                                                                                                                                              Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                               let traversableEnvT1 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_traversableEnvT(),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Traversable2"),
                                                                                                                                                                              Sharpurs_Prelude::unbox(dictTraversableWithIndex)),
                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_TraversableWithIndexusd_Dict(),
                                                                                                                                &&&add(string("traverseWithIndex"),
                                                                                                                                       &&Func1::new({
                                                                                                                                                        let dictTraversableWithIndex
                                                                                                                                                            =
                                                                                                                                                            dictTraversableWithIndex.clone();
                                                                                                                                                        move
                                                                                                                                                            |dictApplicative|
                                                                                                                                                            {
                                                                                                                                                                let Functor0 =
                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                                                &Func1::new({
                                                                                                                                                                                let Functor0
                                                                                                                                                                                    =
                                                                                                                                                                                    Functor0.clone();
                                                                                                                                                                                let dictApplicative
                                                                                                                                                                                    =
                                                                                                                                                                                    dictApplicative.clone();
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
                                                                                                                                                                                                            let matchValue_1:
                                                                                                                                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                   &&&Functor0),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Functor::Data_Functor_functorFn()),
                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
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
                                                                                                                                                                                                                                                                                                                                                         }))),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_TraversableWithIndex::Data_TraversableWithIndex_traverseWithIndex(),
                                                                                                                                                                                                                                                                                                                                                                                         &&&dictTraversableWithIndex),
                                                                                                                                                                                                                                                                                                                                                      &&&dictApplicative),
                                                                                                                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                                                                                                                &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                   }))
                                                                                                                                                                                                        }
                                                                                                                                                                                                })
                                                                                                                                                                            })
                                                                                                                                                            }
                                                                                                                                                    }),
                                                                                                                                       add(string("FunctorWithIndex0"),
                                                                                                                                           &&Func1::new({
                                                                                                                                                            let functorWithIndexEnvT1
                                                                                                                                                                =
                                                                                                                                                                functorWithIndexEnvT1.clone();
                                                                                                                                                            move
                                                                                                                                                                |usd__unused|
                                                                                                                                                                &functorWithIndexEnvT1
                                                                                                                                                        }),
                                                                                                                                           add(string("FoldableWithIndex1"),
                                                                                                                                               &&Func1::new({
                                                                                                                                                                let foldableWithIndexEnvT1
                                                                                                                                                                    =
                                                                                                                                                                    foldableWithIndexEnvT1.clone();
                                                                                                                                                                move
                                                                                                                                                                    |usd__unused_1|
                                                                                                                                                                    &foldableWithIndexEnvT1
                                                                                                                                                            }),
                                                                                                                                               add(string("Traversable2"),
                                                                                                                                                   &&Func1::new({
                                                                                                                                                                    let traversableEnvT1
                                                                                                                                                                        =
                                                                                                                                                                        traversableEnvT1.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |usd__unused_2|
                                                                                                                                                                        &traversableEnvT1
                                                                                                                                                                }),
                                                                                                                                                   empty::<string,
                                                                                                                                                           &dyn Any>())))))
                                                                                           }))
    }
    pub fn Control_Comonad_Env_Trans_extendEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_extendEnvT: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Comonad_Env_Trans_extendEnvT.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictExtend|
                                                                             {
                                                                                 let Functor0 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                             Sharpurs_Prelude::unbox(dictExtend)),
                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined());
                                                                                 let functorEnvT1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_functorEnvT(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictExtend)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                                                  &&&add(string("extend"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let Functor0
                                                                                                                                              =
                                                                                                                                              Functor0.clone();
                                                                                                                                          let dictExtend
                                                                                                                                              =
                                                                                                                                              dictExtend.clone();
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
                                                                                                                                                                      let matchValue_1:
                                                                                                                                                                              LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                                                      let e =
                                                                                                                                                                          match matchValue_1.as_ref()
                                                                                                                                                                              {
                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                         _)
                                                                                                                                                                              =>
                                                                                                                                                                              x.clone(),
                                                                                                                                                                          };
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                          &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT()),
                                                                                                                                                                                                       &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                        &&&Functor0),
                                                                                                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                                                                                                                                                           &&&dictExtend),
                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
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
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&e)),
                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT())),
                                                                                                                                                                                                                                                                                                                                     &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                        })))))
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      }),
                                                                                                                         add(string("Functor0"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let functorEnvT1
                                                                                                                                                  =
                                                                                                                                                  functorEnvT1.clone();
                                                                                                                                              move
                                                                                                                                                  |usd__unused|
                                                                                                                                                  &functorEnvT1
                                                                                                                                          }),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>())))
                                                                             }))
    }
    pub fn Control_Comonad_Env_Trans_comonadTransEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_comonadTransEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_comonadTransEnvT.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_ComonadTransusd_Dict(),
                                                                                                    &&&add(string("lower"),
                                                                                                           &&Func1::new(move
                                                                                                                            |dictComonad|
                                                                                                                            &Func1::new(move
                                                                                                                                            |v|
                                                                                                                                            &match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                 {
                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                            x)
                                                                                                                                                 =>
                                                                                                                                                 x.clone(),
                                                                                                                                             })),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
    }
    pub fn Control_Comonad_Env_Trans_comonadEnvT() -> &dyn Any {
        static Control_Comonad_Env_Trans_comonadEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Trans_comonadEnvT.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictComonad|
                                                                              {
                                                                                  let extendEnvT1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_extendEnvT(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictComonad)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_Comonadusd_Dict(),
                                                                                                                   &&&add(string("extract"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictComonad
                                                                                                                                               =
                                                                                                                                               dictComonad.clone();
                                                                                                                                           move
                                                                                                                                               |v|
                                                                                                                                               {
                                                                                                                                                   let matchValue:
                                                                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                       &&&dictComonad),
                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                           {
                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                      x)
                                                                                                                                                                                           =>
                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                       })
                                                                                                                                               }
                                                                                                                                       }),
                                                                                                                          add(string("Extend0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let extendEnvT1
                                                                                                                                                   =
                                                                                                                                                   extendEnvT1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &extendEnvT1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
}
