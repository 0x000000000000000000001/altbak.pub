pub mod PureScript_Data_Semigroup {
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
    use crate::module_38c1e8e1::PureScript_Data_Void;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Data_Semigroup_FFI {
        use super::*;
        use fable_library_rust::Array_::copyTo;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::NativeArray_::new_init;
        use fable_library_rust::String_::append;
        use fable_library_rust::System::Array;
        pub fn concatString<a: Clone + 'static, b: Clone +
                            'static>(a: a, b: b) -> string {
            append(Sharpurs_Prelude::unbox(&&a), Sharpurs_Prelude::unbox(&&b))
        }
        pub fn concatArray(xs: &dyn Any, ys: &dyn Any) -> &dyn Any {
            let arrX = Sharpurs_Prelude::unbox(xs);
            let arrY = Sharpurs_Prelude::unbox(ys);
            if count(arrX.clone()) == 0_i32 {
                ys.clone()
            } else {
                if count(arrY.clone()) == 0_i32 {
                    xs.clone()
                } else {
                    let res =
                        new_init(&defaultOf(),
                                 count(arrX.clone()) + count(arrY.clone()));
                    copyTo(arrX.clone(), 0_i32, res.clone(), 0_i32,
                           count(arrX.clone()));
                    copyTo(arrY.clone(), 0_i32, res.clone(), count(arrX),
                           count(arrY));
                    &res
                }
            }
        }
    }
    pub fn Data_Semigroup_concatArray() -> &dyn Any {
        static Data_Semigroup_concatArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_concatArray.get_or_init(||
                                                   &Func1::new(move |arg0|
                                                                   &Func1::new({
                                                                                   let arg0
                                                                                       =
                                                                                       arg0.clone();
                                                                                   move
                                                                                       |arg1|
                                                                                       &PureScript_Data_Semigroup::Data_Semigroup_FFI::concatArray(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                   &Sharpurs_Prelude::unbox(arg1))
                                                                               })))
    }
    pub fn Data_Semigroup_concatString() -> &dyn Any {
        static Data_Semigroup_concatString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_concatString.get_or_init(||
                                                    &Func1::new(move |arg0|
                                                                    &Func1::new({
                                                                                    let arg0
                                                                                        =
                                                                                        arg0.clone();
                                                                                    move
                                                                                        |arg1|
                                                                                        &PureScript_Data_Semigroup::Data_Semigroup_FFI::concatString(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                     Sharpurs_Prelude::unbox(arg1))
                                                                                })))
    }
    pub fn Data_Semigroup_SemigroupRecordusd_Dict() -> &dyn Any {
        static Data_Semigroup_SemigroupRecordusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_SemigroupRecordusd_Dict.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               x.clone()))
    }
    pub fn Data_Semigroup_Semigroupusd_Dict() -> &dyn Any {
        static Data_Semigroup_Semigroupusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Semigroupusd_Dict.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Semigroup_semigroupVoid() -> &dyn Any {
        static Data_Semigroup_semigroupVoid: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupVoid.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                      &&&add(string("append"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              &PureScript_Data_Void::Data_Void_absurd()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Semigroup_semigroupUnit() -> &dyn Any {
        static Data_Semigroup_semigroupUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupUnit.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                      &&&add(string("append"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              &Func1::new(move
                                                                                                                              |v1|
                                                                                                                              &PureScript_Data_Unit::Data_Unit_unit())),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Semigroup_semigroupString() -> &dyn Any {
        static Data_Semigroup_semigroupString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupString.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                        &&&add(string("append"),
                                                                                               &&PureScript_Data_Semigroup::Data_Semigroup_concatString(),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))
    }
    pub fn Data_Semigroup_semigroupRecordNil() -> &dyn Any {
        static Data_Semigroup_semigroupRecordNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupRecordNil.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_SemigroupRecordusd_Dict(),
                                                                                           &&&add(string("appendRecord"),
                                                                                                  &&Func1::new(move
                                                                                                                   |v|
                                                                                                                   &Func1::new(move
                                                                                                                                   |v1|
                                                                                                                                   &Func1::new(move
                                                                                                                                                   |v2|
                                                                                                                                                   &empty::<LrcPtr<dyn IComparable>,
                                                                                                                                                            &dyn Any>()))),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Semigroup_semigroupProxy() -> &dyn Any {
        static Data_Semigroup_semigroupProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupProxy.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                       &&&add(string("append"),
                                                                                              &&Func1::new(move
                                                                                                               |v|
                                                                                                               &Func1::new(move
                                                                                                                               |v1|
                                                                                                                               &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Semigroup_semigroupArray() -> &dyn Any {
        static Data_Semigroup_semigroupArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupArray.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                       &&&add(string("append"),
                                                                                              &&PureScript_Data_Semigroup::Data_Semigroup_concatArray(),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Semigroup_appendRecord() -> &dyn Any {
        static Data_Semigroup_appendRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_appendRecord.get_or_init(||
                                                    &Func1::new(move |dict|
                                                                    find(string("appendRecord"),
                                                                         Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semigroup_semigroupRecord() -> &dyn Any {
        static Data_Semigroup_semigroupRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupRecord.get_or_init(||
                                                       &Func1::new(move
                                                                       |usd__unused|
                                                                       &Func1::new(move
                                                                                       |dictSemigroupRecord|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                        &&&add(string("append"),
                                                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_appendRecord(),
                                                                                                                                                                                                    dictSemigroupRecord),
                                                                                                                                                                 &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>())))))
    }
    pub fn Data_Semigroup_append() -> &dyn Any {
        static Data_Semigroup_append: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_append.get_or_init(||
                                              &Func1::new(move |dict|
                                                              find(string("append"),
                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Semigroup_semigroupFn() -> &dyn Any {
        static Data_Semigroup_semigroupFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupFn.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictSemigroup|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                    &&&add(string("append"),
                                                                                                           &&Func1::new({
                                                                                                                            let dictSemigroup
                                                                                                                                =
                                                                                                                                dictSemigroup.clone();
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
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                               &&&dictSemigroup),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                               x)),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&g,
                                                                                                                                                                                                                                            x))
                                                                                                                                                                })
                                                                                                                                            })
                                                                                                                        }),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>()))))
    }
    pub fn Data_Semigroup_semigroupRecordCons() -> &dyn Any {
        static Data_Semigroup_semigroupRecordCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_semigroupRecordCons.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictIsSymbol|
                                                                           &Func1::new({
                                                                                           let dictIsSymbol
                                                                                               =
                                                                                               dictIsSymbol.clone();
                                                                                           move
                                                                                               |usd__unused|
                                                                                               &Func1::new(move
                                                                                                               |dictSemigroupRecord|
                                                                                                               &Func1::new({
                                                                                                                               let dictSemigroupRecord
                                                                                                                                   =
                                                                                                                                   dictSemigroupRecord.clone();
                                                                                                                               move
                                                                                                                                   |dictSemigroup|
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_SemigroupRecordusd_Dict(),
                                                                                                                                                                    &&&add(string("appendRecord"),
                                                                                                                                                                           &&Func1::new({
                                                                                                                                                                                            let dictSemigroup
                                                                                                                                                                                                =
                                                                                                                                                                                                dictSemigroup.clone();
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
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_appendRecord(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&dictSemigroupRecord),
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
                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                     &&&dictSemigroup),
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
}
