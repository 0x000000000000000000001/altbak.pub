pub mod PureScript_Control_Comonad_Env_Class {
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
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Env_Class_lower() -> &dyn Any {
        static Control_Comonad_Env_Class_lower: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_lower.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_lower(),
                                                                                         &&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_comonadTransStoreT()))
    }
    pub fn Control_Comonad_Env_Class_ComonadAskusd_Dict() -> &dyn Any {
        static Control_Comonad_Env_Class_ComonadAskusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_ComonadAskusd_Dict.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Control_Comonad_Env_Class_ComonadEnvusd_Dict() -> &dyn Any {
        static Control_Comonad_Env_Class_ComonadEnvusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_ComonadEnvusd_Dict.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Control_Comonad_Env_Class_local() -> &dyn Any {
        static Control_Comonad_Env_Class_local: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_local.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("local"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Comonad_Env_Class_comonadAskTuple() -> &dyn Any {
        static Control_Comonad_Env_Class_comonadAskTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_comonadAskTuple.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ComonadAskusd_Dict(),
                                                                                                   &&&add(string("ask"),
                                                                                                          &&PureScript_Data_Tuple::Data_Tuple_fst(),
                                                                                                          add(string("Comonad0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Data_Tuple::Data_Tuple_comonadTuple()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Control_Comonad_Env_Class_comonadEnvTuple() -> &dyn Any {
        static Control_Comonad_Env_Class_comonadEnvTuple:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_comonadEnvTuple.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ComonadEnvusd_Dict(),
                                                                                                   &&&add(string("local"),
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
                                                                                                                                                   let matchValue_1:
                                                                                                                                                           LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                   &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
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
                                                                                                                                                                                                            }))
                                                                                                                                               }
                                                                                                                                       })),
                                                                                                          add(string("ComonadAsk0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_comonadAskTuple()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Control_Comonad_Env_Class_comonadAskEnvT() -> &dyn Any {
        static Control_Comonad_Env_Class_comonadAskEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_comonadAskEnvT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictComonad|
                                                                                 {
                                                                                     let comonadEnvT =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_comonadEnvT(),
                                                                                                                          dictComonad);
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ComonadAskusd_Dict(),
                                                                                                                      &&&add(string("ask"),
                                                                                                                             &&Func1::new(move
                                                                                                                                              |v|
                                                                                                                                              {
                                                                                                                                                  let x =
                                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_fst(),
                                                                                                                                                                                   &&&x)
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
                                                                                                                                         &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Comonad_Env_Class_comonadEnvEnvT() -> &dyn Any {
        static Control_Comonad_Env_Class_comonadEnvEnvT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_comonadEnvEnvT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictComonad|
                                                                                 {
                                                                                     let comonadAskEnvT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_comonadAskEnvT(),
                                                                                                                          dictComonad);
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ComonadEnvusd_Dict(),
                                                                                                                      &&&add(string("local"),
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
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Trans::Control_Comonad_Env_Trans_EnvT(),
                                                                                                                                                                                                       &&{
                                                                                                                                                                                                             let matchValue_3:
                                                                                                                                                                                                                     LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&matchValue_1);
                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                      &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                     &match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                      }))
                                                                                                                                                                                                         })
                                                                                                                                                                  }
                                                                                                                                                          })),
                                                                                                                             add(string("ComonadAsk0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let comonadAskEnvT1
                                                                                                                                                      =
                                                                                                                                                      comonadAskEnvT1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &comonadAskEnvT1
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Comonad_Env_Class_ask() -> &dyn Any {
        static Control_Comonad_Env_Class_ask: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_ask.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("ask"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Comonad_Env_Class_asks() -> &dyn Any {
        static Control_Comonad_Env_Class_asks: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_asks.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictComonadAsk|
                                                                       &Func1::new({
                                                                                       let dictComonadAsk
                                                                                           =
                                                                                           dictComonadAsk.clone();
                                                                                       move
                                                                                           |f|
                                                                                           &Func1::new({
                                                                                                           let f
                                                                                                               =
                                                                                                               f.clone();
                                                                                                           move
                                                                                                               |x|
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ask(),
                                                                                                                                                                                                                      &&&dictComonadAsk),
                                                                                                                                                                                   x))
                                                                                                       })
                                                                                   })))
    }
    pub fn Control_Comonad_Env_Class_comonadAskStoreT() -> &dyn Any {
        static Control_Comonad_Env_Class_comonadAskStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_comonadAskStoreT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictComonadAsk|
                                                                                   {
                                                                                       let Comonad0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                   Sharpurs_Prelude::unbox(dictComonadAsk)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let comonadStoreT =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_comonadStoreT(),
                                                                                                                            &&&Comonad0);
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ComonadAskusd_Dict(),
                                                                                                                        &&&add(string("ask"),
                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ask(),
                                                                                                                                                                                                                                       dictComonadAsk)),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_lower(),
                                                                                                                                                                                                    &&&Comonad0)),
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
    pub fn Control_Comonad_Env_Class_comonadEnvStoreT() -> &dyn Any {
        static Control_Comonad_Env_Class_comonadEnvStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_comonadEnvStoreT.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictComonadEnv|
                                                                                   {
                                                                                       let comonadAskStoreT1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_comonadAskStoreT(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("ComonadAsk0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(dictComonadEnv)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ComonadEnvusd_Dict(),
                                                                                                                        &&&add(string("local"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let dictComonadEnv
                                                                                                                                                    =
                                                                                                                                                    dictComonadEnv.clone();
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
                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_StoreT(),
                                                                                                                                                                                                             &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_local(),
                                                                                                                                                                                                                                                                                                                                                                              &&&dictComonadEnv),
                                                                                                                                                                                                                                                                                                                                           &&&matchValue),
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
                                                                                                                                                                })
                                                                                                                                            }),
                                                                                                                               add(string("ComonadAsk0"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let comonadAskStoreT1
                                                                                                                                                        =
                                                                                                                                                        comonadAskStoreT1.clone();
                                                                                                                                                    move
                                                                                                                                                        |usd__unused|
                                                                                                                                                        &comonadAskStoreT1
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>())))
                                                                                   }))
    }
    pub fn Control_Comonad_Env_Class_comonadAskTracedT() -> &dyn Any {
        static Control_Comonad_Env_Class_comonadAskTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_comonadAskTracedT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictComonadAsk|
                                                                                    {
                                                                                        let ask1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ask(),
                                                                                                                             dictComonadAsk);
                                                                                        let Comonad0 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Comonad0"),
                                                                                                                                    Sharpurs_Prelude::unbox(dictComonadAsk)),
                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                                        let comonadTracedT =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_comonadTracedT(),
                                                                                                                             &&&Comonad0);
                                                                                        &Func1::new({
                                                                                                        let Comonad0
                                                                                                            =
                                                                                                            Comonad0.clone();
                                                                                                        let ask1
                                                                                                            =
                                                                                                            ask1.clone();
                                                                                                        let comonadTracedT
                                                                                                            =
                                                                                                            comonadTracedT.clone();
                                                                                                        move
                                                                                                            |dictMonoid|
                                                                                                            {
                                                                                                                let comonadTracedT1 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&comonadTracedT,
                                                                                                                                                     dictMonoid);
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ComonadAskusd_Dict(),
                                                                                                                                                 &&&add(string("ask"),
                                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                             &&&ask1),
                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_lower(),
                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_comonadTransTracedT(),
                                                                                                                                                                                                                                                                                                   dictMonoid)),
                                                                                                                                                                                                                             &&&Comonad0)),
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
    pub fn Control_Comonad_Env_Class_comonadEnvTracedT() -> &dyn Any {
        static Control_Comonad_Env_Class_comonadEnvTracedT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Env_Class_comonadEnvTracedT.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictComonadEnv|
                                                                                    {
                                                                                        let comonadAskTracedT1 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_comonadAskTracedT(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("ComonadAsk0"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(dictComonadEnv)),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        &Func1::new({
                                                                                                        let comonadAskTracedT1
                                                                                                            =
                                                                                                            comonadAskTracedT1.clone();
                                                                                                        let dictComonadEnv
                                                                                                            =
                                                                                                            dictComonadEnv.clone();
                                                                                                        move
                                                                                                            |dictMonoid|
                                                                                                            {
                                                                                                                let comonadAskTracedT2 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&comonadAskTracedT1,
                                                                                                                                                     dictMonoid);
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_ComonadEnvusd_Dict(),
                                                                                                                                                 &&&add(string("local"),
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
                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_TracedT(),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Env_Class::Control_Comonad_Env_Class_local(),
                                                                                                                                                                                                                                                                                                                                           &&&dictComonadEnv),
                                                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                                                     &&&matchValue_1))
                                                                                                                                                                                             }
                                                                                                                                                                                     })),
                                                                                                                                                        add(string("ComonadAsk0"),
                                                                                                                                                            &&Func1::new({
                                                                                                                                                                             let comonadAskTracedT2
                                                                                                                                                                                 =
                                                                                                                                                                                 comonadAskTracedT2.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |usd__unused|
                                                                                                                                                                                 &comonadAskTracedT2
                                                                                                                                                                         }),
                                                                                                                                                            empty::<string,
                                                                                                                                                                    &dyn Any>())))
                                                                                                            }
                                                                                                    })
                                                                                    }))
    }
}
