pub mod PureScript_Control_Comonad_Traced_Class {
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
    use crate::module_bc5f9127::PureScript_Control_Monad_Identity_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Traced_Class_ComonadTracedusd_Dict() -> &dyn Any {
        static Control_Comonad_Traced_Class_ComonadTracedusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_ComonadTracedusd_Dict.get_or_init(||
                                                                           &Func1::new(move
                                                                                           |x|
                                                                                           x.clone()))
    }
    pub fn Control_Comonad_Traced_Class_track() -> &dyn Any {
        static Control_Comonad_Traced_Class_track: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_track.get_or_init(||
                                                           &Func1::new(move
                                                                           |dict|
                                                                           find(string("track"),
                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Comonad_Traced_Class_tracks() -> &dyn Any {
        static Control_Comonad_Traced_Class_tracks: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Comonad_Traced_Class_tracks.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictComonadTraced|
                                                                            {
                                                                                let Comonad0 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                            Sharpurs_Prelude::unbox(dictComonadTraced)),
                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                                &Func1::new({
                                                                                                let Comonad0
                                                                                                    =
                                                                                                    Comonad0.clone();
                                                                                                let dictComonadTraced
                                                                                                    =
                                                                                                    dictComonadTraced.clone();
                                                                                                move
                                                                                                    |f|
                                                                                                    &Func1::new({
                                                                                                                    let f
                                                                                                                        =
                                                                                                                        f.clone();
                                                                                                                    move
                                                                                                                        |w|
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_track(),
                                                                                                                                                                                                                               &&&dictComonadTraced),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                  &&&f),
                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                                                                                                     &&&Comonad0),
                                                                                                                                                                                                                                                                  w))),
                                                                                                                                                         w)
                                                                                                                })
                                                                                            })
                                                                            }))
    }
    pub fn Control_Comonad_Traced_Class_lowerTrack() -> &dyn Any {
        static Control_Comonad_Traced_Class_lowerTrack:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_lowerTrack.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictComonadTrans|
                                                                                {
                                                                                    let lower =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_lower(),
                                                                                                                         dictComonadTrans);
                                                                                    &Func1::new({
                                                                                                    let lower
                                                                                                        =
                                                                                                        lower.clone();
                                                                                                    move
                                                                                                        |dictComonadTraced|
                                                                                                        {
                                                                                                            let lower1 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&lower,
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictComonadTraced)),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                            &Func1::new({
                                                                                                                            let dictComonadTraced
                                                                                                                                =
                                                                                                                                dictComonadTraced.clone();
                                                                                                                            let lower1
                                                                                                                                =
                                                                                                                                lower1.clone();
                                                                                                                            move
                                                                                                                                |m|
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_track(),
                                                                                                                                                                                                                                                                          &&&dictComonadTraced),
                                                                                                                                                                                                                                       m)),
                                                                                                                                                                 &&&lower1)
                                                                                                                        })
                                                                                                        }
                                                                                                })
                                                                                }))
    }
    pub fn Control_Comonad_Traced_Class_lowerTrack1() -> &dyn Any {
        static Control_Comonad_Traced_Class_lowerTrack1:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_lowerTrack1.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_lowerTrack(),
                                                                                                  &&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_comonadTransStoreT()))
    }
    pub fn Control_Comonad_Traced_Class_lowerTrack2() -> &dyn Any {
        static Control_Comonad_Traced_Class_lowerTrack2:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_lowerTrack2.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_lowerTrack(),
                                                                                                  &&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_comonadTransIdentityT()))
    }
    pub fn Control_Comonad_Traced_Class_lowerTrack3() -> &dyn Any {
        static Control_Comonad_Traced_Class_lowerTrack3:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_lowerTrack3.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_lowerTrack(),
                                                                                                  &&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_comonadTransEnvT()))
    }
    pub fn Control_Comonad_Traced_Class_listens() -> &dyn Any {
        static Control_Comonad_Traced_Class_listens: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Comonad_Traced_Class_listens.get_or_init(||
                                                             &Func1::new(move
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
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let matchValue =
                                                                                                                             Sharpurs_Prelude::unbox(&&f);
                                                                                                                         let matchValue_1 =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_TracedT(),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                   &&&dictFunctor),
                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                  |g|
                                                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                                                  let g
                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                      g.clone();
                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                      |t|
                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                                                                                                                                               t),
                                                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                               t)))
                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                             &&&matchValue_1))
                                                                                                                     }
                                                                                                             })
                                                                                         })))
    }
    pub fn Control_Comonad_Traced_Class_listen() -> &dyn Any {
        static Control_Comonad_Traced_Class_listen: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Comonad_Traced_Class_listen.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFunctor|
                                                                            &Func1::new({
                                                                                            let dictFunctor
                                                                                                =
                                                                                                dictFunctor.clone();
                                                                                            move
                                                                                                |v|
                                                                                                {
                                                                                                    let tr =
                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_TracedT(),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                              &&&dictFunctor),
                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                             |f|
                                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                                             let f
                                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                                 f.clone();
                                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                                 |t|
                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                          t),
                                                                                                                                                                                                                                                                                                         t.clone()))
                                                                                                                                                                                                                                         }))),
                                                                                                                                                                        &&&tr))
                                                                                                }
                                                                                        })))
    }
    pub fn Control_Comonad_Traced_Class_comonadTracedTracedT() -> &dyn Any {
        static Control_Comonad_Traced_Class_comonadTracedTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_comonadTracedTracedT.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |dictComonad|
                                                                                          {
                                                                                              let comonadTracedT =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_comonadTracedT(),
                                                                                                                                   dictComonad);
                                                                                              &Func1::new({
                                                                                                              let comonadTracedT
                                                                                                                  =
                                                                                                                  comonadTracedT.clone();
                                                                                                              let dictComonad
                                                                                                                  =
                                                                                                                  dictComonad.clone();
                                                                                                              move
                                                                                                                  |dictMonoid|
                                                                                                                  {
                                                                                                                      let comonadTracedT1 =
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&comonadTracedT,
                                                                                                                                                           dictMonoid);
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_ComonadTracedusd_Dict(),
                                                                                                                                                       &&&add(string("track"),
                                                                                                                                                              &&Func1::new(move
                                                                                                                                                                               |t|
                                                                                                                                                                               &Func1::new({
                                                                                                                                                                                               let t
                                                                                                                                                                                                   =
                                                                                                                                                                                                   t.clone();
                                                                                                                                                                                               move
                                                                                                                                                                                                   |v|
                                                                                                                                                                                                   {
                                                                                                                                                                                                       let matchValue =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&t);
                                                                                                                                                                                                       let matchValue_1 =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                                                                                                              &&&dictComonad),
                                                                                                                                                                                                                                                                           &&&matchValue_1),
                                                                                                                                                                                                                                        &&&matchValue)
                                                                                                                                                                                                   }
                                                                                                                                                                                           })),
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
                                                                                                                                                                          &dyn Any>())))
                                                                                                                  }
                                                                                                          })
                                                                                          }))
    }
    pub fn Control_Comonad_Traced_Class_comonadTracedStoreT() -> &dyn Any {
        static Control_Comonad_Traced_Class_comonadTracedStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_comonadTracedStoreT.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |dictComonadTraced|
                                                                                         {
                                                                                             let comonadStoreT =
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_comonadStoreT(),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                                                            Sharpurs_Prelude::unbox(dictComonadTraced)),
                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_ComonadTracedusd_Dict(),
                                                                                                                              &&&add(string("track"),
                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_lowerTrack1(),
                                                                                                                                                                       dictComonadTraced),
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
                                                                                                                                                 &dyn Any>())))
                                                                                         }))
    }
    pub fn Control_Comonad_Traced_Class_comonadTracedIdentityT() -> &dyn Any {
        static Control_Comonad_Traced_Class_comonadTracedIdentityT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_comonadTracedIdentityT.get_or_init(||
                                                                            &Func1::new(move
                                                                                            |dictComonadTraced|
                                                                                            {
                                                                                                let comonadIdentityT =
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Identity_Trans::Control_Monad_Identity_Trans_comonadIdentityT(),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                                                               Sharpurs_Prelude::unbox(dictComonadTraced)),
                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_ComonadTracedusd_Dict(),
                                                                                                                                 &&&add(string("track"),
                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_lowerTrack2(),
                                                                                                                                                                          dictComonadTraced),
                                                                                                                                        add(string("Comonad0"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let comonadIdentityT
                                                                                                                                                                 =
                                                                                                                                                                 comonadIdentityT.clone();
                                                                                                                                                             move
                                                                                                                                                                 |usd__unused|
                                                                                                                                                                 &comonadIdentityT
                                                                                                                                                         }),
                                                                                                                                            empty::<string,
                                                                                                                                                    &dyn Any>())))
                                                                                            }))
    }
    pub fn Control_Comonad_Traced_Class_comonadTracedEnvT() -> &dyn Any {
        static Control_Comonad_Traced_Class_comonadTracedEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_Class_comonadTracedEnvT.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictComonadTraced|
                                                                                       {
                                                                                           let comonadEnvT =
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_comonadEnvT(),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                                                          Sharpurs_Prelude::unbox(dictComonadTraced)),
                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_ComonadTracedusd_Dict(),
                                                                                                                            &&&add(string("track"),
                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Class::Control_Comonad_Traced_Class_lowerTrack3(),
                                                                                                                                                                     dictComonadTraced),
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
                                                                                                                                               &dyn Any>())))
                                                                                       }))
    }
    pub fn Control_Comonad_Traced_Class_censor() -> &dyn Any {
        static Control_Comonad_Traced_Class_censor: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Comonad_Traced_Class_censor.get_or_init(||
                                                            &Func1::new(move
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
                                                                                                                    |v|
                                                                                                                    {
                                                                                                                        let matchValue =
                                                                                                                            Sharpurs_Prelude::unbox(&&f);
                                                                                                                        let matchValue_1 =
                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_TracedT(),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                  &&&dictFunctor),
                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                 |v1|
                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                                                                                                                                                  v1))),
                                                                                                                                                                                            &&&matchValue_1))
                                                                                                                    }
                                                                                                            })
                                                                                        })))
    }
}
