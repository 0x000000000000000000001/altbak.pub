pub mod PureScript_Control_Comonad_Store_Trans {
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
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Store_Trans_StoreT() -> &dyn Any {
        static Control_Comonad_Store_Trans_StoreT: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Trans_StoreT.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Control_Comonad_Store_Trans_runStoreT() -> &dyn Any {
        static Control_Comonad_Store_Trans_runStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Trans_runStoreT.get_or_init(||
                                                              &Func1::new(move
                                                                              |v|
                                                                              &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Control_Comonad_Store_Trans_newtypeStoreT() -> &dyn Any {
        static Control_Comonad_Store_Trans_newtypeStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Trans_newtypeStoreT.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                   &&&add(string("Coercible0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &Sharpurs_Prelude::Prim_undefined()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
    }
    pub fn Control_Comonad_Store_Trans_functorStoreT() -> &dyn Any {
        static Control_Comonad_Store_Trans_functorStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Trans_functorStoreT.get_or_init(||
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
                                                                                                                                                                                                                                           &&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_StoreT()),
                                                                                                                                                                                                        &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                         &&&dictFunctor),
                                                                                                                                                                                                                                                                                                                                      &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                        |h|
                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                            h),
                                                                                                                                                                                                                                                                                                                                                                                         &&&matchValue))),
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
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>()))))
    }
    pub fn Control_Comonad_Store_Trans_extendStoreT() -> &dyn Any {
        static Control_Comonad_Store_Trans_extendStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Trans_extendStoreT.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictExtend|
                                                                                 {
                                                                                     let functorStoreT1 =
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_functorStoreT(),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                    Sharpurs_Prelude::unbox(dictExtend)),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                                                      &&&add(string("extend"),
                                                                                                                             &&Func1::new({
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
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                              &&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_StoreT()),
                                                                                                                                                                                                           &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                                                                                                                            &&&dictExtend),
                                                                                                                                                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                           |w_prime|
                                                                                                                                                                                                                                                                                                                                                           &Func1::new({
                                                                                                                                                                                                                                                                                                                                                                           let w_prime
                                                                                                                                                                                                                                                                                                                                                                               =
                                                                                                                                                                                                                                                                                                                                                                               w_prime.clone();
                                                                                                                                                                                                                                                                                                                                                                           move
                                                                                                                                                                                                                                                                                                                                                                               |s_prime|
                                                                                                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_StoreT()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&w_prime,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             s_prime.clone()))))
                                                                                                                                                                                                                                                                                                                                                                       }))),
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
                                                                                                                             add(string("Functor0"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let functorStoreT1
                                                                                                                                                      =
                                                                                                                                                      functorStoreT1.clone();
                                                                                                                                                  move
                                                                                                                                                      |usd__unused|
                                                                                                                                                      &functorStoreT1
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>())))
                                                                                 }))
    }
    pub fn Control_Comonad_Store_Trans_comonadTransStoreT() -> &dyn Any {
        static Control_Comonad_Store_Trans_comonadTransStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Trans_comonadTransStoreT.get_or_init(||
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Trans_Class::Control_Comonad_Trans_Class_ComonadTransusd_Dict(),
                                                                                                        &&&add(string("lower"),
                                                                                                               &&Func1::new(move
                                                                                                                                |dictComonad|
                                                                                                                                {
                                                                                                                                    let Functor0 =
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(dictComonad)),
                                                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                                                                    &Func1::new({
                                                                                                                                                    let Functor0
                                                                                                                                                        =
                                                                                                                                                        Functor0.clone();
                                                                                                                                                    move
                                                                                                                                                        |v|
                                                                                                                                                        {
                                                                                                                                                            let matchValue:
                                                                                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                   &&&Functor0),
                                                                                                                                                                                                                                &&&Func1::new(move
                                                                                                                                                                                                                                                  |v1|
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                      v1),
                                                                                                                                                                                                                                                                                   &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                      }))),
                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                })
                                                                                                                                                        }
                                                                                                                                                })
                                                                                                                                }),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>())))
    }
    pub fn Control_Comonad_Store_Trans_comonadStoreT() -> &dyn Any {
        static Control_Comonad_Store_Trans_comonadStoreT:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Store_Trans_comonadStoreT.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |dictComonad|
                                                                                  {
                                                                                      let extendStoreT1 =
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad_Store_Trans::Control_Comonad_Store_Trans_extendStoreT(),
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
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                                                              &&&dictComonad),
                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                              }),
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
                                                                                                                                                   let extendStoreT1
                                                                                                                                                       =
                                                                                                                                                       extendStoreT1.clone();
                                                                                                                                                   move
                                                                                                                                                       |usd__unused|
                                                                                                                                                       &extendStoreT1
                                                                                                                                               }),
                                                                                                                                  empty::<string,
                                                                                                                                          &dyn Any>())))
                                                                                  }))
    }
}
