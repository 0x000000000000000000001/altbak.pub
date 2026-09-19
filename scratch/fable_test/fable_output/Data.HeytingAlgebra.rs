pub mod PureScript_Data_HeytingAlgebra {
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
    use fable_library_rust::Util_::Lazy;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    use fable_library_rust::System::Lazy_1;
    pub mod Data_HeytingAlgebra_FFI {
        use super::*;
        pub fn boolConj(a: &dyn Any, b: &dyn Any) -> &dyn Any {
            &if Sharpurs_Prelude::unbox(a) {
                 Sharpurs_Prelude::unbox(b)
             } else { false }
        }
        pub fn boolDisj(a: &dyn Any, b: &dyn Any) -> &dyn Any {
            &if Sharpurs_Prelude::unbox(a) {
                 true
             } else { Sharpurs_Prelude::unbox(b) }
        }
        pub fn boolNot(a: &dyn Any) -> &dyn Any {
            &!Sharpurs_Prelude::unbox(a)
        }
    }
    pub fn Data_HeytingAlgebra_boolConj() -> &dyn Any {
        static Data_HeytingAlgebra_boolConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_boolConj.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &Func1::new({
                                                                                     let arg0
                                                                                         =
                                                                                         arg0.clone();
                                                                                     move
                                                                                         |arg1|
                                                                                         &PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_FFI::boolConj(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                            &Sharpurs_Prelude::unbox(arg1))
                                                                                 })))
    }
    pub fn Data_HeytingAlgebra_boolDisj() -> &dyn Any {
        static Data_HeytingAlgebra_boolDisj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_boolDisj.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &Func1::new({
                                                                                     let arg0
                                                                                         =
                                                                                         arg0.clone();
                                                                                     move
                                                                                         |arg1|
                                                                                         &PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_FFI::boolDisj(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                            &Sharpurs_Prelude::unbox(arg1))
                                                                                 })))
    }
    pub fn Data_HeytingAlgebra_boolNot() -> &dyn Any {
        static Data_HeytingAlgebra_boolNot: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_boolNot.get_or_init(||
                                                    &Func1::new(move |arg0|
                                                                    &PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_FFI::boolNot(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Data_HeytingAlgebra_HeytingAlgebraRecordusd_Dict() -> &dyn Any {
        static Data_HeytingAlgebra_HeytingAlgebraRecordusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_HeytingAlgebraRecordusd_Dict.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |x|
                                                                                         x.clone()))
    }
    pub fn Data_HeytingAlgebra_HeytingAlgebrausd_Dict() -> &dyn Any {
        static Data_HeytingAlgebra_HeytingAlgebrausd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_HeytingAlgebrausd_Dict.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |x|
                                                                                   x.clone()))
    }
    pub fn Data_HeytingAlgebra_ttRecord() -> &dyn Any {
        static Data_HeytingAlgebra_ttRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_ttRecord.get_or_init(||
                                                     &Func1::new(move |dict|
                                                                     find(string("ttRecord"),
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_tt() -> &dyn Any {
        static Data_HeytingAlgebra_tt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_tt.get_or_init(||
                                               &Func1::new(move |dict|
                                                               find(string("tt"),
                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_notRecord() -> &dyn Any {
        static Data_HeytingAlgebra_notRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_notRecord.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("notRecord"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_not() -> &dyn Any {
        static Data_HeytingAlgebra_not: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_not.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("not"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_impliesRecord() -> &dyn Any {
        static Data_HeytingAlgebra_impliesRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_impliesRecord.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("impliesRecord"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_implies() -> &dyn Any {
        static Data_HeytingAlgebra_implies: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_implies.get_or_init(||
                                                    &Func1::new(move |dict|
                                                                    find(string("implies"),
                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraUnit() -> &dyn Any {
        static Data_HeytingAlgebra_heytingAlgebraUnit:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_heytingAlgebraUnit.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebrausd_Dict(),
                                                                                                &&&add(string("ff"),
                                                                                                       &&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                                       add(string("tt"),
                                                                                                           &&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                                           add(string("implies"),
                                                                                                               &&Func1::new(move
                                                                                                                                |v|
                                                                                                                                &Func1::new(move
                                                                                                                                                |v1|
                                                                                                                                                &PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                               add(string("conj"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |v_1|
                                                                                                                                    &Func1::new(move
                                                                                                                                                    |v1_1|
                                                                                                                                                    &PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                   add(string("disj"),
                                                                                                                       &&Func1::new(move
                                                                                                                                        |v_2|
                                                                                                                                        &Func1::new(move
                                                                                                                                                        |v1_2|
                                                                                                                                                        &PureScript_Data_Unit::Data_Unit_unit())),
                                                                                                                       add(string("not"),
                                                                                                                           &&Func1::new(move
                                                                                                                                            |v_3|
                                                                                                                                            &PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                           empty::<string,
                                                                                                                                   &dyn Any>()))))))))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraRecordNil() -> &dyn Any {
        static Data_HeytingAlgebra_heytingAlgebraRecordNil:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_heytingAlgebraRecordNil.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebraRecordusd_Dict(),
                                                                                                     &&&add(string("conjRecord"),
                                                                                                            &&Func1::new(move
                                                                                                                             |v|
                                                                                                                             &Func1::new(move
                                                                                                                                             |v1|
                                                                                                                                             &Func1::new(move
                                                                                                                                                             |v2|
                                                                                                                                                             &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                                      &dyn Any>()))),
                                                                                                            add(string("disjRecord"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |v_1|
                                                                                                                                 &Func1::new(move
                                                                                                                                                 |v1_1|
                                                                                                                                                 &Func1::new(move
                                                                                                                                                                 |v2_1|
                                                                                                                                                                 &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                                          &dyn Any>()))),
                                                                                                                add(string("ffRecord"),
                                                                                                                    &&Func1::new(move
                                                                                                                                     |v_2|
                                                                                                                                     &Func1::new(move
                                                                                                                                                     |v1_2|
                                                                                                                                                     &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                              &dyn Any>())),
                                                                                                                    add(string("impliesRecord"),
                                                                                                                        &&Func1::new(move
                                                                                                                                         |v_3|
                                                                                                                                         &Func1::new(move
                                                                                                                                                         |v1_3|
                                                                                                                                                         &Func1::new(move
                                                                                                                                                                         |v2_2|
                                                                                                                                                                         &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                                                  &dyn Any>()))),
                                                                                                                        add(string("notRecord"),
                                                                                                                            &&Func1::new(move
                                                                                                                                             |v_4|
                                                                                                                                             &Func1::new(move
                                                                                                                                                             |v1_4|
                                                                                                                                                             &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                                      &dyn Any>())),
                                                                                                                            add(string("ttRecord"),
                                                                                                                                &&Func1::new(move
                                                                                                                                                 |v_5|
                                                                                                                                                 &Func1::new(move
                                                                                                                                                                 |v1_5|
                                                                                                                                                                 &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                                          &dyn Any>())),
                                                                                                                                empty::<string,
                                                                                                                                        &dyn Any>()))))))))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraProxy() -> &dyn Any {
        static Data_HeytingAlgebra_heytingAlgebraProxy:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_heytingAlgebraProxy.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebrausd_Dict(),
                                                                                                 &&&add(string("conj"),
                                                                                                        &&Func1::new(move
                                                                                                                         |v|
                                                                                                                         &Func1::new(move
                                                                                                                                         |v1|
                                                                                                                                         &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                                        add(string("disj"),
                                                                                                            &&Func1::new(move
                                                                                                                             |v_1|
                                                                                                                             &Func1::new(move
                                                                                                                                             |v1_1|
                                                                                                                                             &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                                            add(string("implies"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |v_2|
                                                                                                                                 &Func1::new(move
                                                                                                                                                 |v1_2|
                                                                                                                                                 &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                                                add(string("ff"),
                                                                                                                    &&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor),
                                                                                                                    add(string("not"),
                                                                                                                        &&Func1::new(move
                                                                                                                                         |v_3|
                                                                                                                                         &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                        add(string("tt"),
                                                                                                                            &&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>()))))))))
    }
    pub fn Data_HeytingAlgebra_ffRecord() -> &dyn Any {
        static Data_HeytingAlgebra_ffRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_ffRecord.get_or_init(||
                                                     &Func1::new(move |dict|
                                                                     find(string("ffRecord"),
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_ff() -> &dyn Any {
        static Data_HeytingAlgebra_ff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_ff.get_or_init(||
                                               &Func1::new(move |dict|
                                                               find(string("ff"),
                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_disjRecord() -> &dyn Any {
        static Data_HeytingAlgebra_disjRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_disjRecord.get_or_init(||
                                                       &Func1::new(move |dict|
                                                                       find(string("disjRecord"),
                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_disj() -> &dyn Any {
        static Data_HeytingAlgebra_disj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_disj.get_or_init(||
                                                 &Func1::new(move |dict|
                                                                 find(string("disj"),
                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraBoolean_004048() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebrausd_Dict(),
                                         &&&add(string("ff"), &&false,
                                                add(string("tt"), &&true,
                                                    add(string("implies"),
                                                        &&Func1::new({
                                                                         let Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1
                                                                             =
                                                                             Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1.clone();
                                                                         move
                                                                             |a|
                                                                             &Func1::new({
                                                                                             let Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1
                                                                                                 =
                                                                                                 Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1.clone();
                                                                                             let a
                                                                                                 =
                                                                                                 a.clone();
                                                                                             move
                                                                                                 |b|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                        &&&Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1.Value),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                           &&&Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1.Value),
                                                                                                                                                                                                        &&&a)),
                                                                                                                                  b)
                                                                                         })
                                                                     }),
                                                        add(string("conj"),
                                                            &&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_boolConj(),
                                                            add(string("disj"),
                                                                &&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_boolDisj(),
                                                                add(string("not"),
                                                                    &&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_boolNot(),
                                                                    empty::<string,
                                                                            &dyn Any>())))))))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1.get_or_init(||
                                                                               Lazy(Data_HeytingAlgebra_heytingAlgebraBoolean_004048.clone()))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraBoolean() -> &dyn Any {
        static Data_HeytingAlgebra_heytingAlgebraBoolean:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_heytingAlgebraBoolean.get_or_init(||
                                                                  Data_HeytingAlgebra_heytingAlgebraBoolean_004048_002d1.Value)
    }
    pub fn Data_HeytingAlgebra_conjRecord() -> &dyn Any {
        static Data_HeytingAlgebra_conjRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_conjRecord.get_or_init(||
                                                       &Func1::new(move |dict|
                                                                       find(string("conjRecord"),
                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraRecord() -> &dyn Any {
        static Data_HeytingAlgebra_heytingAlgebraRecord:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_heytingAlgebraRecord.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |usd__unused|
                                                                                 &Func1::new(move
                                                                                                 |dictHeytingAlgebraRecord|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebrausd_Dict(),
                                                                                                                                  &&&add(string("ff"),
                                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ffRecord(),
                                                                                                                                                                                                                                                 dictHeytingAlgebraRecord),
                                                                                                                                                                                                              &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                           &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                         add(string("tt"),
                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ttRecord(),
                                                                                                                                                                                                                                                     dictHeytingAlgebraRecord),
                                                                                                                                                                                                                  &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                               &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                             add(string("conj"),
                                                                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conjRecord(),
                                                                                                                                                                                                                      dictHeytingAlgebraRecord),
                                                                                                                                                                                   &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                 add(string("disj"),
                                                                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disjRecord(),
                                                                                                                                                                                                                          dictHeytingAlgebraRecord),
                                                                                                                                                                                       &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                     add(string("implies"),
                                                                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_impliesRecord(),
                                                                                                                                                                                                                              dictHeytingAlgebraRecord),
                                                                                                                                                                                           &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                         add(string("not"),
                                                                                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_notRecord(),
                                                                                                                                                                                                                                  dictHeytingAlgebraRecord),
                                                                                                                                                                                               &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                             empty::<string,
                                                                                                                                                                     &dyn Any>()))))))))))
    }
    pub fn Data_HeytingAlgebra_conj() -> &dyn Any {
        static Data_HeytingAlgebra_conj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_conj.get_or_init(||
                                                 &Func1::new(move |dict|
                                                                 find(string("conj"),
                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraFunction() -> &dyn Any {
        static Data_HeytingAlgebra_heytingAlgebraFunction:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_heytingAlgebraFunction.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictHeytingAlgebra|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebrausd_Dict(),
                                                                                                                    &&&add(string("ff"),
                                                                                                                           &&Func1::new({
                                                                                                                                            let dictHeytingAlgebra
                                                                                                                                                =
                                                                                                                                                dictHeytingAlgebra.clone();
                                                                                                                                            move
                                                                                                                                                |v|
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ff(),
                                                                                                                                                                                 &&&dictHeytingAlgebra)
                                                                                                                                        }),
                                                                                                                           add(string("tt"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let dictHeytingAlgebra
                                                                                                                                                    =
                                                                                                                                                    dictHeytingAlgebra.clone();
                                                                                                                                                move
                                                                                                                                                    |v_1|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_tt(),
                                                                                                                                                                                     &&&dictHeytingAlgebra)
                                                                                                                                            }),
                                                                                                                               add(string("implies"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let dictHeytingAlgebra
                                                                                                                                                        =
                                                                                                                                                        dictHeytingAlgebra.clone();
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
                                                                                                                                                                                                |a|
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_implies(),
                                                                                                                                                                                                                                                                                                       &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                       a)),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                                                    a))
                                                                                                                                                                                        })
                                                                                                                                                                    })
                                                                                                                                                }),
                                                                                                                                   add(string("conj"),
                                                                                                                                       &&Func1::new({
                                                                                                                                                        let dictHeytingAlgebra
                                                                                                                                                            =
                                                                                                                                                            dictHeytingAlgebra.clone();
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
                                                                                                                                                                                                    |a_1|
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                           &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f_1,
                                                                                                                                                                                                                                                                                                           a_1)),
                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&g_1,
                                                                                                                                                                                                                                                                        a_1))
                                                                                                                                                                                            })
                                                                                                                                                                        })
                                                                                                                                                    }),
                                                                                                                                       add(string("disj"),
                                                                                                                                           &&Func1::new({
                                                                                                                                                            let dictHeytingAlgebra
                                                                                                                                                                =
                                                                                                                                                                dictHeytingAlgebra.clone();
                                                                                                                                                            move
                                                                                                                                                                |f_2|
                                                                                                                                                                &Func1::new({
                                                                                                                                                                                let f_2
                                                                                                                                                                                    =
                                                                                                                                                                                    f_2.clone();
                                                                                                                                                                                move
                                                                                                                                                                                    |g_2|
                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                    let g_2
                                                                                                                                                                                                        =
                                                                                                                                                                                                        g_2.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |a_2|
                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                                                                                                                               &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&f_2,
                                                                                                                                                                                                                                                                                                               a_2)),
                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&g_2,
                                                                                                                                                                                                                                                                            a_2))
                                                                                                                                                                                                })
                                                                                                                                                                            })
                                                                                                                                                        }),
                                                                                                                                           add(string("not"),
                                                                                                                                               &&Func1::new({
                                                                                                                                                                let dictHeytingAlgebra
                                                                                                                                                                    =
                                                                                                                                                                    dictHeytingAlgebra.clone();
                                                                                                                                                                move
                                                                                                                                                                    |f_3|
                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                    let f_3
                                                                                                                                                                                        =
                                                                                                                                                                                        f_3.clone();
                                                                                                                                                                                    move
                                                                                                                                                                                        |a_3|
                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                            &&&dictHeytingAlgebra),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&f_3,
                                                                                                                                                                                                                                                            a_3))
                                                                                                                                                                                })
                                                                                                                                                            }),
                                                                                                                                               empty::<string,
                                                                                                                                                       &dyn Any>())))))))))
    }
    pub fn Data_HeytingAlgebra_heytingAlgebraRecordCons() -> &dyn Any {
        static Data_HeytingAlgebra_heytingAlgebraRecordCons:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_HeytingAlgebra_heytingAlgebraRecordCons.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictIsSymbol|
                                                                                     &Func1::new({
                                                                                                     let dictIsSymbol
                                                                                                         =
                                                                                                         dictIsSymbol.clone();
                                                                                                     move
                                                                                                         |usd__unused|
                                                                                                         &Func1::new(move
                                                                                                                         |dictHeytingAlgebraRecord|
                                                                                                                         &Func1::new({
                                                                                                                                         let dictHeytingAlgebraRecord
                                                                                                                                             =
                                                                                                                                             dictHeytingAlgebraRecord.clone();
                                                                                                                                         move
                                                                                                                                             |dictHeytingAlgebra|
                                                                                                                                             {
                                                                                                                                                 let ff1 =
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ff(),
                                                                                                                                                                                      dictHeytingAlgebra);
                                                                                                                                                 let tt1 =
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_tt(),
                                                                                                                                                                                      dictHeytingAlgebra);
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebraRecordusd_Dict(),
                                                                                                                                                                                  &&&add(string("conjRecord"),
                                                                                                                                                                                         &&Func1::new({
                                                                                                                                                                                                          let dictHeytingAlgebra
                                                                                                                                                                                                              =
                                                                                                                                                                                                              dictHeytingAlgebra.clone();
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
                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conjRecord(),
                                                                                                                                                                                                                                                                                                                                                                                                    &&&dictHeytingAlgebraRecord),
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
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&get_,
                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&ra)),
                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&get_,
                                                                                                                                                                                                                                                                                                                                                                                                rb))),
                                                                                                                                                                                                                                                                                       &&&tail)
                                                                                                                                                                                                                                                  }
                                                                                                                                                                                                                                          }))
                                                                                                                                                                                                      }),
                                                                                                                                                                                         add(string("disjRecord"),
                                                                                                                                                                                             &&Func1::new({
                                                                                                                                                                                                              let dictHeytingAlgebra
                                                                                                                                                                                                                  =
                                                                                                                                                                                                                  dictHeytingAlgebra.clone();
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
                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disjRecord(),
                                                                                                                                                                                                                                                                                                                                                                                                        &&&dictHeytingAlgebraRecord),
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
                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&get__1,
                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&ra_1)),
                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&get__1,
                                                                                                                                                                                                                                                                                                                                                                                                    rb_1))),
                                                                                                                                                                                                                                                                                           &&&tail_1)
                                                                                                                                                                                                                                                      }
                                                                                                                                                                                                                                              }))
                                                                                                                                                                                                          }),
                                                                                                                                                                                             add(string("impliesRecord"),
                                                                                                                                                                                                 &&Func1::new({
                                                                                                                                                                                                                  let dictHeytingAlgebra
                                                                                                                                                                                                                      =
                                                                                                                                                                                                                      dictHeytingAlgebra.clone();
                                                                                                                                                                                                                  move
                                                                                                                                                                                                                      |v_2|
                                                                                                                                                                                                                      &Func1::new(move
                                                                                                                                                                                                                                      |ra_2|
                                                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                                                      let ra_2
                                                                                                                                                                                                                                                          =
                                                                                                                                                                                                                                                          ra_2.clone();
                                                                                                                                                                                                                                                      move
                                                                                                                                                                                                                                                          |rb_2|
                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                              let tail_2 =
                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_impliesRecord(),
                                                                                                                                                                                                                                                                                                                                                                                                            &&&dictHeytingAlgebraRecord),
                                                                                                                                                                                                                                                                                                                                                                         &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                                                                      &&&ra_2),
                                                                                                                                                                                                                                                                                                   rb_2);
                                                                                                                                                                                                                                                              let key_2 =
                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                                      &&&dictIsSymbol),
                                                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                                              let insert_2 =
                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                                   &&&key_2);
                                                                                                                                                                                                                                                              let get__2 =
                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeGet(),
                                                                                                                                                                                                                                                                                                   &&&key_2);
                                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&insert_2,
                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_implies(),
                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&get__2,
                                                                                                                                                                                                                                                                                                                                                                                                                                           &&&ra_2)),
                                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&get__2,
                                                                                                                                                                                                                                                                                                                                                                                                        rb_2))),
                                                                                                                                                                                                                                                                                               &&&tail_2)
                                                                                                                                                                                                                                                          }
                                                                                                                                                                                                                                                  }))
                                                                                                                                                                                                              }),
                                                                                                                                                                                                 add(string("ffRecord"),
                                                                                                                                                                                                     &&Func1::new({
                                                                                                                                                                                                                      let ff1
                                                                                                                                                                                                                          =
                                                                                                                                                                                                                          ff1.clone();
                                                                                                                                                                                                                      move
                                                                                                                                                                                                                          |v_3|
                                                                                                                                                                                                                          &Func1::new(move
                                                                                                                                                                                                                                          |row|
                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                              let tail_3 =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ffRecord(),
                                                                                                                                                                                                                                                                                                                                                         &&&dictHeytingAlgebraRecord),
                                                                                                                                                                                                                                                                                                                      &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                   row);
                                                                                                                                                                                                                                              let key_3 =
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                      &&&dictIsSymbol),
                                                                                                                                                                                                                                                                                   &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                                                                                      &&&key_3),
                                                                                                                                                                                                                                                                                                                  &&&ff1),
                                                                                                                                                                                                                                                                               &&&tail_3)
                                                                                                                                                                                                                                          })
                                                                                                                                                                                                                  }),
                                                                                                                                                                                                     add(string("notRecord"),
                                                                                                                                                                                                         &&Func1::new({
                                                                                                                                                                                                                          let dictHeytingAlgebra
                                                                                                                                                                                                                              =
                                                                                                                                                                                                                              dictHeytingAlgebra.clone();
                                                                                                                                                                                                                          move
                                                                                                                                                                                                                              |v_4|
                                                                                                                                                                                                                              &Func1::new(move
                                                                                                                                                                                                                                              |row_1|
                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                  let tail_4 =
                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_notRecord(),
                                                                                                                                                                                                                                                                                                                                                             &&&dictHeytingAlgebraRecord),
                                                                                                                                                                                                                                                                                                                          &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                       row_1);
                                                                                                                                                                                                                                                  let key_4 =
                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                          &&&dictIsSymbol),
                                                                                                                                                                                                                                                                                       &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                                  let insert_4 =
                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                       &&&key_4);
                                                                                                                                                                                                                                                  let get__3 =
                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeGet(),
                                                                                                                                                                                                                                                                                       &&&key_4);
                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&insert_4,
                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&get__3,
                                                                                                                                                                                                                                                                                                                                                                                            row_1))),
                                                                                                                                                                                                                                                                                   &&&tail_4)
                                                                                                                                                                                                                                              })
                                                                                                                                                                                                                      }),
                                                                                                                                                                                                         add(string("ttRecord"),
                                                                                                                                                                                                             &&Func1::new({
                                                                                                                                                                                                                              let tt1
                                                                                                                                                                                                                                  =
                                                                                                                                                                                                                                  tt1.clone();
                                                                                                                                                                                                                              move
                                                                                                                                                                                                                                  |v_5|
                                                                                                                                                                                                                                  &Func1::new(move
                                                                                                                                                                                                                                                  |row_2|
                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                      let tail_5 =
                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ttRecord(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictHeytingAlgebraRecord),
                                                                                                                                                                                                                                                                                                                              &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                           row_2);
                                                                                                                                                                                                                                                      let key_5 =
                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                              &&&dictIsSymbol),
                                                                                                                                                                                                                                                                                           &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                                                                                              &&&key_5),
                                                                                                                                                                                                                                                                                                                          &&&tt1),
                                                                                                                                                                                                                                                                                       &&&tail_5)
                                                                                                                                                                                                                                                  })
                                                                                                                                                                                                                          }),
                                                                                                                                                                                                             empty::<string,
                                                                                                                                                                                                                     &dyn Any>())))))))
                                                                                                                                             }
                                                                                                                                     }))
                                                                                                 })))
    }
}
