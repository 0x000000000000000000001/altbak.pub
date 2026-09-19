pub mod PureScript_Data_Eq {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Data_Eq_FFI {
        use super::*;
        use fable_library_rust::NativeArray_::count;
        pub fn eqIntImpl(a: &dyn Any, b: &dyn Any) -> bool {
            Sharpurs_Prelude::unbox(a) == Sharpurs_Prelude::unbox(b)
        }
        pub fn eqNumberImpl(a: &dyn Any, b: &dyn Any) -> bool {
            Sharpurs_Prelude::unbox(a) == Sharpurs_Prelude::unbox(b)
        }
        pub fn eqStringImpl(a: &dyn Any, b: &dyn Any) -> bool {
            Sharpurs_Prelude::unbox(a) == Sharpurs_Prelude::unbox(b)
        }
        pub fn eqCharImpl(a: &dyn Any, b: &dyn Any) -> bool {
            Sharpurs_Prelude::unbox(a) == Sharpurs_Prelude::unbox(b)
        }
        pub fn eqBooleanImpl(a: &dyn Any, b: &dyn Any) -> bool {
            Sharpurs_Prelude::unbox(a) == Sharpurs_Prelude::unbox(b)
        }
        pub fn eqArrayImpl(f: &dyn Any, xs: &dyn Any, ys: &dyn Any)
         -> &dyn Any {
            let xs_ = Sharpurs_Prelude::unbox(xs);
            let ys_ = Sharpurs_Prelude::unbox(ys);
            if count(xs_.clone()) != count(ys_.clone()) {
                &false
            } else {
                let eq: MutCell<bool> = MutCell::new(true);
                for i in 0_i32..=count(xs_.clone()) - 1_i32 {
                    if !Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                                                    &xs_[i].clone()),
                                                                                  &ys_[i].clone()))
                       {
                        eq.set(false);
                    };
                }
                &eq.get()
            }
        }
    }
    pub fn Data_Eq_eqArrayImpl() -> &dyn Any {
        static Data_Eq_eqArrayImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqArrayImpl.get_or_init(||
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
                                                                                                    &PureScript_Data_Eq::Data_Eq_FFI::eqArrayImpl(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                  &Sharpurs_Prelude::unbox(&arg1),
                                                                                                                                                  &Sharpurs_Prelude::unbox(arg2))
                                                                                            })
                                                                        })))
    }
    pub fn Data_Eq_eqBooleanImpl() -> &dyn Any {
        static Data_Eq_eqBooleanImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqBooleanImpl.get_or_init(||
                                              &Func1::new(move |arg0|
                                                              &Func1::new({
                                                                              let arg0
                                                                                  =
                                                                                  arg0.clone();
                                                                              move
                                                                                  |arg1|
                                                                                  &PureScript_Data_Eq::Data_Eq_FFI::eqBooleanImpl(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                  &Sharpurs_Prelude::unbox(arg1))
                                                                          })))
    }
    pub fn Data_Eq_eqCharImpl() -> &dyn Any {
        static Data_Eq_eqCharImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqCharImpl.get_or_init(||
                                           &Func1::new(move |arg0|
                                                           &Func1::new({
                                                                           let arg0
                                                                               =
                                                                               arg0.clone();
                                                                           move
                                                                               |arg1|
                                                                               &PureScript_Data_Eq::Data_Eq_FFI::eqCharImpl(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                            &Sharpurs_Prelude::unbox(arg1))
                                                                       })))
    }
    pub fn Data_Eq_eqIntImpl() -> &dyn Any {
        static Data_Eq_eqIntImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqIntImpl.get_or_init(||
                                          &Func1::new(move |arg0|
                                                          &Func1::new({
                                                                          let arg0
                                                                              =
                                                                              arg0.clone();
                                                                          move
                                                                              |arg1|
                                                                              &PureScript_Data_Eq::Data_Eq_FFI::eqIntImpl(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                          &Sharpurs_Prelude::unbox(arg1))
                                                                      })))
    }
    pub fn Data_Eq_eqNumberImpl() -> &dyn Any {
        static Data_Eq_eqNumberImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqNumberImpl.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &Func1::new({
                                                                             let arg0
                                                                                 =
                                                                                 arg0.clone();
                                                                             move
                                                                                 |arg1|
                                                                                 &PureScript_Data_Eq::Data_Eq_FFI::eqNumberImpl(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                &Sharpurs_Prelude::unbox(arg1))
                                                                         })))
    }
    pub fn Data_Eq_eqStringImpl() -> &dyn Any {
        static Data_Eq_eqStringImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqStringImpl.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &Func1::new({
                                                                             let arg0
                                                                                 =
                                                                                 arg0.clone();
                                                                             move
                                                                                 |arg1|
                                                                                 &PureScript_Data_Eq::Data_Eq_FFI::eqStringImpl(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                &Sharpurs_Prelude::unbox(arg1))
                                                                         })))
    }
    pub fn Data_Eq_EqRecordusd_Dict() -> &dyn Any {
        static Data_Eq_EqRecordusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_EqRecordusd_Dict.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Eq_Equsd_Dict() -> &dyn Any {
        static Data_Eq_Equsd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Equsd_Dict.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Eq_Eq1usd_Dict() -> &dyn Any {
        static Data_Eq_Eq1usd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_Eq1usd_Dict.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Eq_eqVoid() -> &dyn Any {
        static Data_Eq_eqVoid: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_eqVoid.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                        &&&add(string("eq"),
                                                                               &&Func1::new(move
                                                                                                |v|
                                                                                                &Func1::new(move
                                                                                                                |v1|
                                                                                                                &true)),
                                                                               empty::<string,
                                                                                       &dyn Any>())))
    }
    pub fn Data_Eq_eqUnit() -> &dyn Any {
        static Data_Eq_eqUnit: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_eqUnit.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                        &&&add(string("eq"),
                                                                               &&Func1::new(move
                                                                                                |v|
                                                                                                &Func1::new(move
                                                                                                                |v1|
                                                                                                                &true)),
                                                                               empty::<string,
                                                                                       &dyn Any>())))
    }
    pub fn Data_Eq_eqString() -> &dyn Any {
        static Data_Eq_eqString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqString.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                          &&&add(string("eq"),
                                                                                 &&PureScript_Data_Eq::Data_Eq_eqStringImpl(),
                                                                                 empty::<string,
                                                                                         &dyn Any>())))
    }
    pub fn Data_Eq_eqRowNil() -> &dyn Any {
        static Data_Eq_eqRowNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqRowNil.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_EqRecordusd_Dict(),
                                                                          &&&add(string("eqRecord"),
                                                                                 &&Func1::new(move
                                                                                                  |v|
                                                                                                  &Func1::new(move
                                                                                                                  |v1|
                                                                                                                  &Func1::new(move
                                                                                                                                  |v2|
                                                                                                                                  &true))),
                                                                                 empty::<string,
                                                                                         &dyn Any>())))
    }
    pub fn Data_Eq_eqRecord() -> &dyn Any {
        static Data_Eq_eqRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqRecord.get_or_init(||
                                         &Func1::new(move |dict|
                                                         find(string("eqRecord"),
                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Eq_eqRec() -> &dyn Any {
        static Data_Eq_eqRec: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_eqRec.get_or_init(||
                                      &Func1::new(move |usd__unused|
                                                      &Func1::new(move
                                                                      |dictEqRecord|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                       &&&add(string("eq"),
                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eqRecord(),
                                                                                                                                                                                   dictEqRecord),
                                                                                                                                                &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))))
    }
    pub fn Data_Eq_eqProxy() -> &dyn Any {
        static Data_Eq_eqProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqProxy.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                         &&&add(string("eq"),
                                                                                &&Func1::new(move
                                                                                                 |v|
                                                                                                 &Func1::new(move
                                                                                                                 |v1|
                                                                                                                 &true)),
                                                                                empty::<string,
                                                                                        &dyn Any>())))
    }
    pub fn Data_Eq_eqNumber() -> &dyn Any {
        static Data_Eq_eqNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqNumber.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                          &&&add(string("eq"),
                                                                                 &&PureScript_Data_Eq::Data_Eq_eqNumberImpl(),
                                                                                 empty::<string,
                                                                                         &dyn Any>())))
    }
    pub fn Data_Eq_eqInt() -> &dyn Any {
        static Data_Eq_eqInt: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_eqInt.get_or_init(||
                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                       &&&add(string("eq"),
                                                                              &&PureScript_Data_Eq::Data_Eq_eqIntImpl(),
                                                                              empty::<string,
                                                                                      &dyn Any>())))
    }
    pub fn Data_Eq_eqChar() -> &dyn Any {
        static Data_Eq_eqChar: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_eqChar.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                        &&&add(string("eq"),
                                                                               &&PureScript_Data_Eq::Data_Eq_eqCharImpl(),
                                                                               empty::<string,
                                                                                       &dyn Any>())))
    }
    pub fn Data_Eq_eqBoolean() -> &dyn Any {
        static Data_Eq_eqBoolean: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqBoolean.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                           &&&add(string("eq"),
                                                                                  &&PureScript_Data_Eq::Data_Eq_eqBooleanImpl(),
                                                                                  empty::<string,
                                                                                          &dyn Any>())))
    }
    pub fn Data_Eq_eq1() -> &dyn Any {
        static Data_Eq_eq1: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_eq1.get_or_init(||
                                    &Func1::new(move |dict|
                                                    find(string("eq1"),
                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Eq_eq() -> &dyn Any {
        static Data_Eq_eq: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_eq.get_or_init(||
                                   &Func1::new(move |dict|
                                                   find(string("eq"),
                                                        Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Eq_eqArray() -> &dyn Any {
        static Data_Eq_eqArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqArray.get_or_init(||
                                        &Func1::new(move |dictEq|
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                         &&&add(string("eq"),
                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eqArrayImpl(),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                     dictEq)),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Eq_eq1Array() -> &dyn Any {
        static Data_Eq_eq1Array: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eq1Array.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                          &&&add(string("eq1"),
                                                                                 &&Func1::new(move
                                                                                                  |dictEq|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eqArray(),
                                                                                                                                                                      dictEq))),
                                                                                 empty::<string,
                                                                                         &dyn Any>())))
    }
    pub fn Data_Eq_eqRowCons() -> &dyn Any {
        static Data_Eq_eqRowCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Eq_eqRowCons.get_or_init(||
                                          &Func1::new(move |dictEqRecord|
                                                          &Func1::new({
                                                                          let dictEqRecord
                                                                              =
                                                                              dictEqRecord.clone();
                                                                          move
                                                                              |usd__unused|
                                                                              &Func1::new(move
                                                                                              |dictIsSymbol|
                                                                                              &Func1::new({
                                                                                                              let dictIsSymbol
                                                                                                                  =
                                                                                                                  dictIsSymbol.clone();
                                                                                                              move
                                                                                                                  |dictEq|
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_EqRecordusd_Dict(),
                                                                                                                                                   &&&add(string("eqRecord"),
                                                                                                                                                          &&Func1::new({
                                                                                                                                                                           let dictEq
                                                                                                                                                                               =
                                                                                                                                                                               dictEq.clone();
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
                                                                                                                                                                                                                       let tail =
                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eqRecord(),
                                                                                                                                                                                                                                                                                                                                                                     &&&dictEqRecord),
                                                                                                                                                                                                                                                                                                                                  &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                               &&&ra),
                                                                                                                                                                                                                                                            rb);
                                                                                                                                                                                                                       let key =
                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                               &&&dictIsSymbol),
                                                                                                                                                                                                                                                            &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                       let get_ =
                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeGet(),
                                                                                                                                                                                                                                                            &&&key);
                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictEq),
                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&get_,
                                                                                                                                                                                                                                                                                                                                                                                                    &&&ra)),
                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&get_,
                                                                                                                                                                                                                                                                                                                                                                 rb))),
                                                                                                                                                                                                                                                        &&&tail)
                                                                                                                                                                                                                   }
                                                                                                                                                                                                           }))
                                                                                                                                                                       }),
                                                                                                                                                          empty::<string,
                                                                                                                                                                  &dyn Any>()))
                                                                                                          }))
                                                                      })))
    }
    pub fn Data_Eq_notEq() -> &dyn Any {
        static Data_Eq_notEq: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_notEq.get_or_init(||
                                      &Func1::new(move |dictEq|
                                                      &Func1::new({
                                                                      let dictEq
                                                                          =
                                                                          dictEq.clone();
                                                                      move |x|
                                                                          &Func1::new({
                                                                                          let x
                                                                                              =
                                                                                              x.clone();
                                                                                          move
                                                                                              |y|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                     &&&PureScript_Data_Eq::Data_Eq_eqBoolean()),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                           &&&dictEq),
                                                                                                                                                                                                                                        &&&x),
                                                                                                                                                                                                     y)),
                                                                                                                               &&&false)
                                                                                      })
                                                                  })))
    }
    pub fn Data_Eq_notEq1() -> &dyn Any {
        static Data_Eq_notEq1: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Eq_notEq1.get_or_init(||
                                       &Func1::new(move |dictEq1|
                                                       &Func1::new({
                                                                       let dictEq1
                                                                           =
                                                                           dictEq1.clone();
                                                                       move
                                                                           |dictEq|
                                                                           &Func1::new({
                                                                                           let dictEq
                                                                                               =
                                                                                               dictEq.clone();
                                                                                           move
                                                                                               |x|
                                                                                               &Func1::new({
                                                                                                               let x
                                                                                                                   =
                                                                                                                   x.clone();
                                                                                                               move
                                                                                                                   |y|
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                          &&&PureScript_Data_Eq::Data_Eq_eqBoolean()),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                                   &&&dictEq1),
                                                                                                                                                                                                                                                                                                &&&dictEq),
                                                                                                                                                                                                                                                             &&&x),
                                                                                                                                                                                                                          y)),
                                                                                                                                                    &&&false)
                                                                                                           })
                                                                                       })
                                                                   })))
    }
}
