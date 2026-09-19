pub mod PureScript_Control_Comonad_Store_Class {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ca268cc0::PureScript_Control_Comonad_Env_Trans;
    use crate::module_a7b6ae02::PureScript_Control_Comonad_Store_Trans;
    use crate::module_7d1f6438::PureScript_Control_Comonad_Traced_Trans;
    use crate::module_35294a53::PureScript_Control_Comonad_Trans_Class;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Store_Class_lower() -> &dyn Any {
        static Control_Comonad_Store_Class_lower: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_lower.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_lower(),
                                                                                           &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_comonadTransEnvT()))
    }
    pub fn Control_Comonad_Store_Class_ComonadStoreusd_Dict() -> &dyn Any {
        static Control_Comonad_Store_Class_ComonadStoreusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_ComonadStoreusd_Dict.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |x|
                                                                                         x.clone()))
    }
    pub fn Control_Comonad_Store_Class_pos() -> &dyn Any {
        static Control_Comonad_Store_Class_pos: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_pos.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("pos"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Comonad_Store_Class_peek() -> &dyn Any {
        static Control_Comonad_Store_Class_peek: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_peek.get_or_init(||
                                                         &Func1::new(move
                                                                         |dict|
                                                                         find(string("peek"),
                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Comonad_Store_Class_peeks() -> &dyn Any {
        static Control_Comonad_Store_Class_peeks: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_peeks.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictComonadStore|
                                                                          &Func1::new({
                                                                                          let dictComonadStore
                                                                                              =
                                                                                              dictComonadStore.clone();
                                                                                          move
                                                                                              |f|
                                                                                              &Func1::new({
                                                                                                              let f
                                                                                                                  =
                                                                                                                  f.clone();
                                                                                                              move
                                                                                                                  |x|
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_peek(),
                                                                                                                                                                                                                         &&&dictComonadStore),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                            &&&f),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_pos(),
                                                                                                                                                                                                                                                                                               &&&dictComonadStore),
                                                                                                                                                                                                                                                            x))),
                                                                                                                                                   x)
                                                                                                          })
                                                                                      })))
    }
    pub fn Control_Comonad_Store_Class_seeks() -> &dyn Any {
        static Control_Comonad_Store_Class_seeks: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_seeks.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictComonadStore|
                                                                          {
                                                                              let duplicate =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_duplicate(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                                                                                                              Sharpurs_Prelude::unbox(dictComonadStore)),
                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              &Func1::new({
                                                                                              let dictComonadStore
                                                                                                  =
                                                                                                  dictComonadStore.clone();
                                                                                              let duplicate
                                                                                                  =
                                                                                                  duplicate.clone();
                                                                                              move
                                                                                                  |f|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_peeks(),
                                                                                                                                                                                                                                            &&&dictComonadStore),
                                                                                                                                                                                                         f)),
                                                                                                                                   &&&duplicate)
                                                                                          })
                                                                          }))
    }
    pub fn Control_Comonad_Store_Class_seek() -> &dyn Any {
        static Control_Comonad_Store_Class_seek: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_seek.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictComonadStore|
                                                                         {
                                                                             let duplicate =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_duplicate(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                                                                                                             Sharpurs_Prelude::unbox(dictComonadStore)),
                                                                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             &Func1::new({
                                                                                             let dictComonadStore
                                                                                                 =
                                                                                                 dictComonadStore.clone();
                                                                                             let duplicate
                                                                                                 =
                                                                                                 duplicate.clone();
                                                                                             move
                                                                                                 |s|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_peek(),
                                                                                                                                                                                                                                           &&&dictComonadStore),
                                                                                                                                                                                                        s)),
                                                                                                                                  &&&duplicate)
                                                                                         })
                                                                         }))
    }
    pub fn Control_Comonad_Store_Class_experiment() -> &dyn Any {
        static Control_Comonad_Store_Class_experiment:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_experiment.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictComonadStore|
                                                                               {
                                                                                   let peek1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_peek(),
                                                                                                                        dictComonadStore);
                                                                                   &Func1::new({
                                                                                                   let dictComonadStore
                                                                                                       =
                                                                                                       dictComonadStore.clone();
                                                                                                   let peek1
                                                                                                       =
                                                                                                       peek1.clone();
                                                                                                   move
                                                                                                       |dictFunctor|
                                                                                                       &Func1::new({
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
                                                                                                                                               |x|
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                      &&&dictFunctor),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                         &&&peek1),
                                                                                                                                                                                                                                                      x)),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_pos(),
                                                                                                                                                                                                                                                                                         &&&dictComonadStore),
                                                                                                                                                                                                                                                      x)))
                                                                                                                                       })
                                                                                                                   })
                                                                                               })
                                                                               }))
    }
    pub fn Control_Comonad_Store_Class_comonadStoreTracedT() -> &dyn Any {
        static Control_Comonad_Store_Class_comonadStoreTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_comonadStoreTracedT.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictComonadStore|
                                                                                        {
                                                                                            let pos1 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_pos(),
                                                                                                                                 dictComonadStore);
                                                                                            let Comonad0 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                        Sharpurs_Prelude::unbox(dictComonadStore)),
                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                                            let comonadTracedT =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_comonadTracedT(),
                                                                                                                                 &&&Comonad0);
                                                                                            &Func1::new({
                                                                                                            let Comonad0
                                                                                                                =
                                                                                                                Comonad0.clone();
                                                                                                            let comonadTracedT
                                                                                                                =
                                                                                                                comonadTracedT.clone();
                                                                                                            let dictComonadStore
                                                                                                                =
                                                                                                                dictComonadStore.clone();
                                                                                                            let pos1
                                                                                                                =
                                                                                                                pos1.clone();
                                                                                                            move
                                                                                                                |dictMonoid|
                                                                                                                {
                                                                                                                    let lower1 =
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_lower(),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_comonadTransTracedT(),
                                                                                                                                                                                                                               dictMonoid)),
                                                                                                                                                         &&&Comonad0);
                                                                                                                    let comonadTracedT1 =
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&comonadTracedT,
                                                                                                                                                         dictMonoid);
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_ComonadStoreusd_Dict(),
                                                                                                                                                     &&&add(string("pos"),
                                                                                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                 &&&pos1),
                                                                                                                                                                                              &&&lower1),
                                                                                                                                                            add(string("peek"),
                                                                                                                                                                &&Func1::new({
                                                                                                                                                                                 let lower1
                                                                                                                                                                                     =
                                                                                                                                                                                     lower1.clone();
                                                                                                                                                                                 move
                                                                                                                                                                                     |s|
                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_peek(),
                                                                                                                                                                                                                                                                                                                               &&&dictComonadStore),
                                                                                                                                                                                                                                                                                            s)),
                                                                                                                                                                                                                      &&&lower1)
                                                                                                                                                                             }),
                                                                                                                                                                add(string("Comonad0"),
                                                                                                                                                                    &&Func1::new({
                                                                                                                                                                                     let comonadTracedT1
                                                                                                                                                                                         =
                                                                                                                                                                                         comonadTracedT1.clone();
                                                                                                                                                                                     move
                                                                                                                                                                                         |usd__unused|
                                                                                                                                                                                         &comonadTracedT1
                                                                                                                                                                                 }),
                                                                                                                                                                    empty::<string,
                                                                                                                                                                            &dyn Any>()))))
                                                                                                                }
                                                                                                        })
                                                                                        }))
    }
    pub fn Control_Comonad_Store_Class_comonadStoreStoreT() -> &dyn Any {
        static Control_Comonad_Store_Class_comonadStoreStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_comonadStoreStoreT.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictComonad|
                                                                                       {
                                                                                           let comonadStoreT =
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_comonadStoreT(),
                                                                                                                                dictComonad);
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_ComonadStoreusd_Dict(),
                                                                                                                            &&&add(string("pos"),
                                                                                                                                   &&Func1::new(move
                                                                                                                                                    |v|
                                                                                                                                                    &match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                         {
                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                    x)
                                                                                                                                                         =>
                                                                                                                                                         x.clone(),
                                                                                                                                                     }),
                                                                                                                                   add(string("peek"),
                                                                                                                                       &&Func1::new({
                                                                                                                                                        let dictComonad
                                                                                                                                                            =
                                                                                                                                                            dictComonad.clone();
                                                                                                                                                        move
                                                                                                                                                            |s_1|
                                                                                                                                                            &Func1::new({
                                                                                                                                                                            let s_1
                                                                                                                                                                                =
                                                                                                                                                                                s_1.clone();
                                                                                                                                                                            move
                                                                                                                                                                                |v_1|
                                                                                                                                                                                {
                                                                                                                                                                                    let matchValue_1 =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&s_1);
                                                                                                                                                                                    let matchValue_2:
                                                                                                                                                                                            LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                        Sharpurs_Prelude::unbox(v_1);
                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                                                                                           &&&dictComonad),
                                                                                                                                                                                                                                                        &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                     &&&matchValue_1)
                                                                                                                                                                                }
                                                                                                                                                                        })
                                                                                                                                                    }),
                                                                                                                                       add(string("Comonad0"),
                                                                                                                                           &&Func1::new({
                                                                                                                                                            let comonadStoreT
                                                                                                                                                                =
                                                                                                                                                                comonadStoreT.clone();
                                                                                                                                                            move
                                                                                                                                                                |usd__unused|
                                                                                                                                                                &comonadStoreT
                                                                                                                                                        }),
                                                                                                                                           empty::<string,
                                                                                                                                                   &dyn Any>()))))
                                                                                       }))
    }
    pub fn Control_Comonad_Store_Class_comonadStoreEnvT() -> &dyn Any {
        static Control_Comonad_Store_Class_comonadStoreEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Class_comonadStoreEnvT.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictComonadStore|
                                                                                     {
                                                                                         let Comonad0 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                     Sharpurs_Prelude::unbox(dictComonadStore)),
                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined());
                                                                                         let lower1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_lower(),
                                                                                                                              &&&Comonad0);
                                                                                         let comonadEnvT =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_comonadEnvT(),
                                                                                                                              &&&Comonad0);
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_ComonadStoreusd_Dict(),
                                                                                                                          &&&add(string("pos"),
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_pos(),
                                                                                                                                                                                                                                         dictComonadStore)),
                                                                                                                                                                   &&&lower1),
                                                                                                                                 add(string("peek"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let dictComonadStore
                                                                                                                                                          =
                                                                                                                                                          dictComonadStore.clone();
                                                                                                                                                      let lower1
                                                                                                                                                          =
                                                                                                                                                          lower1.clone();
                                                                                                                                                      move
                                                                                                                                                          |s|
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Class::Control_Comonad_Store_Class_peek(),
                                                                                                                                                                                                                                                                                                    &&&dictComonadStore),
                                                                                                                                                                                                                                                                 s)),
                                                                                                                                                                                           &&&lower1)
                                                                                                                                                  }),
                                                                                                                                     add(string("Comonad0"),
                                                                                                                                         &&Func1::new({
                                                                                                                                                          let comonadEnvT
                                                                                                                                                              =
                                                                                                                                                              comonadEnvT.clone();
                                                                                                                                                          move
                                                                                                                                                              |usd__unused|
                                                                                                                                                              &comonadEnvT
                                                                                                                                                      }),
                                                                                                                                         empty::<string,
                                                                                                                                                 &dyn Any>()))))
                                                                                     }))
    }
}
