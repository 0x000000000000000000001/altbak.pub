pub mod PureScript_Data_Semiring {
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
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Data_Semiring_FFI {
        use super::*;
        pub fn intAdd<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> i32 {
            Sharpurs_Prelude::unbox(&&a) + Sharpurs_Prelude::unbox(&&b)
        }
        pub fn intMul<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> i32 {
            Sharpurs_Prelude::unbox(&&a) * Sharpurs_Prelude::unbox(&&b)
        }
        pub fn numAdd<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> f64 {
            Sharpurs_Prelude::unbox(&&a) + Sharpurs_Prelude::unbox(&&b)
        }
        pub fn numMul<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> f64 {
            Sharpurs_Prelude::unbox(&&a) * Sharpurs_Prelude::unbox(&&b)
        }
    }
    pub fn Data_Semiring_intAdd() -> &dyn Any {
        static Data_Semiring_intAdd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_intAdd.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &Func1::new({
                                                                             let arg0
                                                                                 =
                                                                                 arg0.clone();
                                                                             move
                                                                                 |arg1|
                                                                                 &PureScript_Data_Semiring::Data_Semiring_FFI::intAdd(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                      Sharpurs_Prelude::unbox(arg1))
                                                                         })))
    }
    pub fn Data_Semiring_intMul() -> &dyn Any {
        static Data_Semiring_intMul: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_intMul.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &Func1::new({
                                                                             let arg0
                                                                                 =
                                                                                 arg0.clone();
                                                                             move
                                                                                 |arg1|
                                                                                 &PureScript_Data_Semiring::Data_Semiring_FFI::intMul(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                      Sharpurs_Prelude::unbox(arg1))
                                                                         })))
    }
    pub fn Data_Semiring_numAdd() -> &dyn Any {
        static Data_Semiring_numAdd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_numAdd.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &Func1::new({
                                                                             let arg0
                                                                                 =
                                                                                 arg0.clone();
                                                                             move
                                                                                 |arg1|
                                                                                 &PureScript_Data_Semiring::Data_Semiring_FFI::numAdd(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                      Sharpurs_Prelude::unbox(arg1))
                                                                         })))
    }
    pub fn Data_Semiring_numMul() -> &dyn Any {
        static Data_Semiring_numMul: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_numMul.get_or_init(||
                                             &Func1::new(move |arg0|
                                                             &Func1::new({
                                                                             let arg0
                                                                                 =
                                                                                 arg0.clone();
                                                                             move
                                                                                 |arg1|
                                                                                 &PureScript_Data_Semiring::Data_Semiring_FFI::numMul(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                      Sharpurs_Prelude::unbox(arg1))
                                                                         })))
    }
    pub fn Data_Semiring_SemiringRecordusd_Dict() -> &dyn Any {
        static Data_Semiring_SemiringRecordusd_Dict: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semiring_SemiringRecordusd_Dict.get_or_init(||
                                                             &Func1::new(move
                                                                             |x|
                                                                             x.clone()))
    }
    pub fn Data_Semiring_Semiringusd_Dict() -> &dyn Any {
        static Data_Semiring_Semiringusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_Semiringusd_Dict.get_or_init(||
                                                       &Func1::new(move |x|
                                                                       x.clone()))
    }
    pub fn Data_Semiring_zeroRecord() -> &dyn Any {
        static Data_Semiring_zeroRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_zeroRecord.get_or_init(||
                                                 &Func1::new(move |dict|
                                                                 find(string("zeroRecord"),
                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_zero() -> &dyn Any {
        static Data_Semiring_zero: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_zero.get_or_init(||
                                           &Func1::new(move |dict|
                                                           find(string("zero"),
                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_semiringUnit() -> &dyn Any {
        static Data_Semiring_semiringUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_semiringUnit.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                    &&&add(string("add"),
                                                                                           &&Func1::new(move
                                                                                                            |v|
                                                                                                            &Func1::new(move
                                                                                                                            |v1|
                                                                                                                            &PureScript_Data_Unit::Data_Unit_unit())),
                                                                                           add(string("zero"),
                                                                                               &&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                               add(string("mul"),
                                                                                                   &&Func1::new(move
                                                                                                                    |v_1|
                                                                                                                    &Func1::new(move
                                                                                                                                    |v1_1|
                                                                                                                                    &PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                   add(string("one"),
                                                                                                       &&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))))
    }
    pub fn Data_Semiring_semiringRecordNil() -> &dyn Any {
        static Data_Semiring_semiringRecordNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_semiringRecordNil.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_SemiringRecordusd_Dict(),
                                                                                         &&&add(string("addRecord"),
                                                                                                &&Func1::new(move
                                                                                                                 |v|
                                                                                                                 &Func1::new(move
                                                                                                                                 |v1|
                                                                                                                                 &Func1::new(move
                                                                                                                                                 |v2|
                                                                                                                                                 &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                          &dyn Any>()))),
                                                                                                add(string("mulRecord"),
                                                                                                    &&Func1::new(move
                                                                                                                     |v_1|
                                                                                                                     &Func1::new(move
                                                                                                                                     |v1_1|
                                                                                                                                     &Func1::new(move
                                                                                                                                                     |v2_1|
                                                                                                                                                     &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                              &dyn Any>()))),
                                                                                                    add(string("oneRecord"),
                                                                                                        &&Func1::new(move
                                                                                                                         |v_2|
                                                                                                                         &Func1::new(move
                                                                                                                                         |v1_2|
                                                                                                                                         &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                  &dyn Any>())),
                                                                                                        add(string("zeroRecord"),
                                                                                                            &&Func1::new(move
                                                                                                                             |v_3|
                                                                                                                             &Func1::new(move
                                                                                                                                             |v1_3|
                                                                                                                                             &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                      &dyn Any>())),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>()))))))
    }
    pub fn Data_Semiring_semiringProxy() -> &dyn Any {
        static Data_Semiring_semiringProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_semiringProxy.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                     &&&add(string("add"),
                                                                                            &&Func1::new(move
                                                                                                             |v|
                                                                                                             &Func1::new(move
                                                                                                                             |v1|
                                                                                                                             &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                            add(string("mul"),
                                                                                                &&Func1::new(move
                                                                                                                 |v_1|
                                                                                                                 &Func1::new(move
                                                                                                                                 |v1_1|
                                                                                                                                 &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                                add(string("one"),
                                                                                                    &&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor),
                                                                                                    add(string("zero"),
                                                                                                        &&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))))))
    }
    pub fn Data_Semiring_semiringNumber() -> &dyn Any {
        static Data_Semiring_semiringNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_semiringNumber.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                      &&&add(string("add"),
                                                                                             &&PureScript_Data_Semiring::Data_Semiring_numAdd(),
                                                                                             add(string("zero"),
                                                                                                 &&0.0_f64,
                                                                                                 add(string("mul"),
                                                                                                     &&PureScript_Data_Semiring::Data_Semiring_numMul(),
                                                                                                     add(string("one"),
                                                                                                         &&1.0_f64,
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))))
    }
    pub fn Data_Semiring_semiringInt() -> &dyn Any {
        static Data_Semiring_semiringInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_semiringInt.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                   &&&add(string("add"),
                                                                                          &&PureScript_Data_Semiring::Data_Semiring_intAdd(),
                                                                                          add(string("zero"),
                                                                                              &&0_i32,
                                                                                              add(string("mul"),
                                                                                                  &&PureScript_Data_Semiring::Data_Semiring_intMul(),
                                                                                                  add(string("one"),
                                                                                                      &&1_i32,
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))))
    }
    pub fn Data_Semiring_oneRecord() -> &dyn Any {
        static Data_Semiring_oneRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_oneRecord.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("oneRecord"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_one() -> &dyn Any {
        static Data_Semiring_one: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_one.get_or_init(||
                                          &Func1::new(move |dict|
                                                          find(string("one"),
                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_mulRecord() -> &dyn Any {
        static Data_Semiring_mulRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_mulRecord.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("mulRecord"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_mul() -> &dyn Any {
        static Data_Semiring_mul: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_mul.get_or_init(||
                                          &Func1::new(move |dict|
                                                          find(string("mul"),
                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_addRecord() -> &dyn Any {
        static Data_Semiring_addRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_addRecord.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("addRecord"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_semiringRecord() -> &dyn Any {
        static Data_Semiring_semiringRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_semiringRecord.get_or_init(||
                                                     &Func1::new(move
                                                                     |usd__unused|
                                                                     &Func1::new(move
                                                                                     |dictSemiringRecord|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                                                      &&&add(string("add"),
                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_addRecord(),
                                                                                                                                                                                                  dictSemiringRecord),
                                                                                                                                                               &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                             add(string("mul"),
                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mulRecord(),
                                                                                                                                                                                                      dictSemiringRecord),
                                                                                                                                                                   &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                 add(string("one"),
                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_oneRecord(),
                                                                                                                                                                                                                                             dictSemiringRecord),
                                                                                                                                                                                                          &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                       &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                     add(string("zero"),
                                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zeroRecord(),
                                                                                                                                                                                                                                                 dictSemiringRecord),
                                                                                                                                                                                                              &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                           &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                         empty::<string,
                                                                                                                                                 &dyn Any>()))))))))
    }
    pub fn Data_Semiring_add() -> &dyn Any {
        static Data_Semiring_add: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_add.get_or_init(||
                                          &Func1::new(move |dict|
                                                          find(string("add"),
                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semiring_semiringFn() -> &dyn Any {
        static Data_Semiring_semiringFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_semiringFn.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictSemiring|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                                  &&&add(string("add"),
                                                                                                         &&Func1::new({
                                                                                                                          let dictSemiring
                                                                                                                              =
                                                                                                                              dictSemiring.clone();
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
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                             &&&dictSemiring),
                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                             x)),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                          x))
                                                                                                                                                              })
                                                                                                                                          })
                                                                                                                      }),
                                                                                                         add(string("zero"),
                                                                                                             &&Func1::new({
                                                                                                                              let dictSemiring
                                                                                                                                  =
                                                                                                                                  dictSemiring.clone();
                                                                                                                              move
                                                                                                                                  |v|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                                   &&&dictSemiring)
                                                                                                                          }),
                                                                                                             add(string("mul"),
                                                                                                                 &&Func1::new({
                                                                                                                                  let dictSemiring
                                                                                                                                      =
                                                                                                                                      dictSemiring.clone();
                                                                                                                                  move
                                                                                                                                      |f_1|
                                                                                                                                      &Func1::new({
                                                                                                                                                      let f_1
                                                                                                                                                          =
                                                                                                                                                          f_1.clone();
                                                                                                                                                      move
                                                                                                                                                          |g_1|
                                                                                                                                                          &Func1::new({
                                                                                                                                                                          let g_1
                                                                                                                                                                              =
                                                                                                                                                                              g_1.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |x_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                     &&&dictSemiring),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                                                                                     x_1)),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&g_1,
                                                                                                                                                                                                                                                  x_1))
                                                                                                                                                                      })
                                                                                                                                                  })
                                                                                                                              }),
                                                                                                                 add(string("one"),
                                                                                                                     &&Func1::new({
                                                                                                                                      let dictSemiring
                                                                                                                                          =
                                                                                                                                          dictSemiring.clone();
                                                                                                                                      move
                                                                                                                                          |v_1|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                           &&&dictSemiring)
                                                                                                                                  }),
                                                                                                                     empty::<string,
                                                                                                                             &dyn Any>())))))))
    }
    pub fn Data_Semiring_semiringRecordCons() -> &dyn Any {
        static Data_Semiring_semiringRecordCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semiring_semiringRecordCons.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictIsSymbol|
                                                                         &Func1::new({
                                                                                         let dictIsSymbol
                                                                                             =
                                                                                             dictIsSymbol.clone();
                                                                                         move
                                                                                             |usd__unused|
                                                                                             &Func1::new(move
                                                                                                             |dictSemiringRecord|
                                                                                                             &Func1::new({
                                                                                                                             let dictSemiringRecord
                                                                                                                                 =
                                                                                                                                 dictSemiringRecord.clone();
                                                                                                                             move
                                                                                                                                 |dictSemiring|
                                                                                                                                 {
                                                                                                                                     let one1 =
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                          dictSemiring);
                                                                                                                                     let zero1 =
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                                          dictSemiring);
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_SemiringRecordusd_Dict(),
                                                                                                                                                                      &&&add(string("addRecord"),
                                                                                                                                                                             &&Func1::new({
                                                                                                                                                                                              let dictSemiring
                                                                                                                                                                                                  =
                                                                                                                                                                                                  dictSemiring.clone();
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
                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_addRecord(),
                                                                                                                                                                                                                                                                                                                                                                                        &&&dictSemiringRecord),
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
                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&dictSemiring),
                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&get_,
                                                                                                                                                                                                                                                                                                                                                                                                                       &&&ra)),
                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&get_,
                                                                                                                                                                                                                                                                                                                                                                                    rb))),
                                                                                                                                                                                                                                                                           &&&tail)
                                                                                                                                                                                                                                      }
                                                                                                                                                                                                                              }))
                                                                                                                                                                                          }),
                                                                                                                                                                             add(string("mulRecord"),
                                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                                  let dictSemiring
                                                                                                                                                                                                      =
                                                                                                                                                                                                      dictSemiring.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |v_1|
                                                                                                                                                                                                      &Func1::new(move
                                                                                                                                                                                                                      |ra_1|
                                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                                      let ra_1
                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                          ra_1.clone();
                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                          |rb_1|
                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                              let tail_1 =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mulRecord(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&dictSemiringRecord),
                                                                                                                                                                                                                                                                                                                                                         &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                                                      &&&ra_1),
                                                                                                                                                                                                                                                                                   rb_1);
                                                                                                                                                                                                                                              let key_1 =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                      &&&dictIsSymbol),
                                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                              let insert_1 =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                   &&&key_1);
                                                                                                                                                                                                                                              let get__1 =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeGet(),
                                                                                                                                                                                                                                                                                   &&&key_1);
                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&insert_1,
                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&dictSemiring),
                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&get__1,
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&ra_1)),
                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&get__1,
                                                                                                                                                                                                                                                                                                                                                                                        rb_1))),
                                                                                                                                                                                                                                                                               &&&tail_1)
                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                  }))
                                                                                                                                                                                              }),
                                                                                                                                                                                 add(string("oneRecord"),
                                                                                                                                                                                     &&Func1::new({
                                                                                                                                                                                                      let one1
                                                                                                                                                                                                          =
                                                                                                                                                                                                          one1.clone();
                                                                                                                                                                                                      move
                                                                                                                                                                                                          |v_2|
                                                                                                                                                                                                          &Func1::new(move
                                                                                                                                                                                                                          |v1|
                                                                                                                                                                                                                          {
                                                                                                                                                                                                                              let tail_2 =
                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_oneRecord(),
                                                                                                                                                                                                                                                                                                                                         &&&dictSemiringRecord),
                                                                                                                                                                                                                                                                                                      &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                              let key_2 =
                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                      &&&dictIsSymbol),
                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                                                                      &&&key_2),
                                                                                                                                                                                                                                                                                                  &&&one1),
                                                                                                                                                                                                                                                               &&&tail_2)
                                                                                                                                                                                                                          })
                                                                                                                                                                                                  }),
                                                                                                                                                                                     add(string("zeroRecord"),
                                                                                                                                                                                         &&Func1::new({
                                                                                                                                                                                                          let zero1
                                                                                                                                                                                                              =
                                                                                                                                                                                                              zero1.clone();
                                                                                                                                                                                                          move
                                                                                                                                                                                                              |v_3|
                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                              |v1_1|
                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                  let tail_3 =
                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zeroRecord(),
                                                                                                                                                                                                                                                                                                                                             &&&dictSemiringRecord),
                                                                                                                                                                                                                                                                                                          &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                       &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                  let key_3 =
                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                          &&&dictIsSymbol),
                                                                                                                                                                                                                                                                       &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                                                                          &&&key_3),
                                                                                                                                                                                                                                                                                                      &&&zero1),
                                                                                                                                                                                                                                                                   &&&tail_3)
                                                                                                                                                                                                                              })
                                                                                                                                                                                                      }),
                                                                                                                                                                                         empty::<string,
                                                                                                                                                                                                 &dyn Any>())))))
                                                                                                                                 }
                                                                                                                         }))
                                                                                     })))
    }
}
