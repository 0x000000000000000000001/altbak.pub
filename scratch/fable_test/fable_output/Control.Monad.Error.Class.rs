pub mod PureScript_Control_Monad_Error_Class {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_173929b2::PureScript_Data_Either;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_ceb943a5::PureScript_Effect_Exception;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Error_Class_MonadThrowusd_Dict() -> &dyn Any {
        static Control_Monad_Error_Class_MonadThrowusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_MonadThrowusd_Dict.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Control_Monad_Error_Class_MonadErrorusd_Dict() -> &dyn Any {
        static Control_Monad_Error_Class_MonadErrorusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_MonadErrorusd_Dict.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Control_Monad_Error_Class_throwError() -> &dyn Any {
        static Control_Monad_Error_Class_throwError: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Error_Class_throwError.get_or_init(||
                                                             &Func1::new(move
                                                                             |dict|
                                                                             find(string("throwError"),
                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Error_Class_monadThrowMaybe() -> &dyn Any {
        static Control_Monad_Error_Class_monadThrowMaybe:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_monadThrowMaybe.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadThrowusd_Dict(),
                                                                                                   &&&add(string("throwError"),
                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                            &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                          add(string("Monad0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Data_Maybe::Data_Maybe_monadMaybe()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Control_Monad_Error_Class_monadThrowEither() -> &dyn Any {
        static Control_Monad_Error_Class_monadThrowEither:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_monadThrowEither.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadThrowusd_Dict(),
                                                                                                    &&&add(string("throwError"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__arg1|
                                                                                                                            &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone()))),
                                                                                                           add(string("Monad0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Either::Data_Either_monadEither()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Control_Monad_Error_Class_monadThrowEffect() -> &dyn Any {
        static Control_Monad_Error_Class_monadThrowEffect:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_monadThrowEffect.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadThrowusd_Dict(),
                                                                                                    &&&add(string("throwError"),
                                                                                                           &&PureScript_Effect_Exception::Effect_Exception_throwException(),
                                                                                                           add(string("Monad0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Effect::Effect_monadEffect()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Control_Monad_Error_Class_monadErrorMaybe() -> &dyn Any {
        static Control_Monad_Error_Class_monadErrorMaybe:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_monadErrorMaybe.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadErrorusd_Dict(),
                                                                                                   &&&add(string("catchError"),
                                                                                                          &&Func1::new(move
                                                                                                                           |v|
                                                                                                                           &Func1::new({
                                                                                                                                           let v
                                                                                                                                               =
                                                                                                                                               v.clone();
                                                                                                                                           move
                                                                                                                                               |v1|
                                                                                                                                               {
                                                                                                                                                   let matchValue:
                                                                                                                                                           LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                   let matchValue_1 =
                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                   match matchValue.as_ref()
                                                                                                                                                       {
                                                                                                                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                       =>
                                                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)),
                                                                                                                                                       _
                                                                                                                                                       =>
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                        &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                   }
                                                                                                                                               }
                                                                                                                                       })),
                                                                                                          add(string("MonadThrow0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_monadThrowMaybe()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Control_Monad_Error_Class_monadErrorEither() -> &dyn Any {
        static Control_Monad_Error_Class_monadErrorEither:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_monadErrorEither.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadErrorusd_Dict(),
                                                                                                    &&&add(string("catchError"),
                                                                                                           &&Func1::new(move
                                                                                                                            |v|
                                                                                                                            &Func1::new({
                                                                                                                                            let v
                                                                                                                                                =
                                                                                                                                                v.clone();
                                                                                                                                            move
                                                                                                                                                |v1|
                                                                                                                                                {
                                                                                                                                                    let matchValue:
                                                                                                                                                            LrcPtr<Data_Either_Either> =
                                                                                                                                                        Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                    let matchValue_1 =
                                                                                                                                                        Sharpurs_Prelude::unbox(v1);
                                                                                                                                                    match matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                        =>
                                                                                                                                                        &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)),
                                                                                                                                                        Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)
                                                                                                                                                        =>
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                         &&matchValue_0_0),
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                        })),
                                                                                                           add(string("MonadThrow0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_monadThrowEither()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Control_Monad_Error_Class_monadErrorEffect() -> &dyn Any {
        static Control_Monad_Error_Class_monadErrorEffect:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_monadErrorEffect.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_MonadErrorusd_Dict(),
                                                                                                    &&&add(string("catchError"),
                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                             &&&PureScript_Effect_Exception::Effect_Exception_catchException()),
                                                                                                           add(string("MonadThrow0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_monadThrowEffect()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Control_Monad_Error_Class_liftMaybe() -> &dyn Any {
        static Control_Monad_Error_Class_liftMaybe: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Error_Class_liftMaybe.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictMonadThrow|
                                                                            {
                                                                                let pure_var =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                Sharpurs_Prelude::unbox(dictMonadThrow)),
                                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                &Func1::new({
                                                                                                let dictMonadThrow
                                                                                                    =
                                                                                                    dictMonadThrow.clone();
                                                                                                let pure_var
                                                                                                    =
                                                                                                    pure_var.clone();
                                                                                                move
                                                                                                    |error|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                                                                                                              &&&dictMonadThrow),
                                                                                                                                                                                                           error)),
                                                                                                                                     &&&pure_var)
                                                                                            })
                                                                            }))
    }
    pub fn Control_Monad_Error_Class_liftEither() -> &dyn Any {
        static Control_Monad_Error_Class_liftEither: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Error_Class_liftEither.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictMonadThrow|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                                                    dictMonadThrow)),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonadThrow)),
                                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined())))))
    }
    pub fn Control_Monad_Error_Class_catchError() -> &dyn Any {
        static Control_Monad_Error_Class_catchError: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Error_Class_catchError.get_or_init(||
                                                             &Func1::new(move
                                                                             |dict|
                                                                             find(string("catchError"),
                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Error_Class_catchJust() -> &dyn Any {
        static Control_Monad_Error_Class_catchJust: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Error_Class_catchJust.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictMonadError|
                                                                            {
                                                                                let MonadThrow0 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadThrow0"),
                                                                                                                            Sharpurs_Prelude::unbox(dictMonadError)),
                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                                &Func1::new({
                                                                                                let MonadThrow0
                                                                                                    =
                                                                                                    MonadThrow0.clone();
                                                                                                let dictMonadError
                                                                                                    =
                                                                                                    dictMonadError.clone();
                                                                                                move
                                                                                                    |p|
                                                                                                    &Func1::new({
                                                                                                                    let p
                                                                                                                        =
                                                                                                                        p.clone();
                                                                                                                    move
                                                                                                                        |act|
                                                                                                                        &Func1::new({
                                                                                                                                        let act
                                                                                                                                            =
                                                                                                                                            act.clone();
                                                                                                                                        move
                                                                                                                                            |handler|
                                                                                                                                            {
                                                                                                                                                let handle =
                                                                                                                                                    &Func1::new({
                                                                                                                                                                    let handler
                                                                                                                                                                        =
                                                                                                                                                                        handler.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |e|
                                                                                                                                                                        {
                                                                                                                                                                            let matchValue:
                                                                                                                                                                                    LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                                                           e));
                                                                                                                                                                            match matchValue.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                =>
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&handler,
                                                                                                                                                                                                                 &&matchValue_1_0),
                                                                                                                                                                                _
                                                                                                                                                                                =>
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                                                                                                                                                    &&&MonadThrow0),
                                                                                                                                                                                                                 e),
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                });
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_catchError(),
                                                                                                                                                                                                                                                       &&&dictMonadError),
                                                                                                                                                                                                                    &&&act),
                                                                                                                                                                                 &&&handle)
                                                                                                                                            }
                                                                                                                                    })
                                                                                                                })
                                                                                            })
                                                                            }))
    }
    pub fn Control_Monad_Error_Class_try() -> &dyn Any {
        static Control_Monad_Error_Class_try: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_try.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonadError|
                                                                      {
                                                                          let Monad0 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadThrow0"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonadError)),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                          let Functor0 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                          let pure_var =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                         Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()));
                                                                          &Func1::new({
                                                                                          let Functor0
                                                                                              =
                                                                                              Functor0.clone();
                                                                                          let dictMonadError
                                                                                              =
                                                                                              dictMonadError.clone();
                                                                                          let pure_var
                                                                                              =
                                                                                              pure_var.clone();
                                                                                          move
                                                                                              |a|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_catchError(),
                                                                                                                                                                                                     &&&dictMonadError),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                           &&&Functor0),
                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                          |usd__arg1|
                                                                                                                                                                                                                                                          &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                     a)),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                     &&&pure_var),
                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                    |usd__arg1_1|
                                                                                                                                                                                    &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_1.clone())))))
                                                                                      })
                                                                      }))
    }
    pub fn Control_Monad_Error_Class_withResource() -> &dyn Any {
        static Control_Monad_Error_Class_withResource:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Error_Class_withResource.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictMonadError|
                                                                               {
                                                                                   let MonadThrow0 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadThrow0"),
                                                                                                                               Sharpurs_Prelude::unbox(dictMonadError)),
                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined());
                                                                                   let Monad0 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                               Sharpurs_Prelude::unbox(&&MonadThrow0)),
                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined());
                                                                                   let Bind1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                               Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined());
                                                                                   let try1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_try(),
                                                                                                                        dictMonadError);
                                                                                   let throwError1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Error_Class::Control_Monad_Error_Class_throwError(),
                                                                                                                        &&&MonadThrow0);
                                                                                   let pure_var =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Monad0)),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                   &Func1::new({
                                                                                                   let Bind1
                                                                                                       =
                                                                                                       Bind1.clone();
                                                                                                   let pure_var
                                                                                                       =
                                                                                                       pure_var.clone();
                                                                                                   let throwError1
                                                                                                       =
                                                                                                       throwError1.clone();
                                                                                                   let try1
                                                                                                       =
                                                                                                       try1.clone();
                                                                                                   move
                                                                                                       |acquire|
                                                                                                       &Func1::new({
                                                                                                                       let acquire
                                                                                                                           =
                                                                                                                           acquire.clone();
                                                                                                                       move
                                                                                                                           |release|
                                                                                                                           &Func1::new({
                                                                                                                                           let release
                                                                                                                                               =
                                                                                                                                               release.clone();
                                                                                                                                           move
                                                                                                                                               |kleisli|
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                      &&&Bind1),
                                                                                                                                                                                                                   &&&acquire),
                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                  let kleisli
                                                                                                                                                                                                      =
                                                                                                                                                                                                      kleisli.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |resource|
                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                             &&&Bind1),
                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                &&&try1),
                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&kleisli,
                                                                                                                                                                                                                                                                                                                                                resource))),
                                                                                                                                                                                                                                       &&&Func1::new({
                                                                                                                                                                                                                                                         let resource
                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                             resource.clone();
                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                             |result|
                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_discard(),
                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Control_Bind::Control_Bind_discardUnit()),
                                                                                                                                                                                                                                                                                                                                                                    &&&Bind1),
                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&release,
                                                                                                                                                                                                                                                                                                                                                                    &&&resource)),
                                                                                                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                                                                                                let result
                                                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                                                    result.clone();
                                                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                                                    |usd__unused|
                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Either::Data_Either_either(),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&throwError1),
                                                                                                                                                                                                                                                                                                                                                                                        &&&pure_var),
                                                                                                                                                                                                                                                                                                                                                     &&&result)
                                                                                                                                                                                                                                                                                                            }))
                                                                                                                                                                                                                                                     }))
                                                                                                                                                                                              }))
                                                                                                                                       })
                                                                                                                   })
                                                                                               })
                                                                               }))
    }
}
