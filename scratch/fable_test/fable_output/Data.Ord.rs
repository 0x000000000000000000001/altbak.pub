pub mod PureScript_Data_Ord {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub mod Data_Ord_FFI {
        use super::*;
        pub fn ordIntImpl<a: Clone + 'static, b: Clone + 'static, c: Clone +
                          'static>(lt: a, eq: a, gt: a, x: b, y: c) -> a {
            let x_: i32 = Sharpurs_Prelude::unbox(&&x);
            let y_: i32 = Sharpurs_Prelude::unbox(&&y);
            if x_ < y_ { lt } else { if x_ == y_ { eq } else { gt } }
        }
        pub fn ordNumberImpl<a: Clone + 'static, b: Clone + 'static,
                             c: Clone +
                             'static>(lt: a, eq: a, gt: a, x: b, y: c) -> a {
            let x_: f64 = Sharpurs_Prelude::unbox(&&x);
            let y_: f64 = Sharpurs_Prelude::unbox(&&y);
            if x_ < y_ { lt } else { if x_ == y_ { eq } else { gt } }
        }
        pub fn ordStringImpl<a: Clone + 'static, b: Clone + 'static,
                             c: Clone +
                             'static>(lt: a, eq: a, gt: a, x: b, y: c) -> a {
            let x_: string = Sharpurs_Prelude::unbox(&&x);
            let y_: string = Sharpurs_Prelude::unbox(&&y);
            if x_.clone() < y_.clone() {
                lt
            } else { if x_ == y_ { eq } else { gt } }
        }
        pub fn ordCharImpl<a: Clone + 'static, b: Clone + 'static, c: Clone +
                           'static>(lt: a, eq: a, gt: a, x: b, y: c) -> a {
            let x_: char = Sharpurs_Prelude::unbox(&&x);
            let y_: char = Sharpurs_Prelude::unbox(&&y);
            if x_ < y_ { lt } else { if x_ == y_ { eq } else { gt } }
        }
        pub fn ordBooleanImpl<a: Clone + 'static, b: Clone + 'static,
                              c: Clone +
                              'static>(lt: a, eq: a, gt: a, x: b, y: c) -> a {
            let x_: bool = Sharpurs_Prelude::unbox(&&x);
            let y_: bool = Sharpurs_Prelude::unbox(&&y);
            if x_ < y_ { lt } else { if x_ == y_ { eq } else { gt } }
        }
        pub fn ordArrayImpl() -> &dyn Any {
            static ordArrayImpl: MutCell<Option<&dyn Any>> =
                MutCell::new(None);
            ordArrayImpl.get_or_init(|| Sharpurs_Prelude::undefined())
        }
    }
    pub fn Data_Ord_ordArrayImpl() -> &dyn Any {
        static Data_Ord_ordArrayImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordArrayImpl.get_or_init(||
                                              &PureScript_Data_Ord::Data_Ord_FFI::ordArrayImpl())
    }
    pub fn Data_Ord_ordBooleanImpl() -> &dyn Any {
        static Data_Ord_ordBooleanImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordBooleanImpl.get_or_init(||
                                                &Func1::new(move |arg0|
                                                                &Func1::new({
                                                                                let arg0
                                                                                    =
                                                                                    arg0.clone();
                                                                                move
                                                                                    |arg1|
                                                                                    &Func1::new({
                                                                                                    let arg1
                                                                                                        =
                                                                                                        arg1.clone();
                                                                                                    move
                                                                                                        |arg2|
                                                                                                        &Func1::new({
                                                                                                                        let arg2
                                                                                                                            =
                                                                                                                            arg2.clone();
                                                                                                                        move
                                                                                                                            |arg3|
                                                                                                                            &Func1::new({
                                                                                                                                            let arg3
                                                                                                                                                =
                                                                                                                                                arg3.clone();
                                                                                                                                            move
                                                                                                                                                |arg4|
                                                                                                                                                &PureScript_Data_Ord::Data_Ord_FFI::ordBooleanImpl(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&arg1),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&arg2),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&arg3),
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(arg4))
                                                                                                                                        })
                                                                                                                    })
                                                                                                })
                                                                            })))
    }
    pub fn Data_Ord_ordCharImpl() -> &dyn Any {
        static Data_Ord_ordCharImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordCharImpl.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &Func1::new({
                                                                             let arg0
                                                                                 =
                                                                                 arg0.clone();
                                                                             move
                                                                                 |arg1|
                                                                                 &Func1::new({
                                                                                                 let arg1
                                                                                                     =
                                                                                                     arg1.clone();
                                                                                                 move
                                                                                                     |arg2|
                                                                                                     &Func1::new({
                                                                                                                     let arg2
                                                                                                                         =
                                                                                                                         arg2.clone();
                                                                                                                     move
                                                                                                                         |arg3|
                                                                                                                         &Func1::new({
                                                                                                                                         let arg3
                                                                                                                                             =
                                                                                                                                             arg3.clone();
                                                                                                                                         move
                                                                                                                                             |arg4|
                                                                                                                                             &PureScript_Data_Ord::Data_Ord_FFI::ordCharImpl(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                                                             Sharpurs_Prelude::unbox(&arg1),
                                                                                                                                                                                             Sharpurs_Prelude::unbox(&arg2),
                                                                                                                                                                                             Sharpurs_Prelude::unbox(&arg3),
                                                                                                                                                                                             Sharpurs_Prelude::unbox(arg4))
                                                                                                                                     })
                                                                                                                 })
                                                                                             })
                                                                         })))
    }
    pub fn Data_Ord_ordIntImpl() -> &dyn Any {
        static Data_Ord_ordIntImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordIntImpl.get_or_init(||
                                            &Func1::new(move |arg0|
                                                            &Func1::new({
                                                                            let arg0
                                                                                =
                                                                                arg0.clone();
                                                                            move
                                                                                |arg1|
                                                                                &Func1::new({
                                                                                                let arg1
                                                                                                    =
                                                                                                    arg1.clone();
                                                                                                move
                                                                                                    |arg2|
                                                                                                    &Func1::new({
                                                                                                                    let arg2
                                                                                                                        =
                                                                                                                        arg2.clone();
                                                                                                                    move
                                                                                                                        |arg3|
                                                                                                                        &Func1::new({
                                                                                                                                        let arg3
                                                                                                                                            =
                                                                                                                                            arg3.clone();
                                                                                                                                        move
                                                                                                                                            |arg4|
                                                                                                                                            &PureScript_Data_Ord::Data_Ord_FFI::ordIntImpl(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&arg1),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&arg2),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(&arg3),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(arg4))
                                                                                                                                    })
                                                                                                                })
                                                                                            })
                                                                        })))
    }
    pub fn Data_Ord_ordNumberImpl() -> &dyn Any {
        static Data_Ord_ordNumberImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordNumberImpl.get_or_init(||
                                               &Func1::new(move |arg0|
                                                               &Func1::new({
                                                                               let arg0
                                                                                   =
                                                                                   arg0.clone();
                                                                               move
                                                                                   |arg1|
                                                                                   &Func1::new({
                                                                                                   let arg1
                                                                                                       =
                                                                                                       arg1.clone();
                                                                                                   move
                                                                                                       |arg2|
                                                                                                       &Func1::new({
                                                                                                                       let arg2
                                                                                                                           =
                                                                                                                           arg2.clone();
                                                                                                                       move
                                                                                                                           |arg3|
                                                                                                                           &Func1::new({
                                                                                                                                           let arg3
                                                                                                                                               =
                                                                                                                                               arg3.clone();
                                                                                                                                           move
                                                                                                                                               |arg4|
                                                                                                                                               &PureScript_Data_Ord::Data_Ord_FFI::ordNumberImpl(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&arg1),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&arg2),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&arg3),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(arg4))
                                                                                                                                       })
                                                                                                                   })
                                                                                               })
                                                                           })))
    }
    pub fn Data_Ord_ordStringImpl() -> &dyn Any {
        static Data_Ord_ordStringImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordStringImpl.get_or_init(||
                                               &Func1::new(move |arg0|
                                                               &Func1::new({
                                                                               let arg0
                                                                                   =
                                                                                   arg0.clone();
                                                                               move
                                                                                   |arg1|
                                                                                   &Func1::new({
                                                                                                   let arg1
                                                                                                       =
                                                                                                       arg1.clone();
                                                                                                   move
                                                                                                       |arg2|
                                                                                                       &Func1::new({
                                                                                                                       let arg2
                                                                                                                           =
                                                                                                                           arg2.clone();
                                                                                                                       move
                                                                                                                           |arg3|
                                                                                                                           &Func1::new({
                                                                                                                                           let arg3
                                                                                                                                               =
                                                                                                                                               arg3.clone();
                                                                                                                                           move
                                                                                                                                               |arg4|
                                                                                                                                               &PureScript_Data_Ord::Data_Ord_FFI::ordStringImpl(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&arg1),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&arg2),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&arg3),
                                                                                                                                                                                                 Sharpurs_Prelude::unbox(arg4))
                                                                                                                                       })
                                                                                                                   })
                                                                                               })
                                                                           })))
    }
    pub fn Data_Ord_eqRec() -> &dyn Any {
        static Data_Ord_eqRec: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Ord_eqRec.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eqRec(),
                                                                        &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Ord_OrdRecordusd_Dict() -> &dyn Any {
        static Data_Ord_OrdRecordusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_OrdRecordusd_Dict.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Ord_Ordusd_Dict() -> &dyn Any {
        static Data_Ord_Ordusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Ordusd_Dict.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Ord_Ord1usd_Dict() -> &dyn Any {
        static Data_Ord_Ord1usd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Ord1usd_Dict.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Ord_ordVoid() -> &dyn Any {
        static Data_Ord_ordVoid: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordVoid.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                          &&&add(string("compare"),
                                                                                 &&Func1::new(move
                                                                                                  |v|
                                                                                                  &Func1::new(move
                                                                                                                  |v1|
                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))),
                                                                                 add(string("Eq0"),
                                                                                     &&Func1::new(move
                                                                                                      |usd__unused|
                                                                                                      &PureScript_Data_Eq::Data_Eq_eqVoid()),
                                                                                     empty::<string,
                                                                                             &dyn Any>()))))
    }
    pub fn Data_Ord_ordUnit() -> &dyn Any {
        static Data_Ord_ordUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordUnit.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                          &&&add(string("compare"),
                                                                                 &&Func1::new(move
                                                                                                  |v|
                                                                                                  &Func1::new(move
                                                                                                                  |v1|
                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))),
                                                                                 add(string("Eq0"),
                                                                                     &&Func1::new(move
                                                                                                      |usd__unused|
                                                                                                      &PureScript_Data_Eq::Data_Eq_eqUnit()),
                                                                                     empty::<string,
                                                                                             &dyn Any>()))))
    }
    pub fn Data_Ord_ordString() -> &dyn Any {
        static Data_Ord_ordString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordString.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                            &&&add(string("compare"),
                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordStringImpl(),
                                                                                                                                                                                           &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)),
                                                                                                                                                        &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)),
                                                                                                                     &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)),
                                                                                   add(string("Eq0"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Data_Eq::Data_Eq_eqString()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Data_Ord_ordRecordNil() -> &dyn Any {
        static Data_Ord_ordRecordNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordRecordNil.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_OrdRecordusd_Dict(),
                                                                               &&&add(string("compareRecord"),
                                                                                      &&Func1::new(move
                                                                                                       |v|
                                                                                                       &Func1::new(move
                                                                                                                       |v1|
                                                                                                                       &Func1::new(move
                                                                                                                                       |v2|
                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)))),
                                                                                      add(string("EqRecord0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &PureScript_Data_Eq::Data_Eq_eqRowNil()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Ord_ordProxy() -> &dyn Any {
        static Data_Ord_ordProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordProxy.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                           &&&add(string("compare"),
                                                                                  &&Func1::new(move
                                                                                                   |v|
                                                                                                   &Func1::new(move
                                                                                                                   |v1|
                                                                                                                   &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))),
                                                                                  add(string("Eq0"),
                                                                                      &&Func1::new(move
                                                                                                       |usd__unused|
                                                                                                       &PureScript_Data_Eq::Data_Eq_eqProxy()),
                                                                                      empty::<string,
                                                                                              &dyn Any>()))))
    }
    pub fn Data_Ord_ordOrdering() -> &dyn Any {
        static Data_Ord_ordOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordOrdering.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                              &&&add(string("compare"),
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
                                                                                                                                      LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                                                              let matchValue_1:
                                                                                                                                      LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                  Sharpurs_Prelude::unbox(v1);
                                                                                                                              if let Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                                                     =
                                                                                                                                     matchValue.as_ref()
                                                                                                                                 {
                                                                                                                                  if let Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                         =
                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                     {
                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                  } else {
                                                                                                                                      if let Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                             =
                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                         {
                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                      } else {
                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                      }
                                                                                                                                  }
                                                                                                                              } else {
                                                                                                                                  if let Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                         =
                                                                                                                                         matchValue.as_ref()
                                                                                                                                     {
                                                                                                                                      if let Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                             =
                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                         {
                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                      } else {
                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                      }
                                                                                                                                  } else {
                                                                                                                                      if let Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                             =
                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                         {
                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                      } else {
                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                      }
                                                                                                                                  }
                                                                                                                              }
                                                                                                                          }
                                                                                                                  })),
                                                                                     add(string("Eq0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Ord_ordNumber() -> &dyn Any {
        static Data_Ord_ordNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordNumber.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                            &&&add(string("compare"),
                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordNumberImpl(),
                                                                                                                                                                                           &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)),
                                                                                                                                                        &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)),
                                                                                                                     &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)),
                                                                                   add(string("Eq0"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Data_Eq::Data_Eq_eqNumber()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Data_Ord_ordInt() -> &dyn Any {
        static Data_Ord_ordInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordInt.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                         &&&add(string("compare"),
                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordIntImpl(),
                                                                                                                                                                                        &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)),
                                                                                                                                                     &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)),
                                                                                                                  &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)),
                                                                                add(string("Eq0"),
                                                                                    &&Func1::new(move
                                                                                                     |usd__unused|
                                                                                                     &PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                    empty::<string,
                                                                                            &dyn Any>()))))
    }
    pub fn Data_Ord_ordChar() -> &dyn Any {
        static Data_Ord_ordChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordChar.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                          &&&add(string("compare"),
                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordCharImpl(),
                                                                                                                                                                                         &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)),
                                                                                                                                                      &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)),
                                                                                                                   &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)),
                                                                                 add(string("Eq0"),
                                                                                     &&Func1::new(move
                                                                                                      |usd__unused|
                                                                                                      &PureScript_Data_Eq::Data_Eq_eqChar()),
                                                                                     empty::<string,
                                                                                             &dyn Any>()))))
    }
    pub fn Data_Ord_ordBoolean() -> &dyn Any {
        static Data_Ord_ordBoolean: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordBoolean.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                             &&&add(string("compare"),
                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordBooleanImpl(),
                                                                                                                                                                                            &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)),
                                                                                                                                                         &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)),
                                                                                                                      &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)),
                                                                                    add(string("Eq0"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused|
                                                                                                         &PureScript_Data_Eq::Data_Eq_eqBoolean()),
                                                                                        empty::<string,
                                                                                                &dyn Any>()))))
    }
    pub fn Data_Ord_compareRecord() -> &dyn Any {
        static Data_Ord_compareRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_compareRecord.get_or_init(||
                                               &Func1::new(move |dict|
                                                               find(string("compareRecord"),
                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Ord_ordRecord() -> &dyn Any {
        static Data_Ord_ordRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordRecord.get_or_init(||
                                           &Func1::new(move |usd__unused|
                                                           &Func1::new(move
                                                                           |dictOrdRecord|
                                                                           {
                                                                               let eqRec1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_eqRec(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("EqRecord0"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictOrdRecord)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                &&&add(string("compare"),
                                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compareRecord(),
                                                                                                                                                                                            dictOrdRecord),
                                                                                                                                                         &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                       add(string("Eq0"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let eqRec1
                                                                                                                                                =
                                                                                                                                                eqRec1.clone();
                                                                                                                                            move
                                                                                                                                                |usd__unused_1|
                                                                                                                                                &eqRec1
                                                                                                                                        }),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>())))
                                                                           })))
    }
    pub fn Data_Ord_compare1() -> &dyn Any {
        static Data_Ord_compare1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_compare1.get_or_init(||
                                          &Func1::new(move |dict|
                                                          find(string("compare1"),
                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Ord_compare() -> &dyn Any {
        static Data_Ord_compare: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_compare.get_or_init(||
                                         &Func1::new(move |dict|
                                                         find(string("compare"),
                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Ord_comparing() -> &dyn Any {
        static Data_Ord_comparing: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_comparing.get_or_init(||
                                           &Func1::new(move |dictOrd|
                                                           &Func1::new({
                                                                           let dictOrd
                                                                               =
                                                                               dictOrd.clone();
                                                                           move
                                                                               |f|
                                                                               &Func1::new({
                                                                                               let f
                                                                                                   =
                                                                                                   f.clone();
                                                                                               move
                                                                                                   |x|
                                                                                                   &Func1::new({
                                                                                                                   let x
                                                                                                                       =
                                                                                                                       x.clone();
                                                                                                                   move
                                                                                                                       |y|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                              &&&dictOrd),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                              &&&x)),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                           y))
                                                                                                               })
                                                                                           })
                                                                       })))
    }
    pub fn Data_Ord_greaterThan() -> &dyn Any {
        static Data_Ord_greaterThan: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_greaterThan.get_or_init(||
                                             &Func1::new(move |dictOrd|
                                                             &Func1::new({
                                                                             let dictOrd
                                                                                 =
                                                                                 dictOrd.clone();
                                                                             move
                                                                                 |a1|
                                                                                 &Func1::new({
                                                                                                 let a1
                                                                                                     =
                                                                                                     a1.clone();
                                                                                                 move
                                                                                                     |a2|
                                                                                                     if let Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                            =
                                                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                             &&&dictOrd),
                                                                                                                                                                                                          &&&a1),
                                                                                                                                                                       a2)).as_ref()
                                                                                                        {
                                                                                                         &true
                                                                                                     } else {
                                                                                                         &false
                                                                                                     }
                                                                                             })
                                                                         })))
    }
    pub fn Data_Ord_greaterThanOrEq() -> &dyn Any {
        static Data_Ord_greaterThanOrEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_greaterThanOrEq.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 &Func1::new({
                                                                                 let dictOrd
                                                                                     =
                                                                                     dictOrd.clone();
                                                                                 move
                                                                                     |a1|
                                                                                     &Func1::new({
                                                                                                     let a1
                                                                                                         =
                                                                                                         a1.clone();
                                                                                                     move
                                                                                                         |a2|
                                                                                                         if let Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                =
                                                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                 &&&dictOrd),
                                                                                                                                                                                                              &&&a1),
                                                                                                                                                                           a2)).as_ref()
                                                                                                            {
                                                                                                             &false
                                                                                                         } else {
                                                                                                             &true
                                                                                                         }
                                                                                                 })
                                                                             })))
    }
    pub fn Data_Ord_lessThan() -> &dyn Any {
        static Data_Ord_lessThan: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_lessThan.get_or_init(||
                                          &Func1::new(move |dictOrd|
                                                          &Func1::new({
                                                                          let dictOrd
                                                                              =
                                                                              dictOrd.clone();
                                                                          move
                                                                              |a1|
                                                                              &Func1::new({
                                                                                              let a1
                                                                                                  =
                                                                                                  a1.clone();
                                                                                              move
                                                                                                  |a2|
                                                                                                  if let Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                         =
                                                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                          &&&dictOrd),
                                                                                                                                                                                                       &&&a1),
                                                                                                                                                                    a2)).as_ref()
                                                                                                     {
                                                                                                      &true
                                                                                                  } else {
                                                                                                      &false
                                                                                                  }
                                                                                          })
                                                                      })))
    }
    pub fn Data_Ord_signum() -> &dyn Any {
        static Data_Ord_signum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_signum.get_or_init(||
                                        &Func1::new(move |dictOrd|
                                                        &Func1::new({
                                                                        let dictOrd
                                                                            =
                                                                            dictOrd.clone();
                                                                        move
                                                                            |dictRing|
                                                                            {
                                                                                let Semiring0 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                            Sharpurs_Prelude::unbox(dictRing)),
                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                                let zero =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                     &&&Semiring0);
                                                                                let Semiring01 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                            Sharpurs_Prelude::unbox(dictRing)),
                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined());
                                                                                let one =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                     &&&Semiring0);
                                                                                &Func1::new({
                                                                                                let Semiring01
                                                                                                    =
                                                                                                    Semiring01.clone();
                                                                                                let dictRing
                                                                                                    =
                                                                                                    dictRing.clone();
                                                                                                let one
                                                                                                    =
                                                                                                    one.clone();
                                                                                                let zero
                                                                                                    =
                                                                                                    zero.clone();
                                                                                                move
                                                                                                    |x|
                                                                                                    {
                                                                                                        let matchValue =
                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                            &&&dictOrd),
                                                                                                                                                                                                         x),
                                                                                                                                                                      &&&zero));
                                                                                                        match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                         &matchValue)
                                                                                                            {
                                                                                                            0_i32
                                                                                                            =>
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                                &&&dictRing),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                                &&&Semiring01)),
                                                                                                            _
                                                                                                            =>
                                                                                                            {
                                                                                                                let matchValue_1 =
                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                                                    &&&dictOrd),
                                                                                                                                                                                                                 x),
                                                                                                                                                                              &&&zero));
                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                 &matchValue_1)
                                                                                                                    {
                                                                                                                    0_i32
                                                                                                                    =>
                                                                                                                    &one,
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    x.clone(),
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                            })
                                                                            }
                                                                    })))
    }
    pub fn Data_Ord_lessThanOrEq() -> &dyn Any {
        static Data_Ord_lessThanOrEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_lessThanOrEq.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              &Func1::new({
                                                                              let dictOrd
                                                                                  =
                                                                                  dictOrd.clone();
                                                                              move
                                                                                  |a1|
                                                                                  &Func1::new({
                                                                                                  let a1
                                                                                                      =
                                                                                                      a1.clone();
                                                                                                  move
                                                                                                      |a2|
                                                                                                      if let Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                             =
                                                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                              &&&dictOrd),
                                                                                                                                                                                                           &&&a1),
                                                                                                                                                                        a2)).as_ref()
                                                                                                         {
                                                                                                          &false
                                                                                                      } else {
                                                                                                          &true
                                                                                                      }
                                                                                              })
                                                                          })))
    }
    pub fn Data_Ord_max() -> &dyn Any {
        static Data_Ord_max: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Ord_max.get_or_init(||
                                     &Func1::new(move |dictOrd|
                                                     &Func1::new({
                                                                     let dictOrd
                                                                         =
                                                                         dictOrd.clone();
                                                                     move |x|
                                                                         &Func1::new({
                                                                                         let x
                                                                                             =
                                                                                             x.clone();
                                                                                         move
                                                                                             |y|
                                                                                             {
                                                                                                 let matchValue:
                                                                                                         LrcPtr<Data_Ordering_Ordering> =
                                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                      &&&dictOrd),
                                                                                                                                                                                                   &&&x),
                                                                                                                                                                y));
                                                                                                 match matchValue.as_ref()
                                                                                                     {
                                                                                                     Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                     =>
                                                                                                     &x,
                                                                                                     Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                     =>
                                                                                                     &x,
                                                                                                     _
                                                                                                     =>
                                                                                                     y.clone(),
                                                                                                 }
                                                                                             }
                                                                                     })
                                                                 })))
    }
    pub fn Data_Ord_min() -> &dyn Any {
        static Data_Ord_min: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Ord_min.get_or_init(||
                                     &Func1::new(move |dictOrd|
                                                     &Func1::new({
                                                                     let dictOrd
                                                                         =
                                                                         dictOrd.clone();
                                                                     move |x|
                                                                         &Func1::new({
                                                                                         let x
                                                                                             =
                                                                                             x.clone();
                                                                                         move
                                                                                             |y|
                                                                                             {
                                                                                                 let matchValue:
                                                                                                         LrcPtr<Data_Ordering_Ordering> =
                                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                      &&&dictOrd),
                                                                                                                                                                                                   &&&x),
                                                                                                                                                                y));
                                                                                                 match matchValue.as_ref()
                                                                                                     {
                                                                                                     Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor
                                                                                                     =>
                                                                                                     &x,
                                                                                                     Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                     =>
                                                                                                     y.clone(),
                                                                                                     _
                                                                                                     =>
                                                                                                     &x,
                                                                                                 }
                                                                                             }
                                                                                     })
                                                                 })))
    }
    pub fn Data_Ord_ordArray() -> &dyn Any {
        static Data_Ord_ordArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordArray.get_or_init(||
                                          &Func1::new(move |dictOrd|
                                                          {
                                                              let eqArray =
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eqArray(),
                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                             Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                               &&&add(string("compare"),
                                                                                                      &{
                                                                                                           let toDelta =
                                                                                                               &Func1::new({
                                                                                                                               let dictOrd
                                                                                                                                   =
                                                                                                                                   dictOrd.clone();
                                                                                                                               move
                                                                                                                                   |x|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let x
                                                                                                                                                       =
                                                                                                                                                       x.clone();
                                                                                                                                                   move
                                                                                                                                                       |y|
                                                                                                                                                       {
                                                                                                                                                           let matchValue:
                                                                                                                                                                   LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                &&&dictOrd),
                                                                                                                                                                                                                                                             &&&x),
                                                                                                                                                                                                                          y));
                                                                                                                                                           match matchValue.as_ref()
                                                                                                                                                               {
                                                                                                                                                               Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                               =>
                                                                                                                                                               &1_i32,
                                                                                                                                                               Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                                               =>
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                &&&1_i32),
                                                                                                                                                               _
                                                                                                                                                               =>
                                                                                                                                                               &0_i32,
                                                                                                                                                           }
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                           });
                                                                                                           &Func1::new({
                                                                                                                           let toDelta
                                                                                                                               =
                                                                                                                               toDelta.clone();
                                                                                                                           move
                                                                                                                               |xs|
                                                                                                                               &Func1::new({
                                                                                                                                               let xs
                                                                                                                                                   =
                                                                                                                                                   xs.clone();
                                                                                                                                               move
                                                                                                                                                   |ys|
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                          &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                       &&&0_i32),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordArrayImpl(),
                                                                                                                                                                                                                                                                                             &&&toDelta),
                                                                                                                                                                                                                                                          &&&xs),
                                                                                                                                                                                                                       ys))
                                                                                                                                           })
                                                                                                                       })
                                                                                                       },
                                                                                                      add(string("Eq0"),
                                                                                                          &&Func1::new({
                                                                                                                           let eqArray
                                                                                                                               =
                                                                                                                               eqArray.clone();
                                                                                                                           move
                                                                                                                               |usd__unused|
                                                                                                                               &eqArray
                                                                                                                       }),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
                                                          }))
    }
    pub fn Data_Ord_ord1Array() -> &dyn Any {
        static Data_Ord_ord1Array: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ord1Array.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                            &&&add(string("compare1"),
                                                                                   &&Func1::new(move
                                                                                                    |dictOrd|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordArray(),
                                                                                                                                                                        dictOrd))),
                                                                                   add(string("Eq10"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Data_Eq::Data_Eq_eq1Array()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Data_Ord_ordRecordCons() -> &dyn Any {
        static Data_Ord_ordRecordCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_ordRecordCons.get_or_init(||
                                               &Func1::new(move
                                                               |dictOrdRecord|
                                                               {
                                                                   let eqRowCons =
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eqRowCons(),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("EqRecord0"),
                                                                                                                                                                                     Sharpurs_Prelude::unbox(dictOrdRecord)),
                                                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                        &&&Sharpurs_Prelude::Prim_undefined());
                                                                   &Func1::new({
                                                                                   let dictOrdRecord
                                                                                       =
                                                                                       dictOrdRecord.clone();
                                                                                   let eqRowCons
                                                                                       =
                                                                                       eqRowCons.clone();
                                                                                   move
                                                                                       |usd__unused|
                                                                                       &Func1::new(move
                                                                                                       |dictIsSymbol|
                                                                                                       {
                                                                                                           let eqRowCons1 =
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&eqRowCons,
                                                                                                                                                dictIsSymbol);
                                                                                                           &Func1::new({
                                                                                                                           let dictIsSymbol
                                                                                                                               =
                                                                                                                               dictIsSymbol.clone();
                                                                                                                           let eqRowCons1
                                                                                                                               =
                                                                                                                               eqRowCons1.clone();
                                                                                                                           move
                                                                                                                               |dictOrd|
                                                                                                                               {
                                                                                                                                   let eqRowCons2 =
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&eqRowCons1,
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_OrdRecordusd_Dict(),
                                                                                                                                                                    &&&add(string("compareRecord"),
                                                                                                                                                                           &&Func1::new({
                                                                                                                                                                                            let dictOrd
                                                                                                                                                                                                =
                                                                                                                                                                                                dictOrd.clone();
                                                                                                                                                                                            move
                                                                                                                                                                                                |v|
                                                                                                                                                                                                &Func1::new(move
                                                                                                                                                                                                                |ra|
                                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                                let ra
                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                    ra.clone();
                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                    |rb|
                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                        let unsafeGet_prime =
                                                                                                                                                                                                                                            &PureScript_Record_Unsafe::Record_Unsafe_unsafeGet();
                                                                                                                                                                                                                                        let key =
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                &&&dictIsSymbol),
                                                                                                                                                                                                                                                                             &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                        let left =
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictOrd),
                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&unsafeGet_prime,
                                                                                                                                                                                                                                                                                                                                                                                      &&&key),
                                                                                                                                                                                                                                                                                                                                                   &&&ra)),
                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&unsafeGet_prime,
                                                                                                                                                                                                                                                                                                                                                   &&&key),
                                                                                                                                                                                                                                                                                                                rb));
                                                                                                                                                                                                                                        let matchValue =
                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_notEq(),
                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                                                                                                                         &&&left),
                                                                                                                                                                                                                                                                                                      &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)));
                                                                                                                                                                                                                                        match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                                         &matchValue)
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            0_i32
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            &left,
                                                                                                                                                                                                                                            _
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compareRecord(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&dictOrdRecord),
                                                                                                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                                                &&&ra),
                                                                                                                                                                                                                                                                             rb),
                                                                                                                                                                                                                                        }
                                                                                                                                                                                                                                    }
                                                                                                                                                                                                                            }))
                                                                                                                                                                                        }),
                                                                                                                                                                           add(string("EqRecord0"),
                                                                                                                                                                               &&Func1::new({
                                                                                                                                                                                                let eqRowCons2
                                                                                                                                                                                                    =
                                                                                                                                                                                                    eqRowCons2.clone();
                                                                                                                                                                                                move
                                                                                                                                                                                                    |usd__unused_1|
                                                                                                                                                                                                    &eqRowCons2
                                                                                                                                                                                            }),
                                                                                                                                                                               empty::<string,
                                                                                                                                                                                       &dyn Any>())))
                                                                                                                               }
                                                                                                                       })
                                                                                                       })
                                                                               })
                                                               }))
    }
    pub fn Data_Ord_clamp() -> &dyn Any {
        static Data_Ord_clamp: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Ord_clamp.get_or_init(||
                                       &Func1::new(move |dictOrd|
                                                       &Func1::new({
                                                                       let dictOrd
                                                                           =
                                                                           dictOrd.clone();
                                                                       move
                                                                           |low|
                                                                           &Func1::new({
                                                                                           let low
                                                                                               =
                                                                                               low.clone();
                                                                                           move
                                                                                               |hi|
                                                                                               &Func1::new({
                                                                                                               let hi
                                                                                                                   =
                                                                                                                   hi.clone();
                                                                                                               move
                                                                                                                   |x|
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_min(),
                                                                                                                                                                                                                          &&&dictOrd),
                                                                                                                                                                                       &&&hi),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_max(),
                                                                                                                                                                                                                                                             &&&dictOrd),
                                                                                                                                                                                                                          &&&low),
                                                                                                                                                                                       x))
                                                                                                           })
                                                                                       })
                                                                   })))
    }
    pub fn Data_Ord_between() -> &dyn Any {
        static Data_Ord_between: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_between.get_or_init(||
                                         &Func1::new(move |dictOrd|
                                                         &Func1::new({
                                                                         let dictOrd
                                                                             =
                                                                             dictOrd.clone();
                                                                         move
                                                                             |low|
                                                                             &Func1::new({
                                                                                             let low
                                                                                                 =
                                                                                                 low.clone();
                                                                                             move
                                                                                                 |hi|
                                                                                                 &Func1::new({
                                                                                                                 let hi
                                                                                                                     =
                                                                                                                     hi.clone();
                                                                                                                 move
                                                                                                                     |x|
                                                                                                                     {
                                                                                                                         let matchValue =
                                                                                                                             Sharpurs_Prelude::unbox(&&low);
                                                                                                                         let matchValue_1 =
                                                                                                                             Sharpurs_Prelude::unbox(&&hi);
                                                                                                                         let matchValue_2 =
                                                                                                                             Sharpurs_Prelude::unbox(x);
                                                                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                            &&&dictOrd),
                                                                                                                                                                                                                         &&&matchValue_2),
                                                                                                                                                                                      &&&matchValue))
                                                                                                                            {
                                                                                                                             &false
                                                                                                                         } else {
                                                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                                                                &&&dictOrd),
                                                                                                                                                                                                                             &&&matchValue_2),
                                                                                                                                                                                          &&&matchValue_1))
                                                                                                                                {
                                                                                                                                 &false
                                                                                                                             } else {
                                                                                                                                 if Sharpurs_Prelude::unbox(&&true)
                                                                                                                                    {
                                                                                                                                     &true
                                                                                                                                 } else {
                                                                                                                                     panic!("{}",
                                                                                                                                            LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Ord.fs"),
                                  Data1: 84_i32,
                                  Data2: 128_i32,}).get_Message(),)
                                                                                                                                 }
                                                                                                                             }
                                                                                                                         }
                                                                                                                     }
                                                                                                             })
                                                                                         })
                                                                     })))
    }
    pub fn Data_Ord_abs() -> &dyn Any {
        static Data_Ord_abs: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Ord_abs.get_or_init(||
                                     &Func1::new(move |dictOrd|
                                                     &Func1::new({
                                                                     let dictOrd
                                                                         =
                                                                         dictOrd.clone();
                                                                     move
                                                                         |dictRing|
                                                                         {
                                                                             let zero =
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                                                            Sharpurs_Prelude::unbox(dictRing)),
                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                             &Func1::new({
                                                                                             let dictRing
                                                                                                 =
                                                                                                 dictRing.clone();
                                                                                             let zero
                                                                                                 =
                                                                                                 zero.clone();
                                                                                             move
                                                                                                 |x|
                                                                                                 {
                                                                                                     let matchValue =
                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                         &&&dictOrd),
                                                                                                                                                                                                      x),
                                                                                                                                                                   &&&zero));
                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                      &matchValue)
                                                                                                         {
                                                                                                         0_i32
                                                                                                         =>
                                                                                                         x.clone(),
                                                                                                         _
                                                                                                         =>
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                             &&&dictRing),
                                                                                                                                          x),
                                                                                                     }
                                                                                                 }
                                                                                         })
                                                                         }
                                                                 })))
    }
}
