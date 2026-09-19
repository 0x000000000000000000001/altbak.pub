pub mod PureScript_Data_Bounded {
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
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Data_Bounded_FFI {
        use super::*;
        pub fn topChar() -> char {
            static topChar: MutCell<Option<char>> = MutCell::new(None);
            topChar.get_or_init(|| '￿')
        }
        pub fn bottomChar() -> char {
            static bottomChar: MutCell<Option<char>> = MutCell::new(None);
            bottomChar.get_or_init(|| '\u{0000}')
        }
        pub fn topNumber() -> f64 {
            static topNumber: MutCell<Option<f64>> = MutCell::new(None);
            topNumber.get_or_init(|| 1.7976931348623157E+308_f64)
        }
        pub fn bottomNumber() -> f64 {
            static bottomNumber: MutCell<Option<f64>> = MutCell::new(None);
            bottomNumber.get_or_init(|| -1.7976931348623157E+308_f64)
        }
        pub fn topInt() -> i32 {
            static topInt: MutCell<Option<i32>> = MutCell::new(None);
            topInt.get_or_init(|| i32::MAX)
        }
        pub fn bottomInt() -> i32 {
            static bottomInt: MutCell<Option<i32>> = MutCell::new(None);
            bottomInt.get_or_init(|| i32::MIN)
        }
    }
    pub fn Data_Bounded_bottomChar() -> &dyn Any {
        static Data_Bounded_bottomChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_bottomChar.get_or_init(||
                                                &PureScript_Data_Bounded::Data_Bounded_FFI::bottomChar())
    }
    pub fn Data_Bounded_bottomInt() -> &dyn Any {
        static Data_Bounded_bottomInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_bottomInt.get_or_init(||
                                               &PureScript_Data_Bounded::Data_Bounded_FFI::bottomInt())
    }
    pub fn Data_Bounded_bottomNumber() -> &dyn Any {
        static Data_Bounded_bottomNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_bottomNumber.get_or_init(||
                                                  &PureScript_Data_Bounded::Data_Bounded_FFI::bottomNumber())
    }
    pub fn Data_Bounded_topChar() -> &dyn Any {
        static Data_Bounded_topChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_topChar.get_or_init(||
                                             &PureScript_Data_Bounded::Data_Bounded_FFI::topChar())
    }
    pub fn Data_Bounded_topInt() -> &dyn Any {
        static Data_Bounded_topInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_topInt.get_or_init(||
                                            &PureScript_Data_Bounded::Data_Bounded_FFI::topInt())
    }
    pub fn Data_Bounded_topNumber() -> &dyn Any {
        static Data_Bounded_topNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_topNumber.get_or_init(||
                                               &PureScript_Data_Bounded::Data_Bounded_FFI::topNumber())
    }
    pub fn Data_Bounded_ordRecord() -> &dyn Any {
        static Data_Bounded_ordRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_ordRecord.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordRecord(),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Bounded_BoundedRecordusd_Dict() -> &dyn Any {
        static Data_Bounded_BoundedRecordusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_BoundedRecordusd_Dict.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Data_Bounded_Boundedusd_Dict() -> &dyn Any {
        static Data_Bounded_Boundedusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_Boundedusd_Dict.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Data_Bounded_topRecord() -> &dyn Any {
        static Data_Bounded_topRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_topRecord.get_or_init(||
                                               &Func1::new(move |dict|
                                                               find(string("topRecord"),
                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bounded_top() -> &dyn Any {
        static Data_Bounded_top: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_top.get_or_init(||
                                         &Func1::new(move |dict|
                                                         find(string("top"),
                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bounded_boundedUnit() -> &dyn Any {
        static Data_Bounded_boundedUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedUnit.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                  &&&add(string("top"),
                                                                                         &&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                         add(string("bottom"),
                                                                                             &&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                             add(string("Ord0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Ord::Data_Ord_ordUnit()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>())))))
    }
    pub fn Data_Bounded_boundedRecordNil() -> &dyn Any {
        static Data_Bounded_boundedRecordNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedRecordNil.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_BoundedRecordusd_Dict(),
                                                                                       &&&add(string("topRecord"),
                                                                                              &&Func1::new(move
                                                                                                               |v|
                                                                                                               &Func1::new(move
                                                                                                                               |v1|
                                                                                                                               &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                        &dyn Any>())),
                                                                                              add(string("bottomRecord"),
                                                                                                  &&Func1::new(move
                                                                                                                   |v_1|
                                                                                                                   &Func1::new(move
                                                                                                                                   |v1_1|
                                                                                                                                   &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                            &dyn Any>())),
                                                                                                  add(string("OrdRecord0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Data_Ord::Data_Ord_ordRecordNil()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))))
    }
    pub fn Data_Bounded_boundedProxy() -> &dyn Any {
        static Data_Bounded_boundedProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedProxy.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                   &&&add(string("bottom"),
                                                                                          &&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor),
                                                                                          add(string("top"),
                                                                                              &&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor),
                                                                                              add(string("Ord0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_Ord::Data_Ord_ordProxy()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))))
    }
    pub fn Data_Bounded_boundedOrdering() -> &dyn Any {
        static Data_Bounded_boundedOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedOrdering.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                      &&&add(string("top"),
                                                                                             &&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                                             add(string("bottom"),
                                                                                                 &&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                                                 add(string("Ord0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Data_Ord::Data_Ord_ordOrdering()),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))))
    }
    pub fn Data_Bounded_boundedNumber() -> &dyn Any {
        static Data_Bounded_boundedNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedNumber.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                    &&&add(string("top"),
                                                                                           &&PureScript_Data_Bounded::Data_Bounded_topNumber(),
                                                                                           add(string("bottom"),
                                                                                               &&PureScript_Data_Bounded::Data_Bounded_bottomNumber(),
                                                                                               add(string("Ord0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>())))))
    }
    pub fn Data_Bounded_boundedInt() -> &dyn Any {
        static Data_Bounded_boundedInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedInt.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                 &&&add(string("top"),
                                                                                        &&PureScript_Data_Bounded::Data_Bounded_topInt(),
                                                                                        add(string("bottom"),
                                                                                            &&PureScript_Data_Bounded::Data_Bounded_bottomInt(),
                                                                                            add(string("Ord0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))))
    }
    pub fn Data_Bounded_boundedChar() -> &dyn Any {
        static Data_Bounded_boundedChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedChar.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                  &&&add(string("top"),
                                                                                         &&PureScript_Data_Bounded::Data_Bounded_topChar(),
                                                                                         add(string("bottom"),
                                                                                             &&PureScript_Data_Bounded::Data_Bounded_bottomChar(),
                                                                                             add(string("Ord0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Ord::Data_Ord_ordChar()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>())))))
    }
    pub fn Data_Bounded_boundedBoolean() -> &dyn Any {
        static Data_Bounded_boundedBoolean: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedBoolean.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                     &&&add(string("top"),
                                                                                            &&true,
                                                                                            add(string("bottom"),
                                                                                                &&false,
                                                                                                add(string("Ord0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &PureScript_Data_Ord::Data_Ord_ordBoolean()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))))
    }
    pub fn Data_Bounded_bottomRecord() -> &dyn Any {
        static Data_Bounded_bottomRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_bottomRecord.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("bottomRecord"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bounded_boundedRecord() -> &dyn Any {
        static Data_Bounded_boundedRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedRecord.get_or_init(||
                                                   &Func1::new(move
                                                                   |usd__unused|
                                                                   &Func1::new(move
                                                                                   |dictBoundedRecord|
                                                                                   {
                                                                                       let ordRecord1 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_ordRecord(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("OrdRecord0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(dictBoundedRecord)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                                                        &&&add(string("top"),
                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_topRecord(),
                                                                                                                                                                                                                                       dictBoundedRecord),
                                                                                                                                                                                                    &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                 &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                               add(string("bottom"),
                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottomRecord(),
                                                                                                                                                                                                                                           dictBoundedRecord),
                                                                                                                                                                                                        &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                     &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                   add(string("Ord0"),
                                                                                                                                       &&Func1::new({
                                                                                                                                                        let ordRecord1
                                                                                                                                                            =
                                                                                                                                                            ordRecord1.clone();
                                                                                                                                                        move
                                                                                                                                                            |usd__unused_1|
                                                                                                                                                            &ordRecord1
                                                                                                                                                    }),
                                                                                                                                       empty::<string,
                                                                                                                                               &dyn Any>()))))
                                                                                   })))
    }
    pub fn Data_Bounded_bottom() -> &dyn Any {
        static Data_Bounded_bottom: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_bottom.get_or_init(||
                                            &Func1::new(move |dict|
                                                            find(string("bottom"),
                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Bounded_boundedRecordCons() -> &dyn Any {
        static Data_Bounded_boundedRecordCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Bounded_boundedRecordCons.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictIsSymbol|
                                                                       &Func1::new({
                                                                                       let dictIsSymbol
                                                                                           =
                                                                                           dictIsSymbol.clone();
                                                                                       move
                                                                                           |dictBounded|
                                                                                           {
                                                                                               let top1 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                    dictBounded);
                                                                                               let bottom1 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                    dictBounded);
                                                                                               let Ord0 =
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                           Sharpurs_Prelude::unbox(dictBounded)),
                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                                               &Func1::new({
                                                                                                               let Ord0
                                                                                                                   =
                                                                                                                   Ord0.clone();
                                                                                                               let bottom1
                                                                                                                   =
                                                                                                                   bottom1.clone();
                                                                                                               let top1
                                                                                                                   =
                                                                                                                   top1.clone();
                                                                                                               move
                                                                                                                   |usd__unused|
                                                                                                                   &Func1::new(move
                                                                                                                                   |usd__unused_1|
                                                                                                                                   &Func1::new(move
                                                                                                                                                   |dictBoundedRecord|
                                                                                                                                                   {
                                                                                                                                                       let ordRecordCons =
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_ordRecordCons(),
                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("OrdRecord0"),
                                                                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictBoundedRecord)),
                                                                                                                                                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                               &&&dictIsSymbol),
                                                                                                                                                                                            &&&Ord0);
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_BoundedRecordusd_Dict(),
                                                                                                                                                                                        &&&add(string("topRecord"),
                                                                                                                                                                                               &&Func1::new({
                                                                                                                                                                                                                let dictBoundedRecord
                                                                                                                                                                                                                    =
                                                                                                                                                                                                                    dictBoundedRecord.clone();
                                                                                                                                                                                                                move
                                                                                                                                                                                                                    |v|
                                                                                                                                                                                                                    &Func1::new(move
                                                                                                                                                                                                                                    |rowProxy|
                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                        let tail =
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_topRecord(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictBoundedRecord),
                                                                                                                                                                                                                                                                                                                &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                             rowProxy);
                                                                                                                                                                                                                                        let key =
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                &&&dictIsSymbol),
                                                                                                                                                                                                                                                                             &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                                                                                &&&key),
                                                                                                                                                                                                                                                                                                            &&&top1),
                                                                                                                                                                                                                                                                         &&&tail)
                                                                                                                                                                                                                                    })
                                                                                                                                                                                                            }),
                                                                                                                                                                                               add(string("bottomRecord"),
                                                                                                                                                                                                   &&Func1::new({
                                                                                                                                                                                                                    let dictBoundedRecord
                                                                                                                                                                                                                        =
                                                                                                                                                                                                                        dictBoundedRecord.clone();
                                                                                                                                                                                                                    move
                                                                                                                                                                                                                        |v_1|
                                                                                                                                                                                                                        &Func1::new(move
                                                                                                                                                                                                                                        |rowProxy_1|
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                            let tail_1 =
                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottomRecord(),
                                                                                                                                                                                                                                                                                                                                                       &&&dictBoundedRecord),
                                                                                                                                                                                                                                                                                                                    &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                 rowProxy_1);
                                                                                                                                                                                                                                            let key_1 =
                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                                                                    &&&dictIsSymbol),
                                                                                                                                                                                                                                                                                 &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeSet(),
                                                                                                                                                                                                                                                                                                                                                    &&&key_1),
                                                                                                                                                                                                                                                                                                                &&&bottom1),
                                                                                                                                                                                                                                                                             &&&tail_1)
                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                }),
                                                                                                                                                                                                   add(string("OrdRecord0"),
                                                                                                                                                                                                       &&Func1::new({
                                                                                                                                                                                                                        let ordRecordCons
                                                                                                                                                                                                                            =
                                                                                                                                                                                                                            ordRecordCons.clone();
                                                                                                                                                                                                                        move
                                                                                                                                                                                                                            |usd__unused_2|
                                                                                                                                                                                                                            &ordRecordCons
                                                                                                                                                                                                                    }),
                                                                                                                                                                                                       empty::<string,
                                                                                                                                                                                                               &dyn Any>()))))
                                                                                                                                                   }))
                                                                                                           })
                                                                                           }
                                                                                   })))
    }
}
