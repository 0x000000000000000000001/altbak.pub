pub mod PureScript_Data_Ring {
    use super::*;
    use fable_library_rust::Interfaces_::System::IComparable;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Data_Ring_FFI {
        use super::*;
        pub fn intSub<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> i32 {
            Sharpurs_Prelude::unbox(&&a) - Sharpurs_Prelude::unbox(&&b)
        }
        pub fn numSub<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> f64 {
            Sharpurs_Prelude::unbox(&&a) - Sharpurs_Prelude::unbox(&&b)
        }
    }
    pub fn Data_Ring_intSub() -> &dyn Any {
        static Data_Ring_intSub: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_intSub.get_or_init(||
                                         &Func1::new(move |arg0|
                                                         &Func1::new({
                                                                         let arg0
                                                                             =
                                                                             arg0.clone();
                                                                         move
                                                                             |arg1|
                                                                             &PureScript_Data_Ring::Data_Ring_FFI::intSub(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                          Sharpurs_Prelude::unbox(arg1))
                                                                     })))
    }
    pub fn Data_Ring_numSub() -> &dyn Any {
        static Data_Ring_numSub: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_numSub.get_or_init(||
                                         &Func1::new(move |arg0|
                                                         &Func1::new({
                                                                         let arg0
                                                                             =
                                                                             arg0.clone();
                                                                         move
                                                                             |arg1|
                                                                             &PureScript_Data_Ring::Data_Ring_FFI::numSub(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                          Sharpurs_Prelude::unbox(arg1))
                                                                     })))
    }
    pub fn Data_Ring_semiringRecord() -> &dyn Any {
        static Data_Ring_semiringRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_semiringRecord.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_semiringRecord(),
                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Ring_RingRecordusd_Dict() -> &dyn Any {
        static Data_Ring_RingRecordusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_RingRecordusd_Dict.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Data_Ring_Ringusd_Dict() -> &dyn Any {
        static Data_Ring_Ringusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_Ringusd_Dict.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Ring_subRecord() -> &dyn Any {
        static Data_Ring_subRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_subRecord.get_or_init(||
                                            &Func1::new(move |dict|
                                                            find(string("subRecord"),
                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Ring_sub() -> &dyn Any {
        static Data_Ring_sub: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Ring_sub.get_or_init(||
                                      &Func1::new(move |dict|
                                                      find(string("sub"),
                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Ring_ringUnit() -> &dyn Any {
        static Data_Ring_ringUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_ringUnit.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                            &&&add(string("sub"),
                                                                                   &&Func1::new(move
                                                                                                    |v|
                                                                                                    &Func1::new(move
                                                                                                                    |v1|
                                                                                                                    &PureScript_Data_Unit::Data_Unit_unit())),
                                                                                   add(string("Semiring0"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Data_Semiring::Data_Semiring_semiringUnit()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Data_Ring_ringRecordNil() -> &dyn Any {
        static Data_Ring_ringRecordNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_ringRecordNil.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_RingRecordusd_Dict(),
                                                                                 &&&add(string("subRecord"),
                                                                                        &&Func1::new(move
                                                                                                         |v|
                                                                                                         &Func1::new(move
                                                                                                                         |v1|
                                                                                                                         &Func1::new(move
                                                                                                                                         |v2|
                                                                                                                                         &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                  &dyn Any>()))),
                                                                                        add(string("SemiringRecord0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &PureScript_Data_Semiring::Data_Semiring_semiringRecordNil()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Ring_ringRecordCons() -> &dyn Any {
        static Data_Ring_ringRecordCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_ringRecordCons.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictIsSymbol|
                                                                 {
                                                                     let semiringRecordCons =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_semiringRecordCons(),
                                                                                                                                             dictIsSymbol),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     &Func1::new({
                                                                                     let dictIsSymbol
                                                                                         =
                                                                                         dictIsSymbol.clone();
                                                                                     let semiringRecordCons
                                                                                         =
                                                                                         semiringRecordCons.clone();
                                                                                     move
                                                                                         |usd__unused|
                                                                                         &Func1::new(move
                                                                                                         |dictRingRecord|
                                                                                                         {
                                                                                                             let semiringRecordCons1 =
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&semiringRecordCons,
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("SemiringRecord0"),
                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictRingRecord)),
                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                             &Func1::new({
                                                                                                                             let dictRingRecord
                                                                                                                                 =
                                                                                                                                 dictRingRecord.clone();
                                                                                                                             let semiringRecordCons1
                                                                                                                                 =
                                                                                                                                 semiringRecordCons1.clone();
                                                                                                                             move
                                                                                                                                 |dictRing|
                                                                                                                                 {
                                                                                                                                     let semiringRecordCons2 =
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&semiringRecordCons1,
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictRing)),
                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_RingRecordusd_Dict(),
                                                                                                                                                                      &&&add(string("subRecord"),
                                                                                                                                                                             &&Func1::new({
                                                                                                                                                                                              let dictRing
                                                                                                                                                                                                  =
                                                                                                                                                                                                  dictRing.clone();
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
                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_subRecord(),
                                                                                                                                                                                                                                                                                                                                                                                        &&&dictRingRecord),
                                                                                                                                                                                                                                                                                                                                                     &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                                                  &&&ra),
                                                                                                                                                                                                                                                                               rb);
                                                                                                                                                                                                                                          let key =
                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                  &&&dictIsSymbol),
                                                                                                                                                                                                                                                                               &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                          let insert =
                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                               &&&key);
                                                                                                                                                                                                                                          let get_ =
                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeGet(),
                                                                                                                                                                                                                                                                               &&&key);
                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&insert,
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&dictRing),
                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&get_,
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&ra)),
                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&get_,
                                                                                                                                                                                                                                                                                                                                                                                    rb))),
                                                                                                                                                                                                                                                                           &&&tail)
                                                                                                                                                                                                                                      }
                                                                                                                                                                                                                              }))
                                                                                                                                                                                          }),
                                                                                                                                                                             add(string("SemiringRecord0"),
                                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                                  let semiringRecordCons2
                                                                                                                                                                                                      =
                                                                                                                                                                                                      semiringRecordCons2.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |usd__unused_1|
                                                                                                                                                                                                      &semiringRecordCons2
                                                                                                                                                                                              }),
                                                                                                                                                                                 empty::<string,
                                                                                                                                                                                         &dyn Any>())))
                                                                                                                                 }
                                                                                                                         })
                                                                                                         })
                                                                                 })
                                                                 }))
    }
    pub fn Data_Ring_ringRecord() -> &dyn Any {
        static Data_Ring_ringRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_ringRecord.get_or_init(||
                                             &Func1::new(move |usd__unused|
                                                             &Func1::new(move
                                                                             |dictRingRecord|
                                                                             {
                                                                                 let semiringRecord1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_semiringRecord(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("SemiringRecord0"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictRingRecord)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                                                                  &&&add(string("sub"),
                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_subRecord(),
                                                                                                                                                                                              dictRingRecord),
                                                                                                                                                           &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                         add(string("Semiring0"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let semiringRecord1
                                                                                                                                                  =
                                                                                                                                                  semiringRecord1.clone();
                                                                                                                                              move
                                                                                                                                                  |usd__unused_1|
                                                                                                                                                  &semiringRecord1
                                                                                                                                          }),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>())))
                                                                             })))
    }
    pub fn Data_Ring_ringProxy() -> &dyn Any {
        static Data_Ring_ringProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_ringProxy.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                             &&&add(string("sub"),
                                                                                    &&Func1::new(move
                                                                                                     |v|
                                                                                                     &Func1::new(move
                                                                                                                     |v1|
                                                                                                                     &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                    add(string("Semiring0"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused|
                                                                                                         &PureScript_Data_Semiring::Data_Semiring_semiringProxy()),
                                                                                        empty::<string,
                                                                                                &dyn Any>()))))
    }
    pub fn Data_Ring_ringNumber() -> &dyn Any {
        static Data_Ring_ringNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_ringNumber.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                              &&&add(string("sub"),
                                                                                     &&PureScript_Data_Ring::Data_Ring_numSub(),
                                                                                     add(string("Semiring0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Ring_ringInt() -> &dyn Any {
        static Data_Ring_ringInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_ringInt.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                           &&&add(string("sub"),
                                                                                  &&PureScript_Data_Ring::Data_Ring_intSub(),
                                                                                  add(string("Semiring0"),
                                                                                      &&Func1::new(move
                                                                                                       |usd__unused|
                                                                                                       &PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                      empty::<string,
                                                                                              &dyn Any>()))))
    }
    pub fn Data_Ring_ringFn() -> &dyn Any {
        static Data_Ring_ringFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_ringFn.get_or_init(||
                                         &Func1::new(move |dictRing|
                                                         {
                                                             let semiringFn =
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_semiringFn(),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                                            Sharpurs_Prelude::unbox(dictRing)),
                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()));
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                                              &&&add(string("sub"),
                                                                                                     &&Func1::new({
                                                                                                                      let dictRing
                                                                                                                          =
                                                                                                                          dictRing.clone();
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
                                                                                                                                                                  |x|
                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                         &&&dictRing),
                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                         x)),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                      x))
                                                                                                                                                          })
                                                                                                                                      })
                                                                                                                  }),
                                                                                                     add(string("Semiring0"),
                                                                                                         &&Func1::new({
                                                                                                                          let semiringFn
                                                                                                                              =
                                                                                                                              semiringFn.clone();
                                                                                                                          move
                                                                                                                              |usd__unused|
                                                                                                                              &semiringFn
                                                                                                                      }),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
                                                         }))
    }
    pub fn Data_Ring_negate() -> &dyn Any {
        static Data_Ring_negate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ring_negate.get_or_init(||
                                         &Func1::new(move |dictRing|
                                                         {
                                                             let Semiring0 =
                                                                 Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                         Sharpurs_Prelude::unbox(dictRing)),
                                                                                                  &&&Sharpurs_Prelude::Prim_undefined());
                                                             &Func1::new({
                                                                             let Semiring0
                                                                                 =
                                                                                 Semiring0.clone();
                                                                             let dictRing
                                                                                 =
                                                                                 dictRing.clone();
                                                                             move
                                                                                 |a|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                        &&&dictRing),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                                                        &&&Semiring0)),
                                                                                                                  a)
                                                                         })
                                                         }))
    }
}
