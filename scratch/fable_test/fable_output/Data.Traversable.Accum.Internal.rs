pub mod PureScript_Data_Traversable_Accum_Internal {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_Traversable_Accum_Internal_StateR() -> &dyn Any {
        static Data_Traversable_Accum_Internal_StateR:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_StateR.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               x.clone()))
    }
    pub fn Data_Traversable_Accum_Internal_StateL() -> &dyn Any {
        static Data_Traversable_Accum_Internal_StateL:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_StateL.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               x.clone()))
    }
    pub fn Data_Traversable_Accum_Internal_stateR() -> &dyn Any {
        static Data_Traversable_Accum_Internal_stateR:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_stateR.get_or_init(||
                                                               &Func1::new(move
                                                                               |v|
                                                                               &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_Traversable_Accum_Internal_stateL() -> &dyn Any {
        static Data_Traversable_Accum_Internal_stateL:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_stateL.get_or_init(||
                                                               &Func1::new(move
                                                                               |v|
                                                                               &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_Traversable_Accum_Internal_functorStateR() -> &dyn Any {
        static Data_Traversable_Accum_Internal_functorStateR:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_functorStateR.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                       &&&add(string("map"),
                                                                                                              &&Func1::new(move
                                                                                                                               |f|
                                                                                                                               &Func1::new({
                                                                                                                                               let f
                                                                                                                                                   =
                                                                                                                                                   f.clone();
                                                                                                                                               move
                                                                                                                                                   |k|
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateR(),
                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                      let k
                                                                                                                                                                                                          =
                                                                                                                                                                                                          k.clone();
                                                                                                                                                                                                      move
                                                                                                                                                                                                          |s|
                                                                                                                                                                                                          {
                                                                                                                                                                                                              let matchValue =
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_stateR(),
                                                                                                                                                                                                                                                                                                                &&&k),
                                                                                                                                                                                                                                                                             s));
                                                                                                                                                                                                              {
                                                                                                                                                                                                                  let activePatternResult =
                                                                                                                                                                                                                      Sharpurs_Prelude::_007cHasProp_007c__007c(string("accum"),
                                                                                                                                                                                                                                                                &matchValue);
                                                                                                                                                                                                                  if activePatternResult.is_some()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                      let activePatternResult_1 =
                                                                                                                                                                                                                          Sharpurs_Prelude::_007cHasProp_007c__007c(string("value"),
                                                                                                                                                                                                                                                                    &matchValue);
                                                                                                                                                                                                                      if activePatternResult_1.is_some()
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                          let a =
                                                                                                                                                                                                                              getValue(activePatternResult_1);
                                                                                                                                                                                                                          let s1 =
                                                                                                                                                                                                                              getValue(activePatternResult);
                                                                                                                                                                                                                          &add(string("accum"),
                                                                                                                                                                                                                               &&s1,
                                                                                                                                                                                                                               add(string("value"),
                                                                                                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                     &&&a),
                                                                                                                                                                                                                                   empty::<string,
                                                                                                                                                                                                                                           &dyn Any>()))
                                                                                                                                                                                                                      } else {
                                                                                                                                                                                                                          panic!("{}",
                                                                                                                                                                                                                                 LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 15_i32,
                                  Data2: 432_i32,}).get_Message(),)
                                                                                                                                                                                                                      }
                                                                                                                                                                                                                  } else {
                                                                                                                                                                                                                      panic!("{}",
                                                                                                                                                                                                                             LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 15_i32,
                                  Data2: 432_i32,}).get_Message(),)
                                                                                                                                                                                                                  }
                                                                                                                                                                                                              }
                                                                                                                                                                                                          }
                                                                                                                                                                                                  }))
                                                                                                                                           })),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))
    }
    pub fn Data_Traversable_Accum_Internal_functorStateL() -> &dyn Any {
        static Data_Traversable_Accum_Internal_functorStateL:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_functorStateL.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                       &&&add(string("map"),
                                                                                                              &&Func1::new(move
                                                                                                                               |f|
                                                                                                                               &Func1::new({
                                                                                                                                               let f
                                                                                                                                                   =
                                                                                                                                                   f.clone();
                                                                                                                                               move
                                                                                                                                                   |k|
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateL(),
                                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                                      let k
                                                                                                                                                                                                          =
                                                                                                                                                                                                          k.clone();
                                                                                                                                                                                                      move
                                                                                                                                                                                                          |s|
                                                                                                                                                                                                          {
                                                                                                                                                                                                              let matchValue =
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_stateL(),
                                                                                                                                                                                                                                                                                                                &&&k),
                                                                                                                                                                                                                                                                             s));
                                                                                                                                                                                                              {
                                                                                                                                                                                                                  let activePatternResult =
                                                                                                                                                                                                                      Sharpurs_Prelude::_007cHasProp_007c__007c(string("accum"),
                                                                                                                                                                                                                                                                &matchValue);
                                                                                                                                                                                                                  if activePatternResult.is_some()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                      let activePatternResult_1 =
                                                                                                                                                                                                                          Sharpurs_Prelude::_007cHasProp_007c__007c(string("value"),
                                                                                                                                                                                                                                                                    &matchValue);
                                                                                                                                                                                                                      if activePatternResult_1.is_some()
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                          let a =
                                                                                                                                                                                                                              getValue(activePatternResult_1);
                                                                                                                                                                                                                          let s1 =
                                                                                                                                                                                                                              getValue(activePatternResult);
                                                                                                                                                                                                                          &add(string("accum"),
                                                                                                                                                                                                                               &&s1,
                                                                                                                                                                                                                               add(string("value"),
                                                                                                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                     &&&a),
                                                                                                                                                                                                                                   empty::<string,
                                                                                                                                                                                                                                           &dyn Any>()))
                                                                                                                                                                                                                      } else {
                                                                                                                                                                                                                          panic!("{}",
                                                                                                                                                                                                                                 LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 17_i32,
                                  Data2: 432_i32,}).get_Message(),)
                                                                                                                                                                                                                      }
                                                                                                                                                                                                                  } else {
                                                                                                                                                                                                                      panic!("{}",
                                                                                                                                                                                                                             LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 17_i32,
                                  Data2: 432_i32,}).get_Message(),)
                                                                                                                                                                                                                  }
                                                                                                                                                                                                              }
                                                                                                                                                                                                          }
                                                                                                                                                                                                  }))
                                                                                                                                           })),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))
    }
    pub fn Data_Traversable_Accum_Internal_applyStateR() -> &dyn Any {
        static Data_Traversable_Accum_Internal_applyStateR:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_applyStateR.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                     &&&add(string("apply"),
                                                                                                            &&Func1::new(move
                                                                                                                             |f|
                                                                                                                             &Func1::new({
                                                                                                                                             let f
                                                                                                                                                 =
                                                                                                                                                 f.clone();
                                                                                                                                             move
                                                                                                                                                 |x|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateR(),
                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                    let x
                                                                                                                                                                                                        =
                                                                                                                                                                                                        x.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |s|
                                                                                                                                                                                                        {
                                                                                                                                                                                                            let matchValue =
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_stateR(),
                                                                                                                                                                                                                                                                                                              &&&x),
                                                                                                                                                                                                                                                                           s));
                                                                                                                                                                                                            {
                                                                                                                                                                                                                let activePatternResult =
                                                                                                                                                                                                                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("accum"),
                                                                                                                                                                                                                                                              &matchValue);
                                                                                                                                                                                                                if activePatternResult.is_some()
                                                                                                                                                                                                                   {
                                                                                                                                                                                                                    let activePatternResult_1 =
                                                                                                                                                                                                                        Sharpurs_Prelude::_007cHasProp_007c__007c(string("value"),
                                                                                                                                                                                                                                                                  &matchValue);
                                                                                                                                                                                                                    if activePatternResult_1.is_some()
                                                                                                                                                                                                                       {
                                                                                                                                                                                                                        let s1 =
                                                                                                                                                                                                                            getValue(activePatternResult);
                                                                                                                                                                                                                        let x_prime =
                                                                                                                                                                                                                            getValue(activePatternResult_1);
                                                                                                                                                                                                                        let matchValue_1 =
                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_stateR(),
                                                                                                                                                                                                                                                                                                                          &&&f),
                                                                                                                                                                                                                                                                                       &&&s1));
                                                                                                                                                                                                                        {
                                                                                                                                                                                                                            let activePatternResult_2 =
                                                                                                                                                                                                                                Sharpurs_Prelude::_007cHasProp_007c__007c(string("accum"),
                                                                                                                                                                                                                                                                          &matchValue_1);
                                                                                                                                                                                                                            if activePatternResult_2.is_some()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                let activePatternResult_3 =
                                                                                                                                                                                                                                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("value"),
                                                                                                                                                                                                                                                                              &matchValue_1);
                                                                                                                                                                                                                                if activePatternResult_3.is_some()
                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                    let f_prime =
                                                                                                                                                                                                                                        getValue(activePatternResult_3);
                                                                                                                                                                                                                                    let s2 =
                                                                                                                                                                                                                                        getValue(activePatternResult_2);
                                                                                                                                                                                                                                    &add(string("accum"),
                                                                                                                                                                                                                                         &&s2,
                                                                                                                                                                                                                                         add(string("value"),
                                                                                                                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&f_prime,
                                                                                                                                                                                                                                                                               &&&x_prime),
                                                                                                                                                                                                                                             empty::<string,
                                                                                                                                                                                                                                                     &dyn Any>()))
                                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                                    panic!("{}",
                                                                                                                                                                                                                                           LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 19_i32,
                                  Data2: 661_i32,}).get_Message(),)
                                                                                                                                                                                                                                }
                                                                                                                                                                                                                            } else {
                                                                                                                                                                                                                                panic!("{}",
                                                                                                                                                                                                                                       LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 19_i32,
                                  Data2: 661_i32,}).get_Message(),)
                                                                                                                                                                                                                            }
                                                                                                                                                                                                                        }
                                                                                                                                                                                                                    } else {
                                                                                                                                                                                                                        panic!("{}",
                                                                                                                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 19_i32,
                                  Data2: 431_i32,}).get_Message(),)
                                                                                                                                                                                                                    }
                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                    panic!("{}",
                                                                                                                                                                                                                           LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 19_i32,
                                  Data2: 431_i32,}).get_Message(),)
                                                                                                                                                                                                                }
                                                                                                                                                                                                            }
                                                                                                                                                                                                        }
                                                                                                                                                                                                }))
                                                                                                                                         })),
                                                                                                            add(string("Functor0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_functorStateR()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Traversable_Accum_Internal_applyStateL() -> &dyn Any {
        static Data_Traversable_Accum_Internal_applyStateL:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_applyStateL.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                     &&&add(string("apply"),
                                                                                                            &&Func1::new(move
                                                                                                                             |f|
                                                                                                                             &Func1::new({
                                                                                                                                             let f
                                                                                                                                                 =
                                                                                                                                                 f.clone();
                                                                                                                                             move
                                                                                                                                                 |x|
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateL(),
                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                    let x
                                                                                                                                                                                                        =
                                                                                                                                                                                                        x.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |s|
                                                                                                                                                                                                        {
                                                                                                                                                                                                            let matchValue =
                                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_stateL(),
                                                                                                                                                                                                                                                                                                              &&&f),
                                                                                                                                                                                                                                                                           s));
                                                                                                                                                                                                            {
                                                                                                                                                                                                                let activePatternResult =
                                                                                                                                                                                                                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("accum"),
                                                                                                                                                                                                                                                              &matchValue);
                                                                                                                                                                                                                if activePatternResult.is_some()
                                                                                                                                                                                                                   {
                                                                                                                                                                                                                    let activePatternResult_1 =
                                                                                                                                                                                                                        Sharpurs_Prelude::_007cHasProp_007c__007c(string("value"),
                                                                                                                                                                                                                                                                  &matchValue);
                                                                                                                                                                                                                    if activePatternResult_1.is_some()
                                                                                                                                                                                                                       {
                                                                                                                                                                                                                        let f_prime =
                                                                                                                                                                                                                            getValue(activePatternResult_1);
                                                                                                                                                                                                                        let s1 =
                                                                                                                                                                                                                            getValue(activePatternResult);
                                                                                                                                                                                                                        let matchValue_1 =
                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_stateL(),
                                                                                                                                                                                                                                                                                                                          &&&x),
                                                                                                                                                                                                                                                                                       &&&s1));
                                                                                                                                                                                                                        {
                                                                                                                                                                                                                            let activePatternResult_2 =
                                                                                                                                                                                                                                Sharpurs_Prelude::_007cHasProp_007c__007c(string("accum"),
                                                                                                                                                                                                                                                                          &matchValue_1);
                                                                                                                                                                                                                            if activePatternResult_2.is_some()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                let activePatternResult_3 =
                                                                                                                                                                                                                                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("value"),
                                                                                                                                                                                                                                                                              &matchValue_1);
                                                                                                                                                                                                                                if activePatternResult_3.is_some()
                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                    let s2 =
                                                                                                                                                                                                                                        getValue(activePatternResult_2);
                                                                                                                                                                                                                                    let x_prime =
                                                                                                                                                                                                                                        getValue(activePatternResult_3);
                                                                                                                                                                                                                                    &add(string("accum"),
                                                                                                                                                                                                                                         &&s2,
                                                                                                                                                                                                                                         add(string("value"),
                                                                                                                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&f_prime,
                                                                                                                                                                                                                                                                               &&&x_prime),
                                                                                                                                                                                                                                             empty::<string,
                                                                                                                                                                                                                                                     &dyn Any>()))
                                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                                    panic!("{}",
                                                                                                                                                                                                                                           LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 21_i32,
                                  Data2: 661_i32,}).get_Message(),)
                                                                                                                                                                                                                                }
                                                                                                                                                                                                                            } else {
                                                                                                                                                                                                                                panic!("{}",
                                                                                                                                                                                                                                       LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 21_i32,
                                  Data2: 661_i32,}).get_Message(),)
                                                                                                                                                                                                                            }
                                                                                                                                                                                                                        }
                                                                                                                                                                                                                    } else {
                                                                                                                                                                                                                        panic!("{}",
                                                                                                                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 21_i32,
                                  Data2: 431_i32,}).get_Message(),)
                                                                                                                                                                                                                    }
                                                                                                                                                                                                                } else {
                                                                                                                                                                                                                    panic!("{}",
                                                                                                                                                                                                                           LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Traversable.Accum.Internal.fs"),
                                  Data1: 21_i32,
                                  Data2: 431_i32,}).get_Message(),)
                                                                                                                                                                                                                }
                                                                                                                                                                                                            }
                                                                                                                                                                                                        }
                                                                                                                                                                                                }))
                                                                                                                                         })),
                                                                                                            add(string("Functor0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_functorStateL()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Traversable_Accum_Internal_applicativeStateR() -> &dyn Any {
        static Data_Traversable_Accum_Internal_applicativeStateR:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_applicativeStateR.get_or_init(||
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                           &&&add(string("pure"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |a|
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateR(),
                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                      let a
                                                                                                                                                                                          =
                                                                                                                                                                                          a.clone();
                                                                                                                                                                                      move
                                                                                                                                                                                          |s|
                                                                                                                                                                                          &add(string("accum"),
                                                                                                                                                                                               s.clone(),
                                                                                                                                                                                               add(string("value"),
                                                                                                                                                                                                   &&a,
                                                                                                                                                                                                   empty::<string,
                                                                                                                                                                                                           &dyn Any>()))
                                                                                                                                                                                  }))),
                                                                                                                  add(string("Apply0"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |usd__unused|
                                                                                                                                       &PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_applyStateR()),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
    }
    pub fn Data_Traversable_Accum_Internal_applicativeStateL() -> &dyn Any {
        static Data_Traversable_Accum_Internal_applicativeStateL:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Traversable_Accum_Internal_applicativeStateL.get_or_init(||
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                           &&&add(string("pure"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |a|
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_StateL(),
                                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                                      let a
                                                                                                                                                                                          =
                                                                                                                                                                                          a.clone();
                                                                                                                                                                                      move
                                                                                                                                                                                          |s|
                                                                                                                                                                                          &add(string("accum"),
                                                                                                                                                                                               s.clone(),
                                                                                                                                                                                               add(string("value"),
                                                                                                                                                                                                   &&a,
                                                                                                                                                                                                   empty::<string,
                                                                                                                                                                                                           &dyn Any>()))
                                                                                                                                                                                  }))),
                                                                                                                  add(string("Apply0"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |usd__unused|
                                                                                                                                       &PureScript_Data_Traversable_Accum_Internal::Data_Traversable_Accum_Internal_applyStateL()),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
    }
}
