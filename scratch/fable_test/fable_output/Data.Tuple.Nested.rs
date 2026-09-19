pub mod PureScript_Data_Tuple_Nested {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Tuple_Nested_uncurry9() -> &dyn Any {
        static Data_Tuple_Nested_uncurry9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry9.get_or_init(||
                                                   &Func1::new(move |f_prime|
                                                                   &Func1::new({
                                                                                   let f_prime
                                                                                       =
                                                                                       f_prime.clone();
                                                                                   move
                                                                                       |v|
                                                                                       {
                                                                                           let matchValue =
                                                                                               Sharpurs_Prelude::unbox(&&f_prime);
                                                                                           let matchValue_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                           let activePatternResult:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_2:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_3:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_4:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_5:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_4.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_6:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_5.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_7:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_6.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                                                                                                                                                                                 &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                                                                                              &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                                                                           &&&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                                                        &&&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                     &&&match activePatternResult_4.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match activePatternResult_5.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                     }),
                                                                                                                                                               &&&match activePatternResult_6.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                 _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                  }),
                                                                                                                            &&&match activePatternResult_7.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                              _)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               })
                                                                                       }
                                                                               })))
    }
    pub fn Data_Tuple_Nested_uncurry8() -> &dyn Any {
        static Data_Tuple_Nested_uncurry8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry8.get_or_init(||
                                                   &Func1::new(move |f_prime|
                                                                   &Func1::new({
                                                                                   let f_prime
                                                                                       =
                                                                                       f_prime.clone();
                                                                                   move
                                                                                       |v|
                                                                                       {
                                                                                           let matchValue =
                                                                                               Sharpurs_Prelude::unbox(&&f_prime);
                                                                                           let matchValue_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                           let activePatternResult:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_2:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_3:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_4:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_5:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_4.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_6:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_5.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                                                                                              &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                                                                           &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                                                        &&&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                     &&&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match activePatternResult_4.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                     }),
                                                                                                                                                               &&&match activePatternResult_5.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                 _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                  }),
                                                                                                                            &&&match activePatternResult_6.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                              _)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               })
                                                                                       }
                                                                               })))
    }
    pub fn Data_Tuple_Nested_uncurry7() -> &dyn Any {
        static Data_Tuple_Nested_uncurry7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry7.get_or_init(||
                                                   &Func1::new(move |f_prime|
                                                                   &Func1::new({
                                                                                   let f_prime
                                                                                       =
                                                                                       f_prime.clone();
                                                                                   move
                                                                                       |v|
                                                                                       {
                                                                                           let matchValue =
                                                                                               Sharpurs_Prelude::unbox(&&f_prime);
                                                                                           let matchValue_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                           let activePatternResult:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_2:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_3:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_4:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_5:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_4.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                                                                           &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                                                        &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                     &&&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match activePatternResult_3.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                     }),
                                                                                                                                                               &&&match activePatternResult_4.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                 _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                  }),
                                                                                                                            &&&match activePatternResult_5.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                              _)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               })
                                                                                       }
                                                                               })))
    }
    pub fn Data_Tuple_Nested_uncurry6() -> &dyn Any {
        static Data_Tuple_Nested_uncurry6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry6.get_or_init(||
                                                   &Func1::new(move |f_prime|
                                                                   &Func1::new({
                                                                                   let f_prime
                                                                                       =
                                                                                       f_prime.clone();
                                                                                   move
                                                                                       |v|
                                                                                       {
                                                                                           let matchValue =
                                                                                               Sharpurs_Prelude::unbox(&&f_prime);
                                                                                           let matchValue_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                           let activePatternResult:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_2:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_3:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_4:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                                                        &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                     &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match activePatternResult_2.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                     }),
                                                                                                                                                               &&&match activePatternResult_3.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                 _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                  }),
                                                                                                                            &&&match activePatternResult_4.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                              _)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               })
                                                                                       }
                                                                               })))
    }
    pub fn Data_Tuple_Nested_uncurry5() -> &dyn Any {
        static Data_Tuple_Nested_uncurry5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry5.get_or_init(||
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
                                                                                           let activePatternResult:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_2:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_3:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                     &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                     }),
                                                                                                                                                               &&&match activePatternResult_2.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                 _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                  }),
                                                                                                                            &&&match activePatternResult_3.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                              _)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               })
                                                                                       }
                                                                               })))
    }
    pub fn Data_Tuple_Nested_uncurry4() -> &dyn Any {
        static Data_Tuple_Nested_uncurry4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry4.get_or_init(||
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
                                                                                           let activePatternResult:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_2:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                     &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match activePatternResult.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                     }),
                                                                                                                                                               &&&match activePatternResult_1.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                 _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                  }),
                                                                                                                            &&&match activePatternResult_2.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                              _)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               })
                                                                                       }
                                                                               })))
    }
    pub fn Data_Tuple_Nested_uncurry3() -> &dyn Any {
        static Data_Tuple_Nested_uncurry3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry3.get_or_init(||
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
                                                                                           let activePatternResult:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           let activePatternResult_1:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                     }),
                                                                                                                                                               &&&match activePatternResult.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                 _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                  }),
                                                                                                                            &&&match activePatternResult_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                              _)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               })
                                                                                       }
                                                                               })))
    }
    pub fn Data_Tuple_Nested_uncurry2() -> &dyn Any {
        static Data_Tuple_Nested_uncurry2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry2.get_or_init(||
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
                                                                                           let activePatternResult:
                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                               Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                 x)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                  });
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                 _)
                                                                                                                                                                      =>
                                                                                                                                                                      x.clone(),
                                                                                                                                                                  }),
                                                                                                                            &&&match activePatternResult.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                              _)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               })
                                                                                       }
                                                                               })))
    }
    pub fn Data_Tuple_Nested_uncurry10() -> &dyn Any {
        static Data_Tuple_Nested_uncurry10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry10.get_or_init(||
                                                    &Func1::new(move |f_prime|
                                                                    &Func1::new({
                                                                                    let f_prime
                                                                                        =
                                                                                        f_prime.clone();
                                                                                    move
                                                                                        |v|
                                                                                        {
                                                                                            let matchValue =
                                                                                                Sharpurs_Prelude::unbox(&&f_prime);
                                                                                            let matchValue_1:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                            let activePatternResult:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            let activePatternResult_1:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            let activePatternResult_2:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            let activePatternResult_3:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            let activePatternResult_4:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            let activePatternResult_5:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_4.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            let activePatternResult_6:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_5.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            let activePatternResult_7:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_6.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            let activePatternResult_8:
                                                                                                    LrcPtr<Data_Tuple_Tuple> =
                                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_7.as_ref()
                                                                                                                                       {
                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                  x)
                                                                                                                                       =>
                                                                                                                                       x.clone(),
                                                                                                                                   });
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                                                                                                                                                  &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                                                                                                                                                               &&&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                                                                                            &&&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                         &&&match activePatternResult_4.as_ref()
                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                      &&&match activePatternResult_5.as_ref()
                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                   &&&match activePatternResult_6.as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                          =>
                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                      }),
                                                                                                                                                                &&&match activePatternResult_7.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                  _)
                                                                                                                                                                       =>
                                                                                                                                                                       x.clone(),
                                                                                                                                                                   }),
                                                                                                                             &&&match activePatternResult_8.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                               _)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                })
                                                                                        }
                                                                                })))
    }
    pub fn Data_Tuple_Nested_uncurry1() -> &dyn Any {
        static Data_Tuple_Nested_uncurry1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_uncurry1.get_or_init(||
                                                   &Func1::new(move |f|
                                                                   &Func1::new({
                                                                                   let f
                                                                                       =
                                                                                       f.clone();
                                                                                   move
                                                                                       |v|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f),
                                                                                                                        &&&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                               {
                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                          _)
                                                                                                                               =>
                                                                                                                               x.clone(),
                                                                                                                           })
                                                                               })))
    }
    pub fn Data_Tuple_Nested_tuple9() -> &dyn Any {
        static Data_Tuple_Nested_tuple9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple9.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &Func1::new({
                                                                                 let a
                                                                                     =
                                                                                     a.clone();
                                                                                 move
                                                                                     |b|
                                                                                     &Func1::new({
                                                                                                     let b
                                                                                                         =
                                                                                                         b.clone();
                                                                                                     move
                                                                                                         |c|
                                                                                                         &Func1::new({
                                                                                                                         let c
                                                                                                                             =
                                                                                                                             c.clone();
                                                                                                                         move
                                                                                                                             |d|
                                                                                                                             &Func1::new({
                                                                                                                                             let d
                                                                                                                                                 =
                                                                                                                                                 d.clone();
                                                                                                                                             move
                                                                                                                                                 |e|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let e
                                                                                                                                                                     =
                                                                                                                                                                     e.clone();
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
                                                                                                                                                                                                             |h|
                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                             let h
                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                 h.clone();
                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                 |i|
                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&g,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&h,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(i.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &PureScript_Data_Unit::Data_Unit_unit()))))))))))))))))))
                                                                                                                                                                                                                         })
                                                                                                                                                                                                     })
                                                                                                                                                                                 })
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_tuple8() -> &dyn Any {
        static Data_Tuple_Nested_tuple8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple8.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &Func1::new({
                                                                                 let a
                                                                                     =
                                                                                     a.clone();
                                                                                 move
                                                                                     |b|
                                                                                     &Func1::new({
                                                                                                     let b
                                                                                                         =
                                                                                                         b.clone();
                                                                                                     move
                                                                                                         |c|
                                                                                                         &Func1::new({
                                                                                                                         let c
                                                                                                                             =
                                                                                                                             c.clone();
                                                                                                                         move
                                                                                                                             |d|
                                                                                                                             &Func1::new({
                                                                                                                                             let d
                                                                                                                                                 =
                                                                                                                                                 d.clone();
                                                                                                                                             move
                                                                                                                                                 |e|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let e
                                                                                                                                                                     =
                                                                                                                                                                     e.clone();
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
                                                                                                                                                                                                             |h|
                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&g,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(h.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &PureScript_Data_Unit::Data_Unit_unit()))))))))))))))))
                                                                                                                                                                                                     })
                                                                                                                                                                                 })
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_tuple7() -> &dyn Any {
        static Data_Tuple_Nested_tuple7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple7.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &Func1::new({
                                                                                 let a
                                                                                     =
                                                                                     a.clone();
                                                                                 move
                                                                                     |b|
                                                                                     &Func1::new({
                                                                                                     let b
                                                                                                         =
                                                                                                         b.clone();
                                                                                                     move
                                                                                                         |c|
                                                                                                         &Func1::new({
                                                                                                                         let c
                                                                                                                             =
                                                                                                                             c.clone();
                                                                                                                         move
                                                                                                                             |d|
                                                                                                                             &Func1::new({
                                                                                                                                             let d
                                                                                                                                                 =
                                                                                                                                                 d.clone();
                                                                                                                                             move
                                                                                                                                                 |e|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let e
                                                                                                                                                                     =
                                                                                                                                                                     e.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |f|
                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                     let f
                                                                                                                                                                                         =
                                                                                                                                                                                         f.clone();
                                                                                                                                                                                     move
                                                                                                                                                                                         |g|
                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(g.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &PureScript_Data_Unit::Data_Unit_unit()))))))))))))))
                                                                                                                                                                                 })
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_tuple6() -> &dyn Any {
        static Data_Tuple_Nested_tuple6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple6.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &Func1::new({
                                                                                 let a
                                                                                     =
                                                                                     a.clone();
                                                                                 move
                                                                                     |b|
                                                                                     &Func1::new({
                                                                                                     let b
                                                                                                         =
                                                                                                         b.clone();
                                                                                                     move
                                                                                                         |c|
                                                                                                         &Func1::new({
                                                                                                                         let c
                                                                                                                             =
                                                                                                                             c.clone();
                                                                                                                         move
                                                                                                                             |d|
                                                                                                                             &Func1::new({
                                                                                                                                             let d
                                                                                                                                                 =
                                                                                                                                                 d.clone();
                                                                                                                                             move
                                                                                                                                                 |e|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let e
                                                                                                                                                                     =
                                                                                                                                                                     e.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |f|
                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(f.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     &PureScript_Data_Unit::Data_Unit_unit()))))))))))))
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_tuple5() -> &dyn Any {
        static Data_Tuple_Nested_tuple5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple5.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &Func1::new({
                                                                                 let a
                                                                                     =
                                                                                     a.clone();
                                                                                 move
                                                                                     |b|
                                                                                     &Func1::new({
                                                                                                     let b
                                                                                                         =
                                                                                                         b.clone();
                                                                                                     move
                                                                                                         |c|
                                                                                                         &Func1::new({
                                                                                                                         let c
                                                                                                                             =
                                                                                                                             c.clone();
                                                                                                                         move
                                                                                                                             |d|
                                                                                                                             &Func1::new({
                                                                                                                                             let d
                                                                                                                                                 =
                                                                                                                                                 d.clone();
                                                                                                                                             move
                                                                                                                                                 |e|
                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(e.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                         &PureScript_Data_Unit::Data_Unit_unit()))))))))))
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_tuple4() -> &dyn Any {
        static Data_Tuple_Nested_tuple4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple4.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &Func1::new({
                                                                                 let a
                                                                                     =
                                                                                     a.clone();
                                                                                 move
                                                                                     |b|
                                                                                     &Func1::new({
                                                                                                     let b
                                                                                                         =
                                                                                                         b.clone();
                                                                                                     move
                                                                                                         |c|
                                                                                                         &Func1::new({
                                                                                                                         let c
                                                                                                                             =
                                                                                                                             c.clone();
                                                                                                                         move
                                                                                                                             |d|
                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(d.clone(),
                                                                                                                                                                                                                                                                                                                                                             &PureScript_Data_Unit::Data_Unit_unit()))))))))
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_tuple3() -> &dyn Any {
        static Data_Tuple_Nested_tuple3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple3.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &Func1::new({
                                                                                 let a
                                                                                     =
                                                                                     a.clone();
                                                                                 move
                                                                                     |b|
                                                                                     &Func1::new({
                                                                                                     let b
                                                                                                         =
                                                                                                         b.clone();
                                                                                                     move
                                                                                                         |c|
                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(c.clone(),
                                                                                                                                                                                                                                                                                 &PureScript_Data_Unit::Data_Unit_unit()))))))
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_tuple2() -> &dyn Any {
        static Data_Tuple_Nested_tuple2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple2.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &Func1::new({
                                                                                 let a
                                                                                     =
                                                                                     a.clone();
                                                                                 move
                                                                                     |b|
                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                             &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(b.clone(),
                                                                                                                                                                                                     &PureScript_Data_Unit::Data_Unit_unit()))))
                                                                             })))
    }
    pub fn Data_Tuple_Nested_tuple10() -> &dyn Any {
        static Data_Tuple_Nested_tuple10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple10.get_or_init(||
                                                  &Func1::new(move |a|
                                                                  &Func1::new({
                                                                                  let a
                                                                                      =
                                                                                      a.clone();
                                                                                  move
                                                                                      |b|
                                                                                      &Func1::new({
                                                                                                      let b
                                                                                                          =
                                                                                                          b.clone();
                                                                                                      move
                                                                                                          |c|
                                                                                                          &Func1::new({
                                                                                                                          let c
                                                                                                                              =
                                                                                                                              c.clone();
                                                                                                                          move
                                                                                                                              |d|
                                                                                                                              &Func1::new({
                                                                                                                                              let d
                                                                                                                                                  =
                                                                                                                                                  d.clone();
                                                                                                                                              move
                                                                                                                                                  |e|
                                                                                                                                                  &Func1::new({
                                                                                                                                                                  let e
                                                                                                                                                                      =
                                                                                                                                                                      e.clone();
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
                                                                                                                                                                                                              |h|
                                                                                                                                                                                                              &Func1::new({
                                                                                                                                                                                                                              let h
                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                  h.clone();
                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                  |i|
                                                                                                                                                                                                                                  &Func1::new({
                                                                                                                                                                                                                                                  let i
                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                      i.clone();
                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                      |j|
                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&g,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&h,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&i,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(j.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      &PureScript_Data_Unit::Data_Unit_unit()))))))))))))))))))))
                                                                                                                                                                                                                                              })
                                                                                                                                                                                                                          })
                                                                                                                                                                                                      })
                                                                                                                                                                                  })
                                                                                                                                                              })
                                                                                                                                          })
                                                                                                                      })
                                                                                                  })
                                                                              })))
    }
    pub fn Data_Tuple_Nested_tuple1() -> &dyn Any {
        static Data_Tuple_Nested_tuple1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_tuple1.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                         &PureScript_Data_Unit::Data_Unit_unit()))))
    }
    pub fn Data_Tuple_Nested_over9() -> &dyn Any {
        static Data_Tuple_Nested_over9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over9.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_2:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_3:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_4:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_5:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_4.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_6:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_5.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_7:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_6.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                _)
                                                                                                                                                     =>
                                                                                                                                                     x.clone(),
                                                                                                                                                 },
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         },
                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_5.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&match activePatternResult_7.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &match activePatternResult_7.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 }))))))))))))))))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Tuple_Nested_over8() -> &dyn Any {
        static Data_Tuple_Nested_over8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over8.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_2:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_3:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_4:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_5:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_4.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_6:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_5.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                _)
                                                                                                                                                     =>
                                                                                                                                                     x.clone(),
                                                                                                                                                 },
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         },
                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_5.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&match activePatternResult_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &match activePatternResult_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         }))))))))))))))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Tuple_Nested_over7() -> &dyn Any {
        static Data_Tuple_Nested_over7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over7.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_2:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_3:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_4:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_5:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_4.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                _)
                                                                                                                                                     =>
                                                                                                                                                     x.clone(),
                                                                                                                                                 },
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         },
                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&match activePatternResult_5.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &match activePatternResult_5.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 }))))))))))))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Tuple_Nested_over6() -> &dyn Any {
        static Data_Tuple_Nested_over6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over6.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_2:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_3:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_4:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                _)
                                                                                                                                                     =>
                                                                                                                                                     x.clone(),
                                                                                                                                                 },
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         },
                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &&&match activePatternResult_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                                                                                                                                                                                        &match activePatternResult_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                         }))))))))))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Tuple_Nested_over5() -> &dyn Any {
        static Data_Tuple_Nested_over5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over5.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_2:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_3:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                _)
                                                                                                                                                     =>
                                                                                                                                                     x.clone(),
                                                                                                                                                 },
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         },
                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                                 &&&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                                                                                                                                &match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                 }))))))))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Tuple_Nested_over4() -> &dyn Any {
        static Data_Tuple_Nested_over4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over4.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_2:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                _)
                                                                                                                                                     =>
                                                                                                                                                     x.clone(),
                                                                                                                                                 },
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         },
                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                         &&&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                                                                        &match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                         }))))))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Tuple_Nested_over3() -> &dyn Any {
        static Data_Tuple_Nested_over3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over3.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        let activePatternResult_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                _)
                                                                                                                                                     =>
                                                                                                                                                     x.clone(),
                                                                                                                                                 },
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         },
                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                 &&&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                &match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 }))))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Tuple_Nested_over2() -> &dyn Any {
        static Data_Tuple_Nested_over2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over2.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        let activePatternResult:
                                                                                                LrcPtr<Data_Tuple_Tuple> =
                                                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                              x)
                                                                                                                                   =>
                                                                                                                                   x.clone(),
                                                                                                                               });
                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                _)
                                                                                                                                                     =>
                                                                                                                                                     x.clone(),
                                                                                                                                                 },
                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                         &&&match activePatternResult.as_ref()
                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                        &match activePatternResult.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         }))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_Tuple_Nested_over10() -> &dyn Any {
        static Data_Tuple_Nested_over10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over10.get_or_init(||
                                                 &Func1::new(move |o|
                                                                 &Func1::new({
                                                                                 let o
                                                                                     =
                                                                                     o.clone();
                                                                                 move
                                                                                     |v|
                                                                                     {
                                                                                         let matchValue =
                                                                                             Sharpurs_Prelude::unbox(&&o);
                                                                                         let matchValue_1:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                         let activePatternResult:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         let activePatternResult_1:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         let activePatternResult_2:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_1.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         let activePatternResult_3:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_2.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         let activePatternResult_4:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_3.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         let activePatternResult_5:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_4.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         let activePatternResult_6:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_5.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         let activePatternResult_7:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_6.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         let activePatternResult_8:
                                                                                                 LrcPtr<Data_Tuple_Tuple> =
                                                                                             Sharpurs_Prelude::_007cUnbox_007c(&match activePatternResult_7.as_ref()
                                                                                                                                    {
                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                               x)
                                                                                                                                    =>
                                                                                                                                    x.clone(),
                                                                                                                                });
                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                 _)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  },
                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult.as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                          },
                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_1.as_ref()
                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                  },
                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_2.as_ref()
                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                  },
                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_4.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_5.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  },
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match activePatternResult_7.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  },
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&match activePatternResult_8.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             }),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &match activePatternResult_8.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          }))))))))))))))))))))
                                                                                     }
                                                                             })))
    }
    pub fn Data_Tuple_Nested_over1() -> &dyn Any {
        static Data_Tuple_Nested_over1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_over1.get_or_init(||
                                                &Func1::new(move |o|
                                                                &Func1::new({
                                                                                let o
                                                                                    =
                                                                                    o.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&o);
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
                                                                            })))
    }
    pub fn Data_Tuple_Nested_get9() -> &dyn Any {
        static Data_Tuple_Nested_get9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get9.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                                                                                            Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                        }).as_ref()
                                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                                               }).as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                      }).as_ref()
                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                             }).as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    }).as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                               =>
                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                           }).as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                 x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  }).as_ref()
                                                                                                             {
                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                        x)
                                                                                                             =>
                                                                                                             x.clone(),
                                                                                                         }).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_get8() -> &dyn Any {
        static Data_Tuple_Nested_get8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get8.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                                                   Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                                               }).as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                      }).as_ref()
                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                             }).as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    }).as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                               =>
                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                           }).as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                 x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  }).as_ref()
                                                                                                             {
                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                        x)
                                                                                                             =>
                                                                                                             x.clone(),
                                                                                                         }).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_get7() -> &dyn Any {
        static Data_Tuple_Nested_get7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get7.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                      }).as_ref()
                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                             }).as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    }).as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                               =>
                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                           }).as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                 x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  }).as_ref()
                                                                                                             {
                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                        x)
                                                                                                             =>
                                                                                                             x.clone(),
                                                                                                         }).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_get6() -> &dyn Any {
        static Data_Tuple_Nested_get6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get6.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                             }).as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    }).as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                               =>
                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                           }).as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                 x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  }).as_ref()
                                                                                                             {
                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                        x)
                                                                                                             =>
                                                                                                             x.clone(),
                                                                                                         }).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_get5() -> &dyn Any {
        static Data_Tuple_Nested_get5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get5.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    }).as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                               =>
                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                           }).as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                 x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  }).as_ref()
                                                                                                             {
                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                        x)
                                                                                                             =>
                                                                                                             x.clone(),
                                                                                                         }).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_get4() -> &dyn Any {
        static Data_Tuple_Nested_get4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get4.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                               {
                                                                                                                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                          x)
                                                                                                                                                                                               =>
                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                           }).as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                 x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  }).as_ref()
                                                                                                             {
                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                        x)
                                                                                                             =>
                                                                                                             x.clone(),
                                                                                                         }).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_get3() -> &dyn Any {
        static Data_Tuple_Nested_get3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get3.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                 x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  }).as_ref()
                                                                                                             {
                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                        x)
                                                                                                             =>
                                                                                                             x.clone(),
                                                                                                         }).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_get2() -> &dyn Any {
        static Data_Tuple_Nested_get2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get2.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                             {
                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                        x)
                                                                                                             =>
                                                                                                             x.clone(),
                                                                                                         }).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_get10() -> &dyn Any {
        static Data_Tuple_Nested_get10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get10.get_or_init(||
                                                &Func1::new(move |v|
                                                                &match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::_007cUnbox_007c(&match Sharpurs_Prelude::unbox(v).as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                                                                                                                      Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                  }).as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                                                                                                                             Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                         }).as_ref()
                                                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                                                }).as_ref()
                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                       }).as_ref()
                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                  Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                             x)
                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                              }).as_ref()
                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                         Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                     }).as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                            }).as_ref()
                                                                                                                                                       {
                                                                                                                                                       Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                  x)
                                                                                                                                                       =>
                                                                                                                                                       x.clone(),
                                                                                                                                                   }).as_ref()
                                                                                                              {
                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                         x)
                                                                                                              =>
                                                                                                              x.clone(),
                                                                                                          }).as_ref()
                                                                     {
                                                                     Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                _)
                                                                     =>
                                                                     x.clone(),
                                                                 }))
    }
    pub fn Data_Tuple_Nested_get1() -> &dyn Any {
        static Data_Tuple_Nested_get1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_get1.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::unbox(v).as_ref()
                                                                    {
                                                                    Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                               _)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Tuple_Nested_curry9() -> &dyn Any {
        static Data_Tuple_Nested_curry9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry9.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
                                                                                 move
                                                                                     |f_prime|
                                                                                     &Func1::new({
                                                                                                     let f_prime
                                                                                                         =
                                                                                                         f_prime.clone();
                                                                                                     move
                                                                                                         |a|
                                                                                                         &Func1::new({
                                                                                                                         let a
                                                                                                                             =
                                                                                                                             a.clone();
                                                                                                                         move
                                                                                                                             |b|
                                                                                                                             &Func1::new({
                                                                                                                                             let b
                                                                                                                                                 =
                                                                                                                                                 b.clone();
                                                                                                                                             move
                                                                                                                                                 |c|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let c
                                                                                                                                                                     =
                                                                                                                                                                     c.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |d|
                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                     let d
                                                                                                                                                                                         =
                                                                                                                                                                                         d.clone();
                                                                                                                                                                                     move
                                                                                                                                                                                         |e|
                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                         let e
                                                                                                                                                                                                             =
                                                                                                                                                                                                             e.clone();
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
                                                                                                                                                                                                                                                     |h|
                                                                                                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                                                                                                     let h
                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                         h.clone();
                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                         |i|
                                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&f_prime,
                                                                                                                                                                                                                                                                                                          &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&g,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&h,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(i.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &z)))))))))))))))))))
                                                                                                                                                                                                                                                                 })
                                                                                                                                                                                                                                             })
                                                                                                                                                                                                                         })
                                                                                                                                                                                                     })
                                                                                                                                                                                 })
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_curry8() -> &dyn Any {
        static Data_Tuple_Nested_curry8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry8.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
                                                                                 move
                                                                                     |f_prime|
                                                                                     &Func1::new({
                                                                                                     let f_prime
                                                                                                         =
                                                                                                         f_prime.clone();
                                                                                                     move
                                                                                                         |a|
                                                                                                         &Func1::new({
                                                                                                                         let a
                                                                                                                             =
                                                                                                                             a.clone();
                                                                                                                         move
                                                                                                                             |b|
                                                                                                                             &Func1::new({
                                                                                                                                             let b
                                                                                                                                                 =
                                                                                                                                                 b.clone();
                                                                                                                                             move
                                                                                                                                                 |c|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let c
                                                                                                                                                                     =
                                                                                                                                                                     c.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |d|
                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                     let d
                                                                                                                                                                                         =
                                                                                                                                                                                         d.clone();
                                                                                                                                                                                     move
                                                                                                                                                                                         |e|
                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                         let e
                                                                                                                                                                                                             =
                                                                                                                                                                                                             e.clone();
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
                                                                                                                                                                                                                                                     |h|
                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&f_prime,
                                                                                                                                                                                                                                                                                      &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&g,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(h.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &z)))))))))))))))))
                                                                                                                                                                                                                                             })
                                                                                                                                                                                                                         })
                                                                                                                                                                                                     })
                                                                                                                                                                                 })
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_curry7() -> &dyn Any {
        static Data_Tuple_Nested_curry7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry7.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
                                                                                 move
                                                                                     |f_prime|
                                                                                     &Func1::new({
                                                                                                     let f_prime
                                                                                                         =
                                                                                                         f_prime.clone();
                                                                                                     move
                                                                                                         |a|
                                                                                                         &Func1::new({
                                                                                                                         let a
                                                                                                                             =
                                                                                                                             a.clone();
                                                                                                                         move
                                                                                                                             |b|
                                                                                                                             &Func1::new({
                                                                                                                                             let b
                                                                                                                                                 =
                                                                                                                                                 b.clone();
                                                                                                                                             move
                                                                                                                                                 |c|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let c
                                                                                                                                                                     =
                                                                                                                                                                     c.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |d|
                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                     let d
                                                                                                                                                                                         =
                                                                                                                                                                                         d.clone();
                                                                                                                                                                                     move
                                                                                                                                                                                         |e|
                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                         let e
                                                                                                                                                                                                             =
                                                                                                                                                                                                             e.clone();
                                                                                                                                                                                                         move
                                                                                                                                                                                                             |f|
                                                                                                                                                                                                             &Func1::new({
                                                                                                                                                                                                                             let f
                                                                                                                                                                                                                                 =
                                                                                                                                                                                                                                 f.clone();
                                                                                                                                                                                                                             move
                                                                                                                                                                                                                                 |g|
                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&f_prime,
                                                                                                                                                                                                                                                                  &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(g.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &z)))))))))))))))
                                                                                                                                                                                                                         })
                                                                                                                                                                                                     })
                                                                                                                                                                                 })
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_curry6() -> &dyn Any {
        static Data_Tuple_Nested_curry6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry6.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
                                                                                 move
                                                                                     |f_prime|
                                                                                     &Func1::new({
                                                                                                     let f_prime
                                                                                                         =
                                                                                                         f_prime.clone();
                                                                                                     move
                                                                                                         |a|
                                                                                                         &Func1::new({
                                                                                                                         let a
                                                                                                                             =
                                                                                                                             a.clone();
                                                                                                                         move
                                                                                                                             |b|
                                                                                                                             &Func1::new({
                                                                                                                                             let b
                                                                                                                                                 =
                                                                                                                                                 b.clone();
                                                                                                                                             move
                                                                                                                                                 |c|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let c
                                                                                                                                                                     =
                                                                                                                                                                     c.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |d|
                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                     let d
                                                                                                                                                                                         =
                                                                                                                                                                                         d.clone();
                                                                                                                                                                                     move
                                                                                                                                                                                         |e|
                                                                                                                                                                                         &Func1::new({
                                                                                                                                                                                                         let e
                                                                                                                                                                                                             =
                                                                                                                                                                                                             e.clone();
                                                                                                                                                                                                         move
                                                                                                                                                                                                             |f|
                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&f_prime,
                                                                                                                                                                                                                                              &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(f.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                &z)))))))))))))
                                                                                                                                                                                                     })
                                                                                                                                                                                 })
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_curry5() -> &dyn Any {
        static Data_Tuple_Nested_curry5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry5.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
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
                                                                                                                             |b|
                                                                                                                             &Func1::new({
                                                                                                                                             let b
                                                                                                                                                 =
                                                                                                                                                 b.clone();
                                                                                                                                             move
                                                                                                                                                 |c|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let c
                                                                                                                                                                     =
                                                                                                                                                                     c.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |d|
                                                                                                                                                                     &Func1::new({
                                                                                                                                                                                     let d
                                                                                                                                                                                         =
                                                                                                                                                                                         d.clone();
                                                                                                                                                                                     move
                                                                                                                                                                                         |e|
                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                          &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(e.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &z)))))))))))
                                                                                                                                                                                 })
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_curry4() -> &dyn Any {
        static Data_Tuple_Nested_curry4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry4.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
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
                                                                                                                             |b|
                                                                                                                             &Func1::new({
                                                                                                                                             let b
                                                                                                                                                 =
                                                                                                                                                 b.clone();
                                                                                                                                             move
                                                                                                                                                 |c|
                                                                                                                                                 &Func1::new({
                                                                                                                                                                 let c
                                                                                                                                                                     =
                                                                                                                                                                     c.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |d|
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                      &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(d.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                        &z)))))))))
                                                                                                                                                             })
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_curry3() -> &dyn Any {
        static Data_Tuple_Nested_curry3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry3.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
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
                                                                                                                             |b|
                                                                                                                             &Func1::new({
                                                                                                                                             let b
                                                                                                                                                 =
                                                                                                                                                 b.clone();
                                                                                                                                             move
                                                                                                                                                 |c|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                  &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                            &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(c.clone(),
                                                                                                                                                                                                                                                                                                                                                            &z)))))))
                                                                                                                                         })
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_curry2() -> &dyn Any {
        static Data_Tuple_Nested_curry2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry2.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
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
                                                                                                                             |b|
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                              &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(b.clone(),
                                                                                                                                                                                                                                                                                &z)))))
                                                                                                                     })
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Tuple_Nested_curry10() -> &dyn Any {
        static Data_Tuple_Nested_curry10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry10.get_or_init(||
                                                  &Func1::new(move |z|
                                                                  &Func1::new({
                                                                                  let z
                                                                                      =
                                                                                      z.clone();
                                                                                  move
                                                                                      |f_prime|
                                                                                      &Func1::new({
                                                                                                      let f_prime
                                                                                                          =
                                                                                                          f_prime.clone();
                                                                                                      move
                                                                                                          |a|
                                                                                                          &Func1::new({
                                                                                                                          let a
                                                                                                                              =
                                                                                                                              a.clone();
                                                                                                                          move
                                                                                                                              |b|
                                                                                                                              &Func1::new({
                                                                                                                                              let b
                                                                                                                                                  =
                                                                                                                                                  b.clone();
                                                                                                                                              move
                                                                                                                                                  |c|
                                                                                                                                                  &Func1::new({
                                                                                                                                                                  let c
                                                                                                                                                                      =
                                                                                                                                                                      c.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |d|
                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                      let d
                                                                                                                                                                                          =
                                                                                                                                                                                          d.clone();
                                                                                                                                                                                      move
                                                                                                                                                                                          |e|
                                                                                                                                                                                          &Func1::new({
                                                                                                                                                                                                          let e
                                                                                                                                                                                                              =
                                                                                                                                                                                                              e.clone();
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
                                                                                                                                                                                                                                                      |h|
                                                                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                                                                      let h
                                                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                                                          h.clone();
                                                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                                                          |i|
                                                                                                                                                                                                                                                                          &Func1::new({
                                                                                                                                                                                                                                                                                          let i
                                                                                                                                                                                                                                                                                              =
                                                                                                                                                                                                                                                                                              i.clone();
                                                                                                                                                                                                                                                                                          move
                                                                                                                                                                                                                                                                                              |j|
                                                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&f_prime,
                                                                                                                                                                                                                                                                                                                               &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&b,
                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&c,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&d,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&e,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&f,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&g,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&h,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&i,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(j.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &z)))))))))))))))))))))
                                                                                                                                                                                                                                                                                      })
                                                                                                                                                                                                                                                                  })
                                                                                                                                                                                                                                              })
                                                                                                                                                                                                                          })
                                                                                                                                                                                                      })
                                                                                                                                                                                  })
                                                                                                                                                              })
                                                                                                                                          })
                                                                                                                      })
                                                                                                  })
                                                                              })))
    }
    pub fn Data_Tuple_Nested_curry1() -> &dyn Any {
        static Data_Tuple_Nested_curry1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Nested_curry1.get_or_init(||
                                                 &Func1::new(move |z|
                                                                 &Func1::new({
                                                                                 let z
                                                                                     =
                                                                                     z.clone();
                                                                                 move
                                                                                     |f|
                                                                                     &Func1::new({
                                                                                                     let f
                                                                                                         =
                                                                                                         f.clone();
                                                                                                     move
                                                                                                         |a|
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                          &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                                                                                                    &z)))
                                                                                                 })
                                                                             })))
    }
}
