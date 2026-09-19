pub mod PureScript_Data_Array_ST_Iterator {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_fcf3066b::PureScript_Control_Monad_ST_Internal;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_acc76ca5::PureScript_Data_Array_ST;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug,)]
    pub enum Data_Array_ST_Iterator_Iterator {
        Data_Array_ST_Iterator_Iteratorusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Array_ST_Iterator_not() -> &dyn Any {
        static Data_Array_ST_Iterator_not: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_not.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                    &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()))
    }
    pub fn Data_Array_ST_Iterator_void() -> &dyn Any {
        static Data_Array_ST_Iterator_void: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_void.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                                     &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()))
    }
    pub fn Data_Array_ST_Iterator_void1() -> &dyn Any {
        static Data_Array_ST_Iterator_void1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_void1.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_void(),
                                                                                      &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()))
    }
    pub fn Data_Array_ST_Iterator_Iterator() -> &dyn Any {
        static Data_Array_ST_Iterator_Iterator: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_Iterator.get_or_init(||
                                                        &Func1::new(move
                                                                        |usd__arg1|
                                                                        Func1::new({
                                                                                       let usd__arg1
                                                                                           =
                                                                                           usd__arg1.clone();
                                                                                       move
                                                                                           |usd__arg2|
                                                                                           &LrcPtr::new(PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator::Data_Array_ST_Iterator_Iteratorusd_Ctor(usd__arg1,
                                                                                                                                                                                                                    usd__arg2.clone()))
                                                                                   })))
    }
    pub fn Data_Array_ST_Iterator_peek() -> &dyn Any {
        static Data_Array_ST_Iterator_peek: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_peek.get_or_init(||
                                                    &Func1::new(move |v|
                                                                    {
                                                                        let matchValue:
                                                                                LrcPtr<PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator> =
                                                                            Sharpurs_Prelude::unbox(v);
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                               &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_read(),
                                                                                                                                                                               &&&match matchValue.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                      PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator::Data_Array_ST_Iterator_Iteratorusd_Ctor(_,
                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                      =>
                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                  })),
                                                                                                         &&&Func1::new(move
                                                                                                                           |i|
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                               &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&match matchValue.as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                      PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator::Data_Array_ST_Iterator_Iteratorusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                      =>
                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                  },
                                                                                                                                                                                               i))))
                                                                    }))
    }
    pub fn Data_Array_ST_Iterator_next() -> &dyn Any {
        static Data_Array_ST_Iterator_next: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_next.get_or_init(||
                                                    &Func1::new(move |v|
                                                                    {
                                                                        let matchValue:
                                                                                LrcPtr<PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator> =
                                                                            Sharpurs_Prelude::unbox(v);
                                                                        let currentIndex =
                                                                            match matchValue.as_ref()
                                                                                {
                                                                                PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator::Data_Array_ST_Iterator_Iteratorusd_Ctor(_,
                                                                                                                                                                                            x)
                                                                                =>
                                                                                x.clone(),
                                                                            };
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                               &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_read(),
                                                                                                                                                                               &&&currentIndex)),
                                                                                                         &&&Func1::new({
                                                                                                                           let currentIndex
                                                                                                                               =
                                                                                                                               currentIndex.clone();
                                                                                                                           move
                                                                                                                               |i|
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                      &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_modify(),
                                                                                                                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                                                                                                                           |v1|
                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                                                               v1),
                                                                                                                                                                                                                                                                                                                            &&&1_i32))),
                                                                                                                                                                                                                                      &&&currentIndex)),
                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                  let i
                                                                                                                                                                                      =
                                                                                                                                                                                      i.clone();
                                                                                                                                                                                  move
                                                                                                                                                                                      |usd__unused|
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                          &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_applicativeST()),
                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator::Data_Array_ST_Iterator_Iteratorusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                             },
                                                                                                                                                                                                                                                          &&&i))
                                                                                                                                                                              }))
                                                                                                                       }))
                                                                    }))
    }
    pub fn Data_Array_ST_Iterator_pushWhile() -> &dyn Any {
        static Data_Array_ST_Iterator_pushWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_pushWhile.get_or_init(||
                                                         &Func1::new(move |p|
                                                                         &Func1::new({
                                                                                         let p
                                                                                             =
                                                                                             p.clone();
                                                                                         move
                                                                                             |iter|
                                                                                             &Func1::new({
                                                                                                             let iter
                                                                                                                 =
                                                                                                                 iter.clone();
                                                                                                             move
                                                                                                                 |array|
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                        &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_new(),
                                                                                                                                                                                                                        &&&false)),
                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                    let array
                                                                                                                                                                        =
                                                                                                                                                                        array.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |break_|
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_while(),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_not()),
                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_read(),
                                                                                                                                                                                                                                                                                                                  break_))),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                  &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_peek(),
                                                                                                                                                                                                                                                                                                                  &&&iter)),
                                                                                                                                                                                                                                            &&&Func1::new({
                                                                                                                                                                                                                                                              let break_
                                                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                                                  break_.clone();
                                                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                                                  |mx|
                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                      let matchValue:
                                                                                                                                                                                                                                                                              LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(mx);
                                                                                                                                                                                                                                                                      if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                             matchValue.as_ref()
                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                          if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                              Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                              _
                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                                                                                                                                                          }))
                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST::Data_Array_ST_push(),
                                                                                                                                                                                                                                                                                                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                                                               Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                               _
                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                                                                                                                                                                     &&&array)),
                                                                                                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                 |usd__unused|
                                                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_void()),
                                                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_next(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&iter))))
                                                                                                                                                                                                                                                                          } else {
                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_void1()),
                                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_write(),
                                                                                                                                                                                                                                                                                                                                                                                     &&&true),
                                                                                                                                                                                                                                                                                                                                                  &&&break_))
                                                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                                                      } else {
                                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_void1()),
                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_write(),
                                                                                                                                                                                                                                                                                                                                                                                 &&&true),
                                                                                                                                                                                                                                                                                                                                              &&&break_))
                                                                                                                                                                                                                                                                      }
                                                                                                                                                                                                                                                                  }
                                                                                                                                                                                                                                                          })))
                                                                                                                                                                }))
                                                                                                         })
                                                                                     })))
    }
    pub fn Data_Array_ST_Iterator_pushAll() -> &dyn Any {
        static Data_Array_ST_Iterator_pushAll: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_pushAll.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_pushWhile(),
                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                           &&&true)))
    }
    pub fn Data_Array_ST_Iterator_iterator() -> &dyn Any {
        static Data_Array_ST_Iterator_iterator: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_iterator.get_or_init(||
                                                        &Func1::new(move |f|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                               &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                 |usd__arg1|
                                                                                                                                                                                                 Func1::new({
                                                                                                                                                                                                                let usd__arg1
                                                                                                                                                                                                                    =
                                                                                                                                                                                                                    usd__arg1.clone();
                                                                                                                                                                                                                move
                                                                                                                                                                                                                    |usd__arg2|
                                                                                                                                                                                                                    &LrcPtr::new(PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_Iterator::Data_Array_ST_Iterator_Iteratorusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                             usd__arg2.clone()))
                                                                                                                                                                                                            })),
                                                                                                                                                                               f)),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_new(),
                                                                                                                                            &&&0_i32))))
    }
    pub fn Data_Array_ST_Iterator_iterate() -> &dyn Any {
        static Data_Array_ST_Iterator_iterate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_iterate.get_or_init(||
                                                       &Func1::new(move |iter|
                                                                       &Func1::new({
                                                                                       let iter
                                                                                           =
                                                                                           iter.clone();
                                                                                       move
                                                                                           |f|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                  &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_new(),
                                                                                                                                                                                                  &&&false)),
                                                                                                                            &&&Func1::new({
                                                                                                                                              let f
                                                                                                                                                  =
                                                                                                                                                  f.clone();
                                                                                                                                              move
                                                                                                                                                  |break_|
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_while(),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_not()),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_read(),
                                                                                                                                                                                                                                                                                            break_))),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_bindST()),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_next(),
                                                                                                                                                                                                                                                                                            &&&iter)),
                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                        let break_
                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                            break_.clone();
                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                            |mx|
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                let matchValue:
                                                                                                                                                                                                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(mx);
                                                                                                                                                                                                                                                match matchValue.as_ref()
                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                    Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_void1()),
                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_write(),
                                                                                                                                                                                                                                                                                                                                                           &&&true),
                                                                                                                                                                                                                                                                                                                        &&&break_)),
                                                                                                                                                                                                                                                    Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                            Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                            _
                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                }
                                                                                                                                                                                                                                            }
                                                                                                                                                                                                                                    })))
                                                                                                                                          }))
                                                                                   })))
    }
    pub fn Data_Array_ST_Iterator_exhausted() -> &dyn Any {
        static Data_Array_ST_Iterator_exhausted: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_ST_Iterator_exhausted.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                   &&&PureScript_Control_Monad_ST_Internal::Control_Monad_ST_Internal_functorST()),
                                                                                                                                                                &&&PureScript_Data_Maybe::Data_Maybe_isNothing())),
                                                                                          &&&PureScript_Data_Array_ST_Iterator::Data_Array_ST_Iterator_peek()))
    }
}
